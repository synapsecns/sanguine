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
	"time"
)

// Updater updates the updater.
type Updater struct {
	indexers   map[string]indexer.DomainIndexer
	producers  map[string]UpdateProducer
	submitters map[string]AttestationSubmitter
	signer     signer.Signer
}

// RefreshInterval is how long to wait before refreshing.
//TODO: This should be done in config.
var RefreshInterval = 1 * time.Second

// NewUpdater creates a new updater.
func NewUpdater(ctx context.Context, cfg config.Config) (_ Updater, err error) {
	updater := Updater{
		indexers:   make(map[string]indexer.DomainIndexer),
		producers:  make(map[string]UpdateProducer),
		submitters: make(map[string]AttestationSubmitter),
	}

	updater.signer, err = config.SignerFromConfig(cfg.Signer)
	if err != nil {
		return Updater{}, fmt.Errorf("could not create updater: %w", err)
	}

	dbType, err := sql.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return Updater{}, fmt.Errorf("could not get legacyDB type: %w", err)
	}

	dbHandle, err := sql.NewStoreFromConfig(ctx, dbType, cfg.Database.ConnString)
	if err != nil {
		return Updater{}, fmt.Errorf("could not connect to legacyDB: %w", err)
	}

	for name, domain := range cfg.Domains {
		domainClient, err := evm.NewEVM(ctx, name, domain)
		if err != nil {
			return Updater{}, fmt.Errorf("could not create updater for: %w", err)
		}

		legacyDBHandle, err := pebble.NewMessageDB(cfg.Database.DBPath, name)
		if err != nil {
			return Updater{}, fmt.Errorf("can not create messageDB: %w", err)
		}

		updater.indexers[name] = indexer.NewDomainIndexer(dbHandle, domainClient, RefreshInterval)
		updater.producers[name] = NewAttestationProducer(domainClient, dbHandle, updater.signer, RefreshInterval)
		updater.submitters[name] = NewAttestationSubmitter(domainClient, legacyDBHandle, dbHandle, updater.signer, RefreshInterval)
	}

	return updater, nil
}

// Start starts the updater.{.
func (u Updater) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	for i := range u.indexers {
		i := i // capture func literal
		g.Go(func() error {
			//nolint: wrapcheck
			return u.indexers[i].SyncMessages(ctx)
		})
	}

	for i := range u.producers {
		i := i // capture func literal
		g.Go(func() error {
			//nolint: wrapcheck
			return u.producers[i].Start(ctx)
		})
	}

	for i := range u.submitters {
		i := i // capture func literal
		g.Go(func() error {
			//nolint: wrapcheck
			return u.submitters[i].Start(ctx)
		})
	}

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("could not start the updater: %w", err)
	}

	return nil
}
