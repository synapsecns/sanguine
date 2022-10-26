package proxy

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/ImVexed/fasturl"
	"github.com/goccy/go-json"
	"github.com/synapsecns/sanguine/services/omnirpc/http"
	"golang.org/x/exp/slices"
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
func (f *Forwarder) newRawResponse(ctx context.Context, body []byte, url string) (*rawResponse, error) {
	// TODO: see if there's a faster way to do this. Canonical json?
	// unmarshall and remarshall
	var rpcMessage JSONRPCMessage
	err := json.Unmarshal(body, &rpcMessage)
	if err != nil {
		return nil, fmt.Errorf("could not parse response: %w", err)
	}

	standardizedResponse, err := standardizeResponse(ctx, *f.rpcRequest, rpcMessage)
	if err != nil {
		return nil, fmt.Errorf("could not standardize response: %w", err)
	}

	return &rawResponse{
		body:     body,
		url:      url,
		hash:     fmt.Sprintf("%x", sha256.Sum256(standardizedResponse)),
		hasError: rpcMessage.Error != nil,
	}, nil
}

const (
	httpSchema  = "http"
	httpsSchema = "https"
)

func (f *Forwarder) forwardRequest(ctx context.Context, endpoint string) (*rawResponse, error) {
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
		return nil, fmt.Errorf("invalid response code: %w", err)
	}

	rawResp, err := f.newRawResponse(ctx, resp.Body(), endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return rawResp, nil
}
