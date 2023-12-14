package db_test

import (
	"errors"
	"github.com/synapsecns/sanguine/services/rfq/relayer/db"
)

func (d *DBSuite) TestBlock() {
	d.RunOnAllDBs(func(testDB db.Service) {
		const testChainID = 5
		_, err := testDB.LatestBlockForChain(d.GetTestContext(), testChainID)
		d.True(errors.Is(err, db.ErrNoLatestBlockForChainID))

		testHeight := 10

		err = testDB.PutLatestBlock(d.GetTestContext(), testChainID, uint64(testHeight))
		d.NoError(err)

		lastHeight, err := testDB.LatestBlockForChain(d.GetTestContext(), testChainID)
		d.NoError(err)
		d.Equal(lastHeight, uint64(testHeight))

		testHeight++
		err = testDB.PutLatestBlock(d.GetTestContext(), testChainID, uint64(testHeight))
		lastHeight, err = testDB.LatestBlockForChain(d.GetTestContext(), testChainID)
		d.NoError(err)
		d.Equal(lastHeight, uint64(testHeight))
	})
}
