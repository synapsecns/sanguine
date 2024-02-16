package node_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"math/big"
	"time"
)

func (n *NodeSuite) TestNodeSuite() {
	// get the user we're going to test as
	auth := n.originChain.GetTxContext(n.GetTestContext(), nil)

	// now, authenticate as the owner so our user can spoof itself as the interchain contract
	originInfo, _ := n.deployManager.GetSynapseModule(n.GetTestContext(), n.originChain)
	ownerAuth := n.originChain.GetTxContext(n.GetTestContext(), originInfo.OwnerPtr())

	tx, err := n.originModule.SetInterchainDB(ownerAuth.TransactOpts, auth.From)
	n.Require().NoError(err)
	n.originChain.WaitForConfirmation(n.GetTestContext(), tx)

	// set value of tx to module fee
	auth.TransactOpts.Value, err = n.originModule.GetModuleFee(&bind.CallOpts{Context: n.GetSuiteContext()}, n.destChain.GetBigChainID())
	n.Require().NoError(err)

	tx, err = n.originModule.RequestVerification(auth.TransactOpts, n.destChain.GetBigChainID(), synapsemodule.InterchainEntry{
		SrcChainId:  n.originChain.GetBigChainID(),
		SrcWriter:   [32]byte{},
		WriterNonce: big.NewInt(0),
		DataHash: [32]byte{
			0x01,
		},
	})
	n.Require().NoError(err)

	// wait for the transaction to be mined
	n.originChain.WaitForConfirmation(n.GetTestContext(), tx)
	time.Sleep(time.Minute * 10)
}
