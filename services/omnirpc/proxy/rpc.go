package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hedzr/cmdr/tool"
	"github.com/invopop/jsonschema"
	"golang.org/x/exp/slices"
	"math/big"
)

// RPCRequest is a raw rpc request format.
type RPCRequest struct {
	ID     json.RawMessage   `json:"id"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}

func init() {
	schema := jsonschema.Reflect(&RPCRequest{})
	rawSchema, err := schema.MarshalJSON()
	if err != nil {
		panic(fmt.Errorf("could not reflect rpc schema: %w", err))
	}
	rpcReqSchema = string(rawSchema)
}

// rpcReqSchema contains the raw rpc request schema.
var rpcReqSchema string

func parseRPCPayload(body []byte) (request RPCRequest, err error) {
	rpcPayload := RPCRequest{}
	err = json.Unmarshal(body, &rpcPayload)
	if err != nil {
		return RPCRequest{}, fmt.Errorf("could not parse json payload: %w, must conform to: %s", err, rpcReqSchema)
	}

	return rpcPayload, nil
}

func isBlockNumConfirmable(arg json.RawMessage) bool {
	// nonConfirmableBlockNumArgs is a list of non numerical block args
	var nonConfirmableBlockNumArgs = []string{"latest", "pending"}

	return !slices.Contains(nonConfirmableBlockNumArgs, tool.StripQuotes(string(arg)))
}

// isFilterArgConfirmable checks if filter.filterCriteria is confirmable.
func isFilterArgConfirmable(arg json.RawMessage) (bool, error) {
	// cast latest block number to a big int for comparison
	latestBlockNumber := new(big.Int).SetInt64(rpc.LatestBlockNumber.Int64())

	filterCriteria := filters.FilterCriteria{}
	err := filterCriteria.UnmarshalJSON(arg)
	if err != nil {
		return false, fmt.Errorf("could not unmarshall filter: %w", err)
	}

	// Block filter requested, construct a single-shot filter
	if filterCriteria.BlockHash != nil {
		return true, nil
	}

	usesLatest := filterCriteria.FromBlock.Cmp(latestBlockNumber) == 0 || filterCriteria.ToBlock.Cmp(latestBlockNumber) == 0
	return !usesLatest, nil
}

func isConfirmable(body []byte) (bool, error) {
	payload, err := parseRPCPayload(body)
	if err != nil {
		return false, fmt.Errorf("could not parse payload: %w", err)
	}

	// TODO: handle batch methods
	// TODO: should we error on default?
	switch payload.Method {
	case "eth_getBlockByNumber", "eth_getBlockTransactionCountByNumber":
		return isBlockNumConfirmable(payload.Params[0]), nil
	case "eth_blockNumber", "eth_syncing", "eth_gasPrice", "eth_maxPriorityFeePerGas", "eth_estimateGas":
		return false, nil
	case "eth_getBalance", "eth_getCode", "eth_getTransactionCount", "eth_call":
		return isBlockNumConfirmable(payload.Params[1]), nil
	case "eth_getStorageAt":
		return isBlockNumConfirmable(payload.Params[2]), nil
	case "eth_getLogs":
		return isFilterArgConfirmable(payload.Params[0])
	// not confirmable because tx could be pending. We might want to handle w/ omnicast though
	// left separate for comment
	case "eth_sendRawTransaction":
		return false, nil
	}
	return true, nil
}
