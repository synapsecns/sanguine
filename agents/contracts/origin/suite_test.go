package origin_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/synapse-node/testutils"
	"testing"
)

// OriginSuite is the origin test suite.
type OriginSuite struct {
	*testutils.TestSuite
	originContract *origin.OriginRef
	testBackend    backends.SimulatedTestBackend
}

// NewOriginSuite creates a end-to-end test suite.
func NewOriginSuite(tb testing.TB) *OriginSuite {
	tb.Helper()
	return &OriginSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func (h *OriginSuite) SetupTest() {
	h.TestSuite.SetupTest()

	deployManager := testutil.NewDeployManager(h.T())

	h.testBackend = simulated.NewSimulatedBackend(h.GetTestContext(), h.T())
	_, h.originContract = deployManager.GetOrigin(h.GetTestContext(), h.testBackend)
}

// TestOriginSuite runs the integration test suite.
func TestOriginSuite(t *testing.T) {
	suite.Run(t, NewOriginSuite(t))
}
