package testutil_test

import (
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
)

// TestDependencies asserts all dependencies are included in contracts.
// TODO: this should be included in ethergo.
func (s *SimulatedSuite) TestDependencies() {
	manager.AssertDependenciesCorrect(s.GetTestContext(), s.T(), func() manager.IDeployManager {
		return testutil.NewDeployManager(s.T())
	})
}
