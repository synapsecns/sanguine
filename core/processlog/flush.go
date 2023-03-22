package processlog

import (
	"bufio"
	"io"
	"time"
)

func NewFlushReadCloser(src io.Reader, flushSize int, flushInterval time.Duration) io.ReadCloser {
	reader := bufio.NewReader(src)
	flushTicker := time.NewTicker(flushInterval)
	flushChannel := make(chan bool)

	go func() {
		buffer := make([]byte, flushSize)
		written := 0
		for {
			select {
			case <-flushTicker.C:
				if written > 0 {
					flushChannel <- true
				}
			default:
				n, err := reader.Read(buffer[written:])
				if err != nil {
					if err == io.EOF && written > 0 {
						flushChannel <- true
					}
					flushTicker.Stop()
					close(flushChannel)
					return
				}
				written += n
				if written == flushSize {
					flushChannel <- true
					written = 0
				}
			}
		}
	}()

	return &flushReadCloser{reader, flushChannel}
}

type flushReadCloser struct {
	reader       *bufio.Reader
	flushChannel chan bool
}

func (f *flushReadCloser) Read(p []byte) (n int, err error) {
	n, err = f.reader.Read(p)
	if err == io.EOF {
		if len(p) > n || <-f.flushChannel {
			m, flushErr := f.reader.Read(p[n:])
			n += m
			if flushErr != nil {
				return n, flushErr
			}
		}
	}
	return
}

func (f *flushReadCloser) Close() error {
	f.reader.Reset(nil)

	return nil
}
