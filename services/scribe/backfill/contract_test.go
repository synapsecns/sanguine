package backfill_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/mocks"
	"math/big"
	"os"
)

// TestFailedStore tests that the ChainBackfiller continues backfilling after a failed store.

func (b BackfillSuite) TestFailedStore() {
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

	simulatedChain := geth.NewEmbeddedBackendForChainID(b.GetTestContext(), b.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backfill.DialBackend(b.GetTestContext(), simulatedChain.RPCAddress(), b.metrics)
	Nil(b.T(), err)

	simulatedChain.FundAccount(b.GetTestContext(), b.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(b.GetTestContext(), nil)
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}
	simulatedChainArr := []backfill.ScribeBackend{simulatedClient, simulatedClient}
	chainConfig := config.ChainConfig{
		ChainID:              chainID,
		ContractChunkSize:    1,
		ContractSubChunkSize: 1,
	}
	backfiller, err := backfill.NewContractBackfiller(chainConfig, contractConfig, mockDB, simulatedChainArr, b.metrics)
	Nil(b.T(), err)

	tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Get the block that the last transaction was executed in.
	txBlockNumber, err := b.getTxBlockNumber(simulatedChain, tx)
	Nil(b.T(), err)
	err = backfiller.Backfill(b.GetTestContext(), contractConfig.StartBlock, txBlockNumber)
	NotNil(b.T(), err)

	// Check to ensure that StoreLastIndexed was never called.
	mockDB.AssertNotCalled(b.T(), "StoreLastIndexed", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

// TestGetLogsSimulated tests the GetLogs function using a simulated blockchain.
//
//nolint:cyclop
func (b BackfillSuite) TestGetLogsSimulated() {
	// Get simulated blockchain, deploy the test contract, and set up test variables.
	simulatedChain := geth.NewEmbeddedBackendForChainID(b.GetSuiteContext(), b.T(), big.NewInt(3))
	simulatedClient, err := backfill.DialBackend(b.GetTestContext(), simulatedChain.RPCAddress(), b.metrics)
	Nil(b.T(), err)

	simulatedChain.FundAccount(b.GetTestContext(), b.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(b.GetTestContext(), nil)
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}
	simulatedChainArr := []backfill.ScribeBackend{simulatedClient, simulatedClient}
	chainConfig := config.ChainConfig{
		ChainID:              3,
		ContractChunkSize:    1,
		ContractSubChunkSize: 1,
	}

	backfiller, err := backfill.NewContractBackfiller(chainConfig, contractConfig, b.testDB, simulatedChainArr, b.metrics)
	Nil(b.T(), err)

	// Emit five events, and then fetch them with GetLogs. The first two will be fetched first,
	// then the last three after.
	tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Get the block that the second transaction was executed in.
	txBlockNumberA, err := b.getTxBlockNumber(simulatedChain, tx)
	Nil(b.T(), err)

	tx, err = testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{10}, big.NewInt(11), big.NewInt(12))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	tx, err = testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(13), big.NewInt(14), big.NewInt(15))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Get the block that the last transaction was executed in.
	txBlockNumberB, err := b.getTxBlockNumber(simulatedChain, tx)
	Nil(b.T(), err)

	// Get the logs for the first two events.
	collectedLogs := []types.Log{}
	logs, done := backfiller.GetLogs(b.GetTestContext(), contractConfig.StartBlock, txBlockNumberA)

	for {
		select {
		case <-b.GetTestContext().Done():
			b.T().Error("test timed out")
		case log := <-logs:
			collectedLogs = append(collectedLogs, log)
		case <-done:
			goto Next
		}
	}

Next:

	// Check to see if 2 logs were collected.
	Equal(b.T(), 2, len(collectedLogs))

	// Get the logs for the last three events.
	collectedLogs = []types.Log{}
	logs, done = backfiller.GetLogs(b.GetTestContext(), txBlockNumberA+1, txBlockNumberB)

	for {
		select {
		case <-b.GetTestContext().Done():
			b.T().Error("test timed out")
		case log := <-logs:
			collectedLogs = append(collectedLogs, log)
		case <-done:
			goto Done
		}
	}

Done:

	// Check to see if 3 logs were collected.
	Equal(b.T(), 3, len(collectedLogs))
}

// TestContractBackfill tests using a contractBackfiller for recording receipts and logs in a database.

func (b BackfillSuite) TestContractBackfill() {
	// Get simulated blockchain, deploy the test contract, and set up test variables.
	simulatedChain := geth.NewEmbeddedBackendForChainID(b.GetSuiteContext(), b.T(), big.NewInt(142))
	simulatedClient, err := backfill.DialBackend(b.GetTestContext(), simulatedChain.RPCAddress(), b.metrics)
	Nil(b.T(), err)

	simulatedChain.FundAccount(b.GetTestContext(), b.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(b.GetTestContext(), nil)

	// Set config.
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}

	simulatedChainArr := []backfill.ScribeBackend{simulatedClient, simulatedClient}
	chainConfig := config.ChainConfig{
		ChainID:              142,
		ContractChunkSize:    1,
		ContractSubChunkSize: 1,
	}
	backfiller, err := backfill.NewContractBackfiller(chainConfig, contractConfig, b.testDB, simulatedChainArr, b.metrics)
	b.Require().NoError(err)

	// Emit events for the backfiller to read.
	tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	tx, err = testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(b.T(), err)

	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Emit two logs in one receipt.
	tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
	Nil(b.T(), err)

	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Get the block that the last transaction was executed in.
	txBlockNumber, err := b.getTxBlockNumber(simulatedChain, tx)
	Nil(b.T(), err)

	// Backfill the events. The `0` will be replaced with the startBlock from the config.
	err = backfiller.Backfill(b.GetTestContext(), contractConfig.StartBlock, txBlockNumber)
	Nil(b.T(), err)

	// Get all receipts.
	receipts, err := b.testDB.RetrieveReceiptsWithFilter(b.GetTestContext(), db.ReceiptFilter{}, 1)
	Nil(b.T(), err)

	// Check to see if 3 receipts were collected.
	Equal(b.T(), 4, len(receipts))

	// Get all logs.
	logs, err := b.testDB.RetrieveLogsWithFilter(b.GetTestContext(), db.LogFilter{}, 1)
	Nil(b.T(), err)

	// Check to see if 4 logs were collected.
	Equal(b.T(), 5, len(logs))

	// Check to see if the last receipt has two logs.
	Equal(b.T(), 2, len(receipts[0].Logs))

	// Ensure last indexed block is correct.
	lastIndexed, err := b.testDB.RetrieveLastIndexed(b.GetTestContext(), testContract.Address(), uint32(testContract.ChainID().Uint64()))
	Nil(b.T(), err)
	Equal(b.T(), txBlockNumber, lastIndexed)
}

// TestTxTypeNotSupported tests how the contract backfiller handles a transaction type that is not supported.
//
// nolint:dupl
func (b BackfillSuite) TestTxTypeNotSupported() {
	if os.Getenv("CI") != "" {
		b.T().Skip("Network test flake")
	}

	var backendClient backfill.ScribeBackend
	omnirpcURL := "https://rpc.interoperability.institute/confirmations/1/rpc/42161"
	backendClient, err := backfill.DialBackend(b.GetTestContext(), omnirpcURL, b.metrics)
	Nil(b.T(), err)

	// This config is using this block https://arbiscan.io/block/6262099
	// and this tx https://arbiscan.io/tx/0x8800222adf9578fb576db0bd7fb4860fe89932549be084a3313939c03e4d279d
	// with a unique Arbitrum type to verify that anomalous tx type is handled correctly.
	contractConfig := config.ContractConfig{
		Address:    "0xA67b7147DcE20D6F25Fd9ABfBCB1c3cA74E11f0B",
		StartBlock: 6262099,
	}

	chainConfig := config.ChainConfig{
		ChainID:               42161,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{contractConfig},
	}
	backendClientArr := []backfill.ScribeBackend{backendClient, backendClient}
	chainBackfiller, err := backfill.NewChainBackfiller(b.testDB, backendClientArr, chainConfig, 1, b.metrics)
	Nil(b.T(), err)
	err = chainBackfiller.Backfill(b.GetTestContext(), &contractConfig.StartBlock, false)
	Nil(b.T(), err)

	logs, err := b.testDB.RetrieveLogsWithFilter(b.GetTestContext(), db.LogFilter{}, 1)
	Nil(b.T(), err)
	Equal(b.T(), 4, len(logs))
	receipts, err := b.testDB.RetrieveReceiptsWithFilter(b.GetTestContext(), db.ReceiptFilter{}, 1)
	Nil(b.T(), err)
	Equal(b.T(), 1, len(receipts))
}

// TestTxTypeNotSupported tests how the contract backfiller handles a transaction type that is not supported.
//
// nolint:dupl
func (b BackfillSuite) TestInvalidTxVRS() {
	if os.Getenv("CI") != "" {
		b.T().Skip("Network test flake")
	}

	var backendClient backfill.ScribeBackend
	omnirpcURL := "https://rpc.interoperability.institute/confirmations/1/rpc/1313161554"
	backendClient, err := backfill.DialBackend(b.GetTestContext(), omnirpcURL, b.metrics)
	Nil(b.T(), err)

	// This config is using this block https://aurorascan.dev/block/58621373
	// and this tx https://aurorascan.dev/tx/0x687282d7bd6c3d591f9ad79784e0983afabcac2a9074d368b7ca3d7caf4edee5
	// to test handling of the v,r,s tx not found error.
	contractConfig := config.ContractConfig{
		Address:    "0xaeD5b25BE1c3163c907a471082640450F928DDFE",
		StartBlock: 58621373,
	}

	chainConfig := config.ChainConfig{
		ChainID:               1313161554,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{contractConfig},
	}
	backendClientArr := []backfill.ScribeBackend{backendClient, backendClient}
	chainBackfiller, err := backfill.NewChainBackfiller(b.testDB, backendClientArr, chainConfig, 1, b.metrics)
	Nil(b.T(), err)

	err = chainBackfiller.Backfill(b.GetTestContext(), &contractConfig.StartBlock, false)
	Nil(b.T(), err)

	logs, err := b.testDB.RetrieveLogsWithFilter(b.GetTestContext(), db.LogFilter{}, 1)
	Nil(b.T(), err)
	Equal(b.T(), 9, len(logs))
	receipts, err := b.testDB.RetrieveReceiptsWithFilter(b.GetTestContext(), db.ReceiptFilter{}, 1)
	Nil(b.T(), err)
	Equal(b.T(), 1, len(receipts))
}
func (b BackfillSuite) getTxBlockNumber(chain backends.SimulatedTestBackend, tx *types.Transaction) (uint64, error) {
	receipt, err := chain.TransactionReceipt(b.GetTestContext(), tx.Hash())
	if err != nil {
		return 0, fmt.Errorf("error getting receipt for tx: %w", err)
	}
	return receipt.BlockNumber.Uint64(), nil
}

// TestContractBackfill tests using a contractBackfiller for recording receipts and logs in a database.
func (b BackfillSuite) TestContractBackfillFromPreIndexed() {
	// Get simulated blockchain, deploy the test contract, and set up test variables.
	simulatedChain := geth.NewEmbeddedBackendForChainID(b.GetSuiteContext(), b.T(), big.NewInt(142))
	simulatedClient, err := backfill.DialBackend(b.GetTestContext(), simulatedChain.RPCAddress(), b.metrics)
	Nil(b.T(), err)

	simulatedChain.FundAccount(b.GetTestContext(), b.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(b.GetTestContext(), nil)

	// Set config.
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}

	simulatedChainArr := []backfill.ScribeBackend{simulatedClient, simulatedClient}
	chainConfig := config.ChainConfig{
		ChainID:              142,
		ContractChunkSize:    1,
		ContractSubChunkSize: 1,
	}
	backfiller, err := backfill.NewContractBackfiller(chainConfig, contractConfig, b.testDB, simulatedChainArr, b.metrics)
	Nil(b.T(), err)

	// Emit events for the backfiller to read.
	tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(b.T(), err)

	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Emit two logs in one receipt.
	tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
	Nil(b.T(), err)

	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Get the block that the last transaction was executed in.
	txBlockNumber, err := b.getTxBlockNumber(simulatedChain, tx)
	Nil(b.T(), err)

	err = b.testDB.StoreLastIndexed(b.GetTestContext(), common.HexToAddress(contractConfig.Address), chainConfig.ChainID, txBlockNumber)
	Nil(b.T(), err)

	tx, err = testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(10), big.NewInt(11), big.NewInt(12))
	Nil(b.T(), err)

	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{13}, big.NewInt(14), big.NewInt(15))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Emit two logs in one receipt.
	tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(16), big.NewInt(17), big.NewInt(18))
	Nil(b.T(), err)

	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Emit two logs in one receipt.
	tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(19), big.NewInt(20), big.NewInt(21))
	Nil(b.T(), err)

	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Get the block that the last transaction was executed in.
	txBlockNumber, err = b.getTxBlockNumber(simulatedChain, tx)
	Nil(b.T(), err)

	err = backfiller.Backfill(b.GetTestContext(), contractConfig.StartBlock, txBlockNumber)
	Nil(b.T(), err)

	// Get all receipts.
	receipts, err := b.testDB.RetrieveReceiptsWithFilter(b.GetTestContext(), db.ReceiptFilter{}, 1)
	Nil(b.T(), err)

	// Check to see if 3 receipts were collected.
	Equal(b.T(), 4, len(receipts))

	// Get all logs.
	logs, err := b.testDB.RetrieveLogsWithFilter(b.GetTestContext(), db.LogFilter{}, 1)
	Nil(b.T(), err)

	// Check to see if 4 logs were collected.
	Equal(b.T(), 6, len(logs))

	// Check to see if the last receipt has two logs.
	Equal(b.T(), 2, len(receipts[0].Logs))

	// Ensure last indexed block is correct.
	lastIndexed, err := b.testDB.RetrieveLastIndexed(b.GetTestContext(), testContract.Address(), uint32(testContract.ChainID().Uint64()))
	Nil(b.T(), err)
	Equal(b.T(), txBlockNumber, lastIndexed)
}
