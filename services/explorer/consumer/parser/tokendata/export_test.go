package tokendata

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

// RetrieveBridgeTokenData wraps retrieveBridgeTokenData for testing bridge token data retrieval.
func (t *tokenDataServiceImpl) RetrieveBridgeTokenData(parentCtx context.Context, chainID uint32, token common.Address) (ImmutableTokenData, error) {
	return t.retrieveBridgeTokenData(parentCtx, chainID, token)
}

// NewImmutableToken creates a new immutable token for testing.
func NewImmutableToken(decimals uint8, tokenID string) ImmutableTokenData {
	return immutableTokenImpl{
		decimals: decimals,
		tokenID:  tokenID,
	}
}
