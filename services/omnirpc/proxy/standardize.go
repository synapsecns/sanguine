package proxy

import (
	"context"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"

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
	ID      int             `json:"id,omitempty"`
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

// rpcBlock is an eth rpc block.
type rpcBlock struct {
	Hash         common.Hash      `json:"hash"`
	Transactions []rpcTransaction `json:"transactions"`
	UncleHashes  []common.Hash    `json:"uncles"`
}

// rpcBlockNoTx is an eth rpc block without transactions (used by eth_getBlockByNumber calls with tx flag set to false).
type rpcBlockNoTx struct {
	Hash        common.Hash   `json:"hash"`
	UncleHashes []common.Hash `json:"uncles"`
}

// fullRPCBlock is used to ensure parity by encoding both the header and the block.
type fullRPCBlock struct {
	Block  rpcBlock      `json:"rpc_block"`
	Header *types.Header `json:"header"`
}

// fullRPCBlock is used to ensure parity by encoding both the header and the block.
type fullRPCBlockNoTx struct {
	Block  rpcBlockNoTx  `json:"rpc_block"`
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
//
//nolint:gocognit,cyclop
func standardizeResponse(ctx context.Context, req *rpc.Request, rpcMessage JSONRPCMessage) (out []byte, err error) {
	// TODO: use a sync.swap for acquiring/releasing these structs
	method := req.Method

OUTER:
	switch client.RPCMethod(method) {
	case client.ChainIDMethod, client.BlockNumberMethod, client.TransactionCountByHashMethod, client.GetBalanceMethod, client.GasPriceMethod, client.MaxPriorityMethod:
		var result hexutil.Big
		if err = json.Unmarshal(rpcMessage.Result, &result); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(result)
	case client.StorageAtMethod, client.GetCodeMethod:
		var result hexutil.Bytes
		if err = json.Unmarshal(rpcMessage.Result, &result); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(result)
	case client.TransactionCountMethod, client.EstimateGasMethod:
		var result hexutil.Uint64
		if err = json.Unmarshal(rpcMessage.Result, &result); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(result)
	case client.PendingTransactionCountMethod:
		var result hexutil.Uint
		if err = json.Unmarshal(rpcMessage.Result, &result); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(result)
	case client.TransactionByHashMethod, client.TransactionByBlockHashAndIndexMethod:
		var rpcBody rpcTransaction
		if err = json.Unmarshal(rpcMessage.Result, &rpcBody); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(rpcBody)
	case client.GetLogsMethod:
		var result []types.Log
		if err = json.Unmarshal(rpcMessage.Result, &result); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(result)
	case client.TransactionReceiptByHashMethod:
		var rpcBody *types.Receipt
		if err = json.Unmarshal(rpcMessage.Result, &rpcBody); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(rpcBody)
	case client.SyncProgressMethod:
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
	case client.FeeHistoryMethod:
		var rpcBody feeHistoryResultMarshaling
		if err := json.Unmarshal(rpcMessage.Result, &rpcBody); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(rpcBody)

	case client.BlockByHashMethod, client.BlockByNumberMethod:
		var head *types.Header
		var rpcBody rpcBlock
		var rpcBlockNoTx rpcBlockNoTx
		groupCtx, _ := errgroup.WithContext(ctx)
		groupCtx.Go(func() error {
			if err = json.Unmarshal(rpcMessage.Result, &head); err != nil {
				return fmt.Errorf("could not parse: %w", err)
			}
			return nil
		})

		var txFlag bool
		err := json.Unmarshal(req.Params[1], &txFlag)
		if txFlag {
			groupCtx.Go(func() error {
				if err = json.Unmarshal(rpcMessage.Result, &rpcBody); err != nil {
					return fmt.Errorf("could not parse: %w", err)
				}
				return nil
			})
		} else {
			groupCtx.Go(func() error {
				if err = json.Unmarshal(rpcMessage.Result, &rpcBlockNoTx); err != nil {
					return fmt.Errorf("could not parse: %w", err)
				}
				return nil
			})
		}

		err = groupCtx.Wait()
		if err != nil {
			//nolint: wrapcheck
			return nil, err
		}

		if head == nil {
			return nil, errors.New("header was empty")
		}

		// If tx flag is false, return output block with fullRPCBlockNoTx type.
		if !txFlag {
			outputBlock := fullRPCBlockNoTx{
				Block:  rpcBlockNoTx,
				Header: head,
			}
			out, err = json.Marshal(outputBlock)

			// Bypass the following uncle block verification.
			return out, nil
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

		outputBlock := fullRPCBlock{
			Block:  rpcBody,
			Header: head,
		}
		out, err = json.Marshal(outputBlock)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshall full block: %w", err)
		}
	// we don't do anything here, kept for exhaustiveness
	case client.CallMethod, client.SendRawTransactionMethod:
		return out, nil
	case client.NetVersionMethod, client.SubscribeMethod:
		return out, fmt.Errorf("method %s is not supported", rpcMessage.Method)
	}

	//nolint: wrapcheck
	return out, err
}
