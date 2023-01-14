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

// AttestationCollectorAttestationScanner fetches attestations for particular origin-destination pair.
// TODO: this needs to become an interface.
type AttestationCollectorAttestationScanner struct {
	// attestationDomain allows access to the attestation contract
	attestationDomain domains.DomainClient
	// domain of origin destination
	originID uint32
	// domain of target destination
	destinationID uint32
	// db is the synapse db
	db db.SynapseDB
	// signer is the signer
	signer signer.Signer
	// interval waits for an interval
	interval time.Duration
}

// NewAttestationCollectorAttestationScanner creates a new attestation collector attestation scanner.
func NewAttestationCollectorAttestationScanner(domain domains.DomainClient, originID, destinationID uint32, db db.SynapseDB, signer signer.Signer, interval time.Duration) AttestationCollectorAttestationScanner {
	return AttestationCollectorAttestationScanner{
		attestationDomain: domain,
		originID:          originID,
		destinationID:     destinationID,
		db:                db,
		signer:            signer,
		interval:          interval,
	}
}

// Start starts the AttestationCollectorAttestationScanner.
func (a AttestationCollectorAttestationScanner) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logger.Info("Guard AttestationCollectorAttestationScanner exiting without error")
			return nil
		case <-time.After(a.interval):
			err := a.update(ctx)
			if err != nil {
				logger.Errorf("Guard AttestationCollectorAttestationScanner exiting with error: %v", err)
				return err
			}
		}
	}
}

// FindLatestNonce fetches the latest cached nonce for a given chain.
// TODO (joe): there was a bug in this code not covered by current tests.
// Make sure to add a test that covers when latestNonce is not zero.
func (a AttestationCollectorAttestationScanner) FindLatestNonce(ctx context.Context) (nonce uint32, err error) {
	latestNonce, err := a.db.RetrieveLatestCachedNonce(ctx, a.originID, a.destinationID)
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
func (a AttestationCollectorAttestationScanner) update(ctx context.Context) error {
	latestNonce, err := a.FindLatestNonce(ctx)
	if err != nil {
		return fmt.Errorf("could not find latest root: %w", err)
	}

	// TODO (joe): Currently we are scanning all nonces in order. Later, we really want to get the latest
	// attestation after the latestNonce if any exists.
	nextNonce := latestNonce + 1
	root, err := a.attestationDomain.AttestationCollector().GetRoot(ctx, a.originID, a.destinationID, nextNonce)
	if err != nil {
		return fmt.Errorf("error getting root for origin %d, destination %d, nonce %d: %w", a.originID, a.destinationID, nextNonce, err)
	}
	if root == [32]byte{} {
		return nil
	}

	signedAttestation, err := a.attestationDomain.AttestationCollector().GetAttestation(ctx, a.originID, a.destinationID, nextNonce)
	if err != nil {
		return fmt.Errorf("erroring getting attestation found for origin %d, destination %d, nonce %d: %w", a.originID, a.destinationID, nextNonce, err)
	}

	// TODO (joe): Check if attestation is valid on Origin before saving it to the DB.
	// Either do this here or in the next worker.
	err = a.db.StoreExistingSignedInProgressAttestation(ctx, signedAttestation)
	if err != nil {
		return fmt.Errorf("could not store in-progress attestations: %w", err)
	}
	return nil
}
