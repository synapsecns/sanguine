package multimock

import (
	"context"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/testutil"
	"github.com/synapsecns/synapse-node/testutils/backends"
	"github.com/synapsecns/synapse-node/testutils/backends/simulated"
	"testing"
)

// Mocker is used to mock out interactions cross-chain.
type Mocker struct {
	// t is the testing object. Since we use mocker exlusively for testing
	// we don't always bubble up errors. These are instead caught through assertions
	// htat will call fail on the test object
	t *testing.T
	// deployManager manages deployments
	deployManager *testutil.DeployManager
	// homeBackend is the backend where the home contract is stored
	homeBackend backends.SimulatedTestBackend
	// homeContract is the contract deployed on the home chain
	homeContract *home.HomeRef
}

// NewMocker creates a new mocker.
func NewMocker(ctx context.Context, t *testing.T) Mocker {
	t.Helper()

	mocker := Mocker{
		t:             t,
		deployManager: testutil.NewDeployManager(t),
		homeBackend:   simulated.NewSimulatedBackend(ctx, t),
	}

	_, mocker.homeContract = mocker.deployManager.GetHome(ctx, mocker.homeBackend)
	return mocker
}
