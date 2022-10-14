package testutil

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/message"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
)

// GetBridgeConfigV3 gets a typecast bridgeconfig contract.
func (d *DeployManager) GetBridgeConfigV3(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *bridgeconfig.BridgeConfigRef) {
	d.T().Helper()

	bridgeConfigContract := d.GetContractRegistry(backend).Get(ctx, BridgeConfigTypeV3)

	bridgeConfigHandle, ok := bridgeConfigContract.ContractHandle().(*bridgeconfig.BridgeConfigRef)
	assert.True(d.T(), ok)

	return bridgeConfigContract, bridgeConfigHandle
}

// GetSynapseBridge gets a typecast bridge contract.
func (d *DeployManager) GetSynapseBridge(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *bridge.BridgeRef) {
	d.T().Helper()

	bridgeContract := d.GetContractRegistry(backend).Get(ctx, SynapseBridgeType)

	bridgeHandle, ok := bridgeContract.ContractHandle().(*bridge.BridgeRef)
	assert.True(d.T(), ok)

	return bridgeContract, bridgeHandle
}

// GetSwapFlashLoan gets a typecast swap contract.
func (d *DeployManager) GetSwapFlashLoan(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *swap.SwapRef) {
	d.T().Helper()

	swapContract := d.GetContractRegistry(backend).Get(ctx, SwapFlashLoanType)

	swapHandle, ok := swapContract.ContractHandle().(*swap.SwapRef)
	assert.True(d.T(), ok)

	return swapContract, swapHandle
}

// GetMessageBusUpgradeable gets a typecast swap contract.
func (d *DeployManager) GetMessageBusUpgradeable(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *message.MessageRef) {
	d.T().Helper()

	messageContract := d.GetContractRegistry(backend).Get(ctx, MessageBusUpgradeableType)

	messageHandle, ok := messageContract.ContractHandle().(*message.MessageRef)
	assert.True(d.T(), ok)

	return messageContract, messageHandle
}
