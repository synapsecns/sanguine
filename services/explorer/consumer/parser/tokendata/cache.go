package tokendata

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/synapsecns/sanguine/services/explorer/contracts/erc20"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"golang.org/x/sync/errgroup"
)

// Service provides data about tokens using either a cache or bridgeconfig
// cache keys sare always ${KEY_NAME}_CHAIN_ID_ADDRESS so unless a token changes tokenID's
// (not the other way around), data is guaranteed to be accurate.
//
//go:generate go run github.com/vektra/mockery/v2 --name=Service --output=mocks --case=underscore
type Service interface {
	// GetBridgeTokenData attempts to get token data from the cache otherwise its fetched from the bridge config
	GetBridgeTokenData(ctx context.Context, chainID uint32, token common.Address) (ImmutableTokenData, error)
	// GetPoolTokenData attempts to get pool token data from the cache otherwise its fetched from the erc20 interface
	GetPoolTokenData(ctx context.Context, chainID uint32, token common.Address, swapService fetcher.SwapService) (ImmutableTokenData, error)
	// GetCCTPTokenData attempts to get the token symbol from the cctp contract
	GetCCTPTokenData(ctx context.Context, chainID uint32, token common.Address, backend bind.ContractBackend) (ImmutableTokenData, error)
}

const (
	cacheSize      = 3000
	maxAttemptTime = time.Minute * 5
	maxAttempts    = 10
)

type tokenDataServiceImpl struct {
	// tokenCache is the tokenCache of the tokenDataServices
	tokenCache *lru.TwoQueueCache[string, ImmutableTokenData]
	// fetcher is the fetcher used to fetch data from the bridge config contract
	service fetcher.Service
	// tokenSymbolToIDs is a mapping of token symbols to token IDs.
	tokenSymbolToIDs map[string]string
}

// NewTokenDataService creates a new token data service.
func NewTokenDataService(service fetcher.Service, tokenSymbolToIDs map[string]string) (Service, error) {
	cache, err := lru.New2Q[string, ImmutableTokenData](cacheSize)
	if err != nil {
		return nil, fmt.Errorf("could not create token data cache: %w", err)
	}

	return &tokenDataServiceImpl{
		tokenCache:       cache,
		service:          service,
		tokenSymbolToIDs: tokenSymbolToIDs,
	}, nil
}

// GetBridgeTokenData attempts to get token data from the cache otherwise it is fetched from the bridge config.
func (t *tokenDataServiceImpl) GetBridgeTokenData(ctx context.Context, chainID uint32, token common.Address) (ImmutableTokenData, error) {
	key := cacheKey(chainID, token)
	if data, ok := t.tokenCache.Get(key); ok {
		return data, nil
	}

	tokenData, err := t.fetchAndCacheTokenData(ctx, key, func(ctx context.Context) (ImmutableTokenData, error) {
		return t.retrieveBridgeTokenData(ctx, chainID, token)
	})

	return tokenData, err
}

// GetPoolTokenData attempts to get pool token data from the cache otherwise it is fetched from the erc20 interface for that token.
func (t *tokenDataServiceImpl) GetPoolTokenData(ctx context.Context, chainID uint32, token common.Address, swapService fetcher.SwapService) (ImmutableTokenData, error) {
	key := cacheKey(chainID, token)
	if data, ok := t.tokenCache.Get(key); ok {
		return data, nil
	}

	tokenData, err := t.fetchAndCacheTokenData(ctx, key, func(ctx context.Context) (ImmutableTokenData, error) {
		return t.retrievePoolTokenData(ctx, token, swapService)
	})

	return tokenData, err
}

// GetCCTPTokenData attempts to get cctp token data from the cache otherwise it is fetched using the cctp ref.
func (t *tokenDataServiceImpl) GetCCTPTokenData(ctx context.Context, chainID uint32, token common.Address, backend bind.ContractBackend) (ImmutableTokenData, error) {
	key := cacheKey(chainID, token)
	if data, ok := t.tokenCache.Get(key); ok {
		return data, nil
	}

	tokenData, err := t.fetchAndCacheTokenData(ctx, key, func(ctx context.Context) (ImmutableTokenData, error) {
		return t.retrieveCCTPTokenData(ctx, token, chainID, backend)
	})

	return tokenData, err
}

