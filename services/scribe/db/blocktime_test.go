package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/scribe/db"
)

func (t *DBSuite) TestStoreRetrieveBlockTime() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		chainIDA := gofakeit.Uint32()
		chainIDB := gofakeit.Uint32()
		blockTime := uint64(gofakeit.Uint32())
		// Store 10 blocks for both chains.
		for i := uint64(0); i < 10; i++ {
			err := testDB.StoreBlockTime(t.GetTestContext(), chainIDA, i, blockTime+i)
			Nil(t.T(), err)
			err = testDB.StoreBlockTime(t.GetTestContext(), chainIDB, i, blockTime+(i*2))
			Nil(t.T(), err)
		}

		// Ensure the block time for the chain ID matches the one stored.
		for i := uint64(0); i < 10; i++ {
			retrievedBlockTimeA, err := testDB.RetrieveBlockTime(t.GetTestContext(), chainIDA, i)
			Nil(t.T(), err)
			Equal(t.T(), retrievedBlockTimeA, blockTime+i)
			retrievedBlockTimeB, err := testDB.RetrieveBlockTime(t.GetTestContext(), chainIDB, i)
			Nil(t.T(), err)
			Equal(t.T(), retrievedBlockTimeB, blockTime+(i*2))
		}

		lastBlockA, err := testDB.RetrieveLastBlockStored(t.GetTestContext(), chainIDA)
		Nil(t.T(), err)
		Equal(t.T(), lastBlockA, uint64(9))
		lastBlockB, err := testDB.RetrieveLastBlockStored(t.GetTestContext(), chainIDB)
		Nil(t.T(), err)
		Equal(t.T(), lastBlockB, uint64(9))

		firstBlockA, err := testDB.RetrieveFirstBlockStored(t.GetTestContext(), chainIDA)
		Nil(t.T(), err)
		Equal(t.T(), firstBlockA, uint64(0))
		firstBlockB, err := testDB.RetrieveFirstBlockStored(t.GetTestContext(), chainIDB)
		Nil(t.T(), err)
		Equal(t.T(), firstBlockB, uint64(0))
	})
}

func (t *DBSuite) TestRetrieveBlockTimesCountForChain() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		chainIDA := gofakeit.Uint32()
		chainIDB := gofakeit.Uint32()
		blockTime := uint64(gofakeit.Uint32())
		// Store 10 blocks for both chains.
		for i := uint64(0); i < 10; i++ {
			err := testDB.StoreBlockTime(t.GetTestContext(), chainIDA, i, blockTime+i)
			Nil(t.T(), err)
			err = testDB.StoreBlockTime(t.GetTestContext(), chainIDB, i, blockTime+(i*2))
			Nil(t.T(), err)
		}

		// Ensure the block time for the chain ID matches the one stored.
		for i := uint64(0); i < 10; i++ {
			retrievedBlockTimeA, err := testDB.RetrieveBlockTime(t.GetTestContext(), chainIDA, i)
			Nil(t.T(), err)
			Equal(t.T(), retrievedBlockTimeA, blockTime+i)
			retrievedBlockTimeB, err := testDB.RetrieveBlockTime(t.GetTestContext(), chainIDB, i)
			Nil(t.T(), err)
			Equal(t.T(), retrievedBlockTimeB, blockTime+(i*2))
		}

		blockTimeCountA, err := testDB.RetrieveBlockTimesCountForChain(t.GetTestContext(), chainIDA)
		Nil(t.T(), err)
		Equal(t.T(), int64(10), blockTimeCountA)
		blockTimeCountB, err := testDB.RetrieveBlockTimesCountForChain(t.GetTestContext(), chainIDB)
		Nil(t.T(), err)
		Equal(t.T(), int64(10), blockTimeCountB)
	})
}
