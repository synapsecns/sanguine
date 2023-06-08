package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mocktokenmessenger"
	"testing"

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

// NewDeployManager creates a deploy manager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	parentManager := manager.NewDeployerManager(t,
		NewMockMessageTransmitterDeployer, NewSynapseCCTPDeployer, NewMockTokenMessengerDeployer,
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
func (d MockMessageTransmitterDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		// define the domain as the chain id!
		return mockmessagetransmitter.DeployMockMessageTransmitter(transactOps, backend, uint32(d.Backend().GetChainID()))
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		// remember what I said about vm.ContractRef!
		return mockmessagetransmitter.NewMockMessageTransmitterRef(address, backend)
	})
}

// NewMockTokenMessengerDeployer deploys the mocktokenmessenger.
func NewMockTokenMessengerDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return MockTokenMessengerDeployer{deployer.NewSimpleDeployer(registry, backend, MockTokenMessengerType)}
}

type MockTokenMessengerDeployer struct {
	*deployer.BaseDeployer
}

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

func (m MockTokenMessengerDeployer) Dependencies() []contracts.ContractType {
	return []contracts.ContractType{MockMessageTransmitterType}
}

// MockMessageTransmitterDeployer deploys the mockmessagetransmitter.
type SynapseCCTPDeployer struct {
	*deployer.BaseDeployer
}

// NewMockMessageTransmitterDeployer deploys the light inbox contract.
func NewSynapseCCTPDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return SynapseCCTPDeployer{deployer.NewSimpleDeployer(registry, backend, SynapseCCTPType)}
}

// Deploy deploys the light manager contract.
func (d SynapseCCTPDeployer) Deploy(ctx context.Context) (contracts.DeployedContract, error) {
	return d.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (address common.Address, tx *types.Transaction, data interface{}, err error) {
		tokenMessenger := d.Registry().Get(ctx, MockTokenMessengerType)

		// define the domain as the chain id!
		return cctp.DeploySynapseCCTP(transactOps, backend, tokenMessenger.Address())
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		// remember what I said about vm.ContractRef!
		return cctp.NewSynapseCCTPRef(address, backend)
	})
}

func (d SynapseCCTPDeployer) Dependencies() []contracts.ContractType {
	return d.RecursiveDependencies([]contracts.ContractType{MockTokenMessengerType})
}
