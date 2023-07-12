package backend_test

import (
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/testutil"

	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/util"
	"k8s.io/apimachinery/pkg/util/sets"
)

func (b *BackendSuite) TestLogsInRange() {
	testBackend := geth.NewEmbeddedBackend(b.GetTestContext(), b.T())
	// start an omnirpc proxy and run 10 test transactions so we can batch call blocks 1-10
	var wg sync.WaitGroup
	wg.Add(2)

	const desiredBlockHeight = 10

	var commonAddress common.Address
	go func() {
		defer wg.Done()
		newContract, err := testutil.PopulateWithLogs(b.GetTestContext(), testBackend, desiredBlockHeight, b.T(), b.manager)
		Nil(b.T(), err)
		commonAddress = *newContract
	}()

	var host string
	go func() {
		defer wg.Done()
		host = testutil.StartOmnirpcServer(b.GetTestContext(), testBackend, b.T())
	}()

	wg.Wait()

	scribeBackend, err := backend.DialBackend(b.GetTestContext(), host, b.metrics)
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
	res, err := backend.GetLogsInRange(b.GetTestContext(), scribeBackend, []common.Address{commonAddress}, chainID.Uint64(), blockRanges)
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

func (b *BackendSuite) TestLogsInRangeWithMultipleContracts() {
	testBackend := geth.NewEmbeddedBackend(b.GetTestContext(), b.T())
	// start an omnirpc proxy and run 10 test transactions so we can batch call blocks 1-10
	var wg sync.WaitGroup
	wg.Add(2)

	const desiredBlockHeight = 10

	var contractAddress1 common.Address
	var contractAddress2 common.Address

	go func() {
		defer wg.Done()
		newContract1, newContract2, err := testutil.MultiContractPopulateWithLogs(b.GetTestContext(), testBackend, desiredBlockHeight, b.T(), b.manager)
		Nil(b.T(), err)
		contractAddress1 = *newContract1
		contractAddress2 = *newContract2
	}()

	var host string
	go func() {
		defer wg.Done()
		host = testutil.StartOmnirpcServer(b.GetTestContext(), testBackend, b.T())
	}()

	wg.Wait()

	scribeBackend, err := backend.DialBackend(b.GetTestContext(), host, b.metrics)
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
	res, err := backend.GetLogsInRange(b.GetTestContext(), scribeBackend, []common.Address{contractAddress1, contractAddress2}, chainID.Uint64(), blockRanges)
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
	testIntRange := backend.MakeRange(0, 4)
	Equal(t, []int{0, 1, 2, 3, 4}, testIntRange)

	testUintRange := backend.MakeRange(uint16(10), uint16(12))
	Equal(t, testUintRange, []uint16{10, 11, 12})
}

func (b *BackendSuite) TestBlockHashesInRange() {
	testBackend := geth.NewEmbeddedBackend(b.GetTestContext(), b.T())

	// start an omnirpc proxy and run 10 test tranactions so we can batch call blocks
	//  1-10
	var wg sync.WaitGroup
	wg.Add(2)

	const desiredBlockHeight = 10

	go func() {
		defer wg.Done()
		testutil.ReachBlockHeight(b.GetTestContext(), testBackend, desiredBlockHeight, b.T())
	}()

	var host string
	go func() {
		defer wg.Done()
		host = testutil.StartOmnirpcServer(b.GetTestContext(), testBackend, b.T())
	}()

	wg.Wait()

	scribeBackend, err := backend.DialBackend(b.GetTestContext(), host, b.metrics)
	Nil(b.T(), err)

	res, err := backend.BlockHashesInRange(b.GetTestContext(), scribeBackend, 1, 10)
	Nil(b.T(), err)

	// use to make sure we don't double use values
	intSet := sets.NewInt64()

	itr := res.Iterator()

	for !itr.Done() {
		index, _, _ := itr.Next()
		Falsef(b.T(), intSet.Has(int64(index)), "%d appears at least twice", index)

	}
}
