package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
)

//// StoreAttestation stores an attestation.
//func (s Store) StoreAttestation(ctx context.Context, attestation agentsTypes.Attestation, blockNumber uint64, blockTime uint64) error {
//	dbAttestation := agentsTypesAttestationToAttestation(attestation, blockNumber, blockTime)
//
//	dbTx := s.DB().WithContext(ctx).
//		Clauses(clause.OnConflict{
//			Columns: []clause.Column{
//				{Name: ChainIDFieldName}, {Name: DestinationFieldName}, {Name: NonceFieldName}, {Name: RootFieldName},
//			},
//			DoNothing: true,
//		}).
//		Create(&dbAttestation)
//
//	if dbTx.Error != nil {
//		return fmt.Errorf("failed to store attestation: %w", dbTx.Error)
//	}
//
//	return nil
//}

//// GetAttestation gets an attestation from the database.
//func (s Store) GetAttestation(ctx context.Context, attestationMask types.DBAttestation) (*agentsTypes.Attestation, error) {
//	var attestation Attestation
//
//	dbAttestationMask := DBAttestationToAttestation(attestationMask)
//	dbTx := s.DB().WithContext(ctx).
//		Model(&attestation).
//		Where(&dbAttestationMask).
//		Limit(1).
//		Scan(&attestation)
//	if dbTx.Error != nil {
//		return nil, fmt.Errorf("failed to get attestation: %w", dbTx.Error)
//	}
//	if dbTx.RowsAffected == 0 {
//		//nolint:nilnil
//		return nil, nil
//	}
//
//	attestKey := agentsTypes.AttestationKey{
//		Origin:      attestation.ChainID,
//		Destination: attestation.Destination,
//		Nonce:       attestation.Nonce,
//	}
//
//	receivedAttestation := agentsTypes.NewAttestation(attestKey.GetRawKey(), common.HexToHash(attestation.Root))
//
//	return &receivedAttestation, nil
//}

// GetAttestationBlockNumber gets the block number of an attestation.
func (s Store) GetAttestationBlockNumber(ctx context.Context, attestationMask types.DBAttestation) (*uint64, error) {
	var attestation Attestation

	dbAttestationMask := DBAttestationToAttestation(attestationMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&attestation).
		Where(&dbAttestationMask).
		Scan(&attestation)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get attestation block number: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil
	}

	return &attestation.DestinationBlockNumber, nil
}

// GetAttestationBlockTime gets the block time of an attestation.
func (s Store) GetAttestationBlockTime(ctx context.Context, attestationMask types.DBAttestation) (*uint64, error) {
	var attestation Attestation

	dbAttestationMask := DBAttestationToAttestation(attestationMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&attestation).
		Where(&dbAttestationMask).
		Scan(&attestation)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get attestation block time: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil
	}

	return &attestation.DestinationBlockTime, nil
}

// GetAttestationForNonceOrGreater gets the lowest nonce attestation that is greater than or equal to the given nonce.
func (s Store) GetAttestationForNonceOrGreater(ctx context.Context, attestationMask types.DBAttestation) (nonce *uint32, blockTime *uint64, err error) {
	var attestation Attestation

	dbAttestationMask := DBAttestationToAttestation(attestationMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&attestation).
		Where(&dbAttestationMask).
		Where(fmt.Sprintf("%s >= ?", NonceFieldName), attestationMask.Nonce).
		Order(fmt.Sprintf("%s ASC", BlockNumberFieldName)).
		Limit(1).
		Scan(&attestation)
	if dbTx.Error != nil {
		return nil, nil, fmt.Errorf("failed to get attestation for nonce or greater: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		return nil, nil, nil
	}

	nonce = &attestation.Nonce
	blockTime = &attestation.DestinationBlockTime

	return nonce, blockTime, nil
}

// GetAttestationsAboveOrEqualNonce gets attestations in a nonce range.
func (s Store) GetAttestationsAboveOrEqualNonce(ctx context.Context, attestationMask types.DBAttestation, minNonce uint32, page int) ([]types.DBAttestation, error) {
	if page < 1 {
		page = 1
	}

	var attestations []Attestation

	dbAttestationMask := DBAttestationToAttestation(attestationMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&attestations).
		Where(&dbAttestationMask).
		Where(fmt.Sprintf("%s >= ?", NonceFieldName), minNonce).
		Order(fmt.Sprintf("%s ASC", BlockNumberFieldName)).
		Offset((page - 1) * PageSize).
		Limit(PageSize).
		Scan(&attestations)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get attestations in nonce range: %w", dbTx.Error)
	}

	dbAttestations := make([]types.DBAttestation, len(attestations))
	for i := range attestations {
		dbAttestations[i] = AttestationToDBAttestation(attestations[i])
	}

	return dbAttestations, nil
}

