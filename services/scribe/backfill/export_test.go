package backfill

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

// GetLogs exports logs for testing.
func (c ContractBackfiller) GetLogs(ctx context.Context, startHeight, endHeight uint64) (<-chan types.Log, <-chan bool) {
	return c.getLogs(ctx, startHeight, endHeight)
}
