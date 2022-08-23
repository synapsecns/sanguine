package multibackend

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"time"
)

// EmptyBlock mines an empty block at time. Must be greater than previous block time.
func (b *SimulatedBackend) EmptyBlock(blockTime time.Time) {
	// Get the last block
	block, err := b.blockByHash(context.TODO(), b.pendingBlock.ParentHash())
	if err != nil {
		panic("could not fetch parent")
	}

	// create a new chain with the block
	blocks, _ := core.GenerateChain(b.config, block, ethash.NewFaker(), b.database, 1, func(number int, block *core.BlockGen) {
		var prevBlockTime int64
		fmt.Println(block.Number())
		if block.Number().Uint64() != 1 {
			prevBlockTime = int64(block.PrevBlock(-1).Time())
		}
		block.OffsetTime(blockTime.Unix() - prevBlockTime)
	})
	stateDB, _ := b.blockchain.State()

	// add the empty block
	b.pendingBlock = blocks[0]
	b.pendingState, _ = state.New(b.pendingBlock.Root(), stateDB.Database(), nil)
	b.Commit()
}
