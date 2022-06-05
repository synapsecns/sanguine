package testutil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/contracts/test/homeharness"
	"github.com/synapsecns/sanguine/core/contracts/test/messageharness"
	"github.com/synapsecns/sanguine/core/contracts/xappconfig"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/synapse-node/testutils/backends"
)

// HomeDeployer deploys the home contract.
type HomeDeployer struct {
	*deployer.BaseDeployer
}

// NewHomeDeployer deploys the home contract.
func NewHomeDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return HomeDeployer{deployer.NewSimpleDeployer(registry, backend, HomeType)}
}

// Deploy deploys the home contract.
func (d HomeDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return home.DeployHome(transactOps, backend, uint32(d.Backend().GetChainID()))
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return home.NewHomeRef(address, backend)
	})
}

// XAppConfigDeployer deploys the XAppConfig contract.
type XAppConfigDeployer struct {
	*deployer.BaseDeployer
}

// NewXAppConfigDeployer creates an XAppConfig deployer.
func NewXAppConfigDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return XAppConfigDeployer{deployer.NewSimpleDeployer(registry, backend, XAppConfigType)}
}

// Deploy deploys the xapp config contract.
func (d XAppConfigDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	homeContract := d.Registry().Get(ctx, HomeType)

	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		xAppAddress, deployTx, handle, err := xappconfig.DeployXAppConfig(transactOps, backend)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy xapp config: %w", err)
		}

		// wait for xapp deployment
		d.Backend().WaitForConfirmation(ctx, deployTx)

		// verify we're setting to the owner
		xAppOwner, err := handle.Owner(&bind.CallOpts{Context: ctx})
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not get xapp owner: %w", err)
		}

		auth := d.Backend().GetTxContext(ctx, &xAppOwner)

		// set the home to the previously deployed contract
		tx, err := handle.SetHome(auth.TransactOpts, homeContract.Address())
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not set home on xapp %s: %w", xAppAddress, err)
		}
		d.Backend().WaitForConfirmation(ctx, tx)

		return xAppAddress, deployTx, handle, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return xappconfig.NewXAppConfigRef(address, backend)
	})
}

// Dependencies gets dependencies of the xappconfig contract.
func (d XAppConfigDeployer) Dependencies() []deployer.ContractType {
	return d.RecursiveDependencies([]deployer.ContractType{HomeType})
}

// MessageHarnessDeployer deploys the message harness for testing.
type MessageHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewMessageHarnessDeployer creates a message harness deployer.
func NewMessageHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MessageHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, MessageHarnessType)}
}

// Deploy deploys the message harness deployer.
func (d MessageHarnessDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return messageharness.DeployMessageHarness(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return messageharness.NewMessageHarnessRef(address, backend)
	})
}

// HomeHarnessDeployer deploys the home harness for testing.
type HomeHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewHomeHarnessDeployer deploys the new home harness.
func NewHomeHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return HomeDeployer{deployer.NewSimpleDeployer(registry, backend, HomeHarnessType)}
}

// Deploy deploys the home harness.
func (h HomeHarnessDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return h.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return homeharness.DeployHomeHarness(transactOps, backend, 1)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return homeharness.NewHomeHarnessRef(address, backend)
	})
}
