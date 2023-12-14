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
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/fastbridgemock"
	"testing"
)

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	// TODO: add contracts here
	parentManager := manager.NewDeployerManager(t, NewFastBridgeDeployer, NewMockERC20Deployer, NewMockFastBridgeDeployer, NewWETH9Deployer, NewUSDTDeployer, NewUSDCDeployer)
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

type MockFastBridgeDeployer struct {
	*deployer.BaseDeployer
}

func NewMockFastBridgeDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MockFastBridgeDeployer{
		deployer.NewSimpleDeployer(registry, backend, FastBridgeMockType),
	}
}

func (m MockFastBridgeDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return m.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return fastbridgemock.DeployFastBridgeMock(transactOps, backend, transactOps.From)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return fastbridgemock.NewFastBridgeMockRef(address, backend)
	})
}
