package notary

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/db/datastore/sql"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/indexer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"golang.org/x/sync/errgroup"
	"time"
)

// Notary updates the origin contract.
type Notary struct {
	indexers   map[string]indexer.DomainIndexer
	producers  map[string]AttestationProducer
	submitters map[string]AttestationSubmitter
	signer     signer.Signer
}

// RefreshInterval is how long to wait before refreshing.
//TODO: This should be done in config.
var RefreshInterval = 1 * time.Second

// NewNotary creates a new notary.
func NewNotary(ctx context.Context, cfg config.Config) (_ Notary, err error) {
	notary := Notary{
		indexers:   make(map[string]indexer.DomainIndexer),
		producers:  make(map[string]AttestationProducer),
		submitters: make(map[string]AttestationSubmitter),
	}

	notary.signer, err = config.SignerFromConfig(cfg.Signer)
	if err != nil {
		return Notary{}, fmt.Errorf("could not create notary: %w", err)
	}

	dbType, err := sql.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return Notary{}, fmt.Errorf("could not get legacyDB type: %w", err)
	}

	dbHandle, err := sql.NewStoreFromConfig(ctx, dbType, cfg.Database.ConnString)
	if err != nil {
		return Notary{}, fmt.Errorf("could not connect to legacyDB: %w", err)
	}

	for name, domain := range cfg.Domains {
		domainClient, err := evm.NewEVM(ctx, name, domain)
		if err != nil {
			return Notary{}, fmt.Errorf("could not create notary for: %w", err)
		}

		notary.indexers[name] = indexer.NewDomainIndexer(dbHandle, domainClient, RefreshInterval)
		notary.producers[name] = NewAttestationProducer(domainClient, dbHandle, notary.signer, RefreshInterval)
		// TODO: this needs to be on a separate chain so it'll need to use a different domain client. Config needs to be modified
		notary.submitters[name] = NewAttestationSubmitter(domainClient, dbHandle, notary.signer, RefreshInterval)
	}

	return notary, nil
}

// Start starts the notary.{.
func (u Notary) Start(ctx context.Context) error {
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
		return fmt.Errorf("could not start the notary: %w", err)
	}

	return nil
}
