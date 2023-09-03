// Package manager manages deployers to make them as simple as possible
package manager

import (
	"context"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
)

// IDeployManager is responsible for deploying contracts.
type IDeployManager interface {
	// T returns the testing object.
	T() *testing.T
	// SetT sets the testing object.
	SetT(t *testing.T)
	// BulkDeploy deploys all contracts.
	BulkDeploy(ctx context.Context, testBackends []backends.SimulatedTestBackend, contracts ...contracts.ContractType)
	// GetContractRegistry returns the contract registry for the given backend.
	GetContractRegistry(backend backends.SimulatedTestBackend) deployer.ContractRegistry
	// Get returns the deployed contract for the given contract type.
	Get(ctx context.Context, backend backends.SimulatedTestBackend, contractType contracts.ContractType) contracts.DeployedContract
	// GetDeployedContracts returns all deployed contracts.
	GetDeployedContracts() (res map[uint32][]contracts.DeployedContract)
}

// DeployerManager is responsible for wrapping contract registry with easy to use getters that correctly cast the handles.
// since ContractRegistry is meant to be kept pure and go does not support generics, the sole function is to provide
// handler wrappers around the registry. This will no longer be required when go supports generics: https://blog.golang.org/generics-proposal
// TODO: go 1.20 supports generics, this can be improved.
type DeployerManager struct {
	// t is the testing object
	t *testing.T
	// registries stores the contract registries
	registries map[string]deployer.ContractRegistry
	// structMux prevents race conditions
	structMux sync.RWMutex
	// deployers adds a list of default deployers
	deployers []DeployerFunc
}

// DeployerFunc defines a deployer we can use.
type DeployerFunc func(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer

// NewDeployerManager creates a new deployment helper.
func NewDeployerManager(t *testing.T, deployers ...DeployerFunc) (d *DeployerManager) {
	t.Helper()
	d = &DeployerManager{
		t:          t,
		structMux:  sync.RWMutex{},
		registries: make(map[string]deployer.ContractRegistry),
		deployers:  deployers,
	}
	return d
}

// T is the testing object.
func (d *DeployerManager) T() *testing.T {
	return d.t
}

// SetT sets the testing object.
func (d *DeployerManager) SetT(t *testing.T) {
	t.Helper()
	d.t = t
}

// BulkDeploy synchronously deploys a bunch of contracts as quickly as possible to speed up tests.
// in a future version this will utilize dependency trees. Returns nothing when complete.
func (d *DeployerManager) BulkDeploy(ctx context.Context, testBackends []backends.SimulatedTestBackend, contracts ...contracts.ContractType) {
	wg := sync.WaitGroup{}
	for _, backend := range testBackends {
		wg.Add(1)
		go func(backend backends.SimulatedTestBackend) {
			defer wg.Done()
			cr := d.GetContractRegistry(backend)

			for _, contract := range contracts {
				cr.Get(ctx, contract)
			}
		}(backend)
	}
	wg.Wait()
}

// GetContractRegistry gets a contract registry for a backend and creates it if it does not exist.
func (d *DeployerManager) GetContractRegistry(backend backends.SimulatedTestBackend) deployer.ContractRegistry {
	d.structMux.Lock()
	defer d.structMux.Unlock()
	// if registry exists, return it
	contractRegistry, ok := d.registries[backend.GetBigChainID().String()]
	if ok {
		return contractRegistry
	}

	contractRegistry = deployer.NewContractRegistry(d.t, backend)

	for _, d := range d.deployers {
		contractRegistry.RegisterContractDeployer(d(contractRegistry, backend))
	}

	d.registries[backend.GetBigChainID().String()] = contractRegistry
	return contractRegistry
}

// Get gets the contract from the registry.
func (d *DeployerManager) Get(ctx context.Context, backend backends.SimulatedTestBackend, contractType contracts.ContractType) contracts.DeployedContract {
	return d.GetContractRegistry(backend).Get(ctx, contractType)
}

// GetDeployedContracts gets all deployed contracts by domain.
func (d *DeployerManager) GetDeployedContracts() (res map[uint32][]contracts.DeployedContract) {
	d.structMux.RLock()
	defer d.structMux.RUnlock()

	res = make(map[uint32][]contracts.DeployedContract)

	for _, registry := range d.registries {
		for _, contract := range registry.GetDeployedContracts() {
			chainID := uint32(contract.ChainID().Uint64())
			res[chainID] = append(res[chainID], contract)
		}
	}

	return res
}

// GetContract gets a contract from the registry and casts it to the correct type.
func GetContract[T any](ctx context.Context, tb testing.TB, deployManager IDeployManager, backend backends.SimulatedTestBackend, contractType contracts.ContractType) (contracts.DeployedContract, T) {
	tb.Helper()
	deployedContract := deployManager.Get(ctx, backend, contractType)
	contractHandle, ok := deployedContract.ContractHandle().(T)
	require.True(tb, ok)

	return deployedContract, contractHandle
}

// DeployManagerFactory is a factory for a deploy manager.
type DeployManagerFactory func() IDeployManager

// AssertDependenciesCorrect asserts that all dependencies of contracts are correct.
func AssertDependenciesCorrect(ctx context.Context, t *testing.T, deployManagerFactory DeployManagerFactory) {
	t.Helper()

	backend := simulated.NewSimulatedBackend(ctx, t)

	deployManager := deployManagerFactory()
	registeredContracts := deployManager.GetContractRegistry(backend).RegisteredDeployers()

	// test until all contacts are done
	for _, contract := range registeredContracts {
		deployManager = deployManagerFactory()
		contractRegistry := deployManager.GetContractRegistry(backend)
		assert.Equal(t, len(contractRegistry.GetDeployedContracts()), 0)

		// the contract is currently on the wrong backend, so we need to make it on the right backend
		dc := contractRegistry.Get(ctx, contract.ContractType())
		assert.Equal(t, dc.ChainID().String(), backend.GetBigChainID().String())

		deployedContracts := contractRegistry.GetDeployedContracts()
		// make sure dependency count is equal (adding our own contract to their expected amount)
		assert.Equal(t, len(deployedContracts), len(contract.Dependencies())+1)

		for _, dep := range contract.Dependencies() {
			_, hasDep := deployedContracts[dep.ID()]
			assert.True(t, hasDep)
		}
	}
}
