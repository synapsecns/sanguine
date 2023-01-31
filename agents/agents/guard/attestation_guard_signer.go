package guard

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// AttestationGuardSigner signs the attestation after it has been verified on origin.
// TODO: this needs to become an interface.
type AttestationGuardSigner struct {
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

// NewAttestationGuardSigner creates a new attestation guard signer.
func NewAttestationGuardSigner(
	originDomain domains.DomainClient,
	attestationDomain domains.DomainClient,
	destinationDomain domains.DomainClient,
	db db.SynapseDB,
	bondedSigner signer.Signer,
	unbondedSigner signer.Signer,
	interval time.Duration) AttestationGuardSigner {
	return AttestationGuardSigner{
		originDomain:      originDomain,
		attestationDomain: attestationDomain,
		destinationDomain: destinationDomain,
		db:                db,
		bondedSigner:      bondedSigner,
		unbondedSigner:    unbondedSigner,
		interval:          interval,
	}
}

// Start starts the AttestationGuardSigner.
func (a AttestationGuardSigner) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logger.Info("Guard AttestationGuardSigner exiting without error")
			return nil
		case <-time.After(a.interval):
			err := a.update(ctx)
			if err != nil {
				logger.Errorf("Guard AttestationGuardSigner exiting with error: %v", err)
				return err
			}
		}
	}
}

// FindNewestGuardUnsignedAndVerifiedAttestation fetches the newest attestation that still needs to be signed by the guard but has been verified on origin.
func (a AttestationGuardSigner) FindNewestGuardUnsignedAndVerifiedAttestation(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveNewestInProgressAttestationIfInState(
		ctx,
		a.originDomain.Config().DomainID,
		a.destinationDomain.Config().DomainID,
		types.AttestationStateGuardUnsignedAndVerified)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not retrieve newest unsigned and verified attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job of the signer
//
//nolint:cyclop
func (a AttestationGuardSigner) update(ctx context.Context) error {
	inProgressAttestationToSign, err := a.FindNewestGuardUnsignedAndVerifiedAttestation(ctx)
	if err != nil {
		return fmt.Errorf("could not find newest unsigned and verified attestation: %w", err)
	}
	if inProgressAttestationToSign == nil {
		return nil
	}

	hashedAttestation, err := types.Hash(inProgressAttestationToSign.SignedAttestation().Attestation())
	if err != nil {
		return fmt.Errorf("could not hash update: %w", err)
	}

	guardSignature, err := a.bondedSigner.SignMessage(ctx, core.BytesToSlice(hashedAttestation), false)
	if err != nil {
		return fmt.Errorf("could not sign message: %w", err)
	}

	signedAttestation := types.NewSignedAttestation(inProgressAttestationToSign.SignedAttestation().Attestation(), []types.Signature{guardSignature}, []types.Signature{})
	signedInProgressAttestation := types.NewInProgressAttestation(signedAttestation, nil, 0)
	err = a.db.UpdateGuardSignature(ctx, signedInProgressAttestation)
	if err != nil {
		return fmt.Errorf("could not store guard signature for attestation: %w", err)
	}

	return nil
}
