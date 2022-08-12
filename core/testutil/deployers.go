package testutil

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/contracts/replicamanager"
	"github.com/synapsecns/sanguine/core/contracts/updatermanager"
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

// ReplicaManagerDeployer deploys the replica manager.
type ReplicaManagerDeployer struct {
	*deployer.BaseDeployer
}

// Dependencies gets a list of dependencies used to deploy the replicamanager contract.
func (r ReplicaManagerDeployer) Dependencies() []deployer.ContractType {
	return []deployer.ContractType{HomeType, UpdaterManagerType}
}

// NewReplicaManagerDeployer creates the deployer for the replica manager.
func NewReplicaManagerDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return ReplicaManagerDeployer{deployer.NewSimpleDeployer(registry, backend, ReplicaManagerType)}
}

// Deploy deploys the replica manager.
func (r ReplicaManagerDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return r.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		replicaAddress, replicaTx, replica, err := replicamanager.DeployReplicaManager(transactOps, backend, uint32(r.Backend().GetChainID()))
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy replica manager: %w", err)
		}

		auth := r.Backend().GetTxContext(ctx, &transactOps.From)
		initTx, err := replica.Initialize(auth.TransactOpts, uint32(r.Registry().Get(ctx, HomeType).ChainID().Uint64()), common.Address{})
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize replica manager: %w", err)
		}
		r.Backend().WaitForConfirmation(ctx, initTx)

		return replicaAddress, replicaTx, replica, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return replicamanager.NewReplicaManagerRef(address, backend)
	})
}
