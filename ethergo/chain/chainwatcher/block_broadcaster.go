package chainwatcher

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

// BlockBroadcaster handles communications between block subcribers and
// clients. It Differs from BlockHeightWatcher in that it does not handle actually
// feteching heights from the rpc or reconnecting. We expose this for client/swap
// to allow multi-rpc block boradcasters.
//
//nolint:containedctx
type BlockBroadcaster struct {
	// firstSet is whether or not height has been set before
	firstSet bool
	// blockChan is the new block channel
	blockChan chan uint64
	// lastHeight is the last height seen by the block height subscriber.
	lastHeight uint64
	// lastHeightMux is used for setting lastHeight.
	// TODO: this should be atomic
	lastHeightMux sync.RWMutex
	// subscriberMux enforces that only one Subscribe() is running at once.
	// A semaphore is used here since mutexes require a Unlock() where as
	// a semaphore allows us to error immediately.
	subscriberMux *semaphore.Weighted
	// blockListeners are listing for new blocks
	blockListeners []*blockListener
	// blockListenersMux is used to make sure events are not-emitted to no-longer existent
	// channels
	blockListenersMux sync.RWMutex
	// chainID for the block height watcher
	chainID uint64
	// ctx is the context object
	ctx context.Context
}

// NewBlockBroadcaster creates a new block broadcaster.
func NewBlockBroadcaster(ctx context.Context, chainID uint64) *BlockBroadcaster {
	return &BlockBroadcaster{
		ctx:           ctx,
		lastHeightMux: sync.RWMutex{},
		subscriberMux: semaphore.NewWeighted(1),
		chainID:       chainID,
		blockChan:     make(chan uint64),
		firstSet:      true,
	}
}

// Emit emits a new block height to all listeners.
func (b *BlockBroadcaster) Emit(height uint64) {
	b.blockListenersMux.RLock()
	defer b.blockListenersMux.RUnlock()
	for _, listener := range b.blockListeners {
		go func(handler chan uint64) {
			handler <- height
		}(listener.producerChan)
	}
}

// Subscribe creates a new block height subscriber.
func (b *BlockBroadcaster) Subscribe() <-chan uint64 {
	ctx, cancel := context.WithCancel(b.ctx)

	b.blockListenersMux.Lock()
	listener := blockListener{
		producerChan:         make(chan uint64),
		listenerChan:         make(chan uint64, 2048),
		hasNewHeightChan:     make(chan bool),
		blockNotifierIsReady: make(chan bool),
		lastHeight:           b.lastHeight,
		ctx:                  ctx,
		cancelFunc:           cancel,
	}
	b.blockListeners = append(b.blockListeners, &listener)
	b.blockListenersMux.Unlock()

	listener.startObserverLoop(b.ctx)

	return listener.listenerChan
}

// Unsubscribe removes a block height subscriber.
func (b *BlockBroadcaster) Unsubscribe(ch <-chan uint64) {
	b.blockListenersMux.Lock()
	defer b.blockListenersMux.Unlock()
	for i, listener := range b.blockListeners {
		if listener.producerChan == ch {
			listener.cancelFunc()
			b.blockListeners = append(b.blockListeners[:i], b.blockListeners[i+1:]...)
			break
		}
	}
}
