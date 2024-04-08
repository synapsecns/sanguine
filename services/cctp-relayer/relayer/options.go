package relayer

import "github.com/synapsecns/sanguine/ethergo/submitter"

// relayerOptions is a struct that holds the options for the relayer.
type relayerOptions struct {
	submitter submitter.TransactionSubmitter
}

// OptionsArgsOption is an option passed into the relayer.
type OptionsArgsOption func(options *relayerOptions)

// WithSubmitter sets the submitter for the relayer.
func WithSubmitter(txSubmitter submitter.TransactionSubmitter) OptionsArgsOption {
	return func(options *relayerOptions) {
		options.submitter = txSubmitter
	}
}

func makeOptions(opts []OptionsArgsOption) *relayerOptions {
	args := &relayerOptions{}
	for _, opt := range opts {
		opt(args)
	}
	return args
}
