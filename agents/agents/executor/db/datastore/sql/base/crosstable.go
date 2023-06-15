package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"math/big"
)

// GetTimestampForMessage gets the timestamp for a message. This is done in multiple logical steps:
// 1. Get all potential snapshot roots for the message (all snapshot roots that are associated to states with
// the same chain ID and a nonce greater than or equal to the message nonce).
// 2. Get the minimum destination block number for all attestations that are associated to the potential snapshot roots.
// 3. Return the timestamp of the attestation with the minimum destination block number.
func (s Store) GetTimestampForMessage(ctx context.Context, chainID, destination, nonce uint32) (*uint64, error) {
	statesTableName, err := dbcommon.GetModelName(s.DB(), &State{})
	if err != nil {
		return nil, fmt.Errorf("failed to get states table name: %w", err)
	}

	attestationsTableName, err := dbcommon.GetModelName(s.DB(), &Attestation{})
	if err != nil {
		return nil, fmt.Errorf("failed to get attestations table name: %w", err)
	}

	var timestamp uint64

	// TODO: Use string formatting to make this query more legible.
	dbTx := s.DB().WithContext(ctx).
		Raw(fmt.Sprintf(
			`SELECT %s FROM %s WHERE %s = (
					SELECT MIN(%s) FROM (
						(SELECT * FROM %s WHERE %s = ? AND %s >= ?) AS stateTable
						INNER JOIN
						(SELECT %s, %s FROM %s WHERE %s = ?) AS attestationTable
						ON stateTable.%s = attestationTable.%s
					)
				) LIMIT 1`,
			DestinationTimestampFieldName, attestationsTableName, DestinationBlockNumberFieldName,
			DestinationBlockNumberFieldName,
			statesTableName, ChainIDFieldName, NonceFieldName,
			SnapshotRootFieldName, DestinationBlockNumberFieldName, attestationsTableName, DestinationFieldName,
			SnapshotRootFieldName, SnapshotRootFieldName,
		), chainID, nonce, destination).
		Scan(&timestamp)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get timestamp for message: %w", dbTx.Error)
	}

	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil
	}

	return &timestamp, nil
}

// GetEarliestStateInRange gets the earliest state with the same snapshot root as an attestation within a nonce range.
// 1. Get all states that are within a nonce range.
// 2. Get the state with the earliest attestation associated to it.
func (s Store) GetEarliestStateInRange(ctx context.Context, chainID, destination, startNonce, endNonce uint32) (*agentsTypes.State, error) {
	statesTableName, err := dbcommon.GetModelName(s.DB(), &State{})
	if err != nil {
		return nil, fmt.Errorf("failed to get states table name: %w", err)
	}

	attestationsTableName, err := dbcommon.GetModelName(s.DB(), &Attestation{})
	if err != nil {
		return nil, fmt.Errorf("failed to get attestations table name: %w", err)
	}

	var state State

	// TODO: Use string formatting to make this query more legible.
	dbTx := s.DB().WithContext(ctx).
		Raw(fmt.Sprintf(
			`SELECT * FROM %s WHERE %s = ? AND %s = (
                     SELECT %s FROM %s WHERE %s = ? AND %s = (
						SELECT MIN(%s) FROM (
							(SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s = ?) AS stateTable
							INNER JOIN
							(SELECT %s, %s FROM %s WHERE %s = ?) as attestationTable
							ON stateTable.%s = attestationTable.%s
						)
					) ORDER BY %s DESC LIMIT 1
				)`,
			statesTableName, ChainIDFieldName, SnapshotRootFieldName,
			SnapshotRootFieldName, attestationsTableName, DestinationFieldName, DestinationBlockNumberFieldName,
			DestinationBlockNumberFieldName,
			SnapshotRootFieldName, statesTableName, NonceFieldName, NonceFieldName, ChainIDFieldName,
			SnapshotRootFieldName, DestinationBlockNumberFieldName, attestationsTableName, DestinationFieldName,
			SnapshotRootFieldName, SnapshotRootFieldName,
			AttestationNonceFieldName,
		), chainID, destination, startNonce, endNonce, chainID, destination).
		Scan(&state)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get earliest state in range: %w", dbTx.Error)
	}

	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil
	}

	gasData := agentsTypes.NewGasData(
		state.GDGasPrice,
		state.GDDataPrice,
		state.GDExecBuffer,
		state.GDAmortAttCost,
		state.GDEtherPrice,
		state.GDMarkup,
	)

	receivedState := agentsTypes.NewState(
		common.HexToHash(state.Root),
		state.ChainID,
		state.Nonce,
		big.NewInt(int64(state.OriginBlockNumber)),
		big.NewInt(int64(state.OriginTimestamp)),
		gasData,
	)

	return &receivedState, nil
}
