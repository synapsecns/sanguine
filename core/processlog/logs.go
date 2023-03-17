package processlog

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ipfs/go-log"
	"golang.org/x/sync/errgroup"
	"io"
	"os"
	"time"
)

const (
	// StdOutSuffix is the standard out file suffix.
	StdOutSuffix = "stdout.log"
	// StdErrSuffix is the standard err suffix.
	StdErrSuffix = "stderr.log"
	// CombinedSuffix is the combined file suffix.
	CombinedSuffix = "combined.log"
)

var logger = log.Logger("proccesslog")

// StartLogs starts the log process.
func StartLogs(opts ...StdStreamLogArgsOption) (_ LogMetadata, err error) {
	input, err := makeArgs(opts)
	if err != nil {
		return nil, err
	}

	li := &logMetadataImpl{
		logDir: input.LogDir,
	}

	g, ctx := errgroup.WithContext(input.Ctx)

	// split stdout so we can write to both log files (combined and separate) and stdout
	splitStdOut := SplitStreams(input.StdOut, 3)
	// split stderr so we can write to both log files (combined and separate) and stdout
	splitStdErr := SplitStreams(input.StdErr, 3)

	g.Go(func() error {
		writePrintFunc(ctx, input.PrintFunc, input.LogFrequency, splitStdOut[0], splitStdErr[0])
		return nil
	})

	handles, err := li.createFiles(input.LogFileName)
	if err != nil {
		return nil, fmt.Errorf("could not create files: %w", err)
	}
	g.Go(func() error {
		return pipeSingleOutput(handles.sdtdOut, splitStdOut[1])
	})

	g.Go(func() error {
		return pipeSingleOutput(handles.stdErr, splitStdErr[1])
	})

	g.Go(func() error {
		return pipeCombinedOutput(ctx, handles.combined, input.Accumulator, splitStdOut[2], splitStdErr[2])
	})

	return li, nil
}

func pipeCombinedOutput(ctx context.Context, file *os.File, acc *bytes.Buffer, readers ...io.ReadCloser) error {
	defer func() {
		_ = file.Close()
	}()
	combinedOut, errChan := CombineStreams(ctx, readers...)
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-errChan:
			if !errors.Is(err, io.EOF) && err != nil {
				logger.Errorf("error combining streams: %v", err)
			}
		case bts, ok := <-combinedOut:
			if !ok {
				return nil
			}
			newData := append(bts, []byte("\n")...)
			if acc != nil {
				acc.Write(newData)
			}

			_, err := file.Write(newData)
			if err != nil {
				return fmt.Errorf("could not write to combined file: %w", err)
			}
		}
	}
}

func pipeSingleOutput(file *os.File, reader io.ReadCloser) (err error) {
	defer func() {
		_ = file.Close()
	}()
	_, err = io.Copy(file, reader)
	if err != nil {
		return fmt.Errorf("error copying stream %w", err)
	}
	return nil
}

// createFiles creates the log files.
func (li *logMetadataImpl) createFiles(logFileName string) (_ *fileHandles, err error) {
	handles := fileHandles{}
	// create the stdout log file
	handles.sdtdOut, err = os.Create(fmt.Sprintf("%s/%s.%s", li.logDir, logFileName, StdOutSuffix))
	if err != nil {
		return nil, fmt.Errorf("error creating file %w", err)
	}

	li.stdoutFile = handles.sdtdOut.Name()

	// create the stderr file
	handles.stdErr, err = os.Create(fmt.Sprintf("%s/%s.%s", li.logDir, logFileName, StdErrSuffix))
	if err != nil {
		return nil, fmt.Errorf("error creating file %w", err)
	}
	li.stderrFile = handles.stdErr.Name()

	// create the combined file
	handles.combined, err = os.Create(fmt.Sprintf("%s/%s.%s", li.logDir, logFileName, CombinedSuffix))
	if err != nil {
		return nil, fmt.Errorf("error creating file %w", err)
	}

	li.combinedFile = handles.combined.Name()
	return &handles, nil
}

type fileHandles struct {
	sdtdOut  *os.File
	stdErr   *os.File
	combined *os.File
}

// writePrintFunc writes the stdout and stderr streams to the combined log file and stdout.
func writePrintFunc(ctx context.Context, printFunc PrintFunc, logFrequency time.Duration, stdOut, stdErr io.ReadCloser) {
	combined, errChan := CombineStreams(ctx, stdOut, stdErr)
	var text []byte
	lastOut := time.Now()

	for {
		select {
		case <-ctx.Done():
			return
		case err := <-errChan:
			if !errors.Is(err, io.EOF) {
				logger.Errorf("error combining streams: %s", err)
			}
		case output, ok := <-combined:
			if !ok {
				if len(text) > 0 {
					// this is effectively a flush
					printFunc(text)
				}
				return
			}

			//nolint: gosimple
			text = append(text, output...)
			if len(text) > 0 && time.Since(lastOut) > logFrequency {
				printFunc(text)
				text = []byte(nil)
			}
		}
	}
}
