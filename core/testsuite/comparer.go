package testsuite

import (
	"github.com/google/go-cmp/cmp"
	"math/big"
)

// BigIntComparer gets the big int comparer for testing.
func BigIntComparer() cmp.Option {
	return cmp.Comparer(func(x *big.Int, y *big.Int) bool {
		if x == nil && y == nil {
			return true
		}
		return x.Cmp(y) == 0
	})
}
