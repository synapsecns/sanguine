package core

import "context"

// Validator defines an interface used for validating a struct.
type Validator interface {
	// IsValid determines whether the config is valid.
	IsValid(ctx context.Context) (ok bool, err error)
}
