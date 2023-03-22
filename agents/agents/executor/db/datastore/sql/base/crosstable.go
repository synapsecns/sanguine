package base

import (
	"context"
	"fmt"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
)

// GetTimestampForMessage gets the timestamp for a message. This is done in multiple logical steps:
// 1. Get all potential snapshot roots for the message (all snapshot roots that are associated to states with
// the same chain ID and a nonce greater than or equal to the message nonce).
// 2. Get the minimum destination block number for all attestations that are associated to the potential snapshot roots.
// 3. Return the timestamp of the attestation with the minimum destination block number.
func (s Store) GetTimestampForMessage(ctx context.Context, chainID, destination, nonce uint32, tablePrefix string) (*uint64, error) {
	var timestamp uint64

	statesTableName := "states"
	attestationsTableName := "attestations"

	if tablePrefix != "" {
		statesTableName = fmt.Sprintf("%s_%s", tablePrefix, statesTableName)
		attestationsTableName = fmt.Sprintf("%s_%s", tablePrefix, attestationsTableName)
	}

	dbTx := s.DB().WithContext(ctx).
		Raw(fmt.Sprintf(
			`SELECT %s FROM %s WHERE %s = ? AND %s = (
					SELECT MIN(%s) FROM (
						(SELECT * FROM %s WHERE %s = ? AND %s >= ?) AS stateTable
						INNER JOIN
						(SELECT %s, %s FROM %s) AS attestationTable
						ON stateTable.%s = attestationTable.%s
					)
				)`,
			DestinationTimestampFieldName, attestationsTableName, DestinationFieldName, DestinationBlockNumberFieldName,
			DestinationBlockNumberFieldName,
			statesTableName, ChainIDFieldName, NonceFieldName,
			SnapshotRootFieldName, DestinationBlockNumberFieldName, attestationsTableName,
			SnapshotRootFieldName, SnapshotRootFieldName,
		), destination, chainID, nonce).
		Scan(&timestamp)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get timestamp for message: %w", dbTx.Error)
	}

	return &timestamp, nil
}

// GetEarliestStateInRange gets the earliest state with the same snapshot root as an attestation within a nonce range.
// This is done in multiple logical steps:
// 1. Get all snapshot roots that are within a nonce range.
// 2. Get the earliest snapshot root from the list of snapshot roots.
// 3. Get the state with the earliest snapshot root.

// 1. Get all states that are within a nonce range.
// 2. Get the state with the earliest attestation associated to it.
func (s Store) GetEarliestStateInRange(ctx context.Context, chainID, destination, startNonce, endNonce uint32) (*agentsTypes.State, error) {
	// var state agentsTypes.State

	return nil, nil
}
