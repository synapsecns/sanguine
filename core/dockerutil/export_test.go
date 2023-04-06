package dockerutil

import (
	"context"
	"github.com/synapsecns/sanguine/core/processlog"
	"io"
)

// NewOptionError exports an optionError for testing.
func NewOptionError(err error) error {
	return optionError{err: err}
}

// NewContextError exports a contextError for testing.
func NewContextError(err string) error {
	return newContextError(err)
}

// NewResourceError exports a resourceError for testing.
func NewResourceError(err string) error {
	return newResourceError(err)
}

// NewPoolError exports a poolError for testing.
func NewPoolError(err string) error {
	return newPoolError(err)
}

// NewCallbackError exports a callbackError for testing.
func NewCallbackError(err string) error {
	return newCallbackError(err)
}

// BuildProcessLogOpts exports buildProcessLogOpts for testing.
func BuildProcessLogOpts(ctx context.Context, stdoutReader, stderrReader io.ReadCloser, opts []processlog.StdStreamLogArgsOption) (out []processlog.StdStreamLogArgsOption) {
	return buildProcessLogOpts(ctx, stdoutReader, stderrReader, opts)
}
