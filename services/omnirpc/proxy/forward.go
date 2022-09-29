package proxy

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/ImVexed/fasturl"
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
}

// newRawResponse produces a response with a unique hash based on json
// regardless of formatting.
func (f *Forwarder) newRawResponse(body []byte, url string) (*rawResponse, error) {
	// TODO: see if there's a faster way to do this. Canonical json?
	// unmarshall and remarshall
	standardizedResponse, err := StandardizeResponse(f.rpcRequest.Method, body)
	if err != nil {
		return nil, fmt.Errorf("could not standardize response: %w", err)
	}

	return &rawResponse{
		body: body,
		url:  url,
		hash: fmt.Sprintf("%x", sha256.Sum256(standardizedResponse)),
	}, nil
}

const (
	httpSchema  = "http"
	httpsSchema = "https"
)

func (f *Forwarder) forwardRequest(ctx context.Context, endpoint, requestID string) (*rawResponse, error) {
	endpointURL, err := fasturl.ParseURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("could not parse endpoint (%s): %w", endpointURL, err)
	}

	allowedProtocols := []string{httpsSchema, httpsSchema}

	// websockets not yet supported
	if !slices.Contains(allowedProtocols, endpointURL.Protocol) {
		return nil, fmt.Errorf("schema must be one of %s, got %s", strings.Join(allowedProtocols, ","), endpointURL.Protocol)
	}

	req := f.client.NewRequest()
	resp, err := req.
		SetContext(ctx).
		SetRequestURI(endpoint).
		SetBody(f.body).
		SetHeaderBytes(http.XRequestID, []byte(requestID)).
		SetHeaderBytes(http.XForwardedFor, http.OmniRPCValue).
		SetHeaderBytes(http.ContentType, http.JSONType).
		SetHeaderBytes(http.Accept, http.JSONType).
		Do()

	if err != nil {
		return nil, fmt.Errorf("could not get response from %s: %w", endpoint, err)
	}

	rawResp, err := f.newRawResponse(resp.Body(), endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return rawResp, nil
}
