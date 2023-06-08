package node_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/node"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"
	"math/big"
	"os"
	"time"
)

// TestLive tests live recording of events.
func (l LiveSuite) TestLive() {
	if os.Getenv("CI") != "" {
		l.T().Skip("Test flake: 1 minute of livefilling may fail on CI")
	}
	chainID := gofakeit.Uint32()
	// We need to set up multiple deploy managers, one for each contract. We will use
	// b.manager for the first contract, and create a new ones for the next two.
	managerB := testutil.NewDeployManager(l.T())
	managerC := testutil.NewDeployManager(l.T())
	// Get simulated blockchain, deploy three test contracts, and set up test variables.
	simulatedChain := geth.NewEmbeddedBackendForChainID(l.GetTestContext(), l.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backfill.DialBackend(l.GetTestContext(), simulatedChain.RPCAddress(), l.metrics)
	Nil(l.T(), err)

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
		Contracts: contractConfigs,
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{chainConfig},
	}

	clients := make(map[uint32][]backfill.ScribeBackend)
	clients[chainID] = append(clients[chainID], simulatedClient)
	clients[chainID] = append(clients[chainID], simulatedClient)

	// Set up the scribe.
	scribe, err := node.NewScribe(l.testDB, clients, scribeConfig, l.metrics)
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

	// Livefill for a minute.
	ctx, cancel := context.WithTimeout(l.GetTestContext(), 1*time.Minute)
	defer cancel()
	_ = scribe.Start(ctx)

	// Check that the events were recorded.
	for _, contract := range contracts {
		// Check the storage of logs.
		logFilter := db.LogFilter{
			ChainID:         chainConfig.ChainID,
			ContractAddress: contract.Address().String(),
		}
		logs, err := l.testDB.RetrieveLogsWithFilter(l.GetTestContext(), logFilter, 1)
		Nil(l.T(), err)
		// There should be 4 logs. One from `EmitEventA`, one from `EmitEventB`, and two
		// from `EmitEventAandB`.
		Equal(l.T(), 4, len(logs))
	}
	// Check the storage of receipts.
	receiptFilter := db.ReceiptFilter{
		ChainID: chainConfig.ChainID,
	}
	receipts, err := l.testDB.RetrieveReceiptsWithFilter(l.GetTestContext(), receiptFilter, 1)
	Nil(l.T(), err)
	// There should be 9 receipts. One from `EmitEventA`, one from `EmitEventB`, and
	// one from `EmitEventAandB`, for each contract.
	Equal(l.T(), 9, len(receipts))
}

func (l LiveSuite) TestRequiredConfirmationSetting() {
	if os.Getenv("CI") != "" {
		l.T().Skip("Test flake: 1 minute of livefilling may fail on CI")
	}
	chainID := gofakeit.Uint32()

	// Emit some events on the simulated blockchain.
	simulatedChain := geth.NewEmbeddedBackendForChainID(l.GetTestContext(), l.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backfill.DialBackend(l.GetTestContext(), simulatedChain.RPCAddress(), l.metrics)
	Nil(l.T(), err)

	simulatedChain.FundAccount(l.GetTestContext(), l.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := l.manager.GetTestContract(l.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(l.GetTestContext(), nil)

	// Set up the config.
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}
	chainConfig := config.ChainConfig{
		ChainID:               chainID,
		RequiredConfirmations: 3,
		Contracts:             []config.ContractConfig{contractConfig},
	}
	scribeConfig := config.Config{
		Chains:                  []config.ChainConfig{chainConfig},
		ConfirmationRefreshRate: 1,
	}

	clients := make(map[uint32][]backfill.ScribeBackend)
	clients[chainID] = append(clients[chainID], simulatedClient)
	clients[chainID] = append(clients[chainID], simulatedClient)

	// Set up the scribe.
	scribe, err := node.NewScribe(l.testDB, clients, scribeConfig, l.metrics)
	Nil(l.T(), err)

	// Emit 5 events.
	for i := 0; i < 5; i++ {
		tx, err := testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		Nil(l.T(), err)
		simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)
	}
	// Process the events, end livefilling after a minute.
	ctx, cancel := context.WithTimeout(l.GetTestContext(), 1*time.Minute)
	defer cancel()
	_ = scribe.Start(ctx)

	// The first 2 events should be confirmed, but the last 3 should not.
	// Check logs.
	logFilter := db.LogFilter{
		ChainID:         chainConfig.ChainID,
		ContractAddress: testContract.Address().String(),
		Confirmed:       true,
	}
	logs, err := l.testDB.RetrieveLogsWithFilter(l.GetTestContext(), logFilter, 1)
	Nil(l.T(), err)
	// There should be 4 logs, two for each event over two blocks.
	Equal(l.T(), 4, len(logs))

	// Check receipts.
	receiptFilter := db.ReceiptFilter{
		ChainID:   chainConfig.ChainID,
		Confirmed: true,
	}
	receipts, err := l.testDB.RetrieveReceiptsWithFilter(l.GetTestContext(), receiptFilter, 1)
	Nil(l.T(), err)
	// There should be 2 receipts, one for each transaction over two blocks.
	Equal(l.T(), 2, len(receipts))

	// Check transactions.
	txFilter := db.EthTxFilter{
		ChainID:   chainConfig.ChainID,
		Confirmed: true,
	}
	txs, err := l.testDB.RetrieveEthTxsWithFilter(l.GetTestContext(), txFilter, 1)
	Nil(l.T(), err)
	// There should be 2 transactions, one for each transaction over two blocks.
	Equal(l.T(), 2, len(txs))

	// Add one more block to the chain by emitting another event.
	tx, err := testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(l.T(), err)
	simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)

	// Process the events.
	err = scribe.ProcessRange(l.GetTestContext(), chainID, chainConfig.RequiredConfirmations)
	Nil(l.T(), err)

	// Check logs.
	logs, err = l.testDB.RetrieveLogsWithFilter(l.GetTestContext(), logFilter, 1)
	Nil(l.T(), err)
	// There should be 6 logs, two for each event over three blocks.
	Equal(l.T(), 6, len(logs))

	// Check receipts.
	receipts, err = l.testDB.RetrieveReceiptsWithFilter(l.GetTestContext(), receiptFilter, 1)
	Nil(l.T(), err)
	// There should be 4 receipts, one for each transaction over three blocks.
	Equal(l.T(), 3, len(receipts))

	// Check transactions.
	txs, err = l.testDB.RetrieveEthTxsWithFilter(l.GetTestContext(), txFilter, 1)
	Nil(l.T(), err)
	// There should be 4 transactions, one for each transaction over three blocks.
	Equal(l.T(), 3, len(txs))
}
