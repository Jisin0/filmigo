// (c) Jisin0
// Imdb client and configurations.

/*
The imdb package is a collection of methods to get data from imdb using it's official api or webscraping.
The only publicly exposed imdb has is for searching hence every other method relies on webscraping.
The AdvancedSearch methods do not use an api either.
*/
package imdb

import (
	"regexp"
	"time"

	"github.com/Jisin0/filmigo/internal/cache"
)

const (
	baseImdbURL = "https://imdb.com"

	statusCodeNotFound = 404
	statusCodeSuccess  = 200
)

var (
	resultTypeTitleRegex = regexp.MustCompile(`^tt\d+`)
	resultTypeNameRegex  = regexp.MustCompile(`^nm\d+`)
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
	// This field is the duration for which cached data is considered valid.
	// Defaluts to 5 * time.Hour.
	CacheExpiration time.Duration
}

const (
	// Default value of cache expiration : 5 hours.
	defaultCacheExpiration = 5 * time.Hour
	// Default file path at which data is stored
	defaultCachePath = "./.imdbcache/"
)

// NewClient returns a new client with given configs.
func NewClient(o ...ImdbClientOpts) *ImdbClient {
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

	return &ImdbClient{
		disabledCaching: disableCaching,
		cache:           NewImdbCache(cacheEpiration),
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
	SearchCache *cache.Cache // Only used for advanced search
}
