package backfill_test

import (
	"context"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/util"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"k8s.io/apimachinery/pkg/util/sets"
)

// startOmnirpcServer boots an omnirpc server for an rpc address.
// the url for this rpc is returned.
func (b *BackfillSuite) startOmnirpcServer(ctx context.Context, backend backends.SimulatedTestBackend) string {
	baseHost := testhelper.NewOmnirpcServer(ctx, b.T(), backend)
	return testhelper.GetURL(baseHost, backend)
}

// ReachBlockHeight reaches a block height on a backend.
func (b *BackfillSuite) ReachBlockHeight(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			b.T().Log(ctx.Err())
			return
		default:
			// continue
		}
		i++
		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))

		latestBlock, err := backend.BlockNumber(ctx)
		Nil(b.T(), err)

		if latestBlock >= desiredBlockHeight {
			return
		}
	}
}

// ReachBlockHeight reaches a block height on a backend.
func (b *BackfillSuite) PopuluateWithLogs(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64) common.Address {
	i := 0
	var address common.Address
	for {
		select {
		case <-ctx.Done():
			b.T().Log(ctx.Err())
			return address
		default:
			// continue
		}
		i++
		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))
		testContract, testRef := b.manager.GetTestContract(b.GetTestContext(), backend)
		address = testContract.Address()
		transactOpts := backend.GetTxContext(b.GetTestContext(), nil)
		tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		Nil(b.T(), err)
		backend.WaitForConfirmation(b.GetTestContext(), tx)

		latestBlock, err := backend.BlockNumber(ctx)
		Nil(b.T(), err)

		if latestBlock >= desiredBlockHeight {
			return address
		}
	}
}

// ReachBlockHeight reaches a block height on a backend.
func (b *BackfillSuite) PopuluateWithLogsWithDifferentContracts(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64) (common.Address, common.Address) {
	i := 0
	var address1 common.Address
	var address2 common.Address

	for {
		select {
		case <-ctx.Done():
			b.T().Log(ctx.Err())
			return address1, address2
		default:
			// continue
		}
		i++
		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))
		testContract1, testRef1 := b.manager.GetTestContract(b.GetTestContext(), backend)
		testContract2, testRef2 := b.manager.GetTestContract2(b.GetTestContext(), backend)

		address1 = testContract1.Address()
		address2 = testContract2.Address()

		transactOpts := backend.GetTxContext(b.GetTestContext(), nil)
		tx1, err := testRef1.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		Nil(b.T(), err)
		tx2, err := testRef2.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		Nil(b.T(), err)

		backend.WaitForConfirmation(b.GetTestContext(), tx1)
		backend.WaitForConfirmation(b.GetTestContext(), tx2)

		latestBlock, err := backend.BlockNumber(ctx)
		Nil(b.T(), err)

		if latestBlock >= desiredBlockHeight {
			return address1, address2
		}
	}
}
func (b *BackfillSuite) TestLogsInRange() {
	testBackend := geth.NewEmbeddedBackend(b.GetTestContext(), b.T())
	// start an omnirpc proxy and run 10 test transactions so we can batch call blocks 1-10
	var wg sync.WaitGroup
	wg.Add(2)

	const desiredBlockHeight = 10

	var commonAddress common.Address
	go func() {
		defer wg.Done()
		commonAddress = b.PopuluateWithLogs(b.GetTestContext(), testBackend, desiredBlockHeight)
	}()

	var host string
	go func() {
		defer wg.Done()
		host = b.startOmnirpcServer(b.GetTestContext(), testBackend)
	}()

	wg.Wait()

	scribeBackend, err := backfill.DialBackend(b.GetTestContext(), host, b.metrics)
	Nil(b.T(), err)

	chainID, err := scribeBackend.ChainID(b.GetTestContext())
	Nil(b.T(), err)
	iterator := util.NewChunkIterator(big.NewInt(int64(1)), big.NewInt(int64(desiredBlockHeight)), 1, true)

	var blockRanges []*util.Chunk
	blockRange := iterator.NextChunk()

	for blockRange != nil {
		blockRanges = append(blockRanges, blockRange)
		blockRange = iterator.NextChunk()
	}
	res, err := backfill.GetLogsInRange(b.GetTestContext(), scribeBackend, []common.Address{commonAddress}, chainID.Uint64(), blockRanges)
	Nil(b.T(), err)

	// use to make sure we don't double use values
	intSet := sets.NewInt64()

	itr := res.Iterator()

	numLogs := 0
	for !itr.Done() {
		numLogs++
		index, chunk := itr.Next()

		Falsef(b.T(), intSet.Has(int64(index)), "%d appears at least twice", index)
		intSet.Insert(int64(index))
		NotNil(b.T(), chunk)
		for range *chunk {
			numLogs++
		}
	}
	Equal(b.T(), 4, numLogs)

}

