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
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
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

// NewBridgeConfigV3Deployer creates a new bridge config v2 client.
func NewBridgeConfigV3Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return BridgeConfigV3Deployer{deployer.NewSimpleDeployer(registry, backend, BridgeConfigTypeV3)}
}

// NewSynapseBridgeDeployer creates a new bridge config v2 client.
func NewSynapseBridgeDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return SynapseBridgeDeployer{deployer.NewSimpleDeployer(registry, backend, SynapseBridgeType)}
}

// NewSwapFlashLoanDeployer creates a new bridge config v2 client.
func NewSwapFlashLoanDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return SwapFlashLoanDeployer{deployer.NewSimpleDeployer(registry, backend, SwapFlashLoanType)}
}

// Deploy deploys bridge config v3
// nolint: dupl
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

// Deploy deploys Synapse Bridge
// nolint: dupl
func (n SynapseBridgeDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return bridge.DeploySynapseBridge(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return bridge.NewBridgeRef(address, backend)
	})
}

// Deploy deploys Swap Flash Loan
// nolint: dupl
func (n SwapFlashLoanDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return swap.DeploySwapFlashLoan(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return swap.NewSwapRef(address, backend)
	})
}

var _ deployer.ContractDeployer = &BridgeConfigV3Deployer{}
var _ deployer.ContractDeployer = &SynapseBridgeDeployer{}
var _ deployer.ContractDeployer = &SwapFlashLoanDeployer{}
