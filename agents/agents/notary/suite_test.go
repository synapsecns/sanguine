package notary_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
)

// NotarySuite tests the notary agent.
type NotarySuite struct {
	*testutil.SimulatedBackendsTestSuite
}

// NewNotarySuite creates a new notary suite.
func NewNotarySuite(tb testing.TB) *NotarySuite {
	tb.Helper()

	return &NotarySuite{
		SimulatedBackendsTestSuite: testutil.NewSimulatedBackendsTestSuite(tb),
	}
}

func (u *NotarySuite) SetupTest() {
	chainwatcher.PollInterval = time.Second

	u.SimulatedBackendsTestSuite.SetupTest()
}

func TestNotarySuite(t *testing.T) {
	suite.Run(t, NewNotarySuite(t))
}
