package testutil

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/rfq/contracts/bridgetransactionv2"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridgev2"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/fastbridgemockv2"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/recipientmock"
)

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

// NewDeployManager creates a new DeployManager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	parentManager := manager.NewDeployerManager(t, NewFastBridgeDeployer, NewMockERC20Deployer, NewMockFastBridgeDeployer, NewRecipientMockDeployer, NewBridgeTransactionV2Deployer, NewWETH9Deployer, NewUSDTDeployer, NewUSDCDeployer, NewDAIDeployer)
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
		adjustedAmount = MustAdjustAmount(ctx, d.T(), amount, handle)

		tx, err := handle.Mint(auth.TransactOpts, mintToAddress, adjustedAmount)
		assert.NoError(d.T(), err)

		backend.WaitForConfirmation(ctx, tx)
	case USDTType:
		contract, handle := d.GetUSDT(ctx, backend)

		adjustedAmount = MustAdjustAmount(ctx, d.T(), amount, handle)

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

		adjustedAmount = MustAdjustAmount(ctx, d.T(), amount, handle)

		auth := backend.GetTxContext(ctx, contract.OwnerPtr())
		// TODO: it's fairly likely we should just configure this w/ max mint rights to minter
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
		adjustedAmount = MustAdjustAmount(ctx, d.T(), amount, handle)

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

// Deploy deploys the fast bridge contract.
func (f FastBridgeDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return f.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return fastbridge.DeployFastBridge(transactOps, backend, transactOps.From)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return fastbridge.NewFastBridgeRef(address, backend)
	})
}

// FastBridgeV2Deployer deplyos a fast bridge contract for testing.
type FastBridgeV2Deployer struct {
	*deployer.BaseDeployer
}

// NewFastBridgeV2Deployer deploys a fast bridge contract.
func NewFastBridgeV2Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return FastBridgeV2Deployer{
		deployer.NewSimpleDeployer(registry, backend, FastBridgeV2Type),
	}
}

// Deploy deploys the fast bridge contract.
func (f FastBridgeV2Deployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return f.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return fastbridgev2.DeployFastBridgeV2(transactOps, backend, transactOps.From)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return fastbridgev2.NewFastBridgeV2Ref(address, backend)
	})
}

// MockFastBridgeDeployer deploys a mock fast bridge contract for testing.
type MockFastBridgeDeployer struct {
	*deployer.BaseDeployer
}

// NewMockFastBridgeDeployer deploys a mock fast bridge contract.
func NewMockFastBridgeDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MockFastBridgeDeployer{
		deployer.NewSimpleDeployer(registry, backend, FastBridgeMockType),
	}
}

// Deploy deploys the mock fast bridge contract.
func (m MockFastBridgeDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return m.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return fastbridgemockv2.DeployFastBridgeMock(transactOps, backend, transactOps.From)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return fastbridgemockv2.NewFastBridgeMockRef(address, backend)
	})
}

// RecipientMockDeployer deploys a mock recipient contract for testing.
type RecipientMockDeployer struct {
	*deployer.BaseDeployer
}

// NewRecipientMockDeployer deploys a mock recipient contract.
func NewRecipientMockDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return RecipientMockDeployer{
		deployer.NewSimpleDeployer(registry, backend, RecipientMockType),
	}
}

// Deploy deploys the recipient mock contract.
func (m RecipientMockDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return m.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return recipientmock.DeployRecipientMock(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return recipientmock.NewRecipientMockRef(address, backend)
	})
}

// BridgeTransactionV2Deployer deploys a bridge transaction contract for testing.
type BridgeTransactionV2Deployer struct {
	*deployer.BaseDeployer
}

// NewBridgeTransactionV2Deployer deploys a mock recipient contract.
func NewBridgeTransactionV2Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return BridgeTransactionV2Deployer{
		deployer.NewSimpleDeployer(registry, backend, BridgeTransactionV2Type),
	}
}

// Deploy deploys the bridge transaction contract.
func (m BridgeTransactionV2Deployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return m.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return bridgetransactionv2.DeployBridgeTransactionV2Harness(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return bridgetransactionv2.NewBridgeTransactionV2Ref(address, backend)
	})
}
