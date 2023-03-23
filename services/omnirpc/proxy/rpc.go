package proxy

import (
	"fmt"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/goccy/go-json"
	"github.com/hashicorp/go-multierror"
	"github.com/hedzr/cmdr/tool"
	"github.com/invopop/jsonschema"
	"golang.org/x/exp/slices"
)

// Confirmable returns the confirmability of an interface.
type Confirmable interface {
	isConfirmable() (bool, error)
}

// RPCRequests is a list or rpc requests.
type RPCRequests []RPCRequest

func (r RPCRequests) isConfirmable() (_ bool, errs error) {
	unconfirmable := false

	for i, request := range r {
		isConfirmable, err := request.isConfirmable()
		if err != nil {
			errs = multierror.Append(errs, fmt.Errorf("request at index %d: %s is not parsable", i, spew.Sprint(request)))
		}

		if !isConfirmable {
			unconfirmable = true
		}
	}

	if errs != nil {
		//nolint:wrapcheck
		return false, errs
	}

	return !unconfirmable, nil
}

// ByID will get an rpc requet by the id.
func (r RPCRequests) ByID(id int) *RPCRequest {
	for _, req := range r {
		if req.ID == id {
			return &req
		}
	}
	return nil
}

// RPCRequest is a raw rpc request format.
type RPCRequest struct {
	ID      int               `json:"id"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
	ChainID uint32            `json:"chainId"`
}

var _ Confirmable = RPCRequests{}

var _ Confirmable = RPCRequest{}

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

// isBatch determines if a request is batch. This method's implementation is borrowed from
// rpc/json.go in ethereum-go where it is unexported.
func isBatch(body []byte) bool {
	for _, c := range body {
		// skip insignificant whitespace (http://www.ietf.org/rfc/rfc4627.txt)
		if c == 0x20 || c == 0x09 || c == 0x0a || c == 0x0d {
			continue
		}
		return c == '['
	}
	return false
}

func parseRPCPayload(body []byte) (_ RPCRequests, err error) {
	if isBatch(body) {
		var rpcPayload []RPCRequest
		err = json.Unmarshal(body, &rpcPayload)
		if err != nil {
			return nil, fmt.Errorf("could not parse batch json payload: %w, must conform to: %s", err, rpcReqSchema)
		}
		return rpcPayload, nil
	}
	rpcRequest := RPCRequest{}
	err = json.Unmarshal(body, &rpcRequest)
	if err != nil {
		return nil, fmt.Errorf("could not parse json payload: %w, must conform to: %s", err, rpcReqSchema)
	}
	return []RPCRequest{rpcRequest}, nil
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

func (r RPCRequest) isConfirmable() (bool, error) {
	// TODO: handle batch methods
	// TODO: should we error on default?
	// TODO: look at RPCMethod.Comparable for lower, necessary?
	//nolint: exhaustive
	switch RPCMethod(r.Method) {
	case BlockByNumberMethod, PendingTransactionCountMethod:
		return isBlockNumConfirmable(r.Params[0]), nil
	case BlockNumberMethod, SyncProgressMethod, GasPriceMethod, MaxPriorityMethod, EstimateGasMethod:
		return false, nil
	case GetBalanceMethod, GetCodeMethod, TransactionCountMethod, CallMethod:
		return isBlockNumConfirmable(r.Params[1]), nil
	case StorageAtMethod:
		return isBlockNumConfirmable(r.Params[2]), nil
	case GetLogsMethod:
		return isFilterArgConfirmable(r.Params[0])
	// not confirmable because tx could be pending. We might want to handle w/ omnicast though
	// left separate for comment
	case SendRawTransactionMethod:
		return false, nil
	}
	return true, nil
}
