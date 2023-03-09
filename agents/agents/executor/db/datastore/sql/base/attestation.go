package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
	"gorm.io/gorm/clause"
	"math/big"
)

// StoreAttestation stores an attestation.
func (s Store) StoreAttestation(ctx context.Context, attestation agentsTypes.Attestation, destination uint32, destinationBlockNumber, destinationTimestamp uint64) error {
	dbAttestation := agentsTypesAttestationToAttestation(attestation, destination, destinationBlockNumber, destinationTimestamp)

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: SnapshotRootFieldName}, {Name: AttestationNonceFieldName},
			},
			DoNothing: true,
		}).
		Create(&dbAttestation)

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store attestation: %w", dbTx.Error)
	}

	return nil
}

// GetAttestation gets an attestation that has fields matching the attestation mask.
func (s Store) GetAttestation(ctx context.Context, attestationMask types.DBAttestation) (*agentsTypes.Attestation, error) {
	var attestation Attestation

	dbAttestationMask := DBAttestationToAttestation(attestationMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&attestation).
		Where(&dbAttestationMask).
		Scan(&attestation)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get attestation: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil
	}

	agentsAttestation := agentsTypes.NewAttestation(
		common.HexToHash(attestation.SnapshotRoot),
		attestation.Height,
		attestation.AttestationNonce,
		big.NewInt(int64(attestation.SummitBlockNumber)),
		big.NewInt(int64(attestation.SummitTimestamp)),
	)

	return &agentsAttestation, nil
}

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

// GetAttestationTimestamp gets the timestamp of an attestation.
func (s Store) GetAttestationTimestamp(ctx context.Context, attestationMask types.DBAttestation) (*uint64, error) {
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

	return &attestation.DestinationTimestamp, nil
}

// GetAttestationMinimumTimestamp takes a list of snapshot roots and returns the timestamp of the attestation with the lowest block number.
func (s Store) GetAttestationMinimumTimestamp(ctx context.Context, attestationMask types.DBAttestation, snapshotRoots [][32]byte) (*uint64, error) {
	var attestation Attestation

	dbAttestationMask := DBAttestationToAttestation(attestationMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&attestation).
		Where(&dbAttestationMask).
		Where(fmt.Sprintf("%s IN ?", SnapshotRootFieldName), snapshotRoots).
		Order(fmt.Sprintf("%s ASC", DestinationBlockNumberFieldName)).
		Limit(1).
		Scan(&attestation)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get attestation minimum timestamp: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil
	}

	return &attestation.DestinationTimestamp, nil
}

// GetEarliestSnapshotFromAttestation takes a list of snapshot roots, checks which one has the lowest block number, and returns that snapshot root back.
func (s Store) GetEarliestSnapshotFromAttestation(ctx context.Context, attestationMask types.DBAttestation, snapshotRoots [][32]byte) (*[32]byte, error) {
	var attestation Attestation

	dbAttestationMask := DBAttestationToAttestation(attestationMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&attestation).
		Where(&dbAttestationMask).
		Where(fmt.Sprintf("%s IN ?", SnapshotRootFieldName), snapshotRoots).
		Order(fmt.Sprintf("%s ASC", DestinationBlockNumberFieldName)).
		Limit(1).
		Scan(&attestation)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get earliest attestation nonce: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil
	}

	snapshotRoot := common.HexToHash(attestation.SnapshotRoot)

	return (*[32]byte)(&snapshotRoot), nil
}

// DBAttestationToAttestation converts a DBAttestation to an Attestation.
func DBAttestationToAttestation(dbAttestation types.DBAttestation) Attestation {
	var attestation Attestation

	if dbAttestation.Destination != nil {
		attestation.Destination = *dbAttestation.Destination
	}

	if dbAttestation.SnapshotRoot != nil {
		attestation.SnapshotRoot = *dbAttestation.SnapshotRoot
	}

	if dbAttestation.Height != nil {
		attestation.Height = *dbAttestation.Height
	}

	if dbAttestation.AttestationNonce != nil {
		attestation.AttestationNonce = *dbAttestation.AttestationNonce
	}

	if dbAttestation.SummitBlockNumber != nil {
		attestation.SummitBlockNumber = *dbAttestation.SummitBlockNumber
	}

	if dbAttestation.SummitTimestamp != nil {
		attestation.SummitTimestamp = *dbAttestation.SummitTimestamp
	}

	if dbAttestation.DestinationBlockNumber != nil {
		attestation.DestinationBlockNumber = *dbAttestation.DestinationBlockNumber
	}

	if dbAttestation.DestinationTimestamp != nil {
		attestation.DestinationTimestamp = *dbAttestation.DestinationTimestamp
	}

	return attestation
}

