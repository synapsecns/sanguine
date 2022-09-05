package observer_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// ObserverSuite defines the basic test suite.
type ObserverSuite struct {
	*testsuite.TestSuite
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *ObserverSuite {
	tb.Helper()
	return &ObserverSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestObserverSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
