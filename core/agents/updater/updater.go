package updater

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/db/datastore/pebble"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/indexer"
	"golang.org/x/sync/errgroup"
)

// Updater updates the updater.
type Updater struct {
	indexers  map[string]indexer.DomainIndexer
	producers map[string]UpdateProducer
}

// NewUpdater creates a new updater.
func NewUpdater(ctx context.Context, cfg config.Config) (Updater, error) {
	updater := Updater{
		indexers:  make(map[string]indexer.DomainIndexer),
		producers: make(map[string]UpdateProducer),
	}
	for name, domain := range cfg.Domains {
		domainClient, err := evm.NewEVM(ctx, name, domain)
		if err != nil {
			return Updater{}, fmt.Errorf("could not create updater for: %w", err)
		}

		dbHandle, err := pebble.NewMessageDB(cfg.DBPath, name)
		if err != nil {
			return Updater{}, fmt.Errorf("can not create db: %w", err)
		}

		updater.indexers[name] = indexer.NewDomainIndexer(dbHandle, domainClient)
		// updater.producers[name] = NewUpdateProducer(domainClient, dbHandle, )
	}

	return updater, nil
}

// Start starts the updater.{.
func (u Updater) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	for _, domainIndexer := range u.indexers {
		g.Go(func() error {
			return domainIndexer.SyncMessages(ctx)
		})
	}

	return nil
}
