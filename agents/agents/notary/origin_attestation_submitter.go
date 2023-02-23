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
			logger.Info("Notary OriginAttestationSubmitter exiting without error")
			return nil
		case <-time.After(a.interval):
			err := a.update(ctx)
			if err != nil {
				logger.Errorf("Notary OriginAttestationSubmitter exiting with error: %v", err)
				return err
			}
		}
	}
}

// FindNewestUnsubmittedAttestation fetches the newest unsubmitted attestation that has been signed.
func (a OriginAttestationSubmitter) FindNewestUnsubmittedAttestation(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveNewestInProgressAttestationIfInState(
		ctx,
		a.originDomain.Config().DomainID,
		a.destinationDomain.Config().DomainID,
		types.AttestationStateNotarySignedUnsubmitted)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find newest unsubmitted signed attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job for the submitter.
//
//nolint:cyclop
func (a OriginAttestationSubmitter) update(ctx context.Context) error {
	inProgressAttestationToSubmit, err := a.FindNewestUnsubmittedAttestation(ctx)
	if err != nil {
		return fmt.Errorf("could not find newest unsubmitted attestation: %w", err)
	}
	if inProgressAttestationToSubmit == nil {
		return nil
	}

	signedAttestation, err := a.attestationDomain.AttestationCollector().GetAttestation(
		ctx,
		inProgressAttestationToSubmit.SignedAttestation().Attestation().Origin(),
		inProgressAttestationToSubmit.SignedAttestation().Attestation().Destination(),
		inProgressAttestationToSubmit.SignedAttestation().Attestation().Nonce())
	if err != nil {
		if err != domains.ErrNoUpdate {
			return fmt.Errorf("could GetAttestation from collector to see if we already signed and submitted: %w", err)
		}
	}

	if signedAttestation == nil || len(signedAttestation.NotarySignatures()) == 0 {
		err = a.attestationDomain.AttestationCollector().SubmitAttestation(ctx, a.unbondedSigner, inProgressAttestationToSubmit.SignedAttestation())
		if err != nil {
			logger.Errorf("Error calling SubmitAttestation on AttestationCollector: %v", err)
			return fmt.Errorf("could not submit attestation: %w", err)
		}
	}

	nowTime := time.Now()
	submittedInProgressAttestation := types.NewInProgressAttestation(inProgressAttestationToSubmit.SignedAttestation(), &nowTime, 0)
	err = a.db.UpdateNotarySubmittedToAttestationCollectorTime(ctx, submittedInProgressAttestation)
	if err != nil {
		return fmt.Errorf("could not store submission time for attestation: %w", err)
	}

	return nil
}
