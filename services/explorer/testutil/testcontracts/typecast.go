package testcontracts

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap/testswap"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
)

// GetBridgeConfigV3 gets a typecast bridgeconfig contract.
func (d *DeployManager) GetBridgeConfigV3(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *bridgeconfig.BridgeConfigRef) {
	d.T().Helper()

	bridgeConfigContract := d.GetContractRegistry(backend).Get(ctx, testutil.BridgeConfigTypeV3)

	bridgeConfigHandle, ok := bridgeConfigContract.ContractHandle().(*bridgeconfig.BridgeConfigRef)
	assert.True(d.T(), ok)

	return bridgeConfigContract, bridgeConfigHandle
}

// GetTestSynapseBridge gets a typecast test bridge contract.
func (d *DeployManager) GetTestSynapseBridge(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testbridge.TestBridgeRef) {
	d.T().Helper()

	bridgeContract := d.GetContractRegistry(backend).Get(ctx, TestSynapseBridgeType)

	bridgeHandle, ok := bridgeContract.ContractHandle().(*testbridge.TestBridgeRef)
	assert.True(d.T(), ok)

	return bridgeContract, bridgeHandle
}

// GetTestSwapFlashLoan gets a typecast test swap contract.
func (d *DeployManager) GetTestSwapFlashLoan(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testswap.TestSwapRef) {
	d.T().Helper()

	swapContract := d.GetContractRegistry(backend).Get(ctx, TestSwapFlashLoanType)

	swapHandle, ok := swapContract.ContractHandle().(*testswap.TestSwapRef)
	assert.True(d.T(), ok)

	return swapContract, swapHandle
}
