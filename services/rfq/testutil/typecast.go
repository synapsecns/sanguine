package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/fastbridgemock"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/mockerc20"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/weth9"
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

// GetMockFastBridge gets the mock fast bridge.
func (d *DeployManager) GetMockFastBridge(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *fastbridgemock.FastBridgeMockRef) {
	d.T().Helper()

	return manager.GetContract[*fastbridgemock.FastBridgeMockRef](ctx, d.T(), d, backend, FastBridgeMockType)
}

// GetWETH9 gets the weth9 contract.
func (d *DeployManager) GetWETH9(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *weth9.Weth9Ref) {
	d.T().Helper()

	return manager.GetContract[*weth9.Weth9Ref](ctx, d.T(), d, backend, WETH9Type)
}
