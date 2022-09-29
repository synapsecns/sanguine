package proxy

import "fmt"

// IsConfirmable exports isConfirmable for testing.
func IsConfirmable(body []byte) (bool, error) {
	parsedPayload, err := parseRPCPayload(body)
	if err != nil {
		return false, fmt.Errorf("could not parse payload: %w", err)
	}
	//nolint: wrapcheck
	return parsedPayload.isConfirmable()
}

// ParseRPCPayload exports parseRPCPayload for testing.
func ParseRPCPayload(body []byte) (_ *RPCRequest, err error) {
	//nolint: wrapcheck
	return parseRPCPayload(body)
}

// RawResponse exports rawResponse for testing.
type RawResponse interface {
	Body() []byte
	URL() string
	Hash() string
}

func (r rawResponse) Body() []byte {
	return r.body
}

func (r rawResponse) URL() string {
	return r.url
}

func (r rawResponse) Hash() string {
	return r.hash
}

var _ RawResponse = rawResponse{}
