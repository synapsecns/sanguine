package watcher

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
	"testing"
	"time"
)

// defaultPollInterval gets the original poll interval so it can ber reset after the test.
var defaultPollInterval time.Duration

func init() {
	defaultPollInterval = chainwatcher.PollInterval
}

// GetDefaultPollInterval gets the default poll interval so it can be set for tests.
func GetDefaultPollInterval() time.Duration {
	return defaultPollInterval
}

// TestContractWatcher is a wrapped contract watcher with exported methods for testing.
type TestContractWatcher interface {
	chainwatcher.ContractWatcher
	AddListener(address common.Address, ch chan interface{})
	AddProducer(ctx context.Context, address common.Address, producedEvents <-chan types.Log) error
}

// NewTestContractWatcher exports contractWatcherImpl methods for testing.
func NewTestContractWatcher(ctx context.Context, tb testing.TB, contractListener ContractFilterer, heightWatcher chainwatcher.BlockHeightWatcher, requiredConfirmations uint) TestContractWatcher {
	tb.Helper()
	contractWatcher := NewContractWatcher(ctx, contractListener, heightWatcher, requiredConfirmations)
	castWatcher, ok := contractWatcher.(*contractWatcherImpl)
	True(tb, ok)
	return castWatcher
}

// AddListener adds a listener for testing.
func (c *contractWatcherImpl) AddListener(address common.Address, ch chan interface{}) {
	c.addListener(address, ch)
}

// AddProducer exports addProducer for testing.
func (c *contractWatcherImpl) AddProducer(ctx context.Context, address common.Address, producedEvents <-chan types.Log) error {
	return c.addProducer(ctx, address, producedEvents)
}
