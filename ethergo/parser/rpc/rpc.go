package rpc

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/invopop/jsonschema"
)

// Requests is a list or rpc requests.
type Requests []Request

// Method returns the method of the rpc request.
func (r Requests) Method() string {
	reqLength := len(r)
	if reqLength == 0 {
		return "none"
	}

	if reqLength > 1 {
		return "batch"
	}

	return r[0].Method
}

// ByID will get an rpc requet by the id.
func (r Requests) ByID(id int) *Request {
	for _, req := range r {
		if req.ID == id {
			return &req
		}
	}
	return nil
}

// Request is a raw rpc request format.
type Request struct {
	ID     int               `json:"id"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}

func init() {
	schema := jsonschema.Reflect(&Request{})
	rawSchema, err := schema.MarshalJSON()
	if err != nil {
		panic(fmt.Errorf("could not reflect rpc schema: %w", err))
	}
	rpcReqSchema = string(rawSchema)
}

// rpcReqSchema contains the raw rpc request schema.
var rpcReqSchema string

// IsBatch determines if a request is batch. This method's implementation is borrowed from
// rpc/json.go in ethereum-go where it is unexported.
func IsBatch(body []byte) bool {
	for _, c := range body {
		// skip insignificant whitespace (http://www.ietf.org/rfc/rfc4627.txt)
		if c == 0x20 || c == 0x09 || c == 0x0a || c == 0x0d {
			continue
		}
		return c == '['
	}
	return false
}

// ParseRPCPayload parses a raw rpc request body and returns a list of rpc requests.
func ParseRPCPayload(body []byte) (_ Requests, err error) {
	if IsBatch(body) {
		var rpcPayload []Request
		err = json.Unmarshal(body, &rpcPayload)
		if err != nil {
			return nil, fmt.Errorf("could not parse batch json payload: %w, must conform to: %s", err, rpcReqSchema)
		}
		return rpcPayload, nil
	}
	rpcRequest := Request{}
	err = json.Unmarshal(body, &rpcRequest)
	if err != nil {
		return nil, fmt.Errorf("could not parse json payload: %w, must conform to: %s", err, rpcReqSchema)
	}
	return []Request{rpcRequest}, nil
}
