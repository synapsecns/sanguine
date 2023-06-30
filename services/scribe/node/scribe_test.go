package node_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/base"
	"github.com/synapsecns/sanguine/services/scribe/node"
	"net/http"
	"time"
)

// TODO combine these functions with backfill/backend as well as other tests

//// ReachBlockHeight reaches a block height on a backend.
//func (l *LiveSuite) ReachBlockHeight(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64) {
//	i := 0
//	for {
//		select {
//		case <-ctx.Done():
//			l.T().Log(ctx.Err())
//			return
//		default:
//			// continue
//		}
//		i++
//		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))
//
//		latestBlock, err := backend.BlockNumber(ctx)
//		Nil(l.T(), err)
//
//		if latestBlock >= desiredBlockHeight {
//			return
//		}
//	}
//}
//
//// startOmnirpcServer boots an omnirpc server for an rpc address.
//// the url for this rpc is returned.
//func (l *LiveSuite) startOmnirpcServer(ctx context.Context, backend backends.SimulatedTestBackend) string {
//	baseHost := testhelper.NewOmnirpcServer(ctx, l.T(), backend)
//	return testhelper.GetURL(baseHost, backend)
//}
//
//// ReachBlockHeight reaches a block height on a backend.
//func (l *LiveSuite) PopuluateWithLogs(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64) common.Address {
//	i := 0
//	var address common.Address
//	for {
//		select {
//		case <-ctx.Done():
//			l.T().Log(ctx.Err())
//			return address
//		default:
//			// continue
//		}
//		i++
//		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))
//		testContract, testRef := l.manager.GetTestContract(l.GetTestContext(), backend)
//		address = testContract.Address()
//		transactOpts := backend.GetTxContext(l.GetTestContext(), nil)
//		tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
//		Nil(l.T(), err)
//		backend.WaitForConfirmation(l.GetTestContext(), tx)
//
//		latestBlock, err := backend.BlockNumber(ctx)
//		Nil(l.T(), err)
//
//		if latestBlock >= desiredBlockHeight {
//			return address
//		}
//	}
//}
//func (l *LiveSuite) TestGetBlockHashes() {
//	testBackend := geth.NewEmbeddedBackend(l.GetTestContext(), l.T())
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	const desiredBlockHeight = 16
//
//	go func() {
//		defer wg.Done()
//		l.ReachBlockHeight(l.GetTestContext(), testBackend, desiredBlockHeight)
//	}()
//
//	var host string
//	go func() {
//		defer wg.Done()
//		host = l.startOmnirpcServer(l.GetTestContext(), testBackend)
//	}()
//
//	wg.Wait()
//
//	scribeBackend, err := backfill.DialBackend(l.GetTestContext(), host, l.metrics)
//	Nil(l.T(), err)
//	hashes, err := node.GetBlockHashes(l.GetTestContext(), scribeBackend, 1, desiredBlockHeight, 3)
//	Nil(l.T(), err)
//
//	// Check that the number of hashes is as expected
//	Equal(l.T(), desiredBlockHeight, len(hashes))
//
//	// use to make sure we don't double use values
//	hashSet := make(map[string]bool)
//
//	for _, hash := range hashes {
//		_, ok := hashSet[hash]
//		False(l.T(), ok, "hash %s appears at least twice", hash)
//		hashSet[hash] = true
//	}
//}
//
//// TestLive tests live recording of events.
//func (l LiveSuite) TestLive() {
//	if os.Getenv("CI") != "" {
//		l.T().Skip("Test flake: 20 sec of livefilling may fail on CI")
//	}
//	chainID := gofakeit.Uint32()
//	// We need to set up multiple deploy managers, one for each contract. We will use
//	// b.manager for the first contract, and create a new ones for the next two.
//	managerB := testutil.NewDeployManager(l.T())
//	managerC := testutil.NewDeployManager(l.T())
//	// Get simulated blockchain, deploy three test contracts, and set up test variables.
//	simulatedChain := geth.NewEmbeddedBackendForChainID(l.GetTestContext(), l.T(), big.NewInt(int64(chainID)))
//	simulatedClient, err := backfill.DialBackend(l.GetTestContext(), simulatedChain.RPCAddress(), l.metrics)
//	Nil(l.T(), err)
//
//	simulatedChain.FundAccount(l.GetTestContext(), l.wallet.Address(), *big.NewInt(params.Ether))
//	testContractA, testRefA := l.manager.GetTestContract(l.GetTestContext(), simulatedChain)
//	testContractB, testRefB := managerB.GetTestContract(l.GetTestContext(), simulatedChain)
//	testContractC, testRefC := managerC.GetTestContract(l.GetTestContext(), simulatedChain)
//	transactOpts := simulatedChain.GetTxContext(l.GetTestContext(), nil)
//	// Put the contracts into a slice so we can iterate over them.
//	contracts := []contracts.DeployedContract{testContractA, testContractB, testContractC}
//	// Put the test refs into a slice so we can iterate over them.
//	testRefs := []*testcontract.TestContractRef{testRefA, testRefB, testRefC}
//
//	// Set up the config.
//	contractConfigs := config.ContractConfigs{}
//	for _, contract := range contracts {
//		contractConfigs = append(contractConfigs, config.ContractConfig{
//			Address:    contract.Address().String(),
//			StartBlock: 0,
//		})
//	}
//	chainConfig := config.ChainConfig{
//		ChainID:             chainID,
//		Contracts:           contractConfigs,
//		GetBlockBatchAmount: 1,
//		GetLogsBatchAmount:  2,
//	}
//	scribeConfig := config.Config{
//		Chains: []config.ChainConfig{chainConfig},
//	}
//
//	clients := make(map[uint32][]backfill.ScribeBackend)
//	clients[chainID] = append(clients[chainID], simulatedClient)
//	clients[chainID] = append(clients[chainID], simulatedClient)
//
//	// Set up the scribe.
//	scribe, err := node.NewScribe(l.testDB, clients, scribeConfig, l.metrics)
//	Nil(l.T(), err)
//
//	for _, testRef := range testRefs {
//		tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
//		Nil(l.T(), err)
//		simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)
//		tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
//		Nil(l.T(), err)
//		simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)
//		tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
//		Nil(l.T(), err)
//		simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)
//	}
//
//	// Livefill for a minute.
//	ctx, cancel := context.WithTimeout(l.GetTestContext(), 20*time.Second)
//	defer cancel()
//	_ = scribe.Start(ctx)
//
//	// Check that the events were recorded.
//	for _, contract := range contracts {
//		// Check the storage of logs.
//		logFilter := db.LogFilter{
//			ChainID:         chainConfig.ChainID,
//			ContractAddress: contract.Address().String(),
//		}
//		logs, err := l.testDB.RetrieveLogsWithFilter(l.GetTestContext(), logFilter, 1)
//		Nil(l.T(), err)
//		// There should be 4 logs. One from `EmitEventA`, one from `EmitEventB`, and two
//		// from `EmitEventAandB`.
//		Equal(l.T(), 4, len(logs))
//	}
//	// Check the storage of receipts.
//	receiptFilter := db.ReceiptFilter{
//		ChainID: chainConfig.ChainID,
//	}
//	receipts, err := l.testDB.RetrieveReceiptsWithFilter(l.GetTestContext(), receiptFilter, 1)
//	Nil(l.T(), err)
//	// There should be 9 receipts. One from `EmitEventA`, one from `EmitEventB`, and
//	// one from `EmitEventAandB`, for each contract.
//	Equal(l.T(), 9, len(receipts))
//}
//
//func (l LiveSuite) TestConfirmationSimple() {
//	if os.Getenv("CI") != "" {
//		l.T().Skip("Test flake: 20 seconds of livefilling may fail on CI")
//	}
//	chainID := gofakeit.Uint32()
//
//	// Emit some events on the simulated blockchain.
//	simulatedChain := geth.NewEmbeddedBackendForChainID(l.GetTestContext(), l.T(), big.NewInt(int64(chainID)))
//	simulatedClient, err := backfill.DialBackend(l.GetTestContext(), simulatedChain.RPCAddress(), l.metrics)
//	Nil(l.T(), err)
//
//	simulatedChain.FundAccount(l.GetTestContext(), l.wallet.Address(), *big.NewInt(params.Ether))
//	testContract, testRef := l.manager.GetTestContract(l.GetTestContext(), simulatedChain)
//	transactOpts := simulatedChain.GetTxContext(l.GetTestContext(), nil)
//
//	// Set up the config.
//	contractConfig := config.ContractConfig{
//		Address:    testContract.Address().String(),
//		StartBlock: 0,
//	}
//	chainConfig := config.ChainConfig{
//		ChainID:   chainID,
//		Contracts: []config.ContractConfig{contractConfig},
//		ConfirmationConfig: config.ConfirmationConfig{
//			RequiredConfirmations:   100,
//			ConfirmationThreshold:   1,
//			ConfirmationRefreshRate: 1,
//		},
//	}
//	scribeConfig := config.Config{
//		Chains:                  []config.ChainConfig{chainConfig},
//		ConfirmationRefreshRate: 1,
//	}
//
//	clients := make(map[uint32][]backfill.ScribeBackend)
//	clients[chainID] = append(clients[chainID], simulatedClient)
//	clients[chainID] = append(clients[chainID], simulatedClient)
//
//	// Set up the scribe.
//	scribe, err := node.NewScribe(l.testDB, clients, scribeConfig, l.metrics)
//	Nil(l.T(), err)
//
//	// Emit 5 events.
//	for i := 0; i < 5; i++ {
//		tx, err := testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
//		Nil(l.T(), err)
//		simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)
//	}
//	// Process the events, end livefilling after 20 seconds.
//	ctx, cancel := context.WithTimeout(l.GetTestContext(), 20*time.Second)
//	defer cancel()
//	_ = scribe.Start(ctx)
//
//	// Check if values are confirmed
//	logFilter := db.LogFilter{
//		ChainID:         chainConfig.ChainID,
//		ContractAddress: testContract.Address().String(),
//		Confirmed:       true,
//	}
//	logs, err := l.testDB.RetrieveLogsWithFilter(l.GetTestContext(), logFilter, 1)
//	Nil(l.T(), err)
//	Equal(l.T(), 8, len(logs))
//	receiptFilter := db.ReceiptFilter{
//		ChainID:   chainConfig.ChainID,
//		Confirmed: true,
//	}
//	receipts, err := l.testDB.RetrieveReceiptsWithFilter(l.GetTestContext(), receiptFilter, 1)
//	Nil(l.T(), err)
//	Equal(l.T(), 4, len(receipts))
//	txFilter := db.EthTxFilter{
//		ChainID:   chainConfig.ChainID,
//		Confirmed: true,
//	}
//
//	txs, err := l.testDB.RetrieveEthTxsWithFilter(l.GetTestContext(), txFilter, 1)
//	Nil(l.T(), err)
//	Equal(l.T(), 4, len(txs))
//
//	lastConfirmedBlock, err := l.testDB.RetrieveLastConfirmedBlock(l.GetTestContext(), chainConfig.ChainID)
//	Nil(l.T(), err)
//	Equal(l.T(), uint64(8), lastConfirmedBlock)
//
//	lastBlockIndexed, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), testContract.Address(), chainConfig.ChainID)
//	Nil(l.T(), err)
//	Equal(l.T(), uint64(9), lastBlockIndexed)
//}
//
//func (l LiveSuite) TestRequiredConfirmationRemAndAdd() {
//	if os.Getenv("CI") != "" {
//		l.T().Skip("Test flake: 20 seconds of livefilling may fail on CI")
//	}
//	chainID := gofakeit.Uint32()
//
//	// Emit some events on the simulated blockchain.
//	simulatedChain := geth.NewEmbeddedBackendForChainID(l.GetTestContext(), l.T(), big.NewInt(int64(chainID)))
//	simulatedClient, err := backfill.DialBackend(l.GetTestContext(), simulatedChain.RPCAddress(), l.metrics)
//	Nil(l.T(), err)
//
//	simulatedChain.FundAccount(l.GetTestContext(), l.wallet.Address(), *big.NewInt(params.Ether))
//	testContract, testRef := l.manager.GetTestContract(l.GetTestContext(), simulatedChain)
//	transactOpts := simulatedChain.GetTxContext(l.GetTestContext(), nil)
//
//	// Set up the config.
//	contractConfig := config.ContractConfig{
//		Address:    testContract.Address().String(),
//		StartBlock: 0,
//	}
//	chainConfig := config.ChainConfig{
//		ChainID:   chainID,
//		Contracts: []config.ContractConfig{contractConfig},
//		ConfirmationConfig: config.ConfirmationConfig{
//			RequiredConfirmations:   100,
//			ConfirmationThreshold:   1,
//			ConfirmationRefreshRate: 1,
//		},
//	}
//	scribeConfig := config.Config{
//		Chains:                  []config.ChainConfig{chainConfig},
//		ConfirmationRefreshRate: 1,
//	}
//
//	clients := make(map[uint32][]backfill.ScribeBackend)
//	clients[chainID] = append(clients[chainID], simulatedClient)
//	clients[chainID] = append(clients[chainID], simulatedClient)
//
//	// Set up scribe.
//	scribe, err := node.NewScribe(l.testDB, clients, scribeConfig, l.metrics)
//	Nil(l.T(), err)
//
//	for i := 0; i < 5; i++ {
//		tx, err := testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
//		Nil(l.T(), err)
//		simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)
//	}
//	// Process the events, end livefilling after 20 seconds.
//	ctx, cancel := context.WithTimeout(l.GetTestContext(), 20*time.Second)
//	defer cancel()
//
//	invalidBlockHash := common.BigToHash(big.NewInt(11111))
//	invalidReceipt := types.Receipt{
//		ContractAddress: testContract.Address(),
//		BlockHash:       invalidBlockHash,
//		BlockNumber:     big.NewInt(3),
//	}
//	receiptFilter := db.ReceiptFilter{
//		ChainID: chainConfig.ChainID,
//	}
//	// Storing an invalid receipt with a nonsense block hash. The proper behavior will be to evict/rm this receipt upon
//	// confirmation checking and re-backfill the block.
//	err = l.testDB.StoreReceipt(l.GetTestContext(), chainConfig.ChainID, invalidReceipt)
//	Nil(l.T(), err)
//	startingReceipts, err := l.testDB.RetrieveReceiptsWithFilter(l.GetTestContext(), receiptFilter, 1)
//	Nil(l.T(), err)
//	Equal(l.T(), 1, len(startingReceipts))
//
//	_ = scribe.Start(ctx)
//
//	// Check if values are confirmed
//	logFilter := db.LogFilter{
//		ChainID:         chainConfig.ChainID,
//		ContractAddress: testContract.Address().String(),
//		Confirmed:       true,
//	}
//	logs, err := l.testDB.RetrieveLogsWithFilter(l.GetTestContext(), logFilter, 1)
//	Nil(l.T(), err)
//	Equal(l.T(), 8, len(logs))
//
//	receipts, err := l.testDB.RetrieveReceiptsWithFilter(l.GetTestContext(), receiptFilter, 1)
//	Nil(l.T(), err)
//	for _, receipt := range receipts {
//		NotEqual(l.T(), receipt.BlockHash, invalidBlockHash)
//	}
//	Equal(l.T(), 5, len(receipts))
//
//	txFilter := db.EthTxFilter{
//		ChainID:   chainConfig.ChainID,
//		Confirmed: true,
//	}
//	txs, err := l.testDB.RetrieveEthTxsWithFilter(l.GetTestContext(), txFilter, 1)
//	Nil(l.T(), err)
//	Equal(l.T(), 4, len(txs))
//
//	lastConfirmedBlock, err := l.testDB.RetrieveLastConfirmedBlock(l.GetTestContext(), chainConfig.ChainID)
//	Nil(l.T(), err)
//	Equal(l.T(), 9-chainConfig.ConfirmationConfig.ConfirmationThreshold, lastConfirmedBlock)
//
//	lastBlockIndexed, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), testContract.Address(), chainConfig.ChainID)
//	Nil(l.T(), err)
//	Equal(l.T(), uint64(9), lastBlockIndexed)
//}

