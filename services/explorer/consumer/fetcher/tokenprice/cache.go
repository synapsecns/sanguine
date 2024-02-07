package tokenprice

import (
	"context"
	"fmt"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"time"
)

// Service provides price data about tokens using either a cache or defillama
// cache keys sare always ${coin gecko id}_${timestamp}.
type Service interface {
	// GetPriceData attempts to get price data from the cache otherwise it requests defillama
	GetPriceData(context.Context, int, string) *float64
}

const cacheSize = 5000

type tokenPriceServiceImpl struct {
	// tokenCache is the cache of the token prices
	tokenPriceCache *lru.TwoQueueCache[string, float64]
}

// NewPriceDataService creates a new token price data service.
func NewPriceDataService() (Service, error) {
	cache, err := lru.New2Q[string, float64](cacheSize)
	if err != nil {
		return nil, fmt.Errorf("could not create token price data cache: %w", err)
	}

	return &tokenPriceServiceImpl{
		tokenPriceCache: cache,
	}, nil
}

func (t *tokenPriceServiceImpl) GetPriceData(parentCtx context.Context, timestamp int, coinGeckoID string) *float64 {
	ctx, cancel := context.WithTimeout(parentCtx, 3*time.Minute)
	defer cancel()
	key := fmt.Sprintf("%s_%d", coinGeckoID, timestamp)

	// truncate key to save requests
	key = key[:len(key)-3]

	if data, ok := t.tokenPriceCache.Get(key); ok {
		return &data
	}
	tokenPrice := fetcher.GetDefiLlamaData(ctx, timestamp, coinGeckoID)
	if tokenPrice == nil {
		return nil
	}
	t.tokenPriceCache.Add(key, *tokenPrice)

	return tokenPrice
}
