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

// AttestationGuardCollectorSubmitter submits the signed attestation (by both notary and guard)
// to the Attestation Collector.
// TODO: this needs to become an interface.
type AttestationGuardCollectorSubmitter struct {
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

// NewAttestationGuardCollectorSubmitter creates a new attestation guard collector submitter.
func NewAttestationGuardCollectorSubmitter(
	originDomain domains.DomainClient,
	attestationDomain domains.DomainClient,
	destinationDomain domains.DomainClient,
	db db.SynapseDB,
	bondedSigner signer.Signer,
	unbondedSigner signer.Signer,
	interval time.Duration) AttestationGuardCollectorSubmitter {
	return AttestationGuardCollectorSubmitter{
		originDomain:      originDomain,
		attestationDomain: attestationDomain,
		destinationDomain: destinationDomain,
		db:                db,
		bondedSigner:      bondedSigner,
		unbondedSigner:    unbondedSigner,
		interval:          interval,
	}
}

// Start starts the AttestationGuardCollectorSubmitter.
func (a AttestationGuardCollectorSubmitter) Start(ctx context.Context) error {
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

// FindOldestGuardUnsubmittedSignedInProgressAttestation fetches the oldest signed attestation (by both notary and guard)
// that has not yet been submitted to the Attestation Collector.
func (a AttestationGuardCollectorSubmitter) FindOldestGuardUnsubmittedSignedInProgressAttestation(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveOldestGuardUnsubmittedSignedInProgressAttestation(ctx, a.originDomain.Config().DomainID, a.destinationDomain.Config().DomainID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not retrieve oldest signed and unsubmitted attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job of the collector submitter
// nolint: cyclop
func (a AttestationGuardCollectorSubmitter) update(ctx context.Context) error {
	inProgressAttestationToSubmitToCollector, err := a.FindOldestGuardUnsubmittedSignedInProgressAttestation(ctx)
	if err != nil {
		return fmt.Errorf("could not find oldest signed and unsubmitted attestation: %w", err)
	}
	if inProgressAttestationToSubmitToCollector == nil {
		return nil
	}

	guardOnlySignedAttestation := types.NewSignedAttestation(
		inProgressAttestationToSubmitToCollector.SignedAttestation().Attestation(),
		inProgressAttestationToSubmitToCollector.SignedAttestation().GuardSignatures(),
		[]types.Signature{})
	err = a.attestationDomain.AttestationCollector().SubmitAttestation(ctx, a.unbondedSigner, guardOnlySignedAttestation)
	if err != nil {
		fmt.Printf("\nCRONIN SubmitAttestationCollector %v\n", err)
		return fmt.Errorf("could not submit attestation: %w", err)
	}

	nowTime := time.Now()
	submittedInProgressAttestation := types.NewInProgressAttestation(
		inProgressAttestationToSubmitToCollector.SignedAttestation(),
		inProgressAttestationToSubmitToCollector.OriginDispatchBlockNumber(),
		&nowTime,
		0)

	err = a.db.UpdateGuardSubmittedToAttestationCollectorTime(ctx, submittedInProgressAttestation)
	if err != nil {
		return fmt.Errorf("could not store submission time for attestation: %w", err)
	}

	return nil
}
