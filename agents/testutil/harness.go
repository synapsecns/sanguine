package testutil

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/agents/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/summitharness"

	"github.com/synapsecns/sanguine/agents/contracts/test/snapshotharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/stateharness"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/headerharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/messageharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
	"github.com/synapsecns/sanguine/agents/contracts/test/testclient"
	"github.com/synapsecns/sanguine/agents/contracts/test/tipsharness"
	"github.com/synapsecns/sanguine/agents/testutil/agentstestcontract"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
)

// MessageHarnessDeployer deploys the message harness for testing.
type MessageHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewMessageHarnessDeployer creates a message harness deployer.
func NewMessageHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MessageHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, MessageHarnessType)}
}

// Deploy deploys the message harness deployer.
func (d MessageHarnessDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return messageharness.DeployMessageHarness(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return messageharness.NewMessageHarnessRef(address, backend)
	})
}

// OriginHarnessDeployer deploys the origin harness for testing.
type OriginHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewOriginHarnessDeployer deploys a new origin harness.
func NewOriginHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return OriginHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, OriginHarnessType)}
}

// Deploy deploys the origin harness.
// nolint:dupl
func (o OriginHarnessDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return o.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		address, tx, rawHandle, err := originharness.DeployOriginHarness(transactOps, backend, uint32(o.Backend().GetChainID()))
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s: %w", o.ContractType().ContractName(), err)
		}
		o.Backend().WaitForConfirmation(ctx, tx)

		initializeOpts := o.Backend().GetTxContext(ctx, &transactOps.From)
		initializeTx, err := rawHandle.Initialize(initializeOpts.TransactOpts)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize origin (%s) on %s: %w", transactOps.From, o.ContractType().ContractName(), err)
		}
		o.Backend().WaitForConfirmation(ctx, initializeTx)

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return originharness.NewOriginHarnessRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the origin contract.
func (o OriginHarnessDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{}
}

// StateHarnessDeployer deploys the state harness.
type StateHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewStateHarnessDeployer creates a new deployer for the state harness.
func NewStateHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return StateHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, StateHarnessType)}
}

// Deploy deploys the state harness.
func (a StateHarnessDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return a.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return stateharness.DeployStateHarness(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return stateharness.NewStateHarnessRef(address, backend)
	})
}

// SnapshotHarnessDeployer deploys the snapshot harness.
type SnapshotHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewSnapshotHarnessDeployer creates a new deployer for the snapshot harness.
func NewSnapshotHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return SnapshotHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, SnapshotHarnessType)}
}

// Deploy deploys the snapshot harness.
func (a SnapshotHarnessDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return a.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return snapshotharness.DeploySnapshotHarness(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return snapshotharness.NewSnapshotHarnessRef(address, backend)
	})
}

// AttestationHarnessDeployer deploys the attestation harness for testing.
type AttestationHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewAttestationHarnessDeployer creates a new deployer for the attestation harness.
func NewAttestationHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return AttestationHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, AttestationHarnessType)}
}

// Deploy deploys the attestation harness.
func (a AttestationHarnessDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return a.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return attestationharness.DeployAttestationHarness(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return attestationharness.NewAttestationHarnessRef(address, backend)
	})
}

// TipsHarnessDeployer deploys the tip harness for tester.
type TipsHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewTipsHarnessDeployer creates a new deployer for the attestation harness.
func NewTipsHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TipsHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, TipsHarnessType)}
}

// Deploy deploys the attestation harness.
func (a TipsHarnessDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return a.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return tipsharness.DeployTipsHarness(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return tipsharness.NewTipsHarnessRef(address, backend)
	})
}

// DestinationHarnessDeployer deploys the destination harness.
type DestinationHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewDestinationHarnessDeployer creates a new deployer for the destination harness.
func NewDestinationHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return DestinationHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, DestinationHarnessType)}
}

