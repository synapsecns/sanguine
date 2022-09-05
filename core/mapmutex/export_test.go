package mapmutex

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

// TestMapMutex extends untypedMapMutex for testing.
type TestMapMutex interface {
	untypedMapMutex
	GetMa() map[interface{}]*mentry
}

// NewTestMapMutex wraps map mutex and casts it to a test map mutex for testing.
func NewTestMapMutex(tb testing.TB) TestMapMutex {
	tb.Helper()
	mapMux := newMapMutex()
	testMapMux, ok := mapMux.(TestMapMutex)
	True(tb, ok)
	return testMapMux
}

// GetMa exports ma from the map mutex for testing.
func (m *untypedMapMutexImpl) GetMa() map[interface{}]*mentry {
	return m.ma
}
