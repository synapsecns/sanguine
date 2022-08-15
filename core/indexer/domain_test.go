package indexer_test

import (
	"context"
	"crypto/rand"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/indexer"
	"github.com/synapsecns/sanguine/core/types"
	"github.com/synapsecns/synapse-node/testutils/utils"
	"math"
	"math/big"
	"testing"
	"time"
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
		recipientAddress:  common.BytesToHash(utils.NewMockAddress().Bytes()),
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

func (i IndexerSuite) TestSyncMessages() {
	_, lastBlock := i.NewTestDispatches(25)

	db, err := sqlite.NewSqliteStore(i.GetTestContext(), filet.TmpDir(i.T(), ""))
	Nil(i.T(), err)

	_, homeContract := i.deployManager.GetOrigin(i.GetTestContext(), i.testBackend)

	domainClient, err := evm.NewEVM(i.GetTestContext(), "test", config.DomainConfig{
		DomainID:              1,
		Type:                  types.EVM.String(),
		RequiredConfirmations: 0,
		OriginAddress:         homeContract.Address().String(),
		RPCUrl:                i.testBackend.RPCAddress(),
		StartHeight:           0,
	})
	Nil(i.T(), err)

	domainIndexer := indexer.NewDomainIndexer(db, domainClient, 0)

	go func() {
		ctx, cancel := context.WithTimeout(i.GetTestContext(), time.Second*20)
		defer cancel()

		// this will error because of context cancellation
		_ = domainIndexer.SyncMessages(ctx)
	}()

	// wait until all blocks are indexed
	i.Eventually(func() bool {
		time.Sleep(time.Second * 1)

		testHeight, _ := db.GetMessageLatestBlockEnd(i.GetTestContext(), domainClient.Config().DomainID)

		return testHeight >= lastBlock
	})

	// TODO: something w/ retrieve dispatches from db
}

func TestUint32Max(t *testing.T) {
	// fuzz
	for i := 0; i < 50; i++ {
		small := gofakeit.Uint32()
		// we can't assert greater then max
		if small == math.MaxUint32 {
			continue
		}

		larger := small + 1

		Equal(t, indexer.MaxUint32(small, larger), larger)
	}

	// edge case
	Equal(t, indexer.MaxUint32(math.MaxUint32, 4), uint32(math.MaxUint32))
}
