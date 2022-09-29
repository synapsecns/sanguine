package http

import (
	"context"
	"github.com/synapsecns/sanguine/core/bytemap"
)

// CaptureClient is a mock client used for checking response values
type CaptureClient struct {
	requests     []*CapturedRequest
	responseFunc MakeResponseFunc
}

// MakeResponseFunc is used for mocking responses
type MakeResponseFunc func(c CapturedRequest) (Response, error)

func NewCaptureClient(responseFunc MakeResponseFunc) CaptureClient {
	return CaptureClient{requests: []*CapturedRequest{}, responseFunc: responseFunc}
}

func (c *CaptureClient) NewRequest() Request {
	request := CapturedRequest{
		Client: c,
	}
	c.requests = append(c.requests, &request)
	return request
}

// CapturedRequest stores all request data for testing
type CapturedRequest struct {
	// ClientContains the capture client object
	Client *CaptureClient
	// Body is the request body
	Body []byte
	// Context is the request set by the client
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
}

var _ Client = &CaptureClient{}

func (c CapturedRequest) SetBody(body []byte) Request {
	c.Body = body
	return c
}

func (c CapturedRequest) SetContext(ctx context.Context) Request {
	c.Context = ctx
	return c
}

func (c CapturedRequest) SetHeader(key, value string) Request {
	c.StringHeaders[key] = value
	return c
}

func (c CapturedRequest) SetHeaderBytes(key, value []byte) Request {
	c.ByteHeaders.Put(key, value)
	return c
}

func (c CapturedRequest) SetRequestURI(uri string) Request {
	c.RequestURI = uri
	return c
}

func (c CapturedRequest) Do() (Response, error) {
	//nolint: wrapcheck
	return c.Client.responseFunc(c)
}

var _ Request = &CapturedRequest{}
