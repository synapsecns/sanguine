package testutil

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/sin-explorer/contracts/interchainclienteventmock"
	"testing"
)

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

// NewDeployManager creates a deploy manager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	return &DeployManager{manager.NewDeployerManager(t, interchainFunctionalDeployer)}
}

// GetInterchainClientEventMock gets a interchain client event mock contract.
func (d *DeployManager) GetInterchainClientEventMock(ctx context.Context, backend backends.SimulatedTestBackend) (contracts.DeployedContract, *interchainclienteventmock.InterchainClientEventMockRef) {
	d.T().Helper()

	return manager.GetContract[*interchainclienteventmock.InterchainClientEventMockRef](ctx, d.T(), d, backend, InterchainClientEventMock)
}

var (
	interchainFunctionalDeployer = deployer.NewFunctionalDeployer(InterchainClientEventMock, func(ctx context.Context, helpers deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return interchainclienteventmock.DeployInterchainClientEventMock(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return interchainclienteventmock.NewInterchainClientEventMockRef(address, backend)
	}, []contracts.ContractType{})
)
