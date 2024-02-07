package testutil

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/agents/contracts/bondingmanager"
	"github.com/synapsecns/sanguine/agents/contracts/gasoracle"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightmanager"

	"github.com/synapsecns/sanguine/ethergo/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/deployer"
)

// SynChainID the id of the SynChain.
// TODO: no longer needs to be hardcoded: https://github.com/synapsecns/sanguine/pull/1280#discussion_r1320882617
const SynChainID uint = 10

// LightInboxDeployer deploys the light inbox contract.
type LightInboxDeployer struct {
	*deployer.BaseDeployer
}

// NewLightInboxDeployer deploys the light inbox contract.
func NewLightInboxDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return LightInboxDeployer{deployer.NewSimpleDeployer(registry, backend, LightInboxType)}
}

// Deploy deploys the light manager contract.
//
//nolint:dupword,dupl,cyclop
func (d LightInboxDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	/*originContract := d.Registry().Get(ctx, OriginType)
	destinationContract := d.Registry().Get(ctx, DestinationType)
	originAddress := originContract.Address()
	destinationAddress := destinationContract.Address()*/
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		if d.Backend().GetChainID() == SynChainID {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s on synchain", d.ContractType().ContractName())
		}

		// deploy the light inbox contract
		var rawHandle *lightinbox.LightInbox
		address, tx, rawHandle, err = lightinbox.DeployLightInbox(transactOps, backend, SynapseChainID)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s: %w", d.ContractType().ContractName(), err)
		}
		d.Backend().WaitForConfirmation(ctx, tx)

		// initialize the origin contract
		/*initializationTx, err := rawHandle.Initialize(transactOps, originAddress, destinationAddress)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize contract: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, initializationTx)*/

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return lightinbox.NewLightInboxRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the light inbox contract.
func (d LightInboxDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{}
}

// LightManagerDeployer deploys the light manager contract.
type LightManagerDeployer struct {
	*deployer.BaseDeployer
}

// NewLightManagerDeployer deploys the light manager contract.
func NewLightManagerDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return LightManagerDeployer{deployer.NewSimpleDeployer(registry, backend, LightManagerType)}
}

// Deploy deploys the light manager contract.
//
//nolint:dupword,dupl,cyclop
func (d LightManagerDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	/*originContract := d.Registry().Get(ctx, OriginType)
	destinationContract := d.Registry().Get(ctx, DestinationType)
	originAddress := originContract.Address()
	destinationAddress := destinationContract.Address()*/
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		if d.Backend().GetChainID() == SynChainID {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s on synchain", d.ContractType().ContractName())
		}

		// deploy the light manager contract
		var rawHandle *lightmanager.LightManager
		address, tx, rawHandle, err = lightmanager.DeployLightManager(transactOps, backend, SynapseChainID)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s: %w", d.ContractType().ContractName(), err)
		}
		d.Backend().WaitForConfirmation(ctx, tx)

		// initialize the origin contract
		/*initializationTx, err := rawHandle.Initialize(transactOps, originAddress, destinationAddress)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize contract: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, initializationTx)*/

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return lightmanager.NewLightManagerRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the light manager contract.
func (d LightManagerDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{}
}

// InboxDeployer deploys the inbox contract.
type InboxDeployer struct {
	*deployer.BaseDeployer
}

// NewInboxDeployer deploys the inbox contract.
func NewInboxDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return InboxDeployer{deployer.NewSimpleDeployer(registry, backend, InboxType)}
}

// Deploy deploys the inbox contract.
// nolint:dupl,cyclop,dupword
func (d InboxDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	/*originContract := d.Registry().Get(ctx, OriginType)
	destinationContract := d.Registry().Get(ctx, DestinationType)
	originAddress := originContract.Address()
	destinationAddress := destinationContract.Address()*/
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		if d.Backend().GetChainID() != SynChainID {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s on non-synchain", d.ContractType().ContractName())
		}

		// deploy the inbox contract
		var rawHandle *inbox.Inbox
		address, tx, rawHandle, err = inbox.DeployInbox(transactOps, backend, SynapseChainID)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s: %w", d.ContractType().ContractName(), err)
		}
		d.Backend().WaitForConfirmation(ctx, tx)

		// initialize the origin contract
		/*initializationTx, err := rawHandle.Initialize(transactOps, originAddress, destinationAddress)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize contract: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, initializationTx)*/

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return inbox.NewInboxRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the inbox contract.
func (d InboxDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{}
}

