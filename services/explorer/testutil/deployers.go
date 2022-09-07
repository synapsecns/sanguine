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
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
)

type BridgeConfigV3Deployer struct {
	*deployer.BaseDeployer
}

// NewBridgeConfigV3Deployer creates a new bridge config v2 client.
func NewBridgeConfigV3Deployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return BridgeConfigV3Deployer{deployer.NewSimpleDeployer(registry, backend, BridgeConfigTypeV3)}
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

var _ deployer.ContractDeployer = &BridgeConfigV3Deployer{}
