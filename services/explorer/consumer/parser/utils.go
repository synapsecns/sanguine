package parser

import (
	"context"
	"database/sql"
	"fmt"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/yaml.v2"
	"math/big"
	"strconv"
)

// ErrUnknownTopic is returned when the topic is unknown.
const ErrUnknownTopic = "unknown topic"

// Parser parses events and stores them.
type Parser interface {
	// Parse parses the logs and returns the parsed data.
	Parse(ctx context.Context, log ethTypes.Log, chainID uint32) (interface{}, error)
	// ParserType returns the type of the parser.
	ParserType() string
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

// ToNullInt64 is a helper function to convert values to null string.
func ToNullInt64[T int64 | uint64](val *T) sql.NullInt64 {
	var newNullInt sql.NullInt64

	if val != nil {
		newNullInt.Valid = true
		newNullInt.Int64 = int64(*val)
	} else {
		newNullInt.Valid = false
	}

	return newNullInt
}

// ToNullInt32 is a helper function to convert values to null string.
func ToNullInt32[T int32 | uint32](val *T) sql.NullInt32 {
	var newNullInt sql.NullInt32

	if val != nil {
		newNullInt.Valid = true
		newNullInt.Int32 = int32(*val)
	} else {
		newNullInt.Valid = false
	}

	return newNullInt
}

// ParseYaml opens yaml file with coin gecko ID mapping and returns it.
func ParseYaml(yamlFile []byte) (res map[string]string, err error) {
	if err != nil {
		return nil, fmt.Errorf("error opening yaml file %w", err)
	}

	err = yaml.Unmarshal(yamlFile, &res)

	if err != nil {
		return nil, fmt.Errorf("error unmarshalling yaml file %w", err)
	}

	return res, nil
}

// GetAmountUSD computes the USD value of a token amount.
func GetAmountUSD(amount *big.Int, decimals uint8, price *float64) *float64 {
	decimalMultiplier := new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	adjustedAmount := new(big.Float).Quo(new(big.Float).SetInt(amount), decimalMultiplier)
	trueAmount := big.NewFloat(0).Mul(adjustedAmount, big.NewFloat(*price))
	trueAmountStr := trueAmount.SetMode(big.AwayFromZero).Text('f', 5)
	priceFloat, err := strconv.ParseFloat(trueAmountStr, 64)
	if err != nil {
		return nil
	}
	return &priceFloat
}

// GetAdjustedAmount computes the adjusted token amount.
func GetAdjustedAmount(amount *big.Int, decimals uint8) *float64 {
	decimalMultiplier := new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	adjustedAmount := new(big.Float).Quo(new(big.Float).SetInt(amount), decimalMultiplier)
	trueAmountStr := adjustedAmount.SetMode(big.AwayFromZero).Text('f', 5)
	fullAmount, err := strconv.ParseFloat(trueAmountStr, 64)
	if err != nil {
		return nil
	}
	return &fullAmount
}
