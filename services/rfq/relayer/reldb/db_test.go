package reldb_test

import (
	"errors"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

func (d *DBSuite) TestBlock() {
	d.RunOnAllDBs(func(testDB reldb.Service) {
		const testChainID = 5
		_, err := testDB.LatestBlockForChain(d.GetTestContext(), testChainID)
		d.True(errors.Is(err, reldb.ErrNoLatestBlockForChainID))

		testHeight := 10

		err = testDB.PutLatestBlock(d.GetTestContext(), testChainID, uint64(testHeight))
		d.NoError(err)

		lastHeight, err := testDB.LatestBlockForChain(d.GetTestContext(), testChainID)
		d.NoError(err)
		d.Equal(lastHeight, uint64(testHeight))

		testHeight++
		err = testDB.PutLatestBlock(d.GetTestContext(), testChainID, uint64(testHeight))
		d.NoError(err)
		lastHeight, err = testDB.LatestBlockForChain(d.GetTestContext(), testChainID)
		d.NoError(err)
		d.Equal(lastHeight, uint64(testHeight))
	})
}
