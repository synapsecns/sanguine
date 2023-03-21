package testcontracts_test

import (
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/explorer/testutil/testcontracts"
)

// GetDeployedContractsFromRegistry gets any registered contract types that are present in the registry.
func (s *SimulatedSuite) GetDeployedContractsFromRegistry(registry deployer.ContractRegistry) (deployedContracts map[int]contracts.ContractType) {
	deployedContracts = make(map[int]contracts.ContractType)

	for _, contractType := range testcontracts.AllContractTypes {
		if registry.IsContractDeployed(contractType) {
			deployedContracts[contractType.ID()] = contractType
		}
	}
	return deployedContracts
}

// TestDependencies asserts all dependencies are included in contracts.
// TODO: this should be included in ethergo.
func (s *SimulatedSuite) TestDependencies() {
	manager.AssertDependenciesCorrect(s.GetTestContext(), s.T(), func() manager.IDeployManager {
		return testcontracts.NewDeployManager(s.T())
	})
}
