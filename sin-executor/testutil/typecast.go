package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/sin-executor/contracts/executionservice"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchaindb"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/interchainapp"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/interchainmodulemock"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/optionslibexport"
)

// GetInterchainClient gets the interchain client.
func (d *DeployManager) GetInterchainClient(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *interchainclient.InterchainClientRef) {
	d.T().Helper()

	return manager.GetContract[*interchainclient.InterchainClientRef](ctx, d.T(), d, backend, InterchainClient)
}

// GetOptionsLib gets the options library.
func (d *DeployManager) GetOptionsLib(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *optionslibexport.OptionsLibExportRef) {
	d.T().Helper()

	return manager.GetContract[*optionslibexport.OptionsLibExportRef](ctx, d.T(), d, backend, OptionsLib)
}

// GetInterchainDB gets the interchain db.
func (d *DeployManager) GetInterchainDB(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *interchaindb.InterchainDBRef) {
	d.T().Helper()

	return manager.GetContract[*interchaindb.InterchainDBRef](ctx, d.T(), d, backend, InterchainDB)
}

// GetInterchainModuleMock gets the interchain module mock.
func (d *DeployManager) GetInterchainModuleMock(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *interchainmodulemock.InterchainModuleMockRef) {
	d.T().Helper()

	return manager.GetContract[*interchainmodulemock.InterchainModuleMockRef](ctx, d.T(), d, backend, InterchainModuleMock)
}

// GetInterchainAppMock gets the interchain app mock.
func (d *DeployManager) GetInterchainAppMock(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *interchainapp.InterchainAppMockRef) {
	d.T().Helper()

	return manager.GetContract[*interchainapp.InterchainAppMockRef](ctx, d.T(), d, backend, InterchainApp)
}

func (d *DeployManager) GetExecutionService(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *executionservice.ExecutionServiceRef) {
	d.T().Helper()

	return manager.GetContract[*executionservice.ExecutionServiceRef](ctx, d.T(), d, backend, ExecutionService)
}
