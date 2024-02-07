package testutil

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/agents/contracts/test/receiptharness"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/bondingmanager"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightmanager"
	"github.com/synapsecns/sanguine/agents/contracts/test/basemessageharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/bondingmanagerharness"
	gasdataharness "github.com/synapsecns/sanguine/agents/contracts/test/gasdata"
	"github.com/synapsecns/sanguine/agents/contracts/test/lightmanagerharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/requestharness"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"golang.org/x/sync/errgroup"

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
// nolint:dupl
func (d *DeployManager) GetOrigin(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *origin.OriginRef) {
	d.T().Helper()

	return manager.GetContract[*origin.OriginRef](ctx, d.T(), d, backend, OriginType)
}

// GetMessageHarness gets the message harness.
// nolint:dupl
func (d *DeployManager) GetMessageHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *messageharness.MessageHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*messageharness.MessageHarnessRef](ctx, d.T(), d, backend, MessageHarnessType)
}

// GetBaseMessageHarness gets the base message harness.
// nolint:dupl
func (d *DeployManager) GetBaseMessageHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *basemessageharness.BaseMessageHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*basemessageharness.BaseMessageHarnessRef](ctx, d.T(), d, backend, BaseMessageHarnessType)
}

// GetRequestHarness gets the request harness.
// nolint:dupl
func (d *DeployManager) GetRequestHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *requestharness.RequestHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*requestharness.RequestHarnessRef](ctx, d.T(), d, backend, RequestHarnessType)
}

// GetLightInbox gets the light inbox.
// nolint:dupl
func (d *DeployManager) GetLightInbox(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *lightinbox.LightInboxRef) {
	d.T().Helper()

	return manager.GetContract[*lightinbox.LightInboxRef](ctx, d.T(), d, backend, LightInboxType)
}

// GetLightManager gets the light manager.
// nolint:dupl
func (d *DeployManager) GetLightManager(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *lightmanager.LightManagerRef) {
	d.T().Helper()

	return manager.GetContract[*lightmanager.LightManagerRef](ctx, d.T(), d, backend, LightManagerType)
}

// GetLightManagerHarness gets the light manager harness.
// nolint:dupl
func (d *DeployManager) GetLightManagerHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *lightmanagerharness.LightManagerHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*lightmanagerharness.LightManagerHarnessRef](ctx, d.T(), d, backend, LightManagerHarnessType)
}

// GetInbox gets the inbox.
// nolint:dupl
func (d *DeployManager) GetInbox(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *inbox.InboxRef) {
	d.T().Helper()

	return manager.GetContract[*inbox.InboxRef](ctx, d.T(), d, backend, InboxType)
}

// GetBondingManager gets the bonding manager.
// nolint:dupl
func (d *DeployManager) GetBondingManager(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *bondingmanager.BondingManagerRef) {
	d.T().Helper()

	return manager.GetContract[*bondingmanager.BondingManagerRef](ctx, d.T(), d, backend, BondingManagerType)
}

// GetBondingManagerHarness gets the bonding manager harness.
// nolint:dupl
func (d *DeployManager) GetBondingManagerHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *bondingmanagerharness.BondingManagerHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*bondingmanagerharness.BondingManagerHarnessRef](ctx, d.T(), d, backend, BondingManagerHarnessType)
}

// GetOriginHarness gets the origin harness.
// nolint:dupl
func (d *DeployManager) GetOriginHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *originharness.OriginHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*originharness.OriginHarnessRef](ctx, d.T(), d, backend, OriginHarnessType)
}

// GetGasDataHarness gets the gasData harness.
// nolint:dupl
func (d *DeployManager) GetGasDataHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *gasdataharness.GasDataHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*gasdataharness.GasDataHarnessRef](ctx, d.T(), d, backend, GasDataHarnessType)
}

// GetStateHarness gets the state harness.
// nolint:dupl
func (d *DeployManager) GetStateHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *stateharness.StateHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*stateharness.StateHarnessRef](ctx, d.T(), d, backend, StateHarnessType)
}

