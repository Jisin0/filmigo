// (c) Jisin0
// Base Client and methods.

/*
Package omdb is an api wrapper of the single rout Open Movie DataBase (omdapi.com).
The package offers user friendly interfaces to send queries and helper methods to make working with results easier.
*/
package omdb

import (
	"time"

	"github.com/Jisin0/filmigo/cache"
)

// OmdbClient type provides all omdb related operations. Use omdb.NewClient to create one.
type OmdbClient struct {
	disabledCaching bool
	cache           *cache.Cache
	apiKey          string
}

// Options to configure the imdb client's behaviour.
type OmdbClientOpts struct {
	// Set this to true to disable caching results.
	DisableCaching bool
	// This field is the duration for which cached data is considered valid.
	// Defaluts to 5 * time.Hour.
	CacheExpiration time.Duration
}

const (
	// Default value of cache expiration : 5 hours.
	defaultCacheExpiration = 5 * time.Hour
	// Default file path at which data is stored
	defaultCachePath  = "./.omdbcache/"
	statusCodeSuccess = 200
)

// NewClient returns a new client with given configs.
func NewClient(apiKey string, o ...OmdbClientOpts) *OmdbClient {
	var (
		disableCaching bool
		cacheEpiration = defaultCacheExpiration
	)

	if len(o) > 0 {
		disableCaching = o[0].DisableCaching

		if o[0].CacheExpiration > 0 {
			cacheEpiration = o[0].CacheExpiration
		}
	}

	return &OmdbClient{
		disabledCaching: disableCaching,
		cache:           cache.NewCache(defaultCachePath, cacheEpiration),
		apiKey:          apiKey,
	}
}

// Set DisableCaching to true only if you need to. It's highly unrecommended as data provided by imdb is pretty persistent.
func (c *OmdbClient) SetDisableCaching(b bool) {
	c.disabledCaching = b
}

// Modify the cache duration of imdb data.
//
// - timeout (time.Duration) - Duration after which cached data must expire.
func (c *OmdbClient) SetCacheTimeout(t time.Duration) {
	c.cache = cache.NewCache(defaultCachePath, t)
}
