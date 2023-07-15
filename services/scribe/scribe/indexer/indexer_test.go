package indexer_test

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/scribe/indexer"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"os"

	"sync"

	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/mocks"

	"math/big"
)

// TestFailedStore tests that the ChainBackfiller continues backfilling after a failed store.

func (x *IndexerSuite) TestFailedStore() {
	mockDB := new(mocks.EventDB)
	mockDB.
		// on a store receipt call
		On("StoreReceipt", mock.Anything, mock.Anything, mock.Anything).
		Return(fmt.Errorf("failed to store receipt"))
	mockDB.
		// on a store transaction call
		On("StoreEthTx", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(fmt.Errorf("failed to store transaction"))
	mockDB.
		// on a store log call
		On("StoreLogs", mock.Anything, mock.Anything, mock.Anything).
		Return(fmt.Errorf("failed to store log"))
	mockDB.
		// on retrieve last indexed call
		On("RetrieveLastIndexed", mock.Anything, mock.Anything, mock.Anything).
		Return(uint64(0), nil)

	mockDB.On("StoreBlockTime", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	chainID := gofakeit.Uint32()

	simulatedChain := geth.NewEmbeddedBackendForChainID(x.GetTestContext(), x.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backend.DialBackend(x.GetTestContext(), simulatedChain.RPCAddress(), x.metrics)
	Nil(x.T(), err)

	simulatedChain.FundAccount(x.GetTestContext(), x.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := x.manager.GetTestContract(x.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(x.GetTestContext(), nil)
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}
	simulatedChainArr := []backend.ScribeBackend{simulatedClient, simulatedClient}
	chainConfig := config.ChainConfig{
		Confirmations:      1,
		ChainID:            chainID,
		GetLogsBatchAmount: 1,
		StoreConcurrency:   1,
		GetLogsRange:       1,
		Contracts:          []config.ContractConfig{contractConfig},
	}
	blockHeightMeter, err := x.metrics.Meter().NewHistogram(fmt.Sprint("scribe_block_meter", chainConfig.ChainID), "block_histogram", "a block height meter", "blocks")
	Nil(x.T(), err)

	contracts := []common.Address{common.HexToAddress(contractConfig.Address)}
	indexer, err := indexer.NewIndexer(chainConfig, contracts, mockDB, simulatedChainArr, x.metrics, blockHeightMeter)
	Nil(x.T(), err)

	tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	// Get the block that the last transaction was executed in.
	txBlockNumber, err := testutil.GetTxBlockNumber(x.GetTestContext(), simulatedChain, tx)
	Nil(x.T(), err)
	err = indexer.Index(x.GetTestContext(), contractConfig.StartBlock, txBlockNumber)
	NotNil(x.T(), err)

	// Check to ensure that StoreLastIndexed was never called.
	mockDB.AssertNotCalled(x.T(), "StoreLastIndexed", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

// TestGetLogsSimulated tests the GetLogs function using a simulated blockchain.
//
//nolint:cyclop
func (x *IndexerSuite) TestGetLogsSimulated() {
	// Get simulated blockchain, deploy the test contract, and set up test variables.
	simulatedChain := geth.NewEmbeddedBackendForChainID(x.GetSuiteContext(), x.T(), big.NewInt(3))
	simulatedClient, err := backend.DialBackend(x.GetTestContext(), simulatedChain.RPCAddress(), x.metrics)
	Nil(x.T(), err)

	simulatedChain.FundAccount(x.GetTestContext(), x.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := x.manager.GetTestContract(x.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(x.GetTestContext(), nil)
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}
	simulatedChainArr := []backend.ScribeBackend{simulatedClient, simulatedClient}
	chainConfig := config.ChainConfig{
		Confirmations:      1,
		ChainID:            3,
		GetLogsBatchAmount: 1,
		StoreConcurrency:   1,
		GetLogsRange:       1,
		Contracts:          []config.ContractConfig{contractConfig},
	}
	blockHeightMeter, err := x.metrics.Meter().NewHistogram(fmt.Sprint("scribe_block_meter", chainConfig.ChainID), "block_histogram", "a block height meter", "blocks")
	Nil(x.T(), err)

	contracts := []common.Address{common.HexToAddress(contractConfig.Address)}
	contractIndexer, err := indexer.NewIndexer(chainConfig, contracts, x.testDB, simulatedChainArr, x.metrics, blockHeightMeter)
	Nil(x.T(), err)

	// Emit five events, and then fetch them with GetLogs. The first two will be fetched first,
	// then the last three after.
	tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)
	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	// Get the block that the second transaction was executed in.
	txBlockNumberA, err := testutil.GetTxBlockNumber(x.GetTestContext(), simulatedChain, tx)
	Nil(x.T(), err)

	tx, err = testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)
	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{10}, big.NewInt(11), big.NewInt(12))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)
	tx, err = testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(13), big.NewInt(14), big.NewInt(15))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	// Get the block that the last transaction was executed in.
	txBlockNumberB, err := testutil.GetTxBlockNumber(x.GetTestContext(), simulatedChain, tx)
	Nil(x.T(), err)

	// Get the logs for the first two events.
	collectedLogs := []types.Log{}
	logs, errChan := contractIndexer.GetLogs(x.GetTestContext(), contractConfig.StartBlock, txBlockNumberA)

	for {
		select {
		case <-x.GetTestContext().Done():
			x.T().Error("test timed out")
		case log, ok := <-logs:
			if !ok {
				goto Done
			}
			collectedLogs = append(collectedLogs, log)
		case errorFromChan := <-errChan:
			Nil(x.T(), errorFromChan)
		}
	}
Done:
	// Check to see if 2 logs were collected.
	Equal(x.T(), 2, len(collectedLogs))

	// Get the logs for the last three events.
	collectedLogs = []types.Log{}
	logs, errChan = contractIndexer.GetLogs(x.GetTestContext(), txBlockNumberA+1, txBlockNumberB)

	for {
		select {
		case <-x.GetTestContext().Done():
			x.T().Error("test timed out")
		case log, ok := <-logs:
			if !ok {
				goto Done2
			}
			collectedLogs = append(collectedLogs, log)
		case errorFromChan := <-errChan:
			Nil(x.T(), errorFromChan)
		}
	}
Done2:

	// Check to see if 3 logs were collected.
	Equal(x.T(), 3, len(collectedLogs))
}

// TestContractBackfill tests using a contractBackfiller for recording receipts and logs in a database.
func (x *IndexerSuite) TestContractBackfill() {
	// Get simulated blockchain, deploy the test contract, and set up test variables.
	simulatedChain := geth.NewEmbeddedBackendForChainID(x.GetSuiteContext(), x.T(), big.NewInt(142))
	simulatedClient, err := backend.DialBackend(x.GetTestContext(), simulatedChain.RPCAddress(), x.metrics)
	Nil(x.T(), err)

	simulatedChain.FundAccount(x.GetTestContext(), x.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := x.manager.GetTestContract(x.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(x.GetTestContext(), nil)

	// Set config.
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}

	simulatedChainArr := []backend.ScribeBackend{simulatedClient, simulatedClient}
	chainConfig := config.ChainConfig{
		ChainID:              142,
		GetLogsBatchAmount:   1,
		Confirmations:        1,
		StoreConcurrency:     1,
		GetLogsRange:         1,
		ConcurrencyThreshold: 100,
		Contracts:            []config.ContractConfig{contractConfig},
	}
	blockHeightMeter, err := x.metrics.Meter().NewHistogram(fmt.Sprint("scribe_block_meter", chainConfig.ChainID), "block_histogram", "a block height meter", "blocks")
	Nil(x.T(), err)
	contracts := []common.Address{common.HexToAddress(contractConfig.Address)}
	contractIndexer, err := indexer.NewIndexer(chainConfig, contracts,
		x.testDB, simulatedChainArr, x.metrics, blockHeightMeter)
	x.Require().NoError(err)

	// Emit events for the backfiller to read.
	tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	tx, err = testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(x.T(), err)

	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)
	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	// Emit two logs in one receipt.
	tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
	Nil(x.T(), err)

	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	// Get the block that the last transaction was executed in.
	txBlockNumber, err := testutil.GetTxBlockNumber(x.GetTestContext(), simulatedChain, tx)
	Nil(x.T(), err)

	// Backfill the events. The `0` will be replaced with the startBlock from the config.
	err = contractIndexer.Index(x.GetTestContext(), contractConfig.StartBlock, txBlockNumber)
	Nil(x.T(), err)

	// Get all receipts.
	receipts, err := x.testDB.RetrieveReceiptsWithFilter(x.GetTestContext(), db.ReceiptFilter{}, 1)
	Nil(x.T(), err)

	// Check to see if 3 receipts were collected.
	Equal(x.T(), 4, len(receipts))

	// Get all logs.
	logs, err := x.testDB.RetrieveLogsWithFilter(x.GetTestContext(), db.LogFilter{}, 1)
	Nil(x.T(), err)

	// Check to see if 4 logs were collected.
	Equal(x.T(), 5, len(logs))

	// Check to see if the last receipt has two logs.
	Equal(x.T(), 2, len(receipts[0].Logs))

	// Ensure last indexed block is correct.
	lastIndexed, err := x.testDB.RetrieveLastIndexed(x.GetTestContext(), testContract.Address(), uint32(testContract.ChainID().Uint64()))
	Nil(x.T(), err)
	Equal(x.T(), txBlockNumber, lastIndexed)
}

// TestContractBackfill tests using a contractBackfiller for recording receipts and logs in a database.
func (x *IndexerSuite) TestContractBackfillFromPreIndexed() {
	// Get simulated blockchain, deploy the test contract, and set up test variables.
	simulatedChain := geth.NewEmbeddedBackendForChainID(x.GetSuiteContext(), x.T(), big.NewInt(142))
	simulatedClient, err := backend.DialBackend(x.GetTestContext(), simulatedChain.RPCAddress(), x.metrics)
	Nil(x.T(), err)

	simulatedChain.FundAccount(x.GetTestContext(), x.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := x.manager.GetTestContract(x.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(x.GetTestContext(), nil)

	// Set config.
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}

	simulatedChainArr := []backend.ScribeBackend{simulatedClient, simulatedClient}
	chainConfig := config.ChainConfig{
		ChainID:              142,
		GetLogsBatchAmount:   1,
		StoreConcurrency:     1,
		Confirmations:        1,
		GetLogsRange:         1,
		ConcurrencyThreshold: 1,
		Contracts:            []config.ContractConfig{contractConfig},
	}
	blockHeightMeter, err := x.metrics.Meter().NewHistogram(fmt.Sprint("scribe_block_meter", chainConfig.ChainID), "block_histogram", "a block height meter", "blocks")
	Nil(x.T(), err)

	contracts := []common.Address{common.HexToAddress(contractConfig.Address)}
	backfiller, err := indexer.NewIndexer(chainConfig, contracts, x.testDB, simulatedChainArr, x.metrics, blockHeightMeter)
	Nil(x.T(), err)

	// 1 log 1 receipt: r:1 l:1
	tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	// 1 log 1 receipt: r:2 l:2
	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	// 2 logs 1 receipt: r:3 l:4
	tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	// Get the block that the last transaction was executed in.
	txBlockNumber, err := testutil.GetTxBlockNumber(x.GetTestContext(), simulatedChain, tx)
	Nil(x.T(), err)
	err = x.testDB.StoreLastIndexed(x.GetTestContext(), common.HexToAddress(contractConfig.Address), chainConfig.ChainID, txBlockNumber)
	Nil(x.T(), err)

	// 1 log 1 receipt: r:4 l:5
	tx, err = testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(10), big.NewInt(11), big.NewInt(12))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	// 1 log 1 receipt: r:5 l:6
	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{13}, big.NewInt(14), big.NewInt(15))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	// 2 logs 1 receipt: r:6 l:8
	tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(16), big.NewInt(17), big.NewInt(18))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	// 2 logs 1 receipt: r:7 l:10
	tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(19), big.NewInt(20), big.NewInt(21))
	Nil(x.T(), err)
	simulatedChain.WaitForConfirmation(x.GetTestContext(), tx)

	// Get the block that the last transaction was executed in.
	txBlockNumber, err = testutil.GetTxBlockNumber(x.GetTestContext(), simulatedChain, tx)
	Nil(x.T(), err)

	err = backfiller.Index(x.GetTestContext(), contractConfig.StartBlock, txBlockNumber)
	Nil(x.T(), err)

	// Get all receipts.
	receipts, err := x.testDB.RetrieveReceiptsWithFilter(x.GetTestContext(), db.ReceiptFilter{}, 1)
	Nil(x.T(), err)
	Equal(x.T(), 7, len(receipts))

	// Get all logs.
	logs, err := x.testDB.RetrieveLogsWithFilter(x.GetTestContext(), db.LogFilter{}, 1)
	Nil(x.T(), err)

	Equal(x.T(), 10, len(logs))

	// Check to see if the last receipt has two logs (emit a and b).
	Equal(x.T(), 2, len(receipts[0].Logs))

	// Ensure last indexed block is correct.
	lastIndexed, err := x.testDB.RetrieveLastIndexed(x.GetTestContext(), testContract.Address(), uint32(testContract.ChainID().Uint64()))
	Nil(x.T(), err)
	Equal(x.T(), txBlockNumber, lastIndexed)
}

