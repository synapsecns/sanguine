package processlog

import (
	"bufio"
	"fmt"
	"io"
	"sync"
	"time"
)

const (
	pipeBufferSize  = 10
	inactivityFlush = 2 * time.Second // Flush interval if no activity
)

type bufferedPipe struct {
	readCloser  *bufferedReadCloser
	writeCloser io.WriteCloser
}

func newBufferedPipe() *bufferedPipe {
	r, w := io.Pipe()

	rb := bufio.NewReaderSize(r, pipeBufferSize)
	bp := &bufferedPipe{}
	wb := &bufferedWriteCloser{
		Writer: bufio.NewWriterSize(w, pipeBufferSize),
		closer: w,
		pipe:   bp,
	}

	bp.readCloser = &bufferedReadCloser{
		ReadCloser: io.NopCloser(rb),
		pipe:       bp,
	}
	bp.writeCloser = wb

	bp.readCloser.startFlushOnInactivity()

	return bp
}

type bufferedReadCloser struct {
	io.ReadCloser
	pipe  *bufferedPipe
	timer *time.Timer
	mu    sync.Mutex
}

func (brc *bufferedReadCloser) Read(p []byte) (int, error) {
	brc.mu.Lock()
	defer brc.mu.Unlock()

	n, err := brc.ReadCloser.Read(p)
	brc.timer.Reset(inactivityFlush)
	return n, err
}

func (brc *bufferedReadCloser) startFlushOnInactivity() {
	brc.timer = time.NewTimer(inactivityFlush)
	go func() {
		for range brc.timer.C {
			brc.mu.Lock()
			_ = brc.pipe.writeCloser.(*bufferedWriteCloser).Flush()
			brc.mu.Unlock()
			brc.timer.Reset(inactivityFlush)
		}
	}()
}

type bufferedWriteCloser struct {
	*bufio.Writer
	closer io.Closer
	pipe   *bufferedPipe
}

func (bwc *bufferedWriteCloser) Write(p []byte) (int, error) {
	n, err := bwc.Writer.Write(p)
	bwc.pipe.readCloser.timer.Reset(inactivityFlush)
	return n, err
}

func (bwc *bufferedWriteCloser) Close() error {
	bwc.pipe.readCloser.timer.Stop()
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
