package watcher_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// WatcherSuite defines the basic chain suite.
type WatcherSuite struct {
	*testsuite.TestSuite
}

// NewWatcherSuite creates a new chain testing suite.
func NewWatcherSuite(tb testing.TB) *WatcherSuite {
	tb.Helper()
	return &WatcherSuite{testsuite.NewTestSuite(tb)}
}

func TestContractWatcherSuite(t *testing.T) {
	suite.Run(t, NewWatcherSuite(t))
}
