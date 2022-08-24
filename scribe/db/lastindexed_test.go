package db_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/scribe/db"
)

func (t *DBSuite) TestStoreRetrieveLastIndexed() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		addressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		addressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		lastIndexed := gofakeit.Uint64()

		// Store a new contract address and last indexed.
		err := testDB.StoreLastIndexed(t.GetTestContext(), addressA, lastIndexed)
		Nil(t.T(), err)

		// Ensure the last indexed for the contract address matches the one stored.
		retrievedLastIndexed, err := testDB.RetrieveLastIndexed(t.GetTestContext(), addressA)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastIndexed, lastIndexed)

		// Update addressA's last indexed to a new value.
		err = testDB.StoreLastIndexed(t.GetTestContext(), addressA, lastIndexed+1)
		Nil(t.T(), err)

		// Ensure the last indexed for the contract address matches the one stored.
		retrievedLastIndexed, err = testDB.RetrieveLastIndexed(t.GetTestContext(), addressA)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastIndexed, lastIndexed+1)

		// Store a second contract address and last indexed.
		err = testDB.StoreLastIndexed(t.GetTestContext(), addressB, lastIndexed)
		Nil(t.T(), err)

		// Ensure the last indexed for the contract address matches the one stored.
		retrievedLastIndexed, err = testDB.RetrieveLastIndexed(t.GetTestContext(), addressB)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLastIndexed, lastIndexed)
	})
}
