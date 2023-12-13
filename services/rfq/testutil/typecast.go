package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
)

// GetFastBridge gets the pre-created fast bridge contract
func (d *DeployManager) GetFastBridge(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *fastbridge.FastBridgeRef) {
	d.T().Helper()

	return manager.GetContract[*fastbridge.FastBridgeRef](ctx, d.T(), d, backend, FastBridgeType)

}
