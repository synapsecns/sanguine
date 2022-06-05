package indexer

import (
	"context"
	"errors"
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/domains"
)

// UpdateProducer updates a producer.
type UpdateProducer struct {
	// domain allows access to the home contract
	domain domains.DomainClient
	// db contains the db object
	db db.DB
	// intervalSeconds adds an interval
	intervalSeconds uint64
}

// FindLatestRoot finds the latest root.
func (u UpdateProducer) FindLatestRoot() (common.Hash, error) {
	latestRoot, err := u.db.RetrieveLatestRoot()
	if err != nil && errors.Is(err, pebble.ErrNotFound) {
		return common.Hash{}, nil
	} else if err != nil {
		return common.Hash{}, fmt.Errorf("could not retrieve latest root: %w", err)
	}

	return latestRoot, nil
}

// Start starts the update producer.
func (u UpdateProducer) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			latestRoot, err := u.FindLatestRoot()
			if err != nil {
				return fmt.Errorf("could not find latest root: %w", err)
			}

			suggestedUpdate, err := u.domain.Home().ProduceUpdate(ctx)
			if err != nil {
				return fmt.Errorf("could not suggest update: %w", err)
			}

			if suggestedUpdate.PreviousRoot() != latestRoot {
				logger.Debugf("Local root not equal to chain root. Skipping update")
				continue
			}

			// Ensure we have not already signed a conflicting update.
			// Ignore suggested if we have.
			existing, err := u.db.RetrieveProducedUpdate(suggestedUpdate.PreviousRoot())
			if err != nil && !errors.Is(err, pebble.ErrNotFound) {
				return fmt.Errorf("could not get update: %w", err)
				// existing was found
			} else if err == nil {
				if existing.Update().NewRoot() != suggestedUpdate.NewRoot() {
					logger.Infof("Updater ignoring conflicting suggested update. Indicates chain awaiting already produced update. Existing update: %s. Suggested conflicting update: %s", existing.Update().NewRoot(), suggestedUpdate.NewRoot())
				}
				continue
			}

			// sign the update
		}
	}
}