// retrieveTokenData retrieves the token data from the bridge config contract.
func (t *tokenDataServiceImpl) retrieveBridgeTokenData(parentCtx context.Context, chainID uint32, token common.Address) (ImmutableTokenData, error) {
	res := immutableTokenImpl{}

	ctx, cancel := context.WithTimeout(parentCtx, maxAttemptTime)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		//nolint: wrapcheck
		return retry.WithBackoff(ctx, func(ctx context.Context) error {
			tokenData, err := t.service.GetToken(ctx, chainID, token)
			if err != nil {
				return fmt.Errorf("could not get token data: %w", err)
			}

			res.decimals = tokenData.TokenDecimals

			return nil
		}, retry.WithMaxAttemptTime(maxAttemptTime), retry.WithMaxAttempts(maxAttempts))
	})

	g.Go(func() error {
		//nolint: wrapcheck
		return retry.WithBackoff(ctx, func(ctx context.Context) error {
			nullableTokenID, err := t.service.GetTokenID(ctx, big.NewInt(int64(chainID)), token)
			if err != nil {
				return fmt.Errorf("could not get token data: %w", err)
			}

			res.tokenID = *nullableTokenID

			return nil
		}, retry.WithMaxAttemptTime(maxAttemptTime), retry.WithMaxAttempts(maxAttempts))
	})

	err := g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not get token data: %w", err)
	}
	res.tokenAddress = token.String()

	return res, nil
}

// retrievePoolTokenData retrieves the token data from the bridge config contract.
func (t *tokenDataServiceImpl) retrievePoolTokenData(ctx context.Context, token common.Address, swapService fetcher.SwapService) (ImmutableTokenData, error) {
	symbol, decimals, err := swapService.GetTokenMetaData(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("could not get token data: %w", err)
	}

	normalizedSymbol := normalizeSymbol(*symbol)
	tokenID, exists := t.tokenSymbolToIDs[normalizedSymbol]
	if !exists {
		tokenID = normalizedSymbol
	}

	return immutableTokenImpl{
		decimals:     *decimals,
		tokenID:      tokenID,
		tokenAddress: token.Hex(),
	}, nil
}

// retrieveCCTPTokenData retrieves the token data from the cctp contract.
func (t *tokenDataServiceImpl) retrieveCCTPTokenData(ctx context.Context, address common.Address, chainID uint32, backend bind.ContractBackend) (ImmutableTokenData, error) {
	erc20Contract, err := erc20.NewERC20Ref(address, chainID, backend)

	if err != nil {
		return nil, fmt.Errorf("could not create erc20: %w", err)
	}
	decimal, symbol, err := erc20Contract.GetTokenData(ctx)

	if err != nil {
		return nil, fmt.Errorf("could not get token data: %w", err)
	}
	return immutableTokenImpl{
		decimals:     decimal,
		tokenID:      symbol,
		tokenAddress: address.Hex(),
	}, nil
}

// fetchAndCacheTokenData fetches the token data and caches it.
func (t *tokenDataServiceImpl) fetchAndCacheTokenData(ctx context.Context, key string, fetchFunc func(ctx context.Context) (ImmutableTokenData, error)) (ImmutableTokenData, error) {
	tokenData, err := fetchFunc(ctx)
	if err != nil {
		return nil, err // Error already wrapped in fetchFunc
	}
	t.tokenCache.Add(key, tokenData)
	return tokenData, nil
}

// normalizeSymbol normalizes token symbols to a standard representation.
func normalizeSymbol(symbol string) string {
	normalizedSymbols := []string{"dai", "usdc", "nusd", "usdt", "eth", "avax", "movr", "frax", "jewel"}
	symbolLower := strings.ToLower(symbol)
	for _, normSymbol := range normalizedSymbols {
		if strings.Contains(symbolLower, normSymbol) {
			return normSymbol
		}
	}
	return symbolLower
}

// cacheKey generates a standardized cache key for a given token.
func cacheKey(chainID uint32, token common.Address) string {
	return fmt.Sprintf("token_%d_%s", chainID, token.Hex())
}
