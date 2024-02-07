package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/bridgev1"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messagebus"
	"github.com/synapsecns/sanguine/services/explorer/contracts/metaswap"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
)

// GetBridgeConfigV3 gets a typecast bridgeconfig contract.
func (d *DeployManager) GetBridgeConfigV3(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *bridgeconfig.BridgeConfigRef) {
	d.T().Helper()

	return manager.GetContract[*bridgeconfig.BridgeConfigRef](ctx, d.T(), d, backend, BridgeConfigTypeV3)
}

// GetSynapseBridge gets a typecast bridge contract.
func (d *DeployManager) GetSynapseBridge(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *bridge.BridgeRef) {
	d.T().Helper()

	return manager.GetContract[*bridge.BridgeRef](ctx, d.T(), d, backend, SynapseBridgeType)
}

// GetSynapseBridgeV1 gets a typecast bridge v1 contract.
func (d *DeployManager) GetSynapseBridgeV1(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *bridgev1.BridgeRef) {
	d.T().Helper()

	return manager.GetContract[*bridgev1.BridgeRef](ctx, d.T(), d, backend, SynapseBridgeV1Type)
}

// GetSwapFlashLoan gets a typecast swap contract.
func (d *DeployManager) GetSwapFlashLoan(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *swap.SwapRef) {
	d.T().Helper()

	return manager.GetContract[*swap.SwapRef](ctx, d.T(), d, backend, SwapFlashLoanType)
}

// GetMessageBus gets a typecast message bus contract.
func (d *DeployManager) GetMessageBus(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *messagebus.MessageBusRef) {
	d.T().Helper()

	return manager.GetContract[*messagebus.MessageBusRef](ctx, d.T(), d, backend, MessageBusType)
}

// GetMetaSwap gets a typecast meta swap.
func (d *DeployManager) GetMetaSwap(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *metaswap.MetaSwapRef) {
	d.T().Helper()

	return manager.GetContract[*metaswap.MetaSwapRef](ctx, d.T(), d, backend, MetaSwapType)
}

// GetCCTP gets a typecast cctp.
func (d *DeployManager) GetCCTP(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *cctp.CCTPRef) {
	d.T().Helper()

	return manager.GetContract[*cctp.CCTPRef](ctx, d.T(), d, backend, CCTPType)
}
