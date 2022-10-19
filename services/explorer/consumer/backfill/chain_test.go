package backfill_test

import (
	gosql "database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/explorer/consumer"
	"github.com/synapsecns/sanguine/services/explorer/consumer/backfill"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/testutil/testcontracts"
	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
	swapTypes "github.com/synapsecns/sanguine/services/explorer/types/swap"
	"math/big"
)

func arrayToTokenIndexMap(input []*big.Int) map[uint8]string {
	output := map[uint8]string{}
	for i := range input {
		output[uint8(i)] = input[i].String()
	}
	return output
}
func (b *BackfillSuite) TestBackfill() {
	testChainID := b.testBackend.GetBigChainID()
	bridgeContract, bridgeRef := b.testDeployManager.GetTestSynapseBridge(b.GetTestContext(), b.testBackend)
	swapContractA, swapRefA := b.testDeployManager.GetTestSwapFlashLoan(b.GetTestContext(), b.testBackend)
	testDeployManagerB := testcontracts.NewDeployManager(b.T())
	swapContractB, swapRefB := testDeployManagerB.GetTestSwapFlashLoan(b.GetTestContext(), b.testBackend)

	transactOpts := b.testBackend.GetTxContext(b.GetTestContext(), nil)

	// Store every bridge event.
	bridgeTx, err := bridgeRef.TestDeposit(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	depositLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 2)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestRedeem(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(int64(gofakeit.Uint32()))), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	redeemLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 2)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestWithdraw(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	withdrawLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 2)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestMint(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	mintLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 3)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestDepositAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	depositAndSwapLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 6)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestRedeemAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	redeemAndSwapLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 7)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestRedeemAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	redeemAndRemoveLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 8)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestMintAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	mintAndSwapLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 9)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestWithdrawAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	withdrawAndRemoveLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 10)
	Nil(b.T(), err)
	bridgeTx, err = bridgeRef.TestRedeemV2(transactOpts.TransactOpts, [32]byte{byte(gofakeit.Uint64())}, big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	redeemV2Log, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 11)
	Nil(b.T(), err)

	// Store every swap event across two different swap contracts.
	swapTx, err := swapRefA.TestSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	swapLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 0)
	Nil(b.T(), err)
	swapTx, err = swapRefB.TestAddLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	addLiquidityLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 0)
	Nil(b.T(), err)
	swapTx, err = swapRefB.TestRemoveLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	removeLiquidityLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 0)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestRemoveLiquidityOne(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	removeLiquidityOneLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 1)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestRemoveLiquidityImbalance(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	removeLiquidityImbalanceLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 1)
	Nil(b.T(), err)
	swapTx, err = swapRefB.TestNewAdminFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	newAdminFeeLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 5)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestNewSwapFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	newSwapFeeLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 6)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	rampALog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 7)
	Nil(b.T(), err)
	swapTx, err = swapRefB.TestStopRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	stopRampALog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 8)
	Nil(b.T(), err)
	swapTx, err = swapRefA.TestFlashLoan(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	flashLoanLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 9)
	Nil(b.T(), err)

	err = b.eventDB.StoreLastBlockTime(b.GetTestContext(), uint32(testChainID.Uint64()), 12)
	Nil(b.T(), err)

	// set up a ChainBackfiller
	bcf, err := consumer.NewBridgeConfigFetcher(b.bridgeConfigContract.Address(), b.testBackend)
	Nil(b.T(), err)
	bp, err := consumer.NewBridgeParser(b.db, bridgeContract.Address(), *bcf, b.consumerFetcher)
	Nil(b.T(), err)

	// srB is the swap ref for getting token data
	srA, err := consumer.NewSwapFetcher(swapContractA.Address(), b.testBackend)
	Nil(b.T(), err)
	spA, err := consumer.NewSwapParser(b.db, swapContractA.Address(), *srA, b.consumerFetcher)
	Nil(b.T(), err)

	// srB is the swap ref for getting token data
	srB, err := consumer.NewSwapFetcher(swapContractB.Address(), b.testBackend)
	Nil(b.T(), err)
	spB, err := consumer.NewSwapParser(b.db, swapContractB.Address(), *srB, b.consumerFetcher)
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
	bridgeEvents := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Find(&sql.BridgeEvent{}).Count(&count)
	Nil(b.T(), bridgeEvents.Error)
	Equal(b.T(), int64(10), count)
	swapEvents := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Find(&sql.SwapEvent{}).Count(&count)
	Nil(b.T(), swapEvents.Error)
	Equal(b.T(), int64(10), count)

	// Test bridge parity
	err = b.depositParity(depositLog, bp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.redeemParity(redeemLog, bp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.withdrawParity(withdrawLog, bp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.mintParity(mintLog, bp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.depositAndSwapParity(depositAndSwapLog, bp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.redeemAndSwapParity(redeemAndSwapLog, bp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.redeemAndRemoveParity(redeemAndRemoveLog, bp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.mintAndSwapParity(mintAndSwapLog, bp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.withdrawAndRemoveParity(withdrawAndRemoveLog, bp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.redeemV2Parity(redeemV2Log, bp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)

	// Test swap parity
	err = b.swapParity(swapLog, spA, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.addLiquidityParity(addLiquidityLog, spB, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.removeLiquidityParity(removeLiquidityLog, spB, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.removeLiquidityOneParity(removeLiquidityOneLog, spA, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.removeLiquidityImbalanceParity(removeLiquidityImbalanceLog, spA, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.newAdminFeeParity(newAdminFeeLog, spB, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.newSwapFeeParity(newSwapFeeLog, spA, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.rampAParity(rampALog, spA, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.stopRampAParity(stopRampALog, spB, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.flashLoanParity(flashLoanLog, spA, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
}

func (b *BackfillSuite) storeTestLog(tx *types.Transaction, chainID uint32, blockNumber uint64) (*types.Log, error) {
	b.testBackend.WaitForConfirmation(b.GetTestContext(), tx)
	receipt, err := b.testBackend.TransactionReceipt(b.GetTestContext(), tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get receipt for transaction %s: %w", tx.Hash().String(), err)
	}
	receipt.Logs[0].BlockNumber = blockNumber
	err = b.eventDB.StoreLog(b.GetTestContext(), *receipt.Logs[0], chainID)
	if err != nil {
		return nil, fmt.Errorf("error storing swap log: %w", err)
	}
	return receipt.Logs[0], nil
}

// nolint:dupl
func (b *BackfillSuite) depositParity(log *types.Log, parser *consumer.BridgeParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseTokenDeposit(*log)
	_ = parsedLog
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	recipient := gosql.NullString{
		String: parsedLog.To.String(),
		Valid:  true,
	}
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.BridgeEvent{}).
		Where(&sql.BridgeEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       bridgeTypes.DepositEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			Token:           parsedLog.Token.String(),
			Amount:          parsedLog.Amount,

			Recipient:          recipient,
			DestinationChainID: parsedLog.ChainId,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) redeemParity(log *types.Log, parser *consumer.BridgeParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseTokenRedeem(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	recipient := gosql.NullString{
		String: parsedLog.To.String(),
		Valid:  true,
	}
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.BridgeEvent{}).
		Where(&sql.BridgeEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       bridgeTypes.RedeemEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			Token:           parsedLog.Token.String(),
			Amount:          parsedLog.Amount,

			Recipient:          recipient,
			DestinationChainID: parsedLog.ChainId,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) withdrawParity(log *types.Log, parser *consumer.BridgeParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseTokenWithdraw(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	recipient := gosql.NullString{
		String: parsedLog.To.String(),
		Valid:  true,
	}
	kappa := gosql.NullString{
		String: common.Bytes2Hex(parsedLog.Kappa[:]),
		Valid:  true,
	}
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.BridgeEvent{}).
		Where(&sql.BridgeEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       bridgeTypes.WithdrawEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			Token:           parsedLog.Token.String(),
			Amount:          parsedLog.Amount,

			Recipient: recipient,
			Fee:       parsedLog.Fee,
			Kappa:     kappa,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) mintParity(log *types.Log, parser *consumer.BridgeParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseTokenMint(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	recipient := gosql.NullString{
		String: parsedLog.To.String(),
		Valid:  true,
	}
	kappa := gosql.NullString{
		String: common.Bytes2Hex(parsedLog.Kappa[:]),
		Valid:  true,
	}
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.BridgeEvent{}).
		Where(&sql.BridgeEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       bridgeTypes.MintEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			Token:           parsedLog.Token.String(),
			Amount:          parsedLog.Amount,

			Recipient: recipient,
			Fee:       parsedLog.Fee,
			Kappa:     kappa,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) depositAndSwapParity(log *types.Log, parser *consumer.BridgeParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseTokenDepositAndSwap(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	recipient := gosql.NullString{
		String: parsedLog.To.String(),
		Valid:  true,
	}
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.BridgeEvent{}).
		Where(&sql.BridgeEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       bridgeTypes.DepositAndSwapEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			Token:           parsedLog.Token.String(),
			Amount:          parsedLog.Amount,

			Recipient:      recipient,
			TokenIndexFrom: big.NewInt(int64(parsedLog.TokenIndexFrom)),
			TokenIndexTo:   big.NewInt(int64(parsedLog.TokenIndexTo)),
			MinDy:          parsedLog.MinDy,
			Deadline:       parsedLog.Deadline,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) redeemAndSwapParity(log *types.Log, parser *consumer.BridgeParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseTokenRedeemAndSwap(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	recipient := gosql.NullString{
		String: parsedLog.To.String(),
		Valid:  true,
	}
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.BridgeEvent{}).
		Where(&sql.BridgeEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       bridgeTypes.RedeemAndSwapEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			Token:           parsedLog.Token.String(),
			Amount:          parsedLog.Amount,

			Recipient:      recipient,
			TokenIndexFrom: big.NewInt(int64(parsedLog.TokenIndexFrom)),
			TokenIndexTo:   big.NewInt(int64(parsedLog.TokenIndexTo)),
			MinDy:          parsedLog.MinDy,
			Deadline:       parsedLog.Deadline,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) redeemAndRemoveParity(log *types.Log, parser *consumer.BridgeParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseTokenRedeemAndRemove(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	recipient := gosql.NullString{
		String: parsedLog.To.String(),
		Valid:  true,
	}
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.BridgeEvent{}).
		Where(&sql.BridgeEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       bridgeTypes.RedeemAndRemoveEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			Token:           parsedLog.Token.String(),
			Amount:          parsedLog.Amount,

			Recipient:      recipient,
			SwapTokenIndex: big.NewInt(int64(parsedLog.SwapTokenIndex)),
			SwapMinAmount:  parsedLog.SwapMinAmount,
			SwapDeadline:   parsedLog.SwapDeadline,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) mintAndSwapParity(log *types.Log, parser *consumer.BridgeParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseTokenMintAndSwap(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	recipient := gosql.NullString{
		String: parsedLog.To.String(),
		Valid:  true,
	}
	kappa := gosql.NullString{
		String: common.Bytes2Hex(parsedLog.Kappa[:]),
		Valid:  true,
	}
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.BridgeEvent{}).
		Where(&sql.BridgeEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       bridgeTypes.MintAndSwapEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			Token:           parsedLog.Token.String(),
			Amount:          parsedLog.Amount,

			Recipient:      recipient,
			Fee:            parsedLog.Fee,
			TokenIndexFrom: big.NewInt(int64(parsedLog.TokenIndexFrom)),
			TokenIndexTo:   big.NewInt(int64(parsedLog.TokenIndexTo)),
			MinDy:          parsedLog.MinDy,
			Deadline:       parsedLog.Deadline,
			SwapSuccess:    big.NewInt(int64(*consumer.BoolToUint8(&parsedLog.SwapSuccess))),
			Kappa:          kappa,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) withdrawAndRemoveParity(log *types.Log, parser *consumer.BridgeParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseTokenWithdrawAndRemove(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	recipient := gosql.NullString{
		String: parsedLog.To.String(),
		Valid:  true,
	}
	kappa := gosql.NullString{
		String: common.Bytes2Hex(parsedLog.Kappa[:]),
		Valid:  true,
	}
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.BridgeEvent{}).
		Where(&sql.BridgeEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       bridgeTypes.WithdrawAndRemoveEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			Token:           parsedLog.Token.String(),
			Amount:          parsedLog.Amount,

			Recipient:      recipient,
			SwapTokenIndex: big.NewInt(int64(parsedLog.SwapTokenIndex)),
			SwapMinAmount:  parsedLog.SwapMinAmount,
			SwapDeadline:   parsedLog.SwapDeadline,
			SwapSuccess:    big.NewInt(int64(*consumer.BoolToUint8(&parsedLog.SwapSuccess))),
			Kappa:          kappa,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) redeemV2Parity(log *types.Log, parser *consumer.BridgeParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseTokenRedeemV2(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	recipientBytes := gosql.NullString{
		String: common.Bytes2Hex(parsedLog.To[:]),
		Valid:  true,
	}

	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.BridgeEvent{}).
		Where(&sql.BridgeEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       bridgeTypes.RedeemV2Event.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			Token:           parsedLog.Token.String(),
			Amount:          parsedLog.Amount,

			RecipientBytes:     recipientBytes,
			DestinationChainID: parsedLog.ChainId,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) swapParity(log *types.Log, parser *consumer.SwapParser, chainID uint32) error {
	// parse the log

	parsedLog, err := parser.Filterer.ParseTokenSwap(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	buyer := gosql.NullString{
		String: parsedLog.Buyer.String(),
		Valid:  true,
	}

	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.SwapEvent{}).
		Where(&sql.SwapEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       swapTypes.TokenSwapEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),

			Buyer:        buyer,
			TokensSold:   parsedLog.TokensSold,
			TokensBought: parsedLog.TokensBought,
			SoldID:       parsedLog.SoldId,
			BoughtID:     parsedLog.BoughtId,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) addLiquidityParity(log *types.Log, parser *consumer.SwapParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseAddLiquidity(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	provider := gosql.NullString{
		String: parsedLog.Provider.String(),
		Valid:  true,
	}
	var storedLog sql.SwapEvent
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.SwapEvent{}).
		Where(&sql.SwapEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       swapTypes.AddLiquidityEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),

			Provider:      provider,
			Invariant:     parsedLog.Invariant,
			LPTokenSupply: parsedLog.LpTokenSupply,
		}).
		Find(&storedLog).
		Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	Equal(b.T(), arrayToTokenIndexMap(parsedLog.TokenAmounts), storedLog.Amount)
	Equal(b.T(), arrayToTokenIndexMap(parsedLog.Fees), storedLog.AmountFee)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) removeLiquidityParity(log *types.Log, parser *consumer.SwapParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseRemoveLiquidity(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	provider := gosql.NullString{
		String: parsedLog.Provider.String(),
		Valid:  true,
	}

	arrayToTokenIndexMap(parsedLog.TokenAmounts)

	var storedLog sql.SwapEvent
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.SwapEvent{}).
		Where(&sql.SwapEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       swapTypes.RemoveLiquidityEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),

			Provider:      provider,
			LPTokenSupply: parsedLog.LpTokenSupply,
		}).
		Find(&storedLog).
		Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	Equal(b.T(), arrayToTokenIndexMap(parsedLog.TokenAmounts), storedLog.Amount)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) removeLiquidityOneParity(log *types.Log, parser *consumer.SwapParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseRemoveLiquidityOne(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	provider := gosql.NullString{
		String: parsedLog.Provider.String(),
		Valid:  true,
	}

	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.SwapEvent{}).
		Where(&sql.SwapEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       swapTypes.RemoveLiquidityOneEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),

			Provider:      provider,
			LPTokenAmount: parsedLog.LpTokenAmount,
			LPTokenSupply: parsedLog.LpTokenSupply,
			BoughtID:      parsedLog.BoughtId,
			TokensBought:  parsedLog.TokensBought,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) removeLiquidityImbalanceParity(log *types.Log, parser *consumer.SwapParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseRemoveLiquidityImbalance(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	provider := gosql.NullString{
		String: parsedLog.Provider.String(),
		Valid:  true,
	}
	var storedLog sql.SwapEvent
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.SwapEvent{}).
		Where(&sql.SwapEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       swapTypes.RemoveLiquidityImbalanceEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),

			Provider:      provider,
			Invariant:     parsedLog.Invariant,
			LPTokenSupply: parsedLog.LpTokenSupply,
		}).
		Find(&storedLog).
		Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	Equal(b.T(), arrayToTokenIndexMap(parsedLog.TokenAmounts), storedLog.Amount)
	Equal(b.T(), arrayToTokenIndexMap(parsedLog.Fees), storedLog.AmountFee)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) newAdminFeeParity(log *types.Log, parser *consumer.SwapParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseNewAdminFee(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}

	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.SwapEvent{}).
		Where(&sql.SwapEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       swapTypes.NewAdminFeeEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),

			NewAdminFee: parsedLog.NewAdminFee,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) newSwapFeeParity(log *types.Log, parser *consumer.SwapParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseNewSwapFee(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}

	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.SwapEvent{}).
		Where(&sql.SwapEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       swapTypes.NewSwapFeeEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),

			NewSwapFee: parsedLog.NewSwapFee,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) rampAParity(log *types.Log, parser *consumer.SwapParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseRampA(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}

	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.SwapEvent{}).
		Where(&sql.SwapEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       swapTypes.RampAEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),

			OldA:        parsedLog.OldA,
			NewA:        parsedLog.NewA,
			InitialTime: parsedLog.InitialTime,
			FutureTime:  parsedLog.FutureTime,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) stopRampAParity(log *types.Log, parser *consumer.SwapParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseStopRampA(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}

	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.SwapEvent{}).
		Where(&sql.SwapEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       swapTypes.StopRampAEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),

			CurrentA: parsedLog.CurrentA,
			Time:     parsedLog.Time,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) flashLoanParity(log *types.Log, parser *consumer.SwapParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseFlashLoan(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}

	receiver := gosql.NullString{
		String: parsedLog.Receiver.String(),
		Valid:  true,
	}
	amountArray := map[uint8]string{parsedLog.TokenIndex: core.CopyBigInt(parsedLog.Amount).String()}
	feeArray := map[uint8]string{parsedLog.TokenIndex: core.CopyBigInt(parsedLog.AmountFee).String()}
	var storedLog sql.SwapEvent
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.SwapEvent{}).
		Where(&sql.SwapEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       swapTypes.FlashLoanEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),

			Receiver:    receiver,
			ProtocolFee: parsedLog.ProtocolFee,
		}).
		Find(&storedLog).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	Equal(b.T(), amountArray, storedLog.Amount)
	Equal(b.T(), feeArray, storedLog.AmountFee)
	return nil
}
