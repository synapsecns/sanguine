package metrics

import (
	"fmt"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"time"
)

type Meter interface {
	NewCounter(meterName string, counterName string, desc string, units string) (metric.Int64Counter, error)
	NewHistogram(meterName string, histName string, desc string, units string) (metric.Int64Histogram, error)
}

// MeterImpl is an implementation of the MeterProvider interface
type MeterImpl struct {
	mp *sdkmetric.MeterProvider
}

// NewMeter creates a new meter provider.
func NewMeter(serviceName string, interval time.Duration, exporter sdkmetric.Exporter) (*MeterImpl, error) {
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
