package inventory

// balanceFetchOptions is is an underlying struct used for option fetching.
type balanceFetchOptions struct {
	skipCache bool
}

// BalanceFetchArgOption is an option that can be passed into a balance fetch request.
// we do this to allow optional args.
type BalanceFetchArgOption func(options *balanceFetchOptions)

// SkipCache allows someone fetching balance(s) to skip the cache.
func SkipCache() BalanceFetchArgOption {
	return func(options *balanceFetchOptions) {
		options.skipCache = true
	}
}

// makeOptions creates the balance fetch options.
func makeOptions(opts []BalanceFetchArgOption) *balanceFetchOptions {
	args := &balanceFetchOptions{}
	for _, opt := range opts {
		opt(args)
	}
	return args
}
