// nolint: forcetypeassert, dupl
package abiutil_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/parser/abiutil"
	"github.com/synapsecns/sanguine/ethergo/parser/abiutil/internal"
	"gotest.tools/assert"
	"math/big"
)

func (a *AbiSuite) TestUnpackDataToInterface() {
	auth := a.backend.GetTxContext(a.GetTestContext(), nil)

	paramA := new(big.Int).SetUint64(gofakeit.Uint64())
	paramB := common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))
	paramC := new(big.Int).SetUint64(gofakeit.Uint64())
	paramD := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	paramE := []byte(gofakeit.Paragraph(1, 10, 1, " "))
	paramF := gofakeit.Bool()

	tx, err := a.testContract.DoSomethingManyTypes(auth.TransactOpts, paramA, paramB, paramC, paramD, paramE, paramF)

	a.Require().NoError(err)
	a.backend.WaitForConfirmation(a.GetTestContext(), tx)

	receipt, err := a.backend.TransactionReceipt(a.GetTestContext(), tx.Hash())
	a.Require().NoError(err)

	filterNum := receipt.BlockNumber.Uint64()

	filteredEvent, err := a.testContract.FilterSomethingHappenedManyTypes(&bind.FilterOpts{
		Context: a.GetTestContext(),
		Start:   filterNum,
		End:     &filterNum,
	})
	a.Require().NoError(err)

	// unpack into interface
	var originalEvent internal.TestSignatureSomethingHappenedManyTypes
	err = abiutil.UnpackInputDataToInterface(&originalEvent, tx.Data(), a.metadata)
	a.Require().NoError(err)

	// get the first event
	filteredEvent.Next()

	assert.DeepEqual(a.T(), originalEvent.A, filteredEvent.Event.A, testsuite.BigIntComparer())
	assert.DeepEqual(a.T(), originalEvent.B, filteredEvent.Event.B)
	assert.DeepEqual(a.T(), originalEvent.C, filteredEvent.Event.C, testsuite.BigIntComparer())
	assert.DeepEqual(a.T(), common.Bytes2Hex(originalEvent.D[:]), common.Bytes2Hex(filteredEvent.Event.D[:]))
	assert.DeepEqual(a.T(), originalEvent.E, filteredEvent.Event.E)
	assert.DeepEqual(a.T(), originalEvent.F, filteredEvent.Event.F)
}

func (a *AbiSuite) TestDoSomethingNoParamsUnpack() {
	auth := a.backend.GetTxContext(a.GetTestContext(), nil)

	tx, err := a.testContract.DoSomethingWithoutParams(auth.TransactOpts)
	a.Require().NoError(err)
	a.backend.WaitForConfirmation(a.GetTestContext(), tx)

	out, err := abiutil.UnpackInputData(tx.Data(), a.metadata)
	a.Require().NoError(err)

	a.Require().Len(out, 0)
}

func (a *AbiSuite) TestDoSomethingUnpack() {
	auth := a.backend.GetTxContext(a.GetTestContext(), nil)

	paramA := new(big.Int).SetUint64(gofakeit.Uint64())
	paramB := new(big.Int).SetUint64(gofakeit.Uint64())

	tx, err := a.testContract.DoSomething(auth.TransactOpts, paramA, paramB)
	a.Require().NoError(err)
	a.backend.WaitForConfirmation(a.GetTestContext(), tx)

	out, err := abiutil.UnpackInputData(tx.Data(), a.metadata)
	a.Require().NoError(err)
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	assert.DeepEqual(a.T(), out0, paramA, testsuite.BigIntComparer())
	assert.DeepEqual(a.T(), out1, paramB, testsuite.BigIntComparer())
}

func (a *AbiSuite) TestDoSomethingOverloadUnpack() {
	auth := a.backend.GetTxContext(a.GetTestContext(), nil)

	paramA := new(big.Int).SetUint64(gofakeit.Uint64())
	paramB := new(big.Int).SetUint64(gofakeit.Uint64())

	tx, err := a.testContract.DoSomethingOverload(auth.TransactOpts, paramA, paramB)
	a.Require().NoError(err)
	a.backend.WaitForConfirmation(a.GetTestContext(), tx)

	out, err := abiutil.UnpackInputData(tx.Data(), a.metadata)
	a.Require().NoError(err)
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	assert.DeepEqual(a.T(), out0, paramA, testsuite.BigIntComparer())
	assert.DeepEqual(a.T(), out1, paramB, testsuite.BigIntComparer())
}

func (a *AbiSuite) TestDoSomethingOverload0Unpack() {
	auth := a.backend.GetTxContext(a.GetTestContext(), nil)

	paramA := new(big.Int).SetUint64(gofakeit.Uint64())
	paramB := common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))

	tx, err := a.testContract.DoSomethingOverload0(auth.TransactOpts, paramA, paramB)
	a.Require().NoError(err)
	a.backend.WaitForConfirmation(a.GetTestContext(), tx)

	out, err := abiutil.UnpackInputData(tx.Data(), a.metadata)
	a.Require().NoError(err)
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	assert.DeepEqual(a.T(), out0, paramA, testsuite.BigIntComparer())
	assert.DeepEqual(a.T(), out1, paramB)
}

func (a *AbiSuite) TestDoSomethingManyTypeUnpack() {
	auth := a.backend.GetTxContext(a.GetTestContext(), nil)

	paramA := new(big.Int).SetUint64(gofakeit.Uint64())
	paramB := common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))
	paramC := new(big.Int).SetUint64(gofakeit.Uint64())
	paramD := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	paramE := []byte(gofakeit.Paragraph(1, 10, 1, " "))
	paramF := gofakeit.Bool()

	tx, err := a.testContract.DoSomethingManyTypes(auth.TransactOpts, paramA, paramB, paramC, paramD, paramE, paramF)
	a.Require().NoError(err)
	a.backend.WaitForConfirmation(a.GetTestContext(), tx)

	out, err := abiutil.UnpackInputData(tx.Data(), a.metadata)
	a.Require().NoError(err)

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	out4 := *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	assert.DeepEqual(a.T(), out0, paramA, testsuite.BigIntComparer())
	assert.DeepEqual(a.T(), out1, paramB)
	assert.DeepEqual(a.T(), out2, paramC, testsuite.BigIntComparer())
	assert.Equal(a.T(), common.Bytes2Hex(out3[:]), common.Bytes2Hex(paramD.Bytes()))
	assert.DeepEqual(a.T(), out4, paramE)
	assert.DeepEqual(a.T(), out5, paramF)
}
