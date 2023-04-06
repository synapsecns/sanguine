package proxy

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/ImVexed/fasturl"
	"github.com/goccy/go-json"
	"github.com/jftuga/ellipsis"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
	"github.com/synapsecns/sanguine/services/omnirpc/http"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slices"
	goHTTP "net/http"
	"strings"
)

type rawResponse struct {
	// body is the raw body returned by the request
	body []byte
	// url is the request url
	url string
	// hash is a unique hash of the raw response.
	// we use this to check for equality
	hash string
	// hasError is wether or not the response could be deserialized
	hasError bool
}

// newRawResponse produces a response with a unique hash based on json
// regardless of formatting.
func (f *Forwarder) newRawResponse(ctx context.Context, body []byte, url string) (_ *rawResponse, err error) {
	// TODO: see if there's a faster way to do this. Canonical json?
	// TODO: standardize batch request
	// unmarshall and remarshall

	var standardizedResponse []byte
	var hasErr bool

	if rpc.IsBatch(body) {
		standardizedResponse, hasErr, err = f.standardizeBatch(ctx, body)
		if err != nil {
			return nil, fmt.Errorf("could not standardize batch response: %w", err)
		}
	} else {
		var rpcMessage JSONRPCMessage
		err := json.Unmarshal(body, &rpcMessage)
		if err != nil {
			return nil, fmt.Errorf("could not parse response %s: %w", body, err)
		}

		hasErr = rpcMessage.Error != nil

		standardizedResponse, err = standardizeResponse(ctx, &f.rpcRequest[0], rpcMessage)
		if err != nil {
			return nil, fmt.Errorf("could not standardize response: %w", err)
		}
	}

	return &rawResponse{
		body:     body,
		url:      url,
		hash:     fmt.Sprintf("%x", sha256.Sum256(standardizedResponse)),
		hasError: hasErr,
	}, nil
}

// standardizes a batch request. anyErr indicates *any* response in the batch had an error
// (not at the decoding step).
func (f *Forwarder) standardizeBatch(ctx context.Context, body []byte) (res []byte, anyErr bool, err error) {
	dec := json.NewDecoder(bytes.NewReader(body))

	_, err = dec.Token() // skip '['
	if err != nil {
		return nil, true, fmt.Errorf("could not decode %s: %w", ellipsis.Shorten(string(body), 10), err)
	}

	// response.id->response
	responses := make(map[int]json.RawMessage)

	i := 0
	for dec.More() {
		response := new(JSONRPCMessage)
		err = dec.Decode(&response)
		if err != nil {
			return nil, true, fmt.Errorf("could not decode response at index %d: %w", i, err)
		}

		if response.Error != nil {
			anyErr = true
		}

		request := f.rpcRequest.ByID(response.ID)
		if request == nil {
			return nil, true, fmt.Errorf("no request found for id %d", response.ID)
		}

		standardized, err := standardizeResponse(ctx, f.rpcRequest.ByID(response.ID), *response)
		if err != nil {
			return nil, true, fmt.Errorf("could not decode response at index %d: %w", i, err)
		}

		responses[response.ID] = standardized
		i++
	}

	// create the return array
	var standardizedResponses []json.RawMessage

	// return the responses in the order they were requested
	for _, request := range f.rpcRequest {
		standardizedResponses = append(standardizedResponses, responses[request.ID])
	}

	standardizedResponse, err := json.Marshal(standardizedResponses)
	if err != nil {
		return nil, true, fmt.Errorf("could not unmarshall responses: %w", err)
	}

	return standardizedResponse, anyErr, nil
}

const (
	httpSchema  = "http"
	httpsSchema = "https"
)

func (f *Forwarder) forwardRequest(parentCtx context.Context, endpoint string) (_ *rawResponse, err error) {
	ctx, span := f.tracer.Start(parentCtx, "forwardRequest",
		trace.WithAttributes(attribute.String("endpoint", endpoint)),
	)

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	endpointURL, err := fasturl.ParseURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("could not parse endpoint (%s): %w", endpointURL, err)
	}

	allowedProtocols := []string{httpsSchema, httpSchema}

	// websockets not yet supported
	if !slices.Contains(allowedProtocols, endpointURL.Protocol) {
		return nil, fmt.Errorf("schema must be one of %s, got %s", strings.Join(allowedProtocols, ","), endpointURL.Protocol)
	}

	req := f.client.NewRequest()
	resp, err := req.
		SetContext(ctx).
		SetRequestURI(endpoint).
		SetBody(f.body).
		SetHeaderBytes(http.XRequestID, f.requestID).
		SetHeaderBytes(http.XForwardedFor, http.OmniRPCValue).
		SetHeaderBytes(http.ContentType, http.JSONType).
		SetHeaderBytes(http.Accept, http.JSONType).
		Do()
	if err != nil {
		return nil, fmt.Errorf("could not get response from %s: %w", endpoint, err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() > 400 {
		return nil, fmt.Errorf("invalid response code: %d (%s)", resp.StatusCode(), goHTTP.StatusText(resp.StatusCode()))
	}

	rawResp, err := f.newRawResponse(ctx, resp.Body(), endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return rawResp, nil
}
