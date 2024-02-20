package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/committee/contracts/interchaindb"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
)

// GetSynapseModule returns a deployed SynapseModule contract and a reference to the contract.
func (d *DeployManager) GetSynapseModule(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *synapsemodule.SynapseModuleRef) {
	d.T().Helper()

	return manager.GetContract[*synapsemodule.SynapseModuleRef](ctx, d.T(), d, backend, SynapseModuleType)
}

// GetInterchainDB returns a deployed InterchainDB contract and a reference to the contract.
func (d *DeployManager) GetInterchainDB(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *interchaindb.InterchainDBRef) {
	d.T().Helper()

	return manager.GetContract[*interchaindb.InterchainDBRef](ctx, d.T(), d, backend, InterchainDB)
}
