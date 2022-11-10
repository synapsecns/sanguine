package parser

import (
	"context"
	"database/sql"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
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
