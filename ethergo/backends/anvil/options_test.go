package anvil_test

import (
	"github.com/ethereum/go-ethereum/accounts"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"reflect"
	"testing"
)

// TODO: make tests more comprehesnive here.
func TestOptionsBuilder(t *testing.T) {
	optionBuilder := anvil.NewAnvilOptionBuilder()
	_, err := optionBuilder.Build()
	NoError(t, err)
}

func TestGetFunctionName(t *testing.T) {
	funcName := anvil.GetFunctionName(TestGetFunctionName)
	Equal(t, funcName, "TestGetFunctionName")
}

func TestValueToString(t *testing.T) {
	testCases := []struct {
		value    interface{}
		expected string
		err      bool
	}{
		{true, "", false},
		{123, "123", false},
		{123.456, "123.456", false},
		{"hello", "\"hello\"", false},
		{anvil.Latest, anvil.Latest.String(), false},
		{anvil.Fees, anvil.Fees.String(), false},
		{accounts.DefaultRootDerivationPath, "\"m/44'/60'/0'\"", false},
		{struct{ a int }{a: 42}, "", true},
	}

	for _, tc := range testCases {
		v := reflect.ValueOf(tc.value)
		actual, err := anvil.ValueToString(v)
		if tc.err {
			NotNil(t, err)
			continue
		}
		if NoError(t, err) {
			Equal(t, tc.expected, actual)
		}
	}
}

// nolint: scopelint
func TestFieldIsEmpty(t *testing.T) {
	type testStruct struct {
		BoolVal    bool
		IntVal     int
		UintVal    uint
		FloatVal   float64
		StringVal  string
		SliceVal   []int
		MapVal     map[string]int
		StructVal  struct{ Field int }
		PointerVal *int
	}

	tests := []struct {
		name     string
		inputVal interface{}
		expected bool
	}{
		{"bool", false, true},
		{"int", 0, true},
		{"uint", uint(0), true},
		{"float", 0.0, true},
		{"string", "", true},
		{"slice", []int{}, true},
		{"map", map[string]int{}, true},
		{"struct", testStruct{}, true},
		{"pointer", (*int)(nil), true},
		{"non-empty bool", true, false},
		{"non-empty int", 10, false},
		{"non-empty uint", uint(10), false},
		{"non-empty float", 3.14, false},
		{"non-empty string", "hello", false},
		{"non-empty slice", []int{1, 2, 3}, false},
		{"non-empty map", map[string]int{"one": 1}, false},
		{"non-empty struct", testStruct{IntVal: 10}, false},
		{"non-nil pointer", new(int), false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := reflect.ValueOf(test.inputVal)
			if got := anvil.FieldIsEmpty(v); got != test.expected {
				t.Errorf("fieldIsEmpty(%v) = %v, expected %v", test.inputVal, got, test.expected)
			}
		})
	}
}
