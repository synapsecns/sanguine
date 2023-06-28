package node_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/node"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"
	"math/big"
	"os"
	"sync"
	"time"
)

// TODO combine these functions with backfill/backend as well as other tests

// ReachBlockHeight reaches a block height on a backend.
func (l *LiveSuite) ReachBlockHeight(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			l.T().Log(ctx.Err())
			return
		default:
			// continue
		}
		i++
		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))

		latestBlock, err := backend.BlockNumber(ctx)
		Nil(l.T(), err)

		if latestBlock >= desiredBlockHeight {
			return
		}
	}
}

// startOmnirpcServer boots an omnirpc server for an rpc address.
// the url for this rpc is returned.
func (l *LiveSuite) startOmnirpcServer(ctx context.Context, backend backends.SimulatedTestBackend) string {
	baseHost := testhelper.NewOmnirpcServer(ctx, l.T(), backend)
	return testhelper.GetURL(baseHost, backend)
}

// ReachBlockHeight reaches a block height on a backend.
func (l *LiveSuite) PopuluateWithLogs(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64) common.Address {
	i := 0
	var address common.Address
	for {
		select {
		case <-ctx.Done():
			l.T().Log(ctx.Err())
			return address
		default:
			// continue
		}
		i++
		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))
		testContract, testRef := l.manager.GetTestContract(l.GetTestContext(), backend)
		address = testContract.Address()
		transactOpts := backend.GetTxContext(l.GetTestContext(), nil)
		tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		Nil(l.T(), err)
		backend.WaitForConfirmation(l.GetTestContext(), tx)

		latestBlock, err := backend.BlockNumber(ctx)
		Nil(l.T(), err)

		if latestBlock >= desiredBlockHeight {
			return address
		}
	}
}
func (l *LiveSuite) TestGetBlockHashes() {
	testBackend := geth.NewEmbeddedBackend(l.GetTestContext(), l.T())

	var wg sync.WaitGroup
	wg.Add(2)

	const desiredBlockHeight = 16

	go func() {
		defer wg.Done()
		l.ReachBlockHeight(l.GetTestContext(), testBackend, desiredBlockHeight)
	}()

	var host string
	go func() {
		defer wg.Done()
		host = l.startOmnirpcServer(l.GetTestContext(), testBackend)
	}()

	wg.Wait()

	scribeBackend, err := backfill.DialBackend(l.GetTestContext(), host, l.metrics)
	Nil(l.T(), err)
	hashes, err := node.GetBlockHashes(l.GetTestContext(), scribeBackend, 1, desiredBlockHeight, 3)
	Nil(l.T(), err)

	// Check that the number of hashes is as expected
	Equal(l.T(), desiredBlockHeight, len(hashes))

	// use to make sure we don't double use values
	hashSet := make(map[string]bool)

	for _, hash := range hashes {
		_, ok := hashSet[hash]
		False(l.T(), ok, "hash %s appears at least twice", hash)
		hashSet[hash] = true
	}
}

