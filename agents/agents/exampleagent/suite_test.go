package exampleagent_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
)

// ExampleAgentSuite tests the example agent.
type ExampleAgentSuite struct {
	*testutil.SimulatedBackendsTestSuite
}

// NewGuardSuite creates a new guard suite.
func NewExampleAgentSuite(tb testing.TB) *ExampleAgentSuite {
	tb.Helper()

	return &ExampleAgentSuite{
		SimulatedBackendsTestSuite: testutil.NewSimulatedBackendsTestSuite(tb),
	}
}

func (u *ExampleAgentSuite) SetupTest() {
	chainwatcher.PollInterval = time.Second

	u.SimulatedBackendsTestSuite.SetupTest()
}

func TestExampleAgentSuite(t *testing.T) {
	suite.Run(t, NewExampleAgentSuite(t))
}
