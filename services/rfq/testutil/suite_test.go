package testutil

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// TestUtilSuite tests the basic test suite.
type TestUtilSuite struct {
	*testsuite.TestSuite
}

// NewTestUtilSuite creates a new testutil suite.
func NewTestUtilSuite(tb testing.TB) *TestUtilSuite {
	tb.Helper()
	return &TestUtilSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestTestUtilSuite(t *testing.T) {
	suite.Run(t)
}
