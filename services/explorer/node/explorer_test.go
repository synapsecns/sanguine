package node_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap/testswap"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/node"
	"github.com/synapsecns/sanguine/services/explorer/testutil/testcontracts"
	"math/big"
	"os"
)

var testTokens = make(map[uint32]TestToken)

// TestLive tests live recording of events.
func (n NodeSuite) TestLive() {
	if os.Getenv("CI") != "" {
		n.T().Skip("Network / processing test flake")
	}
	chainConfigs := []config.ChainConfig{}
	backends := make(map[uint32]bind.ContractBackend)
	// ethclient.DialContext(ctx, chainConfig.RPCURL)
	for k := range n.testBackends {
		testTokens[k] = TestToken{
			tokenID: gofakeit.FirstName(),
			BridgeConfigV3Token: bridgeconfig.BridgeConfigV3Token{
				ChainId:       big.NewInt(int64(k)),
				TokenAddress:  mocks.MockAddress().String(),
				TokenDecimals: gofakeit.Uint8(),
				MaxSwap:       new(big.Int).SetUint64(gofakeit.Uint64()),
				// TODO: this should probably be smaller than maxswap
				MinSwap:       new(big.Int).SetUint64(gofakeit.Uint64()),
				SwapFee:       new(big.Int).SetUint64(gofakeit.Uint64()),
				MaxSwapFee:    new(big.Int).SetUint64(gofakeit.Uint64()),
				MinSwapFee:    new(big.Int).SetUint64(gofakeit.Uint64()),
				HasUnderlying: gofakeit.Bool(),
				IsUnderlying:  gofakeit.Bool(),
			}}
	}

	var deployInfo contracts.DeployedContract
	deployInfo, bridgeConfigContract := n.deployManager.GetBridgeConfigV3(n.GetTestContext(), n.testBackends[n.blockConfigChainID])
	for _, token := range testTokens {
		auth := n.testBackends[n.blockConfigChainID].GetTxContext(n.GetTestContext(), deployInfo.OwnerPtr())
		tx, err := token.SetTokenConfig(bridgeConfigContract, auth)
		n.Require().NoError(err)
		n.testBackends[n.blockConfigChainID].WaitForConfirmation(n.GetTestContext(), tx)
	}
	for k := range n.testBackends {
		backends[k] = n.testBackends[k]
		bridgeContract, bridgeRef := n.testDeployManager.GetTestSynapseBridge(n.GetTestContext(), n.testBackends[k])
		swapContractA, swapRefA := n.testDeployManager.GetTestSwapFlashLoan(n.GetTestContext(), n.testBackends[k])
		testDeployManagerB := testcontracts.NewDeployManager(n.T())
		swapContractB, swapRefB := testDeployManagerB.GetTestSwapFlashLoan(n.GetTestContext(), n.testBackends[k])
		transactOpts := n.testBackends[k].GetTxContext(n.GetTestContext(), nil)

		contracts := []config.ContractConfig{
			{
				ContractType: "bridge",
				Address:      bridgeContract.Address().String(),
				StartBlock:   0,
			},

			{
				ContractType: "swap",
				Address:      swapContractA.Address().String(),
				StartBlock:   0,
			},
			{
				ContractType: "swap",
				Address:      swapContractB.Address().String(),
				StartBlock:   0,
			},
		}
		chainConfigs = append(chainConfigs, config.ChainConfig{
			ChainID:             k,
			RPCURL:              gofakeit.URL(),
			FetchBlockIncrement: 100,
			MaxGoroutines:       5,
			Contracts:           contracts,
		})
		// go through each contract and save the end height in scribe
		for i := range contracts {
			//  the last block store per contract
			err := n.eventDB.StoreLastIndexed(n.GetTestContext(), common.HexToAddress(contracts[i].Address), k, 12)
			Nil(n.T(), err)
		}
		n.fillBlocks(bridgeRef, swapRefA, swapRefB, transactOpts, k)
	}

	// This structure is for reference
	explorerConfig := config.Config{
		RefreshRate:         2,
		ScribeURL:           n.gqlClient.Client.BaseURL,
		BridgeConfigAddress: deployInfo.Address().String(),
		BridgeConfigChainID: n.blockConfigChainID,
		Chains:              chainConfigs,
	}

	explorerBackfiller, err := node.NewExplorerBackfiller(n.db, explorerConfig, backends)
	n.Nil(err)
	n.NotNil(explorerBackfiller)
	err = explorerBackfiller.Backfill(n.GetTestContext(), false)
	n.Nil(err)
	var counttemp int64
	dd := n.db.UNSAFE_DB().WithContext(n.GetTestContext()).Table("swap_events").Count(&counttemp)
	n.Nil(dd.Error)

	var count int64
	bridgeEvents := n.db.UNSAFE_DB().WithContext(n.GetTestContext()).Find(&sql.BridgeEvent{}).Count(&count)
	Nil(n.T(), bridgeEvents.Error)
	Equal(n.T(), int64(10*len(n.testBackends)), count)

	swapEvents := n.db.UNSAFE_DB().WithContext(n.GetTestContext()).Find(&sql.SwapEvent{}).Count(&count)
	Nil(n.T(), swapEvents.Error)
	Equal(n.T(), int64(10*len(n.testBackends)), count)

	for k := range n.testBackends {
		bridgeEventsChain := n.db.UNSAFE_DB().WithContext(n.GetTestContext()).Model(&sql.BridgeEvent{}).Where(&sql.BridgeEvent{ChainID: k}).Count(&count)
		Nil(n.T(), bridgeEventsChain.Error)
		Equal(n.T(), int64(10), count)

		swapEventsChain := n.db.UNSAFE_DB().WithContext(n.GetTestContext()).Model(&sql.SwapEvent{}).Where(&sql.SwapEvent{ChainID: k}).Count(&count)
		Nil(n.T(), swapEventsChain.Error)
		Equal(n.T(), int64(10), count)
	}
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

	err = n.eventDB.StoreLogs(n.GetTestContext(), chainID, *receipt.Logs[0])
	if err != nil {
		return nil, fmt.Errorf("error storing swap log: %w", err)
	}

	return receipt.Logs[0], nil
}

