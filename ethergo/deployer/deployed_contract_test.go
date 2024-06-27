package deployer_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	counter2 "github.com/synapsecns/sanguine/ethergo/examples/contracttests/counter"
	"math/big"
)

func (d *DeployerSuite) TestNewDeployedContract() {
	const chainID = 1
	backend := simulated.NewSimulatedBackendWithChainID(d.GetTestContext(), d.T(), big.NewInt(chainID))

	counterDeployer := backend.GetTxContext(d.GetTestContext(), nil)
	deploymentAddress, tx, _, err := counter2.DeployCounter(counterDeployer.TransactOpts, backend)
	Nil(d.T(), err)

	handle, err := counter2.NewCounterRef(deploymentAddress, backend)
	Nil(d.T(), err)

	deployed, err := deployer.NewDeployedContract(handle, tx)
	Nil(d.T(), err)

	Equal(d.T(), deploymentAddress, deployed.Address())
	Equal(d.T(), handle, deployed.ContractHandle())
	Equal(d.T(), counterDeployer.TransactOpts.From, deployed.Owner())
	Equal(d.T(), &counterDeployer.TransactOpts.From, deployed.OwnerPtr())
	Equal(d.T(), big.NewInt(chainID), deployed.ChainID())
	Equal(d.T(), tx, deployed.DeployTx())
}
