package indexer_test

import (
	"crypto/rand"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"math/big"
	"testing"
)

// TestDispatch is a test dispatch call.
type TestDispatch struct {
	// domain we're sending to
	domain uint32
	// recipient address on the other chain
	recipientAddress common.Hash
	// raw message
	message []byte
	// optimisticSeconds is the optimistic second count
	optimisticSeconds uint32
	// tips defines the tips to use
	tips types.Tips
}

func NewTestDispatch(tb testing.TB) TestDispatch {
	tb.Helper()

	tips := types.NewTips(
		randomTip(tb),
		randomTip(tb),
		randomTip(tb),
		randomTip(tb),
	)

	return TestDispatch{
		domain:            gofakeit.Uint32(),
		recipientAddress:  common.BytesToHash(mocks.MockAddress().Bytes()),
		message:           []byte(gofakeit.Paragraph(4, 1, 4, " ")),
		optimisticSeconds: gofakeit.Uint32(),
		tips:              tips,
	}
}

// randomTip is a helper method for generating random tip that is less than 1 gwei
// see:  https://stackoverflow.com/a/45428754
func randomTip(tb testing.TB) *big.Int {
	tb.Helper()

	// Max random value = 1 eth
	max := new(big.Int)

	max.Exp(big.NewInt(2), big.NewInt(9), nil).Sub(max, big.NewInt(1))

	// Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	Nil(tb, err)

	return n
}

// Call calls dispatch and returns the block number.
func (d TestDispatch) Call(i IndexerSuite) (blockNumber uint32) {
	auth := i.testBackend.GetTxContext(i.GetTestContext(), nil)

	auth.TransactOpts.Value = types.TotalTips(d.tips)

	encodedTips, err := types.EncodeTips(d.tips)
	Nil(i.T(), err)

	tx, err := i.originContract.Dispatch(auth.TransactOpts, d.domain, d.recipientAddress, d.optimisticSeconds, encodedTips, d.message)
	Nil(i.T(), err)

	i.testBackend.WaitForConfirmation(i.GetTestContext(), tx)

	txReceipt, err := i.testBackend.TransactionReceipt(i.GetTestContext(), tx.Hash())
	Nil(i.T(), err)

	return uint32(txReceipt.BlockNumber.Uint64())
}

func (i IndexerSuite) NewTestDispatches(dispatchCount int) (testDispatches []TestDispatch, lastBlock uint32) {
	for iter := 0; iter < dispatchCount; iter++ {
		testDispatch := NewTestDispatch(i.T())
		lastBlock = testDispatch.Call(i)

		testDispatches = append(testDispatches, testDispatch)
	}

	return testDispatches, lastBlock
}
