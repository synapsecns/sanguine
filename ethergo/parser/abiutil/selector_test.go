package abiutil_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/synapsecns/sanguine/ethergo/parser/abiutil"
	"github.com/synapsecns/sanguine/ethergo/parser/abiutil/internal"
	"math/big"
)

func (a *AbiSuite) TestGetSelectorNoExist() {
	selector, err := abiutil.GetSelectorByName("test", internal.TestSignatureMetaData)
	a.Require().Error(err)
	a.Require().Empty(selector)

	a.Panics(func() {
		_, err = abiutil.GetSelectorByName("test", nil)
	})
}

// getSigsFromParity is a helper function to get the signatures
// it makes sure MustGetSelectorByName matches get selector by name.
func (a *AbiSuite) getSelectorSuccesful(name string, metadata *bind.MetaData) [4]byte {
	selector, err := abiutil.GetSelectorByName(name, metadata)
	a.Require().NoError(err)

	selector2 := abiutil.MustGetSelectorByName(name, metadata)
	a.Require().Equal(selector, selector2)

	return selector
}

func (a *AbiSuite) TestGetSelectorSuccess() {
	expectedSelector, err := a.testContract.TestSignature(&bind.CallOpts{Context: a.GetTestContext()})
	a.Require().NoError(err)

	realSelector := a.getSelectorSuccesful("testSignature", internal.TestSignatureMetaData)
	a.Require().Equal(expectedSelector, realSelector)
}

func (a *AbiSuite) TestGetSelectorArgs() {
	expectedSelector, err := a.testContract.TestSignatureArgs(&bind.CallOpts{Context: a.GetTestContext()}, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	a.Require().NoError(err)

	realSelector := a.getSelectorSuccesful("testSignatureArgs", internal.TestSignatureMetaData)
	a.Require().NotEmpty(realSelector)

	a.Require().Equal(expectedSelector, realSelector)
}

func (a *AbiSuite) TestSignatureOverloadSelector() {
	expectedSelector, err := a.testContract.TestSignatureOverload0(&bind.CallOpts{Context: a.GetTestContext()}, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	a.Require().NoError(err)

	otherExpectedSelector, err := a.testContract.TestSignatureOverload(&bind.CallOpts{Context: a.GetTestContext()}, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	a.Require().NoError(err)

	// sanity check
	a.Require().NotEqual(expectedSelector, otherExpectedSelector)

	realSelector, err := abiutil.GetSelectorByName("testSignatureOverload", internal.TestSignatureMetaData)
	a.Require().Error(err)
	a.Require().NotEqual(realSelector, expectedSelector)
	a.Require().Empty(realSelector)

	a.Panics(func() {
		_ = abiutil.MustGetSelectorByName("testSignatureOverload", internal.TestSignatureMetaData)
	})
}
