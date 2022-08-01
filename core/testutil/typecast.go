package testutil

import (
	"context"

	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/contracts/replicamanager"
	"github.com/synapsecns/sanguine/core/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/core/contracts/test/homeharness"
	"github.com/synapsecns/sanguine/core/contracts/test/messageharness"
	"github.com/synapsecns/sanguine/core/contracts/test/replicamanagerharness"
	"github.com/synapsecns/sanguine/core/contracts/updatermanager"
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

// GetMessageHarness gets the message harness.
func (d *DeployManager) GetMessageHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *messageharness.MessageHarnessRef) {
	d.T().Helper()

	messageHarnessContract := d.GetContractRegistry(backend).Get(ctx, MessageHarnessType)

	messageHarness, ok := messageHarnessContract.ContractHandle().(*messageharness.MessageHarnessRef)
	assert.True(d.T(), ok)

	return messageHarnessContract, messageHarness
}

// GetHomeHarness gets the home harness.
func (d *DeployManager) GetHomeHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *homeharness.HomeHarnessRef) {
	d.T().Helper()

	messageHarnessContract := d.GetContractRegistry(backend).Get(ctx, HomeHarnessType)

	messageHarness, ok := messageHarnessContract.ContractHandle().(*homeharness.HomeHarnessRef)
	assert.True(d.T(), ok)

	return messageHarnessContract, messageHarness
}

// GetAttestationHarness gets the attestation harness.
func (d *DeployManager) GetAttestationHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *attestationharness.AttestationHarnessRef) {
	d.T().Helper()

	attestationHarnessContract := d.GetContractRegistry(backend).Get(ctx, AttestationHarnessType)

	attestationHarness, ok := attestationHarnessContract.ContractHandle().(*attestationharness.AttestationHarnessRef)
	assert.True(d.T(), ok)

	return attestationHarnessContract, attestationHarness
}

// GetReplicaManagerHarness gets the replica manager harness.
func (d *DeployManager) GetReplicaManagerHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *replicamanagerharness.ReplicaManagerHarnessRef) {
	d.T().Helper()

	replicaManagerHarnessContract := d.GetContractRegistry(backend).Get(ctx, ReplicaManagerHarnessType)
	replicaManagerHarness, ok := replicaManagerHarnessContract.ContractHandle().(*replicamanagerharness.ReplicaManagerHarnessRef)
	assert.True(d.T(), ok)

	return replicaManagerHarnessContract, replicaManagerHarness
}

// GetUpdaterManager gets the update manager.
func (d *DeployManager) GetUpdaterManager(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *updatermanager.UpdaterManagerRef) {
	d.T().Helper()

	updaterManagerContract := d.GetContractRegistry(backend).Get(ctx, UpdaterManagerType)
	updaterManager, ok := updaterManagerContract.ContractHandle().(*updatermanager.UpdaterManagerRef)
	assert.True(d.T(), ok)

	return updaterManagerContract, updaterManager
}

// GetAttestationCollector gets the attestation collector contract.
func (d *DeployManager) GetAttestationCollector(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *attestationcollector.AttestationCollectorRef) {
	d.T().Helper()

	attestationContract := d.GetContractRegistry(backend).Get(ctx, AttestationCollectorType)
	attestationCollector, ok := attestationContract.ContractHandle().(*attestationcollector.AttestationCollectorRef)
	assert.True(d.T(), ok)

	return attestationContract, attestationCollector
}

// GetReplicaManager gets the replica manager contract.
func (d *DeployManager) GetReplicaManager(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *replicamanager.ReplicaManagerRef) {
	d.T().Helper()

	replicaManagerContract := d.GetContractRegistry(backend).Get(ctx, ReplicaManagerType)
	replicaManager, ok := replicaManagerContract.ContractHandle().(*replicamanager.ReplicaManagerRef)
	assert.True(d.T(), ok)

	return replicaManagerContract, replicaManager
}
