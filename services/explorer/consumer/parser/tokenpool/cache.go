package tokenpool

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/db"
)

// Service provides data about tokens using either a cache or bridgeconfig
// cache keys sare always ${KEY_NAME}_CHAIN_ID_ADDRESS so unless a token changes tokenID's
// (not the other way around), data is guaranteed to be accurate.
type Service interface {
	// GetTokenAddress attempts to get token data from the cache otherwise its fetched from the bridge config
	GetTokenAddress(ctx context.Context, chainID uint32, tokenIndex uint8, contractAddress string) (*common.Address, error)
}

const cacheSize = 3000

// maxAttemptTime is how many times we will attempt to get the token data.
const maxAttemptTime = time.Second * 120
const maxAttempt = 60

type tokenPoolDataServiceImpl struct {
	consumerDB db.ConsumerDB
	// tokenCache is the tokenCache of the tokenDataServices
	poolTokenCache *lru.TwoQueueCache[string, common.Address]
	// fetcher is the fetcher used to fetch data from the bridge config contract
	service fetcher.SwapService
}

// NewPoolTokenDataService creates a new token data service.
func NewPoolTokenDataService(service fetcher.SwapService, consumerDB db.ConsumerDB) (Service, error) {
	cache, err := lru.New2Q[string, common.Address](cacheSize)
	if err != nil {
		return nil, fmt.Errorf("could not create token data cache: %w", err)
	}

	return &tokenPoolDataServiceImpl{
		consumerDB:     consumerDB,
		poolTokenCache: cache,
		service:        service,
	}, nil
}

func (t *tokenPoolDataServiceImpl) GetTokenAddress(parentCtx context.Context, chainID uint32, tokenIndex uint8, contractAddress string) (*common.Address, error) {
	key := fmt.Sprintf("token_%d_%d", chainID, tokenIndex)
	if data, ok := t.poolTokenCache.Get(key); ok {
		return &data, nil
	}
	var tokenAddress *common.Address
	ctx, cancel := context.WithTimeout(parentCtx, maxAttemptTime)
	defer cancel()

	//nolint: wrapcheck
	err := retry.WithBackoff(ctx, func(ctx context.Context) error {
		var err error
		tokenAddress, err = t.service.GetTokenAddress(ctx, tokenIndex)
		if err != nil {
			return fmt.Errorf("could not get token data for index %d on chain %d: %w", chainID, tokenIndex, err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not get token data with retry backoff chainID %d, tokenIndex %d, contractAddress %s: %w", chainID, tokenIndex, contractAddress, err)
	}

	//nolint: wrapcheck
	err = retry.WithBackoff(ctx, func(ctx context.Context) error {
		return t.storeTokenIndex(ctx, chainID, tokenIndex, tokenAddress, contractAddress)
	}, retry.WithMaxAttemptTime(maxAttemptTime), retry.WithMaxAttempts(maxAttempt))
	if err != nil {
		return nil, fmt.Errorf("could not store token index: %w", err)
	}
	t.poolTokenCache.Add(key, *tokenAddress)

	return tokenAddress, nil
}

func (t *tokenPoolDataServiceImpl) storeTokenIndex(parentCtx context.Context, chainID uint32, tokenIndex uint8, tokenAddress *common.Address, contractAddress string) error {
	err := t.consumerDB.StoreTokenIndex(parentCtx, chainID, tokenIndex, tokenAddress.String(), contractAddress)
	if err != nil {
		return fmt.Errorf("could not store token index: %w", err)
	}
	return nil
}
