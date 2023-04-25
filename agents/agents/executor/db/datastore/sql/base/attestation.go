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
				{Name: DestinationFieldName}, {Name: SnapshotRootFieldName}, {Name: AttestationNonceFieldName},
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
		common.HexToHash(attestation.AgentRoot),
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

// GetEarliestSnapshotFromAttestation takes a list of snapshot roots, checks which one has the lowest block number, and returns that snapshot root back.
func (s Store) GetEarliestSnapshotFromAttestation(ctx context.Context, attestationMask types.DBAttestation, snapshotRoots []string) (*[32]byte, error) {
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

	if dbAttestation.AgentRoot != nil {
		attestation.AgentRoot = *dbAttestation.AgentRoot
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
	agentRoot := attestation.AgentRoot
	attestationNonce := attestation.AttestationNonce
	summitBlockNumber := attestation.SummitBlockNumber
	summitTimestamp := attestation.SummitTimestamp
	destinationBlockNumber := attestation.DestinationBlockNumber
	destinationBlockTime := attestation.DestinationTimestamp

	return types.DBAttestation{
		Destination:            &destination,
		SnapshotRoot:           &snapshotRoot,
		AgentRoot:              &agentRoot,
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
	agentRoot := attestation.AgentRoot()

	return Attestation{
		Destination:            destination,
		SnapshotRoot:           common.BytesToHash(snapshotRoot[:]).String(),
		AgentRoot:              common.BytesToHash(agentRoot[:]).String(),
		AttestationNonce:       attestation.Nonce(),
		SummitBlockNumber:      attestation.BlockNumber().Uint64(),
		SummitTimestamp:        attestation.Timestamp().Uint64(),
		DestinationBlockNumber: destinationBlockNumber,
		DestinationTimestamp:   destinationTimestamp,
	}
}
