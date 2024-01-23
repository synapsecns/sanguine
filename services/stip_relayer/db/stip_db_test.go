package db_test

import (
	"time"

	"github.com/synapsecns/sanguine/services/stip-relayer/db"
)

func (d *DBSuite) TestGetSTIPTransactionsNotRebated() {
	d.RunOnAllDBs(func(testDB db.STIPDB) {
		// Arrange: Create and insert a transaction
		expectedTransaction := &db.STIPTransactions{
			// Initialize all fields
			ExecutionID: "123456",
			Address:     "0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE",
			Amount:      100.0,
			AmountUSD:   200.0,
			ArbPrice:    300.0,
			BlockTime:   time.Now(),
			Direction:   "inbound",
			Hash:        "0xabc123",
			Module:      "module1",
			Token:       "token1",
			TokenPrice:  400.0,
			Rebated:     false,
			Nonce:       1,
		}
		err := testDB.InsertNewStipTransactions(d.GetTestContext(), []db.STIPTransactions{*expectedTransaction})
		d.Require().NoError(err)

		// Act: Retrieve transactions by DestChainID and DestTokenAddr that have not been rebated
		transactions, err := testDB.GetSTIPTransactionsNotRebated(d.GetTestContext())
		d.Require().NoError(err)

		// Assert: Check if the retrieved transactions match the inserted transaction
		d.Len(transactions, 1)
		d.Equal(expectedTransaction.Address, transactions[0].Address)
		d.Equal(expectedTransaction.Hash, transactions[0].Hash)
		d.Equal(expectedTransaction.Rebated, transactions[0].Rebated)
	})
}

func (d *DBSuite) TestUpdateSTIPTransactionRebated() {}

func (d *DBSuite) TestInsertNewStipTransactions() {}
