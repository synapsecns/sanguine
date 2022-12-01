package evm_test

import (
	"bytes"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/mocks"
)

// TestDispatch is a test dispatch call.
type TestDispatch struct {
	// domain we're sending to
	domain uint32
	// recipient address on the other chain
	recipientAddress common.Hash
	// raw message
	message []byte
	// optimisticSeconds is the seconds count to use
	optimisticSeconds uint32
}

func NewTestDispatch() TestDispatch {
	return TestDispatch{
		domain:            gofakeit.Uint32(),
		recipientAddress:  common.BytesToHash(mocks.MockAddress().Bytes()),
		message:           []byte(gofakeit.Paragraph(4, 1, 4, " ")),
		optimisticSeconds: gofakeit.Uint32(),
	}
}

// Call calls dispatch and returns the block number.
func (d TestDispatch) Call(i ContractSuite) (blockNumber uint32) {
	auth := i.testBackend.GetTxContext(i.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(i.T(), err)

	tx, err := i.originContract.Dispatch(auth.TransactOpts, d.domain, d.recipientAddress, d.optimisticSeconds, encodedTips, d.message)
	Nil(i.T(), err)
	i.testBackend.WaitForConfirmation(i.GetTestContext(), tx)

	txReceipt, err := i.testBackend.TransactionReceipt(i.GetTestContext(), tx.Hash())
	Nil(i.T(), err)

	return uint32(txReceipt.BlockNumber.Uint64())
}

func (i ContractSuite) NewTestDispatches(dispatchCount int) (testDispatches []TestDispatch, lastBlock uint32) {
	for iter := 0; iter < dispatchCount; iter++ {
		testDispatch := NewTestDispatch()
		lastBlock = testDispatch.Call(i)

		testDispatches = append(testDispatches, testDispatch)
	}

	return testDispatches, lastBlock
}

func (i ContractSuite) TestFetchSortedOriginUpdates() {
	// TODO (joe): Currently we are setting the notary in the origin contract, but eventually this will need
	// set the Notary per destination and also add Guards and the dispatch function would assert that the
	// destination has a Notary and there is at least one Guard

	originIndexer, err := evm.NewOriginContract(i.GetTestContext(), i.testBackend, i.originContract.Address())
	Nil(i.T(), err)

	ownerPtr, err := i.originContract.OriginCaller.Owner(&bind.CallOpts{Context: i.GetTestContext()})
	Nil(i.T(), err)

	originDomain, err := i.originContract.LocalDomain(&bind.CallOpts{Context: i.GetTestContext()})
	Nil(i.T(), err)

	originOwnerAuth := i.testBackend.GetTxContext(i.GetTestContext(), &ownerPtr)
	tx, err := i.originContract.AddNotary(originOwnerAuth.TransactOpts, originDomain, i.signer.Address())
	Nil(i.T(), err)
	i.testBackend.WaitForConfirmation(i.GetTestContext(), tx)

	notaries, err := i.originContract.AllNotaries(&bind.CallOpts{Context: i.GetTestContext()})
	Nil(i.T(), err)
	Len(i.T(), notaries, 1)

	testDispatches, filterTo := i.NewTestDispatches(33)

	messages, err := originIndexer.FetchSortedMessages(i.GetTestContext(), 0, filterTo)
	Nil(i.T(), err)

	Equal(i.T(), len(messages), len(testDispatches))

	for iter, message := range messages {
		testDispatch := testDispatches[iter]
		True(i.T(), bytes.Contains(message.Message(), testDispatch.message))
	}
}
