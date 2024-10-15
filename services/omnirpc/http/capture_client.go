package http

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/bytemap"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// CaptureClient is a mock client used for checking response values.
type CaptureClient struct {
	requests     []*CapturedRequest
	responseFunc MakeResponseFunc
}

// MakeResponseFunc is used for mocking responses.
type MakeResponseFunc func(c *CapturedRequest) (Response, error)

// NewCaptureClient creates  anew client for testing.
func NewCaptureClient(responseFunc MakeResponseFunc) *CaptureClient {
	return &CaptureClient{requests: []*CapturedRequest{}, responseFunc: responseFunc}
}

// Requests turns a list of sent requests. These are not mutation safe.
func (c *CaptureClient) Requests() []*CapturedRequest {
	return c.requests
}

// NewRequest creates a new request.
func (c *CaptureClient) NewRequest() Request {
	request := CapturedRequest{
		Client:        c,
		StringHeaders: make(map[string]string),
	}
	c.requests = append(c.requests, &request)
	return &request
}

// CapturedRequest stores all request data for testing.
type CapturedRequest struct {
	// ClientContains the capture client object
	Client *CaptureClient
	// Body is the request body
	Body []byte
	// Context is the request set by the client
	//nolint:containedctx
	Context context.Context
	// StringHeaders are headers set by SetHeader. Notably, this will not
	// include headers set by SetHeaderBytes
	StringHeaders map[string]string
	// StringHeaders are headers set by SetHeaderBytes. Notably, this will not
	// include headers set by SetHeader
	ByteHeaders bytemap.ByteSliceMap[[]byte]
	// RequestURI is the request uri bytes. Notably, this will not include
	// RequestURI's set by SetRequestURIBytes
	RequestURI string
	// RequestURIBytes is the request uri bytes. Notably, this will not include
	// RequestURI's set by SetRequestURI
	RequestURIBytes []byte
	// Metrics is the metrics handler
	Handler metrics.Handler
}

var _ Client = &CaptureClient{}

// SetBody stores the body for testing.
func (c *CapturedRequest) SetBody(body []byte) Request {
	_, span := c.Handler.Tracer().Start(
		c.Context,
		"SetBody",
		trace.WithAttributes(attribute.String("SetBody", common.Bytes2Hex(body))),
	)
	defer func() {
		metrics.EndSpan(span)
	}()
	c.Body = body
	return c
}

// SetContext stores the context for testing.
func (c *CapturedRequest) SetContext(ctx context.Context) Request {
	_, span := c.Handler.Tracer().Start(
		ctx,
		"SetContext",
	)
	span.AddEvent("SetContext")
	defer func() {
		metrics.EndSpan(span)
	}()
	c.Context = ctx
	return c
}

// SetHeader sets the header for testing.
func (c *CapturedRequest) SetHeader(key, value string) Request {
	_, span := c.Handler.Tracer().Start(
		c.Context,
		"SetHeader",
		trace.WithAttributes(attribute.String("SetHeader", key)),
		trace.WithAttributes(attribute.String("value", value)),
	)
	defer func() {
		metrics.EndSpan(span)
	}()
	c.StringHeaders[key] = value
	return c
}

// SetHeaderBytes sets header bytes for testing.
func (c *CapturedRequest) SetHeaderBytes(key, value []byte) Request {
	_, span := c.Handler.Tracer().Start(
		c.Context,
		"SetHeaderBytes",
		trace.WithAttributes(attribute.String("key", common.Bytes2Hex(key))),
		trace.WithAttributes(attribute.String("value", common.Bytes2Hex(value))),
	)
	defer func() {
		metrics.EndSpan(span)
	}()
	c.ByteHeaders.Put(key, value)
	return c
}

// SetRequestURI stores the request uri.
func (c *CapturedRequest) SetRequestURI(uri string) Request {
	_, span := c.Handler.Tracer().Start(
		c.Context,
		"SetRequestURI",
		trace.WithAttributes(attribute.String("uri", uri)),
	)
	defer func() {
		metrics.EndSpan(span)
	}()
	c.RequestURI = uri
	return c
}

// Do calls responseFunc for testing.
func (c *CapturedRequest) Do() (Response, error) {
	_, span := c.Handler.Tracer().Start(
		c.Context,
		"Do",
	)
	defer func() {
		metrics.EndSpan(span)
	}()

	resp, err := c.Client.responseFunc(c)
	if err != nil {
		return nil, err
	}

	span.SetAttributes(attribute.String("response", common.Bytes2Hex(resp.Body())))

	return resp, err
}

// WithMetrics sets the metrics handler.
func (c *CapturedRequest) WithMetrics(metrics metrics.Handler) Request {
	c.Handler = metrics
	return c
}

var _ Request = &CapturedRequest{}
