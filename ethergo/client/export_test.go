package client

import (
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
)

// ParseCalls exports parsecalls for testing.
func ParseCalls(calls []w3types.Caller) attribute.KeyValue {
	return parseCalls(calls)
}

func (c *clientImpl) GetMetrics() metrics.Handler {
	return c.tracing
}
