package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/scribe/db"
)

func (t *DBSuite) TestRetrieveLastBlockStoredVerbose() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		chainIDA := gofakeit.Uint32()
		chainIDB := gofakeit.Uint32()
		// Store 10 blocks for both chains.

		err := testDB.StoreLastBlockTime(t.GetTestContext(), chainIDA, 9)
		Nil(t.T(), err)
		err = testDB.StoreLastBlockTime(t.GetTestContext(), chainIDB, 9)
		Nil(t.T(), err)

		// Ensure the block time for the chain ID matches the one stored.
		retrievedBlockTimeA, err := testDB.RetrieveLastBlockStoredVerbose(t.GetTestContext(), chainIDA)
		Nil(t.T(), err)
		Equal(t.T(), 9, retrievedBlockTimeA.BlockNumber)
		Equal(t.T(), int(chainIDA), retrievedBlockTimeA.ChainID)

		retrievedBlockTimeB, err := testDB.RetrieveLastBlockStoredVerbose(t.GetTestContext(), chainIDB)
		Nil(t.T(), err)
		Equal(t.T(), 9, retrievedBlockTimeB.BlockNumber)
		Equal(t.T(), int(chainIDB), retrievedBlockTimeB.ChainID)
		NotNil(t.T(), retrievedBlockTimeB.CreatedAt)

		err = testDB.StoreLastBlockTime(t.GetTestContext(), chainIDB, 10)
		Nil(t.T(), err)

		retrievedBlockTimeB, err = testDB.RetrieveLastBlockStoredVerbose(t.GetTestContext(), chainIDB)
		Nil(t.T(), err)
		Equal(t.T(), 10, retrievedBlockTimeB.BlockNumber)
		NotNil(t.T(), retrievedBlockTimeB.UpdatedAt)
	})
}
