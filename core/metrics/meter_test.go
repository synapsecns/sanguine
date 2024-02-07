package metrics_test

import (
	"context"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"testing"
)

// Recording an arbitrary event, such as recording block number.
//
// nolint:dupl
func TestCounter(t *testing.T) {
	mp := &metrics.NullMeterImpl{}

	counter, err := mp.NewCounter("test", "block_counter", "a block counter", "blocks")
	Nil(t, err)
	counter.Add(context.Background(), 1, metric.WithAttributeSet(
		attribute.NewSet(attribute.Int64("block_number", 30), attribute.Int64("chain_id", 1))),
	)
	counter.Add(context.Background(), 1, metric.WithAttributeSet(
		attribute.NewSet(attribute.Int64("block_number", 31), attribute.Int64("chain_id", 1))),
	)
	counter.Add(context.Background(), 1, metric.WithAttributeSet(
		attribute.NewSet(attribute.Int64("block_number", 32), attribute.Int64("chain_id", 1))),
	)
}

// Recording values in a histogram format, such the actual block numbers stored over time.
//
// nolint:dupl
func TestHistogram(t *testing.T) {
	mp := &metrics.NullMeterImpl{}

	histogram, err := mp.NewHistogram("test", "block_histogram", "a block counter", "blocks")
	Nil(t, err)
	histogram.Record(context.Background(), 30, metric.WithAttributeSet(
		attribute.NewSet(attribute.String("contract_address", "0x00"), attribute.Int64("chain_id", 1))),
	)
	histogram.Record(context.Background(), 31, metric.WithAttributeSet(
		attribute.NewSet(attribute.String("contract_address", "0x00"), attribute.Int64("chain_id", 1))),
	)
	histogram.Record(context.Background(), 32, metric.WithAttributeSet(
		attribute.NewSet(attribute.String("contract_address", "0x00"), attribute.Int64("chain_id", 1))),
	)
}
