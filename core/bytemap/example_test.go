package bytemap_test

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/bytemap"
)

func ExampleByteSliceMap_Put_getString() {
	m := &bytemap.ByteSliceMap[string]{}

	// Put value using a string key
	m.PutString("hello", "world")

	// Get value using a string key
	value, ok := m.GetString("hello")

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("value not found")
	}

	// Output: world
}

func ExampleByteSliceMap_Get_bytes() {
	m := &bytemap.ByteSliceMap[string]{}

	// Put value using a byte slice key
	m.Put([]byte("golang"), "awesome")

	// Get value using a byte slice key
	value, ok := m.Get([]byte("golang"))

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("value not found")
	}

	// Output: awesome
}
