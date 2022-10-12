package types

import "github.com/goccy/go-json"

type IRPCRequest interface {
	// GetID gets the raw son contents of the id. This (as with all other raw messages)
	// will include quotes
	// underlying bytes are not mutation safe
	GetID() json.RawMessage
	// GetMethod gets the method name of the rpc call
	GetMethod() RPCMethod
	// GetParams gets the raw json contents of the params
	// underlying bytes are not mutation safe
	GetParams() []json.RawMessage
}
