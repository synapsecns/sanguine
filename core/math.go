package core

import (
	"math"
	"math/big"
)

// BigToDecimals converts a big to decimals
func BigToDecimals(bigInt *big.Int, decimals uint8) float64 {
	// Convert vpriceMetric to *big.Float
	bigVPrice := new(big.Float).SetInt(CopyBigInt(bigInt))

	// Calculate the divisor for decimals
	divisor := new(big.Float).SetFloat64(math.Pow10(int(decimals)))

	// Divide bigVPrice by the divisor to account for decimals
	realVPrice := new(big.Float).Quo(bigVPrice, divisor)

	// Convert the final value to float64
	floatVPrice, _ := realVPrice.Float64()
	return floatVPrice
}
