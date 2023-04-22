package backfill_test

import (
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"
)

// TestChainBackfill tests that the ChainBackfiller can backfill events from a chain.
func (b BackfillSuite) TestChainBackfill() {
	// We need to set up multiple deploy managers, one for each contract. We will use
	// b.manager for the first contract, and create a new ones for the next two.
	managerB := testutil.NewDeployManager(b.T())
	managerC := testutil.NewDeployManager(b.T())

	// Get simulated blockchain, deploy three test contracts, and set up test variables.
	chainID := gofakeit.Uint32()

	simulatedChain := geth.NewEmbeddedBackendForChainID(b.GetTestContext(), b.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backfill.DialBackend(b.GetTestContext(), simulatedChain.RPCAddress(), b.metrics)
	Nil(b.T(), err)

	simulatedChain.FundAccount(b.GetTestContext(), b.wallet.Address(), *big.NewInt(params.Ether))
	testContractA, testRefA := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)
	testContractB, testRefB := managerB.GetTestContract(b.GetTestContext(), simulatedChain)
	testContractC, testRefC := managerC.GetTestContract(b.GetTestContext(), simulatedChain)

	contracts := []contracts.DeployedContract{testContractA, testContractB, testContractC}
	testRefs := []*testcontract.TestContractRef{testRefA, testRefB, testRefC}
	startBlocks := make([]uint64, len(contracts))

	for i, contract := range contracts {
		deployTxHash := contract.DeployTx().Hash()
		receipt, err := simulatedChain.TransactionReceipt(b.GetTestContext(), deployTxHash)
		Nil(b.T(), err)
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
		ChainID:   chainID,
		Contracts: contractConfigs,
	}
	simulatedChainArr := []backfill.ScribeBackend{simulatedClient, simulatedClient}
	chainBackfiller, err := backfill.NewChainBackfiller(b.testDB, simulatedChainArr, chainConfig, 1, b.metrics)
	Nil(b.T(), err)
	b.EmitEventsForAChain(contracts, testRefs, simulatedChain, chainBackfiller, chainConfig, true)
}

// EmitEventsForAChain emits events for a chain. If `backfill` is set to true, the function will store the events
// whilst checking their existence in the database.
func (b BackfillSuite) EmitEventsForAChain(contracts []contracts.DeployedContract, testRefs []*testcontract.TestContractRef, simulatedChain backends.SimulatedTestBackend, chainBackfiller *backfill.ChainBackfiller, chainConfig config.ChainConfig, backfill bool) {
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

	if backfill {
		err := chainBackfiller.Backfill(b.GetTestContext(), nil, false)
		Nil(b.T(), err)

		for _, contract := range contracts {
			logFilter := db.LogFilter{
				ChainID:         chainConfig.ChainID,
				ContractAddress: contract.Address().String(),
			}
			logs, err := b.testDB.RetrieveLogsWithFilter(b.GetTestContext(), logFilter, 1)
			Nil(b.T(), err)

			// There should be 4 logs. One from `EmitEventA`, one from `EmitEventB`, and two from `EmitEventAandB`.
			Equal(b.T(), 4, len(logs))
		}

		receiptFilter := db.ReceiptFilter{
			ChainID: chainConfig.ChainID,
		}
		receipts, err := b.testDB.RetrieveReceiptsWithFilter(b.GetTestContext(), receiptFilter, 1)
		Nil(b.T(), err)

		// There should be 9 receipts from `EmitEventA`,`EmitEventB`, and `EmitEventAandB` for each contract.
		Equal(b.T(), 9, len(receipts))
		totalBlockTimes := uint64(0)
		currBlock, err := simulatedChain.BlockNumber(b.GetTestContext())
		Nil(b.T(), err)
		firstBlock, err := b.testDB.RetrieveFirstBlockStored(b.GetTestContext(), chainBackfiller.ChainID())
		Nil(b.T(), err)

		for blockNum := firstBlock; blockNum <= currBlock; blockNum++ {
			_, err := b.testDB.RetrieveBlockTime(b.GetTestContext(), chainBackfiller.ChainID(), blockNum)
			if err == nil {
				totalBlockTimes++
			}
		}

		// There are `currBlock` - `firstBlock`+1 block times stored (checking after contract gets deployed).
		Equal(b.T(), currBlock-firstBlock+uint64(1), totalBlockTimes)
	}
}
