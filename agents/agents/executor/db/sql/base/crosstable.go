package base

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/imkira/go-interpol"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

// GetTimestampForMessage gets the timestamp for a message. This is done in multiple steps:
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

	query, err := interpol.WithMap(
		`SELECT {destTimestamp} FROM {attTable} WHERE {destBlockNum} = (
					SELECT MIN({destBlockNum}) FROM (
						(SELECT * FROM {stTable} WHERE {chainID} = ? AND {nonce} >= ?) AS stateTable
						INNER JOIN
						(SELECT {snapshotRoot}, {destBlockNum} FROM {attTable} WHERE {destination} = ?) AS attestationTable
						ON stateTable.{snapshotRoot}= attestationTable.{snapshotRoot}
					)
				) ORDER BY {attNonce} DESC LIMIT 1`,
		map[string]string{
			"destTimestamp": DestinationTimestampFieldName,
			"attTable":      attestationsTableName,
			"destBlockNum":  DestinationBlockNumberFieldName,
			"stTable":       statesTableName,
			"chainID":       ChainIDFieldName,
			"nonce":         NonceFieldName,
			"snapshotRoot":  SnapshotRootFieldName,
			"destination":   DestinationFieldName,
			"attNonce":      AttestationNonceFieldName,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to interpolate GetTimestampForMessage query: %w", err)
	}

	dbTx := s.DB().WithContext(ctx).
		Raw(query, chainID, nonce, destination).
		Scan(&timestamp)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get timestamp for message: %w", dbTx.Error)
	}

	fmt.Printf("got timestamp for message with rowsAffected %d: %d", dbTx.RowsAffected, timestamp)

	// if dbTx.RowsAffected == 0 {
	// 	//nolint:nilnil
	// 	return nil, nil
	// }
	if timestamp == 0 {
		//nolint:nilnil
		return nil, nil
	}

	return &timestamp, nil
}

// GetEarliestStateInRange gets the earliest state with the same snapshot root as an attestation within a nonce range.
// 1. Get all states that are within a nonce range.
// 2. Get the state with the earliest attestation associated to it.
func (s Store) GetEarliestStateInRange(ctx context.Context, chainID, destination, startNonce, endNonce uint32) (*agentsTypes.State, *string, error) {
	statesTableName, err := dbcommon.GetModelName(s.DB(), &State{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get states table name: %w", err)
	}

	attestationsTableName, err := dbcommon.GetModelName(s.DB(), &Attestation{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get attestations table name: %w", err)
	}

	var state State

	query, err := interpol.WithMap(
		`SELECT * FROM {stTable} WHERE {chainID} = ? AND {snapshotRoot} = (
                     SELECT {snapshotRoot} FROM {attTable} WHERE {destination} = ? AND ({snapshotRoot}) = (
						SELECT attestationTable.{snapshotRoot} FROM (
							(SELECT {snapshotRoot} FROM {stTable} WHERE {nonce} >= ? AND {nonce} <= ? AND {chainID} = ?) AS stateTable
							INNER JOIN
							(SELECT {snapshotRoot}, {destBlockNum} FROM {attTable} WHERE {destination} = ?) as attestationTable
							ON stateTable.{snapshotRoot} = attestationTable.{snapshotRoot}
						) ORDER BY {destBlockNum} ASC LIMIT 1
					) ORDER BY {attNonce} DESC LIMIT 1
				)`,
		map[string]string{
			"stTable":      statesTableName,
			"chainID":      ChainIDFieldName,
			"snapshotRoot": SnapshotRootFieldName,
			"attTable":     attestationsTableName,
			"destination":  DestinationFieldName,
			"destBlockNum": DestinationBlockNumberFieldName,
			"nonce":        NonceFieldName,
			"attNonce":     AttestationNonceFieldName,
		},
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to interpolate GetEarliestStateInRange query: %w", err)
	}

	dbTx := s.DB().WithContext(ctx).
		Raw(query, chainID, destination, startNonce, endNonce, chainID, destination).
		Scan(&state)
	if dbTx.Error != nil {
		return nil, nil, fmt.Errorf("failed to get earliest state in range: %w", dbTx.Error)
	}

	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil, nil
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

	return &receivedState, &state.SnapshotRoot, nil
}
