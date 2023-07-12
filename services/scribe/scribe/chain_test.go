package scribe_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/scribe"
	"github.com/synapsecns/sanguine/services/scribe/scribe/indexer"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"
	"math/big"

	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
)

// TestContractBackfill tests using a contractBackfiller for recording receipts and logs in a database.
func (s ScribeSuite) TestIndexToBlock() {
	// Get simulated blockchain, deploy the test contract, and set up test variables.
	simulatedChain := geth.NewEmbeddedBackendForChainID(s.GetSuiteContext(), s.T(), big.NewInt(142))
	simulatedClient, err := backend.DialBackend(s.GetTestContext(), simulatedChain.RPCAddress(), s.metrics)
	Nil(s.T(), err)

	simulatedChain.FundAccount(s.GetTestContext(), s.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := s.manager.GetTestContract(s.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(s.GetTestContext(), nil)

	// Set config.
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}

	simulatedChainArr := []backend.ScribeBackend{simulatedClient, simulatedClient}
	chainConfig := config.ChainConfig{
		ChainID:              142,
		GetLogsBatchAmount:   1,
		Confirmations:        0,
		StoreConcurrency:     1,
		GetLogsRange:         1,
		ConcurrencyThreshold: 100,
		Contracts:            []config.ContractConfig{contractConfig},
	}

	chainIndexer, err := scribe.NewChainIndexer(s.testDB, simulatedChainArr, chainConfig, s.metrics)
	Nil(s.T(), err)

	// Emit events for the backfiller to read.
	tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(s.T(), err)
	simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)

	tx, err = testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(s.T(), err)

	simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)
	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
	Nil(s.T(), err)
	simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)

	// Emit two logs in one receipt.
	tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
	Nil(s.T(), err)

	simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)

	// Get the block that the last transaction was executed in.
	txBlockNumber, err := testutil.GetTxBlockNumber(s.GetTestContext(), simulatedChain, tx)
	Nil(s.T(), err)

	// TODO use no-op meter
	blockHeightMeter, err := s.metrics.Meter().NewHistogram(fmt.Sprint("scribe_block_meter", chainConfig.ChainID), "block_histogram", "a block height meter", "blocks")
	Nil(s.T(), err)

	contracts := []common.Address{common.HexToAddress(contractConfig.Address)}
	indexer, err := indexer.NewIndexer(chainConfig, contracts, s.testDB, simulatedChainArr, s.metrics, blockHeightMeter)
	Nil(s.T(), err)

	err = chainIndexer.IndexToBlock(s.GetTestContext(), nil, uint64(0), indexer)
	Nil(s.T(), err)

	// Get all receipts.
	receipts, err := s.testDB.RetrieveReceiptsWithFilter(s.GetTestContext(), db.ReceiptFilter{}, 1)
	Nil(s.T(), err)

	// Check to see if 3 receipts were collected.
	Equal(s.T(), 4, len(receipts))

	// Get all logs.
	logs, err := s.testDB.RetrieveLogsWithFilter(s.GetTestContext(), db.LogFilter{}, 1)
	Nil(s.T(), err)

	// Check to see if 4 logs were collected.
	Equal(s.T(), 5, len(logs))

	// Check to see if the last receipt has two logs.
	Equal(s.T(), 2, len(receipts[0].Logs))

	// Ensure last indexed block is correct.
	lastIndexed, err := s.testDB.RetrieveLastIndexed(s.GetTestContext(), testContract.Address(), uint32(testContract.ChainID().Uint64()))
	Nil(s.T(), err)
	Equal(s.T(), txBlockNumber, lastIndexed)
}

