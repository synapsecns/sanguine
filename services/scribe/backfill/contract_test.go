package backfill_test

import (
	"fmt"
	"math/big"

	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
)

// TestGetLogsSimulated tests the GetLogs function using a simulated blockchain.
//
//nolint:cyclop
func (b BackfillSuite) TestGetLogsSimulated() {
	// Get simulated blockchain, deploy the test contract, and set up test variables.
	simulatedChain := simulated.NewSimulatedBackendWithChainID(b.GetSuiteContext(), b.T(), big.NewInt(3))
	simulatedChain.FundAccount(b.GetTestContext(), b.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(b.GetTestContext(), nil)

	// Set config.
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}

	backfiller, err := backfill.NewContractBackfiller(3, contractConfig.Address, b.testDB, simulatedChain)
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
	simulatedChain := simulated.NewSimulatedBackendWithChainID(b.GetSuiteContext(), b.T(), big.NewInt(142))
	simulatedChain.FundAccount(b.GetTestContext(), b.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(b.GetTestContext(), nil)

	// Set config.
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}

	backfiller, err := backfill.NewContractBackfiller(142, contractConfig.Address, b.testDB, simulatedChain)
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
	// Backfill the events. The `0` will be replaced with the startBlock from the config.
	err = backfiller.Backfill(b.GetTestContext(), contractConfig.StartBlock, txBlockNumber)
	Nil(b.T(), err)
	// Get all receipts.
	receipts, err := b.testDB.UnsafeRetrieveAllReceipts(b.GetTestContext(), false, 0)
	Nil(b.T(), err)
	// Check to see if 3 receipts were collected.
	Equal(b.T(), 3, len(receipts))
	// Get all logs.
	logs, err := b.testDB.UnsafeRetrieveAllLogs(b.GetTestContext(), false, 0, common.Address{})
	Nil(b.T(), err)
	// Check to see if 4 logs were collected.
	Equal(b.T(), 4, len(logs))
	// Check to see if the last receipt has two logs.
	Equal(b.T(), 2, len(receipts[2].Logs))
	// Ensure last indexed block is correct.
	lastIndexed, err := b.testDB.RetrieveLastIndexed(b.GetTestContext(), testContract.Address(), uint32(testContract.ChainID().Uint64()))
	Nil(b.T(), err)
	Equal(b.T(), txBlockNumber, lastIndexed)
}

// TestGetLogsMock tests the GetLogs function using a mocked blockchain for errors.
func (b BackfillSuite) TestGetLogsMock() {
	// TODO: do this with mocks for error handling in GetLogs methods
}

func (b BackfillSuite) getTxBlockNumber(chain *simulated.Backend, tx *types.Transaction) (uint64, error) {
	receipt, err := chain.TransactionReceipt(b.GetTestContext(), tx.Hash())
	if err != nil {
		return 0, fmt.Errorf("error getting receipt for tx: %w", err)
	}
	return receipt.BlockNumber.Uint64(), nil
}
