package fetcher

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
)

// SwapFetcher is the fetcher for token data from the swap contract.
type SwapFetcher struct {
	swap        *swap.SwapRef
	backend     bind.ContractBackend
	swapAddress common.Address
}

// NewSwapFetcher creates a new swap fetcher.
// Backend must be an archive backend.
func NewSwapFetcher(swapAddress common.Address, backend bind.ContractBackend) (*SwapFetcher, error) {
	swap, err := swap.NewSwapRef(swapAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind swap config contract: %w", err)
	}

	return &SwapFetcher{swap, backend, swapAddress}, nil
}

// GetTokenMetaData gets the token from the erc20 token contract given a swap contract token id.
func (s *SwapFetcher) GetTokenMetaData(ctx context.Context, tokenIndex uint8) (*string, *uint8) {
	tokenAddress, err := s.swap.GetToken(&bind.CallOpts{
		Context: ctx,
	}, tokenIndex)
	if err != nil {
		logger.Errorf("could not get token address: %s", err)
		return nil, nil
	}

	erc20caller, err := swap.NewERC20(tokenAddress, s.backend)
	if err != nil {
		logger.Errorf("could not bind erc20 contract: %s", err)
		return nil, nil
	}

	tokenSymbol, err := erc20caller.Symbol(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		logger.Errorf("could not get token symbol: %s", err)
		return &tokenSymbol, nil
	}

	tokenDecimals, err := erc20caller.Decimals(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		logger.Errorf("could not get token decimals: %s", err)
		return &tokenSymbol, nil
	}

	return &tokenSymbol, &tokenDecimals
}
