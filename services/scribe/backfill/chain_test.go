package backfill_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/contracts/testcontract"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
)

// TestChainBackfill tests the ChainBackfiller's ability to backfill a chain.
func (b BackfillSuite) TestChainBackfill() {
	// We need to set up multiple deploy managers, one for each contract. We will use
	// b.manager for the first contract, and create a new ones for the next two.
	managerB := testutil.NewDeployManager(b.T())
	managerC := testutil.NewDeployManager(b.T())
	// Get simulated blockchain, deploy three test contracts, and set up test variables.
	simulatedChain := simulated.NewSimulatedBackendWithChainID(b.GetSuiteContext(), b.T(), big.NewInt(1))
	simulatedChain.FundAccount(b.GetTestContext(), b.wallet.Address(), *big.NewInt(params.Ether))
	testContractA, testRefA := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)
	testContractB, testRefB := managerB.GetTestContract(b.GetTestContext(), simulatedChain)
	testContractC, testRefC := managerC.GetTestContract(b.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(b.GetTestContext(), nil)
	// Put the contracts into a slice so we can iterate over them.
	contracts := []contracts.DeployedContract{testContractA, testContractB, testContractC}
	// Put the test refs into a slice so we can iterate over them.
	testRefs := []*testcontract.TestContractRef{testRefA, testRefB, testRefC}
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

	startBlocks := make([]uint64, len(contracts))
	for i, contract := range contracts {
		deployTxHash := contract.DeployTx().Hash()
		receipt, err := simulatedChain.TransactionReceipt(b.GetTestContext(), deployTxHash)
		Nil(b.T(), err)
		startBlocks[i] = receipt.BlockNumber.Uint64()
	}
	// Set up the ChainConfig for the backfiller.
	contractConfigs := make(config.ContractConfigs)
	contractConfigs["TestContractA"] = config.ContractConfig{
		Address:    testContractA.Address().String(),
		StartBlock: startBlocks[0],
	}
	contractConfigs["TestContractB"] = config.ContractConfig{
		Address:    testContractB.Address().String(),
		StartBlock: startBlocks[1],
	}
	contractConfigs["TestContractC"] = config.ContractConfig{
		Address:    testContractC.Address().String(),
		StartBlock: startBlocks[2],
	}
	chainConfig := config.ChainConfig{
		ChainID:               1,
		RPCUrl:                "an rpc url is not needed for simulated backends",
		ConfirmationThreshold: 0,
		Contracts:             contractConfigs,
	}

	// Set up the ChainBackfiller.
	chainBackfiller, err := backfill.NewChainBackfiller(contracts, b.testDB, simulatedChain, chainConfig)
	Nil(b.T(), err)
	// Backfill the chain.
	lastBlock, err := simulatedChain.BlockNumber(b.GetTestContext())
	Nil(b.T(), err)
	err = chainBackfiller.Backfill(b.GetTestContext(), lastBlock)
	Nil(b.T(), err)

	// Check that the events were written to the database.
	for _, contract := range contracts {
		// Check the storage of logs.
		logs, err := b.testDB.RetrieveAllLogs_Test(b.GetTestContext(), true, chainConfig.ChainID, contract.Address().String())
		Nil(b.T(), err)
		// There should be 4 logs. One from `EmitEventA`, one from `EmitEventB`, and two
		// from `EmitEventAandB`.
		Equal(b.T(), 4, len(logs))
		// Check the storage of receipts.
		receipts, err := b.testDB.RetrieveAllReceipts_Test(b.GetTestContext(), true, chainConfig.ChainID, contract.Address().String())
		Nil(b.T(), err)
		// There should be 3 receipts. One from `EmitEventA`, one from `EmitEventB`, and
		// one from `EmitEventAandB`.
		Equal(b.T(), 3, len(receipts))
	}
	_ = chainBackfiller

	fmt.Println("testContractA", testContractA.Address())
	fmt.Println("testContractB", testContractB.Address())
	fmt.Println("testContractC", testContractC.Address())
	_ = testRefA
	_ = testRefB
	_ = testRefC
	_ = transactOpts
}
