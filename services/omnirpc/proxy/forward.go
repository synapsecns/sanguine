package proxy

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"golang.org/x/exp/slices"
	urlParser "net/url"
	"strings"
)

type rawResponse struct {
	// body is the raw body returned by the request
	body []byte
	// url is the request url
	url string
	// hash is a unique hash of the raw response.
	// we use this to check for equality
	hash [32]byte
}

// newRawResponse produces a response with a unique hash based on json
// regardless of formatting.
func newRawResponse(body []byte, url string) (*rawResponse, error) {
	var unmarshalled interface{}

	// unmarshall and remarshall
	decoder := json.NewDecoder(bytes.NewReader(body))
	err := decoder.Decode(&unmarshalled)
	if err != nil {
		return nil, fmt.Errorf("could not decode json: %w", err)
	}

	remarshalled, err := json.Marshal(unmarshalled)
	if err != nil {
		return nil, fmt.Errorf("could not re-encode json: %w", err)
	}

	return &rawResponse{
		body: body,
		url:  url,
		hash: sha256.Sum256(remarshalled),
	}, nil
}

const (
	httpSchema  = "http"
	httpsSchema = "https"
)

func forwardRequest(ctx context.Context, body []byte, endpoint, header string) (*rawResponse, error) {
	endpointURL, err := urlParser.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("could not parse endpoint (%s): %w", endpointURL, err)
	}

	allowedProtocols := []string{httpSchema, httpsSchema}

	// websockets not yet supported
	if !slices.Contains(allowedProtocols, endpointURL.Scheme) {
		return nil, fmt.Errorf("schema must be one of %s, got %s", strings.Join(allowedProtocols, ","), endpointURL.Scheme)
	}

	client := resty.New()
	resp, err := client.R().
		SetContext(ctx).
		SetBody(body).
		SetHeader("x-forwarded-for", "omnirpc").
		SetHeader(requestIDKey, header).
		Post(endpoint)

	if err != nil {
		return nil, fmt.Errorf("could not get response from %s: %w", endpoint, err)
	}

	rawResp, err := newRawResponse(resp.Body(), endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return rawResp, nil
}
