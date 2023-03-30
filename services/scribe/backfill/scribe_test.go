package backfill_test

import (
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"math/big"
	"sync"

	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"
)

// TestScribeBackfill tests backfilling data from all chains.
//
//nolint:cyclop
func (b BackfillSuite) TestScribeBackfill() {
	// Set up 3 chains, and the simulated backends for each.
	chainA := gofakeit.Uint32()
	chainB := chainA + 1
	chainC := chainB + 1
	chains := []uint32{chainA, chainB, chainC}

	simulatedBackends := make([]*geth.Backend, len(chains))
	simulatedClients := make([]backfill.ScribeBackend, len(chains))

	var wg sync.WaitGroup
	var mux sync.Mutex

	for i, chain := range chains {
		// capture func literals
		chain := chain
		i := i

		wg.Add(1)

		go func() {
			defer wg.Done()
			simulatedBackend := geth.NewEmbeddedBackendForChainID(b.GetTestContext(), b.T(), big.NewInt(int64(chain)))
			simulatedClient, err := backfill.DialBackend(b.GetTestContext(), simulatedBackend.RPCAddress(), b.metrics)
			Nil(b.T(), err)

			mux.Lock()
			defer mux.Unlock()
			simulatedBackends[i] = simulatedBackend
			simulatedClients[i] = simulatedClient
		}()
	}
	wg.Wait()

	type deployedContracts []contracts.DeployedContract
	type contractRefs []*testcontract.TestContractRef
	type startBlocks []uint64
	var allDeployedContracts []deployedContracts
	var allContractRefs []contractRefs
	var allStartBlocks []startBlocks
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
			ChainID:   chain,
			Contracts: allContractConfigs[i],
		}
		allChainConfigs = append(allChainConfigs, chainConfig)
	}
	scribeConfig := config.Config{
		Chains: allChainConfigs,
	}

	// Set up all chain backfillers.
	chainBackfillers := []*backfill.ChainBackfiller{}
	for i, chainConfig := range allChainConfigs {
		simulatedChainArr := []backfill.ScribeBackend{simulatedClients[i], simulatedClients[i]}
		chainBackfiller, err := backfill.NewChainBackfiller(b.testDB, simulatedChainArr, chainConfig, 1, b.metrics)
		Nil(b.T(), err)
		chainBackfillers = append(chainBackfillers, chainBackfiller)
	}

	scribeBackends := make(map[uint32][]backfill.ScribeBackend)
	for i := range simulatedBackends {
		client := simulatedClients[i]
		backend := simulatedBackends[i]

		simulatedChainArr := []backfill.ScribeBackend{client, client}
		scribeBackends[uint32(backend.GetChainID())] = simulatedChainArr
	}

	// Set up the scribe backfiller.
	scribeBackfiller, err := backfill.NewScribeBackfiller(b.testDB, scribeBackends, scribeConfig, b.metrics)
	Nil(b.T(), err)

	// Run the backfill test for each chain.
	for i, chainBackfiller := range chainBackfillers {
		b.EmitEventsForAChain(allDeployedContracts[i], allContractRefs[i], simulatedBackends[i], chainBackfiller, allChainConfigs[i], false)
	}

	// Run the scribe's backfill.
	err = scribeBackfiller.Backfill(b.GetTestContext())
	Nil(b.T(), err)

	// Check that the data was added to the database.
	logs, err := b.testDB.RetrieveLogsWithFilter(b.GetTestContext(), db.LogFilter{}, 1)
	Nil(b.T(), err)
	// There are 4 logs per contract, and 3 contracts per chain. Since there are 3 chains, 4*3*3 = 36 logs.
	Equal(b.T(), 36, len(logs))
	receipts, err := b.testDB.RetrieveReceiptsWithFilter(b.GetTestContext(), db.ReceiptFilter{}, 1)
	Nil(b.T(), err)
	// There are 9 receipts per chain. Since there are 3 chains, 9*3 = 27 receipts.
	Equal(b.T(), 27, len(receipts))

	for _, chainBackfiller := range chainBackfillers {
		totalBlockTimes := uint64(0)
		currBlock, err := scribeBackfiller.Clients()[chainBackfiller.ChainID()][0].BlockNumber(b.GetTestContext())
		Nil(b.T(), err)
		firstBlock, err := b.testDB.RetrieveFirstBlockStored(b.GetTestContext(), chainBackfiller.ChainID())
		Nil(b.T(), err)
		for blockNum := firstBlock; blockNum <= currBlock; blockNum++ {
			_, err := b.testDB.RetrieveBlockTime(b.GetTestContext(), chainBackfiller.ChainID(), blockNum)
			if err == nil {
				totalBlockTimes++
			}
		}
		// There are `currBlock` - `firstBlock`+1 block times stored. events don't get emitted until the contract gets deployed.
		Equal(b.T(), currBlock-firstBlock+uint64(1), totalBlockTimes)
	}
}
