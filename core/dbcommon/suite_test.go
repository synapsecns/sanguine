package dbcommon_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// CoreSuite defines the basic test suite.
type DbSuite struct {
	*testsuite.TestSuite
}

// NewTestDBSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestDBSuite(tb testing.TB) *DbSuite {
	tb.Helper()
	return &DbSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestDBSuite(t *testing.T) {
	suite.Run(t, NewTestDBSuite(t))
}
