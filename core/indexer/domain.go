package indexer

import (
	"context"
	"errors"
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/domains"
)

// domainIndexer indexes a single domain and stores event data in the database.
type domainIndexer struct {
	// db contains the db
	db db.MessageDB
	// domain contains the domain clinet
	domain domains.DomainClient
}

// DomainIndexer indexes a domain.
type DomainIndexer interface {
	SyncMessages(ctx context.Context) error
}

// NewDomainIndexer creates a new domain indexer.
func NewDomainIndexer(db db.MessageDB, domain domains.DomainClient) domainIndexer {
	return domainIndexer{
		db:     db,
		domain: domain,
	}
}

func (d domainIndexer) SyncMessages(ctx context.Context) error {
	// get the latest indexed height for the dmoain. Note: this can differ based on contract, we'll need to switch this to a per contaact setting
	indexedHeight, err := d.db.GetMessageLatestBlockEnd()
	if err != nil && !errors.Is(err, pebble.ErrNotFound) {
		return fmt.Errorf("could not get indexed height: %w", err)
	}

	startHeight := maxUint32(indexedHeight-d.domain.Config().RequiredConfirmations, d.domain.Config().StartHeight)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
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
	sortedMessages, err := d.domain.Home().FetchSortedMessages(ctx, startHeight, tip)
	if err != nil {
		return false, tip, fmt.Errorf("could not sync updates: %w", err)
	}

	if len(sortedMessages) == 0 {
		err := d.db.StoreMessageLatestBlockEnd(tip)
		if err != nil {
			return false, tip, fmt.Errorf("could not store height %d on domain %s: %w", tip, d.domain.Name(), err)
		}

		return true, tip, nil
	}

	for _, message := range sortedMessages {
		err = d.db.StoreLatestMessage(message)
		if err != nil {
			return false, tip, fmt.Errorf("could not get latest message: %w", err)
		}
	}

	// store the tip only after we've stored all the messages
	err = d.db.StoreMessageLatestBlockEnd(tip)
	if err != nil {
		return false, tip, fmt.Errorf("could not store height %d on domain %s: %w", tip, d.domain.Name(), err)
	}

	return true, tip, nil
}

// maxUint32 gets the maximum uint32 value out of two
// TODO: once we upgrade to go 1.18 (currently waiting on golangci-lint), we can use https://bitbucket.org/tentontrain/math/src/master/compare_test.go
func maxUint32(x, y uint32) uint32 {
	if x > y {
		return x
	}
	return y
}
