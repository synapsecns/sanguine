package core_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core"
	"reflect"
	"testing"
)

func ExamplePtrTo() {
	res := core.PtrTo(common.Hash{})
	fmt.Println(reflect.TypeOf(res))
	// Output: *common.Hash
}

func ExamplePtrSlice() {
	res := core.PtrSlice([]common.Hash{})
	fmt.Println(reflect.TypeOf(res))
	// Output: []*common.Hash
}

func TestArePointersEqual(t *testing.T) {
	type TestData struct {
		a    interface{}
		b    interface{}
		want bool
	}

	intVal1 := 1
	intVal2 := 2

	testData := []TestData{
		{&intVal1, &intVal1, true},
		{&intVal1, &intVal2, false},
		{nil, nil, true},
		{&intVal1, nil, false},
		{nil, &intVal1, false},
		{intVal1, intVal1, false},  // non-pointers
		{intVal1, &intVal1, false}, // non-pointer and pointer
		{&intVal1, intVal1, false}, // pointer and non-pointer
	}

	for _, data := range testData {
		got := core.ArePointersEqual(data.a, data.b)
		if got != data.want {
			t.Errorf("ArePointersEqual(%v, %v) = %v; want %v", data.a, data.b, got, data.want)
		}
	}
}

func TestCopyPointer(t *testing.T) {
	original := int64(42)
	originalPtr := &original

	newPtr := core.CopyPointer(originalPtr)

	if newPtr == nil {
		t.Errorf("CopyPointer returned nil")
	}

	if newPtr == originalPtr {
		t.Errorf("CopyPointer returned the same pointer")
	}

	if *newPtr != *originalPtr {
		t.Errorf("CopyPointer did not copy the value correctly, got %d, want %d", *newPtr, *originalPtr)
	}
}

func TestCopyPointerNil(t *testing.T) {
	var originalPtr *int64 = nil

	newPtr := core.CopyPointer(originalPtr)

	if newPtr != nil {
		t.Errorf("CopyPointer did not return nil for nil input")
	}
}
