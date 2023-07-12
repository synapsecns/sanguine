package metrics

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/metric/embedded"
	"time"

	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

// Meter is an interface for counter and histogram.
type Meter interface {
	// NewCounter creates a new meter counter instrument.
	NewCounter(meterName string, counterName string, desc string, units string) (metric.Int64Counter, error)
	// NewHistogram creates a new meter histogram instrument.
	NewHistogram(meterName string, histName string, desc string, units string) (metric.Int64Histogram, error)
}

// MeterImpl is an implementation of the MeterProvider interface.
type MeterImpl struct {
	mp *sdkmetric.MeterProvider
}

// NewOtelMeter creates a new meter provider.
func NewOtelMeter(serviceName string, interval time.Duration, exporter sdkmetric.Exporter) (*MeterImpl, error) {
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
	)
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(resource),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(interval)),
		),
	)
	return &MeterImpl{mp: mp}, nil
}

// NewCounter creates a new meter counter instrument.
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/api.md#counter
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
type NullMeterImpl struct{}

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