// GetSnapshotHarness gets the snapshot harness.
// nolint:dupl
func (d *DeployManager) GetSnapshotHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *snapshotharness.SnapshotHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*snapshotharness.SnapshotHarnessRef](ctx, d.T(), d, backend, SnapshotHarnessType)
}

// GetReceiptHarness gets the receipt harness.
// nolint:dupl
func (d *DeployManager) GetReceiptHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *receiptharness.ReceiptHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*receiptharness.ReceiptHarnessRef](ctx, d.T(), d, backend, ReceiptHarnessType)
}

// GetAttestationHarness gets the attestation harness.
// nolint:dupl
func (d *DeployManager) GetAttestationHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *attestationharness.AttestationHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*attestationharness.AttestationHarnessRef](ctx, d.T(), d, backend, AttestationHarnessType)
}

// GetDestinationHarness gets the destination harness.
// nolint:dupl
func (d *DeployManager) GetDestinationHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *destinationharness.DestinationHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*destinationharness.DestinationHarnessRef](ctx, d.T(), d, backend, DestinationHarnessType)
}

// GetSummitHarness gets the summit harness.
// nolint:dupl
func (d *DeployManager) GetSummitHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *summitharness.SummitHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*summitharness.SummitHarnessRef](ctx, d.T(), d, backend, SummitHarnessType)
}

// GetSummit gets the summit contract.
// nolint:dupl
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
// nolint:dupl
func (d *DeployManager) GetTipsHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *tipsharness.TipsHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*tipsharness.TipsHarnessRef](ctx, d.T(), d, backend, TipsHarnessType)
}

// GetHeaderHarness gets the header harness.
// nolint:dupl
func (d *DeployManager) GetHeaderHarness(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *headerharness.HeaderHarnessRef) {
	d.T().Helper()

	return manager.GetContract[*headerharness.HeaderHarnessRef](ctx, d.T(), d, backend, HeaderHarnessType)
}

// GetAgentsTestContract gets the agents test contract.
// nolint:dupl
func (d *DeployManager) GetAgentsTestContract(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *agentstestcontract.AgentsTestContractRef) {
	d.T().Helper()

	return manager.GetContract[*agentstestcontract.AgentsTestContractRef](ctx, d.T(), d, backend, AgentsTestContractType)
}

// GetTestClient gets the test client.
// nolint:dupl
func (d *DeployManager) GetTestClient(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testclient.TestClientRef) {
	d.T().Helper()

	return manager.GetContract[*testclient.TestClientRef](ctx, d.T(), d, backend, TestClientType)
}

// GetPingPongClient gets the ping pong test client.
// nolint:dupl
func (d *DeployManager) GetPingPongClient(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *pingpongclient.PingPongClientRef) {
	d.T().Helper()

	return manager.GetContract[*pingpongclient.PingPongClientRef](ctx, d.T(), d, backend, PingPongClientType)
}

// InitializeBondingManagerHarnessContract handles initializing the bonding manager harness contract on the "SYN" chain.
// nolint:dupl,cyclop
func (d *DeployManager) InitializeBondingManagerHarnessContract(
	ctx context.Context,
	synChainBackend backends.SimulatedTestBackend) error {
	d.T().Helper()

	deployedOriginHarness, _ := d.GetOriginHarness(ctx, synChainBackend)
	deployedDestinationHarness, _ := d.GetDestinationHarness(ctx, synChainBackend)
	deployedSummitHarness, _ := d.GetSummitHarness(ctx, synChainBackend)
	deployedInbox, inboxContract := d.GetInbox(ctx, synChainBackend)

	deployedBondingManagerHarness, bondingManagerHarnessContract := d.GetBondingManagerHarness(ctx, synChainBackend)
	_, summitHarnessContract := d.GetSummitHarness(ctx, synChainBackend)
	summitHarnessOwnerPtr, err := summitHarnessContract.SummitHarnessCaller.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get bonding manager harness: %w", err)
	}
	summitHarnessOwnerAuth := synChainBackend.GetTxContext(ctx, &summitHarnessOwnerPtr)

	initializeBondingManagerHarnessTx, err := bondingManagerHarnessContract.Initialize(
		summitHarnessOwnerAuth.TransactOpts,
		deployedOriginHarness.Address(),
		deployedDestinationHarness.Address(),
		deployedInbox.Address(),
		deployedSummitHarness.Address())
	if err != nil {
		return fmt.Errorf("could not initialize bonding manager harness: %w", err)
	}
	synChainBackend.WaitForConfirmation(ctx, initializeBondingManagerHarnessTx)

	initializeInboxTx, err := inboxContract.Initialize(
		summitHarnessOwnerAuth.TransactOpts,
		deployedBondingManagerHarness.Address(),
		deployedOriginHarness.Address(),
		deployedDestinationHarness.Address(),
		deployedSummitHarness.Address())
	if err != nil {
		return fmt.Errorf("could not initialize inbox: %w", err)
	}
	synChainBackend.WaitForConfirmation(ctx, initializeInboxTx)

	return nil
}

