package updater

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/db/datastore/pebble"
	"github.com/synapsecns/sanguine/core/db/datastore/sql"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/indexer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"golang.org/x/sync/errgroup"
)

// Updater updates the updater.
type Updater struct {
	indexers   map[string]indexer.DomainIndexer
	producers  map[string]UpdateProducer
	submitters map[string]UpdateSubmitter
	signer     signer.Signer
}

var RefreshInterval uint = 4

// NewUpdater creates a new updater.
func NewUpdater(ctx context.Context, cfg config.Config) (_ Updater, err error) {
	updater := Updater{
		indexers:   make(map[string]indexer.DomainIndexer),
		producers:  make(map[string]UpdateProducer),
		submitters: make(map[string]UpdateSubmitter),
	}

	updater.signer, err = config.SignerFromConfig(cfg.Signer)
	if err != nil {
		return Updater{}, fmt.Errorf("could not create updater: %w", err)
	}

	dbType, err := sql.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return Updater{}, fmt.Errorf("could not get db type: %w", err)
	}

	txQueueDB, err := sql.NewStoreFromConfig(ctx, dbType, cfg.Database.ConnString)
	if err != nil {
		return Updater{}, fmt.Errorf("could not connect to db: %w", err)
	}

	for name, domain := range cfg.Domains {
		domainClient, err := evm.NewEVM(ctx, name, domain)
		if err != nil {
			return Updater{}, fmt.Errorf("could not create updater for: %w", err)
		}

		dbHandle, err := pebble.NewMessageDB(cfg.Database.DBPath, name)
		if err != nil {
			return Updater{}, fmt.Errorf("can not create messageDB: %w", err)
		}

		updater.indexers[name] = indexer.NewDomainIndexer(dbHandle, domainClient)
		updater.producers[name] = NewUpdateProducer(domainClient, dbHandle, updater.signer, RefreshInterval)
		updater.submitters[name] = NewUpdateSubmitter(domainClient, dbHandle, txQueueDB, updater.signer, RefreshInterval)
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

	for _, domainProducers := range u.producers {
		g.Go(func() error {
			return domainProducers.Start(ctx)
		})
	}

	for _, domainSubmitters := range u.submitters {
		g.Go(func() error {
			return domainSubmitters.Start(ctx)
		})
	}

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("could not start the updater: %w", err)
	}

	return nil
}
