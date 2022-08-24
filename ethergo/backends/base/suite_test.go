package base_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// BaseSuite is the base test suite.
type BaseSuite struct {
	*testsuite.TestSuite
}

// NewBaseSuite creates a end-to-end test suite.
func NewBaseSuite(tb testing.TB) *BaseSuite {
	tb.Helper()
	return &BaseSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

// TestBaseSuite runs the integration test suite.
func TestBaseSuite(t *testing.T) {
	suite.Run(t, NewBaseSuite(t))
}
