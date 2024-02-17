package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/optionslibexport"
)

// GetInterchainClient gets the interchain client
func (d *DeployManager) GetInterchainClient(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *interchainclient.InterchainClientRef) {
	d.T().Helper()

	return manager.GetContract[*interchainclient.InterchainClientRef](ctx, d.T(), d, backend, InterchainClient)
}

// GetOptionsLib gets the options library
func (d *DeployManager) GetOptionsLib(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *optionslibexport.OptionsLibExportRef) {
	d.T().Helper()

	return manager.GetContract[*optionslibexport.OptionsLibExportRef](ctx, d.T(), d, backend, OptionsLib)
}
