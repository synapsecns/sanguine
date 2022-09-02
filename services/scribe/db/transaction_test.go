package db_test

import (
	"math/big"

	"github.com/synapsecns/sanguine/services/scribe/db"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

// addressPtr converts an address to a pointer.
func addressPtr(address common.Address) *common.Address {
	return &address
}

// testTxes is a list of txes to test inserting different tx types
// because we are not testing a raw tx.
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

func (t *DBSuite) TestTxInsertion() {
	testWallet, err := wallet.FromRandom()
	Nil(t.T(), err)

	signer := localsigner.NewSigner(testWallet.PrivateKey())

	t.RunOnAllDBs(func(testDB db.EventDB) {
		for _, testTx := range testTxes {
			transactor, err := signer.GetTransactor(testTx.ChainId())
			Nil(t.T(), err)

			signedTx, err := transactor.Signer(signer.Address(), testTx)
			Nil(t.T(), err)

			err = testDB.StoreEthTx(t.GetTestContext(), signedTx, uint32(testTx.ChainId().Uint64()), gofakeit.Uint64())
			Nil(t.T(), err)
			// TODO: retrieve the processed tx
		}
	})
}
