package types_test

import (
	"math/big"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/spatialcurrent/go-math/pkg/math"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/types"
)

func TestTotalTips(t *testing.T) {
	for i := 0; i < 100; i++ {
		// use uint32 to make sure we don't overflow
		vals := []int{int(gofakeit.Uint32()), int(gofakeit.Uint32()), int(gofakeit.Uint32()), int(gofakeit.Uint32())}

		// create a new tips from values above
		tips := types.NewTips(big.NewInt(int64(vals[0])), big.NewInt(int64(vals[1])), big.NewInt(int64(vals[2])), big.NewInt(int64(vals[3])))

		// should be convertible to int64
		realSum, err := math.Sum(vals)
		Nil(t, err)

		//nolint:forcetypeassert
		Equal(t, types.TotalTips(tips).Uint64(), uint64(realSum.(int)))
	}
}
