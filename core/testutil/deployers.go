package testutil

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/synapse-node/testutils/backends"
)

// HomeDeployer deploys the home contract.
type HomeDeployer struct {
	*deployer.BaseDeployer
}

// NewHomeDeployer deploys the home contract.
func NewHomeDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return HomeDeployer{deployer.NewSimpleDeployer(registry, backend, HomeType)}
}

// Deploy deploys the home contract.
func (d HomeDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return home.DeployHome(transactOps, backend, uint32(d.Backend().GetChainID()))
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return home.NewHomeRef(address, backend)
	})
}
