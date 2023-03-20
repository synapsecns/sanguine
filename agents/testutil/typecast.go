package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/manager"

	"github.com/synapsecns/sanguine/agents/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/summitharness"

	"github.com/synapsecns/sanguine/agents/contracts/test/snapshotharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/stateharness"

	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/agents/contracts/test/headerharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
	"github.com/synapsecns/sanguine/agents/contracts/test/testclient"
	"github.com/synapsecns/sanguine/ethergo/contracts"

	"github.com/synapsecns/sanguine/agents/contracts/destination"
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

	return manager.GetContract[*origin.OriginRef](ctx, d.T(), d, backend, OriginType)
}

// GetMessageHarness gets the message harness.
func (d *DeployManager) GetMessageHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *messageharness.MessageHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*messageharness.MessageHarnessRef](ctx, d.T(), d, backend, MessageHarnessType)
}

// GetOriginHarness gets the origin harness.
func (d *DeployManager) GetOriginHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *originharness.OriginHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*originharness.OriginHarnessRef](ctx, d.T(), d, backend, OriginHarnessType)
}

// GetStateHarness gets the state harness.
func (d *DeployManager) GetStateHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *stateharness.StateHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*stateharness.StateHarnessRef](ctx, d.T(), d, backend, StateHarnessType)
}

// GetSnapshotHarness gets the snapshot harness.
func (d *DeployManager) GetSnapshotHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *snapshotharness.SnapshotHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*snapshotharness.SnapshotHarnessRef](ctx, d.T(), d, backend, SnapshotHarnessType)
}

// GetAttestationHarness gets the attestation harness.
func (d *DeployManager) GetAttestationHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *attestationharness.AttestationHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*attestationharness.AttestationHarnessRef](ctx, d.T(), d, backend, AttestationHarnessType)
}

// GetDestinationHarness gets the destination harness.
func (d *DeployManager) GetDestinationHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *destinationharness.DestinationHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*destinationharness.DestinationHarnessRef](ctx, d.T(), d, backend, DestinationHarnessType)
}

// GetSummitHarness gets the summit harness.
func (d *DeployManager) GetSummitHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *summitharness.SummitHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*summitharness.SummitHarnessRef](ctx, d.T(), d, backend, SummitHarnessType)
}

// GetSummit gets the summit contract.
func (d *DeployManager) GetSummit(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *summit.SummitRef) {
	d.T().Helper()

	return manager.GetContract[*summit.SummitRef](ctx, d.T(), d, backend, SummitType)
}

// GetDestination gets the destination contract.
func (d *DeployManager) GetDestination(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *destination.DestinationRef) {
	d.T().Helper()

	return manager.GetContract[*destination.DestinationRef](ctx, d.T(), d, backend, DestinationType)
}

// GetTipsHarness gets the tips harness for testing.
func (d *DeployManager) GetTipsHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *tipsharness.TipsHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*tipsharness.TipsHarnessRef](ctx, d.T(), d, backend, TipsHarnessType)
}

// GetHeaderHarness gets the header harness.
func (d *DeployManager) GetHeaderHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *headerharness.HeaderHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*headerharness.HeaderHarnessRef](ctx, d.T(), d, backend, HeaderHarnessType)
}

// GetAgentsTestContract gets the agents test contract.
func (d *DeployManager) GetAgentsTestContract(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *agentstestcontract.AgentsTestContractRef) {
	d.T().Helper()

	return manager.GetContract[*agentstestcontract.AgentsTestContractRef](ctx, d.T(), d, backend, AgentsTestContractType)
}

// GetTestClient gets the test client.
func (d *DeployManager) GetTestClient(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testclient.TestClientRef) {
	d.T().Helper()

	return manager.GetContract[*testclient.TestClientRef](ctx, d.T(), d, backend, TestClientType)
}

// GetPingPongClient gets the ping pong test client.
func (d *DeployManager) GetPingPongClient(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *pingpongclient.PingPongClientRef) {
	d.T().Helper()

	return manager.GetContract[*pingpongclient.PingPongClientRef](ctx, d.T(), d, backend, TestClientType)
}