// BondingManagerDeployer deploys the bonding manager contract.
type BondingManagerDeployer struct {
	*deployer.BaseDeployer
}

// NewBondingManagerDeployer deploys the bonding manager contract.
func NewBondingManagerDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return BondingManagerDeployer{deployer.NewSimpleDeployer(registry, backend, BondingManagerType)}
}

// Deploy deploys the bonding manager contract.
// nolint:dupl,cyclop,dupword
func (d BondingManagerDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	/*originContract := d.Registry().Get(ctx, OriginType)
	destinationContract := d.Registry().Get(ctx, DestinationType)
	originAddress := originContract.Address()
	destinationAddress := destinationContract.Address()*/
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		if d.Backend().GetChainID() != SynChainID {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s on a chain that is not synchain", d.ContractType().ContractName())
		}

		// deploy the bonding manager contract
		var rawHandle *bondingmanager.BondingManager
		address, tx, rawHandle, err = bondingmanager.DeployBondingManager(transactOps, backend, SynapseChainID)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s: %w", d.ContractType().ContractName(), err)
		}
		d.Backend().WaitForConfirmation(ctx, tx)

		// initialize the origin contract
		/*initializationTx, err := rawHandle.Initialize(transactOps, originAddress, destinationAddress)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize contract: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, initializationTx)*/

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return bondingmanager.NewBondingManagerRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the bonding manager contract.
func (d BondingManagerDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{}
}

// GasOracleDeployer deploys the gas oracle contract.
type GasOracleDeployer struct {
	*deployer.BaseDeployer
}

// NewGasOracleDeployer deploys the gas oracle contract.
func NewGasOracleDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return GasOracleDeployer{deployer.NewSimpleDeployer(registry, backend, GasOracleType)}
}

// Deploy deploys the gas oracle contract.
// nolint:dupl,cyclop,dupword
func (d GasOracleDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	/*originContract := d.Registry().Get(ctx, OriginType)
	originAddress := originContract.Address()*/
	destinationContract := d.Registry().Get(ctx, DestinationType)
	destinationAddress := destinationContract.Address()
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		// deploy the bonding manager contract
		var rawHandle *gasoracle.GasOracle
		address, tx, rawHandle, err = gasoracle.DeployGasOracle(transactOps, backend, SynapseChainID, destinationAddress)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s: %w", d.ContractType().ContractName(), err)
		}
		d.Backend().WaitForConfirmation(ctx, tx)

		// initialize the origin contract
		/*initializationTx, err := rawHandle.Initialize(transactOps, originAddress, destinationAddress)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize contract: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, initializationTx)*/

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return gasoracle.NewGasOracleRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the gas oracle contract.
func (d GasOracleDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{}
}

// OriginDeployer deploys the origin contract.
type OriginDeployer struct {
	*deployer.BaseDeployer
}

// NewOriginDeployer deploys the origin contract.
func NewOriginDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return OriginDeployer{deployer.NewSimpleDeployer(registry, backend, OriginType)}
}

