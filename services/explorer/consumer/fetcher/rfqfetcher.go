package fetcher

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/contracts/fastbridge"
)

// RFQService --output=mocks --case=underscore
type RFQService interface {
	// GetTokenSymbol gets the token symbol from the rfq ref..
	GetTokenSymbol(ctx context.Context, tokenAddress common.Address) (*string, error)
}

// rfqFetcher is the fetcher for token data from the rfq contract.
type rfqFetcher struct {
	rfq        *fastbridge.FastBridge
	backend    bind.ContractBackend
	rfqAddress common.Address
}

// NewRFQFetcher creates a new rfq fetcher.
func NewRFQFetcher(rfqAddress common.Address, backend bind.ContractBackend) (RFQService, error) {
	rfqRef, err := fastbridge.NewFastBridge(rfqAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind rfq contract: %w", err)
	}

	return &rfqFetcher{rfqRef, backend, rfqAddress}, nil
}

// GetTokenSymbol gets the token symbol from the rfq ref.
//
//nolint:all // havent removed context because it breaks other things.
func (p *rfqFetcher) GetTokenSymbol(ctx context.Context, tokenAddress common.Address) (*string, error) {
	// temporary solution since there are no contract functions we can use.
	// Convert the common.Address to a string for comparison.
	addressStr := tokenAddress.Hex()
	// Check if the address matches USDC or ETH and return the symbol directly.
	if strings.EqualFold(addressStr, "0xaf88d065e77c8cC2239327C5EDb3A432268e5831") || strings.EqualFold(addressStr, "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") || strings.EqualFold(addressStr, "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85") || strings.EqualFold(addressStr, "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913") {
		symbol := "USDC"
		return &symbol, nil
	} else if strings.EqualFold(addressStr, "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee") {
		symbol := "ETH"
		return &symbol, nil
	}
	err := fmt.Errorf("could not get rfq token symbol: token address not recognized")
	return nil, err
}
