package destination_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/testutil"
)

// DestinationSuite is the destination test suite.
type DestinationSuite struct {
	*testutil.SimulatedBackendsTestSuite
}

// NewDestinationSuite creates a end-to-end test suite.
func NewDestinationSuite(tb testing.TB) *DestinationSuite {
	tb.Helper()
	return &DestinationSuite{
		SimulatedBackendsTestSuite: testutil.NewSimulatedBackendsTestSuite(tb),
	}
}

func (d *DestinationSuite) SetupTest() {
	d.SimulatedBackendsTestSuite.SetupTest()
}

// TestDestinationSuite runs the integration test suite.
func TestDestinationSuite(t *testing.T) {
	suite.Run(t, NewDestinationSuite(t))
}