// AddAgentsToBondingManagerHarnessContract handles adding the agents to the bonding manager harness contract on the "SYN" chain.
// nolint:dupl,cyclop
func (d *DeployManager) AddAgentsToBondingManagerHarnessContract(
	ctx context.Context,
	synChainBackend backends.SimulatedTestBackend,
	agents []common.Address,
	domains []uint32) ([32]byte, [][][32]byte, []bondingmanagerharness.AgentStatus, error) {
	d.T().Helper()

	if len(agents) != len(domains) {
		return [32]byte{}, [][][32]byte{}, []bondingmanagerharness.AgentStatus{}, fmt.Errorf("agents and domains not same length")
	}

	_, bondingManagerHarnessContract := d.GetBondingManagerHarness(ctx, synChainBackend)

	bondingManagerHarnessOwnerPtr, err := bondingManagerHarnessContract.BondingManagerHarnessCaller.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return [32]byte{}, [][][32]byte{}, []bondingmanagerharness.AgentStatus{}, fmt.Errorf("could not get bonding manager harness: %w", err)
	}
	bondingManagerHarnessOwnerAuth := synChainBackend.GetTxContext(ctx, &bondingManagerHarnessOwnerPtr)

	for i := range agents {
		agent := agents[i]
		domain := domains[i]

		agentProof, err := bondingManagerHarnessContract.GetProof(&bind.CallOpts{Context: ctx}, agent)
		if err != nil {
			return [32]byte{}, [][][32]byte{}, []bondingmanagerharness.AgentStatus{}, fmt.Errorf("could not get agent proof from bonding manager harness: %w", err)
		}
		txAddAgent, err := bondingManagerHarnessContract.AddAgent(bondingManagerHarnessOwnerAuth.TransactOpts, domain, agent, agentProof)
		if err != nil {
			return [32]byte{}, [][][32]byte{}, []bondingmanagerharness.AgentStatus{}, fmt.Errorf("could not add agent to bonding manager harness: %w", err)
		}
		synChainBackend.WaitForConfirmation(ctx, txAddAgent)
	}

	bondingManagerHarnessAgentRoot, err := bondingManagerHarnessContract.AgentRoot(&bind.CallOpts{Context: ctx})
	if err != nil {
		return [32]byte{}, [][][32]byte{}, []bondingmanagerharness.AgentStatus{}, fmt.Errorf("could not get bonding manager agent root: %w", err)
	}

	agentProofs := make([][][32]byte, len(agents))
	agentStatuses := make([]bondingmanagerharness.AgentStatus, len(agents))
	for i := range agents {
		agent := agents[i]
		agentStatus, err := bondingManagerHarnessContract.AgentStatus(&bind.CallOpts{Context: ctx}, agent)
		if err != nil {
			return [32]byte{}, [][][32]byte{}, []bondingmanagerharness.AgentStatus{}, fmt.Errorf("could not get agent status from bonding manager harness: %w", err)
		}
		agentStatuses[i] = agentStatus

		agentProof, err := bondingManagerHarnessContract.GetProof(&bind.CallOpts{Context: ctx}, agent)
		if err != nil {
			return [32]byte{}, [][][32]byte{}, []bondingmanagerharness.AgentStatus{}, fmt.Errorf("could not get agent proof from bonding manager harness: %w", err)
		}
		agentProofs[i] = agentProof
	}

	return bondingManagerHarnessAgentRoot, agentProofs, agentStatuses, nil
}

