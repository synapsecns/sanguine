package rpcinfo

import "context"

// GetLatency gets the latency on a chain.
func GetLatency(ctx context.Context, rpcURL string) (l Result) {
	return getLatency(ctx, rpcURL)
}
