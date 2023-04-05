package backfill

import (
	"context"
	"fmt"
	"github.com/benbjohnson/immutable"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/util"
	"math"
	"math/big"
)

// ScribeBackend is the set of functions that the scribe needs from a client.
type ScribeBackend interface {
	// ChainID gets the chain id from the rpc server.
	ChainID(ctx context.Context) (*big.Int, error)
	// BlockByNumber retrieves a block from the database by number, caching it
	// (associated with its hash) if found.
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	// TransactionByHash checks the pool of pending transactions in addition to the
	// blockchain. The isPending return value indicates whether the transaction has been
	// mined yet. Note that the transaction may not be part of the canonical chain even if
	// it's not pending.
	TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error)
	// TransactionReceipt returns the receipt of a mined transaction. Note that the
	// transaction may not be included in the current canonical chain even if a receipt
	// exists.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	// BlockNumber gets the latest block number.
	BlockNumber(ctx context.Context) (uint64, error)
	// FilterLogs executes a log filter operation, blocking during execution and
	// returning all the results in one batch.
	//
	// TODO(karalabe): Deprecate when the subscription one can return past data too.
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	// HeaderByNumber returns the block header with the given block number.
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	// BatchWithContext batches multiple
	BatchWithContext(ctx context.Context, calls ...w3types.Caller) error
}

// DialBackend returns a scribe backend.
func DialBackend(ctx context.Context, url string, handler metrics.Handler) (ScribeBackend, error) {
	//nolint:wrapcheck
	return client.DialBackend(ctx, url, handler)
}

// GetLogsInRange gets all logs in a range with a single batch request
// in successful cases an immutable list is returned, otherwise an error is returned.
func GetLogsInRange(ctx context.Context, backend ScribeBackend, startHeight uint64, endHeight uint64, subChunkSize uint64, contractAddress common.Address, expectedChainID uint64) (*immutable.List[*[]types.Log], error) {
	blockRange := (endHeight - startHeight) + 1
	subChunkCount := int(math.Ceil(float64(blockRange) / float64(subChunkSize)))
	iterator := util.NewChunkIterator(big.NewInt(int64(startHeight)), big.NewInt(int64(endHeight)), int(subChunkSize)-1, true)
	calls := make([]w3types.Caller, subChunkCount+2)
	results := make([][]types.Log, subChunkCount)
	chainID := new(uint64)
	calls[0] = eth.ChainID().Returns(chainID)

	maxHeight := new(big.Int)
	calls[1] = eth.BlockNumber().Returns(maxHeight)

	subChunkIdx := uint64(0)
	chunk := iterator.NextChunk()

	for chunk != nil {
		filter := ethereum.FilterQuery{
			FromBlock: chunk.StartBlock,
			ToBlock:   chunk.EndBlock,
			Addresses: []common.Address{contractAddress},
		}
		calls[subChunkIdx+2] = eth.Logs(filter).Returns(&results[subChunkIdx])
		subChunkIdx++
		chunk = iterator.NextChunk()
	}

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
