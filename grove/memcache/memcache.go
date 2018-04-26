package memcache

import (
	"sync"
	"sync/atomic"

	"github.com/apache/incubator-trafficcontrol/grove/cacheobj"
	"github.com/apache/incubator-trafficcontrol/grove/lru"

	"github.com/apache/incubator-trafficcontrol/lib/go-log"
)

// MemCache is a threadsafe memory cache with a soft byte limit, enforced via LRU.
type MemCache struct {
	lru          *lru.LRU                      // threadsafe.
	cache        map[string]*cacheobj.CacheObj // mutexed: MUST NOT access without locking cacheM. TODO test performance of sync.Map
	cacheM       sync.RWMutex                  // TODO test performance of one mutex for lru+cache
	sizeBytes    uint64                        // atomic: MUST NOT access without sync.atomic
	maxSizeBytes uint64                        // constant: MUST NOT be modified after creation
	gcChan       chan<- uint64
}

func New(bytes uint64) *MemCache {
	log.Errorf("MemCache.New: creating cache with %d capacity.", bytes)
	gcChan := make(chan uint64, 1)
	c := &MemCache{
		lru:          lru.NewLRU(),
		cache:        map[string]*cacheobj.CacheObj{},
		maxSizeBytes: bytes,
		gcChan:       gcChan,
	}
	go c.gcManager(gcChan)
	return c
}

func (c *MemCache) Get(key string) (*cacheobj.CacheObj, bool) {
	c.cacheM.RLock()
	obj, ok := c.cache[key]
	if ok {
		c.lru.Add(key, obj.Size) // TODO directly call c.ll.MoveToFront
	}
	c.cacheM.RUnlock()
	return obj, ok
}

func (c *MemCache) Peek(key string) (*cacheobj.CacheObj, bool) {
	c.cacheM.RLock()
	obj, ok := c.cache[key]
	c.cacheM.RUnlock()
	return obj, ok
}

func (c *MemCache) Add(key string, val *cacheobj.CacheObj) bool {
	c.cacheM.Lock()
	c.cache[key] = val
	c.cacheM.Unlock()
	oldSize := c.lru.Add(key, val.Size)
	sizeChange := val.Size - oldSize
	if sizeChange == 0 {
		return false
	}
	newSizeBytes := atomic.AddUint64(&c.sizeBytes, sizeChange)
	if newSizeBytes <= c.maxSizeBytes {
		return false
	}
	c.doGC(newSizeBytes)
	return false // TODO remove eviction from interface; it's unnecessary and expensive
}

func (c *MemCache) Size() uint64 { return atomic.LoadUint64(&c.sizeBytes) }
func (c *MemCache) Close()       {}

// doGC kicks off garbage collection if it isn't already. Does not block.
func (c *MemCache) doGC(cacheSizeBytes uint64) {
	select {
	case c.gcChan <- cacheSizeBytes:
	default: // don't block if GC is already running
	}
}

// gcManager is the garbage collection manager function, designed to be run in a goroutine. Never returns.
func (c *MemCache) gcManager(gcChan <-chan uint64) {
	for cacheSizeBytes := range gcChan {
		c.gc(cacheSizeBytes)
	}
}

// gc executes garbage collection, until the cache size is under the max. This should be called in a singleton manager goroutine, so only one goroutine is ever doing garbage collection at any time.
func (c *MemCache) gc(cacheSizeBytes uint64) {
	for cacheSizeBytes > c.maxSizeBytes {
		log.Debugf("MemCache.gc cacheSizeBytes %+v > c.maxSizeBytes %+v\n", cacheSizeBytes, c.maxSizeBytes)
		key, sizeBytes, exists := c.lru.RemoveOldest() // TODO change lru to use strings
		if !exists {
			// should never happen
			log.Errorf("MemCache.gc sizeBytes %v > %v maxSizeBytes, but LRU is empty!? Setting cache size to 0!\n", cacheSizeBytes, c.maxSizeBytes)
			atomic.StoreUint64(&c.sizeBytes, 0)
			return
		}

		log.Debugf("MemCache.gc deleting key '" + key + "'")
		c.cacheM.Lock()
		delete(c.cache, key)
		c.cacheM.Unlock()

		cacheSizeBytes = atomic.AddUint64(&c.sizeBytes, ^uint64(sizeBytes-1)) // subtract sizeBytes
	}
}

func (c *MemCache) Keys() []string {
	return c.lru.Keys()
}

func (c *MemCache) Capacity() uint64 {
	return c.maxSizeBytes
}