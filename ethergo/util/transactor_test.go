package util_test

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/go-cmp/cmp"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/util"
	"go.opentelemetry.io/otel/attribute"
	"gotest.tools/assert"
	"math/big"
	"testing"
)

func TestCopyTransactOpts(t *testing.T) {
	// Test case 1: All fields populated
	opts1 := &bind.TransactOpts{
		From:      common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Nonce:     big.NewInt(1),
		Signer:    nil,
		Value:     big.NewInt(100),
		GasPrice:  big.NewInt(200),
		GasFeeCap: big.NewInt(300),
		GasTipCap: big.NewInt(400),
		GasLimit:  500,
		Context:   context.Background(),
		NoSend:    true,
	}

	copyOpts1 := util.CopyTransactOpts(opts1)
	assertTransactOptsEquality(t, opts1, copyOpts1)

	// Test case 2: Some fields populated, others nil
	opts2 := &bind.TransactOpts{
		Nonce:     nil,
		Signer:    nil,
		Value:     nil,
		GasPrice:  nil,
		GasFeeCap: big.NewInt(300),
		GasTipCap: big.NewInt(400),
		GasLimit:  500,
		Context:   context.Background(),
		NoSend:    true,
	}

	copyOpts2 := util.CopyTransactOpts(opts2)
	assertTransactOptsEquality(t, opts2, copyOpts2)

	// Test case 3: All fields nil
	opts3 := &bind.TransactOpts{
		Nonce:     nil,
		Signer:    nil,
		Value:     nil,
		GasPrice:  nil,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  0,
		Context:   nil,
		NoSend:    true,
	}

	copyOpts3 := util.CopyTransactOpts(opts3)
	assertTransactOptsEquality(t, opts3, copyOpts3)
}

func assertTransactOptsEquality(tb testing.TB, toA, toB *bind.TransactOpts) {
	tb.Helper()
	// Check that the pointer values of the big integer fields are different
	assertBigIntsCopiedEqual(tb, toA.Nonce, toB.Nonce, "Nonce")
	assertBigIntsCopiedEqual(tb, toA.Value, toB.Value, "Value")
	assertBigIntsCopiedEqual(tb, toA.GasPrice, toB.GasPrice, "GasPrice")
	assertBigIntsCopiedEqual(tb, toA.GasFeeCap, toB.GasFeeCap, "GasFeeCap")
	assertBigIntsCopiedEqual(tb, toA.GasTipCap, toB.GasTipCap, "GasTipCap")
	assert.DeepEqual(tb, toA, toB, testsuite.BigIntComparer(), cmp.AllowUnexported(context.Background()))
}

// assertBigIntsCopiedEqual checks that the given big.Ints are equal and that
// they have different pointers.
func assertBigIntsCopiedEqual(tb testing.TB, original *big.Int, newVal *big.Int, fieldName string) {
	tb.Helper()
	if original == nil && newVal == nil {
		return
	}

	if core.ArePointersEqual(original, newVal) {
		tb.Errorf("%s has same pointer as original", fieldName)
	}

	if original.Cmp(newVal) != 0 {
		tb.Errorf("%s is not equal", fieldName)
	}
}

func TestAddressPtrToString(t *testing.T) {
	// Test case 1: Address is nil
	var address *common.Address
	assert.Equal(t, util.AddressPtrToString(address), util.NullFieldAttribute)

	// Test case 2: Address is not nil
	address = core.PtrTo[common.Address](common.HexToAddress("0x1234567890123456789012345678901234567890"))
	assert.Equal(t, util.AddressPtrToString(address), "0x1234567890123456789012345678901234567890")
}

func TestBigPtrToString(t *testing.T) {
	// Test case: num is nil
	var num *big.Int
	expected := util.NullFieldAttribute
	result := util.BigPtrToString(num)
	if result != expected {
		t.Errorf("BigPtrToString(nil) = %q; want %q", result, expected)
	}

	// Test case: num is an integer
	num = big.NewInt(123)
	expected = "123"
	result = util.BigPtrToString(num)
	if result != expected {
		t.Errorf("BigPtrToString(123) = %q; want %q", result, expected)
	}
}

