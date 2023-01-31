package base

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Thor-x86/nullable"
	"github.com/synapsecns/sanguine/agents/db"

	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func getOrderByNonceDesc() string {
	return fmt.Sprintf("`%s` desc", NonceFieldName)
}

// StoreNewInProgressAttestation stores in-progress attestation only if it hasn't already been stored.
func (s Store) StoreNewInProgressAttestation(ctx context.Context, attestation types.Attestation) error {
	// We only want to store if not already stored
	tx := s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: OriginFieldName}, {Name: DestinationFieldName}, {Name: NonceFieldName}},
		DoNothing: true,
	}).Create(&InProgressAttestation{
		IPOrigin:           attestation.Origin(),
		IPDestination:      attestation.Destination(),
		IPNonce:            attestation.Nonce(),
		IPRoot:             core.BytesToSlice(attestation.Root()),
		IPAttestationState: uint32(types.AttestationStateNotaryUnsigned),
	})

	if tx.Error != nil {
		return fmt.Errorf("could not store signed attestations: %w", tx.Error)
	}
	return nil
}

// StoreNewGuardInProgressAttestation stores in-progress attestation only if it hasn't already been stored.
func (s Store) StoreNewGuardInProgressAttestation(ctx context.Context, attestation types.Attestation) error {
	// We only want to store if not already stored
	tx := s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: OriginFieldName}, {Name: DestinationFieldName}, {Name: NonceFieldName}},
		DoNothing: true,
	}).Create(&InProgressAttestation{
		IPOrigin:           attestation.Origin(),
		IPDestination:      attestation.Destination(),
		IPNonce:            attestation.Nonce(),
		IPRoot:             core.BytesToSlice(attestation.Root()),
		IPAttestationState: uint32(types.AttestationStateGuardInitialState),
	})

	if tx.Error != nil {
		return fmt.Errorf("could not store signed attestations: %w", tx.Error)
	}
	return nil
}

// StoreExistingSignedInProgressAttestation stores signed in-progress attestation only if it hasn't already been stored.
func (s Store) StoreExistingSignedInProgressAttestation(ctx context.Context, signedAttestation types.SignedAttestation) error {
	if len(signedAttestation.NotarySignatures()) == 0 {
		return fmt.Errorf("StoreExistingSignedInProgressAttestation called on signedAttestation with no notary signatures")
	}
	sig, err := types.EncodeSignature(signedAttestation.NotarySignatures()[0])
	if err != nil {
		return fmt.Errorf("could not encode notary signature: %w", err)
	}

	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      signedAttestation.Attestation().Origin(),
			IPDestination: signedAttestation.Attestation().Destination(),
			IPNonce:       signedAttestation.Attestation().Nonce(),
		}).
		Where(AttestationStateFieldName, uint32(types.AttestationStateGuardInitialState)).
		Updates(
			InProgressAttestation{
				IPNotarySignature:  sig,
				IPAttestationState: uint32(types.AttestationStateGuardUnsignedAndVerified),
			},
		)

	if tx.Error != nil {
		return fmt.Errorf("could not set notary signature for in-progress attestations: %w", tx.Error)
	}
	return nil
}

// UpdateNotarySignature sets the notary signature of the in-progress Attestation.
//
//nolint:dupl
func (s Store) UpdateNotarySignature(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	if len(inProgressAttestation.SignedAttestation().NotarySignatures()) == 0 {
		return fmt.Errorf("UpdateNotarySignature called on attestation with a nil notary signature")
	}
	sig, err := types.EncodeSignature(inProgressAttestation.SignedAttestation().NotarySignatures()[0])
	if err != nil {
		return fmt.Errorf("could not encode notary signature: %w", err)
	}

	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(AttestationStateFieldName, uint32(types.AttestationStateNotaryUnsigned)).
		Updates(
			InProgressAttestation{
				IPNotarySignature:  sig,
				IPAttestationState: uint32(types.AttestationStateNotarySignedUnsubmitted),
			},
		)

	if tx.Error != nil {
		return fmt.Errorf("could not set notary signature for in-progress attestations: %w", tx.Error)
	}
	return nil
}

