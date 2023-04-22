package evm_test

import (
	"bytes"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
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

func NewTestDispatch(destinationID uint32) TestDispatch {
	return TestDispatch{
		domain:            destinationID,
		recipientAddress:  common.BytesToHash(mocks.MockAddress().Bytes()),
		message:           []byte(gofakeit.Paragraph(4, 1, 4, " ")),
		optimisticSeconds: gofakeit.Uint32(),
	}
}

// Call calls dispatch and returns the block number.
func (d TestDispatch) Call(i ContractSuite) (blockNumber uint32) {
	auth := i.TestBackendOrigin.GetTxContext(i.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(i.T(), err)

	paddedTips := new(big.Int).SetBytes(encodedTips)
	paddedRequest := new(big.Int).SetBytes(d.message)

	// TODO (joe): Figure out what to set for content
	tx, err := i.OriginContract.SendBaseMessage(auth.TransactOpts, d.domain, d.recipientAddress, d.optimisticSeconds, paddedTips, paddedRequest, []byte{})
	Nil(i.T(), err)
	i.TestBackendOrigin.WaitForConfirmation(i.GetTestContext(), tx)

	txReceipt, err := i.TestBackendOrigin.TransactionReceipt(i.GetTestContext(), tx.Hash())
	Nil(i.T(), err)

	return uint32(txReceipt.BlockNumber.Uint64())
}

func (i ContractSuite) NewTestDispatches(dispatchCount int, destinationID uint32) (testDispatches []TestDispatch, lastBlock uint32) {
	for iter := 0; iter < dispatchCount; iter++ {
		testDispatch := NewTestDispatch(destinationID)
		lastBlock = testDispatch.Call(i)

		testDispatches = append(testDispatches, testDispatch)
	}

	return testDispatches, lastBlock
}

func (i ContractSuite) TestFetchSortedOriginUpdates() {
	// TODO (joeallen): FIX ME
	i.T().Skip()
	destinationDomain := uint32(i.TestBackendDestination.GetChainID())
	originIndexer, err := evm.NewOriginContract(i.GetTestContext(), i.TestBackendOrigin, i.OriginContract.Address())
	Nil(i.T(), err)

	testDispatches, filterTo := i.NewTestDispatches(15, destinationDomain)

	messages, err := originIndexer.FetchSortedMessages(i.GetTestContext(), 0, filterTo)
	Nil(i.T(), err)

	Equal(i.T(), len(messages), len(testDispatches))

	for iter, message := range messages {
		testDispatch := testDispatches[iter]
		True(i.T(), bytes.Contains(message.Message(), testDispatch.message))
	}
}
