// (c) Jisin0
// Imdb client and configurations.

package imdb

import (
	"time"

	"github.com/Jisin0/filmigo/cache"
)

const (
	baseImdbURL = "https://imdb.com"
)

// ImdbClient type provides all imdb related operations. Use imdb.NewClient to create one.
type ImdbClient struct {
	disabledCaching bool
	cache           *ImdbCache
}

// Options to configure the imdb client's behaviour.
type ImdbClientOpts struct {

	// Set this to true to disable caching results.
	DisableCaching bool
}

const (
	//Default value of cache expiration : 5 hours.
	defaultCacheExpiration = 5 * time.Hour
	//Default file path at which data is stored
	defaultCachePath = "./.imdbcache/"
)

// NewClient returns a new client with given configs.
func NewClient(o ...ImdbClientOpts) *ImdbClient {
	var disableCaching bool

	if len(o) > 0 {
		disableCaching = o[0].DisableCaching
	}

	return &ImdbClient{
		disabledCaching: disableCaching,
		cache:           NewImdbCache(defaultCacheExpiration),
	}
}

// Set DisableCaching to true only if you need to. It's highly unrecommended as data provided by imdb is pretty persistent.
func (c *ImdbClient) SetDisableCaching(b bool) {
	c.disabledCaching = b
}

// Modify the cache duration of imdb data.
//
// - timeout (time.Duration) - Duration after which cached data must expire.
func (c *ImdbClient) SetCacheTimeout(t time.Duration) {
	c.cache = NewImdbCache(t)
}

// Creates a new imdb cache system with given values.
//
// - timeout (time.Duration) - Duration after which cached data must expire.
func NewImdbCache(timeout time.Duration) *ImdbCache {
	return &ImdbCache{
		MovieCache:  cache.NewCache(defaultCachePath+"/movie/", timeout),
		PersonCache: cache.NewCache(defaultCachePath+"/person/", timeout),
		SearchCache: cache.NewCache(defaultCachePath+"/search/", timeout),
	}
}

type ImdbCache struct {
	MovieCache  *cache.Cache
	PersonCache *cache.Cache
	SearchCache *cache.Cache //Only used for advanced search
}
