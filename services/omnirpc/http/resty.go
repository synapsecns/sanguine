package http

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// RestyClient is a resty client for making requests to the http client.
type RestyClient struct {
	client *resty.Client
}

// NewRestyClient creates a resty client.
// while much slower than fasthttp, this client requests context cancellation.
func NewRestyClient() Client {
	return &RestyClient{client: resty.New()}
}

type restyRequest struct {
	*resty.Request
	endpoint string
}

// NewRequest create a new request.
func (r RestyClient) NewRequest() Request {
	return &restyRequest{
		Request: r.client.R(),
	}
}

// SetHeaderBytes is a wrapper around SetHeadre for bytes.
func (r *restyRequest) SetHeaderBytes(key, value []byte) Request {
	r.Request.SetHeader(string(key), string(value))
	return r
}

func (r *restyRequest) SetBody(body []byte) Request {
	r.Request.SetBody(body)
	return r
}

func (r *restyRequest) SetContext(ctx context.Context) Request {
	r.Request.SetContext(ctx)
	return r
}

func (r *restyRequest) SetHeader(key, value string) Request {
	r.Request.SetHeader(key, value)
	return r
}

func (r *restyRequest) SetRequestURI(uri string) Request {
	r.endpoint = uri
	return r
}

func (r *restyRequest) Do() (Response, error) {
	//nolint: wrapcheck
	return r.Request.Post(r.endpoint)
}

var _ Client = &RestyClient{}
