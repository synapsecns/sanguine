package executor_test

import (
	"fmt"
	"math/big"
	"time"

	"github.com/synapsecns/sanguine/sin-executor/contracts/interchaindb"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/interchainmodulemock"
	"github.com/synapsecns/sanguine/sin-executor/testutil"
)

func (i *InterchainSuite) TestE2E() {
	auth := i.originChain.GetTxContext(i.GetTestContext(), nil)

	receiver := i.addressToBytes32(i.deployManager.Get(i.GetTestContext(), i.destChain, testutil.InterchainApp).Address())

	_, appMock := i.deployManager.GetInterchainAppMock(i.GetTestContext(), i.originChain)
	tx, err := appMock.Send(auth.TransactOpts, receiver, i.destChain.GetBigChainID(), []byte("hello"))
	i.Require().NoError(err)
	i.Require().NoError(err)
	i.originChain.WaitForConfirmation(i.GetTestContext(), tx)

	// get the receipt on the origin chain so we can call mock verify entry
	recp, err := i.originChain.TransactionReceipt(i.GetTestContext(), tx.Hash())
	i.Require().NoError(err)

	idb, _ := i.deployManager.GetInterchainDB(i.GetTestContext(), i.originChain)
	parser, err := interchaindb.NewInterchainDBFilterer(idb.Address(), i.originChain)
	i.Require().NoError(err)

	_, destModule := i.deployManager.GetInterchainModuleMock(i.GetTestContext(), i.destChain)

	didMock := false
	for _, log := range recp.Logs {
		written, err := parser.ParseInterchainEntryWritten(*log)
		if err != nil {
			continue
		}

		_, destDB := i.deployManager.GetInterchainDB(i.GetTestContext(), i.destChain)

		destContext := i.destChain.GetTxContext(i.GetTestContext(), nil)
		mockTX, err := destModule.MockVerifyEntry(destContext.TransactOpts, destDB.Address(), interchainmodulemock.InterchainEntry{
			SrcChainId: written.SrcChainId,
			DbNonce:    written.DbNonce,
			SrcWriter:  written.SrcWriter,
			DataHash:   written.DataHash,
		})
		i.Require().NoError(err)
		didMock = true

		i.destChain.WaitForConfirmation(i.GetTestContext(), mockTX)
		fmt.Print(mockTX.Hash())
	}
	i.Require().True(didMock)

	go func() {
		for {
			time.Sleep(time.Second * 3)
			// do some txes to up block.time
			i.originChain.GetFundedAccount(i.GetTestContext(), big.NewInt(1))
			i.destChain.GetFundedAccount(i.GetTestContext(), big.NewInt(1))
		}
	}()

	time.Sleep(time.Minute * 1)
}
