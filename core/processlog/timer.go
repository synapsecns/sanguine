package processlog

import (
	"io"
	"sync"
	"time"
)

const (
	flushInterval = time.Millisecond * 10
)

type intervalFlushReader struct {
	src     io.Reader
	dst     []byte
	buffer  []byte
	readCh  chan int
	flushCh chan struct{}
	errCh   chan error
	doneCh  chan struct{}
	mu      sync.Mutex
}

func newIntervalFlushReader(r io.Reader) *intervalFlushReader {
	ifr := &intervalFlushReader{
		src:     r,
		dst:     make([]byte, 0, pipeBufferSize),
		buffer:  make([]byte, pipeBufferSize),
		readCh:  make(chan int, 1),
		flushCh: make(chan struct{}, 1),
		errCh:   make(chan error, 1),
		doneCh:  make(chan struct{}),
	}

	go ifr.readLoop()

	return ifr
}

func (ifr *intervalFlushReader) Read(p []byte) (int, error) {
	ifr.mu.Lock()
	defer ifr.mu.Unlock()

	n := copy(p, ifr.dst)
	ifr.dst = ifr.dst[n:]
	if n == 0 && len(ifr.dst) == 0 {
		select {
		case err := <-ifr.errCh:
			return 0, err
		default:
		}
	}
	return n, nil
}

func (ifr *intervalFlushReader) Close() error {
	close(ifr.doneCh)
	return nil
}

func (ifr *intervalFlushReader) readLoop() {
	flushTicker := time.NewTicker(flushInterval)
	defer flushTicker.Stop()

	for {
		select {
		case <-ifr.doneCh:
			return
		case <-flushTicker.C:
			select {
			case ifr.flushCh <- struct{}{}:
			default:
			}
		default:
			n, err := ifr.src.Read(ifr.buffer)
			if err != nil {
				select {
				case ifr.errCh <- err:
				default:
				}
				return
			}

			select {
			case ifr.readCh <- n:
			default:
			}
		}

		select {
		case <-ifr.doneCh:
			return
		case n := <-ifr.readCh:
			ifr.mu.Lock()
			ifr.dst = append(ifr.dst, ifr.buffer[:n]...)
			ifr.mu.Unlock()
		case <-ifr.flushCh:
			ifr.mu.Lock()
			ifr.dst = append(ifr.dst, ifr.buffer[:cap(ifr.buffer)-len(ifr.buffer)]...)
			ifr.mu.Unlock()
			ifr.buffer = ifr.buffer[:cap(ifr.buffer)]

			select {
			case err := <-ifr.errCh:
				if err == io.EOF {
					return
				}
			default:
			}
		}
	}
}
