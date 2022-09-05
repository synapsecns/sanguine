package testutil_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
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
	s.deployManager = testutil.NewDeployManager(s.T())
	wrappedBackend := s.testBackend
	registeredContracts := s.deployManager.GetContractRegistry(wrappedBackend).RegisteredDeployers()

	// test until all contacts are done
	for _, contract := range registeredContracts {
		s.deployManager = testutil.NewDeployManager(s.T())
		contractRegistry := s.deployManager.GetContractRegistry(wrappedBackend)
		Equal(s.T(), len(s.GetDeployedContractsFromRegistry(contractRegistry)), 0)

		// the contract is currently on the wrong backend, so we need to make it on the right backend
		dc := contractRegistry.Get(s.GetTestContext(), contract.ContractType())
		Equal(s.T(), dc.ChainID().String(), wrappedBackend.GetBigChainID().String())

		deployedContracts := s.GetDeployedContractsFromRegistry(contractRegistry)
		// make sure dependency count is equal (adding our own contract to there expected amount)
		Equal(s.T(), len(deployedContracts), len(contract.Dependencies())+1)
		for _, dep := range contract.Dependencies() {
			_, hasDep := deployedContracts[dep.ID()]
			True(s.T(), hasDep)
		}
	}
}
