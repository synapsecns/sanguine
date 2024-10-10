package tokendata

import (
	"context"
	"fmt"
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
type Service interface {
	// GetTokenData attempts to get token data from the cache otherwise its fetched from the bridge config
	GetTokenData(ctx context.Context, chainID uint32, token common.Address) (ImmutableTokenData, error)
	// GetPoolTokenData attempts to get pool token data from the cache otherwise its fetched from the erc20 interface
	GetPoolTokenData(ctx context.Context, chainID uint32, token common.Address, swapService fetcher.SwapService) (ImmutableTokenData, error)
	// GetCCTPTokenData attempts to get the token symbol from the cctp contract
	GetCCTPTokenData(ctx context.Context, chainID uint32, token common.Address, cctpService fetcher.CCTPService) (ImmutableTokenData, error)
}

// TokenData is the data for a token -- specifically used for searching the token data map for Synapse X tokens not in the contract config.
type TokenData struct {
	TokenID      string
	Decimals     uint8
	TokenAddress string
}

const cacheSize = 3000

// maxAttemptTime is how many times we will attempt to get the token data.
const maxAttemptTime = time.Minute * 5
const maxAttempt = 10

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

// GetTokenData attempts to get token data from the cache otherwise it is fetched from the bridge config.
func (t *tokenDataServiceImpl) GetTokenData(ctx context.Context, chainID uint32, token common.Address) (ImmutableTokenData, error) {
	// Fix for tokens not in the config or cache (typically SynapseX)
	mapKey := fmt.Sprintf("%d_%s", chainID, token.String())
	if tokenData, ok := tokenDataMap[mapKey]; ok {
		immutableData := immutableTokenImpl{
			tokenID:      tokenData.TokenID,
			decimals:     tokenData.Decimals,
			tokenAddress: tokenData.TokenAddress,
		}
		t.tokenCache.Add(mapKey, immutableData)
		return immutableData, nil
	}
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

// GetPoolTokenData attempts to get pool token data from the cache otherwise it is fetched from the erc20 interface for that token.
func (t *tokenDataServiceImpl) GetPoolTokenData(ctx context.Context, chainID uint32, token common.Address, swapService fetcher.SwapService) (ImmutableTokenData, error) {
	// Fix for tokens not in the config or cache (typically SynapseX)
	mapKey := fmt.Sprintf("%d_%s", chainID, token.String())
	if tokenData, ok := tokenDataMap[mapKey]; ok {
		immutableData := immutableTokenImpl{
			tokenID:      tokenData.TokenID,
			decimals:     tokenData.Decimals,
			tokenAddress: tokenData.TokenAddress,
		}
		t.tokenCache.Add(mapKey, immutableData)
		return immutableData, nil
	}
	key := fmt.Sprintf("token_%d_%s", chainID, token.Hex())
	if data, ok := t.tokenCache.Get(key); ok {
		return data, nil
	}

	tokenData, err := t.retrievePoolTokenData(ctx, token, swapService)
	if err != nil {
		return nil, fmt.Errorf("could not get token data: %w", err)
	}

	t.tokenCache.Add(key, tokenData)

	return tokenData, nil
}

// GetCCTPTokenData attempts to get cctp token data from the cache otherwise it is fetched using the cctp ref.
func (t *tokenDataServiceImpl) GetCCTPTokenData(ctx context.Context, chainID uint32, token common.Address, cctpService fetcher.CCTPService) (ImmutableTokenData, error) {
	// Fix for tokens not in the config or cache (typically SynapseX)
	mapKey := fmt.Sprintf("%d_%s", chainID, token.String())
	if tokenData, ok := tokenDataMap[mapKey]; ok {
		immutableData := immutableTokenImpl{
			tokenID:      tokenData.TokenID,
			decimals:     tokenData.Decimals,
			tokenAddress: tokenData.TokenAddress,
		}
		t.tokenCache.Add(mapKey, immutableData)
		return immutableData, nil
	}
	key := fmt.Sprintf("token_%d_%s", chainID, token.Hex())
	if data, ok := t.tokenCache.Get(key); ok {
		return data, nil
	}

	tokenData, err := t.retrieveCCTPTokenData(ctx, token, cctpService)
	if err != nil {
		return nil, fmt.Errorf("could not get token data: %w", err)
	}

	t.tokenCache.Add(key, tokenData)

	return tokenData, nil
}

// retrieveTokenData retrieves the token data from the bridge config contract
// this will retry for maxAttemptTime.
func (t *tokenDataServiceImpl) retrieveTokenData(parentCtx context.Context, chainID uint32, token common.Address) (ImmutableTokenData, error) {
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
		}, retry.WithMaxAttemptTime(maxAttemptTime), retry.WithMaxAttempts(maxAttempt))
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
		}, retry.WithMaxAttemptTime(maxAttemptTime), retry.WithMaxAttempts(maxAttempt))
	})

	err := g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not get token data: %w", err)
	}
	res.tokenAddress = token.String()

	return res, nil
}

