package backfill_test

import (
	gosql "database/sql"
	"fmt"
	scribeTypes "github.com/synapsecns/sanguine/services/scribe/types"
	"math/big"

	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher/tokenprice"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokendata"
	"github.com/synapsecns/sanguine/services/explorer/static"
	messageBusTypes "github.com/synapsecns/sanguine/services/explorer/types/messagebus"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/explorer/backfill"
	indexerConfig "github.com/synapsecns/sanguine/services/explorer/config/indexer"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	parserpkg "github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/testutil/testcontracts"
	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
	cctpTypes "github.com/synapsecns/sanguine/services/explorer/types/cctp"
	swapTypes "github.com/synapsecns/sanguine/services/explorer/types/swap"
)

func arrayToTokenIndexMap(input []*big.Int) map[uint8]string {
	output := map[uint8]string{}
	for i := range input {
		output[uint8(i)] = input[i].String()
	}
	return output
}

//nolint:maintidx
func (b *BackfillSuite) TestBackfill() {
	testChainID := b.testBackend.GetBigChainID()

	bridgeContract, bridgeRef := b.testDeployManager.GetTestSynapseBridge(b.GetTestContext(), b.testBackend)
	bridgeV1Contract, bridgeV1Ref := b.testDeployManager.GetTestSynapseBridgeV1(b.GetTestContext(), b.testBackend)
	swapContractA, swapRefA := b.testDeployManager.GetTestSwapFlashLoan(b.GetTestContext(), b.testBackend)
	metaSwapContract, metaSwapRef := b.testDeployManager.GetTestMetaSwap(b.GetTestContext(), b.testBackend)
	messageBusContract, messageBusRef := b.testDeployManager.GetTestMessageBusUpgradeable(b.GetTestContext(), b.testBackend)
	cctpContract, cctpRef := b.testDeployManager.GetTestCCTP(b.GetTestContext(), b.testBackend)
	testDeployManagerB := testcontracts.NewDeployManager(b.T())

	swapContractB, swapRefB := testDeployManagerB.GetTestSwapFlashLoan(b.GetTestContext(), b.testBackend)

	lastBlock := uint64(12)
	transactOpts := b.testBackend.GetTxContext(b.GetTestContext(), nil)

	// Initialize testing config.
	contractConfigBridge := indexerConfig.ContractConfig{
		ContractType: "bridge",
		Address:      bridgeContract.Address().String(),
		StartBlock:   0,
	}
	contractConfigBridgeV1 := indexerConfig.ContractConfig{
		ContractType: "bridge",
		Address:      bridgeV1Contract.Address().String(),
		StartBlock:   0,
	}
	contractConfigSwap1 := indexerConfig.ContractConfig{
		ContractType: "swap",
		Address:      swapContractA.Address().String(),
		StartBlock:   0,
	}
	contractConfigSwap2 := indexerConfig.ContractConfig{
		ContractType: "swap",
		Address:      swapContractB.Address().String(),
		StartBlock:   0,
	}
	contractConfigMetaSwap := indexerConfig.ContractConfig{
		ContractType: "metaswap",
		Address:      metaSwapContract.Address().String(),
		StartBlock:   0,
	}
	contractMessageBus := indexerConfig.ContractConfig{
		ContractType: "messagebus",
		Address:      messageBusContract.Address().String(),
		StartBlock:   0,
	}

	// CCTP config
	contractCCTP := indexerConfig.ContractConfig{
		ContractType: "cctp",
		Address:      cctpContract.Address().String(),
		StartBlock:   0,
	}

	// Create the chain configs
	chainConfigs := []indexerConfig.ChainConfig{
		{
			ChainID:             uint32(testChainID.Uint64()),
			RPCURL:              gofakeit.URL(),
			FetchBlockIncrement: 2,
			MaxGoroutines:       2,
			Contracts:           []indexerConfig.ContractConfig{contractConfigBridge, contractConfigSwap1, contractConfigSwap2, contractMessageBus, contractConfigMetaSwap, contractCCTP},
		},
	}
	chainConfigsV1 := []indexerConfig.ChainConfig{
		{
			ChainID:             uint32(testChainID.Uint64()),
			RPCURL:              gofakeit.URL(),
			FetchBlockIncrement: 2,
			MaxGoroutines:       2,
			Contracts:           []indexerConfig.ContractConfig{contractConfigBridgeV1, contractConfigSwap1, contractConfigSwap2, contractMessageBus, contractConfigMetaSwap},
		},
	}

	// Store blocktimes for testing defillama and timestamp indexing.
	for i := uint64(0); i < 13; i++ {
		err := b.eventDB.StoreBlockTime(b.GetTestContext(), uint32(testChainID.Uint64()), i, i)
		Nil(b.T(), err)
	}

	// Store every bridge event.
	bridgeTx, err := bridgeRef.TestDeposit(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(2)), 1)
	depositLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 2)
	Nil(b.T(), err)

	bridgeTx, err = bridgeRef.TestRedeem(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(int64(gofakeit.Uint32()))), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(2)), 2)
	redeemLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 2)
	Nil(b.T(), err)

	bridgeTx, err = bridgeRef.TestWithdraw(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(2)), 3)
	withdrawLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 2)
	Nil(b.T(), err)

	bridgeTx, err = bridgeRef.TestMint(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(3)), 1)
	mintLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 3)
	Nil(b.T(), err)

	bridgeTx, err = bridgeRef.TestDepositAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(6)), 1)
	depositAndSwapLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 6)
	Nil(b.T(), err)

	bridgeTx, err = bridgeRef.TestRedeemAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(7)), 1)
	redeemAndSwapLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 7)
	Nil(b.T(), err)

	bridgeTx, err = bridgeRef.TestRedeemAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(8)), 1)
	redeemAndRemoveLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 8)
	Nil(b.T(), err)

	bridgeTx, err = bridgeRef.TestMintAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(9)), 1)
	mintAndSwapLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 9)
	Nil(b.T(), err)

	bridgeTx, err = bridgeRef.TestWithdrawAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(10)), 1)
	withdrawAndRemoveLog, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 10)
	Nil(b.T(), err)

	bridgeTx, err = bridgeRef.TestRedeemV2(transactOpts.TransactOpts, [32]byte{byte(gofakeit.Uint64())}, big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(11)), 1)
	redeemV2Log, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 12)
	Nil(b.T(), err)

	// Store every bridge event using the V1 contract.
	bridgeTx, err = bridgeV1Ref.TestDeposit(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(3)), 1)
	depositV1Log, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 3)
	Nil(b.T(), err)

	bridgeTx, err = bridgeV1Ref.TestRedeem(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(int64(gofakeit.Uint32()))), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(3)), 2)
	redeemV1Log, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 3)
	Nil(b.T(), err)

	bridgeTx, err = bridgeV1Ref.TestWithdraw(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(3)), 3)
	withdrawV1Log, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 3)
	Nil(b.T(), err)

	bridgeTx, err = bridgeV1Ref.TestDepositAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(7)), 1)
	depositAndSwapV1Log, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 7)
	Nil(b.T(), err)

	bridgeTx, err = bridgeV1Ref.TestRedeemAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(8)), 1)
	redeemAndSwapV1Log, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 8)
	Nil(b.T(), err)

	bridgeTx, err = bridgeV1Ref.TestRedeemAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(b.T(), err)
	b.storeEthTx(bridgeTx, testChainID, big.NewInt(int64(9)), 1)
	redeemAndRemoveV1Log, err := b.storeTestLog(bridgeTx, uint32(testChainID.Uint64()), 9)
	Nil(b.T(), err)

	// Store every swap event across two different swap contracts.
	swapTx, err := swapRefA.TestSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	b.storeEthTx(swapTx, testChainID, big.NewInt(int64(4)), 1)
	swapLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 4)
	Nil(b.T(), err)

	swapTx, err = swapRefB.TestAddLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	b.storeEthTx(swapTx, testChainID, big.NewInt(int64(4)), 2)
	addLiquidityLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 4)
	Nil(b.T(), err)

	swapTx, err = swapRefB.TestRemoveLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	b.storeEthTx(swapTx, testChainID, big.NewInt(int64(4)), 3)
	removeLiquidityLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 4)
	Nil(b.T(), err)

	swapTx, err = swapRefA.TestRemoveLiquidityOne(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	b.storeEthTx(swapTx, testChainID, big.NewInt(int64(5)), 1)
	removeLiquidityOneLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 5)
	Nil(b.T(), err)

	swapTx, err = swapRefA.TestRemoveLiquidityImbalance(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	b.storeEthTx(swapTx, testChainID, big.NewInt(int64(5)), 2)
	removeLiquidityImbalanceLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 5)
	Nil(b.T(), err)

	swapTx, err = swapRefB.TestNewAdminFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	b.storeEthTx(swapTx, testChainID, big.NewInt(int64(6)), 1)
	newAdminFeeLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 6)
	Nil(b.T(), err)

	swapTx, err = swapRefA.TestNewSwapFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	b.storeEthTx(swapTx, testChainID, big.NewInt(int64(7)), 1)
	newSwapFeeLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 7)
	Nil(b.T(), err)

	swapTx, err = swapRefA.TestRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	b.storeEthTx(swapTx, testChainID, big.NewInt(int64(8)), 1)
	rampALog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 8)
	Nil(b.T(), err)

	swapTx, err = swapRefB.TestStopRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	b.storeEthTx(swapTx, testChainID, big.NewInt(int64(8)), 1)
	stopRampALog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 8)
	Nil(b.T(), err)

	swapTx, err = swapRefA.TestFlashLoan(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	b.storeEthTx(swapTx, testChainID, big.NewInt(int64(12)), 1)
	flashLoanLog, err := b.storeTestLog(swapTx, uint32(testChainID.Uint64()), 12)
	Nil(b.T(), err)

	// Store metaswap swap underlying event.
	metaSwapTx, err := metaSwapRef.TestSwapUnderlying(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(b.T(), err)
	b.storeEthTx(metaSwapTx, testChainID, big.NewInt(int64(4)), 3)
	metaSwapUnderlyingLog, err := b.storeTestLog(metaSwapTx, uint32(testChainID.Uint64()), 4)
	Nil(b.T(), err)

	// Store every message event.
	messageTx, err := messageBusRef.TestExecuted(transactOpts.TransactOpts, [32]byte{byte(gofakeit.Uint64())}, uint8(gofakeit.Number(0, 2)), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint64(), gofakeit.Uint64())
	Nil(b.T(), err)
	executedLog, err := b.storeTestLog(messageTx, uint32(testChainID.Uint64()), 5)
	Nil(b.T(), err)

	messageTx, err = messageBusRef.TestMessageSent(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())}, big.NewInt(int64(gofakeit.Uint32())), []byte(gofakeit.Paragraph(2, 5, 30, " ")), gofakeit.Uint64(), []byte{byte(gofakeit.Uint64())}, big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(b.T(), err)
	messageSentLog, err := b.storeTestLog(messageTx, uint32(testChainID.Uint64()), 6)
	Nil(b.T(), err)

	messageTx, err = messageBusRef.TestCallReverted(transactOpts.TransactOpts, gofakeit.Paragraph(2, 4, 20, " "))
	Nil(b.T(), err)
	callRevertedLog, err := b.storeTestLog(messageTx, uint32(testChainID.Uint64()), 7)
	Nil(b.T(), err)

	// Store every cctp event.
	var requestID [32]byte
	requestIDBytes := common.Hex2Bytes(mocks.NewMockHash(b.T()).String())
	copy(requestID[:], requestIDBytes)

	requestSentTx, err := cctpRef.TestSendCircleToken(transactOpts.TransactOpts, testChainID, common.BigToAddress(big.NewInt(gofakeit.Int64())), 1, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(gofakeit.Int64()), 1, []byte(gofakeit.Paragraph(2, 5, 30, " ")), requestID)
	Nil(b.T(), err)
	b.storeEthTx(requestSentTx, testChainID, big.NewInt(int64(5)), 5)
	requestSentLog, err := b.storeTestLog(requestSentTx, uint32(testChainID.Uint64()), 5)
	Nil(b.T(), err)

	requestFulfilledTx, err := cctpRef.TestReceiveCircleToken(transactOpts.TransactOpts, uint32(testChainID.Int64()), common.BigToAddress(big.NewInt(gofakeit.Int64())), common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(gofakeit.Int64()), common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(gofakeit.Int64()), requestID)
	Nil(b.T(), err)
	b.storeEthTx(requestFulfilledTx, testChainID, big.NewInt(int64(6)), 6)
	requestFulfilledLog, err := b.storeTestLog(requestFulfilledTx, uint32(testChainID.Uint64()), 6)
	Nil(b.T(), err)

	// Go through each contract and save the end height in scribe
	for i := range chainConfigs[0].Contracts {
		//  the last block store per contract
		err = b.eventDB.StoreLastIndexed(b.GetTestContext(), common.HexToAddress(chainConfigs[0].Contracts[i].Address), uint32(testChainID.Uint64()), lastBlock, scribeTypes.IndexingConfirmed)
		Nil(b.T(), err)
	}
	for i := range chainConfigsV1[0].Contracts {
		//  the last block store per contract
		err = b.eventDB.StoreLastIndexed(b.GetTestContext(), common.HexToAddress(chainConfigsV1[0].Contracts[i].Address), uint32(testChainID.Uint64()), lastBlock, scribeTypes.IndexingConfirmed)
		Nil(b.T(), err)
	}

	// Set up a ChainBackfiller
	bcf, err := fetcher.NewBridgeConfigFetcher(b.bridgeConfigContract.Address(), b.bridgeConfigContract)
	Nil(b.T(), err)

	tokenSymbolToIDs, err := parser.ParseYaml(static.GetTokenSymbolToTokenIDConfig())
	Nil(b.T(), err)
	tokenDataService, err := tokendata.NewTokenDataService(bcf, tokenSymbolToIDs)
	Nil(b.T(), err)
	tokenPriceService, err := tokenprice.NewPriceDataService()
	Nil(b.T(), err)

	bp, err := parser.NewBridgeParser(b.db, bridgeContract.Address(), tokenDataService, b.consumerFetcher, tokenPriceService, false)
	Nil(b.T(), err)
	bpv1, err := parser.NewBridgeParser(b.db, bridgeV1Contract.Address(), tokenDataService, b.consumerFetcher, tokenPriceService, false)
	Nil(b.T(), err)

	// srB is the swap ref for getting token data
	srA, err := fetcher.NewSwapFetcher(swapContractA.Address(), b.testBackend, false)
	Nil(b.T(), err)
	spA, err := parser.NewSwapParser(b.db, swapContractA.Address(), false, b.consumerFetcher, srA, tokenDataService, tokenPriceService)
	Nil(b.T(), err)

	// srB is the swap ref for getting token data
	srB, err := fetcher.NewSwapFetcher(swapContractB.Address(), b.testBackend, false)
	Nil(b.T(), err)
	spB, err := parser.NewSwapParser(b.db, swapContractB.Address(), false, b.consumerFetcher, srB, tokenDataService, tokenPriceService)
	Nil(b.T(), err)

	// msr is the meta swap ref for getting token data
	msr, err := fetcher.NewSwapFetcher(metaSwapContract.Address(), b.testBackend, true)
	Nil(b.T(), err)
	msp, err := parser.NewSwapParser(b.db, metaSwapContract.Address(), true, b.consumerFetcher, msr, tokenDataService, tokenPriceService)
	Nil(b.T(), err)

	// msr is the meta swap ref for getting token data
	cr, err := fetcher.NewCCTPFetcher(cctpRef.Address(), b.testBackend)
	Nil(b.T(), err)
	cp, err := parser.NewCCTPParser(b.db, cctpRef.Address(), b.consumerFetcher, cr, tokenDataService, tokenPriceService, false)
	Nil(b.T(), err)

	spMap := map[common.Address]*parser.SwapParser{}
	spMap[swapContractA.Address()] = spA
	spMap[swapContractB.Address()] = spB
	spMap[metaSwapContract.Address()] = msp
	f := fetcher.NewFetcher(b.gqlClient, b.metrics)

	// Set up message bus parser
	mbp, err := parser.NewMessageBusParser(b.db, messageBusContract.Address(), b.consumerFetcher, tokenPriceService)
	Nil(b.T(), err)

	// Test the first chain in the config file
	chainBackfiller := backfill.NewChainBackfiller(b.db, bp, spMap, mbp, cp, f, chainConfigs[0])
	chainBackfillerV1 := backfill.NewChainBackfiller(b.db, bpv1, spMap, mbp, cp, f, chainConfigsV1[0])

	// Backfill the blocks
	var count int64
	err = chainBackfiller.Backfill(b.GetTestContext(), false, 1)

	Nil(b.T(), err)
	swapEvents := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Find(&sql.SwapEvent{}).Count(&count)
	Nil(b.T(), swapEvents.Error)
	Equal(b.T(), int64(11), count)

	bridgeEvents := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Find(&sql.BridgeEvent{}).Count(&count)
	Nil(b.T(), bridgeEvents.Error)
	Equal(b.T(), int64(12), count) // 10 + 2 cctp events

	messageEvents := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Find(&sql.MessageBusEvent{}).Count(&count)
	Nil(b.T(), messageEvents.Error)
	Equal(b.T(), int64(3), count)

	cctpEvents := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Find(&sql.CCTPEvent{}).Count(&count)
	Nil(b.T(), cctpEvents.Error)
	Equal(b.T(), int64(2), count)

	// Test cctp parity
	err = b.sendCircleTokenParity(requestSentLog, cp)
	Nil(b.T(), err)
	err = b.receiveCircleTokenParity(requestFulfilledLog, cp)
	Nil(b.T(), err)

	// Test bridge parity
	err = b.depositParity(depositLog, bp, uint32(testChainID.Uint64()), false)
	Nil(b.T(), err)
	err = b.redeemParity(redeemLog, bp, uint32(testChainID.Uint64()), false)
	Nil(b.T(), err)
	err = b.withdrawParity(withdrawLog, bp, uint32(testChainID.Uint64()), false)
	Nil(b.T(), err)
	err = b.mintParity(mintLog, bp, uint32(testChainID.Uint64()), false)
	Nil(b.T(), err)
	err = b.depositAndSwapParity(depositAndSwapLog, bp, uint32(testChainID.Uint64()), false)
	Nil(b.T(), err)
	err = b.redeemAndSwapParity(redeemAndSwapLog, bp, uint32(testChainID.Uint64()), false)
	Nil(b.T(), err)
	err = b.redeemAndRemoveParity(redeemAndRemoveLog, bp, uint32(testChainID.Uint64()), false)
	Nil(b.T(), err)
	err = b.mintAndSwapParity(mintAndSwapLog, bp, uint32(testChainID.Uint64()), false)
	Nil(b.T(), err)
	err = b.withdrawAndRemoveParity(withdrawAndRemoveLog, bp, uint32(testChainID.Uint64()), false)
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

	// Test meta swap parity
	err = b.metaSwapUnderlyingParity(metaSwapUnderlyingLog, msp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)

	// Test message parity
	err = b.executedParity(executedLog, mbp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.messageSentParity(messageSentLog, mbp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)
	err = b.callRevertedParity(callRevertedLog, mbp, uint32(testChainID.Uint64()))
	Nil(b.T(), err)

	// Test bridge v1 parity
	err = chainBackfillerV1.Backfill(b.GetTestContext(), false, 1)
	Nil(b.T(), err)

	err = b.depositParity(depositV1Log, bpv1, uint32(testChainID.Uint64()), true)
	Nil(b.T(), err)
	err = b.redeemParity(redeemV1Log, bpv1, uint32(testChainID.Uint64()), true)
	Nil(b.T(), err)
	err = b.withdrawParity(withdrawV1Log, bpv1, uint32(testChainID.Uint64()), true)
	Nil(b.T(), err)
	err = b.depositAndSwapParity(depositAndSwapV1Log, bpv1, uint32(testChainID.Uint64()), true)
	Nil(b.T(), err)
	err = b.redeemAndSwapParity(redeemAndSwapV1Log, bpv1, uint32(testChainID.Uint64()), true)
	Nil(b.T(), err)
	err = b.redeemAndRemoveParity(redeemAndRemoveV1Log, bpv1, uint32(testChainID.Uint64()), true)
	Nil(b.T(), err)

	bridgeEvents = b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Find(&sql.BridgeEvent{}).Count(&count)

	Nil(b.T(), bridgeEvents.Error)
	Equal(b.T(), int64(18), count) // 16 + 2 cctp events

	lastBlockStored, err := b.db.GetLastStoredBlock(b.GetTestContext(), uint32(testChainID.Uint64()), chainConfigsV1[0].Contracts[0].Address)

	Nil(b.T(), err)
	Equal(b.T(), lastBlock, lastBlockStored)
}

// storeTestLogs stores the test logs in the database.
func (b *BackfillSuite) storeTestLog(tx *types.Transaction, chainID uint32, blockNumber uint64) (*types.Log, error) {
	b.testBackend.WaitForConfirmation(b.GetTestContext(), tx)
	receipt, err := b.testBackend.TransactionReceipt(b.GetTestContext(), tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get receipt for transaction %s: %w", tx.Hash().String(), err)
	}
	receipt.Logs[0].BlockNumber = blockNumber
	err = b.eventDB.StoreLogs(b.GetTestContext(), chainID, *receipt.Logs[0])
	if err != nil {
		return nil, fmt.Errorf("error storing swap log: %w", err)
	}
	return receipt.Logs[0], nil
}

//nolint:dupl
func (b *BackfillSuite) sendCircleTokenParity(log *types.Log, parser *parser.CCTPParser) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseCircleRequestSent(*log)
	_ = parsedLog
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	var count int64
	sender := gosql.NullString{
		String: parsedLog.Sender.String(),
		Valid:  true,
	}
	nonce := gosql.NullInt64{
		Int64: int64(parsedLog.Nonce),
		Valid: true,
	}

	requestVersion := gosql.NullInt32{
		Int32: int32(parsedLog.RequestVersion),
		Valid: true,
	}
	formattedRequest := gosql.NullString{
		String: common.Bytes2Hex(parsedLog.FormattedRequest),
		Valid:  true,
	}
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.CCTPEvent{}).
		Where(&sql.CCTPEvent{
			ContractAddress:    log.Address.String(),
			BlockNumber:        log.BlockNumber,
			TxHash:             log.TxHash.String(),
			EventType:          cctpTypes.CircleRequestSentEvent.Int(),
			RequestID:          common.Bytes2Hex(parsedLog.RequestID[:]),
			DestinationChainID: parsedLog.ChainId,
			Sender:             sender,
			Nonce:              nonce,
			Token:              parsedLog.Token.String(),
			Amount:             parsedLog.Amount,
			RequestVersion:     requestVersion,
			FormattedRequest:   formattedRequest,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) receiveCircleTokenParity(log *types.Log, parser *parser.CCTPParser) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseCircleRequestFulfilled(*log)
	_ = parsedLog
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	var count int64
	mintToken := gosql.NullString{
		String: parsedLog.MintToken.String(),
		Valid:  true,
	}
	recipient := gosql.NullString{
		String: parsedLog.Recipient.String(),
		Valid:  true,
	}
	domainToChain := []int64{1, 43114, 10, 42161}
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.CCTPEvent{}).
		Where(&sql.CCTPEvent{
			ContractAddress: log.Address.String(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			EventType:       cctpTypes.CircleRequestFulfilledEvent.Int(),
			RequestID:       common.Bytes2Hex(parsedLog.RequestID[:]),
			OriginChainID:   big.NewInt(domainToChain[parsedLog.OriginDomain]),
			MintToken:       mintToken,
			Amount:          parsedLog.Amount,
			Recipient:       recipient,
			Fee:             parsedLog.Fee,
			Token:           parsedLog.Token.String(),
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) depositParity(log *types.Log, parser *parser.BridgeParser, chainID uint32, useV1 bool) error {
	// parse the log
	if useV1 {
		parsedLog, err := parser.FiltererV1.ParseTokenDeposit(*log)
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
func (b *BackfillSuite) redeemParity(log *types.Log, parser *parser.BridgeParser, chainID uint32, useV1 bool) error {
	// parse the log
	if useV1 {
		parsedLog, err := parser.FiltererV1.ParseTokenRedeem(*log)
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
func (b *BackfillSuite) withdrawParity(log *types.Log, parser *parser.BridgeParser, chainID uint32, useV1 bool) error {
	// parse the log
	if useV1 {
		parsedLog, err := parser.FiltererV1.ParseTokenWithdraw(*log)
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
func (b *BackfillSuite) mintParity(log *types.Log, parser *parser.BridgeParser, chainID uint32, useV1 bool) error {
	// parse the log
	if useV1 {
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
func (b *BackfillSuite) depositAndSwapParity(log *types.Log, parser *parser.BridgeParser, chainID uint32, useV1 bool) error {
	// parse the log
	if useV1 {
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
func (b *BackfillSuite) redeemAndSwapParity(log *types.Log, parser *parser.BridgeParser, chainID uint32, useV1 bool) error {
	// parse the log
	if useV1 {
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
func (b *BackfillSuite) redeemAndRemoveParity(log *types.Log, parser *parser.BridgeParser, chainID uint32, useV1 bool) error {
	// parse the log
	if useV1 {
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
func (b *BackfillSuite) mintAndSwapParity(log *types.Log, parser *parser.BridgeParser, chainID uint32, useV1 bool) error {
	// parse the log
	if useV1 {
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
				SwapSuccess:    big.NewInt(int64(*parserpkg.BoolToUint8(&parsedLog.SwapSuccess))),
				Kappa:          kappa,
			}).Count(&count)
		if events.Error != nil {
			return fmt.Errorf("error querying for event: %w", events.Error)
		}
		Equal(b.T(), int64(1), count)
		return nil
	}
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
			SwapSuccess:    big.NewInt(int64(*parserpkg.BoolToUint8(&parsedLog.SwapSuccess))),
			Kappa:          kappa,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) withdrawAndRemoveParity(log *types.Log, parser *parser.BridgeParser, chainID uint32, useV1 bool) error {
	// parse the log
	if useV1 {
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
				SwapSuccess:    big.NewInt(int64(*parserpkg.BoolToUint8(&parsedLog.SwapSuccess))),
				Kappa:          kappa,
			}).Count(&count)
		if events.Error != nil {
			return fmt.Errorf("error querying for event: %w", events.Error)
		}
		Equal(b.T(), int64(1), count)
		return nil
	}
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
			SwapSuccess:    big.NewInt(int64(*parserpkg.BoolToUint8(&parsedLog.SwapSuccess))),
			Kappa:          kappa,
		}).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) redeemV2Parity(log *types.Log, parser *parser.BridgeParser, chainID uint32) error {
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
func (b *BackfillSuite) swapParity(log *types.Log, parser *parser.SwapParser, chainID uint32) error {
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
func (b *BackfillSuite) addLiquidityParity(log *types.Log, parser *parser.SwapParser, chainID uint32) error {
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
	Equal(b.T(), arrayToTokenIndexMap(parsedLog.Fees), storedLog.Fee)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) removeLiquidityParity(log *types.Log, parser *parser.SwapParser, chainID uint32) error {
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
func (b *BackfillSuite) removeLiquidityOneParity(log *types.Log, parser *parser.SwapParser, chainID uint32) error {
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
func (b *BackfillSuite) removeLiquidityImbalanceParity(log *types.Log, parser *parser.SwapParser, chainID uint32) error {
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
	Equal(b.T(), arrayToTokenIndexMap(parsedLog.Fees), storedLog.Fee)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) newAdminFeeParity(log *types.Log, parser *parser.SwapParser, chainID uint32) error {
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
func (b *BackfillSuite) newSwapFeeParity(log *types.Log, parser *parser.SwapParser, chainID uint32) error {
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
func (b *BackfillSuite) rampAParity(log *types.Log, parser *parser.SwapParser, chainID uint32) error {
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
func (b *BackfillSuite) stopRampAParity(log *types.Log, parser *parser.SwapParser, chainID uint32) error {
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
func (b *BackfillSuite) flashLoanParity(log *types.Log, parser *parser.SwapParser, chainID uint32) error {
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
	Equal(b.T(), feeArray, storedLog.Fee)
	return nil
}

// storeEthTx stores the eth transaction so the get sender functionality can be tested.
func (b *BackfillSuite) storeEthTx(tx *types.Transaction, chainID *big.Int, blockNumber *big.Int, index int) {
	err := b.eventDB.StoreEthTx(b.GetTestContext(), tx, uint32(chainID.Uint64()), common.BigToHash(blockNumber), blockNumber.Uint64(), uint64(index))
	Nil(b.T(), err)
}

//nolint:dupl
func (b *BackfillSuite) executedParity(log *types.Log, parser *parser.MessageBusParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseExecuted(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	var storedLog sql.MessageBusEvent
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.MessageBusEvent{}).
		Where(&sql.MessageBusEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       messageBusTypes.ExecutedEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),

			MessageID:     gosql.NullString{String: common.Bytes2Hex(parsedLog.MessageId[:]), Valid: true},
			SourceChainID: big.NewInt(int64(parsedLog.SrcChainId)),
		}).
		Find(&storedLog).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	Equal(b.T(), big.NewInt(int64(parsedLog.SrcNonce)), storedLog.Nonce)
	Equal(b.T(), parsedLog.DstAddress.String(), storedLog.DestinationAddress.String)
	return nil
}

//nolint:dupl
func (b *BackfillSuite) metaSwapUnderlyingParity(log *types.Log, parser *parser.SwapParser, chainID uint32) error {
	// parse the log

	parsedLog, err := parser.FiltererMetaSwap.ParseTokenSwapUnderlying(*log)
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
func (b *BackfillSuite) messageSentParity(log *types.Log, parser *parser.MessageBusParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseMessageSent(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	var storedLog sql.MessageBusEvent
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.MessageBusEvent{}).
		Where(&sql.MessageBusEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       messageBusTypes.MessageSentEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),

			MessageID:     gosql.NullString{String: common.Bytes2Hex(parsedLog.MessageId[:]), Valid: true},
			SourceChainID: parsedLog.SrcChainID,
		}).
		Find(&storedLog).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	Equal(b.T(), common.Bytes2Hex(parsedLog.Message), storedLog.Message.String)
	Equal(b.T(), parsedLog.Fee, storedLog.Fee)
	Equal(b.T(), common.Bytes2Hex(parsedLog.Receiver[:]), storedLog.Receiver.String)

	return nil
}

//nolint:dupl
func (b *BackfillSuite) callRevertedParity(log *types.Log, parser *parser.MessageBusParser, chainID uint32) error {
	// parse the log
	parsedLog, err := parser.Filterer.ParseCallReverted(*log)
	if err != nil {
		return fmt.Errorf("error parsing log: %w", err)
	}
	var storedLog sql.MessageBusEvent
	var count int64
	events := b.db.UNSAFE_DB().WithContext(b.GetTestContext()).Model(&sql.MessageBusEvent{}).
		Where(&sql.MessageBusEvent{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			EventType:       messageBusTypes.CallRevertedEvent.Int(),
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
		}).
		Find(&storedLog).Count(&count)
	if events.Error != nil {
		return fmt.Errorf("error querying for event: %w", events.Error)
	}
	Equal(b.T(), int64(1), count)
	Equal(b.T(), parsedLog.Reason, storedLog.RevertedReason.String)
	return nil
}
