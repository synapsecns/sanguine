package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/core/contracts/origin"
	"github.com/synapsecns/sanguine/core/contracts/test/headerharness"

	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/core/contracts/notarymanager"
	"github.com/synapsecns/sanguine/core/contracts/replicamanager"
	"github.com/synapsecns/sanguine/core/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/core/contracts/test/homeharness"
	"github.com/synapsecns/sanguine/core/contracts/test/messageharness"
	"github.com/synapsecns/sanguine/core/contracts/test/replicamanagerharness"
	"github.com/synapsecns/sanguine/core/contracts/test/tipsharness"
	"github.com/synapsecns/synapse-node/testutils/backends"
)

// GetOrigin gets a typecast origin contract.
func (d *DeployManager) GetOrigin(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *origin.OriginRef) {
	d.T().Helper()

	originContract := d.GetContractRegistry(backend).Get(ctx, OriginType)

	originHandle, ok := originContract.ContractHandle().(*origin.OriginRef)
	assert.True(d.T(), ok)

	return originContract, originHandle
}

// GetMessageHarness gets the message harness.
func (d *DeployManager) GetMessageHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *messageharness.MessageHarnessRef) {
	d.T().Helper()

	messageHarnessContract := d.GetContractRegistry(backend).Get(ctx, MessageHarnessType)

	messageHarness, ok := messageHarnessContract.ContractHandle().(*messageharness.MessageHarnessRef)
	assert.True(d.T(), ok)

	return messageHarnessContract, messageHarness
}

// GetHomeHarness gets the origin harness.
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

// GetNotaryManager gets the update manager.
func (d *DeployManager) GetNotaryManager(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *notarymanager.NotaryManagerRef) {
	d.T().Helper()

	notaryManagerContract := d.GetContractRegistry(backend).Get(ctx, NotaryManagerType)
	notaryManager, ok := notaryManagerContract.ContractHandle().(*notarymanager.NotaryManagerRef)
	assert.True(d.T(), ok)

	return notaryManagerContract, notaryManager
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

// GetTipsHarness gets the tips harness for testing.
func (d *DeployManager) GetTipsHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *tipsharness.TipsHarnessRef) {
	d.T().Helper()

	tipsContract := d.GetContractRegistry(backend).Get(ctx, TipsHarnessType)
	tipsHarness, ok := tipsContract.ContractHandle().(*tipsharness.TipsHarnessRef)
	assert.True(d.T(), ok)

	return tipsContract, tipsHarness
}

// GetHeaderHarness gets the header harness.
func (d *DeployManager) GetHeaderHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract backends.DeployedContract, handle *headerharness.HeaderHarnessRef) {
	d.T().Helper()

	headerHarnessContract := d.GetContractRegistry(backend).Get(ctx, HeaderHarnessType)
	headerHarness, ok := headerHarnessContract.ContractHandle().(*headerharness.HeaderHarnessRef)
	assert.True(d.T(), ok)

	return headerHarnessContract, headerHarness
}