// TestLive tests live recording of events.
func (l LiveSuite) TestLive() {
	if os.Getenv("CI") != "" {
		l.T().Skip("Test flake: 20 sec of livefilling may fail on CI")
	}
	chainID := gofakeit.Uint32()
	// We need to set up multiple deploy managers, one for each contract. We will use
	// b.manager for the first contract, and create a new ones for the next two.
	managerB := testutil.NewDeployManager(l.T())
	managerC := testutil.NewDeployManager(l.T())
	// Get simulated blockchain, deploy three test contracts, and set up test variables.
	simulatedChain := geth.NewEmbeddedBackendForChainID(l.GetTestContext(), l.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backfill.DialBackend(l.GetTestContext(), simulatedChain.RPCAddress(), l.metrics)
	Nil(l.T(), err)

	simulatedChain.FundAccount(l.GetTestContext(), l.wallet.Address(), *big.NewInt(params.Ether))
	testContractA, testRefA := l.manager.GetTestContract(l.GetTestContext(), simulatedChain)
	testContractB, testRefB := managerB.GetTestContract(l.GetTestContext(), simulatedChain)
	testContractC, testRefC := managerC.GetTestContract(l.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(l.GetTestContext(), nil)
	// Put the contracts into a slice so we can iterate over them.
	contracts := []contracts.DeployedContract{testContractA, testContractB, testContractC}
	// Put the test refs into a slice so we can iterate over them.
	testRefs := []*testcontract.TestContractRef{testRefA, testRefB, testRefC}

	// Set up the config.
	contractConfigs := config.ContractConfigs{}
	for _, contract := range contracts {
		contractConfigs = append(contractConfigs, config.ContractConfig{
			Address:    contract.Address().String(),
			StartBlock: 0,
		})
	}
	chainConfig := config.ChainConfig{
		ChainID:             chainID,
		Contracts:           contractConfigs,
		GetBlockBatchAmount: 1,
		GetLogsBatchAmount:  2,
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{chainConfig},
	}

	clients := make(map[uint32][]backfill.ScribeBackend)
	clients[chainID] = append(clients[chainID], simulatedClient)
	clients[chainID] = append(clients[chainID], simulatedClient)

	// Set up the scribe.
	scribe, err := node.NewScribe(l.testDB, clients, scribeConfig, l.metrics)
	Nil(l.T(), err)

	for _, testRef := range testRefs {
		tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		Nil(l.T(), err)
		simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)
		tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
		Nil(l.T(), err)
		simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)
		tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
		Nil(l.T(), err)
		simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)
	}

	// Livefill for a minute.
	ctx, cancel := context.WithTimeout(l.GetTestContext(), 20*time.Second)
	defer cancel()
	_ = scribe.Start(ctx)

	// Check that the events were recorded.
	for _, contract := range contracts {
		// Check the storage of logs.
		logFilter := db.LogFilter{
			ChainID:         chainConfig.ChainID,
			ContractAddress: contract.Address().String(),
		}
		logs, err := l.testDB.RetrieveLogsWithFilter(l.GetTestContext(), logFilter, 1)
		Nil(l.T(), err)
		// There should be 4 logs. One from `EmitEventA`, one from `EmitEventB`, and two
		// from `EmitEventAandB`.
		Equal(l.T(), 4, len(logs))
	}
	// Check the storage of receipts.
	receiptFilter := db.ReceiptFilter{
		ChainID: chainConfig.ChainID,
	}
	receipts, err := l.testDB.RetrieveReceiptsWithFilter(l.GetTestContext(), receiptFilter, 1)
	Nil(l.T(), err)
	// There should be 9 receipts. One from `EmitEventA`, one from `EmitEventB`, and
	// one from `EmitEventAandB`, for each contract.
	Equal(l.T(), 9, len(receipts))
}

func (l LiveSuite) TestConfirmationSimple() {
	if os.Getenv("CI") != "" {
		l.T().Skip("Test flake: 20 seconds of livefilling may fail on CI")
	}
	chainID := gofakeit.Uint32()

	// Emit some events on the simulated blockchain.
	simulatedChain := geth.NewEmbeddedBackendForChainID(l.GetTestContext(), l.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backfill.DialBackend(l.GetTestContext(), simulatedChain.RPCAddress(), l.metrics)
	Nil(l.T(), err)

	simulatedChain.FundAccount(l.GetTestContext(), l.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := l.manager.GetTestContract(l.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(l.GetTestContext(), nil)

	// Set up the config.
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}
	chainConfig := config.ChainConfig{
		ChainID:   chainID,
		Contracts: []config.ContractConfig{contractConfig},
		ConfirmationConfig: config.ConfirmationConfig{
			RequiredConfirmations:   100,
			ConfirmationThreshold:   1,
			ConfirmationRefreshRate: 1,
		},
	}
	scribeConfig := config.Config{
		Chains:                  []config.ChainConfig{chainConfig},
		ConfirmationRefreshRate: 1,
	}

	clients := make(map[uint32][]backfill.ScribeBackend)
	clients[chainID] = append(clients[chainID], simulatedClient)
	clients[chainID] = append(clients[chainID], simulatedClient)

	// Set up the scribe.
	scribe, err := node.NewScribe(l.testDB, clients, scribeConfig, l.metrics)
	Nil(l.T(), err)

	// Emit 5 events.
	for i := 0; i < 5; i++ {
		tx, err := testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		Nil(l.T(), err)
		simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)
	}
	// Process the events, end livefilling after 20 seconds.
	ctx, cancel := context.WithTimeout(l.GetTestContext(), 20*time.Second)
	defer cancel()
	_ = scribe.Start(ctx)

	// Check if values are confirmed
	logFilter := db.LogFilter{
		ChainID:         chainConfig.ChainID,
		ContractAddress: testContract.Address().String(),
		Confirmed:       true,
	}
	logs, err := l.testDB.RetrieveLogsWithFilter(l.GetTestContext(), logFilter, 1)
	Nil(l.T(), err)
	Equal(l.T(), 8, len(logs))
	receiptFilter := db.ReceiptFilter{
		ChainID:   chainConfig.ChainID,
		Confirmed: true,
	}
	receipts, err := l.testDB.RetrieveReceiptsWithFilter(l.GetTestContext(), receiptFilter, 1)
	Nil(l.T(), err)
	Equal(l.T(), 4, len(receipts))
	txFilter := db.EthTxFilter{
		ChainID:   chainConfig.ChainID,
		Confirmed: true,
	}

	txs, err := l.testDB.RetrieveEthTxsWithFilter(l.GetTestContext(), txFilter, 1)
	Nil(l.T(), err)
	Equal(l.T(), 4, len(txs))

	lastConfirmedBlock, err := l.testDB.RetrieveLastConfirmedBlock(l.GetTestContext(), chainConfig.ChainID)
	Nil(l.T(), err)
	Equal(l.T(), uint64(8), lastConfirmedBlock)

	lastBlockIndexed, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), testContract.Address(), chainConfig.ChainID)
	Nil(l.T(), err)
	Equal(l.T(), uint64(9), lastBlockIndexed)
}

