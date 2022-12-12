package backfill

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/exp/constraints"
)

// GetLogs exports logs for testing.
func (c ContractBackfiller) GetLogs(ctx context.Context, startHeight, endHeight uint64) (<-chan types.Log, <-chan int) {
	return c.getLogs(ctx, startHeight, endHeight)
}

// Clients exports clients for testing.
func (s *ScribeBackfiller) Clients() map[uint32][]ScribeBackend {
	return s.clients
}

// ChainID exports chainID for testing.
func (c ChainBackfiller) ChainID() uint32 {
	return c.chainID
}

// MakeRange exports makeRange for testing.
func MakeRange[T constraints.Integer](min, max T) []T {
	return makeRange(min, max)
}
