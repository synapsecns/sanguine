package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/db"
)

func (t *DBSuite) TestStoreRetreiveMessageLatestBlockEnd() {
	const testDomain = 10

	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		height, err := testDB.GetMessageLatestBlockEnd(t.GetTestContext(), testDomain)
		ErrorIs(t.T(), err, db.ErrNoStoredBlockForChain, "expected an error when no height is stored")
		Zerof(t.T(), height, "expected non-existent height")

		testHeight := uint32(gofakeit.Uint16())

		// store again
		err = testDB.StoreMessageLatestBlockEnd(t.GetTestContext(), testDomain, testHeight)
		Nil(t.T(), err)

		// store a different height on another chain to see if we break anything
		err = testDB.StoreMessageLatestBlockEnd(t.GetTestContext(), uint32(testDomain+1+gofakeit.Uint16()), testHeight)
		Nil(t.T(), err)

		height, err = testDB.GetMessageLatestBlockEnd(t.GetTestContext(), testDomain)
		Nil(t.T(), err)
		Equal(t.T(), height, testHeight)
	})
}
