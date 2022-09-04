package mapmutex

import "fmt"

// StringerMapMutex is an implementation of mapMutex for the fmt.Stringer conforming types.
type StringerMapMutex interface {
	Lock(key fmt.Stringer) Unlocker
	TryLock(key fmt.Stringer) (Unlocker, bool)
}

// stringerLockerImpl is the implementation of StringerMapMutex.
type stringerLockerImpl struct {
	mapMux untypedMapMutex
}

func (s stringerLockerImpl) TryLock(key fmt.Stringer) (Unlocker, bool) {
	return s.mapMux.TryLock(key.String())
}

// Lock locks on the string.
func (s stringerLockerImpl) Lock(key fmt.Stringer) Unlocker {
	return s.mapMux.Lock(key.String())
}

// NewStringerMapMutex creates an initialized locker that locks on fmt.String.
func NewStringerMapMutex() StringerMapMutex {
	return &stringerLockerImpl{
		mapMux: newMapMutex(),
	}
}

// StringMapMutex is an implementation of map mutex for string typed values.
type StringMapMutex interface {
	Lock(key string) Unlocker
	TryLock(key string) (Unlocker, bool)
}

// stringMutexImpl locks on a string type.
type stringMutexImpl struct {
	mapMux untypedMapMutex
}

// NewStringMapMutex creates a map mutex for the string type.
func NewStringMapMutex() StringMapMutex {
	return &stringMutexImpl{
		mapMux: newMapMutex(),
	}
}

// Lock locks ona  string value.
func (s stringMutexImpl) Lock(key string) Unlocker {
	return s.mapMux.Lock(key)
}

// TryLock attempts to lock on a string value.
func (s stringMutexImpl) TryLock(key string) (Unlocker, bool) {
	return s.mapMux.TryLock(key)
}

// IntMapMutex is a map mutex that allows locking on an int.
type IntMapMutex interface {
	Lock(key int) Unlocker
	TryLock(key int) (Unlocker, bool)
}

// intMapMux locks on an int.
type intMapMux struct {
	mapMux untypedMapMutex
}

func (i intMapMux) TryLock(key int) (Unlocker, bool) {
	return i.mapMux.TryLock(key)
}

// Lock locks an int map mux.
func (i intMapMux) Lock(key int) Unlocker {
	return i.mapMux.Lock(key)
}

// NewIntMapMutex creates a map mutex for locking on an integer.
func NewIntMapMutex() IntMapMutex {
	return &intMapMux{
		mapMux: newMapMutex(),
	}
}
