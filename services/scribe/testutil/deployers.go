package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract2"

	"github.com/synapsecns/sanguine/ethergo/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"
)

// TestContractDeployer deploys a test contract.
type TestContractDeployer struct {
	*deployer.BaseDeployer
}

// NewTestContractDeployer creates a new test contract deployer.
func NewTestContractDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TestContractDeployer{deployer.NewSimpleDeployer(registry, backend, TestContractType)}
}

// Deploy deploys the test contract.
func (t TestContractDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return t.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testcontract.DeployTestContract(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testcontract.NewTestContractRef(address, backend)
	})
}

// TestContract2Deployer deploys a test contract 2.
type TestContract2Deployer struct {
	*deployer.BaseDeployer
}

// NewTestContract2Deployer creates a new test contract deployer 2.
func NewTestContract2Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TestContract2Deployer{deployer.NewSimpleDeployer(registry, backend, TestContract2Type)}
}

// Deploy deploys the test contract 2.
func (t TestContract2Deployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return t.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testcontract2.DeployTestContract(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testcontract2.NewTestContractRef(address, backend)
	})
}
