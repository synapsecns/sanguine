package rpcinfo_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type LatencySuite struct {
	*testsuite.TestSuite
}

// NewLatencySuite creates a end-to-end test suite.
func NewLatencySuite(tb testing.TB) *LatencySuite {
	tb.Helper()
	return &LatencySuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestLatencySuite(t *testing.T) {
	suite.Run(t, NewLatencySuite(t))
}
