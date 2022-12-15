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
// Right now, for this MVP, this is the simplest path and we can make improvements later.
type Notary struct {
	scanners        map[string]OriginAttestationScanner
	signers         map[string]OriginAttestationSigner
	submitters      map[string]OriginAttestationSubmitter
	verifiers       map[string]OriginAttestationVerifier
	signer          signer.Signer
	destinationID   uint32
	refreshInterval time.Duration
}

// NewNotary creates a new notary.
func NewNotary(ctx context.Context, cfg config.NotaryConfig) (_ Notary, err error) {
	if cfg.RefreshIntervalInSeconds == int64(0) {
		return Notary{}, fmt.Errorf("cfg.refreshInterval cannot be 0")
	}
	notary := Notary{
		scanners:        make(map[string]OriginAttestationScanner),
		signers:         make(map[string]OriginAttestationSigner),
		submitters:      make(map[string]OriginAttestationSubmitter),
		verifiers:       make(map[string]OriginAttestationVerifier),
		refreshInterval: time.Second * time.Duration(cfg.RefreshIntervalInSeconds),
		destinationID:   cfg.DestinationID,
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

		notary.scanners[name] = NewOriginAttestationScanner(domainClient, notary.destinationID, dbHandle, notary.signer, notary.refreshInterval)
		notary.signers[name] = NewOriginAttestationSigner(domainClient, notary.destinationID, dbHandle, notary.signer, notary.refreshInterval)
		notary.submitters[name] = NewOriginAttestationSubmitter(domainClient, notary.destinationID, dbHandle, notary.signer, notary.refreshInterval)
		notary.verifiers[name] = NewOriginAttestationVerifier(domainClient, notary.destinationID, dbHandle, notary.signer, notary.refreshInterval)
	}

	return notary, nil
}

// Start starts the notary.{.
func (u Notary) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	for name := range u.scanners {
		name := name // capture func literal
		g.Go(func() error {
			//nolint: wrapcheck
			return u.scanners[name].Start(ctx)
		})
	}

	for name := range u.signers {
		name := name // capture func literal
		g.Go(func() error {
			//nolint: wrapcheck
			return u.signers[name].Start(ctx)
		})
	}

	for name := range u.submitters {
		name := name // capture func literal
		g.Go(func() error {
			//nolint: wrapcheck
			return u.submitters[name].Start(ctx)
		})
	}

	for name := range u.verifiers {
		name := name // capture func literal
		g.Go(func() error {
			//nolint: wrapcheck
			return u.verifiers[name].Start(ctx)
		})
	}

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("could not start the notary: %w", err)
	}

	return nil
}
