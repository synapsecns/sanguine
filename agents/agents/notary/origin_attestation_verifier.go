package notary

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

// OriginAttestationVerifier verifies that the AttestationCollector in fact posted submitted attestations and if not resubmits.
// TODO: this needs to become an interface.
type OriginAttestationVerifier struct {
	// originDomain allows access to the origin contract on the origin chain
	originDomain domains.DomainClient
	// attestationDomain allows access to the atttestation contract on the SYN chain
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

// NewOriginAttestationVerifier creates a new origin attestation verifier.
func NewOriginAttestationVerifier(
	originDomain domains.DomainClient,
	attestationDomain domains.DomainClient,
	destinationDomain domains.DomainClient,
	db db.SynapseDB,
	bondedSigner signer.Signer,
	unbondedSigner signer.Signer,
	interval time.Duration) OriginAttestationVerifier {
	return OriginAttestationVerifier{
		originDomain:      originDomain,
		attestationDomain: attestationDomain,
		destinationDomain: destinationDomain,
		db:                db,
		bondedSigner:      bondedSigner,
		unbondedSigner:    unbondedSigner,
		interval:          interval,
	}
}

// Start starts the OriginAttestationVerifier.
func (a OriginAttestationVerifier) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logger.Info("Notary OriginAttestationSubmitter exiting without error")
			return nil
		case <-time.After(a.interval):
			err := a.update(ctx)
			if err != nil {
				logger.Errorf("Notary OriginAttestationVerifier exiting with error: %v", err)
				return err
			}
		}
	}
}

// FindNewestUnconfirmedAttestation fetches the newest attestation that still needs to be confirmed.
func (a OriginAttestationVerifier) FindNewestUnconfirmedAttestation(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveNewestInProgressAttestationIfInState(
		ctx,
		a.originDomain.Config().DomainID,
		a.destinationDomain.Config().DomainID,
		types.AttestationStateNotarySubmittedUnconfirmed)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find newest unconfirmed signed attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job for the verifier.
//
//nolint:cyclop
func (a OriginAttestationVerifier) update(ctx context.Context) error {
	// TODO (joe): we want to go through and update attestations for each destination.
	inProgressAttestationToConfirm, err := a.FindNewestUnconfirmedAttestation(ctx)
	if err != nil {
		return fmt.Errorf("could not find newest unconfirmed attestation: %w", err)
	}
	if inProgressAttestationToConfirm == nil {
		return nil
	}

	// TODO (joe): This will need to be updated.
	latestNonce, err := a.attestationDomain.AttestationCollector().GetLatestNonce(ctx, a.originDomain.Config().DomainID, a.destinationDomain.Config().DomainID, a.bondedSigner)
	if err != nil {
		return fmt.Errorf("could not find latest nonce: %w", err)
	}

	if latestNonce >= inProgressAttestationToConfirm.SignedAttestation().Attestation().Nonce() {
		confirmedInProgressAttestation := types.NewInProgressAttestation(inProgressAttestationToConfirm.SignedAttestation(), inProgressAttestationToConfirm.SubmittedToAttestationCollectorTime(), 0)

		err = a.db.MarkNotaryConfirmedOnAttestationCollector(ctx, confirmedInProgressAttestation)
		if err != nil {
			return fmt.Errorf("could not store confirmation block number for attestation: %w", err)
		}
	}

	if a.shouldResubmit(inProgressAttestationToConfirm.SubmittedToAttestationCollectorTime()) {
		err = a.attestationDomain.AttestationCollector().SubmitAttestation(ctx, a.unbondedSigner, inProgressAttestationToConfirm.SignedAttestation())
		if err != nil {
			return fmt.Errorf("could not submit attestation: %w", err)
		}

		nowTime := time.Now()
		submittedInProgressAttestation := types.NewInProgressAttestation(inProgressAttestationToConfirm.SignedAttestation(), &nowTime, 0)
		err = a.db.ReUpdateNotarySubmittedToAttestationCollectorTime(ctx, submittedInProgressAttestation)
		if err != nil {
			return fmt.Errorf("could not store submission time for attestation: %w", err)
		}
	}

	return nil
}

func (a OriginAttestationVerifier) shouldResubmit(sentTime *time.Time) bool {
	// nowTime := time.Now()
	// TODO: (joe) figure this out.
	return false
}
