package guard

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/agents/types"

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

// FindNewestGuardAttestationInInitialState fetches the newest attestation that was suggested from origin.
func (a AttestationCollectorAttestationScanner) FindNewestGuardAttestationInInitialState(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveNewestInProgressAttestationIfInState(
		ctx,
		a.originID,
		a.destinationID,
		types.AttestationStateGuardInitialState)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find newest unsigned and unverified attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job of the scanner
//
//nolint:cyclop
func (a AttestationCollectorAttestationScanner) update(ctx context.Context) error {
	inProgressAttestation, err := a.FindNewestGuardAttestationInInitialState(ctx)
	if err != nil {
		return fmt.Errorf("could not retrieve newest attestation in initial state: %w", err)
	}
	if inProgressAttestation == nil {
		return nil
	}

	nonceToFetch := inProgressAttestation.SignedAttestation().Attestation().Nonce()
	root, err := a.attestationDomain.AttestationCollector().GetRoot(ctx, a.originID, a.destinationID, nonceToFetch)
	if err != nil {
		return fmt.Errorf("error getting root for origin %d, destination %d, nonce %d: %w", a.originID, a.destinationID, nonceToFetch, err)
	}
	if root == [32]byte{} {
		return nil
	}

	signedAttestation, err := a.attestationDomain.AttestationCollector().GetAttestation(ctx, a.originID, a.destinationID, nonceToFetch)
	if err != nil {
		return fmt.Errorf("erroring getting attestation found for origin %d, destination %d, nonce %d: %w", a.originID, a.destinationID, nonceToFetch, err)
	}

	// Either do this here or in the next worker.
	err = a.db.StoreExistingSignedInProgressAttestation(ctx, signedAttestation)
	if err != nil {
		return fmt.Errorf("could not store in-progress attestations: %w", err)
	}
	return nil
}
