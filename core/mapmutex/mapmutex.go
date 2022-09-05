package mapmutex

import (
	"fmt"
	"github.com/LK4d4/trylock"
	"sync"
)

// untypedMapMutex wraps a map of mutexes.  Each key locks separately.
// this implementation is a modified version of https://stackoverflow.com/a/62562831/1011803
// this is unexported on purpose to prevent possible issues with pointers not being equivalent.
// TODO: different implementations need to be moved to ginny: https://flaviocopes.com/golang-generic-generate/
type untypedMapMutex interface {
	Lock(key interface{}) Unlocker
	TryLock(key interface{}) (Unlocker, bool)
}

type untypedMapMutexImpl struct {
	ml sync.Mutex              // lock for entry map
	ma map[interface{}]*mentry // entry map
}

type mentry struct {
	// m point back to untypedMapMutexImpl, so we can synchronize removing this mentry when cnt==0
	m *untypedMapMutexImpl
	// el is an entry-specific lock
	el trylock.Mutex
	// cnt is the reference count
	cnt int
	// key is the key of the memory entry
	key interface{}
}

// Unlocker provides an Unlock method to release the lock.
type Unlocker interface {
	Unlock()
}

// newMapMutex returns an initialized untypedMapMutexImpl.
func newMapMutex() untypedMapMutex {
	return &untypedMapMutexImpl{ma: make(map[interface{}]*mentry)}
}

// Lock acquires a lock corresponding to this key.
// This method will never return nil and Unlock() must be called
// to release the lock when done.
func (m *untypedMapMutexImpl) Lock(key interface{}) Unlocker {
	// read or create entry for this key atomically
	m.ml.Lock()
	e, ok := m.ma[key]
	if !ok {
		e = &mentry{m: m, key: key}
		m.ma[key] = e
	}
	e.cnt++ // ref count
	m.ml.Unlock()

	// acquire lock, will block here until e.cnt==1
	e.el.Lock()

	return e
}

// TryLock tries to acquire the lock, if this can't be done instantly false is returned. Otherwise
// true and the unlocker are returned.
func (m *untypedMapMutexImpl) TryLock(key interface{}) (Unlocker, bool) {
	// read or create entry for this key atomically
	m.ml.Lock()
	defer m.ml.Unlock()

	e, ok := m.ma[key]
	if !ok {
		e = &mentry{m: m, key: key}
		m.ma[key] = e
	}

	if e.el.TryLock() {
		e.cnt++
		return e, true
	}

	return nil, false
}

// Unlock releases the lock for this entry.
func (me *mentry) Unlock() {
	m := me.m

	// decrement and if needed remove entry atomically
	m.ml.Lock()
	e, ok := m.ma[me.key]
	if !ok { // entry must exist
		m.ml.Unlock()
		panic(fmt.Errorf("unlock requested for key=%v but no entry found", me.key))
	}
	e.cnt--        // ref count
	if e.cnt < 1 { // if it hits zero then we own it and remove from map
		delete(m.ma, me.key)
	}
	m.ml.Unlock()

	// now that map stuff is handled, we unlock and let
	// anything else waiting on this key through
	e.el.Unlock()
}
