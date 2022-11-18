package testcontracts_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/services/explorer/testutil/testcontracts"
	"testing"
)

// SimulatedSuite is used to test individual contract deployments to make sure other tests don't break.
type SimulatedSuite struct {
	*testsuite.TestSuite
	// testBackend is the test backend
	testBackend backends.SimulatedTestBackend
	// deployManager is the deploy helper
	deployManager *testcontracts.DeployManager
}

// SetupTest sets up a test.
func (s *SimulatedSuite) SetupTest() {
	s.TestSuite.SetupTest()

	s.testBackend = simulated.NewSimulatedBackend(s.GetTestContext(), s.T())
	s.deployManager = testcontracts.NewDeployManager(s.T())
	s.deployManager.GetContractRegistry(s.testBackend)
}

// NewSimulatedSuite creates a end-to-end test suite.
func NewSimulatedSuite(tb testing.TB) *SimulatedSuite {
	tb.Helper()
	return &SimulatedSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestSimulatedSuite(t *testing.T) {
	suite.Run(t, NewSimulatedSuite(t))
}
