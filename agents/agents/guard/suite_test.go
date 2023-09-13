package guard_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
)

// GuardSuite tests the guard agent.
type GuardSuite struct {
	*testutil.SimulatedBackendsTestSuite
}

// NewGuardSuite creates a new guard suite.
func NewGuardSuite(tb testing.TB) *GuardSuite {
	tb.Helper()

	return &GuardSuite{
		SimulatedBackendsTestSuite: testutil.NewSimulatedBackendsTestSuite(tb),
	}
}

func (g *GuardSuite) SetupTest() {
	chainwatcher.PollInterval = time.Second

	g.SimulatedBackendsTestSuite.SetupTest()
}

func TestGuardSuite(t *testing.T) {
	suite.Run(t, NewGuardSuite(t))
}
