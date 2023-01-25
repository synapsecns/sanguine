package notary

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

// OriginAttestationSigner signs unsigned attestations that have been fetched for particular origin-destination pair.
// TODO: this needs to become an interface.
type OriginAttestationSigner struct {
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

// NewOriginAttestationSigner creates a new origin attestation signer.
func NewOriginAttestationSigner(
	originDomain domains.DomainClient,
	attestationDomain domains.DomainClient,
	destinationDomain domains.DomainClient,
	db db.SynapseDB,
	bondedSigner signer.Signer,
	unbondedSigner signer.Signer,
	interval time.Duration) OriginAttestationSigner {
	return OriginAttestationSigner{
		originDomain:      originDomain,
		attestationDomain: attestationDomain,
		destinationDomain: destinationDomain,
		db:                db,
		bondedSigner:      bondedSigner,
		unbondedSigner:    unbondedSigner,
		interval:          interval,
	}
}

// Start starts the OriginAttestationSigner.
func (a OriginAttestationSigner) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logger.Info("Notary OriginAttestationSigner exiting without error")
			return nil
		case <-time.After(a.interval):
			err := a.update(ctx)
			if err != nil {
				logger.Errorf("Notary OriginAttestationSigner exiting with error: %v", err)
				return err
			}
		}
	}
}

// FindNewestUnsignedAttestation fetches the newest attestation that still needs to be signed.
func (a OriginAttestationSigner) FindNewestUnsignedAttestation(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveNewestInProgressAttestationIfInState(
		ctx,
		a.originDomain.Config().DomainID,
		a.destinationDomain.Config().DomainID,
		types.AttestationStateNotaryUnsigned)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not newest unsigned attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job of th signer
//
//nolint:cyclop
func (a OriginAttestationSigner) update(ctx context.Context) error {
	inProgressAttestationToSign, err := a.FindNewestUnsignedAttestation(ctx)
	if err != nil {
		return fmt.Errorf("could not find newest unsigned attestation: %w", err)
	}

	if inProgressAttestationToSign == nil {
		return nil
	}

	hashedAttestation, err := types.Hash(inProgressAttestationToSign.SignedAttestation().Attestation())
	if err != nil {
		return fmt.Errorf("could not hash update: %w", err)
	}

	signature, err := a.bondedSigner.SignMessage(ctx, core.BytesToSlice(hashedAttestation), false)
	if err != nil {
		return fmt.Errorf("could not sign message: %w", err)
	}

	signedAttestation := types.NewSignedAttestation(inProgressAttestationToSign.SignedAttestation().Attestation(), []types.Signature{}, []types.Signature{signature})
	signedInProgressAttestation := types.NewInProgressAttestation(signedAttestation, nil, 0)
	err = a.db.UpdateNotarySignature(ctx, signedInProgressAttestation)
	if err != nil {
		return fmt.Errorf("could not store notary signature for attestation: %w", err)
	}

	return nil
}
