package http

import (
	"context"
	"strings"
)

// Client contains a post client for interacting with json rpc servers.
type Client interface {
	// NewRequest creates a new request
	NewRequest() Request
}

// Request is a request builder.
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

// Response is a standardized response interface.
//
//go:generate go run github.com/vektra/mockery/v2 --name Response --output ./mocks --case=underscore
type Response interface {
	Body() []byte
	StatusCode() int
}

// ClientType is the client type to use
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=ClientType -linecomment
type ClientType uint16

const (
	// FastHTTP is the fast http client type.
	FastHTTP ClientType = 0 // FastHTTP
	// Resty is the resty client type.
	Resty ClientType = iota // Resty
)

// AllClientTypes is a list of all client types. Since we use stringer
// we can auto generate this at runtime.
var AllClientTypes []ClientType

// set all client types.
func init() {
	for i := 0; i < len(_ClientType_index); i++ {
		clientType := ClientType(i)
		AllClientTypes = append(AllClientTypes, clientType)
	}
}

// NewClient creates a client from the client type
// defaults to fast http.
func NewClient(clientType ClientType) Client {
	switch clientType {
	case FastHTTP:
		return NewFastHTTPClient()
	case Resty:
		return NewRestyClient()
	default:
		return NewRestyClient()
	}
}

// ClientTypeFromString returns a client type from a string.
func ClientTypeFromString(clientType string) ClientType {
	for _, rawType := range AllClientTypes {
		if strings.EqualFold(rawType.String(), clientType) {
			return rawType
		}
	}
	// use unknown type
	return ClientType(len(_ClientType_index) + 2)
}
