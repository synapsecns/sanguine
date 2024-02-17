package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
)

// GetInterchainClient gets the interchain client
func (d *DeployManager) GetInterchainClient(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *interchainclient.InterchainClientRef) {
	d.T().Helper()

	return manager.GetContract[*interchainclient.InterchainClientRef](ctx, d.T(), d, backend, InterchainClient)
}
