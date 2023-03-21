package processlog

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"time"
)

// PrintFunc is the function used to print data to the desired logger.
type PrintFunc func([]byte)

// stdStreamLogArgs define log inputs.
type stdStreamLogArgs struct {
	// StdOut is the standard out stream to split from
	StdOut io.ReadCloser
	// StdErr is the stderr stream to split from
	StdErr io.ReadCloser
	// PrintFunc is how the print function is returned
	PrintFunc PrintFunc
	// LogDir is the directory to log to
	LogDir string
	// LogFileName is the log file name prefix
	LogFileName string
	// LogFrequency is how often to write to the lgos
	LogFrequency time.Duration
	// Accumulator to accumulate logs to.
	Accumulator *bytes.Buffer
	// Ctx is the context. Note: cancellations may lead to errors in your Wait() function
	// nolint: containedctx
	Ctx context.Context
}

var (
	errCtxNil          = errors.New("context cannot be nil")
	errStdOutMustBeSet = errors.New("StdOut must be set")
	//nolint: errName
	errStdErrMustBeSet = errors.New("StdErr must be set")
	errLogDirRequired  = errors.New("log directory must be set")
)

// Validate validates the stdStreamLogArgs struct.
func (s *stdStreamLogArgs) Validate() error {
	if s.Ctx == nil {
		return errCtxNil
	}

	if s.StdOut == nil {
		return errStdOutMustBeSet
	}

	if s.StdErr == nil {
		return errStdErrMustBeSet
	}

	if s.LogDir == "" {
		return errLogDirRequired
	}

	return nil
}

// StdStreamLogArgsOption is a function that modifies a stdStreamLogArgs struct field.
type StdStreamLogArgsOption func(*stdStreamLogArgs)

// WithStdOut returns a function that modifies the StdOut field of a stdStreamLogArgs struct.
func WithStdOut(stdOut io.ReadCloser) StdStreamLogArgsOption {
	return func(args *stdStreamLogArgs) {
		args.StdOut = stdOut
	}
}

// WithStdErr returns a function that modifies the StdErr field of a stdStreamLogArgs struct.
func WithStdErr(stdErr io.ReadCloser) StdStreamLogArgsOption {
	return func(args *stdStreamLogArgs) {
		args.StdErr = stdErr
	}
}

// WithReader returns a function that modifies the StdOut or StdErr field of a stdStreamLogArgs struct.
func WithReader(readerType ReaderType, reader io.ReadCloser) StdStreamLogArgsOption {
	return func(args *stdStreamLogArgs) {
		switch readerType {
		case StdOut:
			WithStdOut(reader)(args)
		case StdErr:
			WithStdErr(reader)(args)
		}
	}
}

// WithPrintFunc returns a function that modifies the PrintFunc field of a stdStreamLogArgs struct.
func WithPrintFunc(printFunc PrintFunc) StdStreamLogArgsOption {
	return func(args *stdStreamLogArgs) {
		args.PrintFunc = printFunc
	}
}

// WithLogDir returns a function that modifies the LogDir field of a stdStreamLogArgs struct.
func WithLogDir(logDir string) StdStreamLogArgsOption {
	return func(args *stdStreamLogArgs) {
		args.LogDir = logDir
	}
}

// WithLogFileName returns a function that modifies the LogFileName field of a stdStreamLogArgs struct.
func WithLogFileName(logFileName string) StdStreamLogArgsOption {
	return func(args *stdStreamLogArgs) {
		args.LogFileName = logFileName
	}
}

// WithLogFrequency returns a function that modifies the LogFrequency field of a stdStreamLogArgs struct.
func WithLogFrequency(logFrequency time.Duration) StdStreamLogArgsOption {
	return func(args *stdStreamLogArgs) {
		args.LogFrequency = logFrequency
	}
}

// WithAccumulator returns a function that modifies the Accumulator field of a stdStreamLogArgs struct.
func WithAccumulator(accumulator *bytes.Buffer) StdStreamLogArgsOption {
	return func(args *stdStreamLogArgs) {
		args.Accumulator = accumulator
	}
}

// WithCtx returns a function that modifies the Ctx field of a stdStreamLogArgs struct.
func WithCtx(ctx context.Context) StdStreamLogArgsOption {
	return func(args *stdStreamLogArgs) {
		args.Ctx = ctx
	}
}

// ReaderType is the type of reader (stdout or stderr).
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=ReaderType -linecomment
type ReaderType uint8

const (
	// StdOut is the stdout reader type.
	StdOut ReaderType = iota // stdout
	// StdErr is the stderr reader type.
	StdErr ReaderType = iota // stderr
)

// HasReader returns true if the given reader type is set.
func HasReader(readerType ReaderType, opts ...StdStreamLogArgsOption) bool {
	args := makeRawArgs(opts)

	switch readerType {
	case StdOut:
		if args.StdOut != nil {
			return true
		}
	case StdErr:
		if args.StdErr != nil {
			return true
		}
	}
	return false
}

// makeArgs creates a new stdStreamLogArgs struct with the given options applied.
// It takes in a variadic list of StdStreamLogArgsOption functions and applies each of them to the new struct instance.
// It returns the modified struct instance.
func makeArgs(opts []StdStreamLogArgsOption) (_ *stdStreamLogArgs, err error) {
	args := makeRawArgs(opts)
	if err := args.Validate(); err != nil {
		return nil, err
	}
	return args, nil
}

// makeRawArgs creates a new stdStreamLogArgs struct with the given options applied.
// it performs no validation.
func makeRawArgs(opts []StdStreamLogArgsOption) *stdStreamLogArgs {
	args := &stdStreamLogArgs{}
	args.LogFileName = "test"
	args.LogFrequency = time.Second
	// not a problem unless it's missing from user too, we'll validate at the end
	args.LogDir, _ = os.MkdirTemp("", "")

	args.LogFileName = "log"
	args.PrintFunc = func(s []byte) {
		// do nothing
	}

	for _, opt := range opts {
		opt(args)
	}

	return args
}
