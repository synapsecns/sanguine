package summit_test

import (
	"testing"

	"github.com/synapsecns/sanguine/agents/testutil"

	"github.com/stretchr/testify/suite"
)

// SummitSuite is the summit test suite.
type SummitSuite struct {
	*testutil.SimulatedBackendsTestSuite
}

// NewSummitSuite creates an end-to-end test suite.
func NewSummitSuite(tb testing.TB) *SummitSuite {
	tb.Helper()
	return &SummitSuite{
		SimulatedBackendsTestSuite: testutil.NewSimulatedBackendsTestSuite(tb),
	}
}

// SetupTest sets up the test.
func (a *SummitSuite) SetupTest() {
	a.SimulatedBackendsTestSuite.SetupTest()
}

// TestSummitSuite runs the integration test suite.
func TestSummitSuite(t *testing.T) {
	suite.Run(t, NewSummitSuite(t))
}