func (x *IndexerSuite) TestGetLogs() {
	const desiredBlockHeight = 10

	var addresses []common.Address
	var err error
	var wg sync.WaitGroup

	wg.Add(2)
	testBackend := geth.NewEmbeddedBackend(x.GetTestContext(), x.T())

	go func() {
		defer wg.Done()
		addresses, _, err = testutil.PopulateWithLogs(x.GetTestContext(), x.T(), testBackend, desiredBlockHeight, []*testutil.DeployManager{x.manager})
		Nil(x.T(), err)
	}()

	var host string
	go func() {
		defer wg.Done()
		host = testutil.StartOmnirpcServer(x.GetTestContext(), x.T(), testBackend)
	}()

	wg.Wait()

	scribeBackend, err := backend.DialBackend(x.GetTestContext(), host, x.metrics)
	Nil(x.T(), err)
	simulatedChainArr := []backend.ScribeBackend{scribeBackend, scribeBackend}

	chainID, err := scribeBackend.ChainID(x.GetTestContext())
	Nil(x.T(), err)

	var contractConfigs []config.ContractConfig
	for _, address := range addresses {
		contractConfig := config.ContractConfig{
			Address: address.String(),
		}
		contractConfigs = append(contractConfigs, contractConfig)
	}

	chainConfig := config.ChainConfig{
		ChainID:            uint32(chainID.Uint64()),
		Confirmations:      1,
		GetLogsBatchAmount: 1,
		StoreConcurrency:   1,
		GetLogsRange:       1,
		Contracts:          contractConfigs,
	}
	blockHeightMeter, err := x.metrics.Meter().NewHistogram(fmt.Sprint("scribe_block_meter", chainConfig.ChainID), "block_histogram", "a block height meter", "blocks")
	Nil(x.T(), err)

	contractBackfiller, err := indexer.NewIndexer(chainConfig, addresses, x.testDB, simulatedChainArr, x.metrics, blockHeightMeter)
	Nil(x.T(), err)

	startHeight, endHeight := uint64(1), uint64(10)
	logsChan, errChan := contractBackfiller.GetLogs(x.GetTestContext(), startHeight, endHeight)

	var logs []types.Log
	var errs []string
loop:
	for {
		select {
		case log, ok := <-logsChan:
			if !ok {
				break loop
			}
			logs = append(logs, log)
		case err, ok := <-errChan:
			if !ok {
				break loop
			}
			errs = append(errs, err)
		}
	}

	Equal(x.T(), 2, len(logs))
	Equal(x.T(), 0, len(errs))

	cancelCtx, cancel := context.WithCancel(x.GetTestContext())
	cancel()

	_, errChan = contractBackfiller.GetLogs(cancelCtx, startHeight, endHeight)
loop2:
	for {
		errStr := <-errChan
		Contains(x.T(), errStr, "context canceled")
		break loop2
	}
}

