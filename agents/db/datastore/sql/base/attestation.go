package base

import (
	"context"
	"errors"
	"fmt"

	"github.com/synapsecns/sanguine/agents/db"

	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StoreSignedAttestations stores signed attestations.
func (s Store) StoreSignedAttestations(ctx context.Context, attestation types.SignedAttestation) error {
	// TODO (joe): Fix this to handle multiple guard and notary sigs??? Or have other type representing single agent signature.
	if len(attestation.GuardSignatures()) != 0 {
		return fmt.Errorf("currently only handle signed attestation with no guard signatures. Num guard sigs: %d", len(attestation.GuardSignatures()))
	}
	if len(attestation.NotarySignatures()) != 1 {
		return fmt.Errorf("currently only handle signed attestation with single notary signature. Num notary sigs: %d", len(attestation.NotarySignatures()))
	}
	guardSig, err := types.EncodeSignature(attestation.GuardSignatures()[0])
	if err != nil {
		return fmt.Errorf("could not encode signature: %w", err)
	}
	notarySig, err := types.EncodeSignature(attestation.NotarySignatures()[0])
	if err != nil {
		return fmt.Errorf("could not encode signature: %w", err)
	}

	tx := s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: DomainIDFieldName}, {Name: NonceFieldName}},
		DoNothing: true,
	}).Create(&SignedAttestation{
		SAOrigin:          attestation.Attestation().Origin(),
		SADestination:     attestation.Attestation().Destination(),
		SANonce:           attestation.Attestation().Nonce(),
		SARoot:            core.BytesToSlice(attestation.Attestation().Root()),
		SAGuardSignature:  guardSig,
		SANotarySignature: notarySig,
	})

	if tx.Error != nil {
		return fmt.Errorf("could not store signed attestations: %w", tx.Error)
	}
	return nil
}

// RetrieveSignedAttestationByNonce retrieves a signed attestation by nonce.
// TODO (joe): This will need to be updated after we make the Global Registry changes.
func (s Store) RetrieveSignedAttestationByNonce(ctx context.Context, domainID, nonce uint32) (attestation types.SignedAttestation, err error) {
	var signedAttestation SignedAttestation
	tx := s.DB().WithContext(ctx).Model(&SignedAttestation{}).Where(&SignedAttestation{
		SAOrigin: domainID,
		SANonce:  nonce,
	}).First(&signedAttestation)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("could not find attestation with nonce %d and domain %d: %w", nonce, domainID, db.ErrNotFound)
		}
		return nil, fmt.Errorf("could not store attestation: %w", tx.Error)
	}
	return signedAttestation, err
}
