package watcher

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"github.com/synapsecns/sanguine/core/observer"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
	"github.com/teivah/onecontext"
)

// ContractFilterer is the filterer object.
type ContractFilterer interface {
	bind.ContractFilterer
	GetBigChainID() *big.Int
}

// contractWatcherImpl handles listening for logs on a contract and registering listeners/
// producers to listen for different events.
type contractWatcherImpl struct {
	// ctx is the overriding context for the contract contractWatcher
	//nolint: containedctx
	ctx context.Context
	// observer handles registering/deregistering listeners
	// in the futures these should be wrapped to avoid type issues downstream
	// see: https://pkg.go.dev/github.com/eapache/channels#Wrap (and: https://stackoverflow.com/a/25594926/1011803)
	// generic typing in go solves this
	//nolint: staticcheck
	observer *observer.StringObserver
	// producers is used to register producers on the channel
	producers map[common.Address]interface{}
	// producerLock locks the producer map
	producerLock sync.RWMutex
	// eventCount is the total number of events observed on all contracts
	eventCount uint64
	// client is the client used for interacting with the chain
	client ContractFilterer
	// heightWatcher is the block height watcher
	heightWatcher chainwatcher.BlockHeightWatcher
	// requiredConfirmations is how many confirmations we wait before finality
	requiredConfirmations uint
}

// NewContractWatcher creates a new contract contractWatcher. A contract contractWatcher defines two types:
// a listener, and a producer. A producer is responsible for emitting events on a contract to
// listeners. There can only be one producer per contract. A listener processes these events.
// there can be (theoretically) unlimited listeners per producer. Listeners are responsible for
// registering and unregistering themselves.
// TODO: consider replacing with: https://pkg.go.dev/github.com/cryptoriums/telliot@v0.3.0/pkg/tracker/events#New
func NewContractWatcher(ctx context.Context, contractListener ContractFilterer, heightWatcher chainwatcher.BlockHeightWatcher, requiredConfirmations uint) chainwatcher.ContractWatcher {
	return &contractWatcherImpl{
		ctx:                   ctx,
		observer:              observer.NewStringObserver(),
		producers:             make(map[common.Address]interface{}),
		client:                contractListener,
		heightWatcher:         heightWatcher,
		requiredConfirmations: requiredConfirmations,
	}
}

// ListenOnContract listens on a contract. The method here uses a string contract to remain chain agnostic.
func (c *contractWatcherImpl) ListenOnContract(ctx context.Context, contract string, eventLog chan interface{}) error {
	ctx, _ = onecontext.Merge(ctx, c.ctx)

	logger.Debugf("listening on contract %s", contract)

	contractAddress := common.HexToAddress(contract)

	// create the listener
	c.addListener(contractAddress, eventLog)

	// create the producer if it doesn't exist
	err := c.createProducerIfNotExists(ctx, contractAddress)
	if err != nil {
		return err
	}

	return nil
}

// createProducerIfNotExists creates a producer for listening to events on contract address in the contract watcher.
func (c *contractWatcherImpl) createProducerIfNotExists(ctx context.Context, contractAddress common.Address) error {
	// generate a merged context
	ctx, cancel := onecontext.Merge(ctx, c.ctx)

	producedEvents := make(chan types.Log)

	err := c.addProducer(ctx, contractAddress, producedEvents)
	// subscription already exists
	//nolint:nilerr
	if err != nil {
		cancel()
		return nil
	}

	heightSubscription := c.heightWatcher.Subscribe()

	// TODO This belongs in a new create producer function that checks is hasProducer is true first.
	// ideally w/ some kind of retries as well

	go func() {
		// prevent context leaks
		// on every new height
		defer c.heightWatcher.Unsubscribe(heightSubscription)
		for {
			select {
			case <-ctx.Done():
				return
			case height := <-heightSubscription:
				logs, _ := c.filterBlocks(ctx, contractAddress, height, heightSubscription)
				for _, log := range logs {
					select {
					case <-ctx.Done():
						return
					case producedEvents <- log:
						continue
					}
				}
			}
		}
	}()

	return nil
}