// TestTxTypeNotSupported tests how the contract backfiller handles a transaction type that is not supported.
//
// nolint:dupl
func (x *IndexerSuite) TestTxTypeNotSupported() {
	if os.Getenv("CI") != "" {
		x.T().Skip("Network test flake")
	}

	var backendClient backend.ScribeBackend
	omnirpcURL := "https://rpc.interoperability.institute/confirmations/1/rpc/42161"
	backendClient, err := backend.DialBackend(x.GetTestContext(), omnirpcURL, x.metrics)
	Nil(x.T(), err)

	// This config is using this block https://arbiscan.io/block/6262099
	// and this tx https://arbiscan.io/tx/0x8800222adf9578fb576db0bd7fb4860fe89932549be084a3313939c03e4d279d
	// with a unique Arbitrum type to verify that anomalous tx type is handled correctly.
	contractConfig := config.ContractConfig{
		Address:    "0xA67b7147DcE20D6F25Fd9ABfBCB1c3cA74E11f0B",
		StartBlock: 6262099,
	}

	chainConfig := config.ChainConfig{
		ChainID:            42161,
		Confirmations:      1,
		GetLogsRange:       1,
		GetLogsBatchAmount: 1,
		Contracts:          []config.ContractConfig{contractConfig},
	}

	addresses := []common.Address{common.HexToAddress(contractConfig.Address)}
	backendClientArr := []backend.ScribeBackend{backendClient, backendClient}
	blockHeightMeter, err := x.metrics.Meter().NewHistogram(fmt.Sprint("scribe_block_meter", chainConfig.ChainID), "block_histogram", "a block height meter", "blocks")
	Nil(x.T(), err)

	contractIndexer, err := indexer.NewIndexer(chainConfig, addresses, x.testDB, backendClientArr, x.metrics, blockHeightMeter)
	Nil(x.T(), err)

	err = contractIndexer.Index(x.GetTestContext(), contractConfig.StartBlock, contractConfig.StartBlock+1)
	Nil(x.T(), err)

	logs, err := x.testDB.RetrieveLogsWithFilter(x.GetTestContext(), db.LogFilter{}, 1)
	Nil(x.T(), err)
	Equal(x.T(), 4, len(logs))
	receipts, err := x.testDB.RetrieveReceiptsWithFilter(x.GetTestContext(), db.ReceiptFilter{}, 1)
	Nil(x.T(), err)
	Equal(x.T(), 1, len(receipts))
}

