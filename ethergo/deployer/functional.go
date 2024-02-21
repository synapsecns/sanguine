package deployer

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
)

// FunctionalDeployer is a new functional deployer.
//
// It is currently experimental and seeks to succeed SimpleDeployer.
// It has the advantage of being pureley functional and not requiring the user to create a struct.
// It currently extends BaseDeployer.
//
// Deprecated: TODO remove this, it looks terrible.
type FunctionalDeployer struct {
	*BaseDeployer
	deployFunc   func(ctx context.Context) (contracts.DeployedContract, error)
	dependencies []contracts.ContractType
}

// IFunctionalDeployer is an interface for a functional deployer.
type IFunctionalDeployer interface {
	Registry() GetOnlyContractRegistry
	Backend() backends.SimulatedTestBackend
	ContractType() contracts.ContractType
}

// DeployerFunc defines a deployer we can use.
type DeployerFunc func(registry GetOnlyContractRegistry, backend backends.SimulatedTestBackend) ContractDeployer

// FunctionalDeployFunc is a function that deploys a contract.
type FunctionalDeployFunc func(ctx context.Context, helpers IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error)

// WrapFunc is a function that wraps a contract.
type WrapFunc func(address common.Address, backend bind.ContractBackend) (interface{}, error)

// NewFunctionalDeployer creates a new functional deployer.
// Deprecated: TODO remove this
func NewFunctionalDeployer(contractType contracts.ContractType, deployFunc FunctionalDeployFunc,
	wrapFunc WrapFunc, autoRecursedDeps []contracts.ContractType) DeployerFunc {

	return func(registry GetOnlyContractRegistry, backend backends.SimulatedTestBackend) ContractDeployer {
		baseDeployer := NewSimpleDeployer(registry, backend, contractType)
		functionalDeployer := &FunctionalDeployer{
			BaseDeployer: baseDeployer,
		}

		functionalDeployer.deployFunc = func(ctx context.Context) (contracts.DeployedContract, error) {
			// nolint: errcheck
			return baseDeployer.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
				// nolint: errcheck

				return deployFunc(ctx, functionalDeployer, transactOps, backend)
			}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
				// nolint: errcheck
				return wrapFunc(address, backend)
			})
		}

		functionalDeployer.dependencies = autoRecursedDeps
		return functionalDeployer
	}
}

// Deploy deploys the contract.
func (f *FunctionalDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return f.deployFunc(ctx)
}

// Dependencies returns the dependencies of the functional deployer.
func (f *FunctionalDeployer) Dependencies() []contracts.ContractType {
	return f.RecursiveDependencies(f.dependencies)
}
