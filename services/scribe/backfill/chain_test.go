package backfill_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/contracts/testcontract"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
)

// TestConfirmations tests that data will not be added if a specified amount of blocks
// have not passed before the block that the data belongs to.
func (b BackfillSuite) TestConfirmations() {
	// Get simulated blockchain, deploy three test contracts, and set up test variables.
	simulatedChain := simulated.NewSimulatedBackendWithChainID(b.GetTestContext(), b.T(), big.NewInt(4))
	simulatedChain.FundAccount(b.GetTestContext(), b.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)
	// Create a second test contract just meant to pass blocks.
	dummyManager := testutil.NewDeployManager(b.T())
	_, dummyRef := dummyManager.GetTestContract(b.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(b.GetTestContext(), nil)
	// Set up the config.
	deployTxHash := testContract.DeployTx().Hash()
	receipt, err := simulatedChain.TransactionReceipt(b.GetTestContext(), deployTxHash)
	Nil(b.T(), err)
	startBlock := receipt.BlockNumber.Uint64()
	contractConfigs := config.ContractConfigs{}
	contractConfigs = append(contractConfigs, config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: startBlock,
	})
	chainConfig := config.ChainConfig{
		ChainID:               4,
		RPCUrl:                "an rpc url is not needed for simulated backends",
		ConfirmationThreshold: 2,
		Contracts:             contractConfigs,
	}

	// Set up the ChainBackfiller.
	chainBackfiller, err := backfill.NewChainBackfiller(4, b.testDB, simulatedChain, chainConfig)
	Nil(b.T(), err)

	// Emit three events from two transactions.
	tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(4), big.NewInt(5), big.NewInt(6))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Use the dummy contract to pass two blocks.
	tx, err = dummyRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	tx, err = dummyRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Backfill the first batch of events.
	latestBlock, err := simulatedChain.BlockNumber(b.GetTestContext())
	Nil(b.T(), err)
	err = chainBackfiller.Backfill(b.GetTestContext(), latestBlock-uint64(chainConfig.ConfirmationThreshold))
	Nil(b.T(), err)

	// Check that the first batch of events were added to the database.
	logs, err := b.testDB.UnsafeRetrieveAllLogs(b.GetTestContext(), true, chainConfig.ChainID, testContract.Address())
	Nil(b.T(), err)
	Equal(b.T(), 3, len(logs))

	receipts, err := b.testDB.UnsafeRetrieveAllReceipts(b.GetTestContext(), true, chainConfig.ChainID)
	Nil(b.T(), err)
	Equal(b.T(), 2, len(receipts))

	// Send one more transaction.
	tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{7}, big.NewInt(8), big.NewInt(9))
	Nil(b.T(), err)
	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	// Backfill before the confirmation threshold has passed.
	latestBlock, err = simulatedChain.BlockNumber(b.GetTestContext())
	Nil(b.T(), err)
	err = chainBackfiller.Backfill(b.GetTestContext(), latestBlock-uint64(chainConfig.ConfirmationThreshold))
	Nil(b.T(), err)

	// Check that the second batch of events were not added to the database.
	logs, err = b.testDB.UnsafeRetrieveAllLogs(b.GetTestContext(), true, chainConfig.ChainID, testContract.Address())
	Nil(b.T(), err)
	Equal(b.T(), 3, len(logs))

	receipts, err = b.testDB.UnsafeRetrieveAllReceipts(b.GetTestContext(), true, chainConfig.ChainID)
	Nil(b.T(), err)
	Equal(b.T(), 2, len(receipts))
}

