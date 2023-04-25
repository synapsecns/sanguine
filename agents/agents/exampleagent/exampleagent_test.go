package exampleagent_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
)

func (u ExampleAgentSuite) TestExampleAgentSimulatedTestSuite() {
	NotNil(u.T(), u.SimulatedBackendsTestSuite)
	agentStatus, err := u.DestinationContract.AgentStatus(&bind.CallOpts{Context: u.GetTestContext()}, u.NotaryBondedSigner.Address())
	Nil(u.T(), err)
	Equal(u.T(), uint32(u.TestBackendDestination.GetChainID()), agentStatus.Domain)

	agentStatusFromEVM, err := u.DestinationDomainClient.LightManager().GetAgentStatus(u.GetTestContext(), u.NotaryBondedSigner)
	Nil(u.T(), err)
	Equal(u.T(), agentStatusFromEVM.Domain(), uint32(u.TestBackendDestination.GetChainID()))
}
