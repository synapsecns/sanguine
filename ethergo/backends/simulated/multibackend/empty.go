package multibackend

import (
	"context"
	"fmt"
	"time"
)

// EmptyBlock mines an empty block at time. Must be greater than previous block time.
func (b *SimulatedBackend) EmptyBlock(blockTime time.Time) {
	// Get the current head block number
	currentBlock, err := b.BlockByNumber(context.Background(), nil)
	if err != nil {
		panic(fmt.Sprintf("could not fetch current block: %v", err))
	}

	// AdjustTime to the specified blockTime
	timeDiff := blockTime.Sub(time.Unix(int64(currentBlock.Time()), 0))
	if timeDiff > 0 {
		err = b.Backend.AdjustTime(timeDiff)
		if err != nil {
			panic(fmt.Sprintf("could not adjust time: %v", err))
		}
	}

	// Mine a new block
	b.Backend.Commit()
}
