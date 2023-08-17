package rpcinfo

import (
	"context"
	"github.com/synapsecns/sanguine/core/metrics"
)

// GetLatency gets the latency on a chain.
func GetLatency(ctx context.Context, rpcURL string) (l Result) {
	return getLatency(ctx, rpcURL, metrics.NewNullHandler())
}
