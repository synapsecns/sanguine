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

func (t *DBSuite) TestStoreAndRetrieveEthTx() {
	testWallet, err := wallet.FromRandom()
	Nil(t.T(), err)

	signer := localsigner.NewSigner(testWallet.PrivateKey())

	t.RunOnAllDBs(func(testDB db.EventDB) {
		for _, testTx := range testTxes {
			transactor, err := signer.GetTransactor(t.GetTestContext(), testTx.ChainId())
			Nil(t.T(), err)

			signedTx, err := transactor.Signer(signer.Address(), testTx)
			Nil(t.T(), err)

			// Store same tx with different blockhash
			err = testDB.StoreEthTx(t.GetTestContext(), signedTx, uint32(testTx.ChainId().Uint64()), common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint64(), gofakeit.Uint64())
			Nil(t.T(), err)

			// err = testDB.StoreEthTxAtHead(t.GetTestContext(), signedTx, uint32(testTx.ChainId().Uint64()), common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint64(), gofakeit.Uint64())
			// Nil(t.T(), err)

			ethTxFilter := db.EthTxFilter{
				ChainID: uint32(testTx.ChainId().Uint64()),
				TxHash:  signedTx.Hash().String(),
			}
			tx, err := testDB.RetrieveEthTxsWithFilter(t.GetTestContext(), ethTxFilter, 1)
			Nil(t.T(), err)
			resA, err := tx[0].Tx.MarshalJSON()
			Nil(t.T(), err)
			resB, err := signedTx.MarshalJSON()
			Nil(t.T(), err)
			Equal(t.T(), resA, resB)
		}
	})
}

func (t *DBSuite) TestDeleteEthTxsForBlockHash() {
	testWallet, err := wallet.FromRandom()
	Nil(t.T(), err)

	signer := localsigner.NewSigner(testWallet.PrivateKey())

	t.RunOnAllDBs(func(testDB db.EventDB) {
		chainID := gofakeit.Uint32()

		// Store a tx.
		testTx := types.NewTx(&types.LegacyTx{
			Nonce:    uint64(0),
			GasPrice: new(big.Int).SetUint64(gofakeit.Uint64()),
			Gas:      gofakeit.Uint64(),
			To:       addressPtr(common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))),
			Value:    new(big.Int).SetUint64(gofakeit.Uint64()),
			Data:     []byte(gofakeit.Paragraph(1, 2, 3, " ")),
		})
		transactor, err := localsigner.NewSigner(testWallet.PrivateKey()).GetTransactor(t.GetTestContext(), testTx.ChainId())
		Nil(t.T(), err)

		signedTx, err := transactor.Signer(signer.Address(), testTx)
		Nil(t.T(), err)

		err = testDB.StoreEthTx(t.GetTestContext(), signedTx, chainID, common.BigToHash(big.NewInt(5)), uint64(0), gofakeit.Uint64())
		Nil(t.T(), err)

		// Ensure the tx is in the database,
		ethTxFilter := db.EthTxFilter{
			ChainID:   chainID,
			BlockHash: common.BigToHash(big.NewInt(5)).String(),
		}
		retrievedTxs, err := testDB.RetrieveEthTxsWithFilter(t.GetTestContext(), ethTxFilter, 1)
		Nil(t.T(), err)
		Equal(t.T(), 1, len(retrievedTxs))

		// Delete the tx.
		err = testDB.DeleteEthTxsForBlockHash(t.GetTestContext(), common.BigToHash(big.NewInt(5)), chainID)
		Nil(t.T(), err)

		// Ensure the tx is not in the database.
		retrievedTxs, err = testDB.RetrieveEthTxsWithFilter(t.GetTestContext(), ethTxFilter, 1)
		Nil(t.T(), err)
		Equal(t.T(), 0, len(retrievedTxs))
	})
}
