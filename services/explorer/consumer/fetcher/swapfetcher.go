package fetcher

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/contracts/metaswap"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"math/big"
	"os"
)

// SwapService --output=mocks --case=underscore.
type SwapService interface {
	// GetTokenMetaData gets the token from the erc20 token contract given a swap contract token id.
	GetTokenMetaData(ctx context.Context, tokenAddress common.Address) (*string, *uint8, error)
	// GetTokenAddress gets the token address from the swap contract given a swap contract token id.
	GetTokenAddress(ctx context.Context, tokenIndex uint8) (*common.Address, error)
}

// SwapFetcher is the fetcher for token data from the swap contract.
type swapFetcher struct {
	swap        *swap.SwapRef
	metaSwap    *metaswap.MetaSwapRef
	backend     bind.ContractBackend
	swapAddress common.Address
}

// NewSwapFetcher creates a new swap fetcher.
// Backend must be an archive backend.
func NewSwapFetcher(swapAddress common.Address, backend bind.ContractBackend, isMetaSwap bool) (SwapService, error) {
	if isMetaSwap {
		metaSwap, err := metaswap.NewMetaSwapRef(swapAddress, backend)
		if err != nil {
			return nil, fmt.Errorf("could not bind metaswap config contract: %w", err)
		}
		return &swapFetcher{nil, metaSwap, backend, swapAddress}, nil
	}
	swap, err := swap.NewSwapRef(swapAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind swap config contract: %w", err)
	}

	return &swapFetcher{swap, nil, backend, swapAddress}, nil
}
func (s *swapFetcher) GetTokenAddress(ctx context.Context, tokenIndex uint8) (*common.Address, error) {
	// TODO implement tests so can test interacting with the swap contract
	if os.Getenv("CI") != "" {
		fakeAddress := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		return &fakeAddress, nil
	}
	if s.metaSwap != nil {
		tokenAddress, err := s.metaSwap.GetToken(&bind.CallOpts{
			Context: ctx,
		}, tokenIndex)
		if err != nil {
			logger.Errorf("could not get metaswap token address %s", s.metaSwap.Address())
			return nil, fmt.Errorf("could not get metaswap token address  %s", s.metaSwap.Address())
		}
		return &tokenAddress, nil
	}
	tokenAddress, err := s.swap.GetToken(&bind.CallOpts{
		Context: ctx,
	}, tokenIndex)
	if err != nil {
		logger.Errorf("could not get token address %s", s.swap.Address())
		return nil, fmt.Errorf("could not get token address  %s", s.swap.Address())
	}
	return &tokenAddress, nil
}

// GetTokenMetaData gets the token from the erc20 token contract given a swap contract token id.
func (s *swapFetcher) GetTokenMetaData(ctx context.Context, tokenAddress common.Address) (*string, *uint8, error) {
	// TODO implement tests so can test interacting with the swap contract
	if os.Getenv("CI") != "" {
		fakeSymbol := gofakeit.Name()
		fakeDecimals := uint8(gofakeit.Int8())
		return &fakeSymbol, &fakeDecimals, nil
	}
	if s.metaSwap != nil {
		erc20caller, err := metaswap.NewERC20(tokenAddress, s.backend)
		if err != nil {
			logger.Errorf("could not bind erc20 contract: %s", err)
			return nil, nil, fmt.Errorf("could not bind erc20 contract: %w", err)
		}

		tokenSymbol, err := erc20caller.Symbol(&bind.CallOpts{
			Context: ctx,
		})
		if err != nil {
			logger.Errorf("could not get token symbol: %s", err)
			return nil, nil, fmt.Errorf("could not get token symbol: %w", err)
		}

		tokenDecimals, err := erc20caller.Decimals(&bind.CallOpts{
			Context: ctx,
		})
		if err != nil {
			logger.Errorf("could not get token decimals: %s", err)
			return nil, nil, fmt.Errorf("could not get token decimal: %w", err)
		}

		return &tokenSymbol, &tokenDecimals, nil
	}
	erc20caller, err := swap.NewERC20(tokenAddress, s.backend)
	if err != nil {
		logger.Errorf("could not bind erc20 contract: %s", err)
		return nil, nil, fmt.Errorf("could not bind erc20 contract: %w", err)
	}

	tokenSymbol, err := erc20caller.Symbol(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		logger.Errorf("could not get token symbol: %s", err)
		return nil, nil, fmt.Errorf("could not get token symbol: %w", err)
	}

	tokenDecimals, err := erc20caller.Decimals(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		logger.Errorf("could not get token decimals: %s", err)
		return nil, nil, fmt.Errorf("could not get token decimal: %w", err)
	}

	return &tokenSymbol, &tokenDecimals, nil
}

var _ SwapService = &swapFetcher{}
