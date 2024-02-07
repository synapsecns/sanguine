package fetcher

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/contracts/cctp"
)

// CCTPService --output=mocks --case=underscore.
type CCTPService interface {
	// GetTokenSymbol gets the token symbol from the cctp ref..
	GetTokenSymbol(ctx context.Context, tokenAddress common.Address) (*string, error)
}

// cctpFetcher is the fetcher for token data from the cctp contract.
type cctpFetcher struct {
	cctp        *cctp.SynapseCCTP
	backend     bind.ContractBackend
	cctpAddress common.Address
}

// NewCCTPFetcher creates a new cctp fetcher.
func NewCCTPFetcher(cctpAddress common.Address, backend bind.ContractBackend) (CCTPService, error) {
	cctpRef, err := cctp.NewSynapseCCTP(cctpAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("could not bind cctp contract: %w", err)
	}

	return &cctpFetcher{cctpRef, backend, cctpAddress}, nil
}
func (c *cctpFetcher) GetTokenSymbol(ctx context.Context, tokenAddress common.Address) (*string, error) {
	symbol, err := c.cctp.TokenToSymbol(&bind.CallOpts{
		Context: ctx,
	}, tokenAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get cctp token symbol: %w", err)
	}

	if symbol == "" {
		payload := NoTokenID
		return &payload, nil
	}

	return &symbol, nil
}
