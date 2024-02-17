package executor_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/optionslibexport"
	"github.com/synapsecns/sanguine/sin-executor/testutil"
	"math/big"
	"time"
)

func (i *InterchainSuite) TestE2E() {

	auth := i.originChain.GetTxContext(i.GetTestContext(), nil)

	// TODO: should be mock app
	receiver := i.addressToBytes32(i.deployManager.Get(i.GetTestContext(), i.destChain, testutil.InterchainAppMock).Address())

	_, optionsLib := i.deployManager.GetOptionsLib(i.GetTestContext(), i.originChain)

	encodedOptions, err := optionsLib.EncodeOptions(&bind.CallOpts{Context: i.GetTestContext()}, optionslibexport.OptionsLibOptions{
		Version:    0,
		GasLimit:   big.NewInt(100000),
		GasAirdrop: big.NewInt(0),
	})
	i.Require().NoError(err)

	originModule := i.deployManager.Get(i.GetTestContext(), i.originChain, testutil.InterchainModuleMock)
	tx, err := i.originModule.InterchainSend(auth.TransactOpts, receiver, i.destChain.GetBigChainID(), []byte("hello"), encodedOptions, []common.Address{originModule.Address()})
	i.Require().NoError(err)
	i.originChain.WaitForConfirmation(i.GetTestContext(), tx)

	time.Sleep(time.Minute * 9)
}
