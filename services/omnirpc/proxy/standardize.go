package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

func StandardizeResponse(method string, body []byte) (out []byte, err error) {
	// TODO: use a sync.pool for acquiring/releasing these structs
	var rpcMessage JSONRPCMessage
	err = json.Unmarshal(body, &rpcMessage)

	switch method {
	case ChainIDMethod:
		var result hexutil.Big
		err := json.Unmarshal(rpcMessage.Result, &result)
		if err != nil {
			return nil, fmt.Errorf("could not parse")
		}
		out, err = json.Marshal(result)
	}

	return out, nil
}
