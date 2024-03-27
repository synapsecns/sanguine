package executor_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/optionslibexport"
	"github.com/synapsecns/sanguine/sin-executor/db"
	"math/big"
	"time"

	"github.com/synapsecns/sanguine/sin-executor/contracts/interchaindb"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/interchainmodulemock"
	"github.com/synapsecns/sanguine/sin-executor/testutil"
)

func (i *InterchainSuite) TestE2E() {
	auth := i.originChain.GetTxContext(i.GetTestContext(), nil)

	message := []byte("hello")

	_, optionsLib := i.deployManager.GetOptionsLib(i.GetTestContext(), i.originChain)
	// must match values in the contract
	encodedOptions, err := optionsLib.EncodeOptions(&bind.CallOpts{Context: i.GetTestContext()}, optionslibexport.OptionsV1{
		GasLimit:   big.NewInt(200000),
		GasAirdrop: big.NewInt(0),
	})
	i.Require().NoError(err)

	_, originClient := i.deployManager.GetInterchainClient(i.GetTestContext(), i.originChain)
	interchainFee, err := originClient.GetInterchainFee(&bind.CallOpts{Context: i.GetTestContext()},
		i.destChain.GetBigChainID(), i.deployManager.Get(i.GetTestContext(), i.originChain, testutil.ExecutionService).Address(), []common.Address{i.deployManager.Get(i.GetTestContext(), i.originChain, testutil.InterchainModuleMock).Address()}, encodedOptions, message)
	i.Require().NoError(err)

	_, appMock := i.deployManager.GetInterchainAppMock(i.GetTestContext(), i.originChain)

	auth.TransactOpts.Value = interchainFee
	tx, err := appMock.SendMessage(auth.TransactOpts, i.destChain.GetBigChainID(), big.NewInt(1_000_000), message)
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
		written, err := parser.ParseInterchainBatchVerificationRequested(*log)
		if err != nil {
			continue
		}
		_, destDB := i.deployManager.GetInterchainDB(i.GetTestContext(), i.destChain)

		destContext := i.destChain.GetTxContext(i.GetTestContext(), nil)
		mockTX, err := destModule.MockVerifyRemoteBatch(destContext.TransactOpts, destDB.Address(), interchainmodulemock.InterchainBatch{
			SrcChainId: i.originChain.GetBigChainID(),
			DbNonce:    written.DbNonce,
			BatchRoot:  written.BatchRoot,
		})
		i.Require().NoError(err)
		didMock = true

		i.destChain.WaitForConfirmation(i.GetTestContext(), mockTX)
	}

	fmt.Printf("cast run %s --rpc-url %s/rpc/1 \n", recp.TxHash, i.omnirpcURL)
	i.Require().True(didMock)

	go func() {
		for {
			time.Sleep(time.Second * 3)
			// do some txes to up block.time
			i.originChain.GetFundedAccount(i.GetTestContext(), big.NewInt(1))
			i.destChain.GetFundedAccount(i.GetTestContext(), big.NewInt(1))
		}
	}()

	i.Eventually(func() bool {
		status, err := i.executor.DB().GetInterchainTXsByStatus(i.GetTestContext(), db.Executed)
		i.Require().NoError(err)

		return 1 == len(status)
	})
}
