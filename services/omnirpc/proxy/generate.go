package proxy

import "io"

// BodyReader is used for generating a mock of the request body that returns an error
//
//go:generate go run github.com/vektra/mockery/v2 --name BodyReader --output ./mocks --case=underscore
type BodyReader interface {
	io.ReadCloser
}
