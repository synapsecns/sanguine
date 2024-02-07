package chainwatcher

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// BlockSubscriberClient defines a method for getting a subscription to the chain-tip height on geth based rpc clients.
type BlockSubscriberClient interface {
	// LatestHeight gets the latest height block. An error is handled by the subscriber
	LatestHeight(ctx context.Context) (uint64, error)
}

// ConditionCheck is a function passed in by the caller that checks the block height against
// their condition and returns a bool indicating whether or not the context should be canceled.
type ConditionCheck func(blockHeight uint64) bool

// BlockHeightWatcher creates a subscription to the block height for geth based chains.
// it uses the observer pattern to allow many subscribers.
//
//go:generate go run github.com/vektra/mockery/v2 --name BlockHeightWatcher --output ./mocks --case=underscore
type BlockHeightWatcher interface {
	// Subscribe creates a new block height subscriber.
	Subscribe() <-chan uint64
	// Unsubscribe removes a block height subscriber.
	Unsubscribe(ch <-chan uint64)
}

type blockHeightWatcherImpl struct {
	*BlockBroadcaster
	// ctx is the master context of the block height subscriber. If this context
	// is canceled all child subscriptions will be canceled
	//nolint: containedctx
	ctx context.Context
	// BlockSubscriberClient is a client used to get/subscribe to the block client
	BlockSubscriberClient
}

// HeightCounterMetricName is the name of the height counter metric.
const HeightCounterMetricName = "height_counter"

// NewBlockHeightWatcher creates a new block height subscriber. This creates a channel for getting the latest heights from
// a subscription and attempts to reconnect on disconnect.
func NewBlockHeightWatcher(ctx context.Context, chainID uint64, reader BlockSubscriberClient) BlockHeightWatcher {
	bls := &blockHeightWatcherImpl{
		ctx:                   ctx,
		BlockSubscriberClient: reader,
		BlockBroadcaster:      NewBlockBroadcaster(ctx, chainID),
	}
	go func() {
		for {
			// TODO this needs to have much better retries
			blocks, err := bls.startBlockSubscriber()
			if err != nil {
				// TODO handle fatal case here + backoff
				continue
			}

			for {
				select {
				case block := <-blocks:
					logger.Debugf("got new block %d on chain %d", block, chainID)
					bls.Emit(block)
				case <-ctx.Done():
					logger.Debugf("context ended on chain: %d", chainID)
					return
				}
			}
		}
	}()
	return bls
}

// UpdateHeight updates the block height and sends any heights
// in between last height and new height to the channel.
// as a reminder: BlockBroadcaster is a counter - heights go up, not down.
// UpdateHeight follows that interface.
func (b *BlockBroadcaster) UpdateHeight(newHeight uint64) {
	b.lastHeightMux.Lock()
	// if height == last height we don't need to do anything
	if newHeight > b.lastHeight {
		// on the first set, don't use 0. Use 1 height lower than this
		// so we don't add every height from 0 to new height
		if b.firstSet {
			b.firstSet = false
			b.lastHeight = newHeight - 1
		}

		// iterate the height + any other heights received in the mean time and add to the channel
		for i := b.lastHeight + 1; i <= newHeight; i++ {
			b.lastHeight = i
			// add new heights to the block chan
			b.blockChan <- newHeight
		}
	}
	b.lastHeightMux.Unlock()
}

// PollInterval is how often to poll. This is exported for testing.
var PollInterval = time.Second * 5

// Subscribe subscribes to new block heights. The first height (current height) is sent immediately.
//
//nolint:cyclop
func (b *blockHeightWatcherImpl) startBlockSubscriber() (<-chan uint64, error) {
	acquired := b.subscriberMux.TryAcquire(1)
	if !acquired {
		return nil, errors.New("only one Subscribe() can be used per blockHeightWatcherImpl")
	}

	subscribe := func() {
		ticker := time.NewTicker(PollInterval)
		for {
			select {
			case <-b.ctx.Done():
				return
			case <-ticker.C:
				ctx, cancel := context.WithTimeout(b.ctx, time.Second*30)

				currentHeight, err := b.LatestHeight(ctx)
				if err != nil {
					logger.Warnf("could not create subscription to new blocks: %v on chain %d", err, b.chainID)
					cancel()
					return
				}

				b.UpdateHeight(currentHeight)
				cancel()
			}
		}
	}

	// get the most recent height
	latest, err := b.LatestHeight(b.ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get most recent latest height: %w", err)
	}

	// send the initial latest height to the channel
	b.lastHeight = latest

	go func() {
		defer b.subscriberMux.Release(1)

		select {
		case <-b.ctx.Done():
			return
		case b.blockChan <- b.lastHeight:
			break
		}

		for {
			select {
			case <-b.ctx.Done():
				return
			default:
				// will latest until we have an error and try to restart, unless of course we have another error.
				subscribe()
			}
		}
	}()

	return b.blockChan, err
}
