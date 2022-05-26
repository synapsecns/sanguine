package testutil

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/contracts/xappconfig"
	"github.com/synapsecns/synapse-node/testutils/backends"
)

// GetHome gets a typecast home contract.
func (d *DeployManager) GetHome(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *home.HomeRef) {
	d.T().Helper()

	homeContract := d.GetContractRegistry(backend).Get(ctx, HomeType)

	homeHandle, ok := homeContract.ContractHandle().(*home.HomeRef)
	assert.True(d.T(), ok)

	return homeContract, homeHandle
}

// GetXAppConfig gets a typecast XAppConfig contract.
func (d *DeployManager) GetXAppConfig(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *xappconfig.XAppConfigRef) {
	d.T().Helper()

	xAppContract := d.GetContractRegistry(backend).Get(ctx, XAppConfigType)

	xAppConfig, ok := xAppContract.ContractHandle().(*xappconfig.XAppConfigRef)
	assert.True(d.T(), ok)

	return xAppContract, xAppConfig
}
