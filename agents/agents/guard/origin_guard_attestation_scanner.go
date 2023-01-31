package guard

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// OriginGuardAttestationScanner fetches merkle roots for particular origin-destination pair.
// TODO: this needs to become an interface.
type OriginGuardAttestationScanner struct {
	// originDomain allows access to the origin contract on the origin chain
	originDomain domains.DomainClient
	// attestationDomain allows access to the attestation contract on the SYN chain
	attestationDomain domains.DomainClient
	// destinationDomain allows access to the destination contract on the destination chain
	destinationDomain domains.DomainClient
	// db is the synapse db
	db db.SynapseDB
	// bondedSigner is the attestation signer that must be a bonded agent
	bondedSigner signer.Signer
	// unbondedSigner is the signer for submitting transactions
	unbondedSigner signer.Signer
	// interval waits for an interval
	interval time.Duration
}

// NewOriginGuardAttestationScanner creates a new origin guard attestation scanner.
func NewOriginGuardAttestationScanner(
	originDomain domains.DomainClient,
	attestationDomain domains.DomainClient,
	destinationDomain domains.DomainClient,
	db db.SynapseDB,
	bondedSigner signer.Signer,
	unbondedSigner signer.Signer,
	interval time.Duration) OriginGuardAttestationScanner {
	return OriginGuardAttestationScanner{
		originDomain:      originDomain,
		attestationDomain: attestationDomain,
		destinationDomain: destinationDomain,
		db:                db,
		bondedSigner:      bondedSigner,
		unbondedSigner:    unbondedSigner,
		interval:          interval,
	}
}

// Start starts the OriginGuardAttestationScanner.
func (a OriginGuardAttestationScanner) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logger.Info("Guard OriginGuardAttestationScanner exiting without error")
			return nil
		case <-time.After(a.interval):
			err := a.update(ctx)
			if err != nil {
				logger.Errorf("Guard OriginGuardAttestationScanner exiting with error: %v", err)
				return err
			}
		}
	}
}

// FindLatestNonce fetches the latest nonce for a given chain.
func (a OriginGuardAttestationScanner) FindLatestNonce(ctx context.Context) (nonce uint32, err error) {
	latestNonce, err := a.db.RetrieveLatestCachedNonce(ctx, a.originDomain.Config().DomainID, a.destinationDomain.Config().DomainID)
	if err != nil {
		if errors.Is(err, db.ErrNoNonceForDomain) {
			return 0, nil
		}
		return 0, fmt.Errorf("could not find latest root: %w", err)
	}
	return latestNonce, nil
}

// update runs the job of the scanner
//
//nolint:cyclop
func (a OriginGuardAttestationScanner) update(ctx context.Context) error {
	latestNonce, err := a.FindLatestNonce(ctx)
	if err != nil {
		return fmt.Errorf("could not find latest root: %w", err)
	}

	attestation, err := a.originDomain.Origin().SuggestAttestation(ctx, a.destinationDomain.Config().DomainID)
	if errors.Is(err, domains.ErrNoUpdate) {
		// no update produced this time
		return nil
	}

	if err != nil {
		return fmt.Errorf("could not get suggested attestation: %w", err)
	}

	if latestNonce > uint32(0) && attestation.Nonce() <= latestNonce {
		// We already have seen this nonce
		return nil
	}

	err = a.db.StoreNewGuardInProgressAttestation(ctx, attestation)
	if err != nil {
		return fmt.Errorf("could not store in-progress attestations: %w", err)
	}

	return nil
}
