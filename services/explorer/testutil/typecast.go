package testutil

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/bridgev1"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messagebus"
	"github.com/synapsecns/sanguine/services/explorer/contracts/metaswap"
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

// GetSynapseBridgeV1 gets a typecast bridge v1 contract.
func (d *DeployManager) GetSynapseBridgeV1(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *bridgev1.BridgeRef) {
	d.T().Helper()
	bridgeContract := d.GetContractRegistry(backend).Get(ctx, SynapseBridgeV1Type)
	bridgeHandle, ok := bridgeContract.ContractHandle().(*bridgev1.BridgeRef)
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

// GetMessageBus gets a typecast message bus contract.
func (d *DeployManager) GetMessageBus(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *messagebus.MessageBusRef) {
	d.T().Helper()
	messageContract := d.GetContractRegistry(backend).Get(ctx, MessageBusType)
	messageHandle, ok := messageContract.ContractHandle().(*messagebus.MessageBusRef)
	assert.True(d.T(), ok)

	return messageContract, messageHandle
}

// GetMetaSwap gets a typecast meta swap.
func (d *DeployManager) GetMetaSwap(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *metaswap.MetaSwapRef) {
	d.T().Helper()
	metaSwapContract := d.GetContractRegistry(backend).Get(ctx, MetaSwapType)
	metaSwapHandle, ok := metaSwapContract.ContractHandle().(*metaswap.MetaSwapRef)
	assert.True(d.T(), ok)

	return metaSwapContract, metaSwapHandle
}
