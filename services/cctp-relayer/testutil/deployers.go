package testutil

import (
	"context"
	"fmt"
	"testing"

	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mockmintburntoken"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mocktokenmessenger"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mocktokenminter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mockmessagetransmitter"
)

// ChainIDDomainMap maps chain IDs to domains.
// see: https://developers.circle.com/stablecoin/docs/cctp-technical-reference#domain
// TODO: make this eaiser to debug in the case of missing domains, etc.
var ChainIDDomainMap = map[uint32]uint32{
	1:     0,
	43114: 1,
}

// NewDeployManager creates a deploy manager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	parentManager := manager.NewDeployerManager(t,
		NewMockMessageTransmitterDeployer, NewSynapseCCTPDeployer, NewMockTokenMessengerDeployer, NewMockMintBurnTokenDeployer, NewMockTokenMinterDeployer,
	)
	return &DeployManager{parentManager}
}

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

// MockMessageTransmitterDeployer deploys the mockmessagetransmitter.
type MockMessageTransmitterDeployer struct {
	*deployer.BaseDeployer
}

// NewMockMessageTransmitterDeployer deploys the light inbox contract.
func NewMockMessageTransmitterDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MockMessageTransmitterDeployer{deployer.NewSimpleDeployer(registry, backend, MockMessageTransmitterType)}
}

// Deploy deploys the light manager contract.
//
//nolint:dupword,dupl,cyclop
func (d MockMessageTransmitterDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		// define the domain as the chain id!
		return mockmessagetransmitter.DeployMockMessageTransmitter(transactOps, backend, ChainIDDomainMap[uint32(d.Backend().GetChainID())])
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		// remember what I said about vm.ContractRef!
		return mockmessagetransmitter.NewMockMessageTransmitterRef(address, backend)
	})
}

// NewMockTokenMessengerDeployer deploys the mocktokenmessenger.
func NewMockTokenMessengerDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MockTokenMessengerDeployer{deployer.NewSimpleDeployer(registry, backend, MockTokenMessengerType)}
}

// MockTokenMessengerDeployer deploys the mocktokenmessenger.
type MockTokenMessengerDeployer struct {
	*deployer.BaseDeployer
}

// Deploy deploys the mock token messenger contract.
//
//nolint:dupword,dupl,cyclop
func (m MockTokenMessengerDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	messageTransmitter := m.Registry().Get(ctx, MockMessageTransmitterType)

	return m.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		// define the domain as the chain id!
		return mocktokenmessenger.DeployMockTokenMessenger(transactOps, backend, messageTransmitter.Address())
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		// remember what I said about vm.ContractRef!
		return mocktokenmessenger.NewMockTokenMessengerRef(address, backend)
	})
}

// Dependencies returns the dependencies of the mocktokenmessenger.
func (m MockTokenMessengerDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{MockMessageTransmitterType}
}

// NewMockMintBurnTokenDeployer deploys the mocktokenmessenger.
func NewMockMintBurnTokenDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MockMintBurnTokenDeployer{deployer.NewSimpleDeployer(registry, backend, MockMintBurnTokenType)}
}

// MockTokenMinterDeployer deploys the mocktokenminter.
type MockTokenMinterDeployer struct {
	*deployer.BaseDeployer
}

// NewMockTokenMinterDeployer deploys the mocktokenminter.
func NewMockTokenMinterDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MockTokenMinterDeployer{deployer.NewSimpleDeployer(registry, backend, MockTokenMinterType)}
}

// Deploy deploys the mock token minter contract.
//
//nolint:dupword,dupl,cyclop
func (m MockTokenMinterDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	tokenMessenger := m.Registry().Get(ctx, MockTokenMessengerType)

	return m.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		// define the domain as the chain id!
		address, tx, handle, err := mocktokenminter.DeployMockTokenMinter(transactOps, backend, tokenMessenger.Address())
		if err != nil {
			return address, tx, handle, fmt.Errorf("could not deploy mock token minter: %w", err)
		}

		messengerOpts := m.Backend().GetTxContext(ctx, tokenMessenger.OwnerPtr())
		messengerHandle, ok := tokenMessenger.ContractHandle().(*mocktokenmessenger.MockTokenMessengerRef)
		if !ok {
			return address, tx, handle, fmt.Errorf("could not case %T to %T", tokenMessenger.ContractHandle(), messengerHandle)
		}

		setTx, err := messengerHandle.SetLocalMinter(messengerOpts.TransactOpts, address)
		if err != nil {
			return address, tx, handle, fmt.Errorf("could not set local minter: %w", err)
		}
		m.Backend().WaitForConfirmation(ctx, setTx)

		return address, tx, handle, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		// remember what I said about vm.ContractRef!
		return mocktokenminter.NewMockTokenMinterRef(address, backend)
	})
}

// Dependencies returns the dependencies of the mock token minter.
func (m MockTokenMinterDeployer) Dependencies() []contracts.ContractType {
	return m.RecursiveDependencies([]contracts.ContractType{MockTokenMessengerType})
}

// MockMintBurnTokenDeployer deploys the mocktokenminter.
type MockMintBurnTokenDeployer struct {
	*deployer.BaseDeployer
}

// Deploy deploys the mock mint burn token contract.
//
//nolint:dupword,dupl,cyclop
func (m MockMintBurnTokenDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	messageTransmitter := m.Registry().Get(ctx, MockMessageTransmitterType)

	return m.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		// define the domain as the chain id!
		return mockmintburntoken.DeployMockMintBurnToken(transactOps, backend, messageTransmitter.Address())
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		// remember what I said about vm.ContractRef!
		return mockmintburntoken.NewMockMintBurnTokenRef(address, backend)
	})
}

// Dependencies returns the dependencies of the mock mint burn token.
func (m MockMintBurnTokenDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{MockMessageTransmitterType}
}

// SynapseCCTPDeployer deploys the synapse cctp contract.
type SynapseCCTPDeployer struct {
	*deployer.BaseDeployer
}

// NewSynapseCCTPDeployer deploys the synapse cctp contract.
func NewSynapseCCTPDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return SynapseCCTPDeployer{deployer.NewSimpleDeployer(registry, backend, SynapseCCTPType)}
}

// Deploy deploys the light manager contract.
//
//nolint:dupword,dupl,cyclop
func (d SynapseCCTPDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		tokenMessenger := d.Registry().Get(ctx, MockTokenMessengerType)

		// make this a dependency so it self registers the remote domains
		_ = d.Registry().Get(ctx, MockTokenMinterType)

		// define the domain as the chain id!
		return cctp.DeploySynapseCCTP(transactOps, backend, tokenMessenger.Address(), transactOps.From)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		// remember what I said about vm.ContractRef!
		return cctp.NewSynapseCCTPRef(address, backend)
	})
}

// Dependencies returns the dependencies of the SynapseCCTP contract.
func (d SynapseCCTPDeployer) Dependencies() []contracts.ContractType {
	return d.RecursiveDependencies([]contracts.ContractType{MockTokenMessengerType, MockTokenMinterType})
}
