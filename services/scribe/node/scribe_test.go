package node_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/node"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"
)

// TestLive tests live recording of events.
func (l LiveSuite) TestLive() {
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

	// Set up the config.
	contractConfigs := config.ContractConfigs{}
	for _, contract := range contracts {
		contractConfigs = append(contractConfigs, config.ContractConfig{
			Address:    contract.Address().String(),
			StartBlock: 0,
		})
	}
	chainConfig := config.ChainConfig{
		ChainID:   chainID,
		RPCUrl:    "an rpc url is not needed for simulated backends",
		Contracts: contractConfigs,
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{chainConfig},
	}

	clients := make(map[uint32]backfill.ScribeBackend)
	clients[chainID] = simulatedChain

	// Set up the scribe.
	scribe, err := node.NewScribe(l.testDB, clients, scribeConfig)
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

	err = scribe.ProcessRange(l.GetTestContext(), chainID)
	Nil(l.T(), err)

	// Check that the events were recorded.
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
