package executor_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	agentsTestutil "github.com/synapsecns/sanguine/agents/testutil"
)

// ExecutorSuite tests the executor agent.
type ExecutorSuite struct {
	*agentsTestutil.SimulatedBackendsTestSuite
}

// NewExecutorSuite creates a new executor suite.
func NewExecutorSuite(tb testing.TB) *ExecutorSuite {
	tb.Helper()

	return &ExecutorSuite{
		SimulatedBackendsTestSuite: agentsTestutil.NewSimulatedBackendsTestSuite(tb),
	}
}

func (e *ExecutorSuite) SetupTest() {
	e.SimulatedBackendsTestSuite.SetupTest()
	e.SetTestTimeout(time.Minute * 5)
}

func TestExecutorSuite(t *testing.T) {
	suite.Run(t, NewExecutorSuite(t))
}