// AttestationToDBAttestation converts an Attestation to a DBAttestation.
func AttestationToDBAttestation(attestation Attestation) types.DBAttestation {
	destination := attestation.Destination
	snapshotRoot := attestation.SnapshotRoot
	height := attestation.Height
	attestationNonce := attestation.AttestationNonce
	summitBlockNumber := attestation.SummitBlockNumber
	summitTimestamp := attestation.SummitTimestamp
	destinationBlockNumber := attestation.DestinationBlockNumber
	destinationBlockTime := attestation.DestinationTimestamp

	return types.DBAttestation{
		Destination:            &destination,
		SnapshotRoot:           &snapshotRoot,
		Height:                 &height,
		AttestationNonce:       &attestationNonce,
		SummitBlockNumber:      &summitBlockNumber,
		SummitTimestamp:        &summitTimestamp,
		DestinationBlockNumber: &destinationBlockNumber,
		DestinationTimestamp:   &destinationBlockTime,
	}
}

// agentsTypesAttestationToAttestation converts an agentsTypes.Attestation to an Attestation.
func agentsTypesAttestationToAttestation(attestation agentsTypes.Attestation, destination uint32, destinationBlockNumber, destinationTimestamp uint64) Attestation {
	snapshotRoot := attestation.SnapshotRoot()

	return Attestation{
		Destination:            destination,
		SnapshotRoot:           common.BytesToHash(snapshotRoot[:]).String(),
		Height:                 attestation.Height(),
		AttestationNonce:       attestation.Nonce(),
		SummitBlockNumber:      attestation.BlockNumber().Uint64(),
		SummitTimestamp:        attestation.Timestamp().Uint64(),
		DestinationBlockNumber: destinationBlockNumber,
		DestinationTimestamp:   destinationTimestamp,
	}
}

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
//
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
//
//// GetAttestationBlockNumber gets the block number of an attestation.
//func (s Store) GetAttestationBlockNumber(ctx context.Context, attestationMask types.DBAttestation) (*uint64, error) {
//	var attestation Attestation
//
//	dbAttestationMask := DBAttestationToAttestation(attestationMask)
//	dbTx := s.DB().WithContext(ctx).
//		Model(&attestation).
//		Where(&dbAttestationMask).
//		Scan(&attestation)
//	if dbTx.Error != nil {
//		return nil, fmt.Errorf("failed to get attestation block number: %w", dbTx.Error)
//	}
//	if dbTx.RowsAffected == 0 {
//		//nolint:nilnil
//		return nil, nil
//	}
//
//	return &attestation.DestinationBlockNumber, nil
//}
//
//// GetAttestationBlockTime gets the block time of an attestation.
//func (s Store) GetAttestationBlockTime(ctx context.Context, attestationMask types.DBAttestation) (*uint64, error) {
//	var attestation Attestation
//
//	dbAttestationMask := DBAttestationToAttestation(attestationMask)
//	dbTx := s.DB().WithContext(ctx).
//		Model(&attestation).
//		Where(&dbAttestationMask).
//		Scan(&attestation)
//	if dbTx.Error != nil {
//		return nil, fmt.Errorf("failed to get attestation block time: %w", dbTx.Error)
//	}
//	if dbTx.RowsAffected == 0 {
//		//nolint:nilnil
//		return nil, nil
//	}
//
//	return &attestation.DestinationTimestamp, nil
//}
//
//// GetAttestationForNonceOrGreater gets the lowest nonce attestation that is greater than or equal to the given nonce.
//func (s Store) GetAttestationForNonceOrGreater(ctx context.Context, attestationMask types.DBAttestation) (nonce *uint32, blockTime *uint64, err error) {
//	var attestation Attestation
//
//	dbAttestationMask := DBAttestationToAttestation(attestationMask)
//	dbTx := s.DB().WithContext(ctx).
//		Model(&attestation).
//		Where(&dbAttestationMask).
//		Where(fmt.Sprintf("%s >= ?", NonceFieldName), attestationMask.Nonce).
//		Order(fmt.Sprintf("%s ASC", DestinationBlockNumberFieldName)).
//		Limit(1).
//		Scan(&attestation)
//	if dbTx.Error != nil {
//		return nil, nil, fmt.Errorf("failed to get attestation for nonce or greater: %w", dbTx.Error)
//	}
//	if dbTx.RowsAffected == 0 {
//		return nil, nil, nil
//	}
//
//	nonce = &attestation.Nonce
//	blockTime = &attestation.DestinationTimestamp
//
//	return nonce, blockTime, nil
//}
//
//// GetAttestationsAboveOrEqualNonce gets attestations in a nonce range.
//func (s Store) GetAttestationsAboveOrEqualNonce(ctx context.Context, attestationMask types.DBAttestation, minNonce uint32, page int) ([]types.DBAttestation, error) {
//	if page < 1 {
//		page = 1
//	}
//
//	var attestations []Attestation
//
//	dbAttestationMask := DBAttestationToAttestation(attestationMask)
//	dbTx := s.DB().WithContext(ctx).
//		Model(&attestations).
//		Where(&dbAttestationMask).
//		Where(fmt.Sprintf("%s >= ?", NonceFieldName), minNonce).
//		Order(fmt.Sprintf("%s ASC", DestinationBlockNumberFieldName)).
//		Offset((page - 1) * PageSize).
//		Limit(PageSize).
//		Scan(&attestations)
//	if dbTx.Error != nil {
//		return nil, fmt.Errorf("failed to get attestations in nonce range: %w", dbTx.Error)
//	}
//
//	dbAttestations := make([]types.DBAttestation, len(attestations))
//	for i := range attestations {
//		dbAttestations[i] = AttestationToDBAttestation(attestations[i])
//	}
//
//	return dbAttestations, nil
//}
//
//// GetEarliestAttestationsNonceInNonceRange gets the earliest attestation (by block number) in a nonce range.
//func (s Store) GetEarliestAttestationsNonceInNonceRange(ctx context.Context, attestationMask types.DBAttestation, minNonce uint32, maxNonce uint32) (*uint32, error) {
//	var attestation Attestation
//
//	dbAttestationMask := DBAttestationToAttestation(attestationMask)
//	dbTx := s.DB().WithContext(ctx).
//		Model(&attestation).
//		Where(&dbAttestationMask).
//		Where(fmt.Sprintf("%s >= ?", NonceFieldName), minNonce).
//		Where(fmt.Sprintf("%s <= ?", NonceFieldName), maxNonce).
//		Order(fmt.Sprintf("%s ASC", DestinationBlockNumberFieldName)).
//		Limit(1).
//		Scan(&attestation)
//	if dbTx.Error != nil {
//		return nil, fmt.Errorf("failed to get earliest attestation in nonce range: %w", dbTx.Error)
//	}
//	if dbTx.RowsAffected == 0 {
//		//nolint:nilnil
//		return nil, nil
//	}
//
//	return &attestation.Nonce, nil
//}
//
//// DBAttestationToAttestation converts a DBAttestation to an Attestation.
//func DBAttestationToAttestation(dbAttestation types.DBAttestation) Attestation {
//	var attestation Attestation
//
//	if dbAttestation.ChainID != nil {
//		attestation.ChainID = *dbAttestation.ChainID
//	}
//
//	if dbAttestation.Destination != nil {
//		attestation.Destination = *dbAttestation.Destination
//	}
//
//	if dbAttestation.Nonce != nil {
//		attestation.Nonce = *dbAttestation.Nonce
//	}
//
//	if dbAttestation.Root != nil {
//		attestation.Root = dbAttestation.Root.String()
//	}
//
//	if dbAttestation.DestinationBlockNumber != nil {
//		attestation.DestinationBlockNumber = *dbAttestation.DestinationBlockNumber
//	}
//
//	if dbAttestation.DestinationTimestamp != nil {
//		attestation.DestinationTimestamp = *dbAttestation.DestinationTimestamp
//	}
//
//	return attestation
//}
//
//// AttestationToDBAttestation converts an Attestation to a DBAttestation.
//func AttestationToDBAttestation(attestation Attestation) types.DBAttestation {
//	chainID := attestation.ChainID
//	destination := attestation.Destination
//	nonce := attestation.Nonce
//	root := common.HexToHash(attestation.Root)
//	blockNumber := attestation.DestinationBlockNumber
//	blockTime := attestation.DestinationTimestamp
//
//	return types.DBAttestation{
//		ChainID:                &chainID,
//		Destination:            &destination,
//		Nonce:                  &nonce,
//		Root:                   &root,
//		DestinationBlockNumber: &blockNumber,
//		DestinationTimestamp:   &blockTime,
//	}
//}
//
//// agentsTypesAttestationToAttestation converts an agentsTypes.Attestation to an Attestation.
//func agentsTypesAttestationToAttestation(attestation agentsTypes.Attestation, blockNumber uint64, blockTime uint64) Attestation {
//	root := attestation.Root()
//	return Attestation{
//		ChainID:                attestation.Origin(),
//		Destination:            attestation.Destination(),
//		Nonce:                  attestation.Nonce(),
//		Root:                   common.BytesToHash(root[:]).String(),
//		DestinationBlockNumber: blockNumber,
//		DestinationTimestamp:   blockTime,
//	}
//}
