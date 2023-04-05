package metrics

import "go.opentelemetry.io/otel/trace"

// EndSpanWithErr ends a span and records an error if one is present.
func EndSpanWithErr(span trace.Span, err error) {
	if span == nil {
		return
	}

	if err != nil {
		span.RecordError(err)
	}

	span.End()
}

// EndSpan ends a span.
func EndSpan(span trace.Span) {
	if span == nil {
		return
	}

	span.End()
}
