package simulated_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type SimulatedSuite struct {
	*testsuite.TestSuite
}

// NewSimulatedSuite creates a end-to-end test suite.
func NewSimulatedSuite(tb testing.TB) *SimulatedSuite {
	tb.Helper()
	return &SimulatedSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestSimulatedSuite(t *testing.T) {
	suite.Run(t, NewSimulatedSuite(t))
}
