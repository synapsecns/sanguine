package testutil

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/testbridgev1"
	"github.com/synapsecns/sanguine/services/explorer/contracts/cctp/testcctp"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messagebus/testmessagebus"
	"github.com/synapsecns/sanguine/services/explorer/contracts/metaswap/testmetaswap"

	"github.com/synapsecns/sanguine/services/explorer/contracts/erc20"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap/testswap"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
)

// BridgeConfigV3Deployer is the type of the bridge config v3 deployer.
type BridgeConfigV3Deployer struct {
	*deployer.BaseDeployer
}

// SynapseBridgeDeployer is the type of the bridge deployer.
type SynapseBridgeDeployer struct {
	*deployer.BaseDeployer
}

// SwapFlashLoanDeployer is the type of the swap flash loan deployer.
type SwapFlashLoanDeployer struct {
	*deployer.BaseDeployer
}

// SynapseBridgeV1Deployer is the type of the swap flash loan deployer.
type SynapseBridgeV1Deployer struct {
	*deployer.BaseDeployer
}

// MessageBusDeployer is the type of the message bus deployer.
type MessageBusDeployer struct {
	*deployer.BaseDeployer
}

// MetaSwapDeployer is the type of the meta swap deployer.
type MetaSwapDeployer struct {
	*deployer.BaseDeployer
}

// CCTPDeployer is the type of the cctp deployer.
type CCTPDeployer struct {
	*deployer.BaseDeployer
}

// ERC20Deployer is the type of the erc20 deployer.
type ERC20Deployer struct {
	*deployer.BaseDeployer
}

// NewBridgeConfigV3Deployer creates a new bridge config v2 client.
func NewBridgeConfigV3Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return BridgeConfigV3Deployer{deployer.NewSimpleDeployer(registry, backend, BridgeConfigTypeV3)}
}

// NewSynapseBridgeDeployer creates a new bridge client.
func NewSynapseBridgeDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return SynapseBridgeDeployer{deployer.NewSimpleDeployer(registry, backend, SynapseBridgeType)}
}

// NewSwapFlashLoanDeployer creates a new flash loan client.
func NewSwapFlashLoanDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return SwapFlashLoanDeployer{deployer.NewSimpleDeployer(registry, backend, SwapFlashLoanType)}
}

// NewSynapseBridgeV1Deployer creates a new bridge v1 client.
func NewSynapseBridgeV1Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return SynapseBridgeV1Deployer{deployer.NewSimpleDeployer(registry, backend, SynapseBridgeV1Type)}
}

// NewMessageBusDeployer creates a new message bus client.
func NewMessageBusDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MessageBusDeployer{deployer.NewSimpleDeployer(registry, backend, MessageBusType)}
}

// NewMetaSwapDeployer creates a new meta swap client.
func NewMetaSwapDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MetaSwapDeployer{deployer.NewSimpleDeployer(registry, backend, MetaSwapType)}
}

// NewCCTPDeployer creates a new cctp client.
func NewCCTPDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return CCTPDeployer{deployer.NewSimpleDeployer(registry, backend, CCTPType)}
}

// NewERC20Deployer creates a new erc20 client.
func NewERC20Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return ERC20Deployer{deployer.NewSimpleDeployer(registry, backend, ERC20Type)}
}

// Deploy deploys bridge config v3 contract
//
//nolint:dupl
func (n BridgeConfigV3Deployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		deployAddress, tx, handler, err := bridgeconfig.DeployBridgeConfigV3(transactOps, backend)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy bridge config: %w", err)
		}

		// wait for confirm, we need this to grant role
		n.Backend().WaitForConfirmation(ctx, tx)

		// https://github.com/synapsecns/synapse-contracts/pull/13 introduces a breaking change where the BRIDGEMANAGER_ROLE is not automatically granted to the
		// deployer. We fix that here by granting the role to the owner
		bridgeManagerRole, err := handler.BRIDGEMANAGERROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not get bridge manager role: %w", err)
		}

		auth := n.Backend().GetTxContext(ctx, &transactOps.From)
		// grant the role
		grantTx, err := handler.GrantRole(auth.TransactOpts, bridgeManagerRole, auth.From)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not grant bridge manager role: %w", err)
		}

		n.Backend().WaitForConfirmation(ctx, grantTx)

		return deployAddress, tx, handler, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return bridgeconfig.NewBridgeConfigRef(address, backend)
	})
}

// Deploy deploys Synapse Bridge contract
//
//nolint:dupl
func (n SynapseBridgeDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testbridge.DeployTestSynapseBridge(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testbridge.NewTestBridgeRef(address, backend)
	})
}

// Deploy deploys Synapse Bridge V1 contract
//
//nolint:dupl
func (n SynapseBridgeV1Deployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testbridgev1.DeployTestSynapseBridgeV1(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testbridgev1.NewTestBridgeV1Ref(address, backend)
	})
}

// Deploy deploys Swap Flash Loan contract
//
//nolint:dupl
func (n SwapFlashLoanDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testswap.DeployTestSwapFlashLoan(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testswap.NewTestSwapRef(address, backend)
	})
}

// Deploy deploys Message Bus contract
//
//nolint:dupl
func (n MessageBusDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testmessagebus.DeployTestMessageBusUpgradeable(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testmessagebus.NewTestMessageBusRef(address, backend)
	})
}

// Deploy deploys Meta Swap contract
//
//nolint:dupl
func (n MetaSwapDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testmetaswap.DeployTestMetaSwap(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testmetaswap.NewTestMetaSwapRef(address, backend)
	})
}

// Deploy deploys CCTP contract
//
//nolint:dupl
func (n CCTPDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	tokenMessengerContract, err := n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return testcctp.DeployMessageTransmitter(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testcctp.NewMessageTransmitter(address, backend)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to deploy tokenMessengerContract %w", err)
	}
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		// Create mock owner
		owner := common.BigToAddress(big.NewInt(gofakeit.Int64()))

		return testcctp.DeployTestSynapseCCTP(transactOps, backend, tokenMessengerContract.Address(), owner)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return testcctp.NewTestCCTPRef(address, backend)
	})
}

// Deploy deploys ERC20 contract
//
//nolint:dupl
func (n ERC20Deployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return erc20.DeployTestERC(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return erc20.NewTestERC(address, backend)
	})
}

var _ deployer.ContractDeployer = &BridgeConfigV3Deployer{}
var _ deployer.ContractDeployer = &SynapseBridgeDeployer{}
var _ deployer.ContractDeployer = &SwapFlashLoanDeployer{}
var _ deployer.ContractDeployer = &SynapseBridgeV1Deployer{}
var _ deployer.ContractDeployer = &MetaSwapDeployer{}
var _ deployer.ContractDeployer = &CCTPDeployer{}
var _ deployer.ContractDeployer = &ERC20Deployer{}
