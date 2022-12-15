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
	// domain allows access to the origin contract
	domain domains.DomainClient
	// domain of target destination
	destinationID uint32
	// db is the synapse db
	db db.SynapseDB
	// signer is the signer
	signer signer.Signer
	// interval waits for an interval
	interval time.Duration
}

// NewOriginAttestationVerifier creates a new origin attestation verifier.
func NewOriginAttestationVerifier(domain domains.DomainClient, destinationID uint32, db db.SynapseDB, signer signer.Signer, interval time.Duration) OriginAttestationVerifier {
	return OriginAttestationVerifier{
		domain:        domain,
		destinationID: destinationID,
		db:            db,
		signer:        signer,
		interval:      interval,
	}
}

// Start starts the OriginAttestationVerifier.
func (a OriginAttestationVerifier) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(a.interval): // TODO: a.interval
			err := a.update(ctx)
			if err != nil {
				return err
			}
		}
	}
}

// FindOldestUnconfirmedAttestation fetches the oldest attestation that still needs to be confirmed.
func (a OriginAttestationVerifier) FindOldestUnconfirmedAttestation(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveOldestUnconfirmedSubmittedInProgressAttestation(ctx, a.domain.Config().DomainID, a.destinationID)
	if err != nil {
		if errors.Is(err, db.ErrNoNonceForDomain) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find oldest unconfirmed signed attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job for the verifier.
// nolint: cyclop
func (a OriginAttestationVerifier) update(ctx context.Context) error {
	// TODO (joe): we want to go through and update attestations for each destination.
	inProgressAttestationToConfirm, err := a.FindOldestUnconfirmedAttestation(ctx)
	if err != nil {
		return fmt.Errorf("could not find oldest unconfirmed attestation: %w", err)
	}
	if inProgressAttestationToConfirm == nil {
		return nil
	}

	// TODO (joe): This will need to be updated. Obviously we want to know when latest nonce was written and then
	// figure out if confirmation is enough in terms of currBlock - blockNumWhenWritten, etc
	latestNonce, currBlock, err := a.domain.AttestationCollector().GetLatestNonce(ctx, a.domain.Config().DomainID, a.destinationID, a.signer)
	if err != nil {
		return fmt.Errorf("could not find latest nonce: %w", err)
	}

	if latestNonce >= inProgressAttestationToConfirm.SignedAttestation().Attestation().Nonce() {
		confirmedInProgressAttestation := types.NewInProgressAttestation(inProgressAttestationToConfirm.SignedAttestation(), inProgressAttestationToConfirm.OriginDispatchBlockNumber(), inProgressAttestationToConfirm.SubmittedToAttestationCollectorTime(), currBlock, 0)

		err = a.db.UpdateConfirmedOnAttestationCollectorBlockNumber(ctx, confirmedInProgressAttestation)
		if err != nil {
			return fmt.Errorf("could not store confirmation block number for attestation: %w", err)
		}
	}

	if a.shouldResubmit(inProgressAttestationToConfirm.SubmittedToAttestationCollectorTime()) {
		err = a.domain.AttestationCollector().SubmitAttestation(ctx, a.signer, inProgressAttestationToConfirm.SignedAttestation())
		if err != nil {
			return fmt.Errorf("could not find submit attestation: %w", err)
		}

		nowTime := time.Now()
		submittedInProgressAttestation := types.NewInProgressAttestation(inProgressAttestationToConfirm.SignedAttestation(), inProgressAttestationToConfirm.OriginDispatchBlockNumber(), &nowTime, 0, 0)
		err = a.db.UpdateSubmittedToAttestationCollectorTime(ctx, submittedInProgressAttestation)
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
