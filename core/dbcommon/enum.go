package dbcommon

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/rung/go-safecast"
	"reflect"
)

// EnumInter ensures an enum has a type to access an integer.
type EnumInter interface {
	Int() uint8
}

// Enum is a type used to define interfaces db enums must conform to.
// it is kept here so it can be used throughout the app.
type Enum interface {
	EnumInter
	fmt.Stringer
	GormDataType() string
	sql.Scanner
	driver.Valuer
}

// EnumDataType is exported here to be passed as GormDataType.
// TODO: support string types
const EnumDataType = "integer"

// EnumValue converts the enum to a value.
func EnumValue(enum EnumInter) (driver.Value, error) {
	converter := driver.NotNull{Converter: driver.DefaultParameterConverter}
	rawValue := enum.Int()
	val, err := converter.ConvertValue(rawValue)
	if err != nil {
		return nil, fmt.Errorf("could not convert rawValue %d of type %T: %w", rawValue, rawValue, err)
	}
	return val, nil
}

// EnumScan converts the enum to a value.
// the returned uint 8 should be used. Note this should only be used for a non-nullable enum.
// TODO: consider using generics to make this more type safe.
func EnumScan(src interface{}) (uint8, error) {
	res := sql.NullInt32{}
	err := res.Scan(src)
	if err != nil {
		return 0, fmt.Errorf("could not scan enum: %w", err)
	}
	if res.Int32 < 0 {
		return 0, fmt.Errorf("enum value cannot be a negative number, got: %w", err)
	}
	cast, err := safecast.Int8(int(res.Int32))
	if err != nil {
		return 0, fmt.Errorf("could not cast %d (type %s) to %s: %w", res.Int32, reflect.TypeOf(res.Int32), reflect.TypeOf(cast), err)
	}
	return uint8(cast), nil
}
