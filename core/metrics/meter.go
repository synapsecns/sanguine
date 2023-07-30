package metrics

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
	"go.opentelemetry.io/otel/metric/noop"
)

// MeterProvider is an interface for creating and registering meters.
// It also allows the provider to be managed.
type MeterProvider interface {
	metric.MeterProvider
	// Shutdown shuts down the MeterProvider flushing all pending telemetry and
	// releasing any held computational resources.
	//
	// This call is idempotent. The first call will perform all flush and
	// releasing operations. Subsequent calls will perform no action and will
	// return an error stating this.
	//
	// Measurements made by instruments from meters this MeterProvider created
	// will not be exported after Shutdown is called.
	//
	// This method honors the deadline or cancellation of ctx. An appropriate
	// error will be returned in these situations. There is no guaranteed that all
	// telemetry be flushed or all resources have been released in these
	// situations.
	//
	// This method is safe to call concurrently.
	Shutdown(ctx context.Context) error
	// ForceFlush flushes all pending telemetry.
	//
	// This method honors the deadline or cancellation of ctx. An appropriate
	// error will be returned in these situations. There is no guaranteed that all
	// telemetry be flushed or all resources have been released in these
	// situations.
	//
	// This method is safe to call concurrently.
	ForceFlush(ctx context.Context) error
}

// Meter is an interface for counter and histogram.
// Deprecated: will be removed in a future version.
type Meter interface {
	// NewCounter creates a new meter counter instrument.
	NewCounter(meterName string, counterName string, desc string, units string) (metric.Int64Counter, error)
	// NewHistogram creates a new meter histogram instrument.
	NewHistogram(meterName string, histName string, desc string, units string) (metric.Int64Histogram, error)
}

// MeterImpl is an implementation of the MeterProvider interface.
type MeterImpl struct {
	mp MeterProvider
}

// NewOtelMeter creates a new meter provider.
// Deprecated: will be removed in a future version.
func NewOtelMeter(mp MeterProvider) *MeterImpl {
	return &MeterImpl{mp: mp}
}

// NewCounter creates a new meter counter instrument.
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/api.md#counter
// Deprecated: will be removed in a future version.
func (m *MeterImpl) NewCounter(meterName string, counterName string, desc string, units string) (metric.Int64Counter, error) {
	counter, err := m.mp.Meter(
		meterName,
	).
		Int64Counter(
			counterName,
			metric.WithDescription(desc),
			metric.WithUnit(units),
		)
	if err != nil {
		return nil, fmt.Errorf("creating counter failed %w", err)
	}
	return counter, nil
}

// NewHistogram creates a new meter histogram instrument.
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/api.md#histogram
// Deprecated: will be removed in a future version.
func (m *MeterImpl) NewHistogram(meterName string, histName string, desc string, units string) (metric.Int64Histogram, error) {
	histogram, err := m.mp.Meter(
		meterName,
	).Int64Histogram(
		histName,
		metric.WithDescription(desc),
		metric.WithUnit(units),
	)
	if err != nil {
		return nil, fmt.Errorf("creating histogram failed %w", err)
	}
	return histogram, nil
}

// NullMeterImpl is a no-op implementation of the Meter interface.
// Deprecated: will be removed in a future version.
type NullMeterImpl struct{}

// Meter providees a no-op implementation of the Meter interface.
func (m *NullMeterImpl) Meter() metric.Meter {
	return noop.Meter{}
}

// NewCounter creates a new meter counter instrument.
func (m *NullMeterImpl) NewCounter(_ string, _ string, _ string, _ string) (metric.Int64Counter, error) {
	return &NullCounter{}, nil
}

// NewHistogram creates a new meter histogram instrument.
func (m *NullMeterImpl) NewHistogram(_ string, _ string, _ string, _ string) (metric.Int64Histogram, error) {
	return &NullHistogram{}, nil
}

// NullCounter is a no-op implementation of the metric.Int64Counter.
type NullCounter struct {
	embedded.Int64Counter
}

// Add is a no-op implementation of Int64Counter Add() function.
func (n *NullCounter) Add(_ context.Context, _ int64, _ ...metric.AddOption) {}

// NullHistogram is a no-op implementation of the metric.Int64Histogram.
type NullHistogram struct {
	embedded.Int64Histogram
}

// Record is a no-op implementation of RecordOption Record() function.
func (n *NullHistogram) Record(_ context.Context, _ int64, _ ...metric.RecordOption) {}
