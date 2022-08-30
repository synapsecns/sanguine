package backfill_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/contracts/testcontract"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
)

// TestScribeConfirmations tests that data will not be added if a specified amount of blocks
// have not passed before the block that the data belongs to.
func (b BackfillSuite) TestScribeConfirmations() {
	b.T().Skip("TODO")
}

// TestScribeBackfill tests backfilling data from all chains.
func (b BackfillSuite) TestScribeBackfill() {
	// Set up 3 chains, and the simulated backends for each.
	chainA := gofakeit.Uint32()
	chainB := chainA + 1
	chainC := chainB + 1
	chains := []uint32{chainA, chainB, chainC}
	simulatedBackends := []*simulated.Backend{}
	for _, chain := range chains {
		simulatedBackends = append(simulatedBackends, simulated.NewSimulatedBackendWithChainID(b.GetTestContext(), b.T(), big.NewInt(int64(chain))))
	}
	type deployedContracts []contracts.DeployedContract
	type contractRefs []*testcontract.TestContractRef
	type startBlocks []uint64
	allDeployedContracts := []deployedContracts{}
	allContractRefs := []contractRefs{}
	allStartBlocks := []startBlocks{}
	// Deploy test contracts to each chain.
	for _, backend := range simulatedBackends {
		// We need to set up multiple deploy managers, one for each contract. We will use
		// b.manager for the first contract, and create a new ones for the next two.
		managerB := testutil.NewDeployManager(b.T())
		managerC := testutil.NewDeployManager(b.T())
		// Set the contracts and contract refs for each chain.
		testContractA, testRefA := b.manager.GetTestContract(b.GetTestContext(), backend)
		testContractB, testRefB := managerB.GetTestContract(b.GetTestContext(), backend)
		testContractC, testRefC := managerC.GetTestContract(b.GetTestContext(), backend)
		testContracts := []contracts.DeployedContract{testContractA, testContractB, testContractC}
		testRefs := []*testcontract.TestContractRef{testRefA, testRefB, testRefC}
		// Set the start blocks for each chain.
		var startBlocks startBlocks
		for _, contract := range testContracts {
			deployTxHash := contract.DeployTx().Hash()
			receipt, err := backend.TransactionReceipt(b.GetTestContext(), deployTxHash)
			Nil(b.T(), err)
			startBlocks = append(startBlocks, receipt.BlockNumber.Uint64())
		}
		allStartBlocks = append(allStartBlocks, startBlocks)

		// Add the contracts and contract refs to the list of all contracts and contract refs.
		allDeployedContracts = append(allDeployedContracts, testContracts)
		allContractRefs = append(allContractRefs, testRefs)
	}

	// Set up the config for the scribe.
	allContractConfigs := []config.ContractConfigs{}
	for i, deployedContracts := range allDeployedContracts {
		var contractConfig config.ContractConfigs
		for j, deployedContract := range deployedContracts {
			contractConfig = append(contractConfig, config.ContractConfig{
				Address:    deployedContract.Address().String(),
				StartBlock: allStartBlocks[i][j],
			})
		}
		allContractConfigs = append(allContractConfigs, contractConfig)
	}
	allChainConfigs := []config.ChainConfig{}
	for i, chain := range chains {
		chainConfig := config.ChainConfig{
			ChainID:               chain,
			RPCUrl:                "an rpc url is not needed for simulated backends",
			ConfirmationThreshold: 0,
			Contracts:             allContractConfigs[i],
		}
		allChainConfigs = append(allChainConfigs, chainConfig)
	}
	scribeConfig := config.Config{
		Chains: allChainConfigs,
	}

	// Set up all chain backfillers.
	chainBackfillers := []*backfill.ChainBackfiller{}
	for i, chainConfig := range allChainConfigs {
		chainBackfiller, err := backfill.NewChainBackfiller(chainConfig.ChainID, b.testDB, simulatedBackends[i], chainConfig)
		Nil(b.T(), err)
		chainBackfillers = append(chainBackfillers, chainBackfiller)
	}

	scribeBackends := make([]backfill.ScribeBackend, len(simulatedBackends))
	for i, backend := range simulatedBackends {
		scribeBackends[i] = backend
	}

	// Set up the scribe backfiller.
	scribeBackfiller, err := backfill.NewScribeBackfiller(b.testDB, scribeBackends, scribeConfig)

	// Run the backfill test for each chain.
	for i, chainBackfiller := range chainBackfillers {
		ChainBackfillTest(b, chains[i], allDeployedContracts[i], allContractRefs[i], simulatedBackends[i], chainBackfiller, allChainConfigs[i])
	}

	// Check that the data was added to the database.
	logs, err := b.testDB.UnsafeRetrieveAllLogs(b.GetTestContext(), false, 0, common.Address{})
	Nil(b.T(), err)
	// There are 4 logs per contract, and 3 contracts per chain. Since there are 3 chains, 4*3*3 = 36 logs.
	Equal(b.T(), 36, len(logs))
	receipts, err := b.testDB.UnsafeRetrieveAllReceipts(b.GetTestContext(), false, 0)
	Nil(b.T(), err)
	// There are 9 receipts per chain. Since there are 3 chains, 9*3 = 27 receipts.
	Equal(b.T(), 27, len(receipts))
}
