package gas

import "math/big"

// Min exports the min method for testing.
func Min(a, b *big.Int) *big.Int {
	return min(a, b)
}
