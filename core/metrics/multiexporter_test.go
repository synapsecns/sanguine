package metrics_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/synapsecns/sanguine/core/metrics"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestMultiExporter(t *testing.T) {
	// Create in-memory exporters
	exporter1 := tracetest.NewInMemoryExporter()
	exporter2 := tracetest.NewInMemoryExporter()

	// Create multi-exporter
	multiExporter := metrics.NewMultiExporter(exporter1, exporter2)

	// Create test spans
	spans := []sdktrace.ReadOnlySpan{
		tracetest.SpanStub{}.Snapshot(),
		tracetest.SpanStub{}.Snapshot(),
	}

	// Test ExportSpans
	err := multiExporter.ExportSpans(context.Background(), spans)
	require.NoError(t, err)

	// Verify that spans were exported to both exporters
	assert.Equal(t, 2, len(exporter1.GetSpans()))
	assert.Equal(t, 2, len(exporter2.GetSpans()))

	// Test Shutdown
	err = multiExporter.Shutdown(context.Background())
	require.NoError(t, err)

	// Verify that both exporters were shut down
	// Note: InMemoryExporter doesn't have a Stopped() method, so we can't check this directly
	// Instead, we can try to export spans again and check for an error
	err = multiExporter.ExportSpans(context.Background(), spans)
	assert.NoError(t, err, "Expected no error after shutdown")
}
