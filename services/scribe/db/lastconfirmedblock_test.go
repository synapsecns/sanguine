package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/scribe/db"
)

func (t *DBSuite) TestStoreRetrieveLastConfirmedBlock() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		chainID := gofakeit.Uint32()

		// Before storing, ensure that the last confirmed block is 0.
		retrievedLastConfirmedBlock, err := testDB.RetrieveLastConfirmedBlock(t.GetTestContext(), chainID)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastConfirmedBlock, uint64(0))

		// Store a new last confirmed block.
		err = testDB.StoreLastConfirmedBlock(t.GetTestContext(), chainID, 1)
		Nil(t.T(), err)

		// Ensure the last confirmed block matches the one stored.
		retrievedLastConfirmedBlock, err = testDB.RetrieveLastConfirmedBlock(t.GetTestContext(), chainID)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastConfirmedBlock, uint64(1))

		// Update the last confirmed block to a new value.
		err = testDB.StoreLastConfirmedBlock(t.GetTestContext(), chainID, 2)
		Nil(t.T(), err)

		// Ensure the last confirmed block matches the one stored.
		retrievedLastConfirmedBlock, err = testDB.RetrieveLastConfirmedBlock(t.GetTestContext(), chainID)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastConfirmedBlock, uint64(2))
	})
}
