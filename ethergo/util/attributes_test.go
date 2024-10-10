package util_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/util"
	"go.opentelemetry.io/otel/attribute"
	"math/big"
	"testing"
)

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

func makeAttrMap(tx *types.Transaction) map[string]attribute.Value {
	mapAttr := make(map[string]attribute.Value)
	attr := util.TxToAttributes(tx)
	for _, a := range attr {
		mapAttr[string(a.Key)] = a.Value
	}
	return mapAttr
}

func (u *UtilSuite) checkEmptyTx(rawTx *types.Transaction) {
	tx := makeAttrMap(rawTx)

	u.Require().Equal(tx[util.HashAttr].AsString(), rawTx.Hash().Hex())
	u.Require().Equal(tx[util.NonceAttr].AsInt64(), int64(0))
	u.Require().Equal(tx[util.GasLimitAttr].AsInt64(), int64(0))
	u.Require().Equal(tx[util.ToAttr].AsString(), util.NullFieldAttribute)
	u.Require().Equal(tx[util.ValueAttr].AsString(), "0")
	u.Require().Equal(tx[util.DataAttr].AsString(), "")

	if rawTx.Type() == types.DynamicFeeTxType {
		u.Require().Equal(tx[util.GasTipCapAttr].AsString(), "0")
		u.Require().Equal(tx[util.GasFeeCapAttr].AsString(), "0")
	}
	if rawTx.Type() == types.LegacyTxType {
		u.Require().Equal(tx[util.GasPriceAttr].AsString(), "0")
	}
}

func (u *UtilSuite) TestTxToAttributesNullFields() {
	u.checkEmptyTx(types.NewTx(&types.DynamicFeeTx{}))
	u.checkEmptyTx(types.NewTx(&types.LegacyTx{}))
}

func (u *UtilSuite) TestTxToAttributesLegacyTX() {
	mockTX := mocks.GetMockTxes(u.GetTestContext(), u.T(), 1, types.LegacyTxType)[0]
	mapAttr := makeAttrMap(mockTX)

	u.Require().Equal(mapAttr[util.HashAttr].AsString(), mockTX.Hash().String())
	u.Require().Equal(mapAttr[util.NonceAttr].AsInt64(), int64(mockTX.Nonce()))
	u.Require().Equal(mapAttr[util.GasLimitAttr].AsInt64(), int64(mockTX.Gas()))
	u.Require().Equal(mapAttr[util.ToAttr].AsString(), mockTX.To().String())
	u.Require().Equal(mapAttr[util.ValueAttr].AsString(), mockTX.Value().String())
	u.Require().Equal(mapAttr[util.DataAttr].AsString(), "")

	u.Require().Equal(mapAttr[util.GasPriceAttr].AsString(), mockTX.GasPrice().String())
	_, hasFeeCap := mapAttr[util.GasFeeCapAttr]
	_, hasTipCap := mapAttr[util.GasTipCapAttr]
	u.Require().False(hasFeeCap)
	u.Require().False(hasTipCap)
	u.Require().NotNil(mapAttr[util.FromAttr])
}

func (u *UtilSuite) TestTxToAttributesDynamicTX() {
	mockTX := mocks.GetMockTxes(u.GetTestContext(), u.T(), 1, types.DynamicFeeTxType)[0]
	mapAttr := makeAttrMap(mockTX)

	u.Require().Equal(mapAttr[util.HashAttr].AsString(), mockTX.Hash().String())
	u.Require().Equal(mapAttr[util.NonceAttr].AsInt64(), int64(mockTX.Nonce()))
	u.Require().Equal(mapAttr[util.GasLimitAttr].AsInt64(), int64(mockTX.Gas()))
	u.Require().Equal(mapAttr[util.ToAttr].AsString(), mockTX.To().String())
	u.Require().Equal(mapAttr[util.ValueAttr].AsString(), mockTX.Value().String())
	u.Require().Equal(mapAttr[util.DataAttr].AsString(), "")

	u.Require().Equal(mapAttr[util.GasFeeCapAttr].AsString(), mockTX.GasFeeCap().String())
	u.Require().Equal(mapAttr[util.GasTipCapAttr].AsString(), mockTX.GasTipCap().String())
	_, hasGasPrice := mapAttr[util.GasPriceAttr]
	u.Require().False(hasGasPrice)
	u.Require().NotNil(mapAttr[util.FromAttr])
}
