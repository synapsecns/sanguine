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
	"github.com/synapsecns/sanguine/services/explorer/testutil/testcontracts"
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
	swapContractA, swapRefA := b.testDeployManager.GetTestSwapFlashLoan(b.GetTestContext(), b.testBackend)
	testDeployManagerB := testcontracts.NewDeployManager(b.T())
	swapContractB, swapRefB := testDeployManagerB.GetTestSwapFlashLoan(b.GetTestContext(), b.testBackend)
	//bridgeContract, bridgeRef := b.deployManager.GetSynapseBridge(b.GetTestContext(), b.testBackend)
	//_ = bridgeContract
	//_ = bridgeRef

	bridgeTopicsMap := bridge.TopicMap()
	swapTopicsMap := swap.TopicMap()

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

	// Store every swap event across two different swap contracts.
	swapTx, err := swapRefA.TestSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.TokenSwapEvent], common.BigToHash(big.NewInt(gofakeit.Int64()))}, swapContractA.Address(), uint32(testChainID.Uint64()), 0)
	Nil(b.T(), err)
	swapTx, err = swapRefB.TestAddLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.TokenSwapEvent], common.BigToHash(big.NewInt(gofakeit.Int64()))}, swapContractB.Address(), uint32(testChainID.Uint64()), 0)
	Nil(b.T(), err)
	swapTx, err = swapRefB.TestRemoveLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.TokenSwapEvent], common.BigToHash(big.NewInt(gofakeit.Int64()))}, swapContractB.Address(), uint32(testChainID.Uint64()), 0)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestRemoveLiquidityOne(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.RemoveLiquidityOneEvent], common.BigToHash(big.NewInt(gofakeit.Int64()))}, swapContractA.Address(), uint32(testChainID.Uint64()), 1)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestRemoveLiquidityImbalance(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.TokenSwapEvent], common.BigToHash(big.NewInt(gofakeit.Int64()))}, swapContractA.Address(), uint32(testChainID.Uint64()), 1)
	Nil(b.T(), err)
	swapTx, err = swapRefB.TestNewAdminFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.NewAdminFeeEvent]}, swapContractB.Address(), uint32(testChainID.Uint64()), 5)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestNewSwapFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.NewSwapFeeEvent]}, swapContractA.Address(), uint32(testChainID.Uint64()), 6)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.RampAEvent]}, swapContractA.Address(), uint32(testChainID.Uint64()), 7)
	Nil(b.T(), err)
	swapTx, err = swapRefB.TestStopRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.StopRampAEvent]}, swapContractB.Address(), uint32(testChainID.Uint64()), 8)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestFlashLoan(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, []common.Hash{swapTopicsMap[swapTypes.FlashLoanEvent], common.BigToHash(big.NewInt(gofakeit.Int64()))}, swapContractA.Address(), uint32(testChainID.Uint64()), 9)
	Nil(b.T(), err)

	// setup a ChainBackfiller
	bcf, err := consumer.NewBridgeConfigFetcher(b.bridgeConfigContract.Address(), b.testBackend)
	bp, err := consumer.NewBridgeParser(b.db, bridgeContract.Address(), *bcf)
	Nil(b.T(), err)
	spA, err := consumer.NewSwapParser(b.db, swapContractA.Address())
	Nil(b.T(), err)
	spB, err := consumer.NewSwapParser(b.db, swapContractB.Address())
	spMap := map[common.Address]*consumer.SwapParser{}
	spMap[swapContractA.Address()] = spA
	spMap[swapContractB.Address()] = spB
	f := consumer.NewFetcher(b.gqlClient)
	chainBackfiller := backfill.NewChainBackfiller(uint32(testChainID.Uint64()), b.db, 3, bp, bridgeContract.Address(), spMap, *f, b.bridgeConfigContract.Address())

	// backfill the blocks
	err = chainBackfiller.Backfill(b.GetTestContext(), 0, 10)
	Nil(b.T(), err)
	var count int64
	swapEvents := b.db.DB().WithContext(b.GetTestContext()).Find(&sql.SwapEvent{}).Count(&count)
	Nil(b.T(), swapEvents.Error)
	Equal(b.T(), int64(10), count)
	//bridgeEvents := b.db.DB().WithContext(b.GetTestContext()).Find(&sql.BridgeEvent{}).Count(&count)
	//Nil(b.T(), bridgeEvents.Error)
	//Equal(b.T(), int64(10), count)
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
