package parser

import (
	"context"
	"database/sql"
	"fmt"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/yaml.v2"
	"math"
	"math/big"
	"os"
)

// Parser parses events and stores them.
type Parser interface {
	// ParseAndStore parses the logs and stores them in the database.
	ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error
}

// BoolToUint8 is a helper function to handle bool to uint8 conversion for clickhouse.
func BoolToUint8(input *bool) *uint8 {
	if input == nil {
		return nil
	}
	if *input {
		one := uint8(1)

		return &one
	}
	zero := uint8(0)

	return &zero
}

// ToNullString is a helper function to convert values to null string.
func ToNullString(str *string) sql.NullString {
	var newNullStr sql.NullString

	if str != nil {
		newNullStr.Valid = true
		newNullStr.String = *str
	} else {
		newNullStr.Valid = false
	}

	return newNullStr
}

// OpenYaml opens yaml file with coin gecko ID mapping and returns it.
func OpenYaml(path string) (map[string]string, error) {
	// nolint:gosec
	input, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("error opening yaml file %w", err)
	}

	var res map[string]string
	err = yaml.Unmarshal(input, &res)

	if err != nil {
		return nil, fmt.Errorf("error unmarshalling yaml file %w", err)
	}

	return res, nil
}

// GetAmountUSD computes the USD value of a token amount.
func GetAmountUSD(amount *big.Int, decimals uint8, price *float64) *float64 {
	trueAmount := (float64(amount.Uint64()) / math.Pow(10.0, float64(decimals))) * *price

	return &trueAmount
}
