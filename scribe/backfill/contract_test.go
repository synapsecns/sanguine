package backfill_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/scribe/backfill"
)

// TestStartHeightForBackfill ensures the start height for backfill is calculated correctly.
func (b BackfillSuite) TestStartHeightForBackfill() {

	// Get simulated blockchain and deploy the test contract.
	simulatedChain := simulated.NewSimulatedBackend(b.GetSuiteContext(), b.T())
	testContract, _ := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)

	// Use the receipt of the contract's deploy tx to get which block it was deployed in.
	deployedBlockNumber, err := b.getTxBlockNumber(simulatedChain, testContract.DeployTx())
	Nil(b.T(), err)

	backfiller := backfill.NewContractBackfiller(b.testDB, testContract, simulatedChain)

	// Since useDB is false, this should return the block number of the contract's deploy tx.
	startHeight, err := backfiller.StartHeightForBackfill(b.GetTestContext(), false)
	Nil(b.T(), err)
	Equal(b.T(), deployedBlockNumber, startHeight)

	// Since useDB is true, but we have not filled in when the contract was last indexed,
	// this should return the block number of the contract's deploy tx.
	startHeight, err = backfiller.StartHeightForBackfill(b.GetTestContext(), true)
	Nil(b.T(), err)
	Equal(b.T(), deployedBlockNumber, startHeight)

	// Now fill in the contract's last indexed block.
	err = b.testDB.StoreLastIndexed(b.GetTestContext(), testContract.Address(), uint32(simulatedChain.GetChainID()), 1000)
	Nil(b.T(), err)

	// Since useDB is false, even though last indexed is filled in,
	// this should still return the block number of the contract's deploy tx.
	startHeight, err = backfiller.StartHeightForBackfill(b.GetTestContext(), false)
	Nil(b.T(), err)
	Equal(b.T(), deployedBlockNumber, startHeight)

	// Now that useDB is true and last indexed is filled in, this should return the last indexed block.
	startHeight, err = backfiller.StartHeightForBackfill(b.GetTestContext(), true)
	Nil(b.T(), err)
	Equal(b.T(), uint64(1000), startHeight)

}

// TestGetLogsSimulated tests the GetLogs function using a simulated blockchain.
//
//nolint:cyclop
func (b BackfillSuite) TestGetLogsSimulated() {
	// Get simulated blockchain, deploy the test contract, and set up test variables.
	simulatedChain := simulated.NewSimulatedBackend(b.GetSuiteContext(), b.T())
	simulatedChain.FundAccount(b.GetTestContext(), b.wallet.Address(), *big.NewInt(params.Ether))
	testContract, ref := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(b.GetTestContext(), nil)

	backfiller := backfill.NewContractBackfiller(b.testDB, testContract, simulatedChain)

	// Emit five events, and then fetch them with GetLogs. The first two will be fetched first,
	// then the last three after.
	tx, err := ref.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	tx, err = ref.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	// Get the block that the second transaction was executed in.
	txBlockNumberA, err := b.getTxBlockNumber(simulatedChain, tx)
	Nil(b.T(), err)

	tx, err = ref.EmitEventA(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	tx, err = ref.EmitEventB(transactOpts.TransactOpts, []byte{10}, big.NewInt(11), big.NewInt(12))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	tx, err = ref.EmitEventA(transactOpts.TransactOpts, big.NewInt(13), big.NewInt(14), big.NewInt(15))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	// Get the block that the last transaction was executed in.
	txBlockNumberB, err := b.getTxBlockNumber(simulatedChain, tx)
	Nil(b.T(), err)

	// Get the logs for the first two events.
	collectedLogs := []types.Log{}
	logs, errors, done := backfiller.GetLogs(b.GetTestContext(), 0, txBlockNumberA)
	for {
		select {
		case <-b.GetTestContext().Done():
			b.T().Error("test timed out")
		case e := <-errors:
			b.T().Error(e)
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
	logs, errors, done = backfiller.GetLogs(b.GetTestContext(), txBlockNumberA+1, txBlockNumberB)
	for {
		select {
		case <-b.GetTestContext().Done():
			b.T().Error("test timed out")
		case e := <-errors:
			b.T().Error(e)
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

// TestGetLogsMock tests the GetLogs function using a mocked blockchain for errors.
func (b BackfillSuite) TestGetLogsMock() {
	// TODO: do this with mocks
}

func (b BackfillSuite) getTxBlockNumber(chain *simulated.Backend, tx *types.Transaction) (uint64, error) {
	receipt, err := chain.TransactionReceipt(b.GetTestContext(), tx.Hash())
	if err != nil {
		return 0, err
	}
	return receipt.BlockNumber.Uint64(), nil
}
