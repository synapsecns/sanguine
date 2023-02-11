package near

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// AuroraClient contains the methods necessary to use the aurora client adapter.
type AuroraClient interface {
	// CallContext performs a JSON-RPC call with the given arguments. If the context is
	// canceled before the call has successfully returned, CallContext returns immediately.
	//
	// The result must be a pointer so that package json can unmarshal into it. You
	// can also pass nil, in which case the result is ignored.
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	// BatchCallContext sends all given requests as a single batch and waits for the server
	// to return a response for all of them. The wait duration is bounded by the
	// context's deadline.
	//
	// In contrast to CallContext, BatchCallContext only returns errors that have occurred
	// while sending the request. Any error specific to a request is reported through the
	// Error field of the corresponding BatchElem.
	//
	// Note that batch calls may not be executed atomically on the server side.
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
	// BlockNumber gets the latest block number
	BlockNumber(ctx context.Context) (uint64, error)
	// TransactionByHash gets a tx by hash
	TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error)
}

// BlockByNumber overrides the get block unmarshalling to be compatible with near.
func BlockByNumber(ctx context.Context, c AuroraClient, number *big.Int) (*types.Block, error) {
	if number == nil {
		latestBlock, err := c.BlockNumber(ctx)
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}

		return getBlock(ctx, c, "eth_getBlockByNumber", toBlockNumArg(new(big.Int).SetUint64(latestBlock)))
	}
	return getBlock(ctx, c, "eth_getBlockByNumber", toBlockNumArg(number))
}

// BlockByHash overrides the get block unmarshalling to be compatible with near.
func BlockByHash(ctx context.Context, c AuroraClient, hash common.Hash) (*types.Block, error) {
	return getBlock(ctx, c, "eth_getBlockByHash", hash, true)
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	pending := big.NewInt(-1)
	if number.Cmp(pending) == 0 {
		return "pending"
	}
	return hexutil.EncodeBig(number)
}

// see https://git.io/Jy1Oc
type rpcBlock struct {
	Hash         common.Hash   `json:"hash"`
	Transactions []common.Hash `json:"transactions"`
	UncleHashes  []common.Hash `json:"uncles"`
}

// getBlock is a temporary workaround for https://github.com/aurora-is-near/aurora-relayer/pull/141
// amd will be removed as soon as possible.
//
//nolint:gocognit,cyclop
func getBlock(ctx context.Context, c AuroraClient, method string, args ...interface{}) (*types.Block, error) {
	var raw json.RawMessage
	err := c.CallContext(ctx, &raw, method, args...)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	} else if len(raw) == 0 {
		//nolint:wrapcheck
		return nil, ethereum.NotFound
	}
	// Decode header and transactions.
	var head *types.Header
	var body rpcBlock
	if err := json.Unmarshal(raw, &head); err != nil {
		//nolint: wrapcheck
		return nil, err
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		//nolint: wrapcheck
		return nil, err
	}

	// near doesn't have uncles: manually overwrite the uncle hash with the empty uncle hash
	head.UncleHash = types.EmptyUncleHash

	if len(body.Transactions) == 0 {
		head.TxHash = types.EmptyRootHash
	}

	// Quick-verify transaction and uncle lists. This mostly helps with debugging the server.
	if head.UncleHash == types.EmptyUncleHash && len(body.UncleHashes) > 0 {
		return nil, fmt.Errorf("server returned non-empty uncle list but block header indicates no uncles")
	}
	if head.UncleHash != types.EmptyUncleHash && len(body.UncleHashes) == 0 {
		return nil, fmt.Errorf("server returned empty uncle list but block header indicates uncles")
	}
	if head.TxHash == types.EmptyRootHash && len(body.Transactions) > 0 {
		return nil, fmt.Errorf("server returned non-empty transaction list but block header indicates no transactions")
	}
	if head.TxHash != types.EmptyRootHash && len(body.Transactions) == 0 {
		return nil, fmt.Errorf("server returned empty transaction list but block header indicates transactions")
	}

	// there are no uncles in aurora
	var uncles []*types.Header

	txes, err := getTransactions(ctx, c, body.Transactions, body.Hash)
	if err != nil {
		return nil, err
	}

	return types.NewBlockWithHeader(head).WithBody(txes, uncles), nil
}
