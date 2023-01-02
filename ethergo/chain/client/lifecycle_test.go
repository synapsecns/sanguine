package client_test

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	clientMocks "github.com/synapsecns/sanguine/ethergo/chain/client/mocks"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/viant/toolbox"
)

// TestLifecycleClient makes sure all methods properly call acquire and release permit. Because this method uses reflect everywhere it requires some clever
// use of args to rpevent duplicate code.
func (c ClientSuite) TestLifecycleClient() {
	mockUnderlyingClient := clientMocks.EVMClient{}
	mockPermitter := clientMocks.Permitter{}
	// no errors
	permitAcquires := mockPermitter.On("AcquirePermit", mock.Anything).Return(nil)
	permitReleases := mockPermitter.On("ReleasePermit")

	lifecycleClient := client.NewLifecycleClient(&mockUnderlyingClient, big.NewInt(0), &mockPermitter, time.Minute)
	evmClienType := reflect.TypeOf((*client.EVMClient)(nil)).Elem()

	// use a mock backend + some clever reflection to mock return values for all these funds
	simulatedBackend := simulated.NewSimulatedBackend(c.GetTestContext(), c.T())
	for i := 0; i < evmClienType.NumMethod(); i++ {
		// reflect the type of the method
		method := evmClienType.Method(i)
		// get that method on the simulated client so we can returns *something*, we need to do this to avoid panicking when release permit is called
		reflectedMethod, ok := reflect.TypeOf(simulatedBackend.Client()).MethodByName(method.Name)
		True(c.T(), ok)
		// create the args array
		args := createArgsArr(reflectedMethod.Type)
		returnVals := createReturnArr(simulatedBackend, reflectedMethod.Func)

		// special cases
		if method.Name == "SendTransaction" {
			mockUnderlyingClient.On(method.Name, args...).
				Return(nil)
			continue
		}
		// save the call
		mockUnderlyingClient.
			// mock using vals reflected vals
			On(method.Name, args...).
			// we don't care what this returns so use the simulated backend value
			Return(returnVals...)
	}

	// test the methods.
	for i := 0; i < evmClienType.NumMethod(); i++ {
		// get the method matching this name on the lifecycle client
		methodName := evmClienType.Method(i).Name

		if shouldSkip(methodName) {
			continue
		}

		lifecycleClientMethod := reflect.ValueOf(lifecycleClient).MethodByName(methodName)
		methodSig, ok := reflect.TypeOf(lifecycleClient).MethodByName(methodName)
		True(c.T(), ok)

		lifecycleClientMethod.Call(c.createMockInputArgs(methodSig))
	}
	// make sure underlying was called
	// defer mockUnderlyingClient.AssertNumberOfCalls(c.T(), methodName, 1)
	// 4 non-locking methods + skipped
	mockPermitter.AssertNumberOfCalls(c.T(), permitReleases.Method, evmClienType.NumMethod()-5)
	mockPermitter.AssertNumberOfCalls(c.T(), permitAcquires.Method, evmClienType.NumMethod()-5)
}

// shouldSkip indicates an untestable method.
func shouldSkip(name string) bool {
	skipMethods := []string{"BatchCallContext", "CallContext", "BatchContext", "SyncProgress", "FeeHistory"}
	return toolbox.HasSliceAnyElements(skipMethods, name)
}

// createMockInputArgs creates mock input arguments for a given func
// for an explanation of why we can't use reflect.new or reflect.zero,  see: https://stackoverflow.com/a/26321245/1011803
//
//nolint:cyclop
func (c ClientSuite) createMockInputArgs(p reflect.Method) (res []reflect.Value) {
	for i := 0; i < p.Type.NumIn(); i++ {
		// ignore the lifecycle client
		if i == 0 {
			continue
		}
		inputType := p.Type.In(i)

		var val interface{}
		switch inputType.String() {
		case "context.Context":
			val = c.GetTestContext()
		case reflect.TypeOf(common.Address{}).String():
			val = mocks.MockAddress()
		case reflect.TypeOf(&big.Int{}).String():
			val = big.NewInt(0)
		case reflect.TypeOf(common.Hash{}).String():
			val = common.BigToHash(big.NewInt(0))
		case reflect.TypeOf(ethereum.CallMsg{}).String():
			val = ethereum.CallMsg{}
		case reflect.TypeOf(ethereum.FilterQuery{}).String():
			val = ethereum.FilterQuery{}
		case reflect.TypeOf(&types.Transaction{}).String():
			val = mocks.GetMockTxes(c.GetTestContext(), c.T(), 1, types.LegacyTxType)[0]
		case reflect.TypeOf(make(chan<- types.Log)).String():
			val = make(chan<- types.Log)
		case reflect.TypeOf(make(chan<- *types.Header)).String():
			val = make(chan<- *types.Header)
		case reflect.TypeOf([]rpc.BatchElem{}).String():
			val = []rpc.BatchElem{}
		case reflect.TypeOf(uint(0)).String():
			val = uint(0)
		case reflect.TypeOf(uint64(0)).String():
			val = uint64(0)
		case reflect.TypeOf([]float64{}).String():
			val = []float64{1, 2}
		default:
			panic(fmt.Errorf("type %s not handled", inputType.String()))
		}
		res = append(res, reflect.ValueOf(val))
	}
	return res
}

func createArgsArr(p reflect.Type) (res []interface{}) {
	for i := 0; i < p.NumIn(); i++ {
		res = append(res, mock.Anything)
	}
	return res
}

// createReturnArr creates a return array with the correct values. Return will call the method n times to get all args
// because Return takes all args but can only return one at a time, this requires some reflection magic.
func createReturnArr(backend *simulated.Backend, p reflect.Value) (res []interface{}) {
	// get an array of all in args
	var inArgs []reflect.Type
	for i := 0; i < p.Type().NumIn(); i++ {
		// the first function is the object itself when reflected.
		if i == 0 {
			continue
		}
		inArgs = append(inArgs, p.Type().In(i))
	}

	for i := 0; i < p.Type().NumOut(); i++ {
		// copy the arg to the local scope, anything above this point will be redefined when func is called
		argIndex := i
		// create a new type with same in args, but one out arg at current index
		newType := reflect.FuncOf(inArgs, []reflect.Type{p.Type().Out(argIndex)}, false)
		// create a new function that only returns the function at index and therefore conforms to the new type defined
		// above
		newFunc := reflect.MakeFunc(newType, func(args []reflect.Value) (results []reflect.Value) {
			// re-add the object
			res := p.Call(append([]reflect.Value{reflect.ValueOf(backend.Chain)}, args...))
			return []reflect.Value{res[argIndex]}
		})
		test, ok := newFunc.Interface().(func(context.Context, common.Address, *big.Int) *big.Int)
		_ = ok
		_ = test
		// append the new function to the return array
		res = append(res, newFunc.Interface())
	}
	return res
}
