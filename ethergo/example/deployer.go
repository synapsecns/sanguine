package example

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/example/counter"
)

// CounterDeployer deploys a counter.
type CounterDeployer struct {
	*deployer.BaseDeployer
}

// NewCounterDeployer creates a deployer for the new counter.
func NewCounterDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return &CounterDeployer{
		deployer.NewSimpleDeployer(registry, backend, CounterType),
	}
}

// Deploy deploys the contract.
func (n *CounterDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	//nolint: wrapcheck
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		//nolint: wrapcheck
		return counter.DeployCounter(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		// this is kept separate because we often want to add an address handle to this so it's compatible with vm.ContractRef
		//nolint: wrapcheck
		return counter.NewCounterRef(address, backend)
	})
}

// compile time assertion.
var _ deployer.ContractDeployer = &CounterDeployer{}
