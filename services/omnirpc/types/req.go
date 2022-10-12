package types

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/invopop/jsonschema"
)

// RPCRequest is a raw rpc request format.
type RPCRequest struct {
	ID     json.RawMessage   `json:"id"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}

func (r RPCRequest) GetID() json.RawMessage {
	return r.ID
}

func (r RPCRequest) GetMethod() RPCMethod {
	return RPCMethod(r.Method)
}

func (r RPCRequest) GetParams() []json.RawMessage {
	return r.Params
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

func ParseRPCPayload(body []byte) (_ *RPCRequest, err error) {
	rpcPayload := RPCRequest{}
	err = json.Unmarshal(body, &rpcPayload)
	if err != nil {
		return nil, fmt.Errorf("could not parse json payload: %w, must conform to: %s", err, rpcReqSchema)
	}

	return &rpcPayload, nil
}

var _ IRPCRequest = &RPCRequest{}
