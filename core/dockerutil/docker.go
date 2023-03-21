package dockerutil

import (
	"context"
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/ory/dockertest/v3/docker"
	"github.com/synapsecns/sanguine/core/processlog"
	"io"
)

var logger = log.Logger("dockertest")

// TailContainerLogs tails the logs of a container and outputs them to a processlog file.
func TailContainerLogs(opts ...Option) error {
	stdoutReader, stdoutWriter := io.Pipe()
	stderrReader, stderrWriter := io.Pipe()

	logOpts := &logOptions{
		logOptions: docker.LogsOptions{
			Stderr:       true,
			Stdout:       true,
			Timestamps:   false,
			RawTerminal:  false,
			Since:        0,
			ErrorStream:  stderrWriter,
			OutputStream: stdoutWriter,
		},
	}

	for _, opt := range opts {
		if err := opt(logOpts); err != nil {
			return optionError{err}
		}
	}
	if err := ValidateOptions(opts...); err != nil {
		return optionError{err}
	}

	// TODO: rest of the options go here
	logInfo, err := processlog.StartLogs(buildProcessLogOpts(logOpts.ctx, stdoutReader, stderrReader, logOpts.processOptions)...)
	if err != nil {
		return fmt.Errorf("failed to get container logs: %w", err)
	}

	if logOpts.callback != nil {
		logOpts.callback(logOpts.ctx, logInfo)
	}

	closeChan := make(chan struct{})
	go func() {
		<-logOpts.ctx.Done()
		close(closeChan)
	}()

	errChan := make(chan error, 1)
	go func() {
		errChan <- logOpts.pool.Client.Logs(logOpts.logOptions)
	}()

	select {
	case <-closeChan:
		return fmt.Errorf("context canceled: %w", logOpts.ctx.Err())
	case err := <-errChan:
		return fmt.Errorf("failed to tail container logs: %w", err)
	}
}

// buildProcessLogOpts builds the processlog options for the given stdout and stderr readers.
// StdOut/StdErr will always override the set readers and emit a warning.
func buildProcessLogOpts(ctx context.Context, stdoutReader, stderrReader io.ReadCloser, opts []processlog.StdStreamLogArgsOption) (out []processlog.StdStreamLogArgsOption) {
	out = append(out, processlog.WithCtx(ctx))
	out = append(out, opts...)

	if processlog.HasReader(processlog.StdOut, opts...) {
		logger.Warn("overriding stdout reader, should not be set")
	}

	if processlog.HasReader(processlog.StdErr, opts...) {
		logger.Warn("overriding stderr reader, should not be set")
	}

	out = append(out, processlog.WithStdOut(stdoutReader))
	out = append(out, processlog.WithStdErr(stderrReader))
	return out
}
