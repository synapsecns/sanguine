package home_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/testutil"
	"github.com/synapsecns/synapse-node/testutils"
	"github.com/synapsecns/synapse-node/testutils/backends"
	"github.com/synapsecns/synapse-node/testutils/backends/simulated"
	"testing"
)

// HomeSuite is the home test suite.
type HomeSuite struct {
	*testutils.TestSuite
	homeContract *home.HomeRef
	testBackend  backends.SimulatedTestBackend
}

// NewHomeSuite creates a end-to-end test suite.
func NewHomeSuite(tb testing.TB) *HomeSuite {
	tb.Helper()
	return &HomeSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func (h *HomeSuite) SetupTest() {
	h.TestSuite.SetupTest()

	deployManager := testutil.NewDeployManager(h.T())

	h.testBackend = simulated.NewSimulatedBackend(h.GetTestContext(), h.T())
	_, h.homeContract = deployManager.GetHome(h.GetTestContext(), h.testBackend)
}

// TestHomeSuite runs the integration test suite.
func TestHomeSuite(t *testing.T) {
	suite.Run(t, NewHomeSuite(t))
}
