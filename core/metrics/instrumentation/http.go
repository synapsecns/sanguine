package instrumentation

import (
	"bytes"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
)

// NewCaptureTransport returns a new http.Transport that captures requests and responses.
// TODO: add tests.
func NewCaptureTransport(underlyingTransport http.RoundTripper, handler metrics.Handler) http.RoundTripper {
	return &captureTransport{
		transport: underlyingTransport,
		metrics:   handler,
	}
}

// captures requests and responses.
type captureTransport struct {
	transport http.RoundTripper
	metrics   metrics.Handler
}

const (
	// RequestSpanName is the name of the span created for each request.
	RequestSpanName = "http.request"
	// RequestEventName is the name of the event created for each request body.
	RequestEventName = "http.request.body"
	// ResponseEventName is the name of the event created for each response body.
	ResponseEventName = "http.response.body"
)

// nolint: cyclop
func (t *captureTransport) RoundTrip(req *http.Request) (_ *http.Response, err error) {
	var response string

	_, span := t.metrics.Tracer().Start(req.Context(), RequestSpanName)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// Perform the HTTP request using the underlying transport
	transport := t.transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	// no need to capture if we're not recording
	if !span.IsRecording() {
		// nolint: wrapcheck
		return transport.RoundTrip(req)
	}

	// Capture the request body
	var requestBody bytes.Buffer
	if req.Body != nil {
		tee := io.TeeReader(req.Body, &requestBody)
		req.Body = io.NopCloser(tee)
		if req.ContentLength == 0 && req.Header.Get("Content-Encoding") == "" {
			req.ContentLength = int64(requestBody.Len())
		}
	}

	resp, err := transport.RoundTrip(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform http request: %w", err)
	}

	// Capture the response body
	//nolint: nestif
	if resp != nil && resp.Body != nil {
		var responseBody bytes.Buffer
		if resp.ContentLength == 0 && resp.Header.Get("Content-Encoding") == "" {
			// Copy the response body to a buffer so we can compute the content length
			tee := io.TeeReader(resp.Body, &responseBody)
			resp.Body = io.NopCloser(tee)
			bodyBytes, err := io.ReadAll(&responseBody)
			if err == nil {
				response = string(bodyBytes)
			}
			resp.ContentLength = int64(len(bodyBytes))
		} else {
			// Copy the response body to a buffer so we can capture the response data
			bodyBytes, err := io.ReadAll(resp.Body)
			if err == nil {
				response = string(bodyBytes)
				resp.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			}
		}
	}

	// Add the request to the list of captured requests

	// Add the request/response body as events to the span
	if requestBody.Len() > 0 {
		span.AddEvent(RequestEventName, trace.WithAttributes(attribute.String("body", requestBody.String())))
	}
	if len(response) > 0 {
		span.AddEvent(ResponseEventName, trace.WithAttributes(attribute.String("body", response)))
	}

	//nolint: wrapcheck
	return resp, err
}