// filterBlocks is a block filterer with a backoff
// returns false for didFilter if the filter was not completed. Keeps retrying in case of an error.
func (c *contractWatcherImpl) filterBlocks(ctx context.Context, contractAddress common.Address, receivedHeight uint64, heightSubscription <-chan uint64) (logs []types.Log, didFilter bool) {
	// backoff in the case of an error
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    30 * time.Second,
	}
	// get the new most recent finalized hegiht (chain tip - required confirmations)
	startHeight := receivedHeight - uint64(c.requiredConfirmations)

	// make sure we didn't overflow the uint64
	if startHeight > receivedHeight {
		return []types.Log{}, false
	}

	endHeight := c.getFilterEndHeight(receivedHeight, heightSubscription)

	// timeout is 0 the first time and set by backoff on subsequent requests
	timeout := time.Duration(0)
	for {
		select {
		case <-ctx.Done():
			return []types.Log{}, false
		case <-time.After(timeout):
			var err error

			logs, err = c.client.FilterLogs(ctx, ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(startHeight)),
				ToBlock:   big.NewInt(int64(endHeight)),
				Addresses: []common.Address{contractAddress},
			})
			if err != nil {
				timeout = b.Duration()
				logger.Errorf("got error %v when fetching logs from %d to %d on address %s will retry in %f seconds", err, startHeight, endHeight, contractAddress, timeout.Seconds())
				continue
			}

			return logs, true
		}
	}
}

// getFilterEndHeight gets the latest height from a channel by draining the channel for 1 ms
// due to latency in getLogs/subsequent calls, or other machine induce latency
// producers can fall behind other nodes on occasion. Without a mechanism to depressurize these channels
// these producers can permanently fall behind. This function attempts to drain the channel of all current messages
// so we can get logs for the range. Ideally, this pressure build ups never actually occur.
//
// at the end, as with the start height required confirmations are subtracted so this does not need to be done by the caller.
func (c *contractWatcherImpl) getFilterEndHeight(startHeight uint64, heightChan <-chan uint64) (endHeight uint64) {
	// populate the default
	endHeight = startHeight

OUTER:
	for {
		select {
		case newHeight := <-heightChan:
			endHeight = newHeight
		default:
			break OUTER
		}
	}

	return endHeight - uint64(c.requiredConfirmations)
}

// emit emits an event to a contract channel.
func (c *contractWatcherImpl) emit(address common.Address, event types.Log) {
	c.observer.Emit(address.String(), event)
}

// addListener adds a listener to the observer for contract address to receive events
// on ch channel. wraps observer.addListener() to prevent exposing Emit() by fully
// extending observer
// TODO: this should not be an interface type.
func (c *contractWatcherImpl) addListener(address common.Address, ch chan interface{}) {
	c.observer.AddListener(address.String(), ch)
}

// RemoveListener removes a listener from the observer for the contract address
// on a ch channel. wraps observer.RemoveListener to prevent exposing Emit() by fully
// extending observer
// TODO: this should not be an interface type.
func (c *contractWatcherImpl) RemoveListener(address common.Address, ch chan interface{}) {
	c.observer.RemoveListener(address.String(), ch)
}

var errProducerAlreadyExists = errors.New("producer already exists")

// addProducer attempts to add a new producer to the contract contractWatcher. Note, since only
// one producer can be registered per address an error will be returned if the producer doe snot
// exist
// Any cancellations on the contractWatcherImpl context or the passed in context will cause the producer to cancel and be removed
// a message will be sent to errorChan on complete.
func (c *contractWatcherImpl) addProducer(ctx context.Context, address common.Address, producedEvents <-chan types.Log) error {
	c.producerLock.Lock()
	defer c.producerLock.Unlock()

	if !c.hasProducer(address) {
		logger.Debugf("adding contract watcher producer for %s", address)

		c.producers[address] = struct{}{}
		go func() {
			// create the combined context of the producer and the contract contractWatcher
			ctx, cancel := onecontext.Merge(ctx, c.ctx)
			defer cancel()

			for {
				select {
				// add an error to the error chan if the context is done
				case <-ctx.Done():
					logger.Warnf("contract watcher for contract %s finished", address.String())
					return
				// Emit any produced events to all listeners
				case event := <-producedEvents:
					atomic.AddUint64(&c.eventCount, 1)
					logger.Warnf("got event in tx %s for contract %s at height %d", event.TxHash, event.Address, event.BlockNumber)
					c.emit(address, event)
				}
			}
		}()
		return nil
	}
	return fmt.Errorf("could not add producer %s, %w", address.String(), errProducerAlreadyExists)
}

// hasProducer returns whether or not the contract producer has a producer already
// any caller to this method should call the lock before calling this method.
func (c *contractWatcherImpl) hasProducer(address common.Address) bool {
	if _, ok := c.producers[address]; ok {
		return true
	}
	return false
}
