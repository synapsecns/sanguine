package agentsintegration_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
)

// AgentsIntegrationSuite tests all the agents working together.
type AgentsIntegrationSuite struct {
	*testutil.SimulatedBackendsTestSuite
}

// NewAgentsIntegrationSuite creates a new agents integration suite.
func NewAgentsIntegrationSuite(tb testing.TB) *AgentsIntegrationSuite {
	tb.Helper()

	return &AgentsIntegrationSuite{
		SimulatedBackendsTestSuite: testutil.NewSimulatedBackendsTestSuite(tb),
	}
}

func (u *AgentsIntegrationSuite) SetupTest() {
	chainwatcher.PollInterval = time.Second

	u.SimulatedBackendsTestSuite.SetupTest()

	if os.Getenv("CI") == "" {
		u.SetTestTimeout(time.Minute * 3)
	}
}

func (u *AgentsIntegrationSuite) SetupSuite() {
	u.SimulatedBackendsTestSuite.SetupSuite()
}

func TestAgentsIntegrationSuite(t *testing.T) {
	suite.Run(t, NewAgentsIntegrationSuite(t))
}
