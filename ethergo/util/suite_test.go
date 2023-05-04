package util_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// UtilSuite defines the basic test suite.
type UtilSuite struct {
	*testsuite.TestSuite
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewUtilSuite(tb testing.TB) *UtilSuite {
	tb.Helper()
	return &UtilSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestUtilSuite(t *testing.T) {
	suite.Run(t, NewUtilSuite(t))
}