// TestTxTypeNotSupported tests how the contract indexerer handles a transaction type that is not supported.
//
// nolint:dupl
func (x IndexerSuite) TestInvalidTxVRS() {
	if os.Getenv("CI") != "" {
		x.T().Skip("Network test flake")
	}

	var backendClient backend.ScribeBackend
	omnirpcURL := "https://rpc.interoperability.institute/confirmations/1/rpc/1313161554"
	backendClient, err := backend.DialBackend(x.GetTestContext(), omnirpcURL, x.metrics)
	Nil(x.T(), err)

	// This config is using this block https://aurorascan.dev/block/58621373
	// and this tx https://aurorascan.dev/tx/0x687282d7bd6c3d591f9ad79784e0983afabcac2a9074d368b7ca3d7caf4edee5
	// to test handling of the v,r,s tx not found error.
	contractConfig := config.ContractConfig{
		Address:    "0xaeD5b25BE1c3163c907a471082640450F928DDFE",
		StartBlock: 58621373,
	}

	chainConfig := config.ChainConfig{
		ChainID:            1313161554,
		Confirmations:      1,
		GetLogsRange:       1,
		GetLogsBatchAmount: 1,
		Contracts:          []config.ContractConfig{contractConfig},
	}
	addresses := []common.Address{common.HexToAddress(contractConfig.Address)}

	backendClientArr := []backend.ScribeBackend{backendClient, backendClient}
	blockHeightMeter, err := x.metrics.Meter().NewHistogram(fmt.Sprint("scribe_block_meter", chainConfig.ChainID), "block_histogram", "a block height meter", "blocks")
	Nil(x.T(), err)

	contractIndexer, err := indexer.NewIndexer(chainConfig, addresses, x.testDB, backendClientArr, x.metrics, blockHeightMeter)
	Nil(x.T(), err)

	err = contractIndexer.Index(x.GetTestContext(), contractConfig.StartBlock, contractConfig.StartBlock+1)
	Nil(x.T(), err)

	logs, err := x.testDB.RetrieveLogsWithFilter(x.GetTestContext(), db.LogFilter{}, 1)
	Nil(x.T(), err)
	Equal(x.T(), 9, len(logs))
	receipts, err := x.testDB.RetrieveReceiptsWithFilter(x.GetTestContext(), db.ReceiptFilter{}, 1)
	Nil(x.T(), err)
	Equal(x.T(), 1, len(receipts))
}
