package tc

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// CRConfig is JSON-serializable as the CRConfig used by Traffic Control.
type CRConfig struct {
	// Config is mostly a map of string values, but may contain an 'soa' key which is a map[string]string, and may contain a 'ttls' key with a value map[string]string. It might not contain these values, so they must be checked for, and all values must be checked by the user and an error returned if the type is unexpected. Be aware, neither the language nor the API provides any guarantees about the type!
	Config           map[string]interface{}               `json:"config,omitempty"`
	ContentServers   map[string]CRConfigTrafficOpsServer  `json:"contentServers,omitempty"`
	ContentRouters   map[string]CRConfigRouter            `json:"contentRouters,omitempty"`
	DeliveryServices map[string]CRConfigDeliveryService   `json:"deliveryServices,omitempty"`
	EdgeLocations    map[string]CRConfigLatitudeLongitude `json:"edgeLocations,omitempty"`
	RouterLocations  map[string]CRConfigLatitudeLongitude `json:"trafficRouterLocations,omitempty"`
	Monitors         map[string]CRConfigMonitor           `json:"monitors,omitempty"`
	Stats            CRConfigStats                        `json:"stats,omitempty"`
}

type CRConfigConfig struct {
	APICacheControlMaxAge                      *string      `json:"api.cache-control.max_age,omitempty"`
	ConsistentDNSRouting                       *string      `json:"consistent.dns.routing,omitempty"`
	CoverageZonePollingIntervalSeconds         *string      `json:"coveragezone.polling.interval,omitempty"`
	CoverageZonePollingURL                     *string      `json:"coveragezone.polling.url,omitempty"`
	DNSSecDynamicResponseExpiration            *string      `json:"dnssec.dynamic.response.expiration,omitempty"`
	DNSSecEnabled                              *string      `json:"dnssec.enabled,omitempty"`
	DomainName                                 *string      `json:"domain_name,omitempty"`
	FederationMappingPollingIntervalSeconds    *string      `json:"federationmapping.polling.interval,omitempty"`
	FederationMappingPollingURL                *string      `json:"federationmapping.polling.url"`
	GeoLocationPollingInterval                 *string      `json:"geolocation.polling.interval,omitempty"`
	GeoLocationPollingURL                      *string      `json:"geolocation.polling.url,omitempty"`
	KeyStoreMaintenanceIntervalSeconds         *string      `json:"keystore.maintenance.interval,omitempty"`
	NeustarPollingIntervalSeconds              *string      `json:"neustar.polling.interval,omitempty"`
	NeustarPollingURL                          *string      `json:"neustar.polling.url,omitempty"`
	SOA                                        *SOA         `json:"soa,omitempty"`
	DNSSecInceptionSeconds                     *string      `json:"dnssec.inception,omitempty"`
	Ttls                                       *CRConfigTTL `json:"ttls,omitempty"`
	Weight                                     *string      `json:"weight,omitempty"`
	ZoneManagerCacheMaintenanceIntervalSeconds *string      `json:"zonemanager.cache.maintenance.interval,omitempty"`
	ZoneManagerThreadpoolScale                 *string      `json:"zonemanager.threadpool.scale,omitempty"`
}

type CRConfigTTL struct {
	ASeconds      *string `json:"A,omitempty"`
	AAAASeconds   *string `json:"AAAA,omitempty"`
	DNSkeySeconds *string `json:"DNSKEY,omitempty"`
	DSSeconds     *string `json:"DS,omitempty"`
	NSSeconds     *string `json:"NS,omitempty"`
	SOASeconds    *string `json:"SOA,omitempty"`
}

type CRConfigRouterStatus string

type CRConfigRouter struct {
	APIPort      *string               `json:"api.port,omitempty"`
	FQDN         *string               `json:"fqdn,omitempty"`
	HTTPSPort    *int                  `json:"httpsPort"`
	HashCount    *int                  `json:"hashCount,omitempty"`
	IP           *string               `json:"ip,omitempty"`
	IP6          *string               `json:"ip6,omitempty"`
	Location     *string               `json:"location,omitempty"`
	Port         *int                  `json:"port,omitempty"`
	Profile      *string               `json:"profile,omitempty"`
	ServerStatus *CRConfigRouterStatus `json:"status,omitempty"`
}

type CRConfigServerStatus string

type CRConfigTrafficOpsServer struct {
	CacheGroup       *string               `json:"cacheGroup,omitempty"`
	Fqdn             *string               `json:"fqdn,omitempty"`
	HashCount        *int                  `json:"hashCount,omitempty"`
	HashId           *string               `json:"hashId,omitempty"`
	HttpsPort        *int                  `json:"httpsPort"`
	InterfaceName    *string               `json:"interfaceName,omitempty"`
	Ip               *string               `json:"ip,omitempty"`
	Ip6              *string               `json:"ip6,omitempty"`
	LocationId       *string               `json:"locationId,omitempty"`
	Port             *int                  `json:"port"`
	Profile          *string               `json:"profile,omitempty"`
	ServerStatus     *CRConfigServerStatus `json:"status,omitempty"`
	ServerType       *string               `json:"type,omitempty"`
	DeliveryServices map[string][]string   `json:"deliveryServices,omitempty"`
	RoutingDisabled  int64                 `json:"routingDisabled"`
}