// UpdateNotarySubmittedToAttestationCollectorTime sets the time attestation was sent to Attestation Collector by the Notary.
func (s Store) UpdateNotarySubmittedToAttestationCollectorTime(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	if inProgressAttestation.SubmittedToAttestationCollectorTime() == nil {
		return fmt.Errorf("UpdateNotarySubmittedToAttestationCollectorTime called on attestation with a nil time")
	}

	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(AttestationStateFieldName, uint32(types.AttestationStateNotarySignedUnsubmitted)).
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

// ReUpdateNotarySubmittedToAttestationCollectorTime sets the time attestation was sent to Attestation Collector by the Notary when resubmitting.
func (s Store) ReUpdateNotarySubmittedToAttestationCollectorTime(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	if inProgressAttestation.SubmittedToAttestationCollectorTime() == nil {
		return fmt.Errorf("UpdateNotarySubmittedToAttestationCollectorTime called on attestation with a nil time")
	}

	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(AttestationStateFieldName, uint32(types.AttestationStateNotarySubmittedUnconfirmed)).
		Updates(
			InProgressAttestation{
				IPSubmittedToAttestationCollectorTime: sql.NullTime{
					Time:  *inProgressAttestation.SubmittedToAttestationCollectorTime(),
					Valid: true,
				},
			},
		)

	if tx.Error != nil {
		return fmt.Errorf("could not update SubmittedToAttestationCollectorTime for in-progress attestations: %w", tx.Error)
	}
	return nil
}

// MarkNotaryConfirmedOnAttestationCollector confirms that the Notary posted the signed attestation on the Attestation Collector.
func (s Store) MarkNotaryConfirmedOnAttestationCollector(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(AttestationStateFieldName, uint32(types.AttestationStateNotarySubmittedUnconfirmed)).
		Updates(
			InProgressAttestation{
				IPAttestationState: uint32(types.AttestationStateNotaryConfirmed),
			},
		)

	if tx.Error != nil {
		return fmt.Errorf("could not execute MarkNotaryConfirmedOnAttestationCollector for in-progress attestation: %w", tx.Error)
	}
	return nil
}

// UpdateGuardSignature sets the guard signature of the in-progress Attestation.
//
//nolint:dupl
func (s Store) UpdateGuardSignature(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	if len(inProgressAttestation.SignedAttestation().GuardSignatures()) == 0 {
		return fmt.Errorf("UpdateGuardSignature called on attestation with a nil guard signature")
	}
	sig, err := types.EncodeSignature(inProgressAttestation.SignedAttestation().GuardSignatures()[0])
	if err != nil {
		return fmt.Errorf("could not encode guard signature: %w", err)
	}

	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(AttestationStateFieldName, uint32(types.AttestationStateGuardUnsignedAndVerified)).
		Updates(
			InProgressAttestation{
				IPGuardSignature:   sig,
				IPAttestationState: uint32(types.AttestationStateGuardSignedUnsubmitted),
			},
		)

	if tx.Error != nil {
		return fmt.Errorf("could not set guard signature for in-progress attestations: %w", tx.Error)
	}
	return nil
}

// RetrieveLatestCachedNonce retrieves the latest nonce cached for given origin-destination pair.
// TODO (joe): Currently, we are grabbing ALL the nonces rather than just asking the origin for the
// most recent one. We are calling getHistoricalRoot with the next nonce that we haven't seen.
// Later, we will replace this with calling suggestAttestation for our particular destination
// and only getting the latest one.
func (s Store) RetrieveLatestCachedNonce(ctx context.Context, originID, destinationID uint32) (_ uint32, err error) {
	if originID == uint32(0) {
		return uint32(0), fmt.Errorf("RetrieveLatestCachedNonce called with 0 origin")
	}
	if destinationID == uint32(0) {
		return uint32(0), fmt.Errorf("RetrieveLatestCachedNonce called with 0 destination")
	}
	var nonce nullable.Uint32

	selectMaxNonce := fmt.Sprintf("max(`%s`)", NonceFieldName)

	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Select(selectMaxNonce).
		Where(&InProgressAttestation{IPOrigin: originID, IPDestination: destinationID}).
		Scan(&nonce)

	if tx.Error != nil {
		return 0, fmt.Errorf("could not get nonce for origin %d and destiniation %d: %w", originID, destinationID, tx.Error)
	}

	// if no nonces, return the corresponding error.
	if nonce.Get() == nil {
		return 0, db.ErrNoNonceForDomain
	}
	return *nonce.Get(), nil
}

// RetrieveInProgressAttestation retrieves a in-progress attestation by <origin, destination, nonce>.
// This is mainly used for testing.
func (s Store) RetrieveInProgressAttestation(ctx context.Context, originID, destinationID, nonce uint32) (attestation types.InProgressAttestation, err error) {
	if originID == uint32(0) {
		return nil, fmt.Errorf("RetrieveInProgressAttestation called with 0 origin")
	}
	if destinationID == uint32(0) {
		return nil, fmt.Errorf("RetrieveInProgressAttestation called with 0 destination")
	}
	if nonce == uint32(0) {
		return nil, fmt.Errorf("RetrieveInProgressAttestation called with 0 nonce")
	}
	var inProgressAttestation InProgressAttestation
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{IPOrigin: originID, IPDestination: destinationID, IPNonce: nonce}).
		First(&inProgressAttestation)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, db.ErrNotFound
		}
		return nil, fmt.Errorf("could not retrieve attestation: %w", tx.Error)
	}
	return inProgressAttestation, err
}

