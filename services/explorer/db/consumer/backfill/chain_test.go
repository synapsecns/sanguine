package backfill_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/backfill"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	"math/big"
)

func (b *BackfillSuite) TestBackfill() {
	contractAddressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	contractAddressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	chainID := gofakeit.Uint32()

	// Store 20 logs
	for blockNumber := 0; blockNumber < 10; blockNumber++ {
		storeLogA := testutil.BuildLog(contractAddressA, uint64(blockNumber), &b.logIndex)
		err := b.eventDB.StoreLog(b.GetTestContext(), storeLogA, chainID)
		Nil(b.T(), err)
		storeLogB := testutil.BuildLog(contractAddressB, uint64(blockNumber), &b.logIndex)
		err = b.eventDB.StoreLog(b.GetTestContext(), storeLogB, chainID)
		Nil(b.T(), err)
	}

	// setup a ChainBackfiller
	_, bridgeConfigContract := b.deployManager.GetBridgeConfigV3(b.GetTestContext(), b.testBackend)

	bcf, err := consumer.NewBridgeConfigFetcher(bridgeConfigContract.Address(), b.testBackend)
	bp, err := consumer.NewBridgeParser(b.db, contractAddressA, *bcf)
	Nil(b.T(), err)
	sp, err := consumer.NewSwapParser(b.db, contractAddressB)
	Nil(b.T(), err)
	spMap := map[common.Address]*consumer.SwapParser{}
	spMap[contractAddressB] = sp
	f := consumer.NewFetcher(b.gqlClient)
	chainBackfiller := backfill.NewChainBackfiller(chainID, b.db, 3, bp, contractAddressA, spMap, *f, b.bridgeConfigContract.Address())

	// backfill the blocks
	err = chainBackfiller.Backfill(b.GetTestContext(), 0, 9)
	Nil(b.T(), err)

	// check that the blocks were backfilled
	// TODO: check that the blocks were backfilled
}