// Deploy deploys the destination harness.
// nolint:dupl
func (d DestinationHarnessDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		address, tx, rawHandle, err := destinationharness.DeployDestinationHarness(transactOps, backend, uint32(d.Backend().GetChainID()))
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s: %w", d.ContractType().ContractName(), err)
		}
		d.Backend().WaitForConfirmation(ctx, tx)

		initializeOpts := d.Backend().GetTxContext(ctx, &transactOps.From)
		initializeTx, err := rawHandle.Initialize(initializeOpts.TransactOpts)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize destination (%s) on %s: %w", transactOps.From, d.ContractType().ContractName(), err)
		}
		d.Backend().WaitForConfirmation(ctx, initializeTx)

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return destinationharness.NewDestinationHarnessRef(address, backend)
	})
}

// SummitHarnessDeployer deploys the summit harness.
type SummitHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewSummitHarnessDeployer creates a new deployer for the summit harness.
func NewSummitHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return SummitHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, SummitHarnessType)}
}

// Deploy deploys the summit harness.
// nolint:dupl
func (d SummitHarnessDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		address, tx, rawHandle, err := summitharness.DeploySummitHarness(transactOps, backend)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s: %w", d.ContractType().ContractName(), err)
		}
		d.Backend().WaitForConfirmation(ctx, tx)

		initializeOpts := d.Backend().GetTxContext(ctx, &transactOps.From)
		initializeTx, err := rawHandle.Initialize(initializeOpts.TransactOpts)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize summit harness (%s) on %s: %w", transactOps.From, d.ContractType().ContractName(), err)
		}
		d.Backend().WaitForConfirmation(ctx, initializeTx)

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return summitharness.NewSummitHarnessRef(address, backend)
	})
}

// HeaderHarnessDeployer deploys the header harness.
type HeaderHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewHeaderHarnessDeployer gets the header harness.
func NewHeaderHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return HeaderHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, HeaderHarnessType)}
}

// Deploy deploys the header harness.
func (h HeaderHarnessDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return h.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return headerharness.DeployHeaderHarness(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return headerharness.NewHeaderHarnessRef(address, backend)
	})
}

// AgentsTestContractDeployer deploys the agents test contract.
type AgentsTestContractDeployer struct {
	*deployer.BaseDeployer
}

// NewAgentsTestContractDeployer gets the agents test contract.
func NewAgentsTestContractDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return AgentsTestContractDeployer{deployer.NewSimpleDeployer(registry, backend, AgentsTestContractType)}
}

// Deploy deploys the agents test contract.
func (h AgentsTestContractDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return h.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return agentstestcontract.DeployAgentsTestContract(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return agentstestcontract.NewAgentsTestContractRef(address, backend)
	})
}

// TestClientDeployer deploys the test client.
type TestClientDeployer struct {
	*deployer.BaseDeployer
}

// NewTestClientDeployer gets the test client.
func NewTestClientDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TestClientDeployer{deployer.NewSimpleDeployer(registry, backend, TestClientType)}
}

// Deploy deploys the test client.
func (h TestClientDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	originHarnessContract := h.Registry().Get(ctx, OriginHarnessType)
	destinationHarnessContract := h.Registry().Get(ctx, DestinationHarnessType)
	originAddress := originHarnessContract.Address()
	destinationAddress := destinationHarnessContract.Address()
	return h.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testclient.DeployTestClient(transactOps, backend, originAddress, destinationAddress)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testclient.NewTestClientRef(address, backend)
	})
}

// PingPongClientDeployer deploys the ping pong test client.
type PingPongClientDeployer struct {
	*deployer.BaseDeployer
}

// NewPingPongClientDeployer gets the ping pong test client.
func NewPingPongClientDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return PingPongClientDeployer{deployer.NewSimpleDeployer(registry, backend, PingPongClientType)}
}

// Deploy deploys the ping pong test client.
func (h PingPongClientDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	originHarnessContract := h.Registry().Get(ctx, OriginHarnessType)
	destinationHarnessContract := h.Registry().Get(ctx, DestinationHarnessType)
	originAddress := originHarnessContract.Address()
	destinationAddress := destinationHarnessContract.Address()
	return h.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return pingpongclient.DeployPingPongClient(transactOps, backend, originAddress, destinationAddress)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return pingpongclient.NewPingPongClientRef(address, backend)
	})
}
