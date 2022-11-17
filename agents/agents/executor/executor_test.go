package executor_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/node"
	"math/big"
)

func (e *ExecutorSuite) TestExecutor() {
	chainIDA := gofakeit.Uint32()
	chainIDB := chainIDA + 1

	simulatedChainA := simulated.NewSimulatedBackendWithChainID(e.GetTestContext(), e.T(), big.NewInt(int64(chainIDA)))
	simulatedChainB := simulated.NewSimulatedBackendWithChainID(e.GetTestContext(), e.T(), big.NewInt(int64(chainIDB)))
	simulatedChainA.FundAccount(e.GetTestContext(), e.wallet.Address(), *big.NewInt(params.Ether))
	simulatedChainB.FundAccount(e.GetTestContext(), e.wallet.Address(), *big.NewInt(params.Ether))
	testContractA, testRefA := e.manager.GetTestContract(e.GetTestContext(), simulatedChainA)
	testContractB, testRefB := e.manager.GetTestContract(e.GetTestContext(), simulatedChainB)
	transactOptsA := simulatedChainA.GetTxContext(e.GetTestContext(), nil)
	transactOptsB := simulatedChainB.GetTxContext(e.GetTestContext(), nil)

	// Emit two events on each chain.
	tx, err := testRefA.EmitEventA(transactOptsA.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	e.Nil(err)
	simulatedChainA.WaitForConfirmation(e.GetTestContext(), tx)
	tx, err = testRefA.EmitEventB(transactOptsA.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
	e.Nil(err)
	simulatedChainA.WaitForConfirmation(e.GetTestContext(), tx)
	tx, err = testRefB.EmitEventAandB(transactOptsB.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
	e.Nil(err)
	simulatedChainB.WaitForConfirmation(e.GetTestContext(), tx)

	contractConfigA := config.ContractConfig{
		Address:    testContractA.Address().String(),
		StartBlock: 0,
	}
	contractConfigB := config.ContractConfig{
		Address:    testContractB.Address().String(),
		StartBlock: 0,
	}
	chainConfigA := config.ChainConfig{
		ChainID:               chainIDA,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{contractConfigA},
	}
	chainConfigB := config.ChainConfig{
		ChainID:               chainIDB,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{contractConfigB},
	}
	chainConfigs := []config.ChainConfig{chainConfigA, chainConfigB}
	scribeConfig := config.Config{
		Chains: chainConfigs,
	}

	clients := map[uint32][]backfill.ScribeBackend{
		chainIDA: {simulatedChainA, simulatedChainA},
		chainIDB: {simulatedChainB, simulatedChainB},
	}

	scribe, err := node.NewScribe(e.testDB, clients, scribeConfig)
	e.Nil(err)

	// Start the Scribe.
	go func() {
		err := scribe.Start(e.GetTestContext())
		e.Nil(err)
	}()

	excA, err := executor.NewExecutor(testContractA.Address(), chainIDA, e.dbPath, "sqlite")
	e.Nil(err)
	excB, err := executor.NewExecutor(testContractB.Address(), chainIDB, e.dbPath, "sqlite")
	e.Nil(err)

	// Start the executor.
	go func() {
		err = excA.Start(e.GetTestContext())
		e.Nil(err)
	}()
	go func() {
		err = excB.Start(e.GetTestContext())
		e.Nil(err)
	}()

	e.Eventually(func() bool {
		if len(excA.LogChan) == 2 {
			logA := <-excA.LogChan
			logB := <-excA.LogChan
			e.Assert().Less(logA.BlockNumber, logB.BlockNumber)
			return true
		}

		return false
	})

	e.Eventually(func() bool {
		if len(excB.LogChan) == 2 {
			logA := <-excB.LogChan
			logB := <-excB.LogChan
			e.Assert().LessOrEqual(logA.BlockNumber, logB.BlockNumber)
			return true
		}

		return false
	})
}
