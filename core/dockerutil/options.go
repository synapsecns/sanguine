package dockerutil

import (
	"context"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/synapsecns/sanguine/core/processlog"
)

// Callback is a callback function that is called when a container log process is started.
type Callback func(ctx context.Context, metadata processlog.LogMetadata)

type logOptions struct {
	logOptions docker.LogsOptions
	//nolint: containedctx
	ctx            context.Context
	pool           *dockertest.Pool
	callback       Callback
	processOptions []processlog.StdStreamLogArgsOption
}

// Option is a function that configures a logOptions struct.
type Option func(*logOptions) error

// WithContext sets the context for the TailContainerLogs function.
func WithContext(ctx context.Context) Option {
	return func(opts *logOptions) error {
		if ctx == nil {
			return newContextError("context is nil")
		}
		opts.ctx = ctx
		return nil
	}
}

// WithResource sets the resource for the TailContainerLogs function.
func WithResource(resource *dockertest.Resource) Option {
	return func(opts *logOptions) error {
		if resource == nil {
			return newResourceError("resource is nil")
		}
		opts.logOptions.Container = resource.Container.ID
		return nil
	}
}

// WithPool sets the swap for the TailContainerLogs function.
func WithPool(pool *dockertest.Pool) Option {
	return func(opts *logOptions) error {
		if pool == nil {
			return newPoolError("swap is nil")
		}
		opts.pool = pool
		return nil
	}
}

// WithStderr sets the Stderr option for the TailContainerLogs function.
func WithStderr(stderr bool) Option {
	return func(opts *logOptions) error {
		opts.logOptions.Stderr = stderr
		return nil
	}
}

// WithStdout sets the Stdout option for the TailContainerLogs function.
func WithStdout(stdout bool) Option {
	return func(opts *logOptions) error {
		opts.logOptions.Stdout = stdout
		return nil
	}
}

// WithFollow sets the Follow option for the TailContainerLogs function.
func WithFollow(follow bool) Option {
	return func(opts *logOptions) error {
		opts.logOptions.Follow = follow
		return nil
	}
}

// WithTimestamps sets the Timestamps option for the TailContainerLogs function.
func WithTimestamps(timestamps bool) Option {
	return func(opts *logOptions) error {
		opts.logOptions.Timestamps = timestamps
		return nil
	}
}

// WithRawTerminal sets the RawTerminal option for the TailContainerLogs function.
func WithRawTerminal(rawTerminal bool) Option {
	return func(opts *logOptions) error {
		opts.logOptions.RawTerminal = rawTerminal
		return nil
	}
}

// WithSince sets the Since option for the TailContainerLogs function.
func WithSince(since int64) Option {
	return func(opts *logOptions) error {
		opts.logOptions.Since = since
		return nil
	}
}

// WithProcessLogOptions sets the process log options for the TailContainerLogs function.
func WithProcessLogOptions(customOpts ...processlog.StdStreamLogArgsOption) Option {
	return func(opts *logOptions) error {
		opts.processOptions = customOpts
		return nil
	}
}

// WithCallback sets the callback function for the TailContainerLogs function.
func WithCallback(callback Callback) Option {
	return func(opts *logOptions) error {
		if callback == nil {
			return newCallbackError("callback is not provided")
		}
		opts.callback = callback
		return nil
	}
}

// ValidateOptions validates the options for the TailContainerLogs function.
func ValidateOptions(opts ...Option) error {
	logOpts := &logOptions{}

	for _, opt := range opts {
		if err := opt(logOpts); err != nil {
			return optionError{err}
		}
	}

	if logOpts.ctx == nil {
		return optionError{newContextError("context is not provided")}
	}

	if logOpts.logOptions.Container == "" {
		return optionError{newResourceError("resource is not provided")}
	}

	if logOpts.pool == nil {
		return optionError{newPoolError("swap is not provided")}
	}

	return nil
}

type optionError struct {
	err error
}

func (e optionError) Error() string {
	return e.err.Error()
}

type contextError struct {
	errMsg string
}

func newContextError(errMsg string) contextError {
	return contextError{errMsg: errMsg}
}

func (e contextError) Error() string {
	return e.errMsg
}

type resourceError struct {
	errMsg string
}

func newResourceError(errMsg string) resourceError {
	return resourceError{errMsg: errMsg}
}

func (e resourceError) Error() string {
	return e.errMsg
}

type poolError struct {
	errMsg string
}

func newPoolError(errMsg string) poolError {
	return poolError{errMsg: errMsg}
}

func (e poolError) Error() string {
	return e.errMsg
}

type callbackError struct {
	errMsg string
}

func newCallbackError(errMsg string) callbackError {
	return callbackError{errMsg: errMsg}
}

func (e callbackError) Error() string {
	return e.errMsg
}