// TestChainIndxer tests that the ChainIndxer can backfill events from a chain.
func (s ScribeSuite) TestChainIndxer() {
	// We need to set up multiple deploy managers, one for each contract. We will use
	// s.manager for the first contract, and create a new ones for the next two.
	managerA := testutil.NewDeployManager(s.T())
	managerB := testutil.NewDeployManager(s.T())
	managerC := testutil.NewDeployManager(s.T())

	// Get simulated blockchain, deploy three test contracts, and set up test variables.
	chainID := gofakeit.Uint32()

	simulatedChain := geth.NewEmbeddedBackendForChainID(s.GetTestContext(), s.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backend.DialBackend(s.GetTestContext(), simulatedChain.RPCAddress(), s.metrics)
	Nil(s.T(), err)

	simulatedChain.FundAccount(s.GetTestContext(), s.wallet.Address(), *big.NewInt(params.Ether))
	testContractA, testRefA := s.manager.GetTestContract(s.GetTestContext(), simulatedChain)
	testContractB, testRefB := managerB.GetTestContract(s.GetTestContext(), simulatedChain)
	testContractC, testRefC := managerC.GetTestContract(s.GetTestContext(), simulatedChain)

	contracts := []contracts.DeployedContract{testContractA, testContractB, testContractC}
	testRefs := []*testcontract.TestContractRef{testRefA, testRefB, testRefC}
	startBlocks := make([]uint64, len(contracts))

	for i, contract := range contracts {
		deployTxHash := contract.DeployTx().Hash()
		receipt, err := simulatedChain.TransactionReceipt(s.GetTestContext(), deployTxHash)
		Nil(s.T(), err)
		startBlocks[i] = receipt.BlockNumber.Uint64()
	}

	contractConfigs := config.ContractConfigs{}

	for i, contract := range contracts {
		contractConfigs = append(contractConfigs, config.ContractConfig{
			Address:    contract.Address().String(),
			StartBlock: startBlocks[i],
		})
	}

	chainConfig := config.ChainConfig{
		ChainID:       chainID,
		Contracts:     contractConfigs,
		Confirmations: 1,
	}
	simulatedChainArr := []backend.ScribeBackend{simulatedClient, simulatedClient}
	chainIndexer, err := scribe.NewChainIndexer(s.testDB, simulatedChainArr, chainConfig, s.metrics)
	Nil(s.T(), err)
	testutil.EmitEventsForAChain(contracts, testRefs, simulatedChain, chainIndexer, chainConfig, true)
}

//// EmitEventsForAChain emits events for a chain. If `backfill` is set to true, the function will store the events
//// whilst checking their existence in the database.
//func (s ScribeSuite) EmitEventsForAChain(contracts []contracts.DeployedContract, testRefs []*testcontract.TestContractRef, simulatedChain backends.SimulatedTestBackend, chainBackfiller *scribe.ChainIndexer, chainConfig config.ChainConfig, backfill bool) {
//	transactOpts := simulatedChain.GetTxContext(s.GetTestContext(), nil)
//
//	// Emit events from each contract.
//	for _, testRef := range testRefs {
//		tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
//		Nil(s.T(), err)
//		simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)
//		tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
//		Nil(s.T(), err)
//		simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)
//		tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
//		Nil(s.T(), err)
//		simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)
//	}
//
//	if backfill {
//		err := chainBackfiller.Index(s.GetTestContext(), nil)
//		Nil(s.T(), err)
//
//		for _, contract := range contracts {
//			logFilter := db.LogFilter{
//				ChainID:         chainConfig.ChainID,
//				ContractAddress: contract.Address().String(),
//			}
//			logs, err := s.testDB.RetrieveLogsWithFilter(s.GetTestContext(), logFilter, 1)
//			Nil(s.T(), err)
//
//			// There should be 4 logs. One from `EmitEventA`, one from `EmitEventB`, and two from `EmitEventAandB`.
//			Equal(s.T(), 4, len(logs))
//		}
//
//		receiptFilter := db.ReceiptFilter{
//			ChainID: chainConfig.ChainID,
//		}
//		receipts, err := s.testDB.RetrieveReceiptsWithFilter(s.GetTestContext(), receiptFilter, 1)
//		Nil(s.T(), err)
//
//		// There should be 9 receipts from `EmitEventA`,`EmitEventB`, and `EmitEventAandB` for each contract.
//		Equal(s.T(), 9, len(receipts))
//		totalBlockTimes := uint64(0)
//		currBlock, err := simulatedChain.BlockNumber(s.GetTestContext())
//		Nil(s.T(), err)
//		firstBlock, err := s.testDB.RetrieveFirstBlockStored(s.GetTestContext(), chainConfig.ChainID)
//		Nil(s.T(), err)
//
//		for blockNum := firstBlock; blockNum <= currBlock; blockNum++ {
//			_, err := s.testDB.RetrieveBlockTime(s.GetTestContext(), chainConfig.ChainID, blockNum)
//			if err == nil {
//				totalBlockTimes++
//			}
//		}
//
//		// There are `currBlock` - `firstBlock`+1 block times stored (checking after contract gets deployed).
//		Equal(s.T(), currBlock-firstBlock+uint64(1), totalBlockTimes)
//	}
//}

// // TestTxTypeNotSupported tests how the contract backfiller handles a transaction type that is not supported.
// //
// // nolint:dupl
//
//	func (x IndexerSuite) TestTxTypeNotSupported() {
//		if os.Getenv("CI") != "" {
//			x.T().Skip("Network test flake")
//		}
//
//		var backendClient backend.ScribeBackend
//		omnirpcURL := "https://rpc.interoperability.institute/confirmations/1/rpc/42161"
//		backendClient, err := backend.DialBackend(x.GetTestContext(), omnirpcURL, x.metrics)
//		Nil(x.T(), err)
//
//		// This config is using this block https://arbiscan.io/block/6262099
//		// and this tx https://arbiscan.io/tx/0x8800222adf9578fb576db0bd7fb4860fe89932549be084a3313939c03e4d279d
//		// with a unique Arbitrum type to verify that anomalous tx type is handled correctly.
//		contractConfig := config.ContractConfig{
//			Address:    "0xA67b7147DcE20D6F25Fd9ABfBCB1c3cA74E11f0B",
//			StartBlock: 6262099,
//		}
//
//		chainConfig := config.ChainConfig{
//			ChainID:       42161,
//			Confirmations: 1,
//			Contracts:     []config.ContractConfig{contractConfig},
//		}
//		backendClientArr := []backend.ScribeBackend{backendClient, backendClient}
//		chainBackfiller, err := backfill.NewChainIndexer(x.testDB, backendClientArr, chainConfig, 1, x.metrics)
//		Nil(x.T(), err)
//		err = chainBackfiller.Index(x.GetTestContext(), &contractConfig.StartBlock, false)
//		Nil(x.T(), err)
//
//		logs, err := x.testDB.RetrieveLogsWithFilter(x.GetTestContext(), db.LogFilter{}, 1)
//		Nil(x.T(), err)
//		Equal(x.T(), 4, len(logs))
//		receipts, err := x.testDB.RetrieveReceiptsWithFilter(x.GetTestContext(), db.ReceiptFilter{}, 1)
//		Nil(x.T(), err)
//		Equal(x.T(), 1, len(receipts))
//	}
//
// // TestTxTypeNotSupported tests how the contract backfiller handles a transaction type that is not supported.
// //
// // nolint:dupl
//
//	func (x IndexerSuite) TestInvalidTxVRS() {
//		if os.Getenv("CI") != "" {
//			x.T().Skip("Network test flake")
//		}
//
//		var backendClient backend.ScribeBackend
//		omnirpcURL := "https://rpc.interoperability.institute/confirmations/1/rpc/1313161554"
//		backendClient, err := backend.DialBackend(x.GetTestContext(), omnirpcURL, x.metrics)
//		Nil(x.T(), err)
//
//		// This config is using this block https://aurorascan.dev/block/58621373
//		// and this tx https://aurorascan.dev/tx/0x687282d7bd6c3d591f9ad79784e0983afabcac2a9074d368b7ca3d7caf4edee5
//		// to test handling of the v,r,s tx not found error.
//		contractConfig := config.ContractConfig{
//			Address:    "0xaeD5b25BE1c3163c907a471082640450F928DDFE",
//			StartBlock: 58621373,
//		}
//
//		chainConfig := config.ChainConfig{
//			ChainID:       1313161554,
//			Confirmations: 1,
//			Contracts:     []config.ContractConfig{contractConfig},
//		}
//		backendClientArr := []backend.ScribeBackend{backendClient, backendClient}
//		chainBackfiller, err := backfill.NewChainIndexer(x.testDB, backendClientArr, chainConfig, 1, x.metrics)
//		Nil(x.T(), err)
//
//		err = chainBackfiller.Index(x.GetTestContext(), &contractConfig.StartBlock, false)
//		Nil(x.T(), err)
//
//		logs, err := x.testDB.RetrieveLogsWithFilter(x.GetTestContext(), db.LogFilter{}, 1)
//		Nil(x.T(), err)
//		Equal(x.T(), 9, len(logs))
//		receipts, err := x.testDB.RetrieveReceiptsWithFilter(x.GetTestContext(), db.ReceiptFilter{}, 1)
//		Nil(x.T(), err)
//		Equal(x.T(), 1, len(receipts))
//	}
