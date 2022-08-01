package testutil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/contracts/updatermanager"
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
	updateManagerContract := d.Registry().Get(ctx, UpdaterManagerType)

	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		// deploy the home contract
		var rawHandle *home.Home
		address, tx, rawHandle, err = home.DeployHome(transactOps, backend, uint32(d.Backend().GetChainID()))
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy home: %w", err)
		}

		// initialize the home contract
		initializationTx, err := rawHandle.Initialize(transactOps, updateManagerContract.Address())
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize contract: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, initializationTx)

		// get the owner of the updater manage rcontract
		updaterTransactOps := d.Backend().GetTxContext(ctx, updateManagerContract.OwnerPtr())

		// set the home contract on the updater manager
		updateManager, ok := updateManagerContract.ContractHandle().(*updatermanager.UpdaterManagerRef)
		if !ok {
			return common.Address{}, nil, nil, fmt.Errorf("could not update contract: %w", err)
		}

		setTx, err := updateManager.SetHome(updaterTransactOps.TransactOpts, address)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not set home: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, setTx)

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return home.NewHomeRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the home contract.
func (d HomeDeployer) Dependencies() []deployer.ContractType {
	return []deployer.ContractType{UpdaterManagerType}
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

// UpdateManagerDeployer deploys the update manager.
type UpdateManagerDeployer struct {
	*deployer.BaseDeployer
}

// NewUpdateManagerDeployer deploys a new update manager.
func NewUpdateManagerDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return UpdateManagerDeployer{deployer.NewSimpleDeployer(registry, backend, UpdaterManagerType)}
}

// Deploy deploys the updater contract.
func (u UpdateManagerDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return u.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return updatermanager.DeployUpdaterManager(transactOps, backend, transactOps.From)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return updatermanager.NewUpdaterManagerRef(address, backend)
	})
}

// AttestationCollectorDeployer deploys the attestation collector.
type AttestationCollectorDeployer struct {
	*deployer.BaseDeployer
}

// NewAttestationCollectorDeployer creates the deployer for  the attestation collecotr.
func NewAttestationCollectorDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return AttestationCollectorDeployer{deployer.NewSimpleDeployer(registry, backend, AttestationCollectorType)}
}

// Deploy deploys the attestation collector.
func (a AttestationCollectorDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return a.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		attestationAddress, attestationTx, collector, err := attestationcollector.DeployAttestationCollector(transactOps, backend)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy attesation collector: %w", err)
		}

		auth := a.Backend().GetTxContext(ctx, &transactOps.From)
		initTx, err := collector.Initialize(auth.TransactOpts)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize attesation collector: %w", err)
		}
		a.Backend().WaitForConfirmation(ctx, initTx)

		return attestationAddress, attestationTx, collector, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return attestationcollector.NewAttestationCollectorRef(address, backend)
	})
}
