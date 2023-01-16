package executor

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// Backend is the backend for the executor.
type Backend interface {
	// HeaderByNumber returns the block header with the given block number.
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}
