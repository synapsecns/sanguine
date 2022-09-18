package backfill_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/backfill"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/parser"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	"math/big"
)

func (t *BackfillSuite) TestBackfill() {
	contractAddressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	contractAddressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	chainID := gofakeit.Uint32()

	// Store 20 logs
	for blockNumber := 0; blockNumber < 10; blockNumber++ {
		storeLogA := testutil.BuildLog(contractAddressA, uint64(blockNumber), &t.logIndex)
		err := t.eventDB.StoreLog(t.GetTestContext(), storeLogA, chainID)
		Nil(t.T(), err)
		storeLogB := testutil.BuildLog(contractAddressB, uint64(blockNumber), &t.logIndex)
		err = t.eventDB.StoreLog(t.GetTestContext(), storeLogB, chainID)
		Nil(t.T(), err)
	}

	// setup a ChainBackfiller
	bcf, err := bridgeconfig.NewFetcher()
	bp, err := parser.NewBridgeParser(t.db, contractAddressA)
	Nil(t.T(), err)
	sp, err := parser.NewSwapParser(t.db, contractAddressB)
	Nil(t.T(), err)
	spMap := map[common.Address]*parser.SwapParser{}
	spMap[contractAddressB] = sp
	f := consumer.NewFetcher(t.gqlClient)
	chainBackfiller := backfill.NewChainBackfiller(chainID, t.db, 3, bp, contractAddressA, spMap, *f)

	// backfill the blocks
	err = chainBackfiller.Backfill(t.GetTestContext(), 0, 9)
	Nil(t.T(), err)

	// check that the blocks were backfilled
	// TODO: check that the blocks were backfilled
}
