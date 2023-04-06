package submitter_test

import (
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
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
						Value: big.NewInt(0),
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
