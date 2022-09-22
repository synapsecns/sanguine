package backfill_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer"
	"github.com/synapsecns/sanguine/services/explorer/consumer/backfill"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
	swapTypes "github.com/synapsecns/sanguine/services/explorer/types/swap"
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

	bridgeContract, bridgeRef := b.testDeployManager.GetTestSynapseBridge(b.GetTestContext(), b.testBackend)
	swapContract, swapRef := b.testDeployManager.GetTestSwapFlashLoan(b.GetTestContext(), b.testBackend)
	//bridgeContract, bridgeRef := b.deployManager.GetSynapseBridge(b.GetTestContext(), b.testBackend)
	//_ = bridgeContract
	//_ = bridgeRef

	//_, _ = b.deployManager.GetSwapFlashLoan(b.GetTestContext(), b.testBackend)

	bridgeTopicsMap := bridge.TopicMap()
	_ = bridgeTopicsMap
	swapTopicsMap := swap.TopicMap()
	_ = swapTopicsMap

	transactOpts := b.testBackend.GetTxContext(b.GetTestContext(), nil)
	transactOpts.TransactOpts.GasLimit = uint64(200000)
	recipient := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	tx, err := bridgeRef.TestDeposit(transactOpts.TransactOpts, recipient, big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(gofakeit.Int64()))
	fmt.Println("RECIP", recipient)
	Nil(b.T(), err)
	fmt.Println(common.HexToAddress(testTokens[0].TokenAddress).String())
	b.testBackend.WaitForConfirmation(b.GetTestContext(), tx)

	txData := tx.Data()
	storeLog := testutil.BuildLog(bridgeContract.Address(), 0, &b.logIndex)
	storeLog.Topics = []common.Hash{bridgeTopicsMap[bridgeTypes.DepositEvent], common.BigToHash(big.NewInt(gofakeit.Int64()))}
	storeLog.Data = txData
	fmt.Println(common.Bytes2Hex(txData))
	testChainID, err := b.testBackend.ChainID(b.GetTestContext())
	Nil(b.T(), err)
	//err = b.eventDB.StoreLog(b.GetTestContext(), storeLog, uint32(testChainID.Uint64()))
	Nil(b.T(), err)

	// Store every swap event.
	swapTx, err := swapRef.TestSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.TokenSwapEvent], common.BigToHash(big.NewInt(gofakeit.Int64()))}, swapContract.Address(), uint32(testChainID.Uint64()), 0)
	Nil(b.T(), err)
	//swapTx, err = swapRef.TestAddLiquidity(transactOpts.TransactOpts, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	//Nil(b.T(), err)
	//err = b.storeTestLog(swapTx, swapTopicsMap[swapTypes.AddLiquidityEvent], swapContract.Address(), uint32(testChainID.Uint64()), 1)
	//Nil(b.T(), err)
	//swapTx, err = swapRef.TestRemoveLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())))
	//Nil(b.T(), err)
	//err = b.storeTestLog(swapTx, swapTopicsMap[swapTypes.RemoveLiquidityEvent], swapContract.Address(), uint32(testChainID.Uint64()), 2)
	//Nil(b.T(), err)
	swapTx, err = swapRef.TestRemoveLiquidityOne(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.RemoveLiquidityOneEvent], common.BigToHash(big.NewInt(gofakeit.Int64()))}, swapContract.Address(), uint32(testChainID.Uint64()), 3)
	Nil(b.T(), err)
	// add remove liquidity imbalance
	swapTx, err = swapRef.TestNewAdminFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.NewAdminFeeEvent]}, swapContract.Address(), uint32(testChainID.Uint64()), 4)
	Nil(b.T(), err)
	swapTx, err = swapRef.TestNewSwapFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.NewSwapFeeEvent]}, swapContract.Address(), uint32(testChainID.Uint64()), 5)
	Nil(b.T(), err)
	swapTx, err = swapRef.TestRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.RampAEvent]}, swapContract.Address(), uint32(testChainID.Uint64()), 6)
	Nil(b.T(), err)
	swapTx, err = swapRef.TestStopRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.StopRampAEvent]}, swapContract.Address(), uint32(testChainID.Uint64()), 7)
	Nil(b.T(), err)

	//// Store 20 logs
	//for blockNumber := 0; blockNumber < 10; blockNumber++ {
	//	//bridgeTopicHash := bridgeTopicsMap[bridgeTypes.EventType(uint8(blockNumber%len(bridgeTopicsMap)))]
	//	//fmt.Println("bridgeTopicHash", bridgeTopicHash)
	//	//storeLogA := testutil.BuildLog(contractAddressA, uint64(blockNumber), &b.logIndex)
	//	//storeLogA.Topics[0] = bridgeTopicHash
	//	//err = b.eventDB.StoreLog(b.GetTestContext(), storeLogA, chainID)
	//	//Nil(b.T(), err)
	//	swapTopicHash := swapTopicsMap[swapTypes.EventType(uint8(blockNumber%len(swapTopicsMap)))]
	//	storeLogB := testutil.BuildLog(swapContract.Address(), uint64(blockNumber), &b.logIndex)
	//	swapStoreLog.Topics = []common.Hash{swapTopicHash, common.BigToHash(big.NewInt(gofakeit.Int64()))}
	//	swapStoreLog.Data = swapTxData
	//	err = b.eventDB.StoreLog(b.GetTestContext(), storeLogB, uint32(testChainID.Uint64()))
	//	Nil(b.T(), err)
	//}

	// setup a ChainBackfiller
	bcf, err := consumer.NewBridgeConfigFetcher(b.bridgeConfigContract.Address(), b.testBackend)
	bp, err := consumer.NewBridgeParser(b.db, bridgeContract.Address(), *bcf)
	Nil(b.T(), err)
	sp, err := consumer.NewSwapParser(b.db, swapContract.Address())
	Nil(b.T(), err)
	spMap := map[common.Address]*consumer.SwapParser{}
	spMap[swapContract.Address()] = sp
	f := consumer.NewFetcher(b.gqlClient)
	chainBackfiller := backfill.NewChainBackfiller(uint32(testChainID.Uint64()), b.db, 3, bp, bridgeContract.Address(), spMap, *f, b.bridgeConfigContract.Address())

	// backfill the blocks
	err = chainBackfiller.Backfill(b.GetTestContext(), 0, 10)
	Nil(b.T(), err)

	// check that the blocks were backfilled
	// TODO: check that the blocks were backfilled
	var count int64
	swapEvents := b.db.DB().WithContext(b.GetTestContext()).Find(&sql.SwapEvent{}).Count(&count)
	Nil(b.T(), swapEvents.Error)
	Equal(b.T(), 10, count)
	bridgeEvents := b.db.DB().WithContext(b.GetTestContext()).Find(&sql.BridgeEvent{})
	Nil(b.T(), bridgeEvents.Error)
	Equal(b.T(), 10, bridgeEvents.RowsAffected)
}

func (b *BackfillSuite) storeTestLog(swapTx *types.Transaction, topics []common.Hash, swapAddress common.Address, chainID uint32, blockNumber uint64) error {
	b.testBackend.WaitForConfirmation(b.GetTestContext(), swapTx)
	swapTxData := swapTx.Data()
	swapStoreLog := testutil.BuildLog(swapAddress, blockNumber, &b.logIndex)
	swapStoreLog.Topics = topics
	swapStoreLog.Data = swapTxData
	err := b.eventDB.StoreLog(b.GetTestContext(), swapStoreLog, chainID)
	if err != nil {
		return err
	}
	return nil
}
