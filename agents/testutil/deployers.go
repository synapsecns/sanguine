package testutil

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/notarymanager"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/synapse-node/testutils/backends"
)

// OriginDeployer deploys the origin contract.
type OriginDeployer struct {
	*deployer.BaseDeployer
}

// NewOriginDeployer deploys the origin contract.
func NewOriginDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return OriginDeployer{deployer.NewSimpleDeployer(registry, backend, OriginType)}
}

// Deploy deploys the origin contract.
func (d OriginDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	notaryManagerContract := d.Registry().Get(ctx, NotaryManagerType)

	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		// deploy the origin contract
		var rawHandle *origin.Origin
		address, tx, rawHandle, err = origin.DeployOrigin(transactOps, backend, uint32(d.Backend().GetChainID()))
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s: %w", d.ContractType().ContractName(), err)
		}

		// initialize the origin contract
		initializationTx, err := rawHandle.Initialize(transactOps, notaryManagerContract.Address())
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize contract: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, initializationTx)

		// get the owner of the notary manager contract
		notaryTransactOps := d.Backend().GetTxContext(ctx, notaryManagerContract.OwnerPtr())

		// set the notary contract on the notary manager
		updateManager, ok := notaryManagerContract.ContractHandle().(*notarymanager.NotaryManagerRef)
		if !ok {
			return common.Address{}, nil, nil, fmt.Errorf("could not update contract: %w", err)
		}

		setTx, err := updateManager.SetOrigin(notaryTransactOps.TransactOpts, address)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not set origin: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, setTx)

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return origin.NewOriginRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the origin contract.
func (d OriginDeployer) Dependencies() []deployer.ContractType {
	return []deployer.ContractType{NotaryManagerType}
}

// NotaryManagerDeployer deploys the update manager.
type NotaryManagerDeployer struct {
	*deployer.BaseDeployer
}

// NewNotaryManagerDeployer deploys a new notary manager.
func NewNotaryManagerDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return NotaryManagerDeployer{deployer.NewSimpleDeployer(registry, backend, NotaryManagerType)}
}

// Deploy deploys the notary contract.
func (n NotaryManagerDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return notarymanager.DeployNotaryManager(transactOps, backend, transactOps.From)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return notarymanager.NewNotaryManagerRef(address, backend)
	})
}

// AttestationCollectorDeployer deploys the attestation collector.
type AttestationCollectorDeployer struct {
	*deployer.BaseDeployer
}

// NewAttestationCollectorDeployer creates the deployer for  the attestation collecotr.
func NewAttestationCollectorDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return AttestationCollectorDeployer{deployer.NewSimpleDeployer(registry, backend, AttestationCollectorType)}
}

// Deploy deploys the attestation collector.
func (a AttestationCollectorDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return a.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		attestationAddress, attestationTx, collector, err := attestationcollector.DeployAttestationCollector(transactOps, backend)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy attestation collector: %w", err)
		}

		auth := a.Backend().GetTxContext(ctx, &transactOps.From)
		initTx, err := collector.Initialize(auth.TransactOpts)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize attestation collector: %w", err)
		}
		a.Backend().WaitForConfirmation(ctx, initTx)

		return attestationAddress, attestationTx, collector, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return attestationcollector.NewAttestationCollectorRef(address, backend)
	})
}

// DestinationDeployer deploys the destination.
type DestinationDeployer struct {
	*deployer.BaseDeployer
}

// NewDestinationDeployer creates the deployer for the destination.
func NewDestinationDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return DestinationDeployer{deployer.NewSimpleDeployer(registry, backend, DestinationType)}
}

// Deploy deploys the destination.
func (d DestinationDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		destinationAddress, destinationTx, destination, err := destination.DeployDestination(transactOps, backend, uint32(d.Backend().GetChainID()))
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy destination: %w", err)
		}

		auth := d.Backend().GetTxContext(ctx, &transactOps.From)
		initTx, err := destination.Initialize(auth.TransactOpts, uint32(d.Registry().Get(ctx, OriginType).ChainID().Uint64()), common.Address{})
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize destination: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, initTx)

		return destinationAddress, destinationTx, destination, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return destination.NewDestinationRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the destination contract.
func (d DestinationDeployer) Dependencies() []deployer.ContractType {
	return []deployer.ContractType{OriginType, NotaryManagerType}
}
