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

// Notary in the current version scans the origins for new messages, signs them, and posts to attestation collector.
// TODO: Note right now, I have threads for each origin-destination pair and do no batching at all
// in terms of calls to the origin.
// Right now, for this MVP, this is the simplest path and we can make improvements later.
type Notary struct {
	scanners        map[string]OriginAttestationScanner
	signers         map[string]OriginAttestationSigner
	submitters      map[string]OriginAttestationSubmitter
	verifiers       map[string]OriginAttestationVerifier
	bondedSigner    signer.Signer
	unbondedSigner  signer.Signer
	refreshInterval time.Duration
}

// NewNotary creates a new notary.
//
//nolint:cyclop
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
	}

	notary.bondedSigner, err = config.SignerFromConfig(ctx, cfg.BondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error with bondedSigner, could not create notary: %w", err)
	}

	notary.unbondedSigner, err = config.SignerFromConfig(ctx, cfg.UnbondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error with unbondedSigner, could not create notary: %w", err)
	}

	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return Notary{}, fmt.Errorf("could not get legacyDB type: %w", err)
	}

	dbHandle, err := sql.NewStoreFromConfig(ctx, dbType, cfg.Database.ConnString, cfg.DBPrefix)
	if err != nil {
		return Notary{}, fmt.Errorf("could not connect to legacyDB: %w", err)
	}

	destinationClient, err := evm.NewEVM(ctx, "destination_client", cfg.DestinationDomain)
	if err != nil {
		return Notary{}, fmt.Errorf("error with destinationClient, could not create notary for: %w", err)
	}
	attestationClient, err := evm.NewEVM(ctx, "attestation_client", cfg.AttestationDomain)
	if err != nil {
		return Notary{}, fmt.Errorf("error with attestationClient, could not create notary for: %w", err)
	}

	err = attestationClient.AttestationCollector().PrimeNonce(ctx, notary.unbondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error trying to PrimeNonce for attestationClient, could not create notary for: %w", err)
	}

	err = destinationClient.Destination().PrimeNonce(ctx, notary.unbondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error trying to PrimeNonce for destinationClient, could not create notary for: %w", err)
	}

	for name, originDomain := range cfg.OriginDomains {
		if originDomain.DomainID == cfg.DestinationDomain.DomainID {
			continue
		}

		originClient, err := evm.NewEVM(ctx, name, originDomain)
		if err != nil {
			return Notary{}, fmt.Errorf("error with originClient, could not create notary for: %w", err)
		}

		notary.scanners[name] = NewOriginAttestationScanner(originClient, attestationClient, destinationClient, dbHandle, notary.bondedSigner, notary.unbondedSigner, notary.refreshInterval)
		notary.signers[name] = NewOriginAttestationSigner(originClient, attestationClient, destinationClient, dbHandle, notary.bondedSigner, notary.unbondedSigner, notary.refreshInterval)
		notary.submitters[name] = NewOriginAttestationSubmitter(originClient, attestationClient, destinationClient, dbHandle, notary.bondedSigner, notary.unbondedSigner, notary.refreshInterval)
		notary.verifiers[name] = NewOriginAttestationVerifier(originClient, attestationClient, destinationClient, dbHandle, notary.bondedSigner, notary.unbondedSigner, notary.refreshInterval)
	}

	return notary, nil
}

// Start starts the notary.
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
		logger.Errorf("Notary exiting with error: %v", err)
		return fmt.Errorf("could not start the notary: %w", err)
	}

	logger.Info("Notary exiting without error")
	return nil
}