func (l LiveSuite) TestRequiredConfirmationRemAndAdd() {
	if os.Getenv("CI") != "" {
		l.T().Skip("Test flake: 20 seconds of livefilling may fail on CI")
	}
	chainID := gofakeit.Uint32()

	// Emit some events on the simulated blockchain.
	simulatedChain := geth.NewEmbeddedBackendForChainID(l.GetTestContext(), l.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backfill.DialBackend(l.GetTestContext(), simulatedChain.RPCAddress(), l.metrics)
	Nil(l.T(), err)

	simulatedChain.FundAccount(l.GetTestContext(), l.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := l.manager.GetTestContract(l.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(l.GetTestContext(), nil)

	// Set up the config.
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}
	chainConfig := config.ChainConfig{
		ChainID:   chainID,
		Contracts: []config.ContractConfig{contractConfig},
		ConfirmationConfig: config.ConfirmationConfig{
			RequiredConfirmations:   100,
			ConfirmationThreshold:   1,
			ConfirmationRefreshRate: 1,
		},
	}
	scribeConfig := config.Config{
		Chains:                  []config.ChainConfig{chainConfig},
		ConfirmationRefreshRate: 1,
	}

	clients := make(map[uint32][]backfill.ScribeBackend)
	clients[chainID] = append(clients[chainID], simulatedClient)
	clients[chainID] = append(clients[chainID], simulatedClient)

	// Set up scribe.
	scribe, err := node.NewScribe(l.testDB, clients, scribeConfig, l.metrics)
	Nil(l.T(), err)

	for i := 0; i < 5; i++ {
		tx, err := testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		Nil(l.T(), err)
		simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)
	}
	// Process the events, end livefilling after 20 seconds.
	ctx, cancel := context.WithTimeout(l.GetTestContext(), 20*time.Second)
	defer cancel()

	invalidBlockHash := common.BigToHash(big.NewInt(11111))
	invalidReceipt := types.Receipt{
		ContractAddress: testContract.Address(),
		BlockHash:       invalidBlockHash,
		BlockNumber:     big.NewInt(3),
	}
	receiptFilter := db.ReceiptFilter{
		ChainID: chainConfig.ChainID,
	}
	// Storing an invalid receipt with a nonsense block hash. The proper behavior will be to evict/rm this receipt upon
	// confirmation checking and re-backfill the block.
	err = l.testDB.StoreReceipt(l.GetTestContext(), chainConfig.ChainID, invalidReceipt)
	Nil(l.T(), err)
	startingReceipts, err := l.testDB.RetrieveReceiptsWithFilter(l.GetTestContext(), receiptFilter, 1)
	Nil(l.T(), err)
	Equal(l.T(), 1, len(startingReceipts))

	_ = scribe.Start(ctx)

	// Check if values are confirmed
	logFilter := db.LogFilter{
		ChainID:         chainConfig.ChainID,
		ContractAddress: testContract.Address().String(),
		Confirmed:       true,
	}
	logs, err := l.testDB.RetrieveLogsWithFilter(l.GetTestContext(), logFilter, 1)
	Nil(l.T(), err)
	Equal(l.T(), 8, len(logs))

	receipts, err := l.testDB.RetrieveReceiptsWithFilter(l.GetTestContext(), receiptFilter, 1)
	Nil(l.T(), err)
	for _, receipt := range receipts {
		NotEqual(l.T(), receipt.BlockHash, invalidBlockHash)
	}
	Equal(l.T(), 5, len(receipts))

	txFilter := db.EthTxFilter{
		ChainID:   chainConfig.ChainID,
		Confirmed: true,
	}
	txs, err := l.testDB.RetrieveEthTxsWithFilter(l.GetTestContext(), txFilter, 1)
	Nil(l.T(), err)
	Equal(l.T(), 4, len(txs))

	lastConfirmedBlock, err := l.testDB.RetrieveLastConfirmedBlock(l.GetTestContext(), chainConfig.ChainID)
	Nil(l.T(), err)
	Equal(l.T(), 9-chainConfig.ConfirmationConfig.ConfirmationThreshold, lastConfirmedBlock)

	lastBlockIndexed, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), testContract.Address(), chainConfig.ChainID)
	Nil(l.T(), err)
	Equal(l.T(), uint64(9), lastBlockIndexed)
}

