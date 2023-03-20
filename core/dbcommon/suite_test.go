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

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *DbSuite {
	tb.Helper()
	return &DbSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestCommonSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
