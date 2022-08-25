package chainlist_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// ChainlistSuite defines the basic test suite.
type ChainlistSuite struct {
	*testsuite.TestSuite
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *ChainlistSuite {
	tb.Helper()
	return &ChainlistSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestChainlistSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
