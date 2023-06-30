package backfill_test

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/chain/client/mocks"
	etherMocks "github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/util"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"math/big"
	"sync"
)

// TestFilterLogsMaxAttempts ensures after the maximum number of attempts, an error is returned.
func (b BackfillSuite) TestFilterLogsMaxAttempts() {
	b.T().Skip("flake")
	chainID := big.NewInt(int64(1))
	simulatedChain := geth.NewEmbeddedBackendForChainID(b.GetTestContext(), b.T(), chainID)
	simulatedClient, err := backfill.DialBackend(b.GetTestContext(), simulatedChain.RPCAddress(), b.metrics)
	Nil(b.T(), err)
	mockFilterer := new(mocks.EVMClient)
	contractAddress := etherMocks.MockAddress()
	config := &config.ChainConfig{
		ChainID:            1,
		GetLogsBatchAmount: 1,
		GetLogsRange:       1,
	}

	rangeFilter := backfill.NewLogFetcher(contractAddress, simulatedClient, big.NewInt(1), big.NewInt(10), config)

	// Use the range filterer created above to create a mock log filter.
	mockFilterer.
		On("FilterLogs", mock.Anything, mock.Anything).
		Return(nil, errors.New("I'm a test error"))
	chunks := []*util.Chunk{{
		StartBlock: big.NewInt(1),
		EndBlock:   big.NewInt(10),
	}}
	logInfo, err := rangeFilter.FetchLogs(b.GetTestContext(), chunks)
	Nil(b.T(), logInfo)
	NotNil(b.T(), err)
}

// TestGetChunkArr ensures that the batching orchestration function (collecting block range chunks into arrays) works properly.
func (b BackfillSuite) TestGetChunkArr() {
	chainID := big.NewInt(int64(1))
	simulatedChain := geth.NewEmbeddedBackendForChainID(b.GetTestContext(), b.T(), chainID)
	simulatedClient, err := backfill.DialBackend(b.GetTestContext(), simulatedChain.RPCAddress(), b.metrics)
	Nil(b.T(), err)
	contractAddress := etherMocks.MockAddress()
	config := &config.ChainConfig{
		ChainID:              1,
		ConcurrencyThreshold: 1,
		GetLogsBatchAmount:   1,
		GetLogsRange:         1,
	}

	startBlock := int64(1)
	endBlock := int64(10)

	rangeFilter := backfill.NewLogFetcher(contractAddress, simulatedClient, big.NewInt(startBlock), big.NewInt(endBlock), config)

	numberOfRequests := int64(0)
	for i := int64(0); i < endBlock; i++ {
		chunks := rangeFilter.GetChunkArr()
		if len(chunks) == 0 {
			break
		}
		Equal(b.T(), len(chunks), int(config.GetLogsBatchAmount))
		numberOfRequests++
	}
	Equal(b.T(), numberOfRequests, endBlock)

	// Test with a larger batch size
	config.GetLogsBatchAmount = 4
	rangeFilter = backfill.NewLogFetcher(contractAddress, simulatedClient, big.NewInt(1), big.NewInt(10), config)
	numberOfRequests = int64(0)
	loopCount := endBlock/int64(config.GetLogsBatchAmount) + 1
	for i := int64(0); i < loopCount; i++ {
		chunks := rangeFilter.GetChunkArr()
		if len(chunks) == 0 {
			break
		}
		if i < loopCount-1 {
			Equal(b.T(), len(chunks), int(config.GetLogsBatchAmount))
		} else {
			Equal(b.T(), len(chunks), int(endBlock%int64(config.GetLogsBatchAmount)))
		}
		numberOfRequests++
	}
	Equal(b.T(), numberOfRequests, loopCount)

	// Test with a larger range size
	config.GetLogsRange = 2
	rangeFilter = backfill.NewLogFetcher(contractAddress, simulatedClient, big.NewInt(1), big.NewInt(10), config)
	numberOfRequests = int64(0)
	loopCount = endBlock/int64(config.GetLogsBatchAmount*config.GetLogsRange) + 1
	for i := int64(0); i < loopCount; i++ {
		chunks := rangeFilter.GetChunkArr()
		if len(chunks) == 0 {
			break
		}
		if i < loopCount-1 {
			Equal(b.T(), len(chunks), int(config.GetLogsBatchAmount))
		} else {
			Equal(b.T(), len(chunks), 1)
		}
		numberOfRequests++
	}
	Equal(b.T(), numberOfRequests, loopCount)
}

// TestGetChunkArr ensures that the batching orchestration function (collecting block range chunks into arrays) works properly.
func (b BackfillSuite) TestFetchLogs() {
	testBackend := geth.NewEmbeddedBackend(b.GetTestContext(), b.T())
	// start an omnirpc proxy and run 10 test transactions so we can batch call blocks 1-10
	var wg sync.WaitGroup
	wg.Add(2)

	const desiredBlockHeight = 10

	var contractAddress common.Address
	go func() {
		defer wg.Done()
		contractAddress = b.PopuluateWithLogs(b.GetTestContext(), testBackend, desiredBlockHeight)
	}()

	var host string
	go func() {
		defer wg.Done()
		host = b.startOmnirpcServer(b.GetTestContext(), testBackend)
	}()

	wg.Wait()

	scribeBackend, err := backfill.DialBackend(b.GetTestContext(), host, b.metrics)
	Nil(b.T(), err)

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
	chainID, err := scribeBackend.ChainID(b.GetTestContext())
	Nil(b.T(), err)
	config := &config.ChainConfig{
		ChainID:              uint32(chainID.Uint64()),
		ConcurrencyThreshold: 1,
		GetLogsBatchAmount:   1,
		GetLogsRange:         2,
	}
	rangeFilter := backfill.NewLogFetcher(contractAddress, scribeBackend, big.NewInt(1), big.NewInt(desiredBlockHeight), config)
	logs, err := rangeFilter.FetchLogs(b.GetTestContext(), chunks)
	Nil(b.T(), err)
	Equal(b.T(), 2, len(logs))

	cancelCtx, cancel := context.WithCancel(b.GetTestContext())
	cancel()

	_, err = rangeFilter.FetchLogs(cancelCtx, chunks)
	NotNil(b.T(), err)
	Contains(b.T(), err.Error(), "context was canceled")
}
