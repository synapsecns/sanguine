package backfill

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

func (c ContractBackfiller) GetLogs(ctx context.Context, startHeight, endHeight uint64) (<-chan types.Log, <-chan bool) {
	logsChan, doneChan := c.getLogs(ctx, startHeight, endHeight)
	return logsChan, doneChan
}
