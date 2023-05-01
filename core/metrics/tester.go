package metrics

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"testing"
)

type testHandler struct {
	*baseHandler
	exporter *tracetest.InMemoryExporter
}

// TestHandler is a handler that can be used for testing traces.
type TestHandler interface {
	Handler
	// GetSpansByName returns all spans with the given name.
	GetSpansByName(name string) (spans []tracetest.SpanStub)
}

// NewTestTracer returns a new test tracer.
func NewTestTracer(ctx context.Context, tb testing.TB) TestHandler {
	tb.Helper()

	th := testHandler{}
	th.exporter = tracetest.NewInMemoryExporter()

	buildInfo := config.NewBuildInfo(config.DefaultVersion, config.DefaultCommit, "test", config.DefaultDate)

	th.baseHandler = newBaseHandler(buildInfo, tracesdk.WithSyncer(th.exporter), tracesdk.WithSampler(tracesdk.AlwaysSample()))
	err := th.baseHandler.Start(ctx)
	assert.Nil(tb, err)

	return &th
}

func (t *testHandler) Type() HandlerType {
	return Null
}

// GetSpansByName returns all spans with the given name.
func (t *testHandler) GetSpansByName(name string) (spans []tracetest.SpanStub) {
	allSpans := t.exporter.GetSpans()
	for _, span := range allSpans {
		if span.Name == name {
			spans = append(spans, span)
		}
	}
	return
}
