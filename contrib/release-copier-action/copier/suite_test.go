package copier_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// CopierSuite defines the basic test suite.
type CopierSuite struct {
	*testsuite.TestSuite
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *CopierSuite {
	tb.Helper()
	return &CopierSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestCopierSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