// Deploy deploys the origin contract.
// nolint:dupl,cyclop
func (d OriginDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	var agentAddress common.Address
	var inboxAddress common.Address
	if d.Backend().GetChainID() == SynChainID {
		bondingManagerContract := d.Registry().Get(ctx, BondingManagerType)
		agentAddress = bondingManagerContract.Address()

		inboxContract := d.Registry().Get(ctx, InboxType)
		inboxAddress = inboxContract.Address()
	} else {
		lightManagerContract := d.Registry().Get(ctx, LightManagerType)
		agentAddress = lightManagerContract.Address()

		lightInboxContract := d.Registry().Get(ctx, LightInboxType)
		inboxAddress = lightInboxContract.Address()
	}
	gasOracleContract := d.Registry().Get(ctx, GasOracleType)
	gasOracleAddress := gasOracleContract.Address()
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		// deploy the origin contract
		var rawHandle *origin.Origin
		address, tx, rawHandle, err = origin.DeployOrigin(transactOps, backend, SynapseChainID, agentAddress, inboxAddress, gasOracleAddress)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s: %w", d.ContractType().ContractName(), err)
		}
		d.Backend().WaitForConfirmation(ctx, tx)

		// initialize the origin contract
		initializationTx, err := rawHandle.Initialize(transactOps)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize contract: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, initializationTx)

		return address, tx, rawHandle, err
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return origin.NewOriginRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the origin contract.
func (d OriginDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{}
}

// SummitDeployer deploys the summit.
type SummitDeployer struct {
	*deployer.BaseDeployer
}

// NewSummitDeployer creates the deployer for  the summit.
func NewSummitDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return SummitDeployer{deployer.NewSimpleDeployer(registry, backend, SummitType)}
}

// Deploy deploys the summit.
//
//nolint:dupword,dupl,cyclop
func (a SummitDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	bondingManagerContract := a.Registry().Get(ctx, BondingManagerType)
	bondingManagerAddress := bondingManagerContract.Address()
	inboxContract := a.Registry().Get(ctx, InboxType)
	inboxAddress := inboxContract.Address()

	return a.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		if a.Backend().GetChainID() != SynChainID {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy %s on nonsynchain", a.ContractType().ContractName())
		}

		summitAddress, summitTx, summit, err := summit.DeploySummit(transactOps, backend, uint32(a.Backend().GetChainID()), bondingManagerAddress, inboxAddress)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy summit: %w", err)
		}

		auth := a.Backend().GetTxContext(ctx, &transactOps.From)
		initTx, err := summit.Initialize(auth.TransactOpts)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize attestation collector: %w", err)
		}
		a.Backend().WaitForConfirmation(ctx, initTx)

		return summitAddress, summitTx, summit, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return summit.NewSummitRef(address, backend)
	})
}

// DestinationDeployer deploys the destination.
type DestinationDeployer struct {
	*deployer.BaseDeployer
}

// NewDestinationDeployer creates the deployer for the destination.
func NewDestinationDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return DestinationDeployer{deployer.NewSimpleDeployer(registry, backend, DestinationType)}
}

// Deploy deploys the destination.
//
//nolint:dupl,dupword
func (d DestinationDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	var agentManagerAddress common.Address
	var inboxAddress common.Address
	if d.Backend().GetChainID() == 10 {
		bondingManagerContract := d.Registry().Get(ctx, BondingManagerType)
		agentManagerAddress = bondingManagerContract.Address()

		inboxContract := d.Registry().Get(ctx, InboxType)
		inboxAddress = inboxContract.Address()
	} else {
		lightManagerContract := d.Registry().Get(ctx, LightManagerType)
		agentManagerAddress = lightManagerContract.Address()

		lightInboxContract := d.Registry().Get(ctx, LightInboxType)
		inboxAddress = lightInboxContract.Address()
	}
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		destinationAddress, destinationTx, destination, err := destination.DeployDestination(transactOps, backend, SynapseChainID, agentManagerAddress, inboxAddress)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy destination: %w", err)
		}

		/*auth := d.Backend().GetTxContext(ctx, &transactOps.From)
		initTx, err := destination.Initialize(auth.TransactOpts)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not initialize destination: %w", err)
		}
		d.Backend().WaitForConfirmation(ctx, initTx)*/

		return destinationAddress, destinationTx, destination, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return destination.NewDestinationRef(address, backend)
	})
}

// Dependencies gets a list of dependencies used to deploy the destination contract.
func (d DestinationDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{OriginType}
}