// TODO finish this test
// TestLivefillParity runs livefill on certain prod chains. Then it checks parity with an explorer API.
// func (l LiveSuite) TestLivefillParity() {
//	originAddress := "0xF3773BE7cb59235Ced272cF324aaeb0A4115280f"
//	destinationAddress := "0xde5BB62aBCF588EC200674757EDB2f6889aCd065"
//	summitAddress := "0x128fF47f1a614c61beC9935898C33B91486aA04e"
//
//	maticOriginStart := uint64(40189736)
//	maticDestinationStart := uint64(40189736)
//	avaxOriginStart := uint64(27262747)
//	avaxDestinationStart := uint64(27262744)
//	opSummitStart := uint64(79864182)
//
//	maticChainID := uint32(137)
//	avaxChainID := uint32(43114)
//	opChainID := uint32(10)
//
//	maticRPCURL := "https://polygon-mainnet.g.alchemy.com/v2/Kmd9QLE1B3CFtVH879DJKsAvv92LV0E2"
//	avaxRPCURL := "https://1rpc.io/avax/c"
//	opRPCURL := "https://optimism-rpc.gateway.pokt.network/"
//
//	maticClient, err := backfill.DialBackend(l.GetTestContext(), maticRPCURL, l.metrics)
//	Nil(l.T(), err)
//	avaxClient, err := backfill.DialBackend(l.GetTestContext(), avaxRPCURL, l.metrics)
//	Nil(l.T(), err)
//	opClient, err := backfill.DialBackend(l.GetTestContext(), opRPCURL, l.metrics)
//	Nil(l.T(), err)
//
//	scribeConfig := config.Config{
//		Chains: []config.ChainConfig{
//			{
//				ChainID:                   maticChainID,
//				RequiredConfirmations:     0,
//				GetLogsRange:      256,
//				GetLogsBatchAmount:         2,
//				ConcurrencyThreshold: 100000,
//				Contracts: []config.ContractConfig{
//					{
//						Address:    originAddress,
//						StartBlock: maticOriginStart,
//					},
//					{
//						Address:    destinationAddress,
//						StartBlock: maticDestinationStart,
//					},
//				},
//			},
//			{
//				ChainID:                   avaxChainID,
//				RequiredConfirmations:     0,
//				GetLogsRange:      256,
//				GetLogsBatchAmount:         2,
//				ConcurrencyThreshold: 100000,
//				Contracts: []config.ContractConfig{
//					{
//						Address:    originAddress,
//						StartBlock: avaxOriginStart,
//					},
//					{
//						Address:    destinationAddress,
//						StartBlock: avaxDestinationStart,
//					},
//				},
//			},
//			{
//				ChainID:                   opChainID,
//				RequiredConfirmations:     0,
//				StoreConcurrency:          1,
//				ConcurrencyThreshold: 100000,
//				Contracts: []config.ContractConfig{
//					{
//						Address:    summitAddress,
//						StartBlock: opSummitStart,
//					},
//				},
//			},
//		},
//	}
//
//	clients := map[uint32][]backfill.ScribeBackend{
//		maticChainID: {maticClient, maticClient},
//		avaxChainID:  {avaxClient, avaxClient},
//		opChainID:    {opClient, opClient},
//	}
//
//	// Get the current block for each chain.
//	maticCurrentBlock, err := maticClient.BlockNumber(l.GetTestContext())
//	Nil(l.T(), err)
//	avaxCurrentBlock, err := avaxClient.BlockNumber(l.GetTestContext())
//	Nil(l.T(), err)
//	opCurrentBlock, err := opClient.BlockNumber(l.GetTestContext())
//	Nil(l.T(), err)
//
//	scribe, err := node.NewScribe(l.testDB, clients, scribeConfig, l.metrics)
//	Nil(l.T(), err)
//
//	killableContext, cancel := context.WithCancel(l.GetTestContext())
//
//	go func() {
//		_ = scribe.Start(killableContext)
//	}()
//
//	waitChan := make(chan bool, 3)
//
//	// Wait for scribe to get to/past the current block for each chain.
//
//	// Wait on Polygon.
//	l.Eventually(func() bool {
//		originBlock, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), common.HexToAddress(originAddress), maticChainID)
//		Nil(l.T(), err)
//		destinationBlock, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), common.HexToAddress(destinationAddress), maticChainID)
//		Nil(l.T(), err)
//
//		if originBlock >= maticCurrentBlock && destinationBlock >= maticCurrentBlock {
//			waitChan <- true
//			return true
//		}
//
//		time.Sleep(5 * time.Second)
//
//		return false
//	})
//
//	// Wait on Avalanche.
//	l.Eventually(func() bool {
//		originBlock, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), common.HexToAddress(originAddress), avaxChainID)
//		Nil(l.T(), err)
//		destinationBlock, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), common.HexToAddress(destinationAddress), avaxChainID)
//		Nil(l.T(), err)
//
//		if originBlock >= avaxCurrentBlock && destinationBlock >= avaxCurrentBlock {
//			waitChan <- true
//			return true
//		}
//
//		time.Sleep(5 * time.Second)
//
//		return false
//	})
//
//	// Wait on Optimism.
//	l.Eventually(func() bool {
//		summitBlock, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), common.HexToAddress(summitAddress), opChainID)
//		Nil(l.T(), err)
//
//		if summitBlock >= opCurrentBlock {
//			waitChan <- true
//			return true
//		}
//
//		time.Sleep(5 * time.Second)
//
//		return false
//	})
//
//	// Do not continue until all chains have reached the current block.
//	<-waitChan
//	<-waitChan
//	<-waitChan
//
//	// Stop the scribe.
//	cancel()
//
//	// Get the last indexed block for each chain.
//	maticLastIndexed, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), common.HexToAddress(originAddress), maticChainID)
//	Nil(l.T(), err)
//	avaxLastIndexed, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), common.HexToAddress(originAddress), avaxChainID)
//	Nil(l.T(), err)
//	opLastIndexed, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), common.HexToAddress(summitAddress), opChainID)
//	Nil(l.T(), err)
//
//	// Get the number of logs stored in the scribe DB for each chain.
//	logFilter := db.LogFilter{
//		ChainID: maticChainID,
//	}
//	maticLogAmount, err := getLogAmount(l.GetTestContext(), l.testDB, logFilter, maticOriginStart, maticCurrentBlock)
//	Nil(l.T(), err)
//
//	logFilter = db.LogFilter{
//		ChainID: avaxChainID,
//	}
//	avaxLogAmount, err := getLogAmount(l.GetTestContext(), l.testDB, logFilter, avaxOriginStart, avaxCurrentBlock)
//	Nil(l.T(), err)
//
//	logFilter = db.LogFilter{
//		ChainID: opChainID,
//	}
//	opLogAmount, err := getLogAmount(l.GetTestContext(), l.testDB, logFilter, opSummitStart, opCurrentBlock)
//	Nil(l.T(), err)
//
//	fmt.Println("Matic last indexed:", maticLastIndexed)
//	fmt.Println("Avalanche last indexed:", avaxLastIndexed)
//	fmt.Println("Optimism last indexed:", opLastIndexed)
//
//	fmt.Println("Matic log amount:", maticLogAmount)
//	fmt.Println("Avalanche log amount:", avaxLogAmount)
//	fmt.Println("Optimism log amount:", opLogAmount)
//
//	// Use explorer API's to do a parity check on the number of logs.
//}
//
// func getLogAmount(ctx context.Context, db db.EventDB, filter db.LogFilter, startBlock, endBlock uint64) (int, error) {
//	page := 1
//
//	var retrievedLogs []*types.Log
//
//	for {
//		logs, err := db.RetrieveLogsInRangeAsc(ctx, filter, startBlock, endBlock, page)
//		if err != nil {
//			return 0, err
//		}
//
//		retrievedLogs = append(retrievedLogs, logs...)
//
//		if len(logs) < base.PageSize {
//			break
//		}
//
//		page++
//	}
//
//	return len(retrievedLogs), nil
//}
