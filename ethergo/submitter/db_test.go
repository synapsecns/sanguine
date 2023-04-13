package submitter_test

import (
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"
	"github.com/synapsecns/sanguine/ethergo/util"
	"math/big"
)

func (t *TXSubmitterDBSuite) TestGetNonceForChainID() {
	t.RunOnAllDBs(func(testDB db.Service) {
		for _, backend := range t.testBackends {
			manager := t.managers[backend.GetChainID()]

			for _, mockAccount := range t.mockAccounts {
				// make sure first nonce errors
				nonce, err := testDB.GetNonceForChainID(t.GetTestContext(), mockAccount.Address, backend.GetBigChainID())
				t.Require().True(errors.Is(err, db.ErrNoNonceForChain))

				for i := 0; i < 4; i++ {
					tx, err := manager.SignTx(types.NewTx(&types.LegacyTx{
						To:    &mockAccount.Address,
						Value: new(big.Int).SetUint64(gofakeit.Uint64()),
					}), backend.Signer(), mockAccount.PrivateKey)

					t.Require().NoError(err)

					// make sure nonces are correct
					err = testDB.PutTX(t.GetTestContext(), tx, db.Pending)
					t.Require().NoError(err)

					nonce, err = testDB.GetNonceForChainID(t.GetTestContext(), mockAccount.Address, backend.GetBigChainID())
					t.Require().NoError(err)
					t.Require().Equal(nonce, tx.Nonce())
				}
			}
		}
	})
}

func (t *TXSubmitterDBSuite) TestGetTransactionsWithLimitPerChainID() {
	t.RunOnAllDBs(func(testDB db.Service) {
		for _, backend := range t.testBackends {
			manager := t.managers[backend.GetChainID()]

			for _, mockAccount := range t.mockAccounts {
				// create some test transactions
				var txs []*types.Transaction
				for i := 0; i < 500; i++ {
					legacyTx := &types.LegacyTx{
						To:    &mockAccount.Address,
						Value: big.NewInt(0),
						Nonce: uint64(i),
					}
					tx, err := manager.SignTx(types.NewTx(legacyTx), backend.Signer(), mockAccount.PrivateKey)
					t.Require().NoError(err)
					txs = append(txs, tx)
				}

				// put the transactions in the database
				for _, tx := range txs {
					err := testDB.PutTX(t.GetTestContext(), tx, db.Pending)
					t.Require().NoError(err)
				}

				// get the transactions with limit per ChainID
				result, err := testDB.GetTXS(t.GetTestContext(), mockAccount.Address, backend.GetBigChainID(), db.Pending)
				t.Require().NoError(err)

				// check that the result has the correct length
				t.Require().Equal(txdb.MaxResultsPerChain, len(result))

				// check that the result is limited per ChainID and address
				for _, tx := range result {
					t.Require().Equal(backend.GetBigChainID(), tx.ChainId(), testsuite.BigIntComparer())
					msg, err := util.TxToCall(tx)
					t.Require().NoError(err)

					t.Require().Equal(mockAccount.Address, msg.From)
				}

				// check that the result is ordered by nonce
				for i := 0; i < len(result)-1; i++ {
					t.Require().Less(result[i].Nonce(), result[i+1].Nonce())
				}
			}
			// TODO: test that the result is limited per ChainID
		}
	})
}
