package http

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
	handler  metrics.Handler
}

// NewRequest create a new request.
func (r RestyClient) NewRequest() Request {
	return &restyRequest{
		Request: r.client.R(),
	}
}

// SetHeaderBytes is a wrapper around SetHeadre for bytes.
func (r *restyRequest) SetHeaderBytes(key, value []byte) Request {
	_, span := r.handler.Tracer().Start(
		r.Request.Context(),
		"SetHeaderBytes",
		trace.WithAttributes(
			attribute.String("key", common.Bytes2Hex(key)),
			attribute.String("value", common.Bytes2Hex(value)),
		))
	defer func() {
		metrics.EndSpan(span)
	}()
	r.Request.SetHeader(string(key), string(value))
	return r
}

func (r *restyRequest) SetBody(body []byte) Request {
	_, span := r.handler.Tracer().Start(
		r.Request.Context(),
		"SetBody",
		trace.WithAttributes(attribute.String("body", common.Bytes2Hex(body))),
	)
	defer func() {
		metrics.EndSpan(span)
	}()
	r.Request.SetBody(body)
	return r
}

func (r *restyRequest) SetContext(ctx context.Context) Request {
	_, span := r.handler.Tracer().Start(
		r.Request.Context(),
		"SetContext",
	)
	span.AddEvent("SetContext")
	defer func() {
		metrics.EndSpan(span)
	}()
	r.Request.SetContext(ctx)
	return r
}

func (r *restyRequest) SetHeader(key, value string) Request {
	_, span := r.handler.Tracer().Start(
		r.Request.Context(),
		"SetHeader",
		trace.WithAttributes(
			attribute.String("SetHeader", key),
			attribute.String("value", value),
		))
	defer func() {
		metrics.EndSpan(span)
	}()
	r.Request.SetHeader(key, value)
	return r
}

func (r *restyRequest) SetRequestURI(uri string) Request {
	_, span := r.handler.Tracer().Start(
		r.Request.Context(),
		"SetRequestURI",
		trace.WithAttributes(attribute.String("uri", uri)),
	)
	defer func() {
		metrics.EndSpan(span)
	}()
	r.endpoint = uri
	return r
}

func (r *restyRequest) Do() (_ Response, err error) {
	_, span := r.handler.Tracer().Start(
		r.Request.Context(),
		"Do",
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

func (r *restyRequest) WithMetrics(metrics metrics.Handler) Request {
	r.handler = metrics
	return r
}

var _ Client = &RestyClient{}
