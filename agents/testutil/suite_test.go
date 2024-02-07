package testutil_test

import (
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"math/big"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
)

// SimulatedSuite is used to test individual contract deployments to make sure other tests don't break.
type SimulatedSuite struct {
	*testsuite.TestSuite
	// testSynBackend is the test syn chain backend
	testSynBackend backends.SimulatedTestBackend
	// testBackend is the test backend
	testBackend backends.SimulatedTestBackend
	// deployManager is the deploy helper
	deployManager *testutil.DeployManager
}

// SetupTest sets up a test.
func (s *SimulatedSuite) SetupTest() {
	s.TestSuite.SetupTest()

	s.testSynBackend = simulated.NewSimulatedBackendWithChainID(s.GetTestContext(), s.T(), big.NewInt(int64(10)))
	s.testBackend = simulated.NewSimulatedBackendWithChainID(s.GetTestContext(), s.T(), big.NewInt(int64(11)))
	s.deployManager = testutil.NewDeployManager(s.T())
	s.deployManager.GetContractRegistry(s.testBackend)
	s.deployManager.GetContractRegistry(s.testSynBackend)
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
