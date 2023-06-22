package metrics

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/sdk/metric/aggregator/lastvalue"
)

type Gauge struct {
	aggregator lastvalue.Aggregator
	maxDelta   int64        // max number of ms between samples before triggering a datadog acttion
	action     func() error // function to call when maxDelta is exceeded (optional)
}

func NewGauge(maxDelta int64, action func() error) *Gauge {
	agg := lastvalue.New(1)
	if action == nil {
		// TODO add datadog otel integration here
	}
	return &Gauge{
		aggregator: agg[0],
		maxDelta:   maxDelta,
		action:     action,
	}
}

func (g *Gauge) Update(ctx context.Context, value uint64) error {
	lastValue, lastStored, err := g.aggregator.LastValue()
	if err != nil {
		return err
	}
	now := time.Now().Unix()
	if (now - lastStored.Unix()) > g.maxDelta {
		err := g.action()
		if err != nil {
			return err
		}
	}

	// Only update last value if its different
	if lastValue != value {
		err = g.aggregator.Update(ctx, value, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
