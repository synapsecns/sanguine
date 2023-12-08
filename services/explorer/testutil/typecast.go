package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridgev1"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/cctp/testcctp"
	"github.com/synapsecns/sanguine/services/explorer/contracts/erc20"
	"github.com/synapsecns/sanguine/services/explorer/contracts/lptoken"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messagebus/testmessagebus"
	"github.com/synapsecns/sanguine/services/explorer/contracts/metaswap/testmetaswap"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap/testswap"
)

// GetBridgeConfigV3 gets a typecast bridgeconfig contract.
func (d *DeployManager) GetBridgeConfigV3(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *bridgeconfig.BridgeConfigRef) {
	d.T().Helper()

	return manager.GetContract[*bridgeconfig.BridgeConfigRef](ctx, d.T(), d, backend, BridgeConfigTypeV3)
}

// GetSynapseBridge gets a typecast bridge contract.
func (d *DeployManager) GetSynapseBridge(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testbridge.TestBridgeRef) {
	d.T().Helper()

	return manager.GetContract[*testbridge.TestBridgeRef](ctx, d.T(), d, backend, SynapseBridgeType)
}

// GetSynapseBridgeV1 gets a typecast bridge v1 contract.
func (d *DeployManager) GetSynapseBridgeV1(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testbridgev1.TestBridgeV1Ref) {
	d.T().Helper()

	return manager.GetContract[*testbridgev1.TestBridgeV1Ref](ctx, d.T(), d, backend, SynapseBridgeV1Type)
}

// GetSwapFlashLoan gets a typecast swap contract.
func (d *DeployManager) GetSwapFlashLoan(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testswap.TestSwapRef) {
	d.T().Helper()

	return manager.GetContract[*testswap.TestSwapRef](ctx, d.T(), d, backend, SwapFlashLoanType)
}

// GetMessageBus gets a typecast message bus contract.
func (d *DeployManager) GetMessageBus(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testmessagebus.TestMessageBusRef) {
	d.T().Helper()

	return manager.GetContract[*testmessagebus.TestMessageBusRef](ctx, d.T(), d, backend, MessageBusType)
}

// GetMetaSwap gets a typecast meta swap.
func (d *DeployManager) GetMetaSwap(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testmetaswap.TestMetaSwapRef) {
	d.T().Helper()

	return manager.GetContract[*testmetaswap.TestMetaSwapRef](ctx, d.T(), d, backend, MetaSwapType)
}

// GetCCTP gets a typecast cctp.
func (d *DeployManager) GetCCTP(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testcctp.TestCCTPRef) {
	d.T().Helper()

	return manager.GetContract[*testcctp.TestCCTPRef](ctx, d.T(), d, backend, CCTPType)
}

// GetERC20A gets a typecast test erc20.
func (d *DeployManager) GetERC20A(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *erc20.TestERC20A) {
	d.T().Helper()
	return manager.GetContract[*erc20.TestERC20A](ctx, d.T(), d, backend, ERC20TypeA)
}

// GetERC20B gets the second typecast test erc20 .
func (d *DeployManager) GetERC20B(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *erc20.TestERC20B) {
	d.T().Helper()
	return manager.GetContract[*erc20.TestERC20B](ctx, d.T(), d, backend, ERC20TypeB)
}

// GetLPToken gets the typecast lp token.
func (d *DeployManager) GetLPToken(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *lptoken.LPTokenRef) {
	d.T().Helper()
	return manager.GetContract[*lptoken.LPTokenRef](ctx, d.T(), d, backend, LPTokenType)
}
