package base

import (
	"context"
	"errors"
	"fmt"

	"github.com/Thor-x86/nullable"
	"github.com/synapsecns/sanguine/agents/db"

	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StoreNewInProgressAttestation stores in-progress attestation only if it hasn't already been stored
func (s Store) StoreNewInProgressAttestation(ctx context.Context, attestation types.Attestation, originDispathBlockNumber uint64) error {
	if originDispathBlockNumber == uint64(0) {
		return fmt.Errorf("StoreNewInProgressAttestation called on attestation with a 0 originDispathBlockNumber")
	}
	// We only want to store if not already stored
	tx := s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "origin"}, {Name: "destination"}, {Name: NonceFieldName}},
		DoNothing: true,
	}).Create(&InProgressAttestation{
		IPOrigin:                    attestation.Origin(),
		IPDestination:               attestation.Destination(),
		IPNonce:                     attestation.Nonce(),
		IPRoot:                      core.BytesToSlice(attestation.Root()),
		IPOriginDispatchBlockNumber: originDispathBlockNumber,
	})

	if tx.Error != nil {
		return fmt.Errorf("could not store signed attestations: %w", tx.Error)
	}
	return nil
}

// UpdateSignature sets the signature of the in-progress Attestation
func (s Store) UpdateSignature(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	if inProgressAttestation.SignedAttestation().Signature() == nil {
		return fmt.Errorf("UpdateSignature called on attestation with a nil signature")
	}
	sig, err := types.EncodeSignature(inProgressAttestation.SignedAttestation().Signature())
	if err != nil {
		return fmt.Errorf("could not encode signature: %w", err)
	}

	signatureNull := fmt.Sprintf("`%s` = '' or `%s` IS NULL", "signature", "signature")
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(signatureNull).
		Update("signature", sig)

	if tx.Error != nil {
		return fmt.Errorf("could not set signature for in-progress attestations: %w", tx.Error)
	}
	return nil
}

// UpdateSubmittedToAttestationCollectorTime sets the time attestation was sent to Attesttion Collector
func (s Store) UpdateSubmittedToAttestationCollectorTime(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	if inProgressAttestation.SubmittedToAttestationCollectorTime() == nil {
		return fmt.Errorf("UpdateSubmittedToAttestationCollectorTime called on attestation with a nil time")
	}

	submittedTimeLessThanNow := fmt.Sprintf("`%s` is NULL or `%s` < ?", "submitted_to_attestation_collector_time", "submitted_to_attestation_collector_time")
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(submittedTimeLessThanNow, inProgressAttestation.SubmittedToAttestationCollectorTime()).
		Update("submitted_to_attestation_collector_time", inProgressAttestation.SubmittedToAttestationCollectorTime())

	if tx.Error != nil {
		return fmt.Errorf("could not update SubmittedToAttestationCollectorTime for in-progress attestations: %w", tx.Error)
	}
	return nil
}

// UpdateConfirmedOnAttestationCollectorBlockNumber sets the block number we confirmed the attestation on the Attesttion Collector
func (s Store) UpdateConfirmedOnAttestationCollectorBlockNumber(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	if inProgressAttestation.ConfirmedOnAttestationCollectorBlockNumber() == uint64(0) {
		return fmt.Errorf("ConfirmedOnAttestationCollectorBlockNumber called on attestation with a 0 ConfirmedOnAttestationCollectorBlockNumber")
	}

	confirmedBlockLessThanNow := fmt.Sprintf("`%s` < ?", "confirmed_on_attestation_collector_block_number")
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(confirmedBlockLessThanNow, inProgressAttestation.ConfirmedOnAttestationCollectorBlockNumber()).
		Update("confirmed_on_attestation_collector_block_number", inProgressAttestation.ConfirmedOnAttestationCollectorBlockNumber())

	if tx.Error != nil {
		return fmt.Errorf("could not set ConfirmedOnAttestationCollectorBlockNumber for in-progress attestation: %w", tx.Error)
	}
	return nil
}

// RetrieveLatestCachedNonce retrieves the latest nonce cached for given origin-destination pair.
func (s Store) RetrieveLatestCachedNonce(ctx context.Context, originID, destinationID uint32) (_ uint32, err error) {
	var nonce nullable.Uint32

	selectMaxNonce := fmt.Sprintf("max(`%s`)", NonceFieldName)

	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Select(selectMaxNonce).
		Where(&InProgressAttestation{IPOrigin: originID, IPDestination: destinationID}).
		Scan(&nonce)

	if tx.Error != nil {
		return 0, fmt.Errorf("could not get nonce for origin %d and destiniation %d: %w", originID, destinationID, tx.Error)
	}

	// if no nonces, return the corresponding eror.
	if nonce.Get() == nil {
		return 0, db.ErrNoNonceForDomain
	}
	return *nonce.Get(), nil
}

