package core

import (
	"math/big"
	"reflect"
)

// PtrTo returns a pointer to the given value.
func PtrTo[T any](v T) *T {
	return &v
}

// ArePointersEqual returns true if the given pointers are equal.
// Will return false if either of the given values are not pointers.
// nolint: cyclop, forcetypeassert
func ArePointersEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	if aValue.Kind() != reflect.Ptr || bValue.Kind() != reflect.Ptr {
		return false
	}

	// zero-value big ints are represented by the same pointer
	if aValue.Type() == reflect.TypeOf((*big.Int)(nil)) && bValue.Type() == reflect.TypeOf((*big.Int)(nil)) {
		aBigInt := aValue.Interface().(*big.Int)
		bBigInt := bValue.Interface().(*big.Int)

		if aBigInt == nil && bBigInt == nil {
			return false
		}
		if aBigInt.IsUint64() && aBigInt.Uint64() == 0 && bBigInt.IsUint64() && bBigInt.Uint64() == 0 {
			return false
		}
	}

	return aValue.Pointer() == bValue.Pointer()
}
