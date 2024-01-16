package core

import (
	"math/big"
	"reflect"
)

// PtrTo returns a pointer to the given value.
func PtrTo[T any](v T) *T {
	return &v
}

// PtrSlice converts every item in the slice to a pointer.
func PtrSlice[T any](slice []T) []*T {
	ptrSlice := make([]*T, len(slice))
	for i, item := range slice {
		ptrSlice[i] = PtrTo(item)
	}
	return ptrSlice
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

// CopyPointer is a generic function that takes a pointer of any type and returns a new pointer to a new value of the same type.
func CopyPointer[T any](originalPtr *T) *T {
	if originalPtr == nil {
		return nil
	}
	newValue := *originalPtr
	return &newValue
}
