package backfill_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/backfill"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	"math/big"
)

//import (
//	"fmt"
//	"github.com/brianvoe/gofakeit/v6"
//	"github.com/ethereum/go-ethereum/common"
//	. "github.com/stretchr/testify/assert"
//	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
//	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
//	"github.com/synapsecns/sanguine/services/explorer/db/consumer"
//	"github.com/synapsecns/sanguine/services/explorer/db/consumer/backfill"
//	"github.com/synapsecns/sanguine/services/explorer/db/sql"
//	"github.com/synapsecns/sanguine/services/explorer/testutil"
//	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
//	swapTypes "github.com/synapsecns/sanguine/services/explorer/types/swap"
//	"math/big"
//)

func (b *BackfillSuite) TestBackfill() {
	chainID := gofakeit.Uint32()

	manager := testutil.NewDeployManager(b.T())
	simulatedChain := simulated.NewSimulatedBackendWithChainID(b.GetTestContext(), b.T(), big.NewInt(int64(chainID)))
	_, _ = manager.GetSynapseBridge(b.GetTestContext(), simulatedChain)
	//bridgeContract, bridgeRef := b.deployManager.GetSynapseBridge(b.GetTestContext(), b.testBackend)
	//_ = bridgeContract
	//_ = bridgeRef

	//_, _ = b.deployManager.GetSwapFlashLoan(b.GetTestContext(), b.testBackend)

	contractAddressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	contractAddressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))

	bridgeTopicsMap := bridge.TopicMap()
	_ = bridgeTopicsMap
	swapTopicsMap := swap.TopicMap()
	_ = swapTopicsMap

	//transactOpts := b.testBackend.GetTxContext(b.GetTestContext(), nil)
	//tx, err := bridgeRef.Deposit(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(gofakeit.Int64()))
	//Nil(b.T(), err)
	//b.testBackend.WaitForConfirmation(b.GetTestContext(), tx)
	//
	//txData := tx.Data()
	//_ = txData

	//storeLog := testutil.BuildLog(bridgeContract.Address(), 0, &b.logIndex)
	//storeLog.Topics[0] = bridgeTopicsMap[bridgeTypes.DepositEvent]
	//storeLog.Data = txData
	//testChainID, err := b.testBackend.ChainID(b.GetTestContext())
	//Nil(b.T(), err)
	//err = b.eventDB.StoreLog(b.GetTestContext(), storeLog, uint32(testChainID.Uint64()))
	//Nil(b.T(), err)

	//// Store 20 logs
	//for blockNumber := 0; blockNumber < 10; blockNumber++ {
	//	bridgeTopicHash := bridgeTopicsMap[bridgeTypes.EventType(uint8(blockNumber%len(bridgeTopicsMap)))]
	//	fmt.Println("bridgeTopicHash", bridgeTopicHash)
	//	storeLogA := testutil.BuildLog(contractAddressA, uint64(blockNumber), &b.logIndex)
	//	storeLogA.Topics[0] = bridgeTopicHash
	//	err := b.eventDB.StoreLog(b.GetTestContext(), storeLogA, chainID)
	//	Nil(b.T(), err)
	//	swapTopicHash := swapTopicsMap[swapTypes.EventType(uint8(blockNumber%len(swapTopicsMap)))]
	//	storeLogB := testutil.BuildLog(contractAddressB, uint64(blockNumber), &b.logIndex)
	//	storeLogB.Topics[0] = swapTopicHash
	//	err = b.eventDB.StoreLog(b.GetTestContext(), storeLogB, chainID)
	//	Nil(b.T(), err)
	//}

	// setup a ChainBackfiller
	bcf, err := consumer.NewBridgeConfigFetcher(b.bridgeConfigContract.Address(), b.testBackend)
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
	swapEvents := b.db.DB().WithContext(b.GetTestContext()).Find(&sql.SwapEvent{})
	Nil(b.T(), swapEvents.Error)
	Equal(b.T(), 10, swapEvents.RowsAffected)
	bridgeEvents := b.db.DB().WithContext(b.GetTestContext()).Find(&sql.BridgeEvent{})
	Nil(b.T(), bridgeEvents.Error)
	Equal(b.T(), 10, bridgeEvents.RowsAffected)
}
