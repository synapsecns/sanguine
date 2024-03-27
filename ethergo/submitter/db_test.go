package submitter_test

import (
	"errors"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/google/uuid"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"

	"math/big"

	"github.com/synapsecns/sanguine/ethergo/util"
)

func (t *TXSubmitterDBSuite) TestGetNonceForChainID() {
	t.RunOnAllDBs(func(testDB db.Service) {
		for _, backend := range t.testBackends {
			manager := t.managers[backend.GetChainID()]

			for _, mockAccount := range t.mockAccounts {
				// make sure first nonce errors
				_, err := testDB.GetNonceForChainID(t.GetTestContext(), mockAccount.Address, backend.GetBigChainID())
				t.Require().True(errors.Is(err, db.ErrNoNonceForChain))

				for i := 0; i < 4; i++ {
					tx, err := manager.SignTx(types.NewTx(&types.LegacyTx{
						To:    &mockAccount.Address,
						Value: new(big.Int).SetUint64(gofakeit.Uint64()),
					}), backend.Signer(), mockAccount.PrivateKey)

					t.Require().NoError(err)

					// make sure nonces are correct
					err = testDB.PutTXS(t.GetTestContext(), db.TX{
						Transaction: tx,
						Status:      db.Pending,
					})
					t.Require().NoError(err)

					dbNonce, err := testDB.GetNonceForChainID(t.GetTestContext(), mockAccount.Address, backend.GetBigChainID())
					t.Require().NoError(err)
					t.Require().Equal(dbNonce, tx.Nonce())
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
				for i := 0; i < 50; i++ {
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
					err := testDB.PutTXS(t.GetTestContext(), db.NewTX(tx, db.Pending, uuid.New().String()))
					t.Require().NoError(err)

					// add a copy of the tx w/ a hardcoded gas price we can use to identify the created at time. This should be returned since it's
					// the latest created at
					copiedTx, err := util.CopyTX(tx, util.WithGasPrice(big.NewInt(1)))
					t.Require().NoError(err)
					// sign it
					copiedTx, err = manager.SignTx(copiedTx, backend.Signer(), mockAccount.PrivateKey, nonce.WithNoBump(true))
					t.Require().NoError(err)
					err = testDB.PutTXS(t.GetTestContext(), db.NewTX(copiedTx, db.Pending, uuid.New().String()))

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
					msg, err := util.TxToCall(tx.Transaction)
					t.Require().NoError(err)

					t.Require().Equal(mockAccount.Address, msg.From)
				}

				// check that the result is ordered by nonce
				for i := 0; i < len(result)-1; i++ {
					t.Require().Less(result[i].Nonce(), result[i+1].Nonce())
					// make sure the gas price is correct and the most recently created has been fetched
					t.Require().Equal(result[i].GasPrice(), big.NewInt(1), testsuite.BigIntComparer())
				}

				// make sure this returns double the number of results, 2 per tx
				// TODO: check nonces
				result, err = testDB.GetAllTXAttemptByStatus(t.GetTestContext(), mockAccount.Address, backend.GetBigChainID(), db.Pending)
				t.Require().NoError(err)
				t.Require().Equal(txdb.MaxResultsPerChain*2, len(result))
			}
		}
	})
}

func (t *TXSubmitterDBSuite) TestGetNonceStatus() {
	t.RunOnAllDBs(func(dbs db.Service) {
		_, err := dbs.GetNonceStatus(t.GetTestContext(), mocks.MockAddress(), big.NewInt(1), 4)
		t.Require().ErrorIs(err, db.ErrNonceNotExist)

		simulatedBackend := simulated.NewSimulatedBackend(t.GetTestContext(), t.T())
		acct := simulatedBackend.GetFundedAccount(t.GetTestContext(), big.NewInt(params.Ether))
		mockTx := mocks.MockTx(t.GetTestContext(), t.T(), simulatedBackend, acct, types.LegacyTxType)

		for i, status := range db.AllStatusTypes() {
			copiedTX, err := util.CopyTX(mockTx, util.WithGasPrice(big.NewInt(int64(i))))
			t.Require().NoError(err)

			copiedTX, err = types.SignTx(copiedTX, simulatedBackend.Signer(), acct.PrivateKey)
			t.Require().NoError(err)

			err = dbs.PutTXS(t.GetTestContext(), db.TX{
				Transaction: copiedTX,
				Status:      status,
			})
			t.Require().NoError(err)

			msg, err := util.TxToCall(copiedTX)
			t.Require().NoError(err)

			nonceStatus, err := dbs.GetNonceStatus(t.GetTestContext(), msg.From, simulatedBackend.GetBigChainID(), mockTx.Nonce())
			t.Require().NoError(err)

			t.Require().Equal(status, nonceStatus)

			txs, err := dbs.GetNonceAttemptsByStatus(t.GetTestContext(), msg.From, simulatedBackend.GetBigChainID(), mockTx.Nonce(), status)
			t.Require().NoError(err)

			t.Require().Equal(1, len(txs))
		}
	})
}

func (t *TXSubmitterDBSuite) TestGetChainIDsByStatus() {
	t.RunOnAllDBs(func(testDB db.Service) {
		chainIDToStatus := map[int64]db.Status{
			1: db.Pending,
			3: db.Stored,
			4: db.FailedSubmit,
		}
		expectedPendingChainIDs := []int64{1}

		for _, mockAccount := range t.mockAccounts {
			for _, backend := range t.testBackends {
				manager := t.managers[backend.GetChainID()]

				// create some test transactions
				var txs []*types.Transaction
				for i := 0; i < 50; i++ {
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
					err := testDB.PutTXS(t.GetTestContext(), db.NewTX(tx, chainIDToStatus[backend.GetBigChainID().Int64()], uuid.New().String()))
					t.Require().NoError(err)
				}
			}

			// check which chainIDs are stored with pending status
			result, err := testDB.GetChainIDsByStatus(t.GetTestContext(), mockAccount.Address, db.Pending)
			t.Require().NoError(err)

			resultInt64 := make([]int64, len(result))
			for i, chainID := range result {
				resultInt64[i] = chainID.Int64()
			}
			t.Equal(expectedPendingChainIDs, resultInt64)
		}
	})
}
