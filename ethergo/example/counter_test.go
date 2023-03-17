package example_test

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/example"
	"github.com/synapsecns/sanguine/ethergo/example/counter"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	// register a test timeout
	testContext, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// since extra deployers don't necessarily deploy anything (only when requested in the GetOnlyContractRegistry)
	// adding them here won't slow anything down. It's recommended you have a global slice of these deployers you register every time.
	deployer := manager.NewDeployerManager(t, example.NewCounterDeployer)

	newTestBackend := simulated.NewSimulatedBackend(testContext, t)

	deployedContract := deployer.Get(testContext, newTestBackend, example.CounterType)
	// if you're using these often, it's recommended you extend manager and add type casted getters here, along with the global registry
	//nolint: forcetypeassert
	counterHandle := deployedContract.ContractHandle().(*counter.CounterRef)

	// first up, let's make sure we're at 0
	count, err := counterHandle.GetCount(&bind.CallOpts{Context: testContext})
	Nil(t, err)
	True(t, count.Int64() == 0)

	// let's increment the counter
	authOpts := newTestBackend.GetTxContext(testContext, nil)
	tx, err := counterHandle.IncrementCounter(authOpts.TransactOpts)
	Nil(t, err)

	newTestBackend.WaitForConfirmation(testContext, tx)

	// we should be at 1
	count, err = counterHandle.GetCount(&bind.CallOpts{Context: testContext})
	Nil(t, err)
	True(t, count.Int64() == 1)
}

func TestDependenciesCorrect(t *testing.T) {
	manager.AssertDependenciesCorrect(context.Background(), t, func() manager.IDeployManager {
		return manager.NewDeployerManager(t, example.NewCounterDeployer)
	})
}
