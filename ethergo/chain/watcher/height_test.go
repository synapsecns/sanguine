package watcher_test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/observer"
	"github.com/synapsecns/sanguine/ethergo/chain/watcher"
	"github.com/teivah/onecontext"
	"math/big"
	"reflect"
	"sync"
	"time"
)

// HeightContext is a context subscription created for a given height
// in the context of this tests, it represents weather or not the block subscriber was able to fetch the test.
type HeightContext struct {
	//nolint: containedctx
	context.Context
	//nolint: unused
	height uint64
}

// TestBlockHeightSubscriber tests the block height subscriber by mocking the
// BlockSubscriber and testing concurrency.
func (s *WatcherSuite) TestBlockHeightSubscriber() {
	s.T().Skip("this flakes on ci")

	// set test wide timeout
	ctx, cancelParent := context.WithTimeout(s.GetTestContext(), time.Second*30)
	defer cancelParent()

	initialHeight := core.CopyBigInt(params.MainnetChainConfig.LondonBlock)
	mockSubscription := NewMockBlockSubscriber(ctx, *initialHeight)
	subscriber := watcher.NewBlockHeightWatcher(ctx, 0, mockSubscription)
	heightChan := subscriber.Subscribe()
	defer subscriber.Unsubscribe(heightChan)

	var reachedLastIteration bool
	const testIterationCount = 50

	// ctxes represent context with cancel for the next block. We check this on
	// every block so we have tpo cancel the first one immediately
	var ctxes []HeightContext

	// escape the loop ont he last iteration
	subscriptionCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	var iteration uint64
OUTER:
	for {
		select {
		case height := <-heightChan:
			iteration++
			Equal(s.T(), height, initialHeight.Uint64()+iteration-1)
			if iteration == testIterationCount {
				reachedLastIteration = true
				cancel()
			} else {
				mockSubscription.PushHeight()
			}
		case <-subscriptionCtx.Done():
			break OUTER
		}
	}

	True(s.T(), reachedLastIteration)

	// wait for catch up if needed
	drainCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	for i, heightCtx := range ctxes {
		select {
		case <-ctx.Done():
			continue
		case <-drainCtx.Done():
			s.T().Errorf("all contexts should be completed. Stopped at height %d (iteration %d)", heightCtx.height, i)
		}
	}
}

// TestBlockHeightErr makes sure the block height watcher continues in the event of an error.
func (s *WatcherSuite) TestBlockHeightErr() {
	s.T().Skip("flake")

	initialHeight := *big.NewInt(1)
	// create mock subscription
	mockSubscription := NewMockBlockSubscriber(s.GetTestContext(), initialHeight)
	subscriber := watcher.NewBlockHeightWatcher(s.GetTestContext(), 0, mockSubscription)

	heightChan := subscriber.Subscribe()
	defer subscriber.Unsubscribe(heightChan)

	var iteration uint64
OUTER:
	for {
		select {
		case <-s.GetTestContext().Done():
			s.T().Errorf("could not finish block %T test", s.TestBlockHeightErr)
		case height := <-heightChan:
			iteration++
			if iteration == 5 || iteration == 7 {
				mockSubscription.TriggerErr()
			}
			Equal(s.T(), height, initialHeight.Uint64()+iteration-1)
			mockSubscription.PushHeight()
			// we triggered errors at 5 and 7 and still got all blocks
			if iteration == 10 {
				break OUTER
			}
		}
	}
}

// MockBlockSubscriber defines a struct that conforms to EthHeadClient
// to test different scenarios related to watching the chain.
type MockBlockSubscriber struct {
	//nolint: containedctx
	ctx context.Context
	// height to produce from. Blocks will increement this height
	height big.Int
	// used to prevent race conditions related to height
	producerMux sync.Mutex
	// every event on this channel represents a new block
	//nolint: staticcheck
	observer *observer.StringObserver
}

const (
	blocksTopic = "blocks"
	errTopic    = "err"
)

// MockSubscription represents a cancellable subscription.
type MockSubscription struct {
	//nolint: containedctx
	ctx        context.Context
	cancelFunc context.CancelFunc
	errChan    chan error
}

func NewMockSubscription(ctx context.Context) *MockSubscription {
	ctx, cancelFunc := context.WithCancel(ctx)
	mockSubscription := &MockSubscription{
		ctx:        ctx,
		errChan:    make(chan error),
		cancelFunc: cancelFunc,
	}

	// listen for context cancellation or unsubscribe
	go func() {
		<-ctx.Done()
		mockSubscription.errChan <- fmt.Errorf("context finished: %w", ctx.Err())
	}()

	return mockSubscription
}

func (m *MockSubscription) Unsubscribe() {
	m.cancelFunc()
}

func (m *MockSubscription) Err() <-chan error {
	return m.errChan
}

// TriggerErr returns an error to the subscription channel for testing.
func (m *MockSubscription) TriggerErr() {
	m.errChan <- errors.New("I'm an error")
}

var _ ethereum.Subscription = &MockSubscription{}

// NewMockBlockSubscriber creates a new block subscriber that allows clients
// to trigger NewBlock() events. Respects context by canceling the subscriber
// mocking at the end of the test. Only height is returned on the mocked blocks.
func NewMockBlockSubscriber(ctx context.Context, initialHeight big.Int) *MockBlockSubscriber {
	return &MockBlockSubscriber{
		ctx:         ctx,
		height:      initialHeight,
		producerMux: sync.Mutex{},
		observer:    observer.NewStringObserver(),
	}
}

// PushHeights pushes count heights.
func (m *MockBlockSubscriber) PushHeights(count int) {
	for i := 0; i < count; i++ {
		m.PushHeight()
	}
}

// PushHeight will push another height to the block subscriber channel for testing.
func (m *MockBlockSubscriber) PushHeight() {
	m.producerMux.Lock()
	defer m.producerMux.Unlock()
	m.height = *big.NewInt(m.height.Int64() + 1)
	m.observer.Emit(blocksTopic, m.height)
}

func (m *MockBlockSubscriber) TriggerErr() {
	m.observer.Emit(errTopic, nil)
}

// SubscribeNewHead creates a new subscription to head.
func (m *MockBlockSubscriber) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (sub ethereum.Subscription, err error) {
	ctx, cancel := onecontext.Merge(ctx, m.ctx)

	mockSubscription := NewMockSubscription(ctx)

	blockListener := make(chan interface{})
	m.observer.AddListener(blocksTopic, blockListener)
	errListener := make(chan interface{})
	m.observer.AddListener(errTopic, errListener)

	go func() {
		defer cancel()
		for {
			select {
			case errListener <- errListener:
				mockSubscription.TriggerErr()
			case rawHeight := <-blockListener:
				height, ok := rawHeight.(big.Int)
				if !ok {
					panic(fmt.Errorf("expected raw height to be %s, got %s", reflect.TypeOf(big.NewInt(0)), reflect.TypeOf(rawHeight)))
				}
				ch <- &types.Header{
					Number: &height,
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return mockSubscription, err
}

func (m *MockBlockSubscriber) LatestHeight(ctx context.Context) (uint64, error) {
	return m.height.Uint64(), nil
}

// BlockByHeader gets the block by header.
func (m *MockBlockSubscriber) HeaderByNumber(_ context.Context, number *big.Int) (*types.Header, error) {
	if number == nil {
		number = &m.height
	}
	return &types.Header{
		Number: number,
	}, nil
}
