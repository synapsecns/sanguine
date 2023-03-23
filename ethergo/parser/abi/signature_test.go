package abi_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/synapsecns/sanguine/ethergo/parser/abi"
	"github.com/synapsecns/sanguine/ethergo/parser/abi/internal"
	"math/big"
)

func (a *AbiSuite) TestGetSelectorNoExist() {
	selector, err := abi.GetSelectorByName("test", internal.TestSignatureMetaData)
	a.Require().Error(err)
	a.Require().Empty(selector)
}

func (a *AbiSuite) TestGetSelectorSuccess() {
	expectedSelector, err := a.signature.TestSignature(&bind.CallOpts{Context: a.GetTestContext()})
	a.Require().NoError(err)

	realSelector, err := abi.GetSelectorByName("testSignature", internal.TestSignatureMetaData)
	a.Require().NotEmpty(realSelector)

	a.Require().NoError(err)
	a.Require().Equal(expectedSelector, realSelector)
}

func (a *AbiSuite) TestGetSelectorArgs() {
	expectedSelector, err := a.signature.TestSignatureArgs(&bind.CallOpts{Context: a.GetTestContext()}, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	a.Require().NoError(err)

	realSelector, err := abi.GetSelectorByName("testSignatureArgs", internal.TestSignatureMetaData)
	a.Require().NoError(err)
	a.Require().NotEmpty(realSelector)

	a.Require().Equal(expectedSelector, realSelector)
}

func (a *AbiSuite) TestSignatureOverload() {
	expectedSelector, err := a.signature.TestSignatureOverload0(&bind.CallOpts{Context: a.GetTestContext()}, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	a.Require().NoError(err)

	otherExpectedSelector, err := a.signature.TestSignatureOverload(&bind.CallOpts{Context: a.GetTestContext()}, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	a.Require().NoError(err)

	// sanity check
	a.Require().NotEqual(expectedSelector, otherExpectedSelector)

	realSelector, err := abi.GetSelectorByName("testSignatureOverload", internal.TestSignatureMetaData)
	a.Require().Error(err)
	a.Require().NotEqual(realSelector, expectedSelector)
	a.Require().Empty(realSelector)
}
