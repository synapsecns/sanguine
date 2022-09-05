package chunk_test

import (
	"fmt"
	"github.com/richardwilkes/toolbox/collection"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/synapse-node/pkg/common"
	"math/big"
	"strings"
	"testing"
)

// make sure are examples don't panic.
func (s ChunkSuite) TestChunkExamples() {
	NotPanics(s.T(), func() {
		ExampleChunk()
	})
}

// ExampleChunk demonstrates ascending and descending chunking.
func ExampleChunk() {
	// a ascending chunk iterator
	chunkIterator := common.NewChunkIterator(big.NewInt(1), big.NewInt(10), 1, true)
	var minMaxChunk, startEndChunks string
	for {
		nextChunk := chunkIterator.NextChunk()
		if nextChunk == nil {
			break
		}

		minMaxChunk += fmt.Sprintf("[%d-%d],", nextChunk.MinBlock(), nextChunk.MaxBlock())
		startEndChunks += fmt.Sprintf("[%d-%d],", nextChunk.StartBlock, nextChunk.EndBlock)
	}
	fmt.Println(strings.TrimSuffix(minMaxChunk, ","))
	fmt.Println(strings.TrimSuffix(startEndChunks, ","))

	// a descending chunk iterator
	chunkIterator = common.NewChunkIterator(big.NewInt(1), big.NewInt(10), 1, false)
	minMaxChunk = ""
	startEndChunks = ""

	for {
		nextChunk := chunkIterator.NextChunk()
		if nextChunk == nil {
			break
		}

		minMaxChunk += fmt.Sprintf("[%d-%d],", nextChunk.MinBlock(), nextChunk.MaxBlock())
		startEndChunks += fmt.Sprintf("[%d-%d],", nextChunk.StartBlock, nextChunk.EndBlock)
	}
	fmt.Println(strings.TrimSuffix(minMaxChunk, ","))
	fmt.Println(strings.TrimSuffix(startEndChunks, ","))
	// output:
	// [1-2],[3-4],[5-6],[7-8],[9-10]
	// [1-2],[3-4],[5-6],[7-8],[9-10]
	// [9-10],[7-8],[5-6],[3-4],[1-2]
	// [10-9],[8-7],[6-5],[4-3],[2-1]
}

// TestChunks runs a test to make sure chunks do not produce dulpicates and correctly end on the last block
// covering start-startBlock fully.
func (s ChunkSuite) TestChunks() {
	s.T().Parallel()
	ascendingTestCases := []bool{false}
	for _, isAcending := range ascendingTestCases {
		// capture the function literal
		isAcending := isAcending
		s.T().Run(fmt.Sprintf("ascending: %v", isAcending), func(t *testing.T) {
			t.Parallel()
			testRange(t, isAcending)
		})
	}
}

// testRange tests a block range.
func testRange(tb testing.TB, isAcending bool) {
	tb.Helper()

	startBlock := big.NewInt(100)
	endBlock := big.NewInt(500000)
	// we use an intset to make sure we hit every chunk
	intSet := collection.NewIntSet()

	chunkIterator := common.NewChunkIterator(startBlock, endBlock, 100, isAcending)
	for {
		nextChunk := chunkIterator.NextChunk()
		if nextChunk == nil {
			break
		}

		chunkStart := nextChunk.StartBlock.Uint64()
		chunkEnd := nextChunk.EndBlock.Uint64()

		// make sure min is always less then or equal to max
		True(tb, nextChunk.MaxBlock().Cmp(nextChunk.MinBlock()) >= 0)

		if isAcending {
			for i := chunkStart; i <= chunkEnd; i++ {
				// the item should not be in the intset
				False(tb, intSet.Contains(int(i)))
				intSet.Add(int(i))
			}

			if nextChunk.EndBlock.Cmp(endBlock) != 0 {
				Equal(tb, chunkEnd-chunkStart, uint64(100))
			}
		} else {
			for i := chunkEnd; i <= chunkStart; i++ {
				// the item should not be in the intset
				False(tb, intSet.Contains(int(i)))
				intSet.Add(int(i))
			}

			if nextChunk.EndBlock.Cmp(startBlock) != 0 {
				Equal(tb, chunkStart-chunkEnd, uint64(100))
			}
		}
	}

	if isAcending {
		Equal(tb, uint64(len(intSet)-1), big.NewInt(0).Sub(endBlock, startBlock).Uint64())
	} else {
		Equal(tb, uint64(len(intSet)-1), big.NewInt(0).Sub(startBlock, endBlock).Uint64())
	}
}
