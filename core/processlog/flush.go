package processlog

import (
	"bytes"
	"io"
	"sync"
	"time"
)

func NewFlushReadCloser(src io.Reader, flushInterval time.Duration, bufferSize int) io.ReadCloser {
	ring := newRingBuffer(bufferSize)
	outputBuffer := bytes.NewBuffer(nil)
	flushTicker := time.NewTicker(flushInterval)
	flushChannel := make(chan struct{})

	var eof bool

	go func() {
		buffer := make([]byte, bufferSize)
		for {
			n, err := src.Read(buffer)
			if n > 0 {
				ring.Write(buffer[:n])
			}
			if err != nil {
				if err != io.EOF {
					return
				}
				eof = true
				break
			}
		}
	}()

	go func() {
		defer close(flushChannel)

		for {
			select {
			case <-flushTicker.C:
				data := ring.Read()
				if len(data) > 0 {
					if _, err := outputBuffer.Write(data); err != nil {
						return
					}
					flushChannel <- struct{}{}
				}
			}
			if eof && ring.Len() == 0 {
				break
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

type ringBuffer struct {
	buffer []byte
	size   int
	head   int
	tail   int
	count  int
	mutex  sync.Mutex
}

func newRingBuffer(size int) *ringBuffer {
	return &ringBuffer{
		buffer: make([]byte, size),
		size:   size,
		head:   0,
		tail:   0,
		count:  0,
	}
}

func (r *ringBuffer) Write(data []byte) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for len(data) > 0 {
		if r.count == r.size {
			r.tail = (r.tail + 1) % r.size
			r.count--
		}
		pos := (r.head + r.count) % r.size
		n := copy(r.buffer[pos:], data)
		data = data[n:]
		r.count += n
	}
}

func (r *ringBuffer) Read() []byte {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	data := make([]byte, r.count)
	if r.count > 0 {
		if r.head+r.count <= r.size {
			copy(data, r.buffer[r.head:r.head+r.count])
		} else {
			n := copy(data, r.buffer[r.head:r.size])
			copy(data[n:], r.buffer[0:r.head+r.count-r.size])
		}
		r.head = (r.head + r.count) % r.size
		r.count = 0
	}
	return data
}

func (r *ringBuffer) Len() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.count
}
