package proxy

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/goccy/go-json"
	"golang.org/x/sync/errgroup"
)

// JSONRPCMessage is A value of this type can a JSON-RPC request, notification, successful response or
// error response. Which one it is depends on the fields.
type JSONRPCMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *JSONError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

// JSONError is used to hold a json error.
type JSONError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// rpcTransaction is an eth rpc transaction (copied from ethclient).
type rpcTransaction struct {
	//nolint: unused
	tx *types.Transaction
	txExtraInfo
}

// txExtraInfo contains extra txinfo (Copied from ethclient).
type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

type rpcBlock struct {
	Hash         common.Hash      `json:"hash"`
	Transactions []rpcTransaction `json:"transactions"`
	UncleHashes  []common.Hash    `json:"uncles"`
}

// fullRPCBlock is used to ensure parity by encoding both the header and the block.
type fullRPCBlock struct {
	Block  rpcBlock      `json:"rpc_block"`
	Header *types.Header `json:"header"`
}

// rpcProgress is a copy of SyncProgressMethod with hex-encoded fields.
// copied from ethclient.
type rpcProgress struct {
	StartingBlock hexutil.Uint64
	CurrentBlock  hexutil.Uint64
	HighestBlock  hexutil.Uint64

	PulledStates hexutil.Uint64
	KnownStates  hexutil.Uint64

	SyncedAccounts      hexutil.Uint64
	SyncedAccountBytes  hexutil.Uint64
	SyncedBytecodes     hexutil.Uint64
	SyncedBytecodeBytes hexutil.Uint64
	SyncedStorage       hexutil.Uint64
	SyncedStorageBytes  hexutil.Uint64
	HealedTrienodes     hexutil.Uint64
	HealedTrienodeBytes hexutil.Uint64
	HealedBytecodes     hexutil.Uint64
	HealedBytecodeBytes hexutil.Uint64
	HealingTrienodes    hexutil.Uint64
	HealingBytecode     hexutil.Uint64
}

// feeHistoryResultMarshaling is used for parity checking against fee history
// copied from ethclient.
type feeHistoryResultMarshaling struct {
	OldestBlock  *hexutil.Big     `json:"oldestBlock"`
	Reward       [][]*hexutil.Big `json:"reward,omitempty"`
	BaseFee      []*hexutil.Big   `json:"baseFeePerGas,omitempty"`
	GasUsedRatio []float64        `json:"gasUsedRatio"`
}

// StandardizeResponse produces a standardized json response for hashing (strips extra fields)
// nolint: gocognit, cyclop
func standardizeResponse(method string, rpcMessage JSONRPCMessage) (out []byte, err error) {
	// TODO: use a sync.pool for acquiring/releasing these structs

OUTER:
	switch RPCMethod(method) {
	case ChainIDMethod, BlockNumberMethod, TransactionCountByHashMethod, GetBalanceMethod, GasPriceMethod, MaxPriorityMethod:
		var result hexutil.Big
		if err = json.Unmarshal(rpcMessage.Result, &result); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(result)
	case StorageAtMethod, GetCodeMethod:
		var result hexutil.Bytes
		if err = json.Unmarshal(rpcMessage.Result, &result); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(result)
	case TransactionCountMethod, EstimateGasMethod:
		var result hexutil.Uint64
		if err = json.Unmarshal(rpcMessage.Result, &result); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(result)
	case PendingTransactionCountMethod:
		var result hexutil.Uint
		if err = json.Unmarshal(rpcMessage.Result, &result); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(result)
	case TransactionByHashMethod, TransactionByBlockHashAndIndexMethod:
		var rpcBody rpcTransaction
		if err = json.Unmarshal(rpcMessage.Result, &rpcBody); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(rpcBody)
	case GetLogsMethod:
		var result []types.Log
		if err = json.Unmarshal(rpcMessage.Result, &result); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(result)
	case TransactionReceiptByHashMethod:
		var rpcBody *types.Receipt
		if err = json.Unmarshal(rpcMessage.Result, &rpcBody); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(rpcBody)
	case SyncProgressMethod:
		var syncing bool
		if err = json.Unmarshal(rpcMessage.Result, &syncing); err == nil {
			out, err = json.Marshal(syncing)
			break OUTER
		}

		var p rpcProgress
		if err = json.Unmarshal(rpcMessage.Result, &p); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}

		out, err = json.Marshal(p)
	case FeeHistoryMethod:
		var rpcBody feeHistoryResultMarshaling
		if err := json.Unmarshal(rpcMessage.Result, &rpcBody); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(rpcBody)

	case BlockByHashMethod, BlockByNumberMethod:
		var head *types.Header
		var rpcBody rpcBlock

		groupCtx, _ := errgroup.WithContext(context.Background())
		groupCtx.Go(func() error {
			if err = json.Unmarshal(rpcMessage.Result, &head); err != nil {
				return fmt.Errorf("could not parse: %w", err)
			}
			return nil
		})
		groupCtx.Go(func() error {
			if err = json.Unmarshal(rpcMessage.Result, &rpcBody); err != nil {
				return fmt.Errorf("could not parse: %w", err)
			}
			return nil
		})

		err = groupCtx.Wait()
		if err != nil {
			//nolint: wrapcheck
			return nil, err
		}

		// Quick-verify transaction and uncle lists. This mostly helps with debugging the server.
		if head.UncleHash == types.EmptyUncleHash && len(rpcBody.UncleHashes) > 0 {
			return nil, fmt.Errorf("server returned non-empty uncle list but block header indicates no uncles")
		}
		if head.UncleHash != types.EmptyUncleHash && len(rpcBody.UncleHashes) == 0 {
			return nil, fmt.Errorf("server returned empty uncle list but block header indicates uncles")
		}
		if head.TxHash == types.EmptyRootHash && len(rpcBody.Transactions) > 0 {
			return nil, fmt.Errorf("server returned non-empty transaction list but block header indicates no transactions")
		}
		if head.TxHash != types.EmptyRootHash && len(rpcBody.Transactions) == 0 {
			return nil, fmt.Errorf("server returned empty transaction list but block header indicates transactions")
		}

		fullBlock := fullRPCBlock{
			Block:  rpcBody,
			Header: head,
		}

		out, err = json.Marshal(fullBlock)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshall full block: %w", err)
		}
	// we don't do anything here, kept for exhaustiveness
	case CallMethod, SendRawTransactionMethod:
		return out, nil
	}

	//nolint: wrapcheck
	return out, err
}
