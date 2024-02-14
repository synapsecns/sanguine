package testutil_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
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
	suite.Run(t, NewTestUtilSuite(t))
}

func (s *TestUtilSuite) TestDependencies() {
	manager.AssertDependenciesCorrect(s.GetTestContext(), s.T(), func() manager.IDeployManager {
		return testutil.NewDeployManager(s.T())
	})
}
