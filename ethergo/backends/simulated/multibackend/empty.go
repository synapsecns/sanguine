package multibackend

import (
	"context"
	"time"
)

// EmptyBlock mines an empty block at time. Must be greater than previous block time.
func (b *SimulatedBackend) EmptyBlock(blockTime time.Time) {
	// Get the last block
	header, err := b.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic("could not fetch current header")
	}

	// Set the block time (adjust the timestamp)
	b.AdjustTime(time.Duration(blockTime.Unix()-int64(header.Time)) * time.Second)

	// Mine an empty block
	b.Commit()
}
