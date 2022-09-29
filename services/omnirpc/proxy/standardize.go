package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
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

// JSONError is used to hold a json error
type JSONError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// rpcTransaction is an eth rpc transaction (copied from ethclient)
type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

// txExtraInfo contains extra txinfo (Copied from ethclient)
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

// fullRpcBlock is used to ensure parity by encoding both the header and the block
type fullRpcBlock struct {
	Block  rpcBlock      `json:"rpc_block"`
	Header *types.Header `json:"header"`
}

func StandardizeResponse(method string, body []byte) (out []byte, err error) {
	// TODO: use a sync.pool for acquiring/releasing these structs
	var rpcMessage JSONRPCMessage
	err = json.Unmarshal(body, &rpcMessage)

	switch method {
	case ChainIDMethod, BlockNumberMethod:
		var result hexutil.Big
		if err := json.Unmarshal(rpcMessage.Result, &result); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(result)
	case TransactionByHashMethod:
		var rpcBody rpcTransaction
		if err := json.Unmarshal(rpcMessage.Result, &rpcBody); err != nil {
			return nil, fmt.Errorf("could not parse: %w", err)
		}
		out, err = json.Marshal(rpcBody)

	case BlockByHashMethod, BlockByNumberMethod:
		var head *types.Header
		var rpcBody rpcBlock

		if err := json.Unmarshal(rpcMessage.Result, &head); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(rpcMessage.Result, &rpcBody); err != nil {
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

		fullBlock := fullRpcBlock{
			Block:  rpcBody,
			Header: head,
		}

		out, err = json.Marshal(fullBlock)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshall full block: %w", err)
		}
	}

	return out, nil
}