// GetEarliestAttestationsNonceInNonceRange gets the earliest attestation (by block number) in a nonce range.
func (s Store) GetEarliestAttestationsNonceInNonceRange(ctx context.Context, attestationMask types.DBAttestation, minNonce uint32, maxNonce uint32) (*uint32, error) {
	var attestation Attestation

	dbAttestationMask := DBAttestationToAttestation(attestationMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&attestation).
		Where(&dbAttestationMask).
		Where(fmt.Sprintf("%s >= ?", NonceFieldName), minNonce).
		Where(fmt.Sprintf("%s <= ?", NonceFieldName), maxNonce).
		Order(fmt.Sprintf("%s ASC", BlockNumberFieldName)).
		Limit(1).
		Scan(&attestation)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get earliest attestation in nonce range: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil
	}

	return &attestation.Nonce, nil
}

// DBAttestationToAttestation converts a DBAttestation to an Attestation.
func DBAttestationToAttestation(dbAttestation types.DBAttestation) Attestation {
	var attestation Attestation

	if dbAttestation.ChainID != nil {
		attestation.ChainID = *dbAttestation.ChainID
	}

	if dbAttestation.Destination != nil {
		attestation.Destination = *dbAttestation.Destination
	}

	if dbAttestation.Nonce != nil {
		attestation.Nonce = *dbAttestation.Nonce
	}

	if dbAttestation.Root != nil {
		attestation.Root = dbAttestation.Root.String()
	}

	if dbAttestation.DestinationBlockNumber != nil {
		attestation.DestinationBlockNumber = *dbAttestation.DestinationBlockNumber
	}

	if dbAttestation.DestinationBlockTime != nil {
		attestation.DestinationBlockTime = *dbAttestation.DestinationBlockTime
	}

	return attestation
}

// AttestationToDBAttestation converts an Attestation to a DBAttestation.
func AttestationToDBAttestation(attestation Attestation) types.DBAttestation {
	chainID := attestation.ChainID
	destination := attestation.Destination
	nonce := attestation.Nonce
	root := common.HexToHash(attestation.Root)
	blockNumber := attestation.DestinationBlockNumber
	blockTime := attestation.DestinationBlockTime

	return types.DBAttestation{
		ChainID:                &chainID,
		Destination:            &destination,
		Nonce:                  &nonce,
		Root:                   &root,
		DestinationBlockNumber: &blockNumber,
		DestinationBlockTime:   &blockTime,
	}
}

//// agentsTypesAttestationToAttestation converts an agentsTypes.Attestation to an Attestation.
//func agentsTypesAttestationToAttestation(attestation agentsTypes.Attestation, blockNumber uint64, blockTime uint64) Attestation {
//	root := attestation.Root()
//	return Attestation{
//		ChainID:                attestation.Origin(),
//		Destination:            attestation.Destination(),
//		Nonce:                  attestation.Nonce(),
//		Root:                   common.BytesToHash(root[:]).String(),
//		DestinationBlockNumber: blockNumber,
//		DestinationBlockTime:   blockTime,
//	}
//}
