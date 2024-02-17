package executor_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/sin-executor/testutil"
	"time"
)

func (i *InterchainSuite) TestE2E() {

	auth := i.originChain.GetTxContext(i.GetTestContext(), nil)

	// TODO: should be mock app
	receiver := addressToBytes32(i.deployManager.Get(i.GetTestContext(), i.destChain, testutil.InterchainAppMock).Address())

	originModule := i.deployManager.Get(i.GetTestContext(), i.originChain, testutil.InterchainModuleMock)
	tx, err := i.originModule.InterchainSend(auth.TransactOpts, receiver, i.destChain.GetBigChainID(), []byte("hello"), []common.Address{originModule.Address()})
	i.Require().NoError(err)
	i.originChain.WaitForConfirmation(i.GetTestContext(), tx)

	time.Sleep(time.Second * 90)
}
