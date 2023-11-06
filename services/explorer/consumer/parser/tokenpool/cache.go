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

// Service provides data about tokens using either a cache or bridgeconfig.
type Service interface {
	GetTokenAddress(ctx context.Context, chainID uint32, tokenIndex uint8, contractAddress string) (*common.Address, error)
}

const (
	cacheSize       = 3000
	maxAttemptTime  = time.Second * 120
	maxAttemptCount = 60
)

type tokenPoolService struct {
	db      db.ConsumerDB
	cache   *lru.TwoQueueCache[string, common.Address]
	fetcher fetcher.SwapService
}

func NewService(fetcher fetcher.SwapService, db db.ConsumerDB) (Service, error) {
	cache, err := lru.New2Q[string, common.Address](cacheSize)
	if err != nil {
		return nil, fmt.Errorf("could not create cache: %w", err)
	}
	return &tokenPoolService{
		db:      db,
		cache:   cache,
		fetcher: fetcher,
	}, nil
}

func (s *tokenPoolService) GetTokenAddress(ctx context.Context, chainID uint32, tokenIndex uint8, contractAddress string) (*common.Address, error) {
	key := s.cacheKey(chainID, tokenIndex)
	if address, ok := s.cache.Get(key); ok {
		return &address, nil
	}

	address, err := s.fetchTokenAddress(ctx, tokenIndex)
	if err != nil {
		return nil, fmt.Errorf("error fetching token address: %w", err)
	}

	if err := s.saveTokenIndex(ctx, chainID, tokenIndex, address, contractAddress); err != nil {
		return nil, fmt.Errorf("error saving token index: %w", err)
	}

	s.cache.Add(key, *address)
	return address, nil
}

func (s *tokenPoolService) fetchTokenAddress(ctx context.Context, tokenIndex uint8) (*common.Address, error) {
	var address *common.Address
	err := retry.WithBackoff(ctx, func(ctx context.Context) error {
		var err error
		address, err = s.fetcher.GetTokenAddress(ctx, tokenIndex)
		return fmt.Errorf("error fetching token address with index %d, err: %w", tokenIndex, err)
	}, retry.WithMaxAttemptTime(maxAttemptTime), retry.WithMaxAttempts(maxAttemptCount))

	return address, fmt.Errorf("token fetch by index: %w", err)
}

func (s *tokenPoolService) saveTokenIndex(ctx context.Context, chainID uint32, tokenIndex uint8, address *common.Address, contractAddress string) error {
	err := s.db.StoreTokenIndex(ctx, chainID, tokenIndex, address.String(), contractAddress)
	if err != nil {
		return fmt.Errorf("error saving token index: %w", err)
	}
	return nil
}

func (s *tokenPoolService) cacheKey(chainID uint32, tokenIndex uint8) string {
	return fmt.Sprintf("token_%d_%d", chainID, tokenIndex)
}
