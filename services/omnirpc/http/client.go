package http

import "context"

type Client interface {
	// NewRequest creates a new request
	NewRequest() Request
}

type Request interface {
	// SetBody sets the request body
	SetBody(body []byte) Request
	// SetContext sets the context for the request
	SetContext(ctx context.Context) Request
	// SetHeader sets the header for the client
	SetHeader(key, value string) Request
	// SetHeaderBytes sets header in bytes to avoid a copy
	SetHeaderBytes(key, value []byte) Request
	// SetRequestURI sets the uri for the request
	SetRequestURI(uri string) Request
	// Do makes the actual request
	Do() (Response, error)
}

type Response interface {
	Body() []byte
}
