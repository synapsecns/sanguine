package evm_test

import (
	"bytes"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/synapse-node/testutils/utils"
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
		recipientAddress:  common.BytesToHash(utils.NewMockAddress().Bytes()),
		message:           []byte(gofakeit.Paragraph(4, 1, 4, " ")),
		optimisticSeconds: gofakeit.Uint32(),
	}
}

// Call calls dispatch and returns the block number.
func (d TestDispatch) Call(i ContractSuite) (blockNumber uint32) {
	auth := i.testBackend.GetTxContext(i.GetTestContext(), nil)

	tx, err := i.homeContract.Dispatch(auth.TransactOpts, d.domain, d.recipientAddress, d.optimisticSeconds, d.message)
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

func (i ContractSuite) TestFetchSortedHomeUpdates() {
	testDispatches, filterTo := i.NewTestDispatches(33)

	homeIndexer, err := evm.NewHomeContract(i.testBackend, i.homeContract.Address())
	Nil(i.T(), err)

	messages, err := homeIndexer.FetchSortedMessages(i.GetTestContext(), 0, filterTo)
	Nil(i.T(), err)

	Equal(i.T(), len(messages), len(testDispatches))

	for iter, message := range messages {
		testDispatch := testDispatches[iter]
		True(i.T(), bytes.Contains(message.Message(), testDispatch.message))
		Equal(i.T(), message.LeafIndex(), uint32(iter))
	}
}
