package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract2"

	"github.com/synapsecns/sanguine/ethergo/backends"
)

// GetTestContract gets the test contract.
func (d *DeployManager) GetTestContract(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testcontract.TestContractRef) {
	d.T().Helper()

	return manager.GetContract[*testcontract.TestContractRef](ctx, d.T(), d, backend, TestContractType)
}

// GetTestContract2 gets the test contract.
func (d *DeployManager) GetTestContract2(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testcontract2.TestContractRef) {
	d.T().Helper()

	return manager.GetContract[*testcontract2.TestContractRef](ctx, d.T(), d, backend, TestContract2Type)
}
