package testutil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/headerharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/messageharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/tipsharness"
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

// OriginHarnessDomain is the domain used for the origin harness.
const OriginHarnessDomain = 1

// OriginHarnessDeployer deploys the origin harness for testing.
type OriginHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewOriginHarnessDeployer deploys a new origin harness.
func NewOriginHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return OriginHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, OriginHarnessType)}
}

// Deploy deploys the origin harness.
func (o OriginHarnessDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	notaryManagerContract := o.Registry().Get(ctx, NotaryManagerType)

	return o.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		address, tx, rawHandle, err := originharness.DeployOriginHarness(transactOps, backend, OriginHarnessDomain)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s: %w", o.ContractType().ContractName(), err)
		}
		o.Backend().WaitForConfirmation(ctx, tx)

		initializeOpts := o.Backend().GetTxContext(ctx, &transactOps.From)
		initializeTx, err := rawHandle.Initialize(initializeOpts.TransactOpts, notaryManagerContract.Address())
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize notary (%s) on %s: %w", transactOps.From, o.ContractType().ContractName(), err)
		}
		o.Backend().WaitForConfirmation(ctx, initializeTx)

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return originharness.NewOriginHarnessRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the origin contract.
func (o OriginHarnessDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{NotaryManagerType}
}

// AttestationHarnessDeployer deploys the attestation harness.
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
func (d DestinationHarnessDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return destinationharness.DeployDestinationHarness(transactOps, backend, uint32(d.Backend().GetChainID()))
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return destinationharness.NewDestinationHarnessRef(address, backend)
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
