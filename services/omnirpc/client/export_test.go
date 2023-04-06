package client

func (r *rpcOptions) Confirmations() int {
	return r.confirmations
}

type RPCOptions interface {
	Confirmations() int
}

var _ RPCOptions = &rpcOptions{}

// MakeOptions exports the makeOptions function for testing.
func MakeOptions(opts []OptionsArgsOption) RPCOptions {
	return makeOptions(opts)
}
