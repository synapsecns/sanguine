package node_test

import (
	"crypto/sha256"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/committee/db"
	"github.com/synapsecns/sanguine/core"
)

func (n *NodeSuite) TestNodeSuite() {
	// get the user we're going to test as
	auth := n.originChain.GetTxContext(n.GetTestContext(), nil)
	_, originDB := n.deployManager.GetInterchainDB(n.GetTestContext(), n.originChain)

	fee, err := originDB.GetInterchainFee(
		&bind.CallOpts{Context: n.GetSuiteContext()},
		n.destChain.GetBigChainID().Uint64(),
		[]common.Address{n.originModule.Address()},
	)
	n.Require().NoError(err)
	auth.TransactOpts.Value = core.CopyBigInt(fee)

	// test a single verification
	tx, err := originDB.WriteEntryWithVerification(
		auth.TransactOpts,
		n.destChain.GetBigChainID().Uint64(),
		sha256.Sum256([]byte("fat")),
		[]common.Address{n.originModule.Address()},
	)
	n.Require().NoError(err)
	// wait for the transaction to be mined
	n.originChain.WaitForConfirmation(n.GetTestContext(), tx)
	recp, err := n.originChain.TransactionReceipt(n.GetTestContext(), tx.Hash())
	n.Require().NoError(err)
	n.Require().Equal(uint64(1), recp.Status)

	n.Eventually(func() bool {
		// mine a block
		n.originChain.GetTxContext(n.GetTestContext(), nil)

		var resStatus []db.SignRequest
		for _, node := range n.nodes {
			resStatus, err = node.DB().GetQuoteResultsByStatus(n.GetTestContext(), db.Completed)
			n.Require().NoError(err)
		}

		return len(resStatus) > 0
	})
}

// func (n *NodeSuite) TestLotsOfVerifications() {
// 	// get the user we're going to test as
// 	auth := n.originChain.GetTxContext(n.GetTestContext(), nil)
// 	_, originDB := n.deployManager.GetInterchainDB(n.GetTestContext(), n.originChain)
//
// 	fee, err := originDB.GetInterchainFee(
// 		&bind.CallOpts{Context: n.GetSuiteContext()},
// 		n.destChain.GetBigChainID().Uint64(),
// 		[]common.Address{n.originModule.Address()},
// 	)
// 	n.Require().NoError(err)
// 	auth.TransactOpts.Value = core.CopyBigInt(fee)
//
// 	for i := 0; i < 10; i++ {
// 		// fire off 100 simultenous transactions
// 		go func() {
// 			tx, err := originDB.WriteEntryWithVerification(
// 				auth.TransactOpts,
// 				n.destChain.GetBigChainID().Uint64(),
// 				sha256.Sum256([]byte("fat")),
// 				[]common.Address{n.originModule.Address()},
// 			)
// 			n.Require().NoError(err)
// 			// wait for the transaction to be mined
// 			n.originChain.WaitForConfirmation(n.GetTestContext(), tx)
// 			recp, err := n.originChain.TransactionReceipt(n.GetTestContext(), tx.Hash())
// 			n.Require().NoError(err)
// 			n.Require().Equal(uint64(1), recp.Status)
// 		}()
// 	}
// }
//
// func (n *NodeSuite) getUserAndOriginDB() (auth backends.AuthType, originDB *interchaindb.InterchainDBRef) {
// 	// get the user we're going to test as
// 	auth = n.originChain.GetTxContext(n.GetTestContext(), nil)
// 	_, originDB = n.deployManager.GetInterchainDB(n.GetTestContext(), n.originChain)
// 	return auth, originDB
// }
//
