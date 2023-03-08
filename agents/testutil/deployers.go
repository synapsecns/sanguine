package testutil

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/ethergo/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/deployer"
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
func (d OriginDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		// deploy the origin contract
		var rawHandle *origin.Origin
		address, tx, rawHandle, err = origin.DeployOrigin(transactOps, backend, uint32(d.Backend().GetChainID()))
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s: %w", d.ContractType().ContractName(), err)
		}
		d.Backend().WaitForConfirmation(ctx, tx)

		// initialize the origin contract
		initializationTx, err := rawHandle.Initialize(transactOps)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize contract: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, initializationTx)

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return origin.NewOriginRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the origin contract.
func (d OriginDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{}
}

// SummitDeployer deploys the summit.
type SummitDeployer struct {
	*deployer.BaseDeployer
}

// NewSummitDeployer creates the deployer for  the summit.
func NewSummitDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return SummitDeployer{deployer.NewSimpleDeployer(registry, backend, SummitType)}
}

// Deploy deploys the summit.
//
//nolint:dupword
func (a SummitDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return a.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		summitAddress, summitTx, summit, err := summit.DeploySummit(transactOps, backend, uint32(a.Backend().GetChainID()))
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy summit: %w", err)
		}

		auth := a.Backend().GetTxContext(ctx, &transactOps.From)
		initTx, err := summit.Initialize(auth.TransactOpts)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize attestation collector: %w", err)
		}
		a.Backend().WaitForConfirmation(ctx, initTx)

		return summitAddress, summitTx, summit, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return summit.NewSummitRef(address, backend)
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
func (d DestinationDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		destinationAddress, destinationTx, destination, err := destination.DeployDestination(transactOps, backend, uint32(d.Backend().GetChainID()))
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy destination: %w", err)
		}

		auth := d.Backend().GetTxContext(ctx, &transactOps.From)
		initTx, err := destination.Initialize(auth.TransactOpts)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize destination: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, initTx)

		// nolint:dupword
		/*setTx, err := destination.AddNotary(auth.TransactOpts, uint32(d.Registry().Get(ctx, OriginType).ChainID().Uint64()), common.Address{})
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not set notary: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, setTx)*/

		return destinationAddress, destinationTx, destination, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return destination.NewDestinationRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the destination contract.
func (d DestinationDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{OriginType}
}
