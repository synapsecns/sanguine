package proxy

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
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
	hash [32]byte
}

// newRawResponse produces a response with a unique hash based on json
// regardless of formatting.
func newRawResponse(body []byte, url string) (*rawResponse, error) {
	// TODO: consider using a syncpool here
	var unmarshalled interface{}

	// TODO: see if there's a faster way to do this. Canonical json?
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

func (f *Forwarder) forwardRequest(ctx context.Context, endpoint, requestID string) (*rawResponse, error) {
	endpointURL, err := fasturl.ParseURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("could not parse endpoint (%s): %w", endpointURL, err)
	}

	allowedProtocols := []string{httpSchema, httpsSchema}

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
		SetHeaderBytes(http.ContentType, http.JsonType).
		SetHeaderBytes(http.Accept, http.JsonType).
		Do()

	if err != nil {
		return nil, fmt.Errorf("could not get response from %s: %w", endpoint, err)
	}

	rawResp, err := newRawResponse(resp.Body(), endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return rawResp, nil
}
