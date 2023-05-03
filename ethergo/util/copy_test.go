package util_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/util"
	"math/big"
	"testing"
)

// nolint: cyclop, gocognit
func (u *UtilSuite) TestCopyTX() {
	mockDynamicTXS := mocks.GetMockTxes(u.GetTestContext(), u.T(), 100, types.DynamicFeeTxType)
	mockLegacyTXes := mocks.GetMockTxes(u.GetTestContext(), u.T(), 100, types.LegacyTxType)

	// delete me
	mockTXes := append(mockDynamicTXS, mockLegacyTXes...)

	for _, tx := range mockTXes {
		// fuzz some data
		txFuzzData := makeFuzzData(u.T(), tx)
		txFuzzData.txType = types.DynamicFeeTxType // delete me

		// build options
		var options []util.CopyOption
		if txFuzzData.nonce != tx.Nonce() {
			options = append(options, util.WithNonce(txFuzzData.nonce))
		}

		if txFuzzData.gasPrice != nil && txFuzzData.gasPrice.Cmp(tx.GasPrice()) != 0 && tx.Type() == types.LegacyTxType {
			options = append(options, util.WithGasPrice(txFuzzData.gasPrice))
		}

		if txFuzzData.feeCap != nil && txFuzzData.feeCap.Cmp(tx.GasFeeCap()) != 0 && tx.Type() == types.DynamicFeeTxType {
			options = append(options, util.WithGasFeeCap(txFuzzData.feeCap))
		}

		if txFuzzData.tipCap != nil && txFuzzData.tipCap.Cmp(tx.GasTipCap()) != 0 && tx.Type() == types.DynamicFeeTxType {
			options = append(options, util.WithGasTipCap(txFuzzData.tipCap))
		}

		if txFuzzData.txType != tx.Type() {
			options = append(options, util.WithTxType(txFuzzData.txType))
		}

		// copy the tx with the options
		newTx, err := util.CopyTX(tx, options...)
		u.Require().NoError(err)

		ogV, ogR, ogS := tx.RawSignatureValues()
		newV, newR, newS := newTx.RawSignatureValues()

		// make sure pointers changed
		assert.False(u.T(), core.ArePointersEqual(tx.ChainId(), newTx.ChainId()))
		assert.False(u.T(), core.ArePointersEqual(tx.Nonce(), newTx.Nonce()))
		assert.False(u.T(), core.ArePointersEqual(tx.GasPrice(), newTx.GasPrice()))
		assert.False(u.T(), core.ArePointersEqual(tx.GasTipCap(), newTx.GasTipCap()))
		assert.False(u.T(), core.ArePointersEqual(tx.GasFeeCap(), newTx.GasFeeCap()))
		assert.False(u.T(), core.ArePointersEqual(ogV, newV))
		assert.False(u.T(), core.ArePointersEqual(ogR, newR))
		assert.False(u.T(), core.ArePointersEqual(ogS, newS))

		// make sure fields without options were not changed
		assert.Equal(u.T(), tx.ChainId(), newTx.ChainId(), testsuite.BigIntComparer())
		assert.Equal(u.T(), tx.To(), newTx.To())
		assert.Equal(u.T(), tx.Value(), newTx.Value(), testsuite.BigIntComparer())
		assert.Equal(u.T(), tx.Data(), newTx.Data(), testsuite.BigIntComparer())

		// they won't be equal in this case
		if tx.Type() == txFuzzData.txType {
			assert.Equal(u.T(), tx.Type(), newTx.Type())
		}
		// make sure new tx matches target
		assert.Equal(u.T(), newTx.Type(), txFuzzData.txType)
		assert.Equal(u.T(), tx.Gas(), newTx.Gas())

		// don't check the v param if we've changed it to return the chainid correctly
		if tx.Type() == txFuzzData.txType {
			assert.Equal(u.T(), ogV, newV, testsuite.BigIntComparer())
			// this wil change slightly, it's not a problem
			assert.Equal(u.T(), tx.AccessList(), newTx.AccessList())

			// make sure the newTX matches the fuzz data
			if !txFuzzData.matchesTX(newTx) {
				u.T().Errorf("expected tx to match fuzz data, but it did not. tx: %v, fuzzData: %v", newTx, txFuzzData)
			}
		}
		assert.Equal(u.T(), ogR, newR, testsuite.BigIntComparer())
		assert.Equal(u.T(), ogS, newS, testsuite.BigIntComparer())
	}
}

// nolint: gocognit, cyclop
func makeFuzzData(tb testing.TB, tx *types.Transaction) fuzzData {
	tb.Helper()

	txType := tx.Type()
	nonce := tx.Nonce()
	if gofakeit.Bool() {
		nonce = gofakeit.Uint64()
	}
	var gasPrice, feeCap, tipCap *big.Int
	if tx.Type() == types.DynamicFeeTxType {
		feeCap = tx.GasFeeCap()
		if gofakeit.Bool() {
			feeCap = new(big.Int).SetUint64(gofakeit.Uint64())
		}

		tipCap = tx.GasTipCap()
		if gofakeit.Bool() {
			tipCap = new(big.Int).SetUint64(gofakeit.Uint64())
		}
	}
	if tx.Type() == types.LegacyTxType {
		gasPrice = tx.GasPrice()

		if gofakeit.Bool() {
			gasPrice = new(big.Int).SetUint64(gofakeit.Uint64())
		}
	}

	if gofakeit.Bool() {
		var err error
		txTypes := []uint8{types.DynamicFeeTxType, types.LegacyTxType}
		txType, err = core.RandomItem[uint8](txTypes)
		assert.Nilf(tb, err, "failed to get random tx type: %v", err)
	}

	return fuzzData{
		nonce:    nonce,
		gasPrice: gasPrice,
		feeCap:   feeCap,
		tipCap:   tipCap,
		txType:   txType,
	}
}

