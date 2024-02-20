package testutil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/committee/contracts/mocks/noopinterchain"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"testing"
)

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

// NewDeployManager creates a new DeployManager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	return &DeployManager{manager.NewDeployerManager(t, NewSynapseModuleDeployer, NewNoOpInterchainDeployer)}
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
		smAddress, smTx, smI, err := synapsemodule.DeploySynapseModule(transactOps, backend)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy SynapseModule: %w", err)
		}

		interchainContract := s.Registry().Get(ctx, NoOpInterchainType)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not get NoOpInterchain contract: %w", err)
		}

		contractRef, err := synapsemodule.NewSynapseModuleRef(smAddress, backend)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not create SynapseModuleRef: %w", err)
		}

		auth := s.Backend().GetTxContext(ctx, &transactOps.From)
		tx, err := contractRef.SetInterchainDB(auth.TransactOpts, interchainContract.Address())
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not set interchain contract: %w", err)
		}

		s.Backend().WaitForConfirmation(ctx, tx)

		return smAddress, smTx, smI, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return synapsemodule.NewSynapseModuleRef(address, backend)
	})
}

// Dependencies returns the dependencies of the SynapseModuleDeployer.
func (s SynapseModuleDeployer) Dependencies() []contracts.ContractType {
	return s.RecursiveDependencies([]contracts.ContractType{NoOpInterchainType})
}

// NoOpInterchainDeployer wraps NoOpInterchainDeployer and allows typed contract handles to be returned.
type NoOpInterchainDeployer struct {
	*deployer.BaseDeployer
}

// NewNoOpInterchainDeployer deploys a NoOpInterchain contract.
func NewNoOpInterchainDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return &NoOpInterchainDeployer{deployer.NewSimpleDeployer(registry, backend, NoOpInterchainType)}
}

// Deploy deploys a NoOpInterchain contract.
func (i NoOpInterchainDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return i.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return noopinterchain.DeployNoOpInterchain(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return noopinterchain.NewNoOpInterchainRef(address, backend)
	})
}
