package testcontracts

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridgev1"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/cctp/testcctp"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messagebus/testmessagebus"
	"github.com/synapsecns/sanguine/services/explorer/contracts/metaswap/testmetaswap"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap/testswap"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
)

// GetBridgeConfigV3 gets a typecast bridgeconfig contract.
func (d *DeployManager) GetBridgeConfigV3(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *bridgeconfig.BridgeConfigRef) {
	d.T().Helper()

	return manager.GetContract[*bridgeconfig.BridgeConfigRef](ctx, d.T(), d, backend, testutil.BridgeConfigTypeV3)
}

// GetTestSynapseBridge gets a typecast test bridge contract.
func (d *DeployManager) GetTestSynapseBridge(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testbridge.TestBridgeRef) {
	d.T().Helper()

	return manager.GetContract[*testbridge.TestBridgeRef](ctx, d.T(), d, backend, TestSynapseBridgeType)
}

// GetTestSwapFlashLoan gets a typecast test swap contract.
func (d *DeployManager) GetTestSwapFlashLoan(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testswap.TestSwapRef) {
	d.T().Helper()

	return manager.GetContract[*testswap.TestSwapRef](ctx, d.T(), d, backend, TestSwapFlashLoanType)
}

// GetTestSynapseBridgeV1 gets a typecast test bridge contract v1.
func (d *DeployManager) GetTestSynapseBridgeV1(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testbridgev1.TestBridgeV1Ref) {
	d.T().Helper()

	return manager.GetContract[*testbridgev1.TestBridgeV1Ref](ctx, d.T(), d, backend, TestSynapseBridgeV1Type)
}

// GetTestMessageBusUpgradeable gets a typecast test message bus contract.
func (d *DeployManager) GetTestMessageBusUpgradeable(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testmessagebus.TestMessageBusRef) {
	d.T().Helper()

	return manager.GetContract[*testmessagebus.TestMessageBusRef](ctx, d.T(), d, backend, TestMessageBusType)
}

// GetTestMetaSwap gets a typecast test meta swap contract.
func (d *DeployManager) GetTestMetaSwap(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testmetaswap.TestMetaSwapRef) {
	d.T().Helper()

	return manager.GetContract[*testmetaswap.TestMetaSwapRef](ctx, d.T(), d, backend, TestMetaSwapType)
}

// GetTestCCTP gets a typecast test cctp contract.
func (d *DeployManager) GetTestCCTP(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testcctp.TestCCTPRef) {
	d.T().Helper()

	return manager.GetContract[*testcctp.TestCCTPRef](ctx, d.T(), d, backend, TestCCTPType)
}
