package testutil

import (
	"context"

	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"

	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
)

// GetTestContract gets the test contract.
func (d *DeployManager) GetTestContract(ctx context.Context, backend backends.SimulatedTestBackend) (contract contracts.DeployedContract, handle *testcontract.Ref) {
	d.T().Helper()

	testContract := d.GetContractRegistry(backend).Get(ctx, TestContractType)
	testContractHandle, ok := testContract.ContractHandle().(*testcontract.Ref)
	assert.True(d.T(), ok)

	return testContract, testContractHandle
}
