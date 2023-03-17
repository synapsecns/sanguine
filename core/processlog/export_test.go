// nolint: wrapcheck
package processlog

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"os"
	"time"
)

// ReadLine exports the readLine function for testing.
func ReadLine(r *bufio.Reader) ([]byte, error) {
	return readLine(r)
}

// CombinedPipe exports buffered pipe for testing
// In buffered pipe the Close() method being defined for both ReadCloser and WriteCloser in the CombinedPipe interface.
// Since the bufferedPipe struct embeds two separate ReadCloser and WriteCloser fields, Go cannot determine which Close() method to use when the CombinedPipe interface is implemented.
// To resolve this, we define separate methods and call close on both in our test.
type CombinedPipe interface {
	io.Reader
	io.Writer
	CloseReader() error
	CloseWriter() error
}

// NewBufferedPipe exports the newBufferedPipe function for testing.
func NewBufferedPipe() CombinedPipe {
	return newBufferedPipe()
}

func (bp *bufferedPipe) CloseReader() error {
	return bp.ReadCloser.Close()
}

func (bp *bufferedPipe) CloseWriter() error {
	return bp.WriteCloser.Close()
}

// MakeArgs exports the makeArgs function for testing.
// only returns the error.
func MakeArgs(opts []StdStreamLogArgsOption) error {
	_, err := makeArgs(opts)
	return err
}

// WritePrintFunc exports the writePrintFunc function for testing.
func WritePrintFunc(ctx context.Context, printFunc PrintFunc, logFrequency time.Duration, stdOut, stdErr io.ReadCloser) {
	writePrintFunc(ctx, printFunc, logFrequency, stdOut, stdErr)
}

// PipeSingleOutput exports the pipeSingleOutput function for testing.
func PipeSingleOutput(file *os.File, reader io.ReadCloser) (err error) {
	return pipeSingleOutput(file, reader)
}

// PipeCombinedOutput exports the pipeCombinedOutput function for testing.
func PipeCombinedOutput(ctx context.Context, file *os.File, acc *bytes.Buffer, readers ...io.ReadCloser) error {
	return pipeCombinedOutput(ctx, file, acc, readers...)
}
