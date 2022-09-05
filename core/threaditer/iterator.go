package threaditer

import (
	"github.com/Soft/iter"
	"sync"
)

// ThreadSafe is a threadsafe iterator adapator used for iterators
// accessed in parallell.
func ThreadSafe[T any](underlyingIter iter.Iterator[T]) iter.Iterator[T] {
	return &threadSafeIter[T]{
		parentIter: underlyingIter,
		mux:        sync.Mutex{},
	}
}

type threadSafeIter[T any] struct {
	parentIter iter.Iterator[T]
	mux        sync.Mutex
}

// Next gets the next item in a thread safe manner.
func (t *threadSafeIter[T]) Next() iter.Option[T] {
	t.mux.Lock()
	defer t.mux.Unlock()

	return t.parentIter.Next()
}
