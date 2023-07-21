package db_test

import (
	"github.com/synapsecns/sanguine/services/scribe/db"
	scribeTypes "github.com/synapsecns/sanguine/services/scribe/types"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
)

func (t *DBSuite) TestStoreRetrieveLastIndexed() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		addressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		addressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		chainID := gofakeit.Uint32()
		lastIndexed := gofakeit.Uint64()

		// Before storing, ensure that the last indexed block is 0.
		retrievedLastIndexed, err := testDB.RetrieveLastIndexed(t.GetTestContext(), addressA, chainID, scribeTypes.IndexingConfirmed)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastIndexed, uint64(0))

		// Store a new contract address and last indexed.
		err = testDB.StoreLastIndexed(t.GetTestContext(), addressA, chainID, lastIndexed, scribeTypes.IndexingConfirmed)
		Nil(t.T(), err)

		// Ensure the last indexed for the contract address matches the one stored.
		retrievedLastIndexed, err = testDB.RetrieveLastIndexed(t.GetTestContext(), addressA, chainID, scribeTypes.IndexingConfirmed)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastIndexed, lastIndexed)

		// Update addressA's last indexed to a new value.
		err = testDB.StoreLastIndexed(t.GetTestContext(), addressA, chainID, lastIndexed+1, scribeTypes.IndexingConfirmed)
		Nil(t.T(), err)

		// Ensure the last indexed for the contract address matches the one stored.
		retrievedLastIndexed, err = testDB.RetrieveLastIndexed(t.GetTestContext(), addressA, chainID, scribeTypes.IndexingConfirmed)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastIndexed, lastIndexed+1)

		// Store a second contract address and last indexed.
		err = testDB.StoreLastIndexed(t.GetTestContext(), addressB, chainID+1, lastIndexed, scribeTypes.IndexingConfirmed)
		Nil(t.T(), err)

		// Ensure the last indexed for the contract address matches the one stored.
		retrievedLastIndexed, err = testDB.RetrieveLastIndexed(t.GetTestContext(), addressB, chainID+1, scribeTypes.IndexingConfirmed)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastIndexed, lastIndexed)
	})
}

func (t *DBSuite) TestStoreRetrieveLastIndexedMultiple() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		addressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		addressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		chainID := gofakeit.Uint32()
		lastIndexed := gofakeit.Uint64()

		// Before storing, ensure that the last indexed block is 0.
		retrievedLastIndexed, err := testDB.RetrieveLastIndexed(t.GetTestContext(), addressA, chainID, scribeTypes.IndexingConfirmed)
		Nil(t.T(), err)
		Equal(t.T(), uint64(0), retrievedLastIndexed)

		// Store a new contract address and last indexed.
		err = testDB.StoreLastIndexedMultiple(t.GetTestContext(), []common.Address{addressA, addressB}, chainID, lastIndexed)
		Nil(t.T(), err)

		// Ensure the last indexed for the contract address matches the one stored.
		retrievedLastIndexedMap, err := testDB.RetrieveLastIndexedMultiple(t.GetTestContext(), []common.Address{addressA, addressB}, chainID)
		Nil(t.T(), err)
		Equal(t.T(), lastIndexed, retrievedLastIndexedMap[addressA])
		Equal(t.T(), lastIndexed, retrievedLastIndexedMap[addressB])
	})
}
