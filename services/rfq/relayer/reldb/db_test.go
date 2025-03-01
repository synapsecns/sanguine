package reldb_test

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/listener"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

func (d *DBSuite) TestBlock() {
	d.RunOnAllDBs(func(testDB reldb.Service) {
		const chainIDA = 5
		const chainIDB = 6
		_, err := testDB.LatestBlockForChain(d.GetTestContext(), chainIDA)
		d.True(errors.Is(err, listener.ErrNoLatestBlockForChainID))
		_, err = testDB.LatestBlockForChain(d.GetTestContext(), chainIDB)
		d.True(errors.Is(err, listener.ErrNoLatestBlockForChainID))

		testHeightA := 10
		testHeightB := 100

		err = testDB.PutLatestBlock(d.GetTestContext(), chainIDA, uint64(testHeightA))
		d.NoError(err)

		err = testDB.PutLatestBlock(d.GetTestContext(), chainIDB, uint64(testHeightB))
		d.NoError(err)

		lastHeight, err := testDB.LatestBlockForChain(d.GetTestContext(), chainIDA)
		d.NoError(err)
		d.Equal(lastHeight, uint64(testHeightA))

		lastHeight, err = testDB.LatestBlockForChain(d.GetTestContext(), chainIDB)
		d.NoError(err)
		d.Equal(lastHeight, uint64(testHeightB))

		testHeightA++
		err = testDB.PutLatestBlock(d.GetTestContext(), chainIDA, uint64(testHeightA))
		d.NoError(err)
		lastHeight, err = testDB.LatestBlockForChain(d.GetTestContext(), chainIDA)
		d.NoError(err)
		d.Equal(lastHeight, uint64(testHeightA))

		testHeightB++
		err = testDB.PutLatestBlock(d.GetTestContext(), chainIDB, uint64(testHeightB))
		d.NoError(err)
		lastHeight, err = testDB.LatestBlockForChain(d.GetTestContext(), chainIDB)
		d.NoError(err)
		d.Equal(lastHeight, uint64(testHeightB))
	})
}

func (d *DBSuite) TestStoreAndUpdateRebalance() {
	d.RunOnAllDBs(func(testDB reldb.Service) {
		rebalance := reldb.Rebalance{
			Origin:       1,
			Destination:  10,
			OriginAmount: big.NewInt(100),
			Status:       reldb.RebalanceInitiated,
		}

		// make sure no rebalances are pending
		pending, err := testDB.GetPendingRebalances(d.GetTestContext(), rebalance.Origin)
		d.Nil(err)
		d.False(len(pending) > 0)
		pending, err = testDB.GetPendingRebalances(d.GetTestContext(), rebalance.Destination)
		d.Nil(err)
		d.False(len(pending) > 0)

		// store rebalance
		err = testDB.StoreRebalance(d.GetTestContext(), rebalance)
		d.Nil(err)
		pending, err = testDB.GetPendingRebalances(d.GetTestContext(), rebalance.Origin)
		d.Nil(err)
		d.True(len(pending) > 0)
		pending, err = testDB.GetPendingRebalances(d.GetTestContext(), rebalance.Destination)
		d.Nil(err)
		d.True(len(pending) > 0)

		// update rebalance to pending
		rebalanceID := "1-1"
		rebalancePending := reldb.Rebalance{
			RebalanceID:  &rebalanceID,
			Origin:       rebalance.Origin,
			OriginTxHash: common.HexToHash("0x123"),
			Status:       reldb.RebalancePending,
		}
		err = testDB.UpdateRebalance(d.GetTestContext(), rebalancePending, true)
		d.Nil(err)
		dbRebalance, err := testDB.GetRebalanceByID(d.GetTestContext(), rebalanceID)
		d.Nil(err)
		d.Equal(rebalancePending.RebalanceID, dbRebalance.RebalanceID)
		d.Equal(rebalancePending.OriginTxHash, dbRebalance.OriginTxHash)
		d.Equal(rebalancePending.Status, dbRebalance.Status)
	})
}

func (d *DBSuite) TestStoreAndUpdateLatestRebalance() {
	d.RunOnAllDBs(func(testDB reldb.Service) {
		// store rebalance
		rebalance := reldb.Rebalance{
			Origin:       1,
			Destination:  10,
			OriginAmount: big.NewInt(100),
			Status:       reldb.RebalanceInitiated,
		}

		err := testDB.StoreRebalance(d.GetTestContext(), rebalance)
		d.Nil(err)
		pending, err := testDB.GetPendingRebalances(d.GetTestContext(), rebalance.Origin)
		d.Nil(err)
		d.True(len(pending) > 0)
		pending, err = testDB.GetPendingRebalances(d.GetTestContext(), rebalance.Destination)
		d.Nil(err)
		d.True(len(pending) > 0)

		// update rebalance to pending without id
		rebalancePending := reldb.Rebalance{
			Origin:       rebalance.Origin,
			Destination:  rebalance.Destination,
			OriginTxHash: common.HexToHash("0x123"),
			Status:       reldb.RebalancePending,
		}
		err = testDB.UpdateLatestRebalance(d.GetTestContext(), rebalancePending)
		d.Nil(err)
		dbRebalances, err := testDB.GetPendingRebalances(d.GetTestContext(), 1)
		dbRebalance := dbRebalances[0]
		d.Nil(err)
		d.Equal(rebalancePending.Origin, dbRebalance.Origin)
		d.Equal(rebalancePending.Destination, dbRebalance.Destination)
		d.Equal(rebalancePending.OriginTxHash, dbRebalance.OriginTxHash)
		d.Equal(rebalancePending.Status, dbRebalance.Status)
	})
}
