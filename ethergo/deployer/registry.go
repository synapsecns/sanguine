package deployer

import (
	"context"
	"github.com/neverlee/keymutex"
	"github.com/stretchr/testify/require"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"sync"
	"testing"
)

// ContractDeployer is a contract deployer for a single contract type.
type ContractDeployer interface {
	// Deploy deploys the contract and returns an error if it cannot be deployed
	Deploy(ctx context.Context) (contracts.DeployedContract, error)
	// ContractType gets the type of the deployed contract
	ContractType() contracts.ContractType
	// Dependencies gets the dependencies of this contract
	Dependencies() []contracts.ContractType
}

// GetOnlyContractRegistry is a contract registry that only allows gets.
type GetOnlyContractRegistry interface {
	// Get gets a contract by type. If the contract is not deployed, a new contract of type is deployed.
	// In cases where an error is present, this error is triggered via the test object in the constructor
	Get(ctx context.Context, contractType contracts.ContractType) contracts.DeployedContract
	// GetRegisteredDeployer gets the deployer for a given contract, returs nil if it doesn't exist
	GetRegisteredDeployer(contractType contracts.ContractType) ContractDeployer
	// GetDeployedContracts gets all deployed contracts in the registry.
	GetDeployedContracts() (res map[int]contracts.DeployedContract)
}

// ContractRegistry handles contract deployment/storage for a specific chain.
//
//go:generate go run github.com/vektra/mockery/v2 --name ContractRegistry --output ./mocks --case=underscore
type ContractRegistry interface {
	GetOnlyContractRegistry
	// Deploy deploys the contract type, but does not register it
	Deploy(ctx context.Context, contractType contracts.ContractType) contracts.DeployedContract
	// Register registers the contract with the contract registry. If you use Get() this isn't
	// required. This method is idempotent and will overwrite any contracts deployed
	Register(contractType contracts.ContractType, contract contracts.DeployedContract)
	// RegisterContractDeployer registers contract types that can be used for deployment. This allows extensibility by
	// non-synapse libraries. This will overwrite previous contract deployers with the same type
	RegisterContractDeployer(deployers ...ContractDeployer)
	// IsContractDeployed checks if a contract is deplyoed yet
	IsContractDeployed(contractType contracts.ContractType) bool
	// RegisteredDeployers gets all deployers registered
	RegisteredDeployers() []ContractDeployer
}

// contractRegistryImpl handles registration/fetching of deployed contracts.
type contractRegistryImpl struct {
	// backend stores the backend of the contract registry
	backend backends.SimulatedTestBackend
	// t is the testing object
	tb testing.TB
	// structMux prevents race conditions in the contract deployer
	structMux sync.RWMutex
	// deployMutex is a keyed mutex that prevents race conditions while deploying contracts (e.g. two contracts are created/
	// returned on the same chain). Note this only runs on Getters
	//nolint: structcheck, unused
	deployMutex *keymutex.KeyMutex
	// deployers stores the contract deploers
	deployers map[int]ContractDeployer
	// deployedContracts are the deployed contracts
	deployedContracts map[int]contracts.DeployedContract
}

func (c *contractRegistryImpl) RegisteredDeployers() (deployers []ContractDeployer) {
	c.structMux.Lock()
	defer c.structMux.Unlock()
	for _, deployer := range c.deployers {
		deployers = append(deployers, deployer)
	}
	return deployers
}

func (c *contractRegistryImpl) IsContractDeployed(contractType contracts.ContractType) bool {
	c.structMux.RLock()
	defer c.structMux.RUnlock()
	_, hasContract := c.deployedContracts[contractType.ID()]
	return hasContract
}

// NewContractRegistry creates a new contract registry.
func NewContractRegistry(tb testing.TB, backend backends.SimulatedTestBackend) ContractRegistry {
	tb.Helper()

	return &contractRegistryImpl{
		backend:           backend,
		tb:                tb,
		structMux:         sync.RWMutex{},
		deployMutex:       keymutex.New(47),
		deployers:         make(map[int]ContractDeployer),
		deployedContracts: make(map[int]contracts.DeployedContract),
	}
}

// Get gets the deployed contract.
func (c *contractRegistryImpl) Get(ctx context.Context, contractType contracts.ContractType) contracts.DeployedContract {
	c.tb.Helper()

	// contractLock this deployedContract
	c.contractLock(contractType)
	defer c.contractUnlock(contractType)

	// check if the deployed contract exists
	c.structMux.RLock()
	deployedContract, hasContract := c.deployedContracts[contractType.ID()]
	c.structMux.RUnlock()

	// if it does return it
	if hasContract {
		return deployedContract
	}

	// if not, deploy a new one
	deployedContract = c.Deploy(ctx, contractType)

	// register it
	c.structMux.Lock()
	c.deployedContracts[contractType.ID()] = deployedContract
	c.structMux.Unlock()

	// and return the new contfract
	return deployedContract
}

func (c *contractRegistryImpl) GetDeployedContracts() (res map[int]contracts.DeployedContract) {
	c.structMux.RLock()
	defer c.structMux.RUnlock()

	return c.deployedContracts
}

// contractLock creates a contractLock from a contract type and locks the mutex.
func (c *contractRegistryImpl) contractLock(contractType contracts.ContractType) {
	c.deployMutex.LockID(uint(contractType.ID()))
}

// contractUnlock unlocks a contractLock based on the contract type.
func (c *contractRegistryImpl) contractUnlock(contractType contracts.ContractType) {
	c.deployMutex.UnlockID(uint(contractType.ID()))
}

func (c *contractRegistryImpl) Deploy(ctx context.Context, contractType contracts.ContractType) contracts.DeployedContract {
	c.tb.Helper()
	c.structMux.RLock()
	deploymentHandle := c.deployers[contractType.ID()]
	c.structMux.RUnlock()

	deployedContract, err := deploymentHandle.Deploy(ctx)
	require.Nil(c.tb, err)

	c.backend.WaitForConfirmation(ctx, deployedContract.DeployTx())
	err = c.backend.VerifyContract(contractType, deployedContract)
	if err != nil {
		logger.Warnf("got error %s while verifying contract, skipping", err)
	}
	logger.Debugf("added contract %s of types %s in tx %s on chain id %s", deployedContract.Address(), contractType.Name(), deployedContract.DeployTx().Hash().String(), deployedContract.ChainID().String())

	return deployedContract
}

func (c *contractRegistryImpl) Register(contractType contracts.ContractType, contract contracts.DeployedContract) {
	c.tb.Helper()
	// contractLock this deployedContract
	c.contractLock(contractType)
	c.structMux.Lock()
	defer c.contractUnlock(contractType)
	defer c.structMux.Unlock()

	c.deployedContracts[contractType.ID()] = contract
}

func (c *contractRegistryImpl) RegisterContractDeployer(deployers ...ContractDeployer) {
	c.structMux.Lock()
	defer c.structMux.Unlock()
	for _, contract := range deployers {
		c.deployers[contract.ContractType().ID()] = contract
	}
}

func (c *contractRegistryImpl) GetRegisteredDeployer(contractType contracts.ContractType) ContractDeployer {
	c.structMux.RLock()
	defer c.structMux.RUnlock()
	for contract, deployer := range c.deployers {
		if contractType.ID() == contract {
			return deployer
		}
	}
	return nil
}

var _ ContractRegistry = &contractRegistryImpl{}
