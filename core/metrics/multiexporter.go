package metrics

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// MultiExporter is an interface that allows exporting spans to multiple OTLP trace exporters.
type MultiExporter interface {
	tracesdk.SpanExporter
	AddExporter(exporter *otlptrace.Exporter)
}

type multiExporter struct {
	exporters []*otlptrace.Exporter
}

// NewMultiExporter creates a new multi exporter that forwards spans to multiple OTLP trace exporters.
// It takes in one or more otlptrace.Exporter instances and ensures that spans are sent to all of them.
// This is useful when you need to send trace data to multiple backends or endpoints.
func NewMultiExporter(exporters ...*otlptrace.Exporter) MultiExporter {
	return &multiExporter{
		exporters: exporters,
	}
}

// ExportSpans exports a batch of spans.
func (m *multiExporter) ExportSpans(ctx context.Context, ss []tracesdk.ReadOnlySpan) error {
	for _, exporter := range m.exporters {
		err := exporter.ExportSpans(ctx, ss)
		if err != nil {
			return fmt.Errorf("could not export spans: %w", err)
		}
	}
	return nil
}

// Shutdown notifies the exporter of a pending halt to operations.
func (m *multiExporter) Shutdown(ctx context.Context) error {
	for _, exporter := range m.exporters {
		err := exporter.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("could not stop exporter: %w", err)
		}
	}
	return nil
}

// AddExporter adds an exporter to the multi exporter.
func (m *multiExporter) AddExporter(exporter *otlptrace.Exporter) {
	m.exporters = append(m.exporters, exporter)
}

var _ tracesdk.SpanExporter = &multiExporter{}
