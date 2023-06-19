package evm_test

import (
	"bytes"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/ethergo/mocks"
)

// TestSent is a test sent call.
type TestSent struct {
	// domain we're sending to
	domain uint32
	// recipient address on the other chain
	recipientAddress common.Hash
	// raw message
	message []byte
	// optimisticSeconds is the seconds count to use
	optimisticSeconds uint32
}

func NewTestSent(destinationID uint32) TestSent {
	return TestSent{
		domain:            destinationID,
		recipientAddress:  common.BytesToHash(mocks.MockAddress().Bytes()),
		message:           []byte(gofakeit.Paragraph(4, 1, 4, " ")),
		optimisticSeconds: gofakeit.Uint32(),
	}
}

// Call calls sent and returns the block number.
func (d TestSent) Call(i ContractSuite) (blockNumber uint32) {
	auth := i.TestBackendOrigin.GetTxContext(i.GetTestContext(), nil)

	paddedRequest := big.NewInt(0)
	tx, err := i.OriginContract.SendBaseMessage(auth.TransactOpts, d.domain, d.recipientAddress, d.optimisticSeconds, paddedRequest, d.message)
	Nil(i.T(), err)
	i.TestBackendOrigin.WaitForConfirmation(i.GetTestContext(), tx)

	txReceipt, err := i.TestBackendOrigin.TransactionReceipt(i.GetTestContext(), tx.Hash())
	Nil(i.T(), err)

	return uint32(txReceipt.BlockNumber.Uint64())
}

func (i ContractSuite) NewTestSents(sentCount int, destinationID uint32) (testSents []TestSent, lastBlock uint32) {
	for iter := 0; iter < sentCount; iter++ {
		testSent := NewTestSent(destinationID)
		lastBlock = testSent.Call(i)

		testSents = append(testSents, testSent)
	}

	return testSents, lastBlock
}

func (i ContractSuite) TestFetchSortedOriginUpdates() {
	// TODO (joeallen): FIX ME
	i.T().Skip()
	destinationDomain := uint32(i.TestBackendDestination.GetChainID())
	originIndexer, err := evm.NewOriginContract(i.GetTestContext(), i.TestBackendOrigin, i.OriginContract.Address())
	Nil(i.T(), err)

	testSents, filterTo := i.NewTestSents(15, destinationDomain)

	messages, err := originIndexer.FetchSortedMessages(i.GetTestContext(), 0, filterTo)
	Nil(i.T(), err)

	Equal(i.T(), len(messages), len(testSents))

	for iter, message := range messages {
		testSent := testSents[iter]
		// TODO: Update this check.
		True(i.T(), bytes.Contains(message.Body(), testSent.message))
	}
}
