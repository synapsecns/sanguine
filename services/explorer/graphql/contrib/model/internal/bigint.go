package internal

import (
	"math"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var BigInt = graphql.NewScalar(graphql.ScalarConfig{
	Name: "BigInt",
	Description: "The `BigInt` scalar type represents non-fractional signed whole numeric. " +
		"values. BigInt can represent values between -(2^53) + 1 and 2^53 - 1. ",
	Serialize:  coerceBigInt,
	ParseValue: coerceBigInt,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.IntValue:
			if intValue, err := strconv.Atoi(valueAST.Value); err == nil {
				return intValue
			}
		}
		return nil
	},
})

func coerceBigInt(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return 1
		}
		return 0
	case *bool:
		if value == nil {
			return nil
		}
		return coerceBigInt(*value)
	case int:
		if value < int(math.MinInt64) || value > int(math.MaxInt64) {
			return nil
		}
		return value
	case *int:
		if value == nil {
			return nil
		}
		return coerceBigInt(*value)
	case int8:
		return int64(value)
	case *int8:
		if value == nil {
			return nil
		}
		return int64(*value)
	case int16:
		return int64(value)
	case *int16:
		if value == nil {
			return nil
		}
		return int64(*value)
	case int32:
		return int64(value)
	case *int32:
		if value == nil {
			return nil
		}
		return int64(*value)
	case int64:
		if value < int64(math.MinInt64) || value > int64(math.MaxInt64) {
			return nil
		}
		return int64(value)
	case *int64:
		if value == nil {
			return nil
		}
		return coerceBigInt(*value)
	case uint:
		if value > math.MaxInt64 {
			return nil
		}
		return int(value)
	case *uint:
		if value == nil {
			return nil
		}
		return coerceBigInt(*value)
	case uint8:
		return int64(value)
	case *uint8:
		if value == nil {
			return nil
		}
		return int(*value)
	case uint16:
		return int(value)
	case *uint16:
		if value == nil {
			return nil
		}
		return int64(*value)
	case uint32:
		if uint64(value) > uint64(math.MaxInt64) {
			return nil
		}
		return int64(value)
	case *uint32:
		if value == nil {
			return nil
		}
		return coerceBigInt(*value)
	case uint64:
		if value > uint64(math.MaxInt64) {
			return nil
		}
		return int64(value)
	case *uint64:
		if value == nil {
			return nil
		}
		return coerceBigInt(*value)
	case float32:
		if value < float32(math.MinInt64) || value > float32(math.MaxInt64) {
			return nil
		}
		return math.RoundToEven(float64(value))
	case *float32:
		if value == nil {
			return nil
		}
		return coerceBigInt(*value)
	case float64:
		if value < float64(math.MinInt64) || value > float64(math.MaxInt64) {
			return nil
		}
		return math.RoundToEven(value)
	case *float64:
		if value == nil {
			return nil
		}
		return coerceBigInt(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return coerceBigInt(val)
	case *string:
		if value == nil {
			return nil
		}
		return coerceBigInt(*value)
	}

	return nil
}
