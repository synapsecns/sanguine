package processlog

import (
	"bufio"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"sync"
)

const pipeBufferSize = 10

type bufferedPipe struct {
	io.ReadCloser
	io.WriteCloser
}

func newBufferedPipe() *bufferedPipe {
	r, w := io.Pipe()

	rb := bufio.NewReaderSize(r, pipeBufferSize)
	wb := &bufferedWriteCloser{
		Writer: bufio.NewWriterSize(w, pipeBufferSize),
		closer: w,
	}

	return &bufferedPipe{
		ReadCloser:  io.NopCloser(rb),
		WriteCloser: wb,
	}
}

type bufferedWriteCloser struct {
	*bufio.Writer
	closer io.Closer
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
	return nil
}

// SplitStreams splits an input into multiple io.readers
// TODO this should return an object with an iterator to prevent reuse.
func SplitStreams(input io.Reader, splitCount int) (outputReaders []io.ReadCloser) {
	// Create a channel to send data from the input reader to the output writers
	dataChan := make(chan []byte, splitCount)

	// Create the output writers
	outputWriters := make([]io.WriteCloser, splitCount)
	outputPipes := make([]*bufferedPipe, splitCount)
	for i := 0; i < splitCount; i++ {
		pipe := newBufferedPipe()
		// Add the reader to the output
		outputReaders = append(outputReaders, pipe.ReadCloser)
		// Add the writer to the output writers
		outputWriters[i] = pipe.WriteCloser
		outputPipes[i] = pipe
	}

	// Start a goroutine to read from the input reader and send data to the output writers
	go func() {
		// Create a buffer to hold data read from the input reader
		buf := make([]byte, pipeBufferSize)

		// Loop until we reach the end of the input reader or an error occurs
		for {
			n, err := input.Read(buf)
			if n > 0 {
				// Copy the data to each output writer
				data := make([]byte, n)
				copy(data, buf[:n])
				for _, writer := range outputWriters {
					_, err := writer.Write(data)
					if err != nil {
						fmt.Printf("error writing data to output writer: %v\n", err)
						return
					}
				}
			}
			if err != nil {
				if err != io.EOF {
					fmt.Printf("error reading data from input reader: %v\n", err)
				}
				// Close the data channel when we're done sending data
				return
			}
		}
	}()

	// Start a goroutine for each output writer to read from its corresponding output reader
	var wg sync.WaitGroup
	for _, pipe := range outputPipes {
		wg.Add(1)
		go func(pipe *bufferedPipe) {
			// Create a buffer to hold data read from the output reader
			buf := make([]byte, pipeBufferSize)
			for {
				// Read data from the output reader
				n, err := pipe.ReadCloser.Read(buf)
				if n > 0 {
					// Send the data to the data channel
					data := make([]byte, n)
					copy(data, buf[:n])
					dataChan <- data
				}
				if err != nil {
					if err != io.EOF {
						fmt.Printf("error reading data from output reader: %v\n", err)
					}
					wg.Done()
					return
				}
			}
		}(pipe)
	}

	// Wait for all the output writers to finish writing before closing the data channel
	go func() {
		wg.Wait()
		close(dataChan)
	}()

	// Return a slice of ReadClosers that read from the data channel
	for i := 0; i < splitCount; i++ {
		outputReaders[i] = newSliceReader(dataChan)
	}

	return outputReaders
}

// newSliceReader returns a ReadCloser that reads data from a channel of byte slices
func newSliceReader(dataChan <-chan []byte) io.ReadCloser {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		for data := range dataChan {
			_, err := pw.Write(data)
			if err != nil {
				fmt.Printf("error writing data to pipe: %v\n", err)
				return
			}
		}
	}()
	return pr
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