// TestChainBackfill tests that the ChainBackfiller can backfill events from a chain.
func (b BackfillSuite) TestChainBackfill() {
	chainID := gofakeit.Uint32()
	// We need to set up multiple deploy managers, one for each contract. We will use
	// b.manager for the first contract, and create a new ones for the next two.
	managerB := testutil.NewDeployManager(b.T())
	managerC := testutil.NewDeployManager(b.T())
	// Get simulated blockchain, deploy three test contracts, and set up test variables.
	simulatedChain := simulated.NewSimulatedBackendWithChainID(b.GetTestContext(), b.T(), big.NewInt(int64(chainID)))
	simulatedChain.FundAccount(b.GetTestContext(), b.wallet.Address(), *big.NewInt(params.Ether))
	testContractA, testRefA := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)
	testContractB, testRefB := managerB.GetTestContract(b.GetTestContext(), simulatedChain)
	testContractC, testRefC := managerC.GetTestContract(b.GetTestContext(), simulatedChain)
	// Put the contracts into a slice so we can iterate over them.
	contracts := []contracts.DeployedContract{testContractA, testContractB, testContractC}
	// Put the test refs into a slice so we can iterate over them.
	testRefs := []*testcontract.TestContractRef{testRefA, testRefB, testRefC}

	startBlocks := make([]uint64, len(contracts))
	for i, contract := range contracts {
		deployTxHash := contract.DeployTx().Hash()
		receipt, err := simulatedChain.TransactionReceipt(b.GetTestContext(), deployTxHash)
		Nil(b.T(), err)
		startBlocks[i] = receipt.BlockNumber.Uint64()
	}
	// Set up the ChainConfig for the backfiller.
	// contractConfigs := make(config.ContractConfigs)
	contractConfigs := config.ContractConfigs{}
	for i, contract := range contracts {
		contractConfigs = append(contractConfigs, config.ContractConfig{
			Address:    contract.Address().String(),
			StartBlock: startBlocks[i],
		})
	}
	chainConfig := config.ChainConfig{
		ChainID:               chainID,
		RPCUrl:                "an rpc url is not needed for simulated backends",
		ConfirmationThreshold: 0,
		Contracts:             contractConfigs,
	}

	// Set up the ChainBackfiller.
	chainBackfiller, err := backfill.NewChainBackfiller(chainID, b.testDB, simulatedChain, chainConfig)
	Nil(b.T(), err)

	ChainBackfillTest(b, chainID, contracts, testRefs, simulatedChain, chainBackfiller, chainConfig)
}

// ChainBackfillTest tests the ChainBackfiller's ability to backfill a chain.
func ChainBackfillTest(b BackfillSuite, chainID uint32, contracts []contracts.DeployedContract, testRefs []*testcontract.TestContractRef, simulatedChain *simulated.Backend, chainBackfiller *backfill.ChainBackfiller, chainConfig config.ChainConfig) {
	transactOpts := simulatedChain.GetTxContext(b.GetTestContext(), nil)
	// Emit events from each contract.
	for _, testRef := range testRefs {
		tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		Nil(b.T(), err)
		simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
		tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
		Nil(b.T(), err)
		simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
		tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
		Nil(b.T(), err)
		simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)
	}

	// Backfill the chain.
	lastBlock, err := simulatedChain.BlockNumber(b.GetTestContext())
	Nil(b.T(), err)
	err = chainBackfiller.Backfill(b.GetTestContext(), lastBlock)
	Nil(b.T(), err)

	// Check that the events were written to the database.
	for _, contract := range contracts {
		// Check the storage of logs.
		logs, err := b.testDB.UnsafeRetrieveAllLogs(b.GetTestContext(), true, chainConfig.ChainID, contract.Address())
		Nil(b.T(), err)
		// There should be 4 logs. One from `EmitEventA`, one from `EmitEventB`, and two
		// from `EmitEventAandB`.
		Equal(b.T(), 4, len(logs))
	}
	// Check the storage of receipts.
	receipts, err := b.testDB.UnsafeRetrieveAllReceipts(b.GetTestContext(), true, chainConfig.ChainID)
	Nil(b.T(), err)
	// There should be 9 receipts. One from `EmitEventA`, one from `EmitEventB`, and
	// one from `EmitEventAandB`, for each contract.
	Equal(b.T(), 9, len(receipts))
}
