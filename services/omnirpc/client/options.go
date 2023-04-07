package client

// rpcOptions is a struct that holds the options for the RPC client.
type rpcOptions struct {
	confirmations int
}

// OptionsArgsOption is an option passed into the client.
type OptionsArgsOption func(options *rpcOptions)

// WithDefaultConfirmations sets the default number of confirmations.
func WithDefaultConfirmations(confirmations int) OptionsArgsOption {
	return func(options *rpcOptions) {
		options.confirmations = confirmations
	}
}

func makeOptions(opts []OptionsArgsOption) *rpcOptions {
	args := &rpcOptions{
		confirmations: 0,
	}
	for _, opt := range opts {
		opt(args)
	}
	return args
}
