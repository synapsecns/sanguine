package indexer_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/indexer"
	"github.com/synapsecns/sanguine/core/types"
	"github.com/synapsecns/synapse-node/testutils/utils"
	"math"
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
}

func NewTestDispatch() TestDispatch {
	return TestDispatch{
		domain:           gofakeit.Uint32(),
		recipientAddress: common.BytesToHash(utils.NewMockAddress().Bytes()),
		message:          []byte(gofakeit.Paragraph(4, 1, 4, " ")),
	}
}

// Call calls dispatch and returns the block number.
func (d TestDispatch) Call(i IndexerSuite) (blockNumber uint32) {
	auth := i.testBackend.GetTxContext(i.GetTestContext(), nil)

	tx, err := i.homeContract.Dispatch(auth.TransactOpts, d.domain, d.recipientAddress, d.message)
	Nil(i.T(), err)

	i.testBackend.WaitForConfirmation(i.GetTestContext(), tx)

	txReceipt, err := i.testBackend.TransactionReceipt(i.GetTestContext(), tx.Hash())
	Nil(i.T(), err)

	return uint32(txReceipt.BlockNumber.Uint64())
}

func (i IndexerSuite) NewTestDispatches(dispatchCount int) (testDispatches []TestDispatch, lastBlock uint32) {
	for iter := 0; iter < dispatchCount; iter++ {
		testDispatch := NewTestDispatch()
		lastBlock = testDispatch.Call(i)

		testDispatches = append(testDispatches, testDispatch)
	}

	return testDispatches, lastBlock
}

func (i IndexerSuite) TestSyncMessages() {
	i.NewTestDispatches(25)

	testDB, err := db.NewDB(filet.TmpDir(i.T(), ""), "")
	Nil(i.T(), err)

	_, xAppConfig := i.deployManager.GetXAppConfig(i.GetTestContext(), i.testBackend)

	domainClient, err := evm.NewEVM(i.GetTestContext(), "test", config.DomainConfig{
		DomainID:              1,
		Type:                  types.EVM.String(),
		RequiredConfirmations: 0,
		XAppConfigAddress:     xAppConfig.Address().String(),
		RPCUrl:                i.testBackend.RPCAddress(),
		StartHeight:           0,
	})
	Nil(i.T(), err)

	domainIndexer := indexer.NewDomainIndexer(testDB, domainClient)

	Nil(i.T(), domainIndexer.SyncMessages(i.GetTestContext()))
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
