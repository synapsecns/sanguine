package tokendata

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

// RetrieveTokenData wraps retrieveTokenData for testing.
func (t *tokenDataServiceImpl) RetrieveTokenData(parentCtx context.Context, chainID uint32, token common.Address) (ImmutableTokenData, error) {
	return t.retrieveTokenData(parentCtx, chainID, token)
}

// NewImmutableToken creates a new immutable token for testing.
func NewImmutableToken(decimals uint8, tokenID string) ImmutableTokenData {
	return immutableTokenImpl{
		decimals: decimals,
		tokenID:  tokenID,
	}
}
