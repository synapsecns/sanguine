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

// AttestationGuardDestinationSubmitter submits the signed attestation (by both notary and guard)
// to the Destination.
// TODO: this needs to become an interface.
type AttestationGuardDestinationSubmitter struct {
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

// NewAttestationGuardDestinationSubmitter creates a new attestation guard destination submitter.
func NewAttestationGuardDestinationSubmitter(
	originDomain domains.DomainClient,
	attestationDomain domains.DomainClient,
	destinationDomain domains.DomainClient,
	db db.SynapseDB,
	bondedSigner signer.Signer,
	unbondedSigner signer.Signer,
	interval time.Duration) AttestationGuardDestinationSubmitter {
	return AttestationGuardDestinationSubmitter{
		originDomain:      originDomain,
		attestationDomain: attestationDomain,
		destinationDomain: destinationDomain,
		db:                db,
		bondedSigner:      bondedSigner,
		unbondedSigner:    unbondedSigner,
		interval:          interval,
	}
}

// Start starts the AttestationGuardDestinationSubmitter.
func (a AttestationGuardDestinationSubmitter) Start(ctx context.Context) error {
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

// FindOldestGuardConfirmedOnCollector fetches the oldest signed attestation (by both notary and guard)
// that has been submitted to and confirmed on the Attestation Collector.
func (a AttestationGuardDestinationSubmitter) FindOldestGuardConfirmedOnCollector(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveOldestGuardConfirmedOnCollector(ctx, a.originDomain.Config().DomainID, a.destinationDomain.Config().DomainID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not retrieve oldest confirmed attestation on attestation collector: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job of the destination submitter
//
//nolint:cyclop
func (a AttestationGuardDestinationSubmitter) update(ctx context.Context) error {
	inProgressAttestationToSubmitToDestination, err := a.FindOldestGuardConfirmedOnCollector(ctx)
	if err != nil {
		return fmt.Errorf("could not find oldest confirmed-on-collector attestation: %w", err)
	}
	if inProgressAttestationToSubmitToDestination == nil {
		return nil
	}

	// TODO (joe): Double check that attestation has both guard and notary signatures.

	err = a.destinationDomain.Destination().SubmitAttestation(ctx, a.unbondedSigner, inProgressAttestationToSubmitToDestination.SignedAttestation())
	if err != nil {
		return fmt.Errorf("could not submit attestation to destination: %w", err)
	}

	submittedInProgressAttestation := types.NewInProgressAttestation(
		inProgressAttestationToSubmitToDestination.SignedAttestation(),
		inProgressAttestationToSubmitToDestination.OriginDispatchBlockNumber(),
		nil,
		0)

	err = a.db.UpdateSubmittedToDestinationTime(ctx, submittedInProgressAttestation)
	if err != nil {
		return fmt.Errorf("could not store submission time for attestation: %w", err)
	}

	return nil
}