type fuzzData struct {
	nonce    uint64
	gasPrice *big.Int
	feeCap   *big.Int
	tipCap   *big.Int
	txType   uint8
}

func (f *fuzzData) matchesTX(tx *types.Transaction) bool {
	if f.nonce != tx.Nonce() {
		return false
	}
	if f.gasPrice != nil && f.gasPrice.Cmp(tx.GasPrice()) != 0 {
		return false
	}
	if f.txType == tx.Type() {
		if f.feeCap != nil && f.feeCap.Cmp(tx.GasFeeCap()) != 0 {
			return false
		}

		if f.tipCap != nil && f.tipCap.Cmp(tx.GasTipCap()) != 0 {
			return false
		}
	}
	return true
}

func TestMakeOptionsNil(t *testing.T) {
	opts := util.MakeOptions()
	if opts.Nonce() != nil {
		t.Errorf("expected nonce to be nil, got %v", opts.Nonce())
	}
	if opts.GasPrice() != nil {
		t.Errorf("expected gas price to be nil, got %v", opts.GasPrice())
	}
	if opts.GasFeeCap() != nil {
		t.Errorf("expected gas fee cap to be nil, got %v", opts.GasFeeCap())
	}
	if opts.GasTipCap() != nil {
		t.Errorf("expected gas tip cap to be nil, got %v", opts.GasTipCap())
	}
}

func TestMakeOptionsWithNonce(t *testing.T) {
	opts := util.MakeOptions(util.WithNonce(1234))

	if opts.Nonce() == nil {
		t.Errorf("expected nonce to be set, got nil")
	}
	if *opts.Nonce() != 1234 {
		t.Errorf("expected nonce to be 1234, got %v", *opts.Nonce())
	}
	if opts.GasPrice() != nil {
		t.Errorf("expected gas price to be nil, got %v", opts.GasPrice())
	}
	if opts.GasFeeCap() != nil {
		t.Errorf("expected gas fee cap to be nil, got %v", opts.GasFeeCap())
	}
	if opts.GasTipCap() != nil {
		t.Errorf("expected gas tip cap to be nil, got %v", opts.GasTipCap())
	}
}

func TestMakeOptionsWithTxType(t *testing.T) {
	txTypes := []uint8{types.LegacyTxType, types.DynamicFeeTxType, types.AccessListTxType}
	for _, txType := range txTypes {
		opts := util.MakeOptions(util.WithTxType(txType))

		if opts.Nonce() != nil {
			t.Errorf("expected nonce to not be set, got %v", opts.Nonce())
		}

		if opts.TxType() == nil {
			t.Errorf("expected tx type to be set, got nil")
		}

		if *opts.TxType() != txType {
			t.Errorf("expected tx type to be %v, got %v", txType, *opts.TxType())
		}

		if opts.GasPrice() != nil {
			t.Errorf("expected gas price to be nil, got %v", opts.GasPrice())
		}
		if opts.GasFeeCap() != nil {
			t.Errorf("expected gas fee cap to be nil, got %v", opts.GasFeeCap())
		}
		if opts.GasTipCap() != nil {
			t.Errorf("expected gas tip cap to be nil, got %v", opts.GasTipCap())
		}
	}
}

func TestMakeOptionsWithGasPrice(t *testing.T) {
	gasPrice := big.NewInt(100)
	opts := util.MakeOptions(util.WithGasPrice(gasPrice))

	if opts.Nonce() != nil {
		t.Errorf("expected nonce to be nil, got %v", opts.Nonce())
	}
	if opts.GasPrice() == nil {
		t.Errorf("expected gas price to be set, got nil")
	}
	if opts.GasPrice().Cmp(gasPrice) != 0 {
		t.Errorf("expected gas price to be %v, got %v", gasPrice, opts.GasPrice())
	}
	if opts.GasFeeCap() != nil {
		t.Errorf("expected gas fee cap to be nil, got %v", opts.GasFeeCap())
	}
	if opts.GasTipCap() != nil {
		t.Errorf("expected gas tip cap to be nil, got %v", opts.GasTipCap())
	}
}

func TestMakeOptionsWithGasFeeCap(t *testing.T) {
	gasFeeCap := big.NewInt(200)
	gasTipCap := big.NewInt(300)
	opts := util.MakeOptions(util.WithGasFeeCap(gasFeeCap), util.WithGasTipCap(gasTipCap))

	if opts.Nonce() != nil {
		t.Errorf("expected nonce to be nil, got %v", opts.Nonce())
	}
	if opts.GasPrice() != nil {
		t.Errorf("expected gas price to be nil, got %v", opts.GasPrice())
	}
	if opts.GasFeeCap() == nil {
		t.Errorf("expected gas fee cap to be set, got nil")
	}
	if opts.GasFeeCap().Cmp(gasFeeCap) != 0 {
		t.Errorf("expected gas fee cap to be %v, got %v", gasFeeCap, opts.GasFeeCap())
	}
	if opts.GasTipCap() != gasTipCap {
		t.Errorf("expected gas tip cap to be nil, got %v", opts.GasTipCap())
	}
}
