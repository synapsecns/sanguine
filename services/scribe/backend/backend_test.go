package backend_test

import (
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/testutil"

	"math/big"
	"sync"
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/util"
	"k8s.io/apimachinery/pkg/util/sets"
)

func (b *BackendSuite) TestLogsInRange() {
	const desiredBlockHeight = 10

	var testChainHandler *testutil.TestChainHandler
	var err error
	var wg sync.WaitGroup

	wg.Add(2)
	testBackend := geth.NewEmbeddedBackend(b.GetTestContext(), b.T())
	go func() {
		defer wg.Done()
		testChainHandler, err = testutil.PopulateWithLogs(b.GetTestContext(), b.T(), testBackend, desiredBlockHeight, []*testutil.DeployManager{b.manager})
		Nil(b.T(), err)
	}()

	var host string
	go func() {
		defer wg.Done()
		host = testutil.StartOmnirpcServer(b.GetTestContext(), b.T(), testBackend)
	}()

	wg.Wait()

	scribeBackend, err := backend.DialBackend(b.GetTestContext(), host, b.metrics)
	Nil(b.T(), err)

	chainID, err := scribeBackend.ChainID(b.GetTestContext())
	Nil(b.T(), err)

	lastBlock, err := testBackend.BlockNumber(b.GetTestContext())
	Nil(b.T(), err)

	iterator := util.NewChunkIterator(big.NewInt(int64(1)), big.NewInt(int64(lastBlock)), 1, true)

	var blockRanges []*util.Chunk
	blockRange := iterator.NextChunk()

	for blockRange != nil {
		blockRanges = append(blockRanges, blockRange)
		blockRange = iterator.NextChunk()
	}

	res, err := backend.GetLogsInRange(b.GetTestContext(), scribeBackend, testChainHandler.Addresses, chainID.Uint64(), blockRanges, nil)
	Nil(b.T(), err)

	// use to make sure we don't double use values
	intSet := sets.NewInt64()

	itr := res.Iterator()

	numLogs := 0
	for !itr.Done() {
		index, chunk := itr.Next()

		Falsef(b.T(), intSet.Has(int64(index)), "%d appears at least twice", index)
		intSet.Insert(int64(index))
		NotNil(b.T(), chunk)
		for range *chunk {
			numLogs++
		}
	}
	Equal(b.T(), int(testChainHandler.EventsEmitted[testChainHandler.Addresses[0]]), numLogs)
}

func (b *BackendSuite) TestLogsInRangeWithMultipleContracts() {
	const desiredBlockHeight = 10

	var testChainHandler *testutil.TestChainHandler
	var err error
	var wg sync.WaitGroup

	wg.Add(2)
	testBackend := geth.NewEmbeddedBackend(b.GetTestContext(), b.T())

	managerB := testutil.NewDeployManager(b.T())
	managerC := testutil.NewDeployManager(b.T())
	managers := []*testutil.DeployManager{b.manager, managerB, managerC}

	go func() {
		defer wg.Done()
		testChainHandler, err = testutil.PopulateWithLogs(b.GetTestContext(), b.T(), testBackend, desiredBlockHeight, managers)
		Nil(b.T(), err)
	}()

	var host string
	go func() {
		defer wg.Done()
		host = testutil.StartOmnirpcServer(b.GetTestContext(), b.T(), testBackend)
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
	res, err := backend.GetLogsInRange(b.GetTestContext(), scribeBackend, testChainHandler.Addresses, chainID.Uint64(), blockRanges, nil)
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
			logs[logAddr]++
			numLogs++
		}
	}
	Equal(b.T(), len(managers), numLogs)

	// Check if there's a log for each of the contracts
	for i := range testChainHandler.Addresses {
		Equal(b.T(), int(testChainHandler.EventsEmitted[testChainHandler.Addresses[i]]), logs[testChainHandler.Addresses[i].String()])
	}
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
		err := testutil.ReachBlockHeight(b.GetTestContext(), b.T(), testBackend, desiredBlockHeight)
		Nil(b.T(), err)
	}()

	var host string
	go func() {
		defer wg.Done()
		host = testutil.StartOmnirpcServer(b.GetTestContext(), b.T(), testBackend)
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
