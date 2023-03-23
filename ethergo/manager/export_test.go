package manager

import (
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/deployer"
)

// SetContractRegistry sets the contract registry for the given backend.
func (d *DeployerManager) SetContractRegistry(backend backends.SimulatedTestBackend, registry deployer.ContractRegistry) {
	d.registries[backend.GetBigChainID().String()] = registry
}
