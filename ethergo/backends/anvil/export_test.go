package anvil

import "reflect"

// GetFunctionName exports getFunctionName for testing.
func GetFunctionName(temp interface{}) string {
	return getFunctionName(temp)
}

// ValueToString exports valueToString for testing.
func ValueToString(v reflect.Value) (string, error) {
	return valueToString(v)
}

// FieldIsEmpty exports fieldIsEmpty for testing.
func FieldIsEmpty(v reflect.Value) bool {
	return fieldIsEmpty(v)
}

// AllHardforks exports allHardforks for testing.
func AllHardforks() []Hardfork {
	return allHardforks
}
