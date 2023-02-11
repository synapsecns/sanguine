package indexer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"bitbucket.org/tentontrain/math"

	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/agents/domains"
)

// domainIndexer indexes a single domain and stores event data in the database.
type domainIndexer struct {
	// db contains the new synapsedb
	db db.SynapseDB
	// domain contains the domain clinet
	domain domains.DomainClient
	// interval is the number of seconds
	interval time.Duration
}

// DomainIndexer indexes a domain.
type DomainIndexer interface {
	SyncMessages(ctx context.Context) error
}

// NewDomainIndexer creates a new domain indexer.
//
//nolint:golint,revive
func NewDomainIndexer(db db.SynapseDB, domain domains.DomainClient, interval time.Duration) DomainIndexer {
	return domainIndexer{
		db:       db,
		domain:   domain,
		interval: interval,
	}
}

func (d domainIndexer) SyncMessages(ctx context.Context) error {
	// get the latest indexed height for the dmoain. Note: this can differ based on contract, we'll need to switch this to a per contaact setting
	indexedHeight, err := d.db.GetMessageLatestBlockEnd(ctx, d.domain.Config().DomainID)
	if err != nil && !errors.Is(err, db.ErrNoStoredBlockForChain) {
		return fmt.Errorf("could not get indexed height: %w", err)
	}

	startHeight := math.Max(indexedHeight-d.domain.Config().RequiredConfirmations, d.domain.Config().StartHeight)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(d.interval):
			// TODO: this needs some sort of backoff
			ok, endHeight, err := d.checkAndStoreMessages(ctx, startHeight)
			if err != nil {
				logger.Warn(err)
				continue
			}

			if ok {
				startHeight = endHeight
			}
		}
	}
}

// checkAndStoreMessages is the sync update loop.
func (d domainIndexer) checkAndStoreMessages(ctx context.Context, startHeight uint32) (ok bool, endHeight uint32, err error) {
	tip, err := d.domain.BlockNumber(ctx)
	if err != nil {
		return false, endHeight, fmt.Errorf("could not get latest block number: %w", err)
	}

	if tip <= startHeight {
		return true, startHeight, nil
	}

	// TODO: handle required confs
	sortedMessages, err := d.domain.Origin().FetchSortedMessages(ctx, startHeight, tip)
	if err != nil {
		return false, tip, fmt.Errorf("could not sync updates: %w", err)
	}

	if len(sortedMessages) == 0 {
		err := d.db.StoreMessageLatestBlockEnd(ctx, d.domain.Config().DomainID, tip)
		if err != nil {
			return false, tip, fmt.Errorf("could not store height %d on domain %s: %w", tip, d.domain.Name(), err)
		}

		return true, tip, nil
	}

	for _, message := range sortedMessages {
		err = d.db.StoreCommittedMessage(ctx, d.domain.Config().DomainID, message)
		if err != nil {
			return false, tip, fmt.Errorf("could not get latest message: %w", err)
		}
	}

	// store the tip only after we've stored all the messages
	err = d.db.StoreMessageLatestBlockEnd(ctx, d.domain.Config().DomainID, tip)
	if err != nil {
		return false, tip, fmt.Errorf("could not store height %d on domain %s: %w", tip, d.domain.Name(), err)
	}

	return true, tip, nil
}
