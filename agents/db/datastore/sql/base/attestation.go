package base

import (
	"context"
	"errors"
	"fmt"

	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StoreSignedAttestations stores signed attestations.
func (s Store) StoreSignedAttestations(ctx context.Context, attestation types.SignedAttestation) error {
	sig, err := types.EncodeSignature(attestation.Signature())
	if err != nil {
		return fmt.Errorf("could not encode signature: %w", err)
	}

	tx := s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: DomainIDFieldName}, {Name: NonceFieldName}},
		DoNothing: true,
	}).Create(&SignedAttestation{
		SADomain:    attestation.Attestation().Domain(),
		SANonce:     attestation.Attestation().Nonce(),
		SARoot:      core.BytesToSlice(attestation.Attestation().Root()),
		SASignature: sig,
	})

	if tx.Error != nil {
		return fmt.Errorf("could not store signed attestations: %w", tx.Error)
	}
	return nil
}

// RetrieveSignedAttestationByNonce retrieves a signed attestation by nonce.
func (s Store) RetrieveSignedAttestationByNonce(ctx context.Context, domainID, nonce uint32) (attestation types.SignedAttestation, err error) {
	var signedAttestation SignedAttestation
	tx := s.DB().WithContext(ctx).Model(&SignedAttestation{}).Where(&SignedAttestation{
		SADomain: domainID,
		SANonce:  nonce,
	}).First(&signedAttestation)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("could not find attestation with nonce %d and domain %d: %w", nonce, domainID, dbcommon.ErrNotFound)
		}
		return nil, fmt.Errorf("could not store attestation: %w", tx.Error)
	}
	return signedAttestation, err
}
