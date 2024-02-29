package exampleagent_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
)

func (u ExampleAgentSuite) TestExampleAgentSimulatedTestSuite() {
	NotNil(u.T(), u.SimulatedBackendsTestSuite)

	notaryStatus, err := u.BondingManagerOnSummit.AgentStatus(&bind.CallOpts{Context: u.GetTestContext()}, u.NotaryBondedSigner.Address())
	Nil(u.T(), err)
	Equal(u.T(), uint32(u.TestBackendDestination.GetChainID()), notaryStatus.Domain)

	notaryStatusFromEVM, err := u.SummitDomainClient.BondingManager().GetAgentStatus(u.GetTestContext(), u.NotaryBondedSigner.Address())
	Nil(u.T(), err)
	Equal(u.T(), notaryStatusFromEVM.Domain(), uint32(u.TestBackendDestination.GetChainID()))

	guardStatus, err := u.SummitContract.AgentStatus(&bind.CallOpts{Context: u.GetTestContext()}, u.GuardBondedSigner.Address())
	Nil(u.T(), err)
	Equal(u.T(), uint32(0), guardStatus.Domain)

	guardStatusFromEVM, err := u.SummitDomainClient.BondingManager().GetAgentStatus(u.GetTestContext(), u.GuardBondedSigner.Address())
	Nil(u.T(), err)
	Equal(u.T(), guardStatusFromEVM.Domain(), uint32(0))
}
