package executor_test

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchaindb"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/interchainmodulemock"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/optionslibexport"
	"github.com/synapsecns/sanguine/sin-executor/testutil"
)

func (i *InterchainSuite) TestE2E() {
	auth := i.originChain.GetTxContext(i.GetTestContext(), nil)

	// TODO: should be mock app
	receiver := i.addressToBytes32(i.deployManager.Get(i.GetTestContext(), i.destChain, testutil.InterchainAppMock).Address())

	_, optionsLib := i.deployManager.GetOptionsLib(i.GetTestContext(), i.originChain)

	encodedOptions, err := optionsLib.EncodeOptions(&bind.CallOpts{Context: i.GetTestContext()}, optionslibexport.OptionsV1{
		GasLimit:   big.NewInt(100000),
		GasAirdrop: big.NewInt(0),
	})
	i.Require().NoError(err)

	originModule := i.deployManager.Get(i.GetTestContext(), i.originChain, testutil.InterchainModuleMock)
	tx, err := i.originModule.InterchainSend(auth.TransactOpts, receiver, i.destChain.GetBigChainID(), []byte("hello"), encodedOptions, []common.Address{originModule.Address()})
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

		_, originDB := i.deployManager.GetInterchainDB(i.GetTestContext(), i.originChain)

		destContext := i.destChain.GetTxContext(i.GetTestContext(), nil)
		mockTX, err := destModule.MockVerifyEntry(destContext.TransactOpts, originDB.Address(), interchainmodulemock.InterchainEntry{
			SrcChainId: written.SrcChainId,
			DbNonce:    written.DbNonce,
			SrcWriter:  written.SrcWriter,
			DataHash:   written.DataHash,
		})
		i.Require().NoError(err)
		didMock = true

		i.destChain.WaitForConfirmation(i.GetTestContext(), mockTX)
	}

	i.Require().True(didMock)

	go func() {
		for {
			time.Sleep(time.Second * 3)
			// do some txes to up block.tmie
			i.originChain.GetFundedAccount(i.GetTestContext(), big.NewInt(1))
			i.destChain.GetFundedAccount(i.GetTestContext(), big.NewInt(1))
		}
	}()

	time.Sleep(time.Minute * 9)
}