// TODO finish this test
// TestLivefillParity runs livefill on certain prod chains. Then it checks parity with an explorer API.
func (l LiveSuite) TestLivefillParity() {
	//ethRPCURL := "https://1rpc.io/eth"
	//arbRPCURL := "https://endpoints.omniatech.io/v1/arbitrum/one/public"
	//maticRPCURL := "https://poly-rpc.gateway.pokt.network"
	//avaxRPCURL := "https://avalanche.public-rpc.com"

	ethRPCURL := "https://rpc.interoperability.institute/confirmations/1/rpc/1"
	arbRPCURL := "https://rpc.interoperability.institute/confirmations/1/rpc/42161"
	maticRPCURL := "https://rpc.interoperability.institute/confirmations/1/rpc/137"
	avaxRPCURL := "https://rpc.interoperability.institute/confirmations/1/rpc/43114"
	bscRPCURL := "https://rpc.interoperability.institute/confirmations/1/rpc/56"

	blockRange := uint64(1000)

	ethClient, err := backfill.DialBackend(l.GetTestContext(), ethRPCURL, l.metrics)
	Nil(l.T(), err)
	arbClient, err := backfill.DialBackend(l.GetTestContext(), arbRPCURL, l.metrics)
	Nil(l.T(), err)
	maticClient, err := backfill.DialBackend(l.GetTestContext(), maticRPCURL, l.metrics)
	Nil(l.T(), err)
	avaxClient, err := backfill.DialBackend(l.GetTestContext(), avaxRPCURL, l.metrics)
	Nil(l.T(), err)
	bscClient, err := backfill.DialBackend(l.GetTestContext(), bscRPCURL, l.metrics)
	Nil(l.T(), err)

	ethID := uint32(1)
	//opID := uint32(10)
	bscID := uint32(56)
	arbID := uint32(42161)
	maticID := uint32(137)
	avaxID := uint32(43114)
	chains := []uint32{ethID, bscID, arbID, maticID, avaxID}

	// Get the current block for each chain.
	ethCurrentBlock, err := ethClient.BlockNumber(l.GetTestContext())
	Nil(l.T(), err)
	arbCurrentBlock, err := arbClient.BlockNumber(l.GetTestContext())
	Nil(l.T(), err)
	maticCurrentBlock, err := maticClient.BlockNumber(l.GetTestContext())
	Nil(l.T(), err)
	avaxCurrentBlock, err := avaxClient.BlockNumber(l.GetTestContext())
	Nil(l.T(), err)
	bscCurrentBlock, err := bscClient.BlockNumber(l.GetTestContext())
	Nil(l.T(), err)
	//opCurrentBlock, err := opClient.BlockNumber(l.GetTestContext())
	//Nil(l.T(), err)

	latestBlocks := map[uint32]uint64{
		ethID:   ethCurrentBlock,
		arbID:   arbCurrentBlock,
		maticID: maticCurrentBlock,
		avaxID:  avaxCurrentBlock,
		bscID:   bscCurrentBlock,
	}
	clients := map[uint32][]backfill.ScribeBackend{
		ethID:   {ethClient, ethClient},
		bscID:   {bscClient, bscClient},
		arbID:   {arbClient, arbClient},
		maticID: {maticClient, maticClient},
		avaxID:  {avaxClient, avaxClient},
	}

	apiUrls := map[uint32]string{
		ethID:   "https://api.etherscan.io/api",
		bscID:   "https://api.arbiscan.io/api",
		arbID:   "https://api.snowtrace.io/api",
		maticID: "https://api.bscscan.com/api",
		avaxID:  "https://api.polygonscan.com/api",
	}
	fmt.Println("latest", latestBlocks)
	scribeConfig := config.Config{
		RefreshRate: 1,
		Chains: []config.ChainConfig{
			{
				ChainID: ethID,
				ConfirmationConfig: config.ConfirmationConfig{
					ConfirmationThreshold:   10,
					ConfirmationRefreshRate: 10,
					RequiredConfirmations:   1,
				},
				GetLogsRange:         1000,
				GetLogsBatchAmount:   3,
				GetBlockBatchAmount:  10,
				ConcurrencyThreshold: 20000,
				Contracts: []config.ContractConfig{
					{
						Address:    "0x2796317b0fF8538F253012862c06787Adfb8cEb6",
						StartBlock: ethCurrentBlock - blockRange,
					},
					{
						Address:    "0x1116898DdA4015eD8dDefb84b6e8Bc24528Af2d8",
						StartBlock: ethCurrentBlock - blockRange,
					},
				},
			},
			{
				ChainID: bscID,
				ConfirmationConfig: config.ConfirmationConfig{
					ConfirmationThreshold:   10,
					ConfirmationRefreshRate: 10,
					RequiredConfirmations:   1,
				},
				GetLogsRange:         256,
				GetLogsBatchAmount:   2,
				ConcurrencyThreshold: 256,
				GetBlockBatchAmount:  10,
				Contracts: []config.ContractConfig{
					{
						Address:    "0x28ec0B36F0819ecB5005cAB836F4ED5a2eCa4D13",
						StartBlock: bscCurrentBlock - blockRange,
					},
					{
						Address:    "0x930d001b7efb225613aC7F35911c52Ac9E111Fa9",
						StartBlock: bscCurrentBlock - blockRange,
					},
				},
			},
			{
				ChainID: arbID,
				ConfirmationConfig: config.ConfirmationConfig{
					ConfirmationThreshold:   10,
					ConfirmationRefreshRate: 10,
					RequiredConfirmations:   1,
				},
				GetLogsRange:         1024,
				GetLogsBatchAmount:   2,
				ConcurrencyThreshold: 20000,
				GetBlockBatchAmount:  10,

				Contracts: []config.ContractConfig{
					{
						Address:    "0x6F4e8eBa4D337f874Ab57478AcC2Cb5BACdc19c9",
						StartBlock: arbCurrentBlock - blockRange,
					},
					{
						Address:    "0x9Dd329F5411466d9e0C488fF72519CA9fEf0cb40",
						StartBlock: arbCurrentBlock - blockRange,
					},
				},
			},
			{
				ChainID: maticID,
				ConfirmationConfig: config.ConfirmationConfig{
					ConfirmationThreshold:   10,
					ConfirmationRefreshRate: 10,
					RequiredConfirmations:   1,
				},
				GetLogsRange:         1000,
				GetLogsBatchAmount:   2,
				GetBlockBatchAmount:  10,
				ConcurrencyThreshold: 1001,
				Contracts: []config.ContractConfig{
					{
						Address:    "0x8F5BBB2BB8c2Ee94639E55d5F41de9b4839C1280",
						StartBlock: maticCurrentBlock - blockRange,
					},
					{
						Address:    "0x85fCD7Dd0a1e1A9FCD5FD886ED522dE8221C3EE5",
						StartBlock: maticCurrentBlock - blockRange,
					},
				},
			},
			{
				ChainID: avaxID,
				ConfirmationConfig: config.ConfirmationConfig{
					ConfirmationThreshold:   10,
					ConfirmationRefreshRate: 10,
					RequiredConfirmations:   1,
				},
				GetLogsRange:        256,
				GetLogsBatchAmount:  1,
				GetBlockBatchAmount: 10,

				ConcurrencyThreshold: 20000,
				Contracts: []config.ContractConfig{
					{
						Address:    "0xC05e61d0E7a63D27546389B7aD62FdFf5A91aACE",
						StartBlock: avaxCurrentBlock - blockRange,
					},
					{
						Address:    "0x77a7e60555bC18B4Be44C181b2575eee46212d44",
						StartBlock: avaxCurrentBlock - blockRange,
					},
				},
			},
		},
	}
	// Error: could not fetch logs in range 106246688 to 106247687: Block not available past max height: 82830278

	scribe, err := node.NewScribe(l.testDB, clients, scribeConfig, l.metrics)
	Nil(l.T(), err)

	killableContext, cancel := context.WithCancel(l.GetTestContext())

	go func() {
		_ = scribe.Start(killableContext)
	}()

	doneChan := make(chan bool, len(chains))

	// Launch a goroutine for each chain to check whether it has reached the destination block
	for i := range chains {
		go func(index int) {
			for {
				allContractsBackfilled := true
				chain := scribeConfig.Chains[index]
				for _, contract := range chain.Contracts {
					currentBlock, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), common.HexToAddress(contract.Address), chain.ChainID)
					if err != nil {
						fmt.Println("UOOSDHSOUDHJS FAIL", currentBlock)
					}
					Nil(l.T(), err)
					if latestBlocks[chain.ChainID] > currentBlock {
						allContractsBackfilled = false
					}
					//fmt.Println("--------", chain.ChainID, contract.Address, currentBlock, latestBlocks[chain.ChainID])
				}
				if allContractsBackfilled {
					doneChan <- true
					fmt.Println("DONE", chain.ChainID)
					return
				}
				time.Sleep(time.Second)
			}
		}(i)
	}

	for range chains {
		<-doneChan
	}
	cancel()

	logCounts := map[string]int{}
	explorerLogCounts := map[string]int{}
	for i := range chains {
		chain := scribeConfig.Chains[i]
		for _, contract := range chain.Contracts {
			key := fmt.Sprintf("%d_%s", chain.ChainID, contract.Address)
			logFilter := db.LogFilter{
				ChainID:         chains[i],
				ContractAddress: contract.Address,
			}
			fromBlock := latestBlocks[chains[i]] - blockRange
			toBlock := latestBlocks[chains[i]]
			logCount, err := getLogAmount(l.GetTestContext(), l.testDB, logFilter, fromBlock, toBlock)
			Nil(l.T(), err)
			logCounts[key] = logCount

			totalLogs, err := getLogs(l.GetTestContext(), contract.Address, fromBlock, toBlock, apiUrls[chain.ChainID])
			Nil(l.T(), err)
			fmt.Println("totalLogs", totalLogs)
			explorerLogCounts[key] = totalLogs

		}

	}
	fmt.Println(logCounts)
	fmt.Println(explorerLogCounts)

	//maticLogAmount, err := getLogAmount(l.GetTestContext(), l.testDB, logFilter, maticOriginStart, maticCurrentBlock)
	//Nil(l.T(), err)

	//fmt.Println("Matic last indexed:", maticLastIndexed)
	//fmt.Println("Avalanche last indexed:", avaxLastIndexed)
	//fmt.Println("Optimism last indexed:", opLastIndexed)
	//
	//fmt.Println("Matic log amount:", maticLogAmount)
	//fmt.Println("Avalanche log amount:", avaxLogAmount)
	//fmt.Println("Optimism log amount:", opLogAmount)

	// Use explorer API's to do a parity check on the number of logs.
}

func createHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}
}

func processBatch(ctx context.Context, client *http.Client, url string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("error getting data: %w", err)
	}
	resRaw, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resRaw.Body.Close()

	var decodedRes map[string]json.RawMessage
	if err := json.NewDecoder(resRaw.Body).Decode(&decodedRes); err != nil {
		return 0, fmt.Errorf("error decoding response: %w", err)
	}

	var resultSlice []map[string]interface{}
	if err := json.Unmarshal(decodedRes["result"], &resultSlice); err != nil {
		return 0, fmt.Errorf("error unmarshaling result: %w", err)
	}

	return len(resultSlice), nil
}

func getLogs(ctx context.Context, contractAddress string, fromBlock uint64, toBlock uint64, apiUrl string) (int, error) {
	blockRange := toBlock - fromBlock
	batchSize := uint64(500)
	numBatches := blockRange/batchSize + 1
	client := createHTTPClient()
	totalResults := 0

	for i := uint64(0); i < numBatches; i++ {
		startBlock := fromBlock + i*batchSize
		endBlock := startBlock + batchSize - 1
		if endBlock > toBlock {
			endBlock = toBlock
		}
		url := fmt.Sprintf("%s?module=logs&action=getLogs&address=%s&fromBlock=%d&toBlock=%d&page=1",
			apiUrl, contractAddress, startBlock, endBlock)
		fmt.Println("URL", url)
		b := &backoff.Backoff{
			Factor: 2,
			Jitter: true,
			Min:    10 * time.Millisecond,
			Max:    1 * time.Second,
		}
		timeout := time.Duration(0)

	RETRY:
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("context cancelled: %s", ctx.Err())
		case <-time.After(timeout):
			resultCount, err := processBatch(ctx, client, url)
			if err != nil {
				timeout = b.Duration()
				goto RETRY
			}
			totalResults += resultCount
			fmt.Println("totalResults += resultCount", totalResults, resultCount)
		}

		if i < numBatches-1 {
			time.Sleep(1 * time.Second)
		}
	}

	return totalResults, nil
}

