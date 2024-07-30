package inventory

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"sync"
	"time"
)

// inFlightManager stores in-flight quotes and allows retrieval via the db.
// it is thread-safe.
type inFlightManager struct {
	ttl   time.Duration
	db    reldb.Service
	mux   sync.RWMutex
	entry *inFlightQuoteCacheEntry
}

// inFlightQuoteCacheEntry represents an entry in the in-flight quote cache.
type inFlightQuoteCacheEntry struct {
	createdAt time.Time
	quotes    []reldb.QuoteRequest
}

// QuoterOption defines a type for functional options.
type QuoterOption func(*inFlightManager)

// WithTTL sets the TTL for the inFlightManager.
func WithTTL(ttl time.Duration) QuoterOption {
	return func(q *inFlightManager) {
		q.ttl = ttl
	}
}

const defaultTTL = 250 * time.Millisecond

// newInflightManager creates a new inFlightManager with the given options.
func newInflightManager(options ...QuoterOption) *inFlightManager {
	// Default TTL to 250ms
	quoter := &inFlightManager{
		ttl: defaultTTL,
	}

	// Apply options
	for _, opt := range options {
		opt(quoter)
	}

	return quoter
}

func (q *inFlightManager) GetInFlightQuotes(ctx context.Context, skipCache bool) (quotes []reldb.QuoteRequest, err error) {
	if skipCache || q.shouldRefresh() {
		inFlightQuotes, err := q.db.GetQuoteResultsByStatus(ctx, reldb.CommittedPending, reldb.CommittedConfirmed, reldb.RelayStarted)
		if err != nil {
			return nil, fmt.Errorf("could not get in flight quotes: %w", err)
		}
		q.mux.Lock()
		q.entry = &inFlightQuoteCacheEntry{
			createdAt: time.Now(),
			quotes:    inFlightQuotes,
		}
		q.mux.Unlock()
	}

	return q.entry.quotes, nil
}

func (q *inFlightManager) shouldRefresh() bool {
	q.mux.RLock()
	defer q.mux.RUnlock()

	return q.entry == nil || time.Since(q.entry.createdAt) > q.ttl
}
