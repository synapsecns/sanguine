package core_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core"
	"reflect"
)

func ExamplePtrTo() {
	res := core.PtrTo(common.Hash{})
	fmt.Println(reflect.TypeOf(res))
	// Output: *common.Hash
}
