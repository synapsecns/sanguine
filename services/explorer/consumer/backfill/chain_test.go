package backfill_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer"
	"github.com/synapsecns/sanguine/services/explorer/consumer/backfill"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/testutil/testcontracts"
	"math/big"
)

func (b *BackfillSuite) TestBackfill() {
	testChainID, err := b.testBackend.ChainID(b.GetTestContext())
	Nil(b.T(), err)
	bridgeContract, bridgeRef := b.testDeployManager.GetTestSynapseBridge(b.GetTestContext(), b.testBackend)
	swapContractA, swapRefA := b.testDeployManager.GetTestSwapFlashLoan(b.GetTestContext(), b.testBackend)
	testDeployManagerB := testcontracts.NewDeployManager(b.T())
	swapContractB, swapRefB := testDeployManagerB.GetTestSwapFlashLoan(b.GetTestContext(), b.testBackend)

	transactOpts := b.testBackend.GetTxContext(b.GetTestContext(), nil)
	transactOpts.TransactOpts.GasLimit = uint64(200000)

	// Store every bridge event.
	bridgeTx, err := bridgeRef.TestDeposit(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	err = b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 0)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestRedeem(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(int64(gofakeit.Uint32()))), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	err = b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 2)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestWithdraw(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	err = b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 2)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestMint(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	err = b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 2)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestDepositAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	err = b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 6)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestRedeemAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	err = b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 7)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestRedeemAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	err = b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 8)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestMintAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	err = b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 9)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestWithdrawAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	err = b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 10)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestRedeemV2(transactOpts.TransactOpts, [32]byte{byte(gofakeit.Uint64())}, big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	err = b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 11)
	Nil(b.T(), err)

	// Store every swap event across two different swap contracts.
	swapTx, err := swapRefA.TestSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 0)
	Nil(b.T(), err)
	swapTx, err = swapRefB.TestAddLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 0)
	Nil(b.T(), err)
	swapTx, err = swapRefB.TestRemoveLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 0)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestRemoveLiquidityOne(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 1)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestRemoveLiquidityImbalance(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 1)
	Nil(b.T(), err)
	swapTx, err = swapRefB.TestNewAdminFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 5)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestNewSwapFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 6)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 7)
	Nil(b.T(), err)
	swapTx, err = swapRefB.TestStopRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 8)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestFlashLoan(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	err = b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 9)
	Nil(b.T(), err)

	// set up a ChainBackfiller
	bcf, err := consumer.NewBridgeConfigFetcher(b.bridgeConfigContract.Address(), b.testBackend)
	Nil(b.T(), err)
	bp, err := consumer.NewBridgeParser(b.db, bridgeContract.Address(), *bcf)
	Nil(b.T(), err)
	spA, err := consumer.NewSwapParser(b.db, swapContractA.Address())
	Nil(b.T(), err)
	spB, err := consumer.NewSwapParser(b.db, swapContractB.Address())
	Nil(b.T(), err)
	spMap := map[common.Address]*consumer.SwapParser{}
	spMap[swapContractA.Address()] = spA
	spMap[swapContractB.Address()] = spB
	f := consumer.NewFetcher(b.gqlClient)
	chainBackfiller := backfill.NewChainBackfiller(uint32(testChainID.Uint64()), b.db, 3, bp, bridgeContract.Address(), spMap, *f, b.bridgeConfigContract.Address())

	// backfill the blocks
	err = chainBackfiller.Backfill(b.GetTestContext(), 0, 12)
	Nil(b.T(), err)
	var count int64
	swapEvents := b.db.DB().WithContext(b.GetTestContext()).Find(&sql.SwapEvent{}).Count(&count)
	Nil(b.T(), swapEvents.Error)
	Equal(b.T(), int64(10), count)
	bridgeEvents := b.db.DB().WithContext(b.GetTestContext()).Find(&sql.BridgeEvent{}).Count(&count)
	Nil(b.T(), bridgeEvents.Error)
	Equal(b.T(), int64(10), count)
}

func (b *BackfillSuite) storeTestLog(tx *types.Transaction, chainID uint32, blockNumber uint64) error {
	b.testBackend.WaitForConfirmation(b.GetTestContext(), tx)
	receipt, err := b.testBackend.TransactionReceipt(b.GetTestContext(), tx.Hash())
	if err != nil {
		return err
	}
	for _, log := range receipt.Logs {
		log.BlockNumber = blockNumber
		err = b.eventDB.StoreLog(b.GetTestContext(), *log, chainID)
		if err != nil {
			return fmt.Errorf("error storing swap log: %w", err)
		}
		return nil
	}
	return fmt.Errorf("no logs found for tx %s", tx.Hash().String())
}
