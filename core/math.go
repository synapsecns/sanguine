package core

import (
	"math"
	"math/big"
)

// BigToDecimals converts a big to decimals
// TODO: unit test.
func BigToDecimals(bigInt *big.Int, decimals uint8) float64 {
	// Convert vpriceMetric to *big.Float
	bigVal := new(big.Float).SetInt(CopyBigInt(bigInt))

	// Calculate the divisor for decimals
	divisor := new(big.Float).SetFloat64(math.Pow10(int(decimals)))

	// Divide bigVPrice by the divisor to account for decimals
	realVal := new(big.Float).Quo(bigVal, divisor)

	// Convert the final value to float64
	floatVal, _ := realVal.Float64()
	return floatVal
}
