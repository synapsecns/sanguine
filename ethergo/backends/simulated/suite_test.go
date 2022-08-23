package simulated_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"testing"
)

type SimulatedSuite struct {
	*testutils.TestSuite
}

// NewSimulatedSuite creates a end-to-end test suite.
func NewSimulatedSuite(tb testing.TB) *SimulatedSuite {
	tb.Helper()
	return &SimulatedSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func TestSimulatedSuite(t *testing.T) {
	suite.Run(t, NewSimulatedSuite(t))
}
