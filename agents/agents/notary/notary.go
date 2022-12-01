package notary

import (
	"context"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"golang.org/x/sync/errgroup"
)

// Notary updates the origin contract.
// TODO: Note right now, I have threads for each origin-destination pair and do no batching at all
// in terms of calls to the origin.
// Right now, for this MVP, this is the simplest path and we can make improvements later
type Notary struct {
	//indexers   map[string]indexer.DomainIndexer
	//producers  map[string]AttestationProducer
	//submitters map[string]AttestationSubmitter
	scanners   map[string]map[string]OriginAttestationScanner
	signers    map[string]map[string]OriginAttestationSigner
	submitters map[string]map[string]OriginAttestationSubmitter
	verifiers  map[string]map[string]OriginAttestationVerifier
	signer     signer.Signer
}

// RefreshInterval is how long to wait before refreshing.
// TODO: This should be done in config.
var RefreshInterval = 1 * time.Second

// NewNotary creates a new notary.
func NewNotary(ctx context.Context, cfg config.Config) (_ Notary, err error) {
	notary := Notary{
		//indexers:   make(map[string]indexer.DomainIndexer),
		//producers:  make(map[string]AttestationProducer),
		//submitters: make(map[string]AttestationSubmitter),
		scanners:   make(map[string]map[string]OriginAttestationScanner),
		signers:    make(map[string]map[string]OriginAttestationSigner),
		submitters: make(map[string]map[string]OriginAttestationSubmitter),
		verifiers:  make(map[string]map[string]OriginAttestationVerifier),
	}

	notary.signer, err = config.SignerFromConfig(cfg.Signer)
	if err != nil {
		return Notary{}, fmt.Errorf("could not create notary: %w", err)
	}

	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
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

		notary.scanners[name] = make(map[string]OriginAttestationScanner)
		notary.signers[name] = make(map[string]OriginAttestationSigner)
		notary.submitters[name] = make(map[string]OriginAttestationSubmitter)
		notary.verifiers[name] = make(map[string]OriginAttestationVerifier)

		//notary.indexers[name] = indexer.NewDomainIndexer(dbHandle, domainClient, RefreshInterval)
		//notary.producers[name] = NewAttestationProducer(domainClient, dbHandle, notary.signer, RefreshInterval)
		// TODO: this needs to be on a separate chain so it'll need to use a different domain client. Config needs to be modified
		//notary.submitters[name] = NewAttestationSubmitter(domainClient, dbHandle, notary.signer, RefreshInterval)
		for destinationName, destinationDomain := range cfg.Domains {
			if domain.DomainID == destinationDomain.DomainID {
				continue
			}
			notary.scanners[name][destinationName] = NewOriginAttestationScanner(domainClient, destinationDomain.DomainID, dbHandle, notary.signer, RefreshInterval)
			notary.signers[name][destinationName] = NewOriginAttestationSigner(domainClient, destinationDomain.DomainID, dbHandle, notary.signer, RefreshInterval)
			notary.submitters[name][destinationName] = NewOriginAttestationSubmitter(domainClient, destinationDomain.DomainID, dbHandle, notary.signer, RefreshInterval)
			notary.verifiers[name][destinationName] = NewOriginAttestationVerifier(domainClient, destinationDomain.DomainID, dbHandle, notary.signer, RefreshInterval)
		}

	}

	return notary, nil
}

// Start starts the notary.{.
func (u Notary) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	for name := range u.scanners {
		for destinationName := range u.scanners[name] {
			name := name // capture func literal
			destinationName := destinationName
			g.Go(func() error {
				//nolint: wrapcheck
				return u.scanners[name][destinationName].Start(ctx)
			})
		}
	}

	for name := range u.signers {
		for destinationName := range u.signers[name] {
			name := name // capture func literal
			destinationName := destinationName
			g.Go(func() error {
				//nolint: wrapcheck
				return u.signers[name][destinationName].Start(ctx)
			})
		}
	}

	for name := range u.submitters {
		for destinationName := range u.submitters[name] {
			name := name // capture func literal
			destinationName := destinationName
			g.Go(func() error {
				//nolint: wrapcheck
				return u.submitters[name][destinationName].Start(ctx)
			})
		}
	}

	for name := range u.verifiers {
		for destinationName := range u.verifiers[name] {
			name := name // capture func literal
			destinationName := destinationName
			g.Go(func() error {
				//nolint: wrapcheck
				return u.verifiers[name][destinationName].Start(ctx)
			})
		}
	}

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("could not start the notary: %w", err)
	}

	return nil
}
