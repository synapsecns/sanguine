package metrics

import (
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"time"
)

func InitMeter(serviceName string, interval time.Duration) error {
	// TODO configure exporter how we need

	exporter, err := stdout.New()
	if err != nil {
		return err
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
	otel.SetMeterProvider(mp)
	return nil
}

func NewCounter(meterName string, counterName string, desc string, units string) (metric.Int64Counter, error) {
	counter, err := otel.GetMeterProvider().
		Meter(
			meterName,
		).
		Int64Counter(
			counterName,
			metric.WithDescription(desc),
			metric.WithUnit(units),
		)
	if err != nil {
		return nil, err
	}
	return counter, nil
}

func NewHistogram(meterName string, histName string, desc string, units string) (metric.Int64Histogram, error) {
	counter, err := otel.GetMeterProvider().
		Meter(
			meterName,
		).Int64Histogram(
		histName,
		metric.WithDescription(desc),
		metric.WithUnit(units),
	)
	if err != nil {
		return nil, err
	}
	return counter, nil
}
