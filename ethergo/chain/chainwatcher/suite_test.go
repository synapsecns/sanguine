package chainwatcher_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// ChainWatcherSuite defines the basic chain suite.
type ChainWatcherSuite struct {
	*testsuite.TestSuite
}

// NewChainWatcherSuite creates a new chain testing suite.
func NewChainWatcherSuite(tb testing.TB) *ChainWatcherSuite {
	tb.Helper()
	return &ChainWatcherSuite{testsuite.NewTestSuite(tb)}
}

func TestContractWatcherSuite(t *testing.T) {
	suite.Run(t, NewChainWatcherSuite(t))
}
