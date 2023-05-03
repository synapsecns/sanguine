package client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lmittmann/w3"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
)

// captureClient is a wrapper around ethclient that can (but doesn't have to) captures requests and responses.
type captureClient struct {
	ethClient *ethclient.Client
	w3Client  *w3.Client
}

func newCaptureClient(ctx context.Context, url string, handler metrics.Handler, capture bool) (*captureClient, error) {
	client := new(http.Client)

	if capture {
		client.Transport = &captureTransport{
			transport: client.Transport,
			metrics:   handler,
		}
	}
	c, err := metrics.RPCClient(ctx, handler, url, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create rpc client: %w", err)
	}
	// capture config goes here

	ethClient := ethclient.NewClient(c)
	w3Client := w3.NewClient(c)

	return &captureClient{
		ethClient: ethClient,
		w3Client:  w3Client,
	}, nil
}

// TODO: test this, move it to metrics package
// captures requests and responses
// TODO: test.
type captureTransport struct {
	transport http.RoundTripper
	metrics   metrics.Handler
}

const (
	requestSpanName   = "http.request"
	requestEventName  = "http.request.body"
	responseEventName = "http.response.body"
)

// nolint: cyclop
func (t *captureTransport) RoundTrip(req *http.Request) (_ *http.Response, err error) {
	var response string

	_, span := t.metrics.Tracer().Start(req.Context(), requestSpanName)
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
		span.AddEvent(requestEventName, trace.WithAttributes(attribute.String("body", requestBody.String())))
	}
	if len(response) > 0 {
		span.AddEvent(responseEventName, trace.WithAttributes(attribute.String("body", response)))
	}

	//nolint: wrapcheck
	return resp, err
}