func (s *UtilSuite) TestTxToAttributesNullFields() {
	s.checkEmptyTx(types.NewTx(&types.DynamicFeeTx{}))
	s.checkEmptyTx(types.NewTx(&types.LegacyTx{}))
}

func (s *UtilSuite) checkEmptyTx(rawTx *types.Transaction) {
	tx := makeAttrMap(rawTx)

	s.Require().Equal(tx[util.HashAttr].AsString(), rawTx.Hash().Hex())
	s.Require().Equal(tx[util.NonceAttr].AsInt64(), int64(0))
	s.Require().Equal(tx[util.GasLimitAttr].AsInt64(), int64(0))
	s.Require().Equal(tx[util.ToAttr].AsString(), util.NullFieldAttribute)
	s.Require().Equal(tx[util.ValueAttr].AsString(), "0")
	s.Require().Equal(tx[util.DataAttr].AsString(), "")

	if rawTx.Type() == types.DynamicFeeTxType {
		s.Require().Equal(tx[util.GasTipCapAttr].AsString(), "0")
		s.Require().Equal(tx[util.GasFeeCapAttr].AsString(), "0")
	}
	if rawTx.Type() == types.LegacyTxType {
		s.Require().Equal(tx[util.GasPriceAttr].AsString(), "0")
	}
}

func (s *UtilSuite) TestTxToAttributesLegacyTX() {
	mockTX := mocks.GetMockTxes(s.GetTestContext(), s.T(), 1, types.LegacyTxType)[0]
	mapAttr := makeAttrMap(mockTX)

	s.Require().Equal(mapAttr[util.HashAttr].AsString(), mockTX.Hash().String())
	s.Require().Equal(mapAttr[util.NonceAttr].AsInt64(), int64(mockTX.Nonce()))
	s.Require().Equal(mapAttr[util.GasLimitAttr].AsInt64(), int64(mockTX.Gas()))
	s.Require().Equal(mapAttr[util.ToAttr].AsString(), mockTX.To().String())
	s.Require().Equal(mapAttr[util.ValueAttr].AsString(), mockTX.Value().String())
	s.Require().Equal(mapAttr[util.DataAttr].AsString(), "")

	s.Require().Equal(mapAttr[util.GasPriceAttr].AsString(), mockTX.GasPrice().String())
	_, hasFeeCap := mapAttr[util.GasFeeCapAttr]
	_, hasTipCap := mapAttr[util.GasTipCapAttr]
	s.Require().False(hasFeeCap)
	s.Require().False(hasTipCap)
	s.Require().NotNil(mapAttr[util.FromAttr])
}

func (s *UtilSuite) TestTxToAttributesDynamicTX() {
	mockTX := mocks.GetMockTxes(s.GetTestContext(), s.T(), 1, types.DynamicFeeTxType)[0]
	mapAttr := makeAttrMap(mockTX)

	s.Require().Equal(mapAttr[util.HashAttr].AsString(), mockTX.Hash().String())
	s.Require().Equal(mapAttr[util.NonceAttr].AsInt64(), int64(mockTX.Nonce()))
	s.Require().Equal(mapAttr[util.GasLimitAttr].AsInt64(), int64(mockTX.Gas()))
	s.Require().Equal(mapAttr[util.ToAttr].AsString(), mockTX.To().String())
	s.Require().Equal(mapAttr[util.ValueAttr].AsString(), mockTX.Value().String())
	s.Require().Equal(mapAttr[util.DataAttr].AsString(), "")

	s.Require().Equal(mapAttr[util.GasFeeCapAttr].AsString(), mockTX.GasFeeCap().String())
	s.Require().Equal(mapAttr[util.GasTipCapAttr].AsString(), mockTX.GasTipCap().String())
	_, hasGasPrice := mapAttr[util.GasPriceAttr]
	s.Require().False(hasGasPrice)
	s.Require().NotNil(mapAttr[util.FromAttr])
}

func makeAttrMap(tx *types.Transaction) map[string]attribute.Value {
	mapAttr := make(map[string]attribute.Value)
	attr := util.TxToAttributes(tx)
	for _, a := range attr {
		mapAttr[string(a.Key)] = a.Value
	}
	return mapAttr
}
