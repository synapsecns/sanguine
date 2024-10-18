package http

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// RestyClient is a resty client for making requests to the http client.
type RestyClient struct {
	client  *resty.Client
	handler metrics.Handler
}

// NewRestyClient creates a resty client.
// while much slower than fasthttp, this client requests context cancellation.
func NewRestyClient(handler metrics.Handler) Client {
	return &RestyClient{client: resty.New(), handler: handler}
}

type restyRequest struct {
	*resty.Request
	endpoint string
	handler  metrics.Handler
}

// NewRequest create a new request.
func (r RestyClient) NewRequest() Request {
	return &restyRequest{
		Request: r.client.R(),
		handler: r.handler,
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

func (r *restyRequest) Do() (_ Response, err error) {
	_, span := r.handler.Tracer().Start(
		r.Request.Context(),
		"Do",
		trace.WithAttributes(
			attribute.String("uri", r.endpoint),
			attribute.String("headers", fmt.Sprintf("%v", r.Request.Header)),
			attribute.String("body", r.Request.Body.(string)),
		),
	)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()
	resp, err := r.Request.Post(r.endpoint)
	if err != nil {
		return nil, fmt.Errorf("could not get response from %s: %w", r.endpoint, err)
	}
	return resp, nil
}

var _ Client = &RestyClient{}
