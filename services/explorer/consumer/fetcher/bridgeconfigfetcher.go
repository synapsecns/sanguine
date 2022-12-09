package fetcher

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"math/big"
)

// NoTokenID is the string returned when a token id is not found (not an authentic token).
const NoTokenID = "NO_TOKEN_ID"

// BridgeConfigFetcher is the fetcher for the bridge config contract.
type BridgeConfigFetcher struct {
	bridgeConfigRef     *bridgeconfig.BridgeConfigRef
	bridgeConfigAddress common.Address
} // TODO switch bridge config based on block number

// NewBridgeConfigFetcher creates a new config fetcher.
// Backend must be an archive backend.
func NewBridgeConfigFetcher(bridgeConfigAddress common.Address, bridgeConfigRef *bridgeconfig.BridgeConfigRef) (*BridgeConfigFetcher, error) {
	return &BridgeConfigFetcher{bridgeConfigRef, bridgeConfigAddress}, nil
}

// GetTokenID gets the token id from the bridge config contract.
func (b *BridgeConfigFetcher) GetTokenID(ctx context.Context, chainID *big.Int, tokenAddress common.Address) (tokenID *string, err error) {
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
func (b *BridgeConfigFetcher) GetToken(ctx context.Context, chainID uint32, tokenAddress common.Address, blockNumber uint32) (token *bridgeconfig.BridgeConfigV3Token, err error) {
	tok, err := b.bridgeConfigRef.GetTokenByAddress(&bind.CallOpts{
		Context: ctx,
	}, tokenAddress.String(), big.NewInt(int64(chainID)))
	if err != nil {
		return nil, fmt.Errorf("could not get token at block number %d: %w", blockNumber, err)
	}

	return &tok, nil
}
