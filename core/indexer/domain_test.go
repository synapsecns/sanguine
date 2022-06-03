package indexer_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/indexer"
	"math"
	"testing"
)

func TestUint32Max(t *testing.T) {
	// fuzz
	for i := 0; i < 50; i++ {
		small := gofakeit.Uint32()
		// we can't assert greater then max
		if small == math.MaxUint32 {
			continue
		}

		larger := small + 1

		Equal(t, indexer.MaxUint32(small, larger), larger)
	}

	// edge case
	Equal(t, indexer.MaxUint32(math.MaxUint32, 4), uint32(math.MaxUint32))
}
