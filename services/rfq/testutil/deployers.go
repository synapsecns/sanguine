package testutil

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/contracts/mockerc20"
	"testing"
)

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	// TODO: add contracts here
	parentManager := manager.NewDeployerManager(t, NewFastBridgeDeployer, NewMockERC20Deployer)
	return &DeployManager{parentManager}
}

// FastBridgeDeployer deplyos a fast bridge contract for testing.
type FastBridgeDeployer struct {
	*deployer.BaseDeployer
}

// NewFastBridgeDeployer deploys a fast bridge contract.
func NewFastBridgeDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return FastBridgeDeployer{
		deployer.NewSimpleDeployer(registry, backend, FastBridgeType),
	}
}

func (f FastBridgeDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return f.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return fastbridge.DeployFastBridge(transactOps, backend, transactOps.From)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return fastbridge.NewFastBridgeRef(address, backend)
	})
}

// MockERC20Deployer deploys a mock erc20 contract.
type MockERC20Deployer struct {
	*deployer.BaseDeployer
}

func NewMockERC20Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MockERC20Deployer{
		deployer.NewSimpleDeployer(registry, backend, MockERC20Type),
	}
}

// MockERC20Decimals is the default number of mock erc20 decimals.
const MockERC20Decimals uint8 = 10

// MockERC20Name is the name of hte mock erc20.
const MockERC20Name = "token"

func (m MockERC20Deployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return m.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return mockerc20.DeployMockERC20(transactOps, backend, MockERC20Name, MockERC20Decimals)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return mockerc20.NewMockerc20Ref(address, backend)
	})
}