// InitializeBondingManagerContract handles initializing the bonding manager contract on the "syn" chain.
// nolint:dupl,cyclop
func (d *DeployManager) InitializeBondingManagerContract(
	ctx context.Context,
	synChainBackend backends.SimulatedTestBackend) error {
	d.T().Helper()

	deployedOrigin, _ := d.GetOrigin(ctx, synChainBackend)
	deployedDestination, _ := d.GetDestination(ctx, synChainBackend)
	deployedSummit, _ := d.GetSummit(ctx, synChainBackend)
	deployedInbox, inboxContract := d.GetInbox(ctx, synChainBackend)

	deployedBondingManager, bondingManagerContract := d.GetBondingManager(ctx, synChainBackend)
	bondingManagerOwnerPtr, err := bondingManagerContract.BondingManagerCaller.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get bonding manager: %w", err)
	}
	bondingManagerOwnerAuth := synChainBackend.GetTxContext(ctx, &bondingManagerOwnerPtr)

	initializeBondingManagerTx, err := bondingManagerContract.Initialize(
		bondingManagerOwnerAuth.TransactOpts,
		deployedOrigin.Address(),
		deployedDestination.Address(),
		deployedInbox.Address(),
		deployedSummit.Address())
	if err != nil {
		return fmt.Errorf("could not initialize bonding manager: %w", err)
	}
	synChainBackend.WaitForConfirmation(ctx, initializeBondingManagerTx)

	initializeInboxTx, err := inboxContract.Initialize(
		bondingManagerOwnerAuth.TransactOpts,
		deployedBondingManager.Address(),
		deployedOrigin.Address(),
		deployedDestination.Address(),
		deployedSummit.Address())
	if err != nil {
		return fmt.Errorf("could not initialize inbox: %w", err)
	}
	synChainBackend.WaitForConfirmation(ctx, initializeInboxTx)

	return nil
}

// AddAgentsToBondingManagerContract handles adding the agents to the bonding manager contract on the "SYN" chain.
// nolint:dupl,cyclop
func (d *DeployManager) AddAgentsToBondingManagerContract(
	ctx context.Context,
	synChainBackend backends.SimulatedTestBackend,
	agents []common.Address,
	domains []uint32) ([32]byte, [][][32]byte, []bondingmanager.AgentStatus, error) {
	d.T().Helper()

	if len(agents) != len(domains) {
		return [32]byte{}, [][][32]byte{}, []bondingmanager.AgentStatus{}, fmt.Errorf("agents and domains not same length")
	}

	_, bondingManagerContract := d.GetBondingManager(ctx, synChainBackend)
	bondingManagerOwnerPtr, err := bondingManagerContract.BondingManagerCaller.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return [32]byte{}, [][][32]byte{}, []bondingmanager.AgentStatus{}, fmt.Errorf("could not get bonding manager: %w", err)
	}
	bondingManagerOwnerAuth := synChainBackend.GetTxContext(ctx, &bondingManagerOwnerPtr)

	var proof [][32]byte
	for i := range agents {
		agent := agents[i]
		domain := domains[i]

		txAddAgent, err := bondingManagerContract.AddAgent(bondingManagerOwnerAuth.TransactOpts, domain, agent, proof)
		if err != nil {
			return [32]byte{}, [][][32]byte{}, []bondingmanager.AgentStatus{}, fmt.Errorf("could not add agent to bonding manager: %w", err)
		}
		synChainBackend.WaitForConfirmation(ctx, txAddAgent)
	}

	bondingManagerAgentRoot, err := bondingManagerContract.AgentRoot(&bind.CallOpts{Context: ctx})
	if err != nil {
		return [32]byte{}, [][][32]byte{}, []bondingmanager.AgentStatus{}, fmt.Errorf("could not get bonding manager agent root: %w", err)
	}

	agentProofs := make([][][32]byte, len(agents))
	agentStatuses := make([]bondingmanager.AgentStatus, len(agents))
	for i := range agents {
		agent := agents[i]
		agentStatus, err := bondingManagerContract.AgentStatus(&bind.CallOpts{Context: ctx}, agent)
		if err != nil {
			return [32]byte{}, [][][32]byte{}, []bondingmanager.AgentStatus{}, fmt.Errorf("could not get agent status from bonding manager: %w", err)
		}
		agentStatuses[i] = agentStatus

		agentProof, err := bondingManagerContract.GetProof(&bind.CallOpts{Context: ctx}, agent)
		if err != nil {
			return [32]byte{}, [][][32]byte{}, []bondingmanager.AgentStatus{}, fmt.Errorf("could not get agent proof from bonding manager: %w", err)
		}
		agentProofs[i] = agentProof
	}

	return bondingManagerAgentRoot, agentProofs, agentStatuses, nil
}

