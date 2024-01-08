package core_test

import (
	"github.com/synapsecns/sanguine/core"
	"math/big"
	"testing"
)

func TestBigToDecimals(t *testing.T) {
	tests := []struct {
		name     string
		bigInt   *big.Int
		decimals uint8
		want     float64
	}{
		{
			name:     "Basic Conversion",
			bigInt:   big.NewInt(1000000),
			decimals: 2,
			want:     10000.00,
		},
		{
			name:     "Zero Value",
			bigInt:   big.NewInt(0),
			decimals: 5,
			want:     0.0,
		},
		{
			name:     "Large Number",
			bigInt:   big.NewInt(1234567890123456789),
			decimals: 9,
			want:     1234567890.123456789,
		},
		// Add more test cases as needed
	}

	for i := range tests {
		tt := tests[i] // capture func literal
		t.Run(tt.name, func(t *testing.T) {
			if got := core.BigToDecimals(tt.bigInt, tt.decimals); got != tt.want {
				t.Errorf("BigToDecimals() = %v, want %v", got, tt.want)
			}
		})
	}
}
