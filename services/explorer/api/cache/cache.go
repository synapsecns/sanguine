package cache

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru/v2"
)

// Service provides price data about tokens using either a cache or defillama
// cache keys sare always ${coin gecko id}_${timestamp}.
type Service interface {
	// CacheResponse saves a response to cache .
	CacheResponse(string, any) error
	// GetCache attempts to get api response data from the cache.
	GetCache(string) any
}

const cacheSize = 100

type apiCacheServiceImpl struct {
	// responseCache is the cache of the api responses
	responseCache *lru.TwoQueueCache[string, any]
}

// NewAPICacheService creates a new api response data service.
func NewAPICacheService() (Service, error) {
	cache, err := lru.New2Q[string, any](cacheSize)
	if err != nil {
		return nil, fmt.Errorf("could not create api response data service: %w", err)
	}
	return &apiCacheServiceImpl{
		responseCache: cache,
	}, nil
}

func (t *apiCacheServiceImpl) CacheResponse(callID string, data any) error {
	t.responseCache.Add(callID, data)
	fmt.Println("added cache data", callID)
	return nil
}

func (t *apiCacheServiceImpl) GetCache(callID string) any {
	if data, ok := t.responseCache.Get(callID); ok {
		fmt.Println("getting cache data")
		return &data
	}
	return nil
}
