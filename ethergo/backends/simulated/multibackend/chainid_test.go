package multibackend_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	util "github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated/multibackend"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"gotest.tools/assert"
	"math/big"
	"reflect"
	"testing"
)

// TestNewConfigWithChainID tests that the newly generated chainconfig is fully equivalent to the old one
// these are really easy to mess up if you're not careful, especially in a major upgrade and incredibly hard to identify
// so we run a test here
// TODO: also double check for chain mutations, maybe even in a global way?
func (s MultiBackendSuite) TestNewConfigWithChainID() {
	testChainID := big.NewInt(int64(gofakeit.Number(1, 9999)))
	newConfig := multibackend.NewConfigWithChainID(testChainID)

	Zerof(s.T(), newConfig.ChainID.Cmp(testChainID), "newConfig should use new chain id")

	// now use the right chain id and make sure everything is the same
	newConfig.ChainID = util.CopyBigInt(params.AllEthashProtocolChanges.ChainID)
	assert.DeepEqual(s.T(), newConfig, params.AllEthashProtocolChanges, testsuite.BigIntComparer())
	CheckBigIntsAreCopied(s.T(), *newConfig, *params.AllEthashProtocolChanges)

	// make sure we didn't mutate the chain id, which is always 1337
	Equal(s.T(), params.AllEthashProtocolChanges.ChainID.Int64(), int64(1337))
}

// CheckBigIntsAreCopied checks if two structs are equal by comparing all of their fields
// including big.Int values using reflect.DeepEqual, except for those that are *big.Int.
func CheckBigIntsAreCopied(tb testing.TB, a, b interface{}) {
	tb.Helper()

	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)

	// If the types of the two values are not the same, they are not equal
	if aValue.Type() != bValue.Type() {
		tb.Errorf("types are not equal: %v != %v", aValue.Type(), bValue.Type())
	}

	// Iterate over all fields in the struct
	for i := 0; i < aValue.NumField(); i++ {
		field := aValue.Type().Field(i)
		fieldName := field.Name

		fieldValueA := aValue.Field(i)
		fieldValueB := bValue.Field(i)

		// If the field is a *big.Int, use PointersAreEqual to compare it
		if fieldValueA.Type() == reflect.TypeOf((*big.Int)(nil)) {
			//nolint: forcetypeassert
			if util.ArePointersEqual(fieldValueA.Interface().(*big.Int), fieldValueB.Interface().(*big.Int)) {
				tb.Errorf("pointers are  equal for %s: %v != %v", fieldName, fieldValueA.Pointer(), fieldValueB.Pointer())
			}
		}
	}
}

// run a sanity check on our new backend by getting the chainid.
func (s MultiBackendSuite) TestNewSimulatedBackendWithChainID() {
	testChainID := big.NewInt(int64(gofakeit.Number(1, 9999)))

	// 100 million ether
	balance := big.NewInt(0).Mul(big.NewInt(params.Ether), big.NewInt(100000000))
	key := mocks.MockAccount(s.T())

	genesisAlloc := map[common.Address]core.GenesisAccount{
		key.Address: {
			Balance: balance,
		},
	}

	NotPanics(s.T(), func() {
		multibackend.NewSimulatedBackendWithConfig(genesisAlloc, uint64(91712388), multibackend.NewConfigWithChainID(testChainID))
	})
}
