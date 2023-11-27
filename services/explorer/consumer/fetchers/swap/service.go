package swap

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/synapsecns/sanguine/services/explorer/contracts/erc20"
	"github.com/synapsecns/sanguine/services/explorer/contracts/metaswap"
	swapContract "github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"time"
)

// ISwapFetcher gets data for tokens related to swap events.
//
//go:generate go run github.com/vektra/mockery/v2 --name=ISwapFetcher --output=mocks --case=underscore
type ISwapFetcher interface {
	// GetToken gets the token data from the erc20 token contract given a swap contract token id.
	GetToken(ctx context.Context, tokenAddress common.Address) (*uint8, *string, error)
	// GetTokenAddress gets the token address from a given token index.
	GetTokenAddress(ctx context.Context, tokenIndex uint8) (*common.Address, error)
	// ChainID returns the chain id of the swap contract.
	ChainID() uint32
}

const (
	cacheSize       = 3000
	maxAttemptTime  = time.Second * 120
	maxAttemptCount = 60
)

// SwapFetcher is the fetcher for token data from the swap contract.
type swapFetcherImpl struct {
	db           db.ConsumerDB
	addressCache *lru.TwoQueueCache[string, common.Address]
	swap         *swapContract.SwapRef
	metaSwap     *metaswap.MetaSwapRef
	backend      bind.ContractBackend
	chainID      uint32
	swapAddress  common.Address
}

// NewSwapFetcher creates a new swap fetcher.
// Backend must be an archive backend.
func NewSwapFetcher(db db.ConsumerDB, chainID uint32, swapAddress common.Address, backend bind.ContractBackend, isMetaSwap bool) (ISwapFetcher, error) {
	addressCache, cacheErr := lru.New2Q[string, common.Address](cacheSize)
	if cacheErr != nil {
		return nil, fmt.Errorf("could not create cache: %w", cacheErr)
	}
	if isMetaSwap {
		metaSwapRef, err := metaswap.NewMetaSwapRef(swapAddress, backend)
		if err != nil {
			return nil, fmt.Errorf("could not bind metaswap contract: %w", err)
		}
		return &swapFetcherImpl{db, addressCache, nil, metaSwapRef, backend, chainID, swapAddress}, nil
	}
	swapRef, err := swapContract.NewSwapRef(swapAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind swap contract: %w", err)
	}

	return &swapFetcherImpl{db, addressCache, swapRef, nil, backend, chainID, swapAddress}, nil
}
func (s *swapFetcherImpl) ChainID() uint32 {
	return s.chainID
}
func (s *swapFetcherImpl) GetTokenAddress(ctx context.Context, tokenIndex uint8) (*common.Address, error) {
	key := s.cacheKey(s.chainID, tokenIndex)
	var tokenAddress common.Address
	var swapAddress common.Address
	var err error
	var ok bool
	if tokenAddress, ok = s.addressCache.Get(key); ok {
		return &tokenAddress, nil
	}

	if s.metaSwap != nil {
		swapAddress = s.metaSwap.Address()
		tokenAddress, err = s.metaSwap.GetToken(&bind.CallOpts{
			Context: ctx,
		}, tokenIndex)
	} else {
		swapAddress = s.swap.Address()
		tokenAddress, err = s.swap.GetToken(&bind.CallOpts{
			Context: ctx,
		}, tokenIndex)
	}
	if err != nil {
		logger.Errorf("could not get token address %s", swapAddress)
		return nil, fmt.Errorf("could not get token address %s", swapAddress)
	}

	if err := s.saveTokenIndex(ctx, s.chainID, tokenIndex, &tokenAddress, s.swapAddress.String()); err != nil {
		return nil, fmt.Errorf("error saving token index: %w", err)
	}

	s.addressCache.Add(key, tokenAddress)
	return &tokenAddress, nil
}

// GetToken gets the token from the erc20 token contract given a swap contract token id.
func (s *swapFetcherImpl) GetToken(ctx context.Context, tokenAddress common.Address) (*uint8, *string, error) {
	erc20Contract, err := erc20.NewERC20Ref(tokenAddress, s.chainID, s.backend)

	if err != nil {
		return nil, nil, fmt.Errorf("could not create erc20: %w", err)
	}
	decimal, symbol, err := erc20Contract.GetTokenData(ctx)

	return &decimal, &symbol, nil
}

func (s *swapFetcherImpl) saveTokenIndex(ctx context.Context, chainID uint32, tokenIndex uint8, address *common.Address, contractAddress string) error {
	err := s.db.StoreTokenIndex(ctx, chainID, tokenIndex, address.String(), contractAddress)
	if err != nil {
		return fmt.Errorf("error saving token index: %w", err)
	}
	return nil
}

func (s *swapFetcherImpl) cacheKey(chainID uint32, tokenIndex uint8) string {
	return fmt.Sprintf("token_%d_%d", chainID, tokenIndex)
}
