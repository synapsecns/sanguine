package threaditer_test

import (
	"github.com/Soft/iter"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/threaditer"
	"testing"
)

func TestSlice(t *testing.T) {
	// TODO: test thread safety
	it := threaditer.ThreadSafe(iter.Slice([]int{1, 2, 3}))

	Equal(t, it.Next().Unwrap(), 1)
	Equal(t, it.Next().Unwrap(), 2)
	Equal(t, it.Next().Unwrap(), 3)
}
