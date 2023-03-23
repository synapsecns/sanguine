package testutil_test

import (
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/manager"
)

// GetDeployedContractsFromRegistry gets any registered contract types that are present in the registry.
func (s *SimulatedSuite) GetDeployedContractsFromRegistry(registry deployer.ContractRegistry) (deployedContracts map[int]contracts.ContractType) {
	deployedContracts = make(map[int]contracts.ContractType)

	for _, contractType := range testutil.AllContractTypes {
		if registry.IsContractDeployed(contractType) {
			deployedContracts[contractType.ID()] = contractType
		}
	}
	return deployedContracts
}

// TestDependencies asserts all dependencies are included in contracts.
// TODO: this should be included in ethergo.
func (s *SimulatedSuite) TestDependencies() {
	// TODO (joe): Destination now depends on OriginType, so get this test working
	s.T().Skip("TODO")
	manager.AssertDependenciesCorrect(s.GetTestContext(), s.T(), func() manager.IDeployManager {
		return testutil.NewDeployManager(s.T())
	})
}
