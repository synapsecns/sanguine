package deployer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
)

// BaseDeployer is a basic deployment contract. It contains several utility functions including:
// - RecursiveDependencies: can be used to return dependents of dependents (since Dependencies should return all dependents: direct and indirect)
// - DeploySimpleContract: can be used to deploy contracts that don't have to be initialized.
type BaseDeployer struct {
	// registry is the contract registry
	registry GetOnlyContractRegistry
	// backend is the registry
	backend      backends.SimulatedTestBackend
	contractType contracts.ContractType
}

// NewSimpleDeployer creates a new base deployer.
func NewSimpleDeployer(registry GetOnlyContractRegistry, backend backends.SimulatedTestBackend, contractType contracts.ContractType) *BaseDeployer {
	return &BaseDeployer{
		registry:     registry,
		backend:      backend,
		contractType: contractType,
	}
}

// Deploy is a placeholder to ensure function inheritance. Calling this directly will panic.
func (n BaseDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	panic("deploy not implemented in base deployer")
}

// Backend gets the backend of the current deployer instance.
func (n BaseDeployer) Backend() backends.SimulatedTestBackend {
	return n.backend
}

// Registry gets the registry.
func (n BaseDeployer) Registry() GetOnlyContractRegistry {
	return n.registry
}

// ContractType returns the contract type.
func (n BaseDeployer) ContractType() contracts.ContractType {
	return n.contractType
}

// Dependencies returns dependencies for the contract - this should be overridden by base classes if there are dependencies.
func (n BaseDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{}
}

// RecursiveDependencies recursively get dependencies.
func (n BaseDeployer) RecursiveDependencies(dependencies []contracts.ContractType) (res []contracts.ContractType) {
	// check if dependency is already in result
	resultHasDependency := func(dep contracts.ContractType) bool {
		for _, dependency := range res {
			if dep.ID() == dependency.ID() {
				return true
			}
		}
		return false
	}
	// TODO handle cyclicality
	for _, dependency := range dependencies {
		codependencies := append(n.registry.GetRegisteredDeployer(dependency).Dependencies(), dependency)
		for _, dependency := range codependencies {
			if !resultHasDependency(dependency) {
				res = append(res, dependency)
			}
		}
	}
	return res
}

// DeployFunc is the deploy function.
type DeployFunc func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error)

// HandleFunc gets the ref function.
type HandleFunc func(address common.Address, backend bind.ContractBackend) (interface{}, error)

// DeploySimpleContract handles no dependency contract deployments.
// All others must be handled in inheriting structs.
func (n BaseDeployer) DeploySimpleContract(ctx context.Context, deployFunction DeployFunc, handleFunction HandleFunc) (contracts.DeployedContract, error) {
	auth := n.backend.GetTxContext(ctx, nil)
	tmpAddress, tx, _, err := deployFunction(auth.TransactOpts, n.backend)
	if err != nil {
		return nil, fmt.Errorf("could not deploy contract of type %s: %w", n.contractType.ContractName(), err)
	}

	handle, err := handleFunction(tmpAddress, n.backend)
	if err != nil {
		return nil, fmt.Errorf("could not get handke: %w", err)
	}

	return &DeployedContract{
		address:        tmpAddress,
		contractHandle: handle,
		owner:          auth.From,
		chainID:        n.backend.GetBigChainID(),
		deployTx:       tx,
	}, nil
}
