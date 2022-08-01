package testutil

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/core/contracts/test/homeharness"
	"github.com/synapsecns/sanguine/core/contracts/test/messageharness"
	"github.com/synapsecns/sanguine/core/contracts/test/replicamanagerharness"
	"github.com/synapsecns/sanguine/core/contracts/test/tipsharness"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/synapse-node/testutils/backends"
)

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

// HomeHarnessDomain is the domain used for the home harness.
const HomeHarnessDomain = 1

// HomeHarnessDeployer deploys the home harness for testing.
type HomeHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewHomeHarnessDeployer deploys the new home harness.
func NewHomeHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return HomeHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, HomeHarnessType)}
}

// Deploy deploys the home harness.
func (h HomeHarnessDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return h.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return homeharness.DeployHomeHarness(transactOps, backend, HomeHarnessDomain)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return homeharness.NewHomeHarnessRef(address, backend)
	})
}

// AttestationHarnessDeployer deploys the attestation harness.
type AttestationHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewAttestationHarnessDeployer creates a new deployer for the attestation harness.
func NewAttestationHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return AttestationHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, AttestationHarnessType)}
}

// Deploy deploys the attestation harness.
func (a AttestationHarnessDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return a.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return attestationharness.DeployAttestationHarness(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return attestationharness.NewAttestationHarnessRef(address, backend)
	})
}

// TipsHarnessDeployer deploys the tip harness for tester.
type TipsHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewTipsHarnessDeployer creates a new deployer for the attestation harness.
func NewTipsHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return TipsHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, TipsHarnessType)}
}

// Deploy deploys the attestation harness.
func (a TipsHarnessDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return a.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return tipsharness.DeployTipsHarness(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return tipsharness.NewTipsHarnessRef(address, backend)
	})
}

// ReplicaManagerHarnessDeployer deploys the replica manager harness.
type ReplicaManagerHarnessDeployer struct {
	*deployer.BaseDeployer
}

// NewReplicaManagerHarnessDeployer creates a new deployer for the replica manager harness.
func NewReplicaManagerHarnessDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
	return ReplicaManagerHarnessDeployer{deployer.NewSimpleDeployer(registry, backend, ReplicaManagerHarnessType)}
}

// Deploy deploys the replica manager harness.
func (r ReplicaManagerHarnessDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
	return r.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return replicamanagerharness.DeployReplicaManagerHarness(transactOps, backend, uint32(r.Backend().GetChainID()))
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return replicamanagerharness.NewReplicaManagerHarnessRef(address, backend)
	})
}
