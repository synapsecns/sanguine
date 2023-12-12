package listener

import (
	"context"
	"fmt"

	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db"

	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	lru "github.com/hashicorp/golang-lru"
	"github.com/synapsecns/sanguine/ethergo/client"
	relayerTypes "github.com/synapsecns/sanguine/rfq/rfq-relayer/utils"
)

const (
	// DefaultCacheSize is the default size of the LRU cache.
	DefaultCacheSize = 1000
)

// IChainListener is the interface for a ChainListener.
type IChainListener interface {
	ABI() abi.ABI
	ChainID() uint32
	StartListening(ctx context.Context) error
	IterateThroughLogs(logs []types.Log, lastUnconfirmedBlock uint64) error
}

type chainListenerImpl struct {
	db               db.DB
	config           *ChainListenerConfig
	eventChan        chan relayerTypes.WrappedLog
	seenChan         chan relayerTypes.WrappedLog
	unconfirmedCache *lru.Cache
	confirmedCache   *lru.Cache
}

// ChainListenerConfig is a config struct for a ChainListener.
type ChainListenerConfig struct {
	ChainID         uint32
	StartBlock      uint64
	BridgeAddress   common.Address
	Client          client.EVM
	PollInterval    int
	MaxGetLogsRange uint64
	Confirmations   uint64
	ABI             abi.ABI
}

func NewChainListener(config *ChainListenerConfig, db db.DB, eventChan chan relayerTypes.WrappedLog, seenChan chan relayerTypes.WrappedLog) (IChainListener, error) {
	// Create caches
	unconfirmedCache, err := lru.New(DefaultCacheSize)
	if err != nil {
		return nil, fmt.Errorf("could not initialize cache: %w", err)
	}
	confirmedCache, err := lru.New(DefaultCacheSize)
	if err != nil {
		return nil, fmt.Errorf("could not initialize cache: %w", err)
	}

	return &chainListenerImpl{
		db:               db,
		config:           config,
		eventChan:        eventChan,
		seenChan:         seenChan,
		unconfirmedCache: unconfirmedCache,
		confirmedCache:   confirmedCache,
	}, nil
}

func (c *chainListenerImpl) ABI() abi.ABI {
	return c.config.ABI
}
func (c *chainListenerImpl) ChainID() uint32 {
	return c.config.ChainID
}

func (c *chainListenerImpl) StartListening(ctx context.Context) error {
	// Handle initial starting block
	startBlock := c.config.StartBlock
	lastIndexed, err := c.db.GetLastIndexed(ctx, c.config.ChainID, c.config.BridgeAddress)
	if err != nil {
		return fmt.Errorf("could not get last indexed block: %w", err)
	}
	if lastIndexed > startBlock {
		startBlock = lastIndexed
	}

	// Backoff config
	b := newBackoffConfig()

	// Client for cleanliness
	rpcClient := c.config.Client

	// Init poll interval
	pollInterval := time.Duration(0)

	// Start listening for events
	for {
		select {
		case <-ctx.Done(): // context cancellation safe
			return fmt.Errorf("context was canceled")
		case <-time.After(pollInterval * time.Second):
			// Get latest block
			latestBlock, lErr := rpcClient.BlockNumber(ctx)
			if lErr != nil {
				pollInterval = b.Duration()
				logger.Errorf("error getting latest block: %v", lErr)
				continue
			}

			// Check if latest block is the same as start block (for chains with slow block times)
			if latestBlock == startBlock {
				pollInterval = time.Duration(c.config.PollInterval)
				logger.Infof("latest block is the same as start block, sleeping for %v", pollInterval)
				continue
			}

			// Handle if the listener is more than one get logs range behind the head
			endBlock := latestBlock
			lastUnconfirmedBlock := latestBlock - c.config.Confirmations
			if startBlock+c.config.MaxGetLogsRange < latestBlock {
				endBlock = startBlock + c.config.MaxGetLogsRange
				// This will be used as the bottom of the range in the next iteration
				lastUnconfirmedBlock = endBlock
			}

			// Get all logs in the range
			filterQuery := c.buildFilterQuery(big.NewInt(int64(startBlock)), big.NewInt(int64(endBlock)))
			logs, lErr := rpcClient.FilterLogs(ctx, filterQuery)
			if lErr != nil {
				pollInterval = b.Duration()
				logger.Errorf("error filtering logs: %v", lErr)
				continue
			}

			// Iterate through logs
			lErr = c.IterateThroughLogs(logs, lastUnconfirmedBlock)
			if lErr != nil {
				pollInterval = b.Duration()
				logger.Errorf("error iterating through logs: %v", lErr)
				continue
			}

			// Update start block
			startBlock = lastUnconfirmedBlock
			pollInterval = time.Duration(c.config.PollInterval)
		}

	}

}

func (c *chainListenerImpl) buildFilterQuery(fromBlock *big.Int, toBlock *big.Int) ethereum.FilterQuery {
	return ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{c.config.BridgeAddress},
		Topics:    [][]common.Hash{{c.config.ABI.Events["BridgeRequested"].ID, c.config.ABI.Events["BridgeRelayed"].ID}},
	}
}

// IterateThroughLogs iterates through logs and sends them to their respective channels.
func (c *chainListenerImpl) IterateThroughLogs(logs []types.Log, lastUnconfirmedBlock uint64) error {
	for _, log := range logs {
		// Handle if in unconfirmed range
		if log.BlockNumber >= lastUnconfirmedBlock {
			// Check if in cache
			if _, ok := c.unconfirmedCache.Get(log.TxHash.Hex()); ok {
				continue
			}
			// Add to cache
			c.unconfirmedCache.Add(log.TxHash.Hex(), log.BlockNumber)
			// Add to channel
			c.seenChan <- relayerTypes.WrappedLog{Log: log, OriginChainID: c.config.ChainID}
		} else { // Handle if in confirmed range
			// Check if in cache
			if _, ok := c.confirmedCache.Get(log.TxHash.Hex()); ok {
				continue
			}
			// Add to cache
			c.confirmedCache.Add(log.TxHash.Hex(), log.BlockNumber)
			// Add to channel
			c.eventChan <- relayerTypes.WrappedLog{Log: log, OriginChainID: c.config.ChainID}
		}
	}
	return nil
}

func newBackoffConfig() *backoff.Backoff {
	return &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    10 * time.Millisecond,
		Max:    1 * time.Second,
	}

}