func (b *BackfillSuite) TestLogsInRangeWithMultipleContracts() {
	testBackend := geth.NewEmbeddedBackend(b.GetTestContext(), b.T())
	// start an omnirpc proxy and run 10 test transactions so we can batch call blocks 1-10
	var wg sync.WaitGroup
	wg.Add(2)

	const desiredBlockHeight = 10

	var contractAddress1 common.Address
	var contractAddress2 common.Address

	go func() {
		defer wg.Done()
		contractAddress1, contractAddress2 = b.PopuluateWithLogsWithDifferentContracts(b.GetTestContext(), testBackend, desiredBlockHeight)

	}()

	var host string
	go func() {
		defer wg.Done()
		host = b.startOmnirpcServer(b.GetTestContext(), testBackend)
	}()

	wg.Wait()

	scribeBackend, err := backfill.DialBackend(b.GetTestContext(), host, b.metrics)
	Nil(b.T(), err)

	chainID, err := scribeBackend.ChainID(b.GetTestContext())
	Nil(b.T(), err)
	iterator := util.NewChunkIterator(big.NewInt(int64(1)), big.NewInt(int64(desiredBlockHeight)), 1, true)

	var blockRanges []*util.Chunk
	blockRange := iterator.NextChunk()

	for blockRange != nil {
		blockRanges = append(blockRanges, blockRange)
		blockRange = iterator.NextChunk()
	}
	res, err := backfill.GetLogsInRange(b.GetTestContext(), scribeBackend, []common.Address{contractAddress1, contractAddress2}, chainID.Uint64(), blockRanges)
	Nil(b.T(), err)

	// use to make sure we don't double use values
	intSet := sets.NewInt64()

	itr := res.Iterator()

	numLogs := 0
	logs := make(map[string]int)
	for !itr.Done() {
		index, chunk := itr.Next()

		Falsef(b.T(), intSet.Has(int64(index)), "%d appears at least twice", index)
		intSet.Insert(int64(index))
		NotNil(b.T(), chunk)
		for i := range *chunk {
			logAddr := (*chunk)[i].Address.String()
			logs[logAddr] = logs[logAddr] + 1
			numLogs++
		}
	}
	Equal(b.T(), 2, numLogs)
	Equal(b.T(), 1, logs[contractAddress1.String()])
	Equal(b.T(), 1, logs[contractAddress2.String()])

}

func TestMakeRange(t *testing.T) {
	testIntRange := backfill.MakeRange(0, 4)
	Equal(t, []int{0, 1, 2, 3, 4}, testIntRange)

	testUintRange := backfill.MakeRange(uint16(10), uint16(12))
	Equal(t, testUintRange, []uint16{10, 11, 12})
}

func (b *BackfillSuite) TestBlockHashesInRange() {
	testBackend := geth.NewEmbeddedBackend(b.GetTestContext(), b.T())

	// start an omnirpc proxy and run 10 test tranactions so we can batch call blocks
	//  1-10
	var wg sync.WaitGroup
	wg.Add(2)

	const desiredBlockHeight = 10

	go func() {
		defer wg.Done()
		b.ReachBlockHeight(b.GetTestContext(), testBackend, desiredBlockHeight)
	}()

	var host string
	go func() {
		defer wg.Done()
		host = b.startOmnirpcServer(b.GetTestContext(), testBackend)
	}()

	wg.Wait()

	scribeBackend, err := backfill.DialBackend(b.GetTestContext(), host, b.metrics)
	Nil(b.T(), err)

	res, err := backfill.BlockHashesInRange(b.GetTestContext(), scribeBackend, 1, 10)
	Nil(b.T(), err)

	// use to make sure we don't double use values
	intSet := sets.NewInt64()

	itr := res.Iterator()

	for !itr.Done() {
		index, _, _ := itr.Next()
		Falsef(b.T(), intSet.Has(int64(index)), "%d appears at least twice", index)

	}
}
