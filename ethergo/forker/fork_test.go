package forker_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/example"
	"github.com/synapsecns/sanguine/ethergo/example/counter"
	"github.com/synapsecns/sanguine/ethergo/forker"
	"github.com/synapsecns/sanguine/ethergo/manager"
)

func (f *ForkSuite) TestFork() {
	// create an embedded backend
	testContext := f.GetTestContext()
	backend := geth.NewEmbeddedBackend(f.GetTestContext(), f.T())

	err := forker.Fork(testContext, backend.HTTPEndpoint(), 10, func(client *ethclient.Client) {
		// deploy the counter contract
		deployer := manager.NewDeployerManager(f.T(), example.NewCounterDeployer)
		deployedContract := deployer.Get(testContext, backend, example.CounterType)
		// if you're using these often, it's recommended you extend manager and add type casted getters here, along with the global registry
		//nolint: forcetypeassert
		counterHandle := deployedContract.ContractHandle().(*counter.CounterRef)

		// first up, let's make sure we're at 0
		count, err := counterHandle.GetCount(&bind.CallOpts{Context: testContext})
		Nil(f.T(), err)
		True(f.T(), count.Int64() == 0)

		// let's increment the counter
		authOpts := backend.GetTxContext(testContext, nil)
		tx, err := counterHandle.IncrementCounter(authOpts.TransactOpts)
		Nil(f.T(), err)

		backend.WaitForConfirmation(testContext, tx)

		// we should be at 1
		count, err = counterHandle.GetCount(&bind.CallOpts{Context: testContext})
		Nil(f.T(), err)
		True(f.T(), count.Int64() == 1)
	})

	Nil(f.T(), err)

	// using the original backend, check the counter state is 0
}
