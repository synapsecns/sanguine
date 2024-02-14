package testutil

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"testing"
)

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

// NewDeployManager creates a new DeployManager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	return &DeployManager{manager.NewDeployerManager(t, NewSynapseModuleDeployer)}
}

// SynapseModuleDeployer wraps SynapseModuleDeployer and allows typed contract handles to be returned.
type SynapseModuleDeployer struct {
	*deployer.BaseDeployer
}

// NewSynapseModuleDeployer deploys a SynapseModule contract.
func NewSynapseModuleDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return &SynapseModuleDeployer{deployer.NewSimpleDeployer(registry, backend, SynapseModuleType)}
}

// Deploy deploys a SynapseModule contract.
func (s SynapseModuleDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return s.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return synapsemodule.DeploySynapseModule(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return synapsemodule.NewSynapseModuleRef(address, backend)
	})
}
