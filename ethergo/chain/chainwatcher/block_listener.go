package chainwatcher

import (
	"context"
)

// blockListener is a listener.
type blockListener struct {
	// producerChan is the channel the producer sends messages to
	// this should never block for more then the time it takes to add messages to the
	// listener channel
	producerChan chan uint64
	// listenerChan is the channel we send messeages to. To avoid goroutine leaks, this
	// should have at most 1 message at a time (although we don't buffer)
	listenerChan chan uint64
	// lastHeight is the last height received by the listener
	lastHeight uint64
	// ctx is the context of the listener. This should be canceled on an unsubsribe
	//nolint: containedctx
	ctx context.Context
	// cancelFunc is used to cancel the context of the listener. This should be called on unsubscribe
	cancelFunc context.CancelFunc
	// hasNewHeightChan notifier the listener chan producer of a new channel
	hasNewHeightChan chan bool
	// blockNotifierIsReady is used to wait for consumer to be ready to start the producer
	blockNotifierIsReady chan bool
	// initialHeight is the initial height of the block
	initialHeight *uint64
}

// startObserverLoop starts the process that listens on block chan and sends new
// this registers the watcher and handles cases where a consumer is not listening by draining unsent messages.
// separating producer makes leaks much easier to debug (see https://github.com/synapsecns/synapse-node/pull/100/)
// from backfiller logic makes these channels a lot easier to debug.
//
// Note: context here should be equivalent to b.ctx. We explitly set it so the linter detects context leaks.
func (b *blockListener) startObserverLoop(ctx context.Context) {
	// this needs to be separate so the has new height chan works (since channel receiver needs to be instant)
	go b.startBlockNotifier(ctx)
	go b.startNotifierListener(ctx)
}

// startBlockNotifier starts the producer consumer.
func (b *blockListener) startBlockNotifier(ctx context.Context) {
	// onFirstRun is a channel used to set consumerIsReady on first run.
	// because we immediately skip if messages can't be forwarded to the channel
	// we miss not forwarding the first block if startNotifierListener starts before
	// startBlockNotifier. We could add a default block instead, but that would have to run
	// every time and eat up cpu
	onFirstRun := make(chan bool)

	go func() {
		select {
		case <-ctx.Done():
			return
		case onFirstRun <- true:
			return
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-onFirstRun:
			// notify the notifierListener it can start
			b.blockNotifierIsReady <- true
		// take from the producer channel and adds them to the notifies hasNewHeightChan so it can batch messages
		case newHeight := <-b.producerChan:
			if b.initialHeight == nil {
				b.initialHeight = &newHeight
			}

			b.lastHeight = newHeight
			select {
			case b.hasNewHeightChan <- true:
			default:
				// producer is busy. skip it
			}
		}
	}
}

// startNotifierListener starts the listener on the producerChannel
// it sets hasNewHeightChan when a new height is received. It only starts after
// consumerIsReady is true.
func (b *blockListener) startNotifierListener(ctx context.Context) {
	select {
	case <-b.blockNotifierIsReady:
		// we're ready to go
	case <-ctx.Done():
		// context was canceled before we needed to start
		return
	}

	// lastSent is the last height sent to the channel
	var lastSent uint64
	firstSend := true

	for {
		select {
		case <-ctx.Done():
			return
		case <-b.hasNewHeightChan:
			if firstSend {
				// make sure the initial height is set so we know where to send from
				b.listenerChan <- *b.initialHeight
				firstSend = false
				lastSent = *b.initialHeight
			}

			// sent to channel. If receiver block here it's fine because this is only one goroutine
			for i := lastSent; i <= b.lastHeight; i++ {
				if i > lastSent {
					b.listenerChan <- i
					lastSent = i
				}
			}
		}
	}
}
