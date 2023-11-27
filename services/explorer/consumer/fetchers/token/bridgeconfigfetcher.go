package token

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
)

// NoTokenID is the string returned when a token id is not found (not an authentic token).
const NoTokenID = "NO_TOKEN_ID"

// IBridgeConfigFetcher is the interface for fetching data from the bridge config.
//
//go:generate go run github.com/vektra/mockery/v2 --name=IBridgeConfigFetcher --output=mocks --case=underscore
type IBridgeConfigFetcher interface {
	// GetTokenID gets the token id from the bridge config contract.
	GetTokenID(ctx context.Context, chainID *big.Int, tokenAddress common.Address) (tokenID *string, err error)
	// GetToken gets the token from the bridge config contract.
	GetToken(ctx context.Context, chainID uint32, tokenAddress common.Address) (token *bridgeconfig.BridgeConfigV3Token, err error)
}

// bridgeConfigFetcher is the fetcher for the bridge config contract.
type bridgeConfigFetcher struct {
	bridgeConfigRef     *bridgeconfig.BridgeConfigRef
	bridgeConfigAddress common.Address
} // TODO switch bridge config based on block number

// NewBridgeConfigFetcher creates a new config fetcher.
// Backend must be an archive backend.
func NewBridgeConfigFetcher(bridgeConfigAddress common.Address, bridgeConfigRef *bridgeconfig.BridgeConfigRef) (IBridgeConfigFetcher, error) {
	return &bridgeConfigFetcher{bridgeConfigRef, bridgeConfigAddress}, nil
}

// GetTokenID gets the token id from the bridge config contract.
func (b *bridgeConfigFetcher) GetTokenID(ctx context.Context, chainID *big.Int, tokenAddress common.Address) (tokenID *string, err error) {
	tokenIDStr, err := b.bridgeConfigRef.GetTokenID(&bind.CallOpts{
		Context: ctx,
	}, tokenAddress, chainID)
	if err != nil {
		return nil, fmt.Errorf("could not get token id: %w", err)
	}

	if tokenIDStr == "" {
		payload := NoTokenID
		return &payload, nil
	}

	return &tokenIDStr, nil
}

// GetToken gets the token from the bridge config contract. Requires an archived note.
func (b *bridgeConfigFetcher) GetToken(ctx context.Context, chainID uint32, tokenAddress common.Address) (token *bridgeconfig.BridgeConfigV3Token, err error) {
	tok, err := b.bridgeConfigRef.GetTokenByAddress(&bind.CallOpts{
		Context: ctx,
	}, tokenAddress.String(), big.NewInt(int64(chainID)))
	if err != nil {
		return nil, fmt.Errorf("could not get token at block number %w", err)
	}

	return &tok, nil
}

var _ IBridgeConfigFetcher = &bridgeConfigFetcher{}
