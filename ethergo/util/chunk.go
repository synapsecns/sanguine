package util

import (
	"github.com/synapsecns/sanguine/core"
	"math/big"
	"sync"
)

// Chunk represents an individual chunk of startBlock-startBlock.
// TODO: this needs to be moved to scribe.
type Chunk struct {
	// StartBlock for this chunk. It is less then end block in ascending txes and greater in descending.
	StartBlock *big.Int
	// EndBlock for this chunk
	EndBlock *big.Int
	// ascending is whether or not the chunk was made in ascending mode. This is used
	// for min/max without a comparotor operation
	ascending bool
}

// MinBlock returns the minimum of start block and end block. This is useful because start/end change
// based on the ordering of the chunk.
func (c *Chunk) MinBlock() *big.Int {
	if c.ascending {
		return c.StartBlock
	}
	return c.EndBlock
}

// MaxBlock returns the minimum of start block and end block. This is useful because start/end change
// based on the ordering of the chunk.
func (c *Chunk) MaxBlock() *big.Int {
	if c.ascending {
		return c.EndBlock
	}
	return c.StartBlock
}

// baseChunkIterator contains a non-directional chunk iterator.
type baseChunkIterator struct {
	// chunkSize is the chunk size
	chunkSize int
	// mux locks the iterator to prevent concurrent issues
	//nolint: structcheck // triggers a false positive. This is used in ascendingChunkIterator and descendingChunkIterator
	mux sync.Mutex
}

// ascendingChunkIterator returns an iterator going from lowest -> highest.
type ascendingChunkIterator struct {
	*baseChunkIterator
	// endBlock is the last block to chunk at
	endBlock *big.Int
	// lastIteratedBlock is the most recent iterated block. This will always be less then endBlock
	lastIteratedBlock *big.Int
}

// descendingChunkIterator returns an iterator going from highest->lowest.
type descendingChunkIterator struct {
	*baseChunkIterator
	// startBlock is the last block to chunk at
	startBlock *big.Int
	// lastIteratedBlock is the block we most recently iterator over. This will always be greater then start block
	lastIteratedBlock *big.Int
}

// ChunkIterator is an instantiation of a stateful iterator that groups startBlock-endBlock blocks into length
// chunkSize.
type ChunkIterator interface {
	// NextChunk gets the next chunk. If the iterator has completed null will be returned.
	NextChunk() *Chunk
}

// NewChunkIterator creates an iterator for a range of elements (startBlock-endBlock) split into groups the length of chunkSize
// the final chunk has any remaining elements. An iterator is used here as opposed to a slice to save memory while iterating over *very
// large* chunks.This follows the stateful iterator pattern: https://ewencp.org/blog/golang-iterators/index.html
func NewChunkIterator(startBlock, endBlock *big.Int, chunkSize int, ascending bool) ChunkIterator {
	baseIterator := &baseChunkIterator{
		chunkSize: chunkSize,
	}

	if ascending {
		return &ascendingChunkIterator{
			baseChunkIterator: baseIterator,
			endBlock:          core.CopyBigInt(endBlock),
			// lastIteratedBlock is set to StartBlock - 1
			lastIteratedBlock: big.NewInt(0).Sub(startBlock, big.NewInt(1)),
		}
	}

	return &descendingChunkIterator{
		baseChunkIterator: baseIterator,
		startBlock:        core.CopyBigInt(startBlock),
		// lastIteratedBlock is set to endBlock + 1 since we subtract one on every chunk
		lastIteratedBlock: big.NewInt(0).Add(endBlock, big.NewInt(1)),
	}
}

// NextChunk gets the next chunk in descending order. Returns nil if complete.
func (d *descendingChunkIterator) NextChunk() *Chunk {
	d.mux.Lock()
	defer d.mux.Unlock()

	startBlock := big.NewInt(0).Sub(d.lastIteratedBlock, big.NewInt(1))
	// if the last block is greater than the start block, we're done
	if startBlock.Cmp(d.startBlock) < 0 {
		return nil
	}

	endBlock := big.NewInt(0).Sub(startBlock, big.NewInt(int64(d.chunkSize)))

	// if the last block is less then the end block, return the last block
	if endBlock.Cmp(d.startBlock) < 0 {
		endBlock = core.CopyBigInt(d.startBlock)
	}

	d.lastIteratedBlock = core.CopyBigInt(endBlock)

	return &Chunk{
		StartBlock: startBlock,
		EndBlock:   endBlock,
		ascending:  false,
	}
}

// NextChunk gets the next chunk in ascending order. Returns nil if complete.
func (c *ascendingChunkIterator) NextChunk() *Chunk {
	c.mux.Lock()
	defer c.mux.Unlock()

	startBlock := big.NewInt(0).Add(c.lastIteratedBlock, big.NewInt(1))
	// if the last block is greater than the start block, we're done
	if startBlock.Cmp(c.endBlock) > 0 {
		return nil
	}

	endBlock := big.NewInt(0).Add(startBlock, big.NewInt(int64(c.chunkSize)))

	// if the last block is greater then the end block, return the last block
	if endBlock.Cmp(c.endBlock) > 0 {
		endBlock = core.CopyBigInt(c.endBlock)
	}

	c.lastIteratedBlock = core.CopyBigInt(endBlock)

	return &Chunk{
		StartBlock: startBlock,
		EndBlock:   endBlock,
		ascending:  true,
	}
}