// RetrieveNewestInProgressAttestation retrieves the newest in-progress attestation.
func (s Store) RetrieveNewestInProgressAttestation(ctx context.Context, originID, destinationID uint32) (_ types.InProgressAttestation, err error) {
	if originID == uint32(0) {
		return nil, fmt.Errorf("RetrieveNewestInProgressAttestation called with 0 origin")
	}
	if destinationID == uint32(0) {
		return nil, fmt.Errorf("RetrieveNewestInProgressAttestation called with 0 destination")
	}
	var inProgressAttestation InProgressAttestation
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(OriginFieldName, originID).
		Where(DestinationFieldName, destinationID).
		Order(getOrderByNonceDesc()).
		First(&inProgressAttestation)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, db.ErrNotFound
		}
		return nil, fmt.Errorf("could not retrieve attestation: %w", tx.Error)
	}
	return inProgressAttestation, err
}

// RetrieveNewestInProgressAttestationIfInState retrieves the newest in-progress attestation in the given state.
func (s Store) RetrieveNewestInProgressAttestationIfInState(ctx context.Context, originID, destinationID uint32, state types.AttestationState) (_ types.InProgressAttestation, err error) {
	inProgressAttestation, err := s.RetrieveNewestInProgressAttestation(ctx, originID, destinationID)
	if err != nil {
		return nil, err
	}

	if inProgressAttestation.AttestationState() != state {
		return nil, db.ErrNotFound
	}

	return inProgressAttestation, err
}

// UpdateGuardSubmittedToAttestationCollectorTime sets the time the attestation was sent to Attestation Collector by the Guard.
func (s Store) UpdateGuardSubmittedToAttestationCollectorTime(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	if inProgressAttestation.SubmittedToAttestationCollectorTime() == nil {
		return fmt.Errorf("UpdateGuardSubmittedToAttestationCollectorTime called on attestation with a nil time")
	}

	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(AttestationStateFieldName, uint32(types.AttestationStateGuardSignedUnsubmitted)).
		Updates(
			InProgressAttestation{
				IPSubmittedToAttestationCollectorTime: sql.NullTime{
					Time:  *inProgressAttestation.SubmittedToAttestationCollectorTime(),
					Valid: true,
				},
				IPAttestationState: uint32(types.AttestationStateGuardSubmittedToCollectorUnconfirmed),
			},
		)

	if tx.Error != nil {
		return fmt.Errorf("could not update SubmittedToAttestationCollectorTime for in-progress attestations: %w", tx.Error)
	}
	return nil
}

