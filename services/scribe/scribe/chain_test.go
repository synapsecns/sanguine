package scribe_test

import (
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/scribe"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"
)

// TestChainBackfill tests that the ChainBackfiller can backfill events from a chain.
func (s ScribeSuite) TestChainBackfill() {
	// We need to set up multiple deploy managers, one for each contract. We will use
	// s.manager for the first contract, and create a new ones for the next two.
	managerB := testutil.NewDeployManager(s.T())
	managerC := testutil.NewDeployManager(s.T())

	// Get simulated blockchain, deploy three test contracts, and set up test variables.
	chainID := gofakeit.Uint32()

	simulatedChain := geth.NewEmbeddedBackendForChainID(s.GetTestContext(), s.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backend.DialBackend(s.GetTestContext(), simulatedChain.RPCAddress(), s.metrics)
	Nil(s.T(), err)

	simulatedChain.FundAccount(s.GetTestContext(), s.wallet.Address(), *big.NewInt(params.Ether))
	testContractA, testRefA := s.manager.GetTestContract(s.GetTestContext(), simulatedChain)
	testContractB, testRefB := managerB.GetTestContract(s.GetTestContext(), simulatedChain)
	testContractC, testRefC := managerC.GetTestContract(s.GetTestContext(), simulatedChain)

	contracts := []contracts.DeployedContract{testContractA, testContractB, testContractC}
	testRefs := []*testcontract.TestContractRef{testRefA, testRefB, testRefC}
	startBlocks := make([]uint64, len(contracts))

	for i, contract := range contracts {
		deployTxHash := contract.DeployTx().Hash()
		receipt, err := simulatedChain.TransactionReceipt(s.GetTestContext(), deployTxHash)
		Nil(s.T(), err)
		startBlocks[i] = receipt.BlockNumber.Uint64()
	}

	contractConfigs := config.ContractConfigs{}

	for i, contract := range contracts {
		contractConfigs = append(contractConfigs, config.ContractConfig{
			Address:    contract.Address().String(),
			StartBlock: startBlocks[i],
		})
	}

	chainConfig := config.ChainConfig{
		ChainID:       chainID,
		Contracts:     contractConfigs,
		Confirmations: 1,
	}
	simulatedChainArr := []backend.ScribeBackend{simulatedClient, simulatedClient}
	chainIndexer, err := scribe.NewChainIndexer(s.testDB, simulatedChainArr, chainConfig, s.metrics)
	Nil(s.T(), err)
	s.EmitEventsForAChain(contracts, testRefs, simulatedChain, chainIndexer, chainConfig, true)
}

// EmitEventsForAChain emits events for a chain. If `backfill` is set to true, the function will store the events
// whilst checking their existence in the database.
func (s ScribeSuite) EmitEventsForAChain(contracts []contracts.DeployedContract, testRefs []*testcontract.TestContractRef, simulatedChain backends.SimulatedTestBackend, chainBackfiller *scribe.ChainIndexer, chainConfig config.ChainConfig, backfill bool) {
	transactOpts := simulatedChain.GetTxContext(s.GetTestContext(), nil)

	// Emit events from each contract.
	for _, testRef := range testRefs {
		tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		Nil(s.T(), err)
		simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)
		tx, err = testRef.EmitEventB(transactOpts.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
		Nil(s.T(), err)
		simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)
		tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
		Nil(s.T(), err)
		simulatedChain.WaitForConfirmation(s.GetTestContext(), tx)
	}

	if backfill {
		err := chainBackfiller.Index(s.GetTestContext(), nil)
		Nil(s.T(), err)

		for _, contract := range contracts {
			logFilter := db.LogFilter{
				ChainID:         chainConfig.ChainID,
				ContractAddress: contract.Address().String(),
			}
			logs, err := s.testDB.RetrieveLogsWithFilter(s.GetTestContext(), logFilter, 1)
			Nil(s.T(), err)

			// There should be 4 logs. One from `EmitEventA`, one from `EmitEventB`, and two from `EmitEventAandB`.
			Equal(s.T(), 4, len(logs))
		}

		receiptFilter := db.ReceiptFilter{
			ChainID: chainConfig.ChainID,
		}
		receipts, err := s.testDB.RetrieveReceiptsWithFilter(s.GetTestContext(), receiptFilter, 1)
		Nil(s.T(), err)

		// There should be 9 receipts from `EmitEventA`,`EmitEventB`, and `EmitEventAandB` for each contract.
		Equal(s.T(), 9, len(receipts))
		totalBlockTimes := uint64(0)
		currBlock, err := simulatedChain.BlockNumber(s.GetTestContext())
		Nil(s.T(), err)
		firstBlock, err := s.testDB.RetrieveFirstBlockStored(s.GetTestContext(), chainConfig.ChainID)
		Nil(s.T(), err)

		for blockNum := firstBlock; blockNum <= currBlock; blockNum++ {
			_, err := s.testDB.RetrieveBlockTime(s.GetTestContext(), chainConfig.ChainID, blockNum)
			if err == nil {
				totalBlockTimes++
			}
		}

		// There are `currBlock` - `firstBlock`+1 block times stored (checking after contract gets deployed).
		Equal(s.T(), currBlock-firstBlock+uint64(1), totalBlockTimes)
	}
}
