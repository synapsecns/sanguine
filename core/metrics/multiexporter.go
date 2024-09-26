package metrics

import (
	"context"
	"fmt"
	"go.uber.org/multierr"
	"sync"
	"time"

	"go.opentelemetry.io/otel/sdk/trace"
)

// MultiExporter is an interface that allows exporting spans to multiple OTLP trace exporters.
type MultiExporter interface {
	trace.SpanExporter
	AddExporter(exporter trace.SpanExporter)
}

type multiExporter struct {
	exporters []trace.SpanExporter
}

// NewMultiExporter creates a new multi exporter that forwards spans to multiple OTLP trace exporters.
// It takes in one or more trace.SpanExporter instances and ensures that spans are sent to all of them.
// This is useful when you need to send trace data to multiple backends or endpoints.
func NewMultiExporter(exporters ...trace.SpanExporter) MultiExporter {
	return &multiExporter{
		exporters: exporters,
	}
}

const defaultTimeout = 30 * time.Second

// ExportSpans exports a batch of spans.
func (m *multiExporter) ExportSpans(parentCtx context.Context, ss []trace.ReadOnlySpan) error {
	return m.doParallel(parentCtx, func(ctx context.Context, exporter trace.SpanExporter) error {
		return exporter.ExportSpans(ctx, ss)
	})
}

func (m *multiExporter) doParallel(parentCtx context.Context, fn func(context.Context, trace.SpanExporter) error) error {
	ctx, cancel := context.WithTimeout(parentCtx, defaultTimeout)
	defer cancel()

	var wg sync.WaitGroup
	var errors []error
	var mu sync.Mutex

	wg.Add(len(m.exporters))
	for _, exporter := range m.exporters {
		go func(exporter trace.SpanExporter) {
			defer wg.Done()
			err := fn(ctx, exporter)
			if err != nil {
				mu.Lock()
				errors = append(errors, fmt.Errorf("error in doMultiple: %w", err))
				mu.Unlock()
			}
		}(exporter)
	}

	wg.Wait()
	if len(errors) > 0 {
		// nolint: wrapcheck
		return multierr.Combine(errors...)
	}

	return nil
}

// Shutdown notifies the exporter of a pending halt to operations.
func (m *multiExporter) Shutdown(ctx context.Context) error {
	return m.doParallel(ctx, func(ctx context.Context, exporter trace.SpanExporter) error {
		return exporter.Shutdown(ctx)
	})
}

// AddExporter adds an exporter to the multi exporter.
func (m *multiExporter) AddExporter(exporter trace.SpanExporter) {
	m.exporters = append(m.exporters, exporter)
}

var _ trace.SpanExporter = &multiExporter{}
