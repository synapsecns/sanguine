package inventory

// balanceFetchOptions is is an underlying struct used for option fetching.
type balanceFetchOptions struct {
	shouldRefreshBalances bool
	skipDBCache           bool
}

// BalanceFetchArgOption is an option that can be passed into a balance fetch request.
// we do this to allow optional args.
type BalanceFetchArgOption func(options *balanceFetchOptions)

// ShouldRefreshBalances allows someone fetching balance(s) to skip the cache.
func ShouldRefreshBalances() BalanceFetchArgOption {
	return func(options *balanceFetchOptions) {
		options.shouldRefreshBalances = true
	}
}

// SkipDBCache allows someone fetching balance(s) to skip the cache.
func SkipDBCache() BalanceFetchArgOption {
	return func(options *balanceFetchOptions) {
		options.skipDBCache = true
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
