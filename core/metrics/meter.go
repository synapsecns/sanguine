package metrics

import (
	"fmt"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"time"
)

// InitMeter creates and sets a global meter provider.
func InitMeter(serviceName string, interval time.Duration) (*sdkmetric.MeterProvider, error) {
	// TODO configure exporter how we need.

	exporter, err := stdout.New()
	if err != nil {
		return nil, fmt.Errorf("creating exporter failed %w", err)
	}
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
	return mp, nil
}

// NewCounter creates a new meter counter instrument.
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/api.md#counter
func NewCounter(mp *sdkmetric.MeterProvider, meterName string, counterName string, desc string, units string) (metric.Int64Counter, error) {
	counter, err := (*mp).Meter(
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
func NewHistogram(mp *sdkmetric.MeterProvider, meterName string, histName string, desc string, units string) (metric.Int64Histogram, error) {
	histogram, err := (*mp).
		Meter(
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
