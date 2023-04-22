package processlog

import (
	"bufio"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"sync"
	"time"
)

const pipeBufferSize = 64

type bufferedPipe struct {
	io.ReadCloser
	io.WriteCloser
}

func newBufferedPipe() *bufferedPipe {
	r, w := io.Pipe()

	rb := bufio.NewReaderSize(r, pipeBufferSize)
	wb := &bufferedWriteCloser{
		Writer:    bufio.NewWriterSize(w, pipeBufferSize),
		closer:    w,
		closeChan: make(chan bool),
	}

	// use a timer to prevent deadlocks
	timer := time.NewTimer(1 * time.Second)

	go func() {
		for {
			select {
			case <-wb.closeChan:
				return
			case <-timer.C:
				_ = wb.Flush()
			}
		}
	}()

	return &bufferedPipe{
		ReadCloser:  io.NopCloser(rb),
		WriteCloser: wb,
	}
}

type bufferedWriteCloser struct {
	*bufio.Writer
	closer    io.Closer
	closeChan chan bool
}

func (bwc *bufferedWriteCloser) Close() error {
	err := bwc.Writer.Flush()
	if err != nil {
		return fmt.Errorf("could not flush buffered writer: %w", err)
	}
	err = bwc.closer.Close()
	if err != nil {
		return fmt.Errorf("could not close pipe: %w", err)
	}
	bwc.closeChan <- true
	return nil
}

// SplitStreams splits an input into multiple io.readers
// TODO this should return an object with an iterator to prevent reuse.
func SplitStreams(input io.Reader, splitCount int) (outputReaders []io.ReadCloser) {
	var outputWriteClosers []io.WriteCloser
	var outputWriters []io.Writer
	for i := 0; i < splitCount; i++ {
		pipe := newBufferedPipe()
		// add the reader to the output
		outputReaders = append(outputReaders, pipe.ReadCloser)
		// add writer to closer object so they can be closed
		outputWriteClosers = append(outputWriteClosers, pipe.WriteCloser)
		// add outputs to writers so they can be passed into io.multiwriter
		// io.writecloser can not be used for variadic passes
		outputWriters = append(outputWriters, pipe.WriteCloser)
	}

	go func() {
		// close all writers when the process ends
		defer func() {
			for _, writer := range outputWriteClosers {
				_ = writer.Close()
			}
		}()

		// for us to transpose these arguments, they need to be cast correctly in a loop
		mw := io.MultiWriter(outputWriters...)
		// copy the data into the multiwriter
		_, _ = io.Copy(mw, input)
	}()

	return outputReaders
}

// readLine returns a single line (without the ending \n)
// from an input buffered reader.
// An error is returned if there is an error with the
// buffered reader.
func readLine(r *bufio.Reader) ([]byte, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return ln, errors.Wrap(err, "could not read line")
}

// CombineStreams creates a combined stream of two io.readClosers
// this is commonly used for combining stdout and stderr.
// nolint: cyclop
func CombineStreams(ctx context.Context, inputs ...io.ReadCloser) (output chan []byte, errChan chan error) {
	output = make(chan []byte)
	errChan = make(chan error, len(inputs))

	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		go func(in io.ReadCloser) {
			defer func() {
				wg.Done()
				_ = in.Close()
			}()

			r := bufio.NewReader(in)
			ln, err := readLine(r)
			for err == nil {
				select {
				case <-ctx.Done():
					errChan <- ctx.Err()
					return
				default:
					buffer := make([]byte, len(ln))
					copy(buffer, ln)
					output <- buffer
					ln, err = readLine(r)
				}
			}

			errChan <- err
		}(in)
	}

	go func() {
		wg.Wait()
		close(output)
		close(errChan)
	}()

	return output, errChan
}
