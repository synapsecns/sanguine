package tokenpool

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"time"
)

// Service provides data about tokens using either a cache or bridgeconfig
// cache keys sare always ${KEY_NAME}_CHAIN_ID_ADDRESS so unless a token changes tokenID's
// (not the other way around), data is guaranteed to be accurate.
type Service interface {
	// GetTokenAddress attempts to get token data from the cache otherwise its fetched from the bridge config
	GetTokenAddress(ctx context.Context, chainID uint32, tokenIndex uint8) (*common.Address, error)
}

const cacheSize = 3000

type tokenPoolDataServiceImpl struct {
	// tokenCache is the tokenCache of the tokenDataServices
	poolTokenCache *lru.TwoQueueCache[string, common.Address]
	// fetcher is the fetcher used to fetch data from the bridge config contract
	service fetcher.SwapService
}

// NewPoolTokenDataService creates a new token data service.
func NewPoolTokenDataService(service fetcher.SwapService) (Service, error) {
	cache, err := lru.New2Q[string, common.Address](cacheSize)
	if err != nil {
		return nil, fmt.Errorf("could not create token data cache: %w", err)
	}

	return &tokenPoolDataServiceImpl{
		poolTokenCache: cache,
		service:        service,
	}, nil
}

func (t *tokenPoolDataServiceImpl) GetTokenAddress(parentCtx context.Context, chainID uint32, tokenIndex uint8) (*common.Address, error) {
	key := fmt.Sprintf("token_%d_%d", chainID, tokenIndex)
	if data, ok := t.poolTokenCache.Get(key); ok {
		return &data, nil
	}
	var tokenAddress *common.Address
	ctx, cancel := context.WithTimeout(parentCtx, maxAttemptTime)
	defer cancel()
	err := t.retryWithBackoff(ctx, func(ctx context.Context) error {
		var err error
		tokenAddress, err = t.service.GetTokenAddress(ctx, tokenIndex)
		if err != nil {
			return fmt.Errorf("could not get token data: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not get token data with retry backoff: %w", err)
	}
	t.poolTokenCache.Add(key, *tokenAddress)

	return tokenAddress, nil
}

// maxAttemptTime is how many times we will attempt to get the token data.
var maxAttemptTime = time.Second * 10

type retryableFunc func(ctx context.Context) error

// retryWithBackoff will retry to get data with a backoff.
func (t *tokenPoolDataServiceImpl) retryWithBackoff(ctx context.Context, doFunc retryableFunc) error {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    2 * time.Millisecond,
		Max:    5 * time.Second,
	}

	timeout := time.Duration(0)

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("%w while retrying", ctx.Err())
		case <-time.After(timeout):
			err := doFunc(ctx)
			if err != nil {
				timeout = b.Duration()
			} else {
				return nil
			}
		}
	}
}