// InitializeRemoteDeployedHarnessContracts handles initializing the harness contracts for light manager harness and destination harness on remote chain.
// nolint:dupl,cyclop
func (d *DeployManager) InitializeRemoteDeployedHarnessContracts(
	ctx context.Context,
	backend backends.SimulatedTestBackend,
	bondingManagerHarnessAgentRoot [32]byte) error {
	d.T().Helper()

	deployedDestinationHarness, destinationHarnessContract := d.GetDestinationHarness(ctx, backend)

	deployedLightInbox, lightInboxContract := d.GetLightInbox(ctx, backend)
	deployedLightManagerHarness, lightManagerHarnessContract := d.GetLightManagerHarness(ctx, backend)
	deployedOriginHarness, originHarnessContract := d.GetOriginHarness(ctx, backend)
	originHarnessOwnerPtr, err := originHarnessContract.OriginHarnessCaller.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get origin harness: %w", err)
	}
	originHarnessOwnerAuth := backend.GetTxContext(ctx, &originHarnessOwnerPtr)

	initializeLightManagerHarnessTx, err := lightManagerHarnessContract.Initialize(
		originHarnessOwnerAuth.TransactOpts,
		deployedOriginHarness.Address(),
		deployedDestinationHarness.Address(),
		deployedLightInbox.Address())
	if err != nil {
		return fmt.Errorf("could not initialize light manager harness: %w", err)
	}
	backend.WaitForConfirmation(ctx, initializeLightManagerHarnessTx)

	/*destinationHarnessOwnerPtr, err := destinationHarnessContract.DestinationHarnessCaller.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get destination harness: %w", err)
	}
	destinationHarnessOwnerAuth := backend.GetTxContext(ctx, &destinationHarnessOwnerPtr)*/

	initializeDestinationHarnessTx, err := destinationHarnessContract.Initialize(originHarnessOwnerAuth.TransactOpts, bondingManagerHarnessAgentRoot)
	if err != nil {
		return fmt.Errorf("could not initialize destination harness: %w", err)
	}
	backend.WaitForConfirmation(ctx, initializeDestinationHarnessTx)

	initializeLightInboxTx, err := lightInboxContract.Initialize(
		originHarnessOwnerAuth.TransactOpts,
		deployedLightManagerHarness.Address(),
		deployedOriginHarness.Address(),
		deployedDestinationHarness.Address())
	if err != nil {
		return fmt.Errorf("could not initialize light inbox: %w", err)
	}
	backend.WaitForConfirmation(ctx, initializeLightInboxTx)

	return nil
}

