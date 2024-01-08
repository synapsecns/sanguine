package fastbridge_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
	"math/big"
	"testing"
)

// FastBridgeSuite tests the basic test suite.
type FastBridgeSuite struct {
	*testsuite.TestSuite
	backend backends.SimulatedTestBackend
	manager *testutil.DeployManager
}

// NewFastBridgeSuite creates a new FastBridge suite.
func NewFastBridgeSuite(tb testing.TB) *FastBridgeSuite {
	tb.Helper()
	return &FastBridgeSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestFastBridgeSuite(t *testing.T) {
	suite.Run(t, NewFastBridgeSuite(t))
}

func (s *FastBridgeSuite) SetupTest() {
	s.TestSuite.SetupTest()
	s.backend = simulated.NewSimulatedBackendWithChainID(s.GetTestContext(), s.T(), big.NewInt(1))
	s.manager = testutil.NewDeployManager(s.T())
}

// TestStatusEnum makes sure.
func (s *FastBridgeSuite) TestStatusEnum() {
	_, fb := s.manager.GetMockFastBridge(s.GetTestContext(), s.backend)
	for _, status := range fastbridge.GetAllBridgeStatuses() {
		solstatus, err := fb.GetEnumKeyByValue(&bind.CallOpts{Context: s.GetTestContext()}, status.Int())
		s.Require().NoError(err, "error getting enum key by value")
		s.Require().Equal(solstatus, status.String(), "status %s does not match. BridgeStatus enums out of sync.", status)
	}
}
