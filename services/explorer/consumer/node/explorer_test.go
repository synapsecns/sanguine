package node_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"

	"github.com/synapsecns/sanguine/services/explorer/consumer/node"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap/testswap"
	"github.com/synapsecns/sanguine/services/explorer/testutil/testcontracts"
	"math/big"
)

// TestLive tests live recording of events.
func (n NodeSuite) TestLive() {
	fmt.Println(n.testBackends, "testChainID")

	chainConfigs := []config.ChainConfig{}
	backends := make(map[uint32]bind.ContractBackend)
	// ethclient.DialContext(ctx, chainConfig.RPCURL)
	for k := range n.testBackends {
		testChainID := k
		backends[k] = n.testBackends[k]
		fmt.Println(testChainID, "testChainID")
		bridgeContract, bridgeRef := n.testDeployManager.GetTestSynapseBridge(n.GetTestContext(), n.testBackends[k])
		swapContractA, swapRefA := n.testDeployManager.GetTestSwapFlashLoan(n.GetTestContext(), n.testBackends[k])
		testDeployManagerB := testcontracts.NewDeployManager(n.T())
		swapContractB, swapRefB := testDeployManagerB.GetTestSwapFlashLoan(n.GetTestContext(), n.testBackends[k])
		transactOpts := n.testBackends[k].GetTxContext(n.GetTestContext(), nil)

		chainConfigs = append(chainConfigs, config.ChainConfig{
			ChainID:                testChainID,
			FetchBlockIncrement:    3,
			StartBlock:             0,
			BridgeConfigV3Address:  n.bridgeConfigContract.Address().String(),
			SynapseBridgeAddress:   bridgeContract.Address().String(),
			SwapFlashLoanAddresses: []string{swapContractA.Address().String(), swapContractB.Address().String()},
		})
		n.fillBlocks(bridgeRef, swapRefA, swapRefB, transactOpts, testChainID)
	}

	// This structure is for reference
	explorerConfig := config.Config{
		Chains:      chainConfigs,
		RefreshRate: 2,
		ScribeURL:   n.gqlClient.Client.BaseURL,
	}
	explorerBackfiller, err := node.NewExplorerBackfiller(n.GetTestContext(), n.db, explorerConfig, backends)
	n.Nil(err)
	n.NotNil(explorerBackfiller)
	err = explorerBackfiller.Backfill(n.GetTestContext())
	if err != nil {
		n.FailNow(err.Error())
	}
	var count int64
	bridgeEvents := n.db.UNSAFE_DB().WithContext(n.GetTestContext()).Find(&sql.BridgeEvent{}).Count(&count)
	Nil(n.T(), bridgeEvents.Error)
	fmt.Println("COUNT", count)
	swapEvents := n.db.UNSAFE_DB().WithContext(n.GetTestContext()).Find(&sql.SwapEvent{}).Count(&count)
	Nil(n.T(), swapEvents.Error)
	fmt.Println("COUNT", count)

}

// nolinting until parity tests implemented
//
//nolint:unparam
func (n *NodeSuite) storeTestLog(tx *types.Transaction, chainID uint32, blockNumber uint64) (*types.Log, error) {
	n.testBackends[chainID].WaitForConfirmation(n.GetTestContext(), tx)
	receipt, err := n.testBackends[chainID].TransactionReceipt(n.GetTestContext(), tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get receipt for transaction %s: %w", tx.Hash().String(), err)
	}
	receipt.Logs[0].BlockNumber = blockNumber
	err = n.eventDB.StoreLog(n.GetTestContext(), *receipt.Logs[0], chainID)
	if err != nil {
		return nil, fmt.Errorf("error storing swap log: %w", err)
	}
	return receipt.Logs[0], nil
}

func (n NodeSuite) fillBlocks(bridgeRef *testbridge.TestBridgeRef, swapRefA *testswap.TestSwapRef, swapRefB *testswap.TestSwapRef, transactOpts backends.AuthType, chainID uint32) {
	bridgeTx, err := bridgeRef.TestDeposit(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(bridgeTx, chainID, 2)
	Nil(n.T(), err)
	bridgeTx, err = bridgeRef.TestRedeem(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(int64(gofakeit.Uint32()))), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(bridgeTx, chainID, 2)
	Nil(n.T(), err)
	bridgeTx, err = bridgeRef.TestWithdraw(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(n.T(), err)
	_, err = n.storeTestLog(bridgeTx, chainID, 2)
	Nil(n.T(), err)
	bridgeTx, err = bridgeRef.TestMint(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(n.T(), err)
	_, err = n.storeTestLog(bridgeTx, chainID, 3)
	Nil(n.T(), err)
	bridgeTx, err = bridgeRef.TestDepositAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(bridgeTx, chainID, 6)
	Nil(n.T(), err)
	bridgeTx, err = bridgeRef.TestRedeemAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(bridgeTx, chainID, 7)
	Nil(n.T(), err)
	bridgeTx, err = bridgeRef.TestRedeemAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(bridgeTx, chainID, 8)
	Nil(n.T(), err)
	bridgeTx, err = bridgeRef.TestMintAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(n.T(), err)
	_, err = n.storeTestLog(bridgeTx, chainID, 9)
	Nil(n.T(), err)
	bridgeTx, err = bridgeRef.TestWithdrawAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(n.T(), err)
	_, err = n.storeTestLog(bridgeTx, chainID, 10)
	Nil(n.T(), err)
	bridgeTx, err = bridgeRef.TestRedeemV2(transactOpts.TransactOpts, [32]byte{byte(gofakeit.Uint64())}, big.NewInt(int64(gofakeit.Uint32())), common.HexToAddress(testTokens[0].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(bridgeTx, chainID, 11)
	Nil(n.T(), err)

	// Store every swap event across two different swap contracts.
	swapTx, err := swapRefA.TestSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(swapTx, chainID, 0)
	Nil(n.T(), err)
	swapTx, err = swapRefB.TestAddLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(swapTx, chainID, 0)
	Nil(n.T(), err)
	swapTx, err = swapRefB.TestRemoveLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(swapTx, chainID, 0)
	Nil(n.T(), err)
	swapTx, err = swapRefA.TestRemoveLiquidityOne(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(swapTx, chainID, 1)
	Nil(n.T(), err)
	swapTx, err = swapRefA.TestRemoveLiquidityImbalance(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(swapTx, chainID, 1)
	Nil(n.T(), err)
	swapTx, err = swapRefB.TestNewAdminFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(swapTx, chainID, 5)
	Nil(n.T(), err)
	swapTx, err = swapRefA.TestNewSwapFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(swapTx, chainID, 6)
	Nil(n.T(), err)
	swapTx, err = swapRefA.TestRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(swapTx, chainID, 7)
	Nil(n.T(), err)
	swapTx, err = swapRefB.TestStopRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(swapTx, chainID, 8)
	Nil(n.T(), err)
	swapTx, err = swapRefA.TestFlashLoan(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	_, err = n.storeTestLog(swapTx, chainID, 9)
	Nil(n.T(), err)
}
