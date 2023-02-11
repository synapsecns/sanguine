package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"math/big"
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

	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		for _, testTx := range testTxes {
			err := testDB.StoreRawTx(t.GetTestContext(), testTx, testTx.ChainId(), signer.Address())
			Nil(t.T(), err)

			// TODO: retrieve raw tx

			transactor, err := signer.GetTransactor(t.GetTestContext(), testTx.ChainId())
			Nil(t.T(), err)

			signedTx, err := transactor.Signer(signer.Address(), testTx)
			Nil(t.T(), err)

			err = testDB.StoreProcessedTx(t.GetTestContext(), signedTx)
			Nil(t.T(), err)
			// TODO: retrieve the processed tx
		}
	})
}

// / make sure tx doesn't conflict on both chains.
func (t *DBSuite) TestTxNonceQueryMultiChain() {
	fakeTx := testTxes[0]
	fakeTx2 := testTxes[1]

	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		from := common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))

		err := testDB.StoreRawTx(t.GetTestContext(), fakeTx, fakeTx.ChainId(), from)
		Nil(t.T(), err)

		err = testDB.StoreRawTx(t.GetTestContext(), fakeTx2, fakeTx2.ChainId(), from)
		Nil(t.T(), err)

		nonce1, err := testDB.GetNonceForChainID(t.GetTestContext(), from, fakeTx.ChainId())
		Nil(t.T(), err)
		Equal(t.T(), nonce1, fakeTx.Nonce())

		nonce2, err := testDB.GetNonceForChainID(t.GetTestContext(), from, fakeTx2.ChainId())
		Nil(t.T(), err)
		Equal(t.T(), nonce2, fakeTx2.Nonce())
	})
}
