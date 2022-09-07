package backfill_test

import (
	"fmt"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"
)

func (b BackfillSuite) TestParser() {
	chainID := gofakeit.Uint32()
	// Get simulated blockchain, deploy three test contracts, and set up test variables.
	simulatedChain := simulated.NewSimulatedBackendWithChainID(b.GetTestContext(), b.T(), big.NewInt(int64(chainID)))
	_, testRefA := b.manager.GetTestContract(b.GetTestContext(), simulatedChain)

	txContext := simulatedChain.GetTxContext(b.GetTestContext(), nil)
	tx, err := testRefA.EmitEventB(txContext.TransactOpts, []byte(gofakeit.Sentence(10)), new(big.Int).SetUint64(gofakeit.Uint64()), new(big.Int).SetUint64(gofakeit.Uint64()))
	Nil(b.T(), err)

	simulatedChain.WaitForConfirmation(b.GetTestContext(), tx)

	receipt, err := simulatedChain.TransactionReceipt(b.GetTestContext(), tx.Hash())

	for _, log := range receipt.Logs {
		filterer, err := testcontract.NewTestContractFilterer(testRefA.Address(), simulatedChain)
		Nil(b.T(), err)

		// TODO: check topic
		res, err := filterer.ParseEventB(*log)
		fmt.Println(res.Sender.String())
		Nil(b.T(), err)
		NotNil(b.T(), res)
	}
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
	contractConfigs := config.ContractConfigs{}
	for i, contract := range contracts {
		contractConfigs = append(contractConfigs, config.ContractConfig{
			Address:    contract.Address().String(),
			StartBlock: startBlocks[i],
		})
	}
	chainConfig := config.ChainConfig{
		ChainID:   chainID,
		RPCUrl:    "an rpc url is not needed for simulated backends",
		Contracts: contractConfigs,
	}

	// Set up the ChainBackfiller.
	chainBackfiller, err := backfill.NewChainBackfiller(chainID, b.testDB, simulatedChain, chainConfig)
	Nil(b.T(), err)

	b.EmitEventsForAChain(chainID, contracts, testRefs, simulatedChain, chainBackfiller, chainConfig, true)
}

// EmitEventsForAChain emits events for a chain, and if `backfill` is set to true,
// will store the events and check their existence in the database.
func (b BackfillSuite) EmitEventsForAChain(chainID uint32, contracts []contracts.DeployedContract, testRefs []*testcontract.TestContractRef, simulatedChain *simulated.Backend, chainBackfiller *backfill.ChainBackfiller, chainConfig config.ChainConfig, backfill bool) {
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
		// Backfill the chain.
		lastBlock, err := simulatedChain.BlockNumber(b.GetTestContext())
		Nil(b.T(), err)
		err = chainBackfiller.Backfill(b.GetTestContext(), lastBlock)
		Nil(b.T(), err)

		// Check that the events were written to the database.
		for _, contract := range contracts {
			// Check the storage of logs.
			logFilter := db.LogFilter{
				ChainID:         chainConfig.ChainID,
				ContractAddress: contract.Address().String(),
			}
			logs, err := b.testDB.RetrieveLogsWithFilter(b.GetTestContext(), logFilter)
			Nil(b.T(), err)
			// There should be 4 logs. One from `EmitEventA`, one from `EmitEventB`, and two
			// from `EmitEventAandB`.
			Equal(b.T(), 4, len(logs))
		}
		// Check the storage of receipts.
		receiptFilter := db.ReceiptFilter{
			ChainID: chainConfig.ChainID,
		}
		receipts, err := b.testDB.RetrieveReceiptsWithFilter(b.GetTestContext(), receiptFilter)
		Nil(b.T(), err)
		// There should be 9 receipts. One from `EmitEventA`, one from `EmitEventB`, and
		// one from `EmitEventAandB`, for each contract.
		Equal(b.T(), 9, len(receipts))
	}
}
