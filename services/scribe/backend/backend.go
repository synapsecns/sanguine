package backend

import (
	"context"
	"fmt"
	"math/big"

	"github.com/benbjohnson/immutable"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/util"
	"golang.org/x/exp/constraints"
)

// ScribeBackend is the set of functions that the scribe needs from a client.
type ScribeBackend interface {
	// ChainID gets the chain id from the rpc server.
	ChainID(ctx context.Context) (*big.Int, error)
	// BlockNumber gets the latest block number.
	BlockNumber(ctx context.Context) (uint64, error)
	// HeaderByNumber returns the block header with the given block number.
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	// BatchWithContext batches multiple
	BatchWithContext(ctx context.Context, calls ...w3types.Caller) error
}

// DialBackend returns a scribe backend.
func DialBackend(ctx context.Context, url string, handler metrics.Handler) (ScribeBackend, error) {
	//nolint:wrapcheck
	return client.DialBackend(ctx, url, handler, client.Capture(true))
}

// GetLogsInRange gets all logs in a range with a single batch request
// in successful cases an immutable list is returned, otherwise an error is returned.
func GetLogsInRange(ctx context.Context, backend ScribeBackend, contractAddresses []common.Address, expectedChainID uint64, chunks []*util.Chunk, topics [][]common.Hash) (*immutable.List[*[]types.Log], error) {
	calls := make([]w3types.Caller, len(chunks)+2)
	results := make([][]types.Log, len(chunks))
	chainID := new(uint64)
	calls[0] = eth.ChainID().Returns(chainID)

	maxHeight := new(big.Int)
	calls[1] = eth.BlockNumber().Returns(maxHeight)

	for i := 0; i < len(chunks); i++ {
		startBlock := chunks[i].StartBlock
		endBlock := chunks[i].EndBlock

		// handle desc iterator
		if startBlock.Uint64() > endBlock.Uint64() {
			startBlock = chunks[i].EndBlock
			endBlock = chunks[i].StartBlock
		}
		filter := ethereum.FilterQuery{
			FromBlock: startBlock,
			ToBlock:   endBlock,
			Addresses: contractAddresses,
			Topics:    topics,
		}
		calls[i+2] = eth.Logs(filter).Returns(&results[i])
	}

	// for logging purposes
	startHeight := chunks[0].StartBlock.Uint64()
	endHeight := chunks[len(chunks)-1].EndBlock.Uint64()
	if err := backend.BatchWithContext(ctx, calls...); err != nil {
		return nil, fmt.Errorf("could not fetch logs in range %d to %d: %w", startHeight, endHeight, err)
	}
	if expectedChainID != *chainID {
		return nil, fmt.Errorf("could not fetch logs in range %d to %d: Incorrect RPC used for expected chainID: %d, RPC used chainID %d instead", startHeight, endHeight, expectedChainID, chainID)
	}
	if endHeight > maxHeight.Uint64() {
		return nil, fmt.Errorf("could not fetch logs in range %d to %d: Block not available past max height: %d", startHeight, endHeight, maxHeight.Uint64())
	}

	// use an immutable list for additional safety to the caller, don't allocate until batch returns successfully
	res := immutable.NewListBuilder[*[]types.Log]()
	for _, result := range results {
		logChunk := result
		if len(result) > 0 {
			res.Append(&logChunk)
		}
	}

	return res.List(), nil
}

// BlockHashesInRange gets all block hashes in a range with a single batch request
// in successful cases an immutable map is returned of [height->hash], otherwise an error is returned.
func BlockHashesInRange(ctx context.Context, backend ScribeBackend, startHeight uint64, endHeight uint64) (*immutable.Map[uint64, string], error) {
	// performance impact will be negligible here because of external constraints on blocksize
	blocks := MakeRange(startHeight, endHeight)
	bulkSize := len(blocks)
	calls := make([]w3types.Caller, bulkSize)
	results := make([]types.Header, bulkSize)

	for i, blockNumber := range blocks {
		calls[i] = eth.HeaderByNumber(new(big.Int).SetUint64(blockNumber)).Returns(&results[i])
	}

	if err := backend.BatchWithContext(ctx, calls...); err != nil {
		return nil, fmt.Errorf("could not fetch blocks in range %d to %d: %w", startHeight, endHeight, err)
	}

	// use an immutable map for additional safety to the caller, don't allocate until batch returns successfully
	res := immutable.NewMapBuilder[uint64, string](nil)
	for _, result := range results {
		res.Set(result.Number.Uint64(), result.Hash().String())
	}

	return res.Map(), nil
}

// MakeRange returns a range of integers from min to max inclusive.
func MakeRange[T constraints.Integer](min, max T) []T {
	a := make([]T, max-min+1)
	for i := range a {
		a[i] = min + T(i)
	}
	return a
}
