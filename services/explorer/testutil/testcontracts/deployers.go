package testcontracts

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap/testswap"
)

// TestSynapseBridgeDeployer is the type of the test bridge deployer.
type TestSynapseBridgeDeployer struct {
	*deployer.BaseDeployer
}

// NewTestSynapseBridgeDeployer creates a new test bridge deployer.
func NewTestSynapseBridgeDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TestSynapseBridgeDeployer{deployer.NewSimpleDeployer(registry, backend, TestSynapseBridgeType)}
}

// Deploy deploys a test bridge
func (t TestSynapseBridgeDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return t.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testbridge.DeploySynapseBridge(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testbridge.NewBridgeRef(address, backend)
	})
}

var _ deployer.ContractDeployer = &TestSynapseBridgeDeployer{}

// TestSwapFlashLoanDeployer is the type of the test swap deployer.
type TestSwapFlashLoanDeployer struct {
	*deployer.BaseDeployer
}

// NewTestSwapFlashLoanDeployer creates a new test swap deployer.
func NewTestSwapFlashLoanDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TestSwapFlashLoanDeployer{deployer.NewSimpleDeployer(registry, backend, TestSynapseBridgeType)}
}

// Deploy deploys a test swap
func (t TestSwapFlashLoanDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return t.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testswap.DeployTestSwapFlashLoan(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testswap.NewSwapRef(address, backend)
	})
}

var _ deployer.ContractDeployer = &TestSwapFlashLoanDeployer{}
