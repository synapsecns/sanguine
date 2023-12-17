package testutil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/dai"
	mockerc202 "github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/mockerc20"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/usdc"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/usdt"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/weth9"
	"math/big"
)

// MockERC20Deployer deploys a mock erc20 contract.
type MockERC20Deployer struct {
	*deployer.BaseDeployer
}

// NewMockERC20Deployer creates a new mock erc20 deployer.
func NewMockERC20Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MockERC20Deployer{
		deployer.NewSimpleDeployer(registry, backend, MockERC20Type),
	}
}

// MockERC20Decimals is the default number of mock erc20 decimals.
const MockERC20Decimals uint8 = 10

// MockERC20Name is the name of hte mock erc20.
const MockERC20Name = "token"

// Deploy deploys a mock erc20 contract.
func (m MockERC20Deployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return m.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return mockerc202.DeployMockERC20(transactOps, backend, MockERC20Name, MockERC20Decimals)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return mockerc202.NewMockerc20Ref(address, backend)
	})
}

// WETH9Deployer deploys a mock erc20 contract.
type WETH9Deployer struct {
	*deployer.BaseDeployer
}

// NewWETH9Deployer creates a new deployer for weth9.
func NewWETH9Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return WETH9Deployer{
		deployer.NewSimpleDeployer(registry, backend, WETH9Type),
	}
}

// Deploy deploys the weth9 contract.
func (m WETH9Deployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return m.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return weth9.DeployWETH9(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return weth9.NewWeth9Ref(address, backend)
	})
}

// USDTDeployer deploys the usdt token (https://tether.to/) for testing.
type USDTDeployer struct {
	*deployer.BaseDeployer
}

var defaultTetherInitialSupply = big.NewInt(1e6 * 50000000000)

// NewUSDTDeployer creates a new deployer for tether.
func NewUSDTDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return USDTDeployer{deployer.NewSimpleDeployer(registry, backend, USDTType)}
}

// Deploy deploys the usdt token.
func (d USDTDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return usdt.DeployTetherToken(transactOps, backend, defaultTetherInitialSupply, "Tether USD", "USDT", big.NewInt(6))
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return usdt.NewUSDTRef(address, backend)
	})
}

// USDCDeployer deploys the usdc token (https://www.centre.io/usdc) for testing.
type USDCDeployer struct {
	*deployer.BaseDeployer
}

// NewUSDCDeployer creates a new deployer for tether.
func NewUSDCDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return USDCDeployer{deployer.NewSimpleDeployer(registry, backend, USDCType)}
}

// Deploy deploys the usdt token.
func (d USDCDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		tmpAddress, tx, handle, err := usdc.DeployFiatTokenV21(transactOps, backend)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy usdc contract")
		}
		d.Backend().WaitForConfirmation(ctx, tx)

		auth := d.Backend().GetTxContext(ctx, &transactOps.From)

		// see https://etherscan.io/token/0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48#writeProxyContract
		initializedTx, err := handle.Initialize(auth.TransactOpts, "USDC Coin", "USDC", "USD", uint8(6), auth.From, auth.From, auth.From, auth.From)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize usdc contract")
		}
		d.Backend().WaitForConfirmation(ctx, initializedTx)
		return tmpAddress, tx, handle, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return usdc.NewUSDCRef(address, backend)
	})
}

// DAIDeployer deploys a mock erc20 contract.
type DAIDeployer struct {
	*deployer.BaseDeployer
}

// NewDAIDeployer creates a new deployer for dai.
func NewDAIDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return DAIDeployer{
		deployer.NewSimpleDeployer(registry, backend, DAIType),
	}
}

// Deploy deploys the dai contract.
func (m DAIDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return m.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return dai.DeployDai(transactOps, backend, m.Backend().GetBigChainID())
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return dai.NewDaiRef(address, backend)
	})
}
