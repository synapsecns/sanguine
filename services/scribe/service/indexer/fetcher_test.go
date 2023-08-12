package indexer_test

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	scribeTypes "github.com/synapsecns/sanguine/services/scribe/types"
	"time"

	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/chain/client/mocks"
	etherMocks "github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/util"
	"github.com/synapsecns/sanguine/services/scribe/service/indexer"
)

// TestFilterLogsMaxAttempts ensures after the maximum number of attempts, an error is returned.
func (x *IndexerSuite) TestFilterLogsMaxAttempts() {
	x.T().Skip("flake")
	chainID := big.NewInt(int64(1))
	simulatedChain := geth.NewEmbeddedBackendForChainID(x.GetTestContext(), x.T(), chainID)
	simulatedClient, err := backend.DialBackend(x.GetTestContext(), simulatedChain.RPCAddress(), x.metrics)
	Nil(x.T(), err)
	mockFilterer := new(mocks.EVMClient)
	contractAddress := etherMocks.MockAddress()
	config := &scribeTypes.IndexerConfig{
		ChainID:            1,
		GetLogsBatchAmount: 1,
		GetLogsRange:       1,
		Addresses:          []common.Address{contractAddress},
	}

	rangeFilter := indexer.NewLogFetcher(simulatedClient, big.NewInt(1), big.NewInt(10), config, true)

	// Use the range filterer created above to create a mock log filter.
	mockFilterer.
		On("FilterLogs", mock.Anything, mock.Anything).
		Return(nil, errors.New("I'm a test error"))
	chunks := []*util.Chunk{{
		StartBlock: big.NewInt(1),
		EndBlock:   big.NewInt(10),
	}}
	logInfo, err := rangeFilter.FetchLogs(x.GetTestContext(), chunks)
	Nil(x.T(), logInfo)
	NotNil(x.T(), err)
}

// TestGetChunkArr ensures that the batching orchestration function (collecting block range chunks into arrays) works properly.
func (x *IndexerSuite) TestGetChunkArr() {
	chainID := big.NewInt(int64(1))
	simulatedChain := geth.NewEmbeddedBackendForChainID(x.GetTestContext(), x.T(), chainID)
	simulatedClient, err := backend.DialBackend(x.GetTestContext(), simulatedChain.RPCAddress(), x.metrics)
	Nil(x.T(), err)
	contractAddress := etherMocks.MockAddress()
	config := &scribeTypes.IndexerConfig{
		ChainID:              1,
		ConcurrencyThreshold: 1,
		GetLogsBatchAmount:   1,
		GetLogsRange:         1,
		Addresses:            []common.Address{contractAddress},
	}

	startBlock := int64(1)
	endBlock := int64(10)

	rangeFilter := indexer.NewLogFetcher(simulatedClient, big.NewInt(startBlock), big.NewInt(endBlock), config, true)

	numberOfRequests := int64(0)
	for i := int64(0); i < endBlock; i++ {
		chunks := rangeFilter.GetChunkArr()
		if len(chunks) == 0 {
			break
		}
		Equal(x.T(), len(chunks), int(config.GetLogsBatchAmount))
		numberOfRequests++
	}
	Equal(x.T(), numberOfRequests, endBlock)

	// Test with a larger batch size
	config.GetLogsBatchAmount = 4
	rangeFilter = indexer.NewLogFetcher(simulatedClient, big.NewInt(1), big.NewInt(10), config, true)
	numberOfRequests = int64(0)
	loopCount := endBlock/int64(config.GetLogsBatchAmount) + 1
	for i := int64(0); i < loopCount; i++ {
		chunks := rangeFilter.GetChunkArr()
		if len(chunks) == 0 {
			break
		}
		if i < loopCount-1 {
			Equal(x.T(), len(chunks), int(config.GetLogsBatchAmount))
		} else {
			Equal(x.T(), len(chunks), int(endBlock%int64(config.GetLogsBatchAmount)))
		}
		numberOfRequests++
	}
	Equal(x.T(), numberOfRequests, loopCount)

	// Test with a larger range size
	config.GetLogsRange = 2
	rangeFilter = indexer.NewLogFetcher(simulatedClient, big.NewInt(1), big.NewInt(10), config, true)
	numberOfRequests = int64(0)
	loopCount = endBlock/int64(config.GetLogsBatchAmount*config.GetLogsRange) + 1
	for i := int64(0); i < loopCount; i++ {
		chunks := rangeFilter.GetChunkArr()
		if len(chunks) == 0 {
			break
		}
		if i < loopCount-1 {
			Equal(x.T(), len(chunks), int(config.GetLogsBatchAmount))
		} else {
			Equal(x.T(), len(chunks), 1)
		}
		numberOfRequests++
	}
	Equal(x.T(), numberOfRequests, loopCount)
}

