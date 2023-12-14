package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/contracts/mockerc20"
)

// GetFastBridge gets the pre-created fast bridge contract.
func (d *DeployManager) GetFastBridge(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *fastbridge.FastBridgeRef) {
	d.T().Helper()

	return manager.GetContract[*fastbridge.FastBridgeRef](ctx, d.T(), d, backend, FastBridgeType)
}

// GetMockERC20 gets a mock erc20 deployed on a chain.
func (d *DeployManager) GetMockERC20(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *mockerc20.MockERC20Ref) {
	d.T().Helper()

	return manager.GetContract[*mockerc20.MockERC20Ref](ctx, d.T(), d, backend, MockERC20Type)
}
