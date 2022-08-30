package livefill_test

import (
	"math/big"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/contracts/testcontract"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/livefill"
)

// TestChainLivefill tests that the ChainLivefiller can livefill data for a chain.
func (l LivefillSuite) TestChainLivefill() {
	chainID := gofakeit.Uint32()
	// We need to set up multiple deploy managers, one for each contract. We will use
	// b.manager for the first contract, and create a new ones for the next two.
	managerB := testutil.NewDeployManager(l.T())
	managerC := testutil.NewDeployManager(l.T())
	// Get simulated blockchain, deploy three test contracts, and set up test variables.
	simulatedChain := simulated.NewSimulatedBackendWithChainID(l.GetTestContext(), l.T(), big.NewInt(int64(chainID)))
	simulatedChain.FundAccount(l.GetTestContext(), l.wallet.Address(), *big.NewInt(params.Ether))
	testContractA, testRefA := l.manager.GetTestContract(l.GetTestContext(), simulatedChain)
	testContractB, testRefB := managerB.GetTestContract(l.GetTestContext(), simulatedChain)
	testContractC, testRefC := managerC.GetTestContract(l.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(l.GetTestContext(), nil)
	// Put the contracts into a slice so we can iterate over them.
	contracts := []contracts.DeployedContract{testContractA, testContractB, testContractC}
	// Put the test refs into a slice so we can iterate over them.
	testRefs := []*testcontract.TestContractRef{testRefA, testRefB, testRefC}
	// Set up the ChainConfig for the backfiller.
	contractConfigs := config.ContractConfigs{}
	for _, contract := range contracts {
		contractConfigs = append(contractConfigs, config.ContractConfig{
			Address:    contract.Address().String(),
			StartBlock: 0,
		})
	}
	chainConfig := config.ChainConfig{
		ChainID:               chainID,
		RPCUrl:                "an rpc url is not needed for simulated backends",
		ConfirmationThreshold: 0,
		Contracts:             contractConfigs,
	}

	// Set up the ChainLivefiller.
	chainLivefiller, err := livefill.NewChainLivefiller(chainID, l.testDB, simulatedChain, chainConfig)
	Nil(l.T(), err)

	go func() {
		// Start the livefiller.
		err = chainLivefiller.Livefill(l.GetTestContext())
		Nil(l.T(), err)
	}()
	time.Sleep(1 * time.Second)

	// Emit events from each contract.
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

	time.Sleep(1 * time.Second)

	// Check that the events were written to the database.
	for _, contract := range contracts {
		// Check the storage of logs.
		logs, err := l.testDB.UnsafeRetrieveAllLogs(l.GetTestContext(), true, chainConfig.ChainID, contract.Address())
		Nil(l.T(), err)
		// There should be 4 logs. One from `EmitEventA`, one from `EmitEventB`, and two
		// from `EmitEventAandB`.
		Equal(l.T(), 4, len(logs))
	}
	// Check the storage of receipts.
	receipts, err := l.testDB.UnsafeRetrieveAllReceipts(l.GetTestContext(), true, chainConfig.ChainID)
	Nil(l.T(), err)
	// There should be 9 receipts. One from `EmitEventA`, one from `EmitEventB`, and
	// one from `EmitEventAandB`, for each contract.
	Equal(l.T(), 9, len(receipts))
}
