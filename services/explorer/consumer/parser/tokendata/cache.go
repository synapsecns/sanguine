package tokendata

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"golang.org/x/sync/errgroup"
	"math/big"
	"time"
)

// Service provides data about tokens using either a cache or bridgeconfig
// cache keys sare always ${KEY_NAME}_CHAIN_ID_ADDRESS so unless a token changes tokenID's
// (not the other way around), data is guaranteed to be accurate.
type Service interface {
	// GetTokenData attempts to get token data from the cache otherwise its fetched from the bridge config
	GetTokenData(ctx context.Context, chainID uint32, token common.Address) (ImmutableTokenData, error)
}

const cacheSize = 3000

type tokenDataServiceImpl struct {
	// tokenCache is the tokenCache of the tokenDataServices
	tokenCache *lru.TwoQueueCache[string, ImmutableTokenData]
	// fetcher is the fetcher used to fetch data from the bridge config contract
	service fetcher.Service
}

// NewTokenDataService creates a new token data service.
func NewTokenDataService(service fetcher.Service) (Service, error) {
	cache, err := lru.New2Q[string, ImmutableTokenData](cacheSize)
	if err != nil {
		return nil, fmt.Errorf("could not create token data cache: %w", err)
	}

	return &tokenDataServiceImpl{
		tokenCache: cache,
		service:    service,
	}, nil
}

func (t *tokenDataServiceImpl) GetTokenData(ctx context.Context, chainID uint32, token common.Address) (ImmutableTokenData, error) {
	key := fmt.Sprintf("token_%d_%s", chainID, token.Hex())
	if data, ok := t.tokenCache.Get(key); ok {
		return data, nil
	}

	tokenData, err := t.retrieveTokenData(ctx, chainID, token)
	if err != nil {
		return nil, fmt.Errorf("could not get token data: %w", err)
	}

	t.tokenCache.Add(key, tokenData)

	return tokenData, nil
}

// maxAttemptTime is how many times we will attempt to get the token data.
var maxAttemptTime = time.Second * 10

// retrieveTokenData retrieves the token data from the bridge config contract
// this will retry for maxAttemptTime.
func (t *tokenDataServiceImpl) retrieveTokenData(parentCtx context.Context, chainID uint32, token common.Address) (ImmutableTokenData, error) {
	res := immutableTokenImpl{}

	ctx, cancel := context.WithTimeout(parentCtx, maxAttemptTime)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return t.retryWithBackoff(ctx, func(ctx context.Context) error {
			tokenData, err := t.service.GetToken(ctx, chainID, token)
			if err != nil {
				return fmt.Errorf("could not get token data: %w", err)
			}

			res.decimals = tokenData.TokenDecimals

			return nil
		})
	})

	g.Go(func() error {
		return t.retryWithBackoff(ctx, func(ctx context.Context) error {
			nullableTokenID, err := t.service.GetTokenID(ctx, big.NewInt(int64(chainID)), token)
			if err != nil {
				return fmt.Errorf("could not get token data: %w", err)
			}

			res.tokenID = *nullableTokenID

			return nil
		})
	})

	err := g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not get token data: %w", err)
	}

	return res, nil
}

type retryableFunc func(ctx context.Context) error

// retryWithBackoff will retry to get data with a backoff.
func (t *tokenDataServiceImpl) retryWithBackoff(ctx context.Context, doFunc retryableFunc) error {
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