// retrieveTokenData retrieves the token data from the bridge config contract
// this will retry for maxAttemptTime.
//
// nolint:cyclop
func (t *tokenDataServiceImpl) retrievePoolTokenData(parentCtx context.Context, token common.Address, swapService fetcher.SwapService) (ImmutableTokenData, error) {
	res := immutableTokenImpl{}

	ctx, cancel := context.WithTimeout(parentCtx, maxAttemptTime)
	defer cancel()

	err := retry.WithBackoff(ctx, func(ctx context.Context) error {
		symbol, decimals, err := swapService.GetTokenMetaData(ctx, token)
		if err != nil {
			return fmt.Errorf("could not get token data: %w", err)
		}

		if strings.Contains(strings.ToLower(*symbol), "dai") {
			*symbol = "dai"
		}
		if strings.Contains(strings.ToLower(*symbol), "usdc") {
			*symbol = "usdc"
		}
		if strings.Contains(strings.ToLower(*symbol), "nusd") {
			*symbol = "nusd"
		}
		if strings.Contains(strings.ToLower(*symbol), "usdt") {
			*symbol = "usdt"
		}
		if strings.Contains(strings.ToLower(*symbol), "eth") {
			*symbol = "eth"
		}
		if strings.Contains(strings.ToLower(*symbol), "avax") {
			*symbol = "avax"
		}
		if strings.Contains(strings.ToLower(*symbol), "movr") {
			*symbol = "movr"
		}
		if strings.Contains(strings.ToLower(*symbol), "frax") {
			*symbol = "frax"
		}
		if strings.Contains(strings.ToLower(*symbol), "jewel") {
			*symbol = "jewel"
		}

		res.tokenID = t.tokenSymbolToIDs[strings.ToLower(*symbol)]
		res.decimals = *decimals
		res.tokenAddress = token.String()

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not get pool token data: %w", err)
	}

	return res, nil
}

func (t *tokenDataServiceImpl) retrieveCCTPTokenData(parentCtx context.Context, tokenAddress common.Address, cctpService fetcher.CCTPService) (ImmutableTokenData, error) {
	res := immutableTokenImpl{}

	ctx, cancel := context.WithTimeout(parentCtx, maxAttemptTime)
	defer cancel()
	err := retry.WithBackoff(ctx, func(ctx context.Context) error {
		symbol, err := cctpService.GetTokenSymbol(ctx, tokenAddress)
		if err != nil {
			return fmt.Errorf("could not get cctp token: %w", err)
		}
		if strings.Contains(strings.ToLower(*symbol), "usdc") {
			*symbol = "usdc"
		}
		res.tokenID = t.tokenSymbolToIDs[strings.ToLower(*symbol)]
		res.decimals = 6 // TODO, as cctp bridging matures, retrieve this data from on chain somehow.

		return nil
	}, retry.WithMaxAttemptTime(maxAttemptTime), retry.WithMaxAttempts(maxAttempt))
	if err != nil {
		return nil, fmt.Errorf("could not get token data: %w", err)
	}
	res.tokenAddress = tokenAddress.String()

	return res, nil
}

var tokenDataMap = map[string]TokenData{
	"1_0x5f98805A4E8be255a32880FDeC7F6728C6568bA0":     {"LUSD", 18, "0x5f98805A4E8be255a32880FDeC7F6728C6568bA0"},
	"1_0xf939E0A03FB07F59A73314E73794Be0E57ac1b4E":     {"crvUSD", 18, "0xf939E0A03FB07F59A73314E73794Be0E57ac1b4E"},
	"8453_0x417Ac0e078398C154EdFadD9Ef675d30Be60Af93":  {"crvUSD", 18, "0x417Ac0e078398C154EdFadD9Ef675d30Be60Af93"},
	"8453_0xd9aAEc86B65D86f6A7B5B1b0c42FFA531710b6CA":  {"USDbC", 6, "0xd9aAEc86B65D86f6A7B5B1b0c42FFA531710b6CA"},
	"1_0x853d955acef822db058eb8505911ed77f175b99e":     {"FRAX", 18, "0x853d955acef822db058eb8505911ed77f175b99e"},
	"42161_0x17FC002b466eEc40DaE837Fc4bE5c67993ddBd6F": {"FRAX", 18, "0x17FC002b466eEc40DaE837Fc4bE5c67993ddBd6F"},
	"137_0x45c32fA6DF82ead1e2EF74d17b76547EDdFaFF89":   {"FRAX", 18, "0x45c32fA6DF82ead1e2EF74d17b76547EDdFaFF89"},
	"10_0x8c6f28f2F1A3C87F0f938b96d27520d9751ec8d9":    {"sUSD", 18, "0x8c6f28f2F1A3C87F0f938b96d27520d9751ec8d9"},
	"8453_0x50c5725949A6F0c72E6C4a641F24049A917DB0Cb":  {"DAI", 18, "0x50c5725949A6F0c72E6C4a641F24049A917DB0Cb"},
	"1_0xAdF7C35560035944e805D98fF17d58CDe2449389":     {"SPEC", 18, "0xAdF7C35560035944e805D98fF17d58CDe2449389"},
	"8453_0x96419929d7949D6A801A6909c145C8EEf6A40431":  {"SPEC", 18, "0x96419929d7949D6A801A6909c145C8EEf6A40431"},
	"480_0x2cFc85d8E48F8EAB294be644d9E25C3030863003":   {"WLD", 18, "0x2cFc85d8E48F8EAB294be644d9E25C3030863003"},
	"10_0xdC6fF44d5d932Cbd77B52E5612Ba0529DC6226F1":    {"WLD", 18, "0xdC6fF44d5d932Cbd77B52E5612Ba0529DC6226F1"},
	"480_0x79A02482A880bCE3F13e09Da970dC34db4CD24d1":   {"USDC.e", 6, "0x79A02482A880bCE3F13e09Da970dC34db4CD24d1"},
	"480_0x4200000000000000000000000000000000000006":   {"WETH", 18, "0x4200000000000000000000000000000000000006"},
	"1_0x163f8c2467924be0ae7b5347228cabf260318753":     {"WLD", 18, "0x163f8c2467924be0ae7b5347228cabf260318753"},
	// Add additional tokens that are not part of the cache yet (and not by nature in bridge config) here
}
