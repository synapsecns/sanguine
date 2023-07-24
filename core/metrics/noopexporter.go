package metrics

import (
	"context"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/aggregation"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

// noOpExporter is a no-op metric exporter that prevents any metrics from being exported.
type noOpMetricExporter struct{}

func newNoOpExporter() metric.Exporter {
	return noOpMetricExporter{}
}

// Temporality returns the temporality given instrument kind.
func (n noOpMetricExporter) Temporality(kind metric.InstrumentKind) metricdata.Temporality {
	return metric.DefaultTemporalitySelector(kind)
}

// Aggregation returns the aggregation for the given instrument kind.
func (n noOpMetricExporter) Aggregation(kind metric.InstrumentKind) aggregation.Aggregation {
	return metric.DefaultAggregationSelector(kind)
}

// Export exporter (no-op).
func (n noOpMetricExporter) Export(_ context.Context, _ *metricdata.ResourceMetrics) error {
	return nil
}

// ForceFlush exporter (no-op).
func (n noOpMetricExporter) ForceFlush(_ context.Context) error {
	return nil
}

// Shutdown exporter (no-op).
func (n noOpMetricExporter) Shutdown(_ context.Context) error {
	return nil
}

var _ metric.Exporter = &noOpMetricExporter{}