// AddAgentsToLightManagerHarnessContract handles adding the agents to the light manager harness contract on the remote chain.
// nolint:dupl,cyclop
// codebeat:disable[CYCLO]
func (d *DeployManager) AddAgentsToLightManagerHarnessContract(
	ctx context.Context,
	backend backends.SimulatedTestBackend,
	agents []common.Address,
	agentProofs [][][32]byte,
	agentStatuses []bondingmanagerharness.AgentStatus) error {
	d.T().Helper()

	if len(agents) != len(agentProofs) {
		return fmt.Errorf("agents and agentProofs not same length")
	}

	_, lightManagerHarnessContract := d.GetLightManagerHarness(ctx, backend)
	lightManagerHarnessOwnerPtr, err := lightManagerHarnessContract.LightManagerHarnessCaller.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get light manager harness: %w", err)
	}
	lightManagerHarnessOwnerAuth := backend.GetTxContext(ctx, &lightManagerHarnessOwnerPtr)

	for i := range agents {
		agent := agents[i]
		agentProof := agentProofs[i]
		bondingManagerAgentStatus := agentStatuses[i]

		lightManagerAgentStatus := lightmanagerharness.AgentStatus{
			Flag:   bondingManagerAgentStatus.Flag,
			Domain: bondingManagerAgentStatus.Domain,
			Index:  bondingManagerAgentStatus.Index,
		}

		// We want to make the notary do the work of adding the agent and not
		// have it done automatically by the test harness
		if lightManagerAgentStatus.Domain != 0 {
			continue
		}

		txAddAgent, err := lightManagerHarnessContract.UpdateAgentStatus(
			lightManagerHarnessOwnerAuth.TransactOpts,
			agent,
			lightManagerAgentStatus,
			agentProof)
		if err != nil {
			return fmt.Errorf("could not add agent to light manager harness: %w", err)
		}
		backend.WaitForConfirmation(ctx, txAddAgent)
	}

	return nil
}

// InitializeRemoteDeployedContracts handles initializing the contracts for light manager and destination on remote chain.
// nolint:dupl,cyclop
func (d *DeployManager) InitializeRemoteDeployedContracts(
	ctx context.Context,
	backend backends.SimulatedTestBackend,
	bondingManagerAgentRoot [32]byte) error {
	d.T().Helper()

	deployedOrigin, originContract := d.GetOrigin(ctx, backend)
	deployedDestination, destinationContract := d.GetDestination(ctx, backend)

	deployedLightInbox, lightInboxContract := d.GetLightInbox(ctx, backend)
	deployedLightManager, lightManagerContract := d.GetLightManager(ctx, backend)
	lightManagerOwnerPtr, err := lightManagerContract.LightManagerCaller.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get light manager: %w", err)
	}
	lightManagerOwnerAuth := backend.GetTxContext(ctx, &lightManagerOwnerPtr)

	initializeLightManagerTx, err := lightManagerContract.Initialize(
		lightManagerOwnerAuth.TransactOpts,
		deployedOrigin.Address(),
		deployedDestination.Address(),
		deployedLightInbox.Address())
	if err != nil {
		return fmt.Errorf("could not initialize light manager: %w", err)
	}
	backend.WaitForConfirmation(ctx, initializeLightManagerTx)

	originOwnerPtr, err := originContract.OriginCaller.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get origin: %w", err)
	}
	originOwnerAuth := backend.GetTxContext(ctx, &originOwnerPtr)

	initializeDestinationTx, err := destinationContract.Initialize(originOwnerAuth.TransactOpts, bondingManagerAgentRoot)
	if err != nil {
		return fmt.Errorf("could not initialize destination: %w", err)
	}
	backend.WaitForConfirmation(ctx, initializeDestinationTx)

	initializeLightInboxTx, err := lightInboxContract.Initialize(
		originOwnerAuth.TransactOpts,
		deployedLightManager.Address(),
		deployedOrigin.Address(),
		deployedDestination.Address())
	if err != nil {
		return fmt.Errorf("could not initialize light inbox: %w", err)
	}
	backend.WaitForConfirmation(ctx, initializeLightInboxTx)

	return nil
}

