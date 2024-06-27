package metrics

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	"go.opentelemetry.io/otel/attribute"
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

// SpanEventByName returns the value of the first event with the given name.
// it is a helper function for tests.
func SpanEventByName(stub tracetest.SpanStub, name string) *attribute.Value {
	for _, event := range stub.Events {
		if event.Name == name {
			return &event.Attributes[0].Value
		}
	}
	return nil
}

// SpanAttributeByName returns the value of the first attribute with the given name.
// it is a helper function for tests.
func SpanAttributeByName(stub tracetest.SpanStub, name string) *attribute.Value {
	for _, attr := range stub.Attributes {
		if attr.Key == attribute.Key(name) {
			return &attr.Value
		}
	}
	return nil
}

// SpanHasException returns true if the span has an exception event.
// it is a helper function for tests.
func SpanHasException(stub tracetest.SpanStub) bool {
	for _, event := range stub.Events {
		if event.Name == "exception" {
			return true
		}
	}
	return false
}