//TODO: drichardson - reconcile this with the DeliveryService struct in deliveryservices.go
type CRConfigDeliveryService struct {
	AnonymousBlockingEnabled *string                               `json:"anonymousBlockingEnabled,omitempty"`
	CoverageZoneOnly         bool                                  `json:"coverageZoneOnly,string"`
	Dispersion               *CRConfigDispersion                   `json:"dispersion,omitempty"`
	Domains                  []string                              `json:"domains,omitempty"`
	GeoLocationProvider      *string                               `json:"geolocationProvider,omitempty"`
	MatchSets                []*MatchSet                           `json:"matchsets,omitempty"`
	MissLocation             *CRConfigLatitudeLongitudeShort       `json:"missLocation,omitempty"`
	Protocol                 *CRConfigDeliveryServiceProtocol      `json:"protocol,omitempty"`
	RegionalGeoBlocking      *string                               `json:"regionalGeoBlocking,omitempty"`
	ResponseHeaders          map[string]string                     `json:"responseHeaders,omitempty"`
	RequestHeaders           []string                              `json:"requestHeaders,omitempty"`
	Soa                      *SOA                                  `json:"soa,omitempty"`
	SSLEnabled               bool                                  `json:"sslEnabled,string"`
	TTL                      *int                                  `json:"ttl,omitempty"`
	TTLs                     *CRConfigTTL                          `json:"ttls,omitempty"`
	MaxDNSIPsForLocation     *int                                  `json:"maxDnsIpsForLocation,omitempty"`
	IP6RoutingEnabled        *bool                                 `json:"ip6RoutingEnabled,string,omitempty"`
	RoutingName              *string                               `json:"routingName,omitempty"`
	BypassDestination        map[string]*CRConfigBypassDestination `json:"bypassDestination,omitempty"`
	DeepCachingType          *DeepCachingType                      `json:"deepCachingType"`
	GeoEnabled               []CRConfigGeoEnabled                  `json:"geoEnabled,omitempty"`
	GeoLimitRedirectURL      *string                               `json:"geoLimitRedirectURL,omitempty"`
	StaticDNSEntries         []StaticDNSEntry                      `json:"staticDnsEntries,omitempty"`
}

type CRConfigGeoEnabled struct {
	CountryCode string `json:"countryCode"`
}

type StaticDNSEntry struct {
	Name  string `json:"name"`
	TTL   int    `json:"ttl"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CRConfigBypassDestination struct {
	IP    *string `json:"ip,omitempty"`    // only used by DNS DSes
	IP6   *string `json:"ip6,omitempty"`   // only used by DNS DSes
	CName *string `json:"cname,omitempty"` // only used by DNS DSes
	TTL   *int    `json:"ttl,omitempty"`   // only used by DNS DSes
	FQDN  *string `json:"fqdn,omitempty"`  // only used by HTTP DSes
	Port  *string `json:"port,omitempty"`  // only used by HTTP DSes
}

type CRConfigDispersion struct {
	Limit    int  `json:"limit"`
	Shuffled bool `json:"shuffled,string"`
}

type CRConfigLatitudeLongitude struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

type CRConfigLatitudeLongitudeShort struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"long"`
}

type CRConfigDeliveryServiceProtocol struct {
	AcceptHTTP      *bool `json:"acceptHttp,string,omitempty"`
	AcceptHTTPS     bool  `json:"acceptHttps,string"`
	RedirectOnHTTPS bool  `json:"redirectToHttps,string"`
}

type CRConfigMonitor struct {
	FQDN         *string               `json:"fqdn,omitempty"`
	HTTPSPort    *int                  `json:"httpsPort"`
	IP           *string               `json:"ip,omitempty"`
	IP6          *string               `json:"ip6,omitempty"`
	Location     *string               `json:"location,omitempty"`
	Port         *int                  `json:"port,omitempty"`
	Profile      *string               `json:"profile,omitempty"`
	ServerStatus *CRConfigServerStatus `json:"status,omitempty"`
}

type CRConfigStats struct {
	CDNName         *string `json:"CDN_name,omitempty"`
	DateUnixSeconds *int64  `json:"date,omitempty"`
	TMHost          *string `json:"tm_host,omitempty"`
	TMPath          *string `json:"tm_path,omitempty"`
	TMUser          *string `json:"tm_user,omitempty"`
	TMVersion       *string `json:"tm_version,omitempty"`
}
