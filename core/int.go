package core

import "math/big"

// CopyBigInt creates a copy of a big int without mutating the original.
func CopyBigInt(val *big.Int) *big.Int {
	if val == nil {
		return nil
	}
	return new(big.Int).SetBytes(val.Bytes())
}