// AddAgentsToLightManagerContract handles adding the agents to the light manager contract on the remote chain.
// nolint:dupl,cyclop
func (d *DeployManager) AddAgentsToLightManagerContract(
	ctx context.Context,
	backend backends.SimulatedTestBackend,
	agents []common.Address,
	agentProofs [][][32]byte,
	agentStatuses []bondingmanager.AgentStatus) error {
	d.T().Helper()

	if len(agents) != len(agentProofs) {
		return fmt.Errorf("agents and agentProofs not same length")
	}

	_, lightManagerContract := d.GetLightManager(ctx, backend)
	lightManagerOwnerPtr, err := lightManagerContract.LightManagerCaller.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get light manager: %w", err)
	}
	lightManagerOwnerAuth := backend.GetTxContext(ctx, &lightManagerOwnerPtr)

	for i := range agents {
		agent := agents[i]
		agentProof := agentProofs[i]
		bondingManagerAgentStatus := agentStatuses[i]

		lightManagerAgentStatus := lightmanager.AgentStatus{
			Flag:   bondingManagerAgentStatus.Flag,
			Domain: bondingManagerAgentStatus.Domain,
			Index:  bondingManagerAgentStatus.Index,
		}

		txAddAgent, err := lightManagerContract.UpdateAgentStatus(
			lightManagerOwnerAuth.TransactOpts,
			agent,
			lightManagerAgentStatus,
			agentProof)
		if err != nil {
			return fmt.Errorf("could not add agent to light manager: %w", err)
		}
		backend.WaitForConfirmation(ctx, txAddAgent)
	}

	return nil
}

// LoadHarnessContractsOnChains loads the harness contracts to the various chains and initializes them.
// nolint:dupl,cyclop
func (d *DeployManager) LoadHarnessContractsOnChains(
	ctx context.Context,
	synChainBackend backends.SimulatedTestBackend,
	backends []backends.SimulatedTestBackend,
	agents []common.Address,
	agentDomains []uint32) error {
	d.T().Helper()

	err := d.InitializeBondingManagerHarnessContract(ctx, synChainBackend)
	if err != nil {
		return fmt.Errorf("could not initialize bonding manager harness on syn chain: %w", err)
	}

	bondingManagerHarnessAgentRoot, agentProofs, agentStatuses, err := d.AddAgentsToBondingManagerHarnessContract(
		ctx,
		synChainBackend,
		agents,
		agentDomains)

	if err != nil {
		return fmt.Errorf("could not add agents to bonding manager harness on syn chain: %w", err)
	}

	wg := sync.WaitGroup{}
	g, ctx := errgroup.WithContext(ctx)
	for _, b := range backends {
		backend := b
		wg.Add(1)
		g.Go(func() error {
			defer wg.Done()
			err := d.InitializeRemoteDeployedHarnessContracts(
				ctx,
				backend,
				bondingManagerHarnessAgentRoot)
			if err != nil {
				return fmt.Errorf("could not initialize remote deplyed harness contracts: %w", err)
			}

			err = d.AddAgentsToLightManagerHarnessContract(
				ctx,
				backend,
				agents,
				agentProofs,
				agentStatuses)
			if err != nil {
				return fmt.Errorf("could not add agents to remote light manager harness contract: %w", err)
			}
			return nil
		})
	}
	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not initialize remote deplyed harness contracts: %w", err)
	}

	return nil
}

// LoadContractsOnChains loads the contracts to the various chains and initializes them.
// nolint:dupl,cyclop
func (d *DeployManager) LoadContractsOnChains(
	ctx context.Context,
	synChainBackend backends.SimulatedTestBackend,
	backends []backends.SimulatedTestBackend,
	agents []common.Address,
	agentDomains []uint32) error {
	d.T().Helper()

	err := d.InitializeBondingManagerContract(ctx, synChainBackend)
	if err != nil {
		return fmt.Errorf("could not initialize bonding manager on syn chain: %w", err)
	}

	bondingManagerAgentRoot, agentProofs, agentStatuses, err := d.AddAgentsToBondingManagerContract(
		ctx,
		synChainBackend,
		agents,
		agentDomains)

	if err != nil {
		return fmt.Errorf("could not add agents to bonding manager on syn chain: %w", err)
	}

	for _, backend := range backends {
		err := d.InitializeRemoteDeployedContracts(
			ctx,
			backend,
			bondingManagerAgentRoot)
		if err != nil {
			return fmt.Errorf("could not initialize remote deplyed contracts: %w", err)
		}

		err = d.AddAgentsToLightManagerContract(
			ctx,
			backend,
			agents,
			agentProofs,
			agentStatuses)
		if err != nil {
			return fmt.Errorf("could not add agents to remote light manager contract: %w", err)
		}
	}

	return nil
}
