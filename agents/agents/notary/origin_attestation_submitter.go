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

// NewOriginAttestationSubmitter creates a new origin attestation submitter.
func NewOriginAttestationSubmitter(domain domains.DomainClient, destinationID uint32, db db.SynapseDB, signer signer.Signer, interval time.Duration) OriginAttestationSubmitter {
	return OriginAttestationSubmitter{
		domain:        domain,
		destinationID: destinationID,
		db:            db,
		signer:        signer,
		interval:      interval,
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

// FindOldestUnsubmittedAttestation fetches the oldest unsubmitted attestation that has been signed
func (a OriginAttestationSubmitter) FindOldestUnsubmittedAttestation(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveOldestUnsubmittedSignedInProgressAttestation(ctx, a.domain.Config().DomainID, a.destinationID)
	if err != nil {
		if errors.Is(err, db.ErrNoNonceForDomain) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find oldest unsubmitted signed attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job for the submitter
// nolint: cyclop
func (a OriginAttestationSubmitter) update(ctx context.Context) error {
	inProgressAttesationToSubmit, err := a.FindOldestUnsubmittedAttestation(ctx)
	if err != nil {
		return fmt.Errorf("could not find oldest unsubmitted attestation: %w", err)
	}
	if inProgressAttesationToSubmit == nil {
		return nil
	}

	err = a.domain.AttestationCollector().SubmitAttestation(ctx, a.signer, inProgressAttesationToSubmit.SignedAttestation())
	if err != nil {
		return fmt.Errorf("could not find submit attestation: %w", err)
	}

	nowTime := time.Now()
	submittedInProgressAttestation := types.NewInProgressAttestation(inProgressAttesationToSubmit.SignedAttestation(), inProgressAttesationToSubmit.OriginDispatchBlockNumber(), &nowTime, 0)
	err = a.db.UpdateSubmittedToAttestationCollectorTime(ctx, submittedInProgressAttestation)
	if err != nil {
		return fmt.Errorf("could not store submission time for attestation: %w", err)
	}

	return nil
}