func (n NodeSuite) fillBlocks(bridgeRef *testbridge.TestBridgeRef, swapRefA *testswap.TestSwapRef, swapRefB *testswap.TestSwapRef, transactOpts backends.AuthType, chainID uint32) {
	// Store blocktimes for testing defillama and timestamp indexing.
	for i := uint64(0); i < 13; i++ {
		err := n.eventDB.StoreBlockTime(n.GetTestContext(), chainID, i, i)
		Nil(n.T(), err)
	}

	bridgeTx, err := bridgeRef.TestDeposit(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(chainID)), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(n.T(), err)
	n.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(5)), 1)
	_, err = n.storeTestLog(bridgeTx, chainID, 5)
	Nil(n.T(), err)

	bridgeTx, err = bridgeRef.TestRedeem(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(int64(gofakeit.Uint32()))), big.NewInt(int64(chainID)), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(n.T(), err)
	n.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(5)), 2)
	_, err = n.storeTestLog(bridgeTx, chainID, 5)
	Nil(n.T(), err)

	bridgeTx, err = bridgeRef.TestWithdraw(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(n.T(), err)
	n.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(6)), 3)
	_, err = n.storeTestLog(bridgeTx, chainID, 6)
	Nil(n.T(), err)

	bridgeTx, err = bridgeRef.TestMint(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(n.T(), err)
	n.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(6)), 4)
	_, err = n.storeTestLog(bridgeTx, chainID, 6)
	Nil(n.T(), err)

	bridgeTx, err = bridgeRef.TestDepositAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(chainID)), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(n.T(), err)
	n.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(6)), 5)
	_, err = n.storeTestLog(bridgeTx, chainID, 6)
	Nil(n.T(), err)

	bridgeTx, err = bridgeRef.TestRedeemAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(chainID)), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(n.T(), err)
	n.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(7)), 1)
	_, err = n.storeTestLog(bridgeTx, chainID, 7)
	Nil(n.T(), err)

	bridgeTx, err = bridgeRef.TestRedeemAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(chainID)), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(n.T(), err)
	n.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(8)), 1)
	_, err = n.storeTestLog(bridgeTx, chainID, 8)
	Nil(n.T(), err)

	bridgeTx, err = bridgeRef.TestMintAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(n.T(), err)
	n.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(9)), 1)
	_, err = n.storeTestLog(bridgeTx, chainID, 9)
	Nil(n.T(), err)

	bridgeTx, err = bridgeRef.TestWithdrawAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(n.T(), err)
	n.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(10)), 1)
	_, err = n.storeTestLog(bridgeTx, chainID, 10)
	Nil(n.T(), err)

	bridgeTx, err = bridgeRef.TestRedeemV2(transactOpts.TransactOpts, [32]byte{byte(gofakeit.Uint64())}, big.NewInt(int64(chainID)), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(n.T(), err)
	n.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(12)), 1)
	_, err = n.storeTestLog(bridgeTx, chainID, 12)
	Nil(n.T(), err)

	// Store every swap event across two different swap contracts.
	swapTx, err := swapRefA.TestSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	n.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(5)), 1)
	_, err = n.storeTestLog(swapTx, chainID, 5)
	Nil(n.T(), err)

	swapTx, err = swapRefB.TestAddLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	n.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(5)), 2)
	_, err = n.storeTestLog(swapTx, chainID, 5)
	Nil(n.T(), err)

	swapTx, err = swapRefB.TestRemoveLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	n.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(6)), 1)
	_, err = n.storeTestLog(swapTx, chainID, 6)
	Nil(n.T(), err)

	swapTx, err = swapRefA.TestRemoveLiquidityOne(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	n.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(7)), 1)
	_, err = n.storeTestLog(swapTx, chainID, 7)
	Nil(n.T(), err)

	swapTx, err = swapRefA.TestRemoveLiquidityImbalance(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	n.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(8)), 1)
	_, err = n.storeTestLog(swapTx, chainID, 8)
	Nil(n.T(), err)

	swapTx, err = swapRefB.TestNewAdminFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	n.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(8)), 2)
	_, err = n.storeTestLog(swapTx, chainID, 8)
	Nil(n.T(), err)

	swapTx, err = swapRefA.TestNewSwapFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	n.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(8)), 3)
	_, err = n.storeTestLog(swapTx, chainID, 8)
	Nil(n.T(), err)

	swapTx, err = swapRefA.TestRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	n.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(9)), 1)
	_, err = n.storeTestLog(swapTx, chainID, 9)
	Nil(n.T(), err)

	swapTx, err = swapRefB.TestStopRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	n.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(9)), 2)
	_, err = n.storeTestLog(swapTx, chainID, 9)
	Nil(n.T(), err)

	swapTx, err = swapRefA.TestFlashLoan(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(n.T(), err)
	n.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(9)), 3)
	_, err = n.storeTestLog(swapTx, chainID, 9)
	Nil(n.T(), err)
}

// storeEthTx stores the eth transaction so the get sender functionality can be tested.
func (n *NodeSuite) storeEthTx(tx *types.Transaction, chainID *big.Int, blockNumber *big.Int, index int) {
	err := n.eventDB.StoreEthTx(n.GetTestContext(), tx, uint32(chainID.Uint64()), common.BigToHash(blockNumber), blockNumber.Uint64(), uint64(index))
	Nil(n.T(), err)
}