// TestGetChunkArr ensures that the batching orchestration function (collecting block range chunks into arrays) works properly.
func (x *IndexerSuite) TestFetchLogs() {
	testBackend := geth.NewEmbeddedBackend(x.GetTestContext(), x.T())
	// start an omnirpc proxy and run 10 test transactions so we can batch call blocks 1-10
	var wg sync.WaitGroup
	var testChainHandler *testutil.TestChainHandler
	var err error
	wg.Add(2)

	const desiredBlockHeight = 10

	go func() {
		defer wg.Done()
		testChainHandler, err = testutil.PopulateWithLogs(x.GetTestContext(), x.T(), testBackend, desiredBlockHeight, []*testutil.DeployManager{x.manager})
		Nil(x.T(), err)
	}()

	var host string
	go func() {
		defer wg.Done()
		host = testutil.StartOmnirpcServer(x.GetTestContext(), x.T(), testBackend)
	}()

	wg.Wait()

	scribeBackend, err := backend.DialBackend(x.GetTestContext(), host, x.metrics)
	Nil(x.T(), err)

	chunks := []*util.Chunk{
		{
			StartBlock: big.NewInt(1),
			EndBlock:   big.NewInt(2),
		},
		{
			StartBlock: big.NewInt(3),
			EndBlock:   big.NewInt(4),
		},
		{
			StartBlock: big.NewInt(5),
			EndBlock:   big.NewInt(6),
		},
		{
			StartBlock: big.NewInt(7),
			EndBlock:   big.NewInt(8),
		},
		{
			StartBlock: big.NewInt(9),
			EndBlock:   big.NewInt(10),
		},
	}
	chainID, err := scribeBackend.ChainID(x.GetTestContext())
	Nil(x.T(), err)
	config := &scribeTypes.IndexerConfig{
		ChainID:              uint32(chainID.Uint64()),
		ConcurrencyThreshold: 1,
		GetLogsBatchAmount:   1,
		GetLogsRange:         2,
		Addresses:            testChainHandler.Addresses,
	}
	rangeFilter := indexer.NewLogFetcher(scribeBackend, big.NewInt(1), big.NewInt(desiredBlockHeight), config, true)
	logs, err := rangeFilter.FetchLogs(x.GetTestContext(), chunks)
	Nil(x.T(), err)
	Equal(x.T(), 2, len(logs))

	cancelCtx, cancel := context.WithCancel(x.GetTestContext())
	cancel()

	_, err = rangeFilter.FetchLogs(cancelCtx, chunks)
	NotNil(x.T(), err)
	Contains(x.T(), err.Error(), "context was canceled")
}

// TestFetchLogsHighVolume tests the behavior of populating and consuming logs from the log fetcher in block ranges with many logs.
func (x *IndexerSuite) TestFetchLogsHighVolume() {
	testBackend := geth.NewEmbeddedBackend(x.GetTestContext(), x.T())
	// start an omnirpc proxy and run 10 test transactions so we can batch call blocks 1-10
	var err error
	host := testutil.StartOmnirpcServer(x.GetTestContext(), x.T(), testBackend)

	scribeBackend, err := backend.DialBackend(x.GetTestContext(), host, x.metrics)
	Nil(x.T(), err)

	chainID, err := scribeBackend.ChainID(x.GetTestContext())
	Nil(x.T(), err)
	config := &scribeTypes.IndexerConfig{
		ChainID:              uint32(chainID.Uint64()),
		ConcurrencyThreshold: 1,
		GetLogsBatchAmount:   1,
		GetLogsRange:         2,
		StoreConcurrency:     6,
		Addresses:            []common.Address{common.BigToAddress(big.NewInt(1))},
	}
	logFetcher := indexer.NewLogFetcher(scribeBackend, big.NewInt(1), big.NewInt(1000), config, true)

	logsChan := logFetcher.GetFetchedLogsChan()

	addContext, addCancel := context.WithTimeout(x.GetTestContext(), 20*time.Second)
	defer addCancel()
	numLogs := 0
	go func() {
		for {
			select {
			case <-addContext.Done():
				// test done
				close(*logsChan)
				return

			case <-time.After(10 * time.Millisecond):
				// add a log
				randomTxHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
				randomLog := testutil.MakeRandomLog(randomTxHash)
				*logsChan <- randomLog
				numLogs++
				// check buffer
				GreaterOrEqual(x.T(), config.StoreConcurrency, len(*logsChan))
			}
		}
	}()
	var collectedLogs []types.Log
	for {
		select {
		case <-x.GetTestContext().Done():
			Error(x.T(), fmt.Errorf("test context was canceled"))
		case <-time.After(1000 * time.Millisecond):
			log, ok := <-*logsChan
			if !ok {
				goto Done
			}
			collectedLogs = append(collectedLogs, log)
		}
	}
Done:
	Equal(x.T(), numLogs, len(collectedLogs))
}
