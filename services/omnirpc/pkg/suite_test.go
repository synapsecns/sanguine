package pkg_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type RPCSuite struct {
	*testsuite.TestSuite
}

// NewRPCSuite creates a end-to-end test suite.
func NewRPCSuite(tb testing.TB) *RPCSuite {
	tb.Helper()
	return &RPCSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestRPCSuite(t *testing.T) {
	suite.Run(t, NewRPCSuite(t))
}
