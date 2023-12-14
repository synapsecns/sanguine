package testutil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/fastbridgemock"
	"math/big"
	"testing"
)

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	// TODO: add contracts here
	parentManager := manager.NewDeployerManager(t, NewFastBridgeDeployer, NewMockERC20Deployer, NewMockFastBridgeDeployer, NewWETH9Deployer, NewUSDTDeployer, NewUSDCDeployer, NewDAIDeployer)
	return &DeployManager{parentManager}
}

// MintToAddress mints an equal amount of tokens to an address. Amount is multiplied by decimals to ensure events.
func (d *DeployManager) MintToAddress(ctx context.Context, backend backends.SimulatedTestBackend, token contracts.ContractType, mintToAddress common.Address, amount *big.Int) (adjustedAmount *big.Int) {
	d.T().Helper()

	//nolint: exhaustive
	switch token {
	case DAIType:
		contract, handle := d.GetDAI(ctx, backend)

		auth := backend.GetTxContext(ctx, contract.OwnerPtr())
		adjustedAmount = mustAdjustAmount(ctx, d.T(), amount, handle)

		tx, err := handle.Mint(auth.TransactOpts, mintToAddress, adjustedAmount)
		assert.NoError(d.T(), err)

		backend.WaitForConfirmation(ctx, tx)
	case USDTType:
		contract, handle := d.GetUSDT(ctx, backend)

		adjustedAmount = mustAdjustAmount(ctx, d.T(), amount, handle)

		auth := backend.GetTxContext(ctx, contract.OwnerPtr())
		tx, err := handle.Redeem(auth.TransactOpts, adjustedAmount)
		assert.NoError(d.T(), err)

		backend.WaitForConfirmation(ctx, tx)

		auth = backend.GetTxContext(ctx, contract.OwnerPtr())
		tx, err = handle.Transfer(auth.TransactOpts, mintToAddress, adjustedAmount)
		assert.NoError(d.T(), err)

		backend.WaitForConfirmation(ctx, tx)

	case USDCType:
		contract, handle := d.GetUSDC(ctx, backend)

		adjustedAmount = mustAdjustAmount(ctx, d.T(), amount, handle)

		auth := backend.GetTxContext(ctx, contract.OwnerPtr())
		tx, err := handle.ConfigureMinter(auth.TransactOpts, contract.Owner(), adjustedAmount)
		assert.NoError(d.T(), err)
		backend.WaitForConfirmation(ctx, tx)
		fmt.Printf("minter configured: %s \n ", tx.Hash())

		auth = backend.GetTxContext(ctx, contract.OwnerPtr())
		tx, err = handle.Mint(auth.TransactOpts, mintToAddress, adjustedAmount)
		assert.NoError(d.T(), err)

		backend.WaitForConfirmation(ctx, tx)
	case MockERC20Type:
		contract, handle := d.GetMockERC20(ctx, backend)
		adjustedAmount = mustAdjustAmount(ctx, d.T(), amount, handle)

		auth := backend.GetTxContext(ctx, contract.OwnerPtr())

		tx, err := handle.Mint(auth.TransactOpts, mintToAddress, adjustedAmount)
		assert.NoError(d.T(), err)

		backend.WaitForConfirmation(ctx, tx)
	default:
		d.T().Errorf("contract type %s not (yet) supported", token.Name())
	}

	return adjustedAmount
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
