package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/scribe/db"
)

func (t *DBSuite) TestStoreRetrieveLastBlockTime() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		chainIDA := gofakeit.Uint32()
		chainIDB := gofakeit.Uint32()
		lastBlockTime := gofakeit.Uint64()

		// Before storing, ensure that the last block time is 0 and err is returned (there is nothing in the database)
		retrievedLastBlockTimeA, err := testDB.RetrieveLastBlockTime(t.GetTestContext(), chainIDA)
		NotNil(t.T(), err)
		Equal(t.T(), retrievedLastBlockTimeA, uint64(0))
		retrievedLastBlockTimeB, err := testDB.RetrieveLastBlockTime(t.GetTestContext(), chainIDB)
		NotNil(t.T(), err)
		Equal(t.T(), retrievedLastBlockTimeB, uint64(0))

		// Store a new chain ID and last block time.
		err = testDB.StoreLastBlockTime(t.GetTestContext(), chainIDA, lastBlockTime)
		Nil(t.T(), err)

		// Store a new chain ID and last block time.
		err = testDB.StoreLastBlockTime(t.GetTestContext(), chainIDB, lastBlockTime)
		Nil(t.T(), err)

		// Ensure the last indexed for the chain ID matches the one stored.
		retrievedLastBlockTimeA, err = testDB.RetrieveLastBlockTime(t.GetTestContext(), chainIDA)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastBlockTimeA, lastBlockTime)
		retrievedLastBlockTimeB, err = testDB.RetrieveLastBlockTime(t.GetTestContext(), chainIDB)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastBlockTimeB, lastBlockTime)

		// Update chainIDA's last block time to a new value.
		err = testDB.StoreLastBlockTime(t.GetTestContext(), chainIDA, lastBlockTime+1)
		Nil(t.T(), err)

		// Ensure the last indexed for the chain ID matches the one stored.
		retrievedLastBlockTimeA, err = testDB.RetrieveLastBlockTime(t.GetTestContext(), chainIDA)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastBlockTimeA, lastBlockTime+1)
		retrievedLastBlockTimeB, err = testDB.RetrieveLastBlockTime(t.GetTestContext(), chainIDB)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastBlockTimeB, lastBlockTime)

		// Store the second chain ID's last block time.
		err = testDB.StoreLastBlockTime(t.GetTestContext(), chainIDB, lastBlockTime)
		Nil(t.T(), err)

		// Ensure the last indexed for the chain ID matches the one stored.
		retrievedLastBlockTimeA, err = testDB.RetrieveLastBlockTime(t.GetTestContext(), chainIDA)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastBlockTimeA, lastBlockTime+1)
		retrievedLastBlockTimeB, err = testDB.RetrieveLastBlockTime(t.GetTestContext(), chainIDB)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastBlockTimeB, lastBlockTime)
	})
}
