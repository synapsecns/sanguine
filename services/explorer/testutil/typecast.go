package testutil

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
)

// GetBridgeConfigV3 gets a typecast bridgeconfig contract.
func (d *DeployManager) GetBridgeConfigV3(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *bridgeconfig.BridgeConfigRef) {
	d.T().Helper()

	bridgeConfigContract := d.GetContractRegistry(backend).Get(ctx, BridgeConfigTypeV3)

	bridgeConfigHandle, ok := bridgeConfigContract.ContractHandle().(*bridgeconfig.BridgeConfigRef)
	assert.True(d.T(), ok)

	return bridgeConfigContract, bridgeConfigHandle
}
