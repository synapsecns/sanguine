package forker_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/example"
	"github.com/synapsecns/sanguine/ethergo/forker"
	"github.com/synapsecns/sanguine/ethergo/manager"
)

func (f *ForkSuite) TestFork() {
	// create an embedded backend
	backend := geth.NewEmbeddedBackend(f.GetTestContext(), f.T())
	deployer := manager.NewDeployerManager(f.T(), example.NewCounterDeployer)
	_ = deployer

	// deploy the counter contract

	err := forker.Fork(f.GetTestContext(), backend.HTTPEndpoint(), 10, func(client *ethclient.Client) {
		// execute a new transaction incrementing the counter

		// check the counter state is 1

		res, err := client.CodeAt(f.GetTestContext(), common.HexToAddress("0x7002B727Ef8F5571Cb5F9D70D13DBEEb4dFAe9d1"), nil)
		Nil(f.T(), err)

		fmt.Println(hexutil.Encode(res))
	})

	Nil(f.T(), err)

	// using the original backend, check the counter state is 0
}
