package watcher_test

import (
	"context"
	"math/big"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
	watcherMocks "github.com/synapsecns/sanguine/ethergo/chain/chainwatcher/mocks"
	chainMocks "github.com/synapsecns/sanguine/ethergo/chain/mocks"
	"github.com/synapsecns/sanguine/ethergo/chain/watcher"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"golang.org/x/sync/semaphore"
)

// filterLogsTracker tracks various calls attempted on filterLogs() mock.
type filterLogsTracker struct {
	tb              testing.TB
	maxCalledHeight uint64
	numberOfCalls   uint64
	calledBlocks    []uint64
	mux             sync.Mutex
}

// newFilterLogsTracker creates a filter logs tracker.
func newFilterLogsTracker(tb testing.TB) filterLogsTracker {
	tb.Helper()
	return filterLogsTracker{
		tb:              tb,
		maxCalledHeight: 0,
		numberOfCalls:   0,
	}
}

// update updates the filter logs tracker and performs some basic assertions.
func (f *filterLogsTracker) update(query ethereum.FilterQuery) {
	f.mux.Lock()
	defer f.mux.Unlock()
	True(f.tb, query.ToBlock.Uint64() >= query.FromBlock.Uint64())
	f.numberOfCalls++
	f.maxCalledHeight = query.ToBlock.Uint64()
	f.calledBlocks = append(f.calledBlocks, query.ToBlock.Uint64())
}

func (s *WatcherSuite) TestContractWatcherRetry() {
	if os.Getenv("CI") != "" {
		s.T().Skip("this flakes on ci. TODO fix this. This should never fail locally.")
	}
	const requiredConfs = 3

	logTracker := newFilterLogsTracker(s.T())

	mockBlockSubscription := NewMockBlockSubscriber(s.GetTestContext(), *big.NewInt(0))
	blockSubscriber := chainwatcher.NewBlockHeightWatcher(s.GetTestContext(), 0, mockBlockSubscription)

	mockEvmClient := new(chainMocks.Chain)
	mockEvmClient.On("GetBigChainID").Return(params.AllEthashProtocolChanges.ChainID)

	mockContract := mocks.MockAddress()

	eventLog := make(chan interface{})
	contractWatcher := watcher.NewTestContractWatcher(s.GetTestContext(), s.T(), mockEvmClient, blockSubscriber, requiredConfs)
	err := contractWatcher.ListenOnContract(s.GetTestContext(), mockContract.String(), eventLog)
	Nil(s.T(), err)

	// let's push some heights where nothing happens and mmake sure it doesn't call more than required confs
	mockEvmClient.On("FilterLogs", mock.Anything, mock.MatchedBy(func(filterQuery ethereum.FilterQuery) bool {
		shouldReturn := filterQuery.ToBlock.Uint64() <= 10 || filterQuery.ToBlock.Uint64() == 12
		if shouldReturn {
			logTracker.update(filterQuery)
		}
		return shouldReturn
	})).Return([]types.Log{}, nil)

	producedHeights := 10
	mockBlockSubscription.PushHeights(producedHeights)
	s.Eventually(func() bool {
		return logTracker.maxCalledHeight == uint64(producedHeights-requiredConfs)
	})

	// push up to the 10 and wait til we get there
	mockBlockSubscription.PushHeights(requiredConfs)
	s.Eventually(func() bool {
		return logTracker.maxCalledHeight == uint64(producedHeights)
	})

	hasSentErr := false
	mockEvmClient.On("FilterLogs", mock.Anything, mock.MatchedBy(func(filterQuery ethereum.FilterQuery) bool {
		hasSentErr = true
		shouldReturn := filterQuery.ToBlock.Uint64() == 11
		if shouldReturn {
			logTracker.update(filterQuery)
		}
		return shouldReturn
	})).Return([]types.Log{}, errors.New("I'm an error")).Once()
	// produce an error on the 11th block and make sure we recover

	mockBlockSubscription.PushHeights(2)

	// on subsequent calls produce successfully
	mockEvmClient.On("FilterLogs", mock.Anything, mock.MatchedBy(func(filterQuery ethereum.FilterQuery) bool {
		if !hasSentErr {
			return false
		}
		shouldReturn := filterQuery.ToBlock.Uint64() == 11
		if shouldReturn {
			logTracker.update(filterQuery)
		}
		return shouldReturn
	})).Return([]types.Log{{Address: mockContract}}, nil)

	s.Eventually(func() bool {
		return logTracker.maxCalledHeight == 12
	})
}

// TestListeners tests a scenario with more than 1 listener listening to the contract observer.
//
//nolint:gocognit,cyclop
func (s *WatcherSuite) TestListeners() {
	// timeout the test after thie period
	ctx, cancel := context.WithTimeout(s.GetTestContext(), 30*time.Second)
	defer cancel()

	contractWatcher := watcher.NewTestContractWatcher(ctx, s.T(), new(chainMocks.Chain), new(watcherMocks.BlockHeightWatcher), 0)

	mockContract := mocks.MockAddress()

	// eventCount is how many events to listen for
	const eventCount = 30
	// listener count is how many listeners to create
	const listenerCount = 10
	// testEvents are the producerChan produces and the listener verify happened
	testEvents := mocks.GetMockLogs(s.T(), eventCount)
	initializationSemaphore := semaphore.NewWeighted(listenerCount)
	// use a waitgroup to manage the listeners
	var wg sync.WaitGroup

	// create the listeners
	for i := 0; i < listenerCount; i++ {
		Nil(s.T(), initializationSemaphore.Acquire(ctx, 1))
		wg.Add(1)
		go func() {
			defer wg.Done()
			// verificationSlice is a copy of test events used to verify the listener received all logs
			verificationSlice := make([]types.Log, len(testEvents))
			copy(verificationSlice, testEvents)

			listener := make(chan interface{})
			contractWatcher.AddListener(mockContract, listener)
			initializationSemaphore.Release(1)

			for {
				select {
				case <-ctx.Done():
					Nil(s.T(), ctx.Err())
					return
				case rawEvent := <-listener:
					// convert the raw event to an event
					event, ok := rawEvent.(types.Log)
					if !ok {
						s.T().Error("could not decode event from channel")
					}

					// verify the event is in the verification slice
					for i, verifiedEvent := range verificationSlice {
						if reflect.DeepEqual(verifiedEvent, event) {
							// remove the event from the verification slice
							verificationSlice = append(verificationSlice[:i], verificationSlice[i+1:]...)
							break
						}
					}

					// all events have been verified
					if len(verificationSlice) == 0 {
						return
					}
				}
			}
		}()
	}

	// wait until all the listeners are initialized
	Nil(s.T(), initializationSemaphore.Acquire(ctx, listenerCount))
	producerChan := make(chan types.Log)
	err := contractWatcher.AddProducer(ctx, mockContract, producerChan)
	Nil(s.T(), err)
	for _, event := range testEvents {
		producerChan <- event
	}
	// attempt to add another producer to the same contract
	NotNil(s.T(), contractWatcher.AddProducer(ctx, mockContract, producerChan))

	wg.Wait()
}
