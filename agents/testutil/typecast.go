package testutil

import (
	"context"

	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/test/headerharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
	"github.com/synapsecns/sanguine/agents/contracts/test/testclient"
	"github.com/synapsecns/sanguine/ethergo/contracts"

	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/messageharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/tipsharness"
	"github.com/synapsecns/sanguine/ethergo/backends"

	"github.com/synapsecns/sanguine/agents/testutil/agentstestcontract"
)

// GetOrigin gets a typecast origin contract.
func (d *DeployManager) GetOrigin(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *origin.OriginRef) {
	d.T().Helper()

	originContract := d.GetContractRegistry(backend).Get(ctx, OriginType)

	originHandle, ok := originContract.ContractHandle().(*origin.OriginRef)
	assert.True(d.T(), ok)

	return originContract, originHandle
}

// GetMessageHarness gets the message harness.
func (d *DeployManager) GetMessageHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *messageharness.MessageHarnessRef) {
	d.T().Helper()

	messageHarnessContract := d.GetContractRegistry(backend).Get(ctx, MessageHarnessType)

	messageHarness, ok := messageHarnessContract.ContractHandle().(*messageharness.MessageHarnessRef)
	assert.True(d.T(), ok)

	return messageHarnessContract, messageHarness
}

// GetOriginHarness gets the origin harness.
func (d *DeployManager) GetOriginHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *originharness.OriginHarnessRef) {
	d.T().Helper()

	originHarnessContract := d.GetContractRegistry(backend).Get(ctx, OriginHarnessType)

	originHarness, ok := originHarnessContract.ContractHandle().(*originharness.OriginHarnessRef)

	assert.True(d.T(), ok)

	return originHarnessContract, originHarness
}

// GetAttestationHarness gets the attestation harness.
func (d *DeployManager) GetAttestationHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *attestationharness.AttestationHarnessRef) {
	d.T().Helper()

	attestationHarnessContract := d.GetContractRegistry(backend).Get(ctx, AttestationHarnessType)

	attestationHarness, ok := attestationHarnessContract.ContractHandle().(*attestationharness.AttestationHarnessRef)
	assert.True(d.T(), ok)

	return attestationHarnessContract, attestationHarness
}

// GetDestinationHarness gets the destination harness.
func (d *DeployManager) GetDestinationHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *destinationharness.DestinationHarnessRef) {
	d.T().Helper()

	destinationHarnessContract := d.GetContractRegistry(backend).Get(ctx, DestinationHarnessType)
	destinationHarness, ok := destinationHarnessContract.ContractHandle().(*destinationharness.DestinationHarnessRef)
	assert.True(d.T(), ok)

	return destinationHarnessContract, destinationHarness
}

// GetAttestationCollector gets the attestation collector contract.
func (d *DeployManager) GetAttestationCollector(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *attestationcollector.AttestationCollectorRef) {
	d.T().Helper()

	attestationContract := d.GetContractRegistry(backend).Get(ctx, AttestationCollectorType)
	attestationCollector, ok := attestationContract.ContractHandle().(*attestationcollector.AttestationCollectorRef)
	assert.True(d.T(), ok)

	return attestationContract, attestationCollector
}

// GetDestination gets the destination contract.
func (d *DeployManager) GetDestination(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *destination.DestinationRef) {
	d.T().Helper()

	destinationContract := d.GetContractRegistry(backend).Get(ctx, DestinationType)
	destination, ok := destinationContract.ContractHandle().(*destination.DestinationRef)
	assert.True(d.T(), ok)

	return destinationContract, destination
}

// GetTipsHarness gets the tips harness for testing.
func (d *DeployManager) GetTipsHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *tipsharness.TipsHarnessRef) {
	d.T().Helper()

	tipsContract := d.GetContractRegistry(backend).Get(ctx, TipsHarnessType)
	tipsHarness, ok := tipsContract.ContractHandle().(*tipsharness.TipsHarnessRef)
	assert.True(d.T(), ok)

	return tipsContract, tipsHarness
}

// GetHeaderHarness gets the header harness.
func (d *DeployManager) GetHeaderHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *headerharness.HeaderHarnessRef) {
	d.T().Helper()

	headerHarnessContract := d.GetContractRegistry(backend).Get(ctx, HeaderHarnessType)
	headerHarness, ok := headerHarnessContract.ContractHandle().(*headerharness.HeaderHarnessRef)
	assert.True(d.T(), ok)

	return headerHarnessContract, headerHarness
}

// GetAgentsTestContract gets the agents test contract.
func (d *DeployManager) GetAgentsTestContract(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *agentstestcontract.AgentsTestContractRef) {
	d.T().Helper()

	agentsTestContract := d.GetContractRegistry(backend).Get(ctx, AgentsTestContractType)
	agentsTestContractHandle, ok := agentsTestContract.ContractHandle().(*agentstestcontract.AgentsTestContractRef)
	assert.True(d.T(), ok)

	return agentsTestContract, agentsTestContractHandle
}

// GetTestClient gets the test client.
func (d *DeployManager) GetTestClient(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testclient.TestClientRef) {
	d.T().Helper()

	registry := d.GetContractRegistry(backend)
	testClient := registry.Get(ctx, TestClientType)
	testClientHandle, ok := testClient.ContractHandle().(*testclient.TestClientRef)
	assert.True(d.T(), ok)

	return testClient, testClientHandle
}

// GetPingPongClient gets the ping pong test client.
func (d *DeployManager) GetPingPongClient(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *pingpongclient.PingPongClientRef) {
	d.T().Helper()

	registry := d.GetContractRegistry(backend)
	pingPongClient := registry.Get(ctx, PingPongClientType)
	pingPongClientHandle, ok := pingPongClient.ContractHandle().(*pingpongclient.PingPongClientRef)
	assert.True(d.T(), ok)

	return pingPongClient, pingPongClientHandle
}
