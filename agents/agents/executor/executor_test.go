package executor_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/params"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
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

	port, err := freeport.GetFreePort()
	e.Nil(err)

	clients := map[uint32][]backfill.ScribeBackend{
		chainIDA: {simulatedChainA, simulatedChainA},
		chainIDB: {simulatedChainB, simulatedChainB},
	}

	exc, err := executor.NewExecutor(e.GetTestContext(), scribeConfig, clients, e.dbPath, "sqlite", uint16(port))
	e.Nil(err)

	// Start the executor.
	go func() {
		err = exc.Start(e.GetTestContext())
		e.Nil(err)
	}()

	e.Eventually(func() bool {
		if len(exc.LogChans[chainIDA]) == 2 && len(exc.LogChans[chainIDB]) == 2 {
			logA := <-exc.LogChans[chainIDA]
			logB := <-exc.LogChans[chainIDA]
			e.Assert().Less(logA.BlockNumber, logB.BlockNumber)
			logC := <-exc.LogChans[chainIDB]
			logD := <-exc.LogChans[chainIDB]
			e.Assert().LessOrEqual(logC.BlockNumber, logD.BlockNumber)
			return true
		}

		return false
	})

}
