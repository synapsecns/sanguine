package testutil

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"

	"github.com/synapsecns/sanguine/ethergo/backends"
)

// GetTestContract gets the test contract.
func (d *DeployManager) GetTestContract(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testcontract.TestContractRef) {
	d.T().Helper()

	return manager.GetContract[*testcontract.TestContractRef](ctx, d.T(), d, backend, TestContractType)
}
