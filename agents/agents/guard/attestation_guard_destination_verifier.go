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

// AttestationGuardDestinationVerifier verifies the signed attestation (by both notary and guard)
// is posted successfully on the Destination.
// TODO: this needs to become an interface.
type AttestationGuardDestinationVerifier struct {
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

// NewAttestationGuardDestinationVerifier creates a new attestation guard destination verifier.
func NewAttestationGuardDestinationVerifier(
	originDomain domains.DomainClient,
	attestationDomain domains.DomainClient,
	destinationDomain domains.DomainClient,
	db db.SynapseDB,
	bondedSigner signer.Signer,
	unbondedSigner signer.Signer,
	interval time.Duration) AttestationGuardDestinationVerifier {
	return AttestationGuardDestinationVerifier{
		originDomain:      originDomain,
		attestationDomain: attestationDomain,
		destinationDomain: destinationDomain,
		db:                db,
		bondedSigner:      bondedSigner,
		unbondedSigner:    unbondedSigner,
		interval:          interval,
	}
}

// Start starts the AttestationGuardDestinationVerifier.
func (a AttestationGuardDestinationVerifier) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logger.Info("Guard AttestationGuardDestinationVerifier exiting without error")
			return nil
		case <-time.After(a.interval):
			err := a.update(ctx)
			if err != nil {
				logger.Errorf("Guard AttestationGuardDestinationVerifier exiting with error: %v", err)
				return err
			}
		}
	}
}

// FindNewestSubmittedToDestinationInProgressAttestation fetches the newest signed attestation (by both notary and guard)
// that has been submitted to the Destination.
func (a AttestationGuardDestinationVerifier) FindNewestSubmittedToDestinationInProgressAttestation(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveNewestInProgressAttestationIfInState(
		ctx,
		a.originDomain.Config().DomainID,
		a.destinationDomain.Config().DomainID,
		types.AttestationStateSubmittedToDestinationUnconfirmed)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not retrieve newest submitted-to-destination attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job of the destination verifier
//
//nolint:cyclop
func (a AttestationGuardDestinationVerifier) update(ctx context.Context) error {
	inProgressAttestationToVerifyOnDestination, err := a.FindNewestSubmittedToDestinationInProgressAttestation(ctx)
	if err != nil {
		return fmt.Errorf("could not find newest sumbitted-to-destination attestation: %w", err)
	}
	if inProgressAttestationToVerifyOnDestination == nil {
		return nil
	}

	// TODO (joe): This will need to be updated.
	submittedAtTime, err := a.destinationDomain.Destination().SubmittedAt(ctx, a.originDomain.Config().DomainID, inProgressAttestationToVerifyOnDestination.SignedAttestation().Attestation().Root())
	if err != nil {
		return fmt.Errorf("error trying to get submitted at time from destination: %w", err)
	}

	if submittedAtTime != nil {
		confirmedInProgressAttestation := types.NewInProgressAttestation(inProgressAttestationToVerifyOnDestination.SignedAttestation(), inProgressAttestationToVerifyOnDestination.SubmittedToAttestationCollectorTime(), 0)

		err = a.db.MarkConfirmedOnDestination(ctx, confirmedInProgressAttestation)
		if err != nil {
			return fmt.Errorf("error in MarkConfirmedOnDestination: %w", err)
		}
	}

	if a.shouldResubmit(inProgressAttestationToVerifyOnDestination.SubmittedToDestinationTime()) {
		// TODO (joe): Figure out what to do here. For now, just error.
		return fmt.Errorf("error: not able to confirm attestation on destination")
	}

	return nil
}

func (a AttestationGuardDestinationVerifier) shouldResubmit(sentTime *time.Time) bool {
	// nowTime := time.Now()
	// TODO: (joe) figure this out.
	return false
}