// RetrieveInProgressAttestation retrieves a in-progress attestation by <origin, destination, nonce>
func (s Store) RetrieveInProgressAttestation(ctx context.Context, originID, destinationID, nonce uint32) (attestation types.InProgressAttestation, err error) {
	var inProgressAttestation InProgressAttestation
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{IPOrigin: originID, IPDestination: destinationID, IPNonce: nonce}).
		First(&inProgressAttestation)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("could not find attestation with nonce %d and origin %d and destination %d: %w", nonce, originID, destinationID, db.ErrNotFound)
		}
		return nil, fmt.Errorf("could not retrieve attestation: %w", tx.Error)
	}
	return inProgressAttestation, err
}

// RetrieveOldestUnsignedInProgressAttestation retrieves the oldest in-progress attestation that has not yet been signed
func (s Store) RetrieveOldestUnsignedInProgressAttestation(ctx context.Context, originID, destinationID uint32) (_ types.InProgressAttestation, err error) {
	signatureNull := fmt.Sprintf("`%s` = '' or `%s` IS NULL", "signature", "signature")
	orderByNonceAsc := fmt.Sprintf("`%s` asc", NonceFieldName)
	var inProgressAttestation InProgressAttestation
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{IPOrigin: originID, IPDestination: destinationID}).
		Where(signatureNull).
		Order(orderByNonceAsc).
		First(&inProgressAttestation)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("could not find unsigned attestation with origin %d and destination %d: %w", originID, destinationID, db.ErrNotFound)
		}
		return nil, fmt.Errorf("could not retrieve attestation: %w", tx.Error)
	}
	return inProgressAttestation, err
}

// RetrieveOldestUnsubmittedSignedInProgressAttestation retrieves the oldest in-progress attestation that has been signed but not yet submitted
func (s Store) RetrieveOldestUnsubmittedSignedInProgressAttestation(ctx context.Context, originID, destinationID uint32) (_ types.InProgressAttestation, err error) {
	orderByNonceAsc := fmt.Sprintf("`%s` asc", NonceFieldName)
	signatureNotNull := fmt.Sprintf("`%s` <> '' and `%s` IS NOT NULL", "signature", "signature")
	submittedTimeIsNull := fmt.Sprintf("`%s` IS NULL", "submitted_to_attestation_collector_time")

	var inProgressAttestation InProgressAttestation
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{IPOrigin: originID, IPDestination: destinationID}).
		Where(signatureNotNull).
		Where(submittedTimeIsNull).
		Order(orderByNonceAsc).
		First(&inProgressAttestation)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("could not find signed attestation waiting to be submitted with origin %d and destination %d: %w", originID, destinationID, db.ErrNotFound)
		}
		return nil, fmt.Errorf("could not retrieve attestation: %w", tx.Error)
	}
	return inProgressAttestation, err
}

// RetrieveOldestUnconfirmedSubmittedInProgressAttestation retrieves the oldest in-progress attestation that has been signed and submitted but not yet confirmed on the AttestationCollector
func (s Store) RetrieveOldestUnconfirmedSubmittedInProgressAttestation(ctx context.Context, originID, destinationID uint32) (_ types.InProgressAttestation, err error) {
	orderByNonceAsc := fmt.Sprintf("`%s` asc", NonceFieldName)
	signatureNotNull := fmt.Sprintf("`%s` <> '' and `%s` IS NOT NULL", "signature", "signature")
	submittedTimeIsNotNull := fmt.Sprintf("`%s` IS NOT NULL", "submitted_to_attestation_collector_time")
	confirmationBlockNumberIsZero := fmt.Sprintf("`%s` = 0", "confirmed_on_attestation_collector_block_number")

	var inProgressAttestation InProgressAttestation
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{IPOrigin: originID, IPDestination: destinationID}).
		Where(signatureNotNull).
		Where(submittedTimeIsNotNull).
		Where(confirmationBlockNumberIsZero).
		Order(orderByNonceAsc).
		First(&inProgressAttestation)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("could not find submitted attestation waiting to be confirmed with origin %d and destination %d: %w", originID, destinationID, db.ErrNotFound)
		}
		return nil, fmt.Errorf("could not retrieve attestation: %w", tx.Error)
	}
	return inProgressAttestation, err
}
