package processlog

import (
	"bytes"
	"io"
	"time"
)

func NewFlushReadCloser(src io.Reader, flushInterval time.Duration, bufferSize int) io.ReadCloser {
	inputChannel := make(chan []byte)
	outputBuffer := bytes.NewBuffer(make([]byte, 0, bufferSize))
	flushTicker := time.NewTicker(flushInterval)
	flushChannel := make(chan struct{}, 1)

	go func() {
		defer close(inputChannel)
		buffer := make([]byte, bufferSize)
		for {
			n, err := src.Read(buffer)
			if n > 0 {
				inputChannel <- buffer[:n]
			}
			if err != nil {
				if err != io.EOF {
					return
				}
				break
			}
		}
	}()

	go func() {
		defer close(flushChannel)

		buffer := make([]byte, bufferSize)
		written := 0

		for {
			select {
			case <-flushTicker.C:
				if written > 0 {
					if _, err := outputBuffer.Write(buffer[:written]); err != nil {
						return
					}
					written = 0
					flushChannel <- struct{}{}
				}
			case data, ok := <-inputChannel:
				if !ok {
					if written > 0 {
						if _, err := outputBuffer.Write(buffer[:written]); err != nil {
							return
						}
						written = 0
						flushChannel <- struct{}{}
					}
					return
				}
				for len(data) > 0 {
					remaining := bufferSize - written
					if len(data) < remaining {
						copy(buffer[written:], data)
						written += len(data)
						data = nil
					} else {
						copy(buffer[written:], data[:remaining])
						written += remaining
						data = data[remaining:]
						if written == bufferSize {
							if _, err := outputBuffer.Write(buffer); err != nil {
								return
							}
							written = 0
							flushChannel <- struct{}{}
						}
					}
				}
			}
		}
	}()

	return &flushReadCloser{
		reader:       outputBuffer,
		flushChannel: flushChannel,
	}
}

type flushReadCloser struct {
	reader       io.Reader
	flushChannel chan struct{}
}

func (f *flushReadCloser) Read(p []byte) (n int, err error) {
	n, err = f.reader.Read(p)
	if err == io.EOF {
		<-f.flushChannel
		if n == 0 {
			return 0, io.EOF
		}
		err = nil
	}
	return
}

func (f *flushReadCloser) Close() error {
	if closer, ok := f.reader.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
