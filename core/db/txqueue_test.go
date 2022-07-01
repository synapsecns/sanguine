package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/db"
	"math/big"
)

// addressPtr converts an address to a pointer
func addressPtr(address common.Address) *common.Address {
	return &address
}

// testTxes is a list of txes to test inserting different tx types
// because we are not testing a raw tx
var testTxes = []*types.Transaction{
	types.NewTx(&types.LegacyTx{
		Nonce:    gofakeit.Uint64(),
		GasPrice: new(big.Int).SetUint64(gofakeit.Uint64()),
		Gas:      gofakeit.Uint64(),
		To:       addressPtr(common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))),
		Value:    new(big.Int).SetUint64(gofakeit.Uint64()),
		Data:     []byte(gofakeit.Paragraph(1, 2, 3, " ")),
	}),
	types.NewTx(&types.AccessListTx{
		ChainID:  big.NewInt(int64(gofakeit.Uint32())),
		Nonce:    gofakeit.Uint64(),
		GasPrice: new(big.Int).SetUint64(gofakeit.Uint64()),
		Gas:      gofakeit.Uint64(),
		To:       addressPtr(common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))),
		Value:    new(big.Int).SetUint64(gofakeit.Uint64()),
		Data:     []byte(gofakeit.Paragraph(1, 2, 3, " ")),
		AccessList: types.AccessList{
			types.AccessTuple{
				Address:     common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())),
				StorageKeys: []common.Hash{common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))},
			},
		},
	}),
	types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(gofakeit.Uint32())),
		Nonce:     gofakeit.Uint64(),
		GasTipCap: new(big.Int).Mul(new(big.Int).SetInt64(int64(gofakeit.Float32Range(1, 10))), big.NewInt(params.GWei)),
		GasFeeCap: new(big.Int).Mul(new(big.Int).SetInt64(int64(gofakeit.Float32Range(10, 100))), big.NewInt(params.GWei)),
		Gas:       gofakeit.Uint64(),
		To:        addressPtr(common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))),
		Value:     new(big.Int).SetUint64(gofakeit.Uint64()),
		Data:      []byte(gofakeit.Paragraph(1, 2, 3, " ")),
		AccessList: types.AccessList{
			types.AccessTuple{
				Address:     common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())),
				StorageKeys: []common.Hash{common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))},
			},
		},
	}),
}

func (t *TxQueueSuite) TestTxInsertion() {
	t.RunOnAllDBs(func(testDB db.TxQueueDB) {
		for _, testTx := range testTxes {
			err := testDB.StoreRawTx(t.GetTestContext(), testTx, testTx.ChainId())
			Nil(t.T(), err)

			// TODO: retreive index
		}

	})
}
