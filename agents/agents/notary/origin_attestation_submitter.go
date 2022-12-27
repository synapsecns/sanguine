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

// OriginAttestationSubmitter submits signed attestations to the attestation collector for particular origin-destination pair.
// TODO: this needs to become an interface.
type OriginAttestationSubmitter struct {
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

// NewOriginAttestationSubmitter creates a new origin attestation submitter.
func NewOriginAttestationSubmitter(
	originDomain domains.DomainClient,
	attestationDomain domains.DomainClient,
	destinationDomain domains.DomainClient,
	db db.SynapseDB,
	bondedSigner signer.Signer,
	unbondedSigner signer.Signer,
	interval time.Duration) OriginAttestationSubmitter {
	return OriginAttestationSubmitter{
		originDomain:      originDomain,
		attestationDomain: attestationDomain,
		destinationDomain: destinationDomain,
		db:                db,
		bondedSigner:      bondedSigner,
		unbondedSigner:    unbondedSigner,
		interval:          interval,
	}
}

// Start starts the OriginAttestationSubmitter.
func (a OriginAttestationSubmitter) Start(ctx context.Context) error {
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

// FindOldestUnsubmittedAttestation fetches the oldest unsubmitted attestation that has been signed.
func (a OriginAttestationSubmitter) FindOldestUnsubmittedAttestation(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveOldestUnsubmittedSignedInProgressAttestation(ctx, a.originDomain.Config().DomainID, a.destinationDomain.Config().DomainID)
	if err != nil {
		if errors.Is(err, db.ErrNoNonceForDomain) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find oldest unsubmitted signed attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job for the submitter.
// nolint: cyclop
func (a OriginAttestationSubmitter) update(ctx context.Context) error {
	inProgressAttestationToSubmit, err := a.FindOldestUnsubmittedAttestation(ctx)
	if err != nil {
		return fmt.Errorf("could not find oldest unsubmitted attestation: %w", err)
	}
	if inProgressAttestationToSubmit == nil {
		return nil
	}

	err = a.attestationDomain.AttestationCollector().SubmitAttestation(ctx, a.unbondedSigner, inProgressAttestationToSubmit.SignedAttestation())
	if err != nil {
		return fmt.Errorf("could not find submit attestation: %w", err)
	}

	nowTime := time.Now()
	submittedInProgressAttestation := types.NewInProgressAttestation(inProgressAttestationToSubmit.SignedAttestation(), inProgressAttestationToSubmit.OriginDispatchBlockNumber(), &nowTime, 0)
	err = a.db.UpdateSubmittedToAttestationCollectorTime(ctx, submittedInProgressAttestation)
	if err != nil {
		return fmt.Errorf("could not store submission time for attestation: %w", err)
	}

	return nil
}
