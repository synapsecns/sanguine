package indexer_test

import (
	"context"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	scribeTypes "github.com/synapsecns/sanguine/services/scribe/types"

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
	"github.com/synapsecns/sanguine/services/scribe/scribe/indexer"
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
	}

	rangeFilter := indexer.NewLogFetcher([]common.Address{contractAddress}, simulatedClient, big.NewInt(1), big.NewInt(10), config)

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
	}

	startBlock := int64(1)
	endBlock := int64(10)

	rangeFilter := indexer.NewLogFetcher([]common.Address{contractAddress}, simulatedClient, big.NewInt(startBlock), big.NewInt(endBlock), config)

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
	rangeFilter = indexer.NewLogFetcher([]common.Address{contractAddress}, simulatedClient, big.NewInt(1), big.NewInt(10), config)
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
	rangeFilter = indexer.NewLogFetcher([]common.Address{contractAddress}, simulatedClient, big.NewInt(1), big.NewInt(10), config)
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
	var addresses []common.Address
	var err error
	wg.Add(2)

	const desiredBlockHeight = 10

	go func() {
		defer wg.Done()
		addresses, _, err = testutil.PopulateWithLogs(x.GetTestContext(), testBackend, desiredBlockHeight, x.T(), []*testutil.DeployManager{x.manager})
		Nil(x.T(), err)
	}()

	var host string
	go func() {
		defer wg.Done()
		host = testutil.StartOmnirpcServer(x.GetTestContext(), testBackend, x.T())
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
	}
	rangeFilter := indexer.NewLogFetcher(addresses, scribeBackend, big.NewInt(1), big.NewInt(desiredBlockHeight), config)
	logs, err := rangeFilter.FetchLogs(x.GetTestContext(), chunks)
	Nil(x.T(), err)
	Equal(x.T(), 2, len(logs))

	cancelCtx, cancel := context.WithCancel(x.GetTestContext())
	cancel()

	_, err = rangeFilter.FetchLogs(cancelCtx, chunks)
	NotNil(x.T(), err)
	Contains(x.T(), err.Error(), "context was canceled")
}
