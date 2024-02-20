package testutil

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/committee/contracts/interchaindb"
	"github.com/synapsecns/sanguine/committee/contracts/mocks/gasoraclemock"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/manager"
)

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

// NewDeployManager creates a new DeployManager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	return &DeployManager{manager.NewDeployerManager(t, NewSynapseModuleDeployer, NewInterchainDBDeployer, NewGasOracleDeployer)}
}

// SynapseModuleDeployer wraps SynapseModuleDeployer and allows typed contract handles to be returned.
type SynapseModuleDeployer struct {
	*deployer.BaseDeployer
}

// NewSynapseModuleDeployer deploys a SynapseModule contract.
func NewSynapseModuleDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return &SynapseModuleDeployer{deployer.NewSimpleDeployer(registry, backend, SynapseModuleType)}
}

// Deploy deploys a SynapseModule contract.
// this is deployed with the NoOpInterchain contract as the interchain contract.
// this can be overridden at any type by owner.
func (s SynapseModuleDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return s.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		interchainContract := s.Registry().Get(ctx, InterchainDB)

		// Deployed with NoOpInterchainDB as INTERCHAIN_DB and deployer address as owner
		smAddress, smTx, smI, err := synapsemodule.DeploySynapseModule(transactOps, backend, interchainContract.Address(), transactOps.From)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy SynapseModule: %w", err)
		}

		auth := s.Backend().GetTxContext(ctx, &transactOps.From)
		gasOracle := s.Registry().Get(ctx, GasOracleMockType)
		setTx, err := smI.SetGasOracle(auth.TransactOpts, gasOracle.Address())
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not set gas oracle: %w", err)
		}
		s.Backend().WaitForConfirmation(ctx, setTx)

		return smAddress, smTx, smI, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return synapsemodule.NewSynapseModuleRef(address, backend)
	})
}

// Dependencies returns the dependencies of the SynapseModuleDeployer.
func (s SynapseModuleDeployer) Dependencies() []contracts.ContractType {
	return s.RecursiveDependencies([]contracts.ContractType{InterchainDB, GasOracleMockType})
}

// InterchainDBDeployer wraps InterchainDBDeployer and allows typed contract handles to be returned.
type InterchainDBDeployer struct {
	*deployer.BaseDeployer
}

// NewInterchainDBDeployer deploys a NoOpInterchain contract.
func NewInterchainDBDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return &InterchainDBDeployer{deployer.NewSimpleDeployer(registry, backend, InterchainDB)}
}

// Deploy deploys a NoOpInterchain contract.
func (i InterchainDBDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return i.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return interchaindb.DeployInterchainDB(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return interchaindb.NewInterchainDBRef(address, backend)
	})
}

// GasOracleDeployer wraps GasOracleDeployer and allows typed contract handles to be returned.
type GasOracleDeployer struct {
	*deployer.BaseDeployer
}

// NewGasOracleDeployer deploys a GasOracle contract.
func NewGasOracleDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return &GasOracleDeployer{deployer.NewSimpleDeployer(registry, backend, GasOracleMockType)}
}

// Deploy deploys a GasOracle contract.
func (i GasOracleDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return i.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return gasoraclemock.DeployGasOracleMock(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return gasoraclemock.NewGasOracleMockRef(address, backend)
	})
}
