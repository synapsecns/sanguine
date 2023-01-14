package guard

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// AttestationDoubleCheckOnOriginVerifier double checks attestations on origin.
// TODO: this needs to become an interface.
type AttestationDoubleCheckOnOriginVerifier struct {
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

// NewAttestationDoubleCheckOnOriginVerifier creates a new attestation double-check-on-origin verifier.
func NewAttestationDoubleCheckOnOriginVerifier(
	originDomain domains.DomainClient,
	attestationDomain domains.DomainClient,
	destinationDomain domains.DomainClient,
	db db.SynapseDB,
	bondedSigner signer.Signer,
	unbondedSigner signer.Signer,
	interval time.Duration) AttestationDoubleCheckOnOriginVerifier {
	return AttestationDoubleCheckOnOriginVerifier{
		originDomain:      originDomain,
		attestationDomain: attestationDomain,
		destinationDomain: destinationDomain,
		db:                db,
		bondedSigner:      bondedSigner,
		unbondedSigner:    unbondedSigner,
		interval:          interval,
	}
}

// Start starts the OriginAttestationScanner.
func (a AttestationDoubleCheckOnOriginVerifier) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(a.interval):
			err := a.update(ctx)
			if err != nil {
				return err
			}
		}
	}
}

// FindOldestGuardUnsignedAndUnverifiedAttestation fetches the oldest attestation that still needs to be signed by the guard and needs to be verified.
func (a AttestationDoubleCheckOnOriginVerifier) FindOldestGuardUnsignedAndUnverifiedAttestation(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveOldestGuardUnsignedAndUnverifiedInProgressAttestation(ctx, a.originDomain.Config().DomainID, a.destinationDomain.Config().DomainID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find oldest unsigned and unverified attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job of the verifier
//
//nolint:cyclop
func (a AttestationDoubleCheckOnOriginVerifier) update(ctx context.Context) error {
	inProgressAttestationToVerify, err := a.FindOldestGuardUnsignedAndUnverifiedAttestation(ctx)
	if err != nil {
		return fmt.Errorf("could not retrieve oldest unsigned and unverified attestation: %w", err)
	}
	if inProgressAttestationToVerify == nil {
		return nil
	}

	// TODO (joe): Currently we are scanning all nonces in order. Later, we really want to get the latest
	// attestation after the latestNonce if any exists.
	attestationFromOrigin, dispatchBlockNumber, err := a.originDomain.Origin().GetHistoricalAttestation(ctx, a.destinationDomain.Config().DomainID, inProgressAttestationToVerify.SignedAttestation().Attestation().Nonce())
	if errors.Is(err, domains.ErrNoUpdate) {
		return fmt.Errorf("FRAUD ALERT! could not get historical attestation with origin %d, destination %d, nonce %d : %w",
			inProgressAttestationToVerify.SignedAttestation().Attestation().Origin(),
			inProgressAttestationToVerify.SignedAttestation().Attestation().Destination(),
			inProgressAttestationToVerify.SignedAttestation().Attestation().Nonce(),
			err)
	}
	if err != nil {
		return fmt.Errorf("could not get historical attestation: %w", err)
	}

	if !a.isConfirmed(dispatchBlockNumber) {
		// not yet confirmed so skip
		return nil
	}

	if attestationFromOrigin.Root() != inProgressAttestationToVerify.SignedAttestation().Attestation().Root() {
		return fmt.Errorf("FRAUD ALERT! Roots do not match for origin %d, destination %d, nonce %d : %w",
			inProgressAttestationToVerify.SignedAttestation().Attestation().Origin(),
			inProgressAttestationToVerify.SignedAttestation().Attestation().Destination(),
			inProgressAttestationToVerify.SignedAttestation().Attestation().Nonce(),
			err)
	}

	nowTime := time.Now()
	inProgressAttestationToMarkVerified := types.NewInProgressAttestation(
		inProgressAttestationToVerify.SignedAttestation(),
		inProgressAttestationToVerify.OriginDispatchBlockNumber(),
		&nowTime,
		0)
	err = a.db.MarkVerifiedOnOrigin(ctx, inProgressAttestationToMarkVerified)
	if err != nil {
		return fmt.Errorf("could not store submission time for attestation: %w", err)
	}

	return nil
}

func (a AttestationDoubleCheckOnOriginVerifier) isConfirmed(txBlockNum uint64) bool {
	// TODO (joe): figure this out
	return true
}
