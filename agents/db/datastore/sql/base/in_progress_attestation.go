package base

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Thor-x86/nullable"
	"github.com/synapsecns/sanguine/agents/db"

	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func getOrderByNonceAsc() string {
	return fmt.Sprintf("`%s` asc", NonceFieldName)
}

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
		IPAttestationState:          uint32(types.AttestationStateNotaryUnsigned),
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

	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where("attestation_state", uint32(types.AttestationStateNotaryUnsigned)).
		Updates(
			InProgressAttestation{
				IPSignature:        sig,
				IPAttestationState: uint32(types.AttestationStateNotarySignedUnsubmitted),
			},
		)

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

	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where("attestation_state", uint32(types.AttestationStateNotarySignedUnsubmitted)).
		Updates(
			InProgressAttestation{
				IPSubmittedToAttestationCollectorTime: sql.NullTime{
					Time:  *inProgressAttestation.SubmittedToAttestationCollectorTime(),
					Valid: true,
				},
				IPAttestationState: uint32(types.AttestationStateNotarySubmittedUnconfirmed),
			},
		)

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

	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where("attestation_state", uint32(types.AttestationStateNotarySubmittedUnconfirmed)).
		Updates(
			InProgressAttestation{
				IPConfirmedOnAttestationCollectorBlockNumber: inProgressAttestation.ConfirmedOnAttestationCollectorBlockNumber(),
				IPAttestationState:                           uint32(types.AttestationStateNotaryConfirmed),
			},
		)

	if tx.Error != nil {
		return fmt.Errorf("could not set ConfirmedOnAttestationCollectorBlockNumber for in-progress attestation: %w", tx.Error)
	}
	return nil
}

// RetrieveLatestCachedNonce retrieves the latest nonce cached for given origin-destination pair.
// TODO (joe): Currently, we are grabbing ALL the nonces rather than just asking the origin for the
// most recent one. We are calling getHistoricalRoot with the next nonce that we haven't seen.
// Later, we will replace this with calling suggestAttestation for our particular destination
// and only getting the latest one.
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
// This is mainly used for testing
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
// TODO (joe): Eventually we will not try to sign ALL the nonces, we really just want the latest one, so we will
// want to replace this with RetrieveNewestUnsignedInProgressAttestation. For Notary MVP, we want to sign all the nonces though
// so we will just get the oldest and go in order
func (s Store) RetrieveOldestUnsignedInProgressAttestation(ctx context.Context, originID, destinationID uint32) (_ types.InProgressAttestation, err error) {
	var inProgressAttestation InProgressAttestation
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      originID,
			IPDestination: destinationID,
		}).
		Where("attestation_state", uint32(types.AttestationStateNotaryUnsigned)).
		Order(getOrderByNonceAsc()).
		First(&inProgressAttestation)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("could not find unsigned attestation with origin %d and destination %d: %w", originID, destinationID, db.ErrNotFound)
		}
		return nil, fmt.Errorf("could not retrieve attestation: %w", tx.Error)
	}
	return inProgressAttestation, err
}

// RetrieveOldestUnsubmittedSignedInProgressAttestation retrieves the oldest in-progress attestation that has been signed but not yet submitted.
func (s Store) RetrieveOldestUnsubmittedSignedInProgressAttestation(ctx context.Context, originID, destinationID uint32) (_ types.InProgressAttestation, err error) {
	var inProgressAttestation InProgressAttestation
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      originID,
			IPDestination: destinationID,
		}).
		Where("attestation_state", uint32(types.AttestationStateNotarySignedUnsubmitted)).
		Order(getOrderByNonceAsc()).
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

	var inProgressAttestation InProgressAttestation
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      originID,
			IPDestination: destinationID,
		}).
		Where("attestation_state", uint32(types.AttestationStateNotarySubmittedUnconfirmed)).
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
