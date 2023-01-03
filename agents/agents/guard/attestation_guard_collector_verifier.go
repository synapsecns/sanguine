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

// AttestationGuardCollectorVerifier verifies the signed attestation (by both notary and guard)
// is posted successfully on the Attestation Collector.
// TODO: this needs to become an interface.
type AttestationGuardCollectorVerifier struct {
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

// NewAttestationGuardCollectorVerifier creates a new attestation guard collector verifier.
func NewAttestationGuardCollectorVerifier(
	originDomain domains.DomainClient,
	attestationDomain domains.DomainClient,
	destinationDomain domains.DomainClient,
	db db.SynapseDB,
	bondedSigner signer.Signer,
	unbondedSigner signer.Signer,
	interval time.Duration) AttestationGuardCollectorVerifier {
	return AttestationGuardCollectorVerifier{
		originDomain:      originDomain,
		attestationDomain: attestationDomain,
		destinationDomain: destinationDomain,
		db:                db,
		bondedSigner:      bondedSigner,
		unbondedSigner:    unbondedSigner,
		interval:          interval,
	}
}

// Start starts the AttestationGuardCollectorVerifier.
func (a AttestationGuardCollectorVerifier) Start(ctx context.Context) error {
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

// FindOldestGuardSubmittedToCollectorInProgressAttestation fetches the oldest signed attestation (by both notary and guard)
// that has been submitted to the Attestation Collector.
func (a AttestationGuardCollectorVerifier) FindOldestGuardSubmittedToCollectorInProgressAttestation(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveOldestGuardSubmittedToCollectorUnconfirmed(ctx, a.originDomain.Config().DomainID, a.destinationDomain.Config().DomainID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not retrieve oldest submitted-to-collector attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job of the collector verifier
//
//nolint:cyclop
func (a AttestationGuardCollectorVerifier) update(ctx context.Context) error {
	inProgressAttestationToVerifyOnCollector, err := a.FindOldestGuardSubmittedToCollectorInProgressAttestation(ctx)
	if err != nil {
		return fmt.Errorf("could not find oldest signed and unsubmitted attestation: %w", err)
	}
	if inProgressAttestationToVerifyOnCollector == nil {
		return nil
	}

	// TODO (joe): This will need to be updated.
	latestNonce, err := a.attestationDomain.AttestationCollector().GetLatestNonce(ctx, a.originDomain.Config().DomainID, a.destinationDomain.Config().DomainID, a.bondedSigner)
	if err != nil {
		return fmt.Errorf("could not find latest nonce: %w", err)
	}

	if latestNonce >= inProgressAttestationToVerifyOnCollector.SignedAttestation().Attestation().Nonce() {
		confirmedInProgressAttestation := types.NewInProgressAttestation(inProgressAttestationToVerifyOnCollector.SignedAttestation(), inProgressAttestationToVerifyOnCollector.OriginDispatchBlockNumber(), inProgressAttestationToVerifyOnCollector.SubmittedToAttestationCollectorTime(), 0)

		err = a.db.MarkGuardConfirmedOnAttestationCollector(ctx, confirmedInProgressAttestation)
		if err != nil {
			return fmt.Errorf("could not store confirmation block number for attestation: %w", err)
		}
	}

	if a.shouldResubmit(inProgressAttestationToVerifyOnCollector.SubmittedToAttestationCollectorTime()) {
		guardOnlySignedAttestation := types.NewSignedAttestation(
			inProgressAttestationToVerifyOnCollector.SignedAttestation().Attestation(),
			inProgressAttestationToVerifyOnCollector.SignedAttestation().GuardSignatures(),
			[]types.Signature{})
		err = a.attestationDomain.AttestationCollector().SubmitAttestation(ctx, a.unbondedSigner, guardOnlySignedAttestation)
		if err != nil {
			return fmt.Errorf("could not submit attestation: %w", err)
		}

		nowTime := time.Now()
		submittedInProgressAttestation := types.NewInProgressAttestation(
			inProgressAttestationToVerifyOnCollector.SignedAttestation(),
			inProgressAttestationToVerifyOnCollector.OriginDispatchBlockNumber(),
			&nowTime,
			0)

		err = a.db.ReUpdateGuardSubmittedToAttestationCollectorTime(ctx, submittedInProgressAttestation)
		if err != nil {
			return fmt.Errorf("could not store submission time for attestation: %w", err)
		}
	}

	return nil
}

func (a AttestationGuardCollectorVerifier) shouldResubmit(sentTime *time.Time) bool {
	// nowTime := time.Now()
	// TODO: (joe) figure this out.
	return false
}