// ReUpdateGuardSubmittedToAttestationCollectorTime sets the time the attestation was sent to Attestation Collector by the Guard when resubmitting.
func (s Store) ReUpdateGuardSubmittedToAttestationCollectorTime(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	if inProgressAttestation.SubmittedToAttestationCollectorTime() == nil {
		return fmt.Errorf("UpdateGuardSubmittedToAttestationCollectorTime called on attestation with a nil time")
	}

	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(AttestationStateFieldName, uint32(types.AttestationStateGuardSubmittedToCollectorUnconfirmed)).
		Updates(
			InProgressAttestation{
				IPSubmittedToAttestationCollectorTime: sql.NullTime{
					Time:  *inProgressAttestation.SubmittedToAttestationCollectorTime(),
					Valid: true,
				},
			},
		)

	if tx.Error != nil {
		return fmt.Errorf("could not update SubmittedToAttestationCollectorTime for in-progress attestations: %w", tx.Error)
	}
	return nil
}

// MarkGuardConfirmedOnAttestationCollector confirms that the Guard posted the signed attestation on the Attestation Collector.
func (s Store) MarkGuardConfirmedOnAttestationCollector(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(AttestationStateFieldName, uint32(types.AttestationStateGuardSubmittedToCollectorUnconfirmed)).
		Updates(
			InProgressAttestation{
				IPAttestationState: uint32(types.AttestationStateGuardConfirmedOnCollector),
			},
		)

	if tx.Error != nil {
		return fmt.Errorf("could not execute MarkGuardConfirmedOnAttestationCollector for in-progress attestation: %w", tx.Error)
	}
	return nil
}

// UpdateSubmittedToDestinationTime sets the time the attestation was sent to the Destination.
func (s Store) UpdateSubmittedToDestinationTime(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	nowTime := time.Now()
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(AttestationStateFieldName, uint32(types.AttestationStateGuardConfirmedOnCollector)).
		Updates(
			InProgressAttestation{
				IPSubmittedToDestinationTime: sql.NullTime{
					Time:  nowTime,
					Valid: true,
				},
				IPAttestationState: uint32(types.AttestationStateSubmittedToDestinationUnconfirmed),
			},
		)

	if tx.Error != nil {
		return fmt.Errorf("could not update SubmittedToDestinationTime for in-progress attestations: %w", tx.Error)
	}
	return nil
}

// MarkConfirmedOnDestination confirms that we posted the signed attestation on the Destination.
func (s Store) MarkConfirmedOnDestination(ctx context.Context, inProgressAttestation types.InProgressAttestation) error {
	tx := s.DB().WithContext(ctx).Model(&InProgressAttestation{}).
		Where(&InProgressAttestation{
			IPOrigin:      inProgressAttestation.SignedAttestation().Attestation().Origin(),
			IPDestination: inProgressAttestation.SignedAttestation().Attestation().Destination(),
			IPNonce:       inProgressAttestation.SignedAttestation().Attestation().Nonce(),
		}).
		Where(AttestationStateFieldName, uint32(types.AttestationStateSubmittedToDestinationUnconfirmed)).
		Updates(
			InProgressAttestation{
				IPAttestationState: uint32(types.AttestationStateConfirmedOnDestination),
			},
		)

	if tx.Error != nil {
		return fmt.Errorf("could not execute MarkConfirmedOnDestination for in-progress attestation: %w", tx.Error)
	}
	return nil
}
