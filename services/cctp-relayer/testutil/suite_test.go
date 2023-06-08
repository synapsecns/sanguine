package testutil_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/cctp-relayer/testutil"
	"testing"
)

// TestUtilSuite defines the basic test suite.
type TestUtilSuite struct {
	*testsuite.TestSuite
}

// NewTestUtilSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestUtilSuite(tb testing.TB) *TestUtilSuite {
	tb.Helper()
	return &TestUtilSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestTestUtilSuite(t *testing.T) {
	suite.Run(t, NewTestUtilSuite(t))
}

// TestDependencies asserts all dependencies are included in contracts.
func (s *TestUtilSuite) TestDependencies() {
	manager.AssertDependenciesCorrect(s.GetTestContext(), s.T(), func() manager.IDeployManager {
		return testutil.NewDeployManager(s.T())
	})
}
