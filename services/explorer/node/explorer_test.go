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
	indexerConfig "github.com/synapsecns/sanguine/services/explorer/config/indexer"
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
func (c NodeSuite) TestLive() {
	if os.Getenv("CI") != "" {
		c.T().Skip("Network / processing test flake")
	}
	chainConfigs := []indexerConfig.ChainConfig{}
	backends := make(map[uint32]bind.ContractBackend)
	// ethclient.DialContext(ctx, chainConfig.RPCURL)
	for k := range c.testBackends {
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
	deployInfo, bridgeConfigContract := c.deployManager.GetBridgeConfigV3(c.GetTestContext(), c.testBackends[c.blockConfigChainID])
	for _, token := range testTokens {
		auth := c.testBackends[c.blockConfigChainID].GetTxContext(c.GetTestContext(), deployInfo.OwnerPtr())
		tx, err := token.SetTokenConfig(bridgeConfigContract, auth)
		c.Require().NoError(err)
		c.testBackends[c.blockConfigChainID].WaitForConfirmation(c.GetTestContext(), tx)
	}
	for k := range c.testBackends {
		backends[k] = c.testBackends[k]
		bridgeContract, bridgeRef := c.testDeployManager.GetTestSynapseBridge(c.GetTestContext(), c.testBackends[k])
		swapContractA, swapRefA := c.testDeployManager.GetTestSwapFlashLoan(c.GetTestContext(), c.testBackends[k])
		testDeployManagerB := testcontracts.NewDeployManager(c.T())
		swapContractB, swapRefB := testDeployManagerB.GetTestSwapFlashLoan(c.GetTestContext(), c.testBackends[k])
		transactOpts := c.testBackends[k].GetTxContext(c.GetTestContext(), nil)

		contracts := []indexerConfig.ContractConfig{
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
		chainConfigs = append(chainConfigs, indexerConfig.ChainConfig{
			ChainID:             k,
			RPCURL:              gofakeit.URL(),
			FetchBlockIncrement: 100,
			MaxGoroutines:       5,
			Contracts:           contracts,
		})
		// go through each contract and save the end height in scribe
		for i := range contracts {
			//  the last block store per contract
			err := c.eventDB.StoreLastIndexed(c.GetTestContext(), common.HexToAddress(contracts[i].Address), k, 12, false)
			Nil(c.T(), err)
		}
		c.fillBlocks(bridgeRef, swapRefA, swapRefB, transactOpts, k)
	}

	// This structure is for reference
	explorerConfig := indexerConfig.Config{
		DefaultRefreshRate:  2,
		ScribeURL:           c.gqlClient.Client.BaseURL,
		BridgeConfigAddress: deployInfo.Address().String(),
		BridgeConfigChainID: c.blockConfigChainID,
		Chains:              chainConfigs,
	}

	explorerBackfiller, err := node.NewExplorerBackfiller(c.db, explorerConfig, backends, c.explorerMetrics)
	c.Nil(err)
	c.NotNil(explorerBackfiller)
	err = explorerBackfiller.Backfill(c.GetTestContext(), false)
	c.Nil(err)
	var counttemp int64
	dd := c.db.UNSAFE_DB().WithContext(c.GetTestContext()).Table("swap_events").Count(&counttemp)
	c.Nil(dd.Error)

	var count int64
	bridgeEvents := c.db.UNSAFE_DB().WithContext(c.GetTestContext()).Find(&sql.BridgeEvent{}).Count(&count)
	Nil(c.T(), bridgeEvents.Error)
	Equal(c.T(), int64(10*len(c.testBackends)), count)

	swapEvents := c.db.UNSAFE_DB().WithContext(c.GetTestContext()).Find(&sql.SwapEvent{}).Count(&count)
	Nil(c.T(), swapEvents.Error)
	Equal(c.T(), int64(10*len(c.testBackends)), count)

	for k := range c.testBackends {
		bridgeEventsChain := c.db.UNSAFE_DB().WithContext(c.GetTestContext()).Model(&sql.BridgeEvent{}).Where(&sql.BridgeEvent{ChainID: k}).Count(&count)
		Nil(c.T(), bridgeEventsChain.Error)
		Equal(c.T(), int64(10), count)

		swapEventsChain := c.db.UNSAFE_DB().WithContext(c.GetTestContext()).Model(&sql.SwapEvent{}).Where(&sql.SwapEvent{ChainID: k}).Count(&count)
		Nil(c.T(), swapEventsChain.Error)
		Equal(c.T(), int64(10), count)
	}
}

// nolinting until parity tests implemented
//
//nolint:unparam
func (c *NodeSuite) storeTestLog(tx *types.Transaction, chainID uint32, blockNumber uint64) (*types.Log, error) {
	c.testBackends[chainID].WaitForConfirmation(c.GetTestContext(), tx)
	receipt, err := c.testBackends[chainID].TransactionReceipt(c.GetTestContext(), tx.Hash())

	if err != nil {
		return nil, fmt.Errorf("failed to get receipt for transaction %s: %w", tx.Hash().String(), err)
	}

	receipt.Logs[0].BlockNumber = blockNumber

	err = c.eventDB.StoreLogs(c.GetTestContext(), chainID, *receipt.Logs[0])
	if err != nil {
		return nil, fmt.Errorf("error storing swap log: %w", err)
	}

	return receipt.Logs[0], nil
}

func (c NodeSuite) fillBlocks(bridgeRef *testbridge.TestBridgeRef, swapRefA *testswap.TestSwapRef, swapRefB *testswap.TestSwapRef, transactOpts backends.AuthType, chainID uint32) {
	// Store blocktimes for testing defillama and timestamp indexing.
	for i := uint64(0); i < 13; i++ {
		err := c.eventDB.StoreBlockTime(c.GetTestContext(), chainID, i, i)
		Nil(c.T(), err)
	}

	bridgeTx, err := bridgeRef.TestDeposit(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(chainID)), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(c.T(), err)
	c.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(5)), 1)
	_, err = c.storeTestLog(bridgeTx, chainID, 5)
	Nil(c.T(), err)

	bridgeTx, err = bridgeRef.TestRedeem(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(int64(gofakeit.Uint32()))), big.NewInt(int64(chainID)), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(c.T(), err)
	c.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(5)), 2)
	_, err = c.storeTestLog(bridgeTx, chainID, 5)
	Nil(c.T(), err)

	bridgeTx, err = bridgeRef.TestWithdraw(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(c.T(), err)
	c.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(6)), 3)
	_, err = c.storeTestLog(bridgeTx, chainID, 6)
	Nil(c.T(), err)

	bridgeTx, err = bridgeRef.TestMint(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(c.T(), err)
	c.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(6)), 4)
	_, err = c.storeTestLog(bridgeTx, chainID, 6)
	Nil(c.T(), err)

	bridgeTx, err = bridgeRef.TestDepositAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(chainID)), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(c.T(), err)
	c.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(6)), 5)
	_, err = c.storeTestLog(bridgeTx, chainID, 6)
	Nil(c.T(), err)

	bridgeTx, err = bridgeRef.TestRedeemAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(chainID)), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(c.T(), err)
	c.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(7)), 1)
	_, err = c.storeTestLog(bridgeTx, chainID, 7)
	Nil(c.T(), err)

	bridgeTx, err = bridgeRef.TestRedeemAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(chainID)), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	Nil(c.T(), err)
	c.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(8)), 1)
	_, err = c.storeTestLog(bridgeTx, chainID, 8)
	Nil(c.T(), err)

	bridgeTx, err = bridgeRef.TestMintAndSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(c.T(), err)
	c.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(9)), 1)
	_, err = c.storeTestLog(bridgeTx, chainID, 9)
	Nil(c.T(), err)

	bridgeTx, err = bridgeRef.TestWithdrawAndRemove(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), [32]byte{byte(gofakeit.Uint64())})
	Nil(c.T(), err)
	c.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(10)), 1)
	_, err = c.storeTestLog(bridgeTx, chainID, 10)
	Nil(c.T(), err)

	bridgeTx, err = bridgeRef.TestRedeemV2(transactOpts.TransactOpts, [32]byte{byte(gofakeit.Uint64())}, big.NewInt(int64(chainID)), common.HexToAddress(testTokens[chainID].TokenAddress), big.NewInt(int64(gofakeit.Uint32())))
	Nil(c.T(), err)
	c.storeEthTx(bridgeTx, big.NewInt(int64(chainID)), big.NewInt(int64(12)), 1)
	_, err = c.storeTestLog(bridgeTx, chainID, 12)
	Nil(c.T(), err)

	// Store every swap event across two different swap contracts.
	swapTx, err := swapRefA.TestSwap(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(c.T(), err)
	c.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(5)), 1)
	_, err = c.storeTestLog(swapTx, chainID, 5)
	Nil(c.T(), err)

	swapTx, err = swapRefB.TestAddLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(c.T(), err)
	c.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(5)), 2)
	_, err = c.storeTestLog(swapTx, chainID, 5)
	Nil(c.T(), err)

	swapTx, err = swapRefB.TestRemoveLiquidity(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())))
	Nil(c.T(), err)
	c.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(6)), 1)
	_, err = c.storeTestLog(swapTx, chainID, 6)
	Nil(c.T(), err)

	swapTx, err = swapRefA.TestRemoveLiquidityOne(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(c.T(), err)
	c.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(7)), 1)
	_, err = c.storeTestLog(swapTx, chainID, 7)
	Nil(c.T(), err)

	swapTx, err = swapRefA.TestRemoveLiquidityImbalance(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, []*big.Int{big.NewInt(int64(gofakeit.Uint64()))}, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(c.T(), err)
	c.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(8)), 1)
	_, err = c.storeTestLog(swapTx, chainID, 8)
	Nil(c.T(), err)

	swapTx, err = swapRefB.TestNewAdminFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(c.T(), err)
	c.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(8)), 2)
	_, err = c.storeTestLog(swapTx, chainID, 8)
	Nil(c.T(), err)

	swapTx, err = swapRefA.TestNewSwapFee(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())))
	Nil(c.T(), err)
	c.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(8)), 3)
	_, err = c.storeTestLog(swapTx, chainID, 8)
	Nil(c.T(), err)

	swapTx, err = swapRefA.TestRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(c.T(), err)
	c.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(9)), 1)
	_, err = c.storeTestLog(swapTx, chainID, 9)
	Nil(c.T(), err)

	swapTx, err = swapRefB.TestStopRampA(transactOpts.TransactOpts, big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(c.T(), err)
	c.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(9)), 2)
	_, err = c.storeTestLog(swapTx, chainID, 9)
	Nil(c.T(), err)

	swapTx, err = swapRefA.TestFlashLoan(transactOpts.TransactOpts, common.BigToAddress(big.NewInt(gofakeit.Int64())), gofakeit.Uint8(), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())), big.NewInt(int64(gofakeit.Uint64())))
	Nil(c.T(), err)
	c.storeEthTx(swapTx, big.NewInt(int64(chainID)), big.NewInt(int64(9)), 3)
	_, err = c.storeTestLog(swapTx, chainID, 9)
	Nil(c.T(), err)
}

// storeEthTx stores the eth transaction so the get sender functionality can be tested.
func (c *NodeSuite) storeEthTx(tx *types.Transaction, chainID *big.Int, blockNumber *big.Int, index int) {
	err := c.eventDB.StoreEthTx(c.GetTestContext(), tx, uint32(chainID.Uint64()), common.BigToHash(blockNumber), blockNumber.Uint64(), uint64(index))
	Nil(c.T(), err)
}