func getLogAmount(ctx context.Context, db db.EventDB, filter db.LogFilter, startBlock uint64, endBlock uint64) (int, error) {
	page := 1
	var retrievedLogs []*types.Log
	for {
		logs, err := db.RetrieveLogsInRangeAsc(ctx, filter, startBlock, endBlock, page)
		if err != nil {
			return 0, err
		}
		retrievedLogs = append(retrievedLogs, logs...)
		if len(logs) < base.PageSize {
			break
		}
		page++
	}
	return len(retrievedLogs), nil
}

//// Get the last indexed block for each chain.
//maticLastIndexed, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), common.HexToAddress(originAddress), maticChainID)
//Nil(l.T(), err)
//avaxLastIndexed, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), common.HexToAddress(originAddress), avaxChainID)
//Nil(l.T(), err)
//opLastIndexed, err := l.testDB.RetrieveLastIndexed(l.GetTestContext(), common.HexToAddress(summitAddress), opChainID)
//Nil(l.T(), err)

// Get the number of logs stored in the scribe DB for each chain.
//logFilter := db.LogFilter{
//ChainID: maticChainID,
//}
//maticLogAmount, err := getLogAmount(l.GetTestContext(), l.testDB, logFilter, maticOriginStart, maticCurrentBlock)
//Nil(l.T(), err)
//
//logFilter = db.LogFilter{
//ChainID: avaxChainID,
//}
//avaxLogAmount, err := getLogAmount(l.GetTestContext(), l.testDB, logFilter, avaxOriginStart, avaxCurrentBlock)
//Nil(l.T(), err)
//
//logFilter = db.LogFilter{
//ChainID: opChainID,
//}
//opLogAmount, err := getLogAmount(l.GetTestContext(), l.testDB, logFilter, opSummitStart, opCurrentBlock)
//Nil(l.T(), err)
