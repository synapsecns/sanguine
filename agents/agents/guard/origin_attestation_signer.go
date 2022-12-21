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

// OriginAttestationSigner signs unsigned attestations that have been fetched for particular origin-destination pair.
// TODO: this needs to become an interface.
type OriginAttestationSigner struct {
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

// NewOriginAttestationSigner creates a new origin attestation signer.
func NewOriginAttestationSigner(domain domains.DomainClient, destinationID uint32, db db.SynapseDB, signer signer.Signer, interval time.Duration) OriginAttestationSigner {
	return OriginAttestationSigner{
		domain:        domain,
		destinationID: destinationID,
		db:            db,
		signer:        signer,
		interval:      interval,
	}
}

// Start starts the OriginAttestationSigner.
func (a OriginAttestationSigner) Start(ctx context.Context) error {
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

// FindOldestUnsignedAttestation fetches the oldest attestation that still needs to be signed.
func (a OriginAttestationSigner) FindOldestUnsignedAttestation(ctx context.Context) (types.InProgressAttestation, error) {
	inProgressAttestation, err := a.db.RetrieveOldestUnsignedInProgressAttestation(ctx, a.domain.Config().DomainID, a.destinationID)
	if err != nil {
		if errors.Is(err, db.ErrNoNonceForDomain) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not oldest unsigned attestation: %w", err)
	}
	return inProgressAttestation, nil
}

// update runs the job of th signer
// nolint: cyclop
func (a OriginAttestationSigner) update(ctx context.Context) error {
	inProgressAttestationToSign, err := a.FindOldestUnsignedAttestation(ctx)
	if err != nil {
		return fmt.Errorf("could not find oldest unsigned attestation: %w", err)
	}
	if inProgressAttestationToSign == nil {
		return nil
	}

	hashedAttestation, err := types.Hash(inProgressAttestationToSign.SignedAttestation().Attestation())
	if err != nil {
		return fmt.Errorf("could not hash update: %w", err)
	}
	signature, err := a.signer.SignMessage(ctx, core.BytesToSlice(hashedAttestation), false)
	if err != nil {
		return fmt.Errorf("could not sign message: %w", err)
	}

	signedAttestation := types.NewSignedAttestation(inProgressAttestationToSign.SignedAttestation().Attestation(), []types.Signature{}, []types.Signature{signature})
	signedInProgressAttestation := types.NewInProgressAttestation(signedAttestation, inProgressAttestationToSign.OriginDispatchBlockNumber(), nil, 0)
	err = a.db.UpdateNotarySignature(ctx, signedInProgressAttestation)
	if err != nil {
		return fmt.Errorf("could not store signature for attestation: %w", err)
	}

	return nil
}
