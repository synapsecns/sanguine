package base

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
	"gorm.io/gorm/clause"
	"math/big"
)

// StoreState stores a state.
func (s Store) StoreState(ctx context.Context, state agentsTypes.State, snapshotRoot [32]byte, proof [][]byte, treeHeight uint32) error {
	dbState, err := AgentsTypesStateToState(state, snapshotRoot, proof, treeHeight)
	if err != nil {
		return fmt.Errorf("failed to convert state to db state: %w", err)
	}

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: SnapshotRootFieldName}, {Name: RootFieldName}, {Name: ChainIDFieldName}, {Name: NonceFieldName},
			},
			DoNothing: true,
		}).
		Create(&dbState)

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store state: %w", dbTx.Error)
	}

	return nil
}

// StoreStates stores multiple states with the same snapshot root.
func (s Store) StoreStates(ctx context.Context, states []agentsTypes.State, snapshotRoot [32]byte, proofs [][][]byte, treeHeight uint32) error {
	var dbStates []State
	for i := range states {
		state, err := AgentsTypesStateToState(states[i], snapshotRoot, proofs[i], treeHeight)
		if err != nil {
			return fmt.Errorf("failed to convert state to db state: %w", err)
		}

		dbStates = append(dbStates, state)
	}

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: SnapshotRootFieldName}, {Name: RootFieldName}, {Name: ChainIDFieldName}, {Name: NonceFieldName},
			},
			DoNothing: true,
		}).
		Create(&dbStates)

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store states: %w", dbTx.Error)
	}

	return nil
}

// GetState gets a state from the database.
func (s Store) GetState(ctx context.Context, stateMask types.DBState) (*agentsTypes.State, error) {
	var state State

	dbStateMask := DBStateToState(stateMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&state).
		Where(&dbStateMask).
		Limit(1).
		Scan(&state)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get state: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil
	}

	receivedState := agentsTypes.NewState(
		common.HexToHash(state.Root),
		state.ChainID,
		state.Nonce,
		big.NewInt(int64(state.OriginBlockNumber)),
		big.NewInt(int64(state.OriginTimestamp)),
	)

	return &receivedState, nil
}

// GetStateMetadata gets the snapshot root, proof, and tree height of a state from the database.
func (s Store) GetStateMetadata(ctx context.Context, stateMask types.DBState) (snapshotRoot *[32]byte, proof *json.RawMessage, treeHeight *uint32, err error) {
	var state State

	dbStateMask := DBStateToState(stateMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&state).
		Where(&dbStateMask).
		Limit(1).
		Scan(&state)
	if dbTx.Error != nil {
		return nil, nil, nil, fmt.Errorf("failed to get state snapshot root: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		return nil, nil, nil, nil
	}

	var snapshotRootB32 [32]byte
	copy(snapshotRoot[:], common.HexToHash(state.SnapshotRoot).Bytes())

	snapshotRoot = &snapshotRootB32
	proof = &state.Proof
	treeHeight = &state.TreeHeight

	return
}

// DBStateToState converts a DBState to a State.
func DBStateToState(dbState types.DBState) State {
	var state State

	if dbState.SnapshotRoot != nil {
		state.SnapshotRoot = *dbState.SnapshotRoot
	}

	if dbState.Root != nil {
		state.Root = *dbState.Root
	}

	if dbState.ChainID != nil {
		state.ChainID = *dbState.ChainID
	}

	if dbState.Nonce != nil {
		state.Nonce = *dbState.Nonce
	}

	if dbState.OriginBlockNumber != nil {
		state.OriginBlockNumber = *dbState.OriginBlockNumber
	}

	if dbState.OriginTimestamp != nil {
		state.OriginTimestamp = *dbState.OriginTimestamp
	}

	if dbState.Proof != nil {
		state.Proof = *dbState.Proof
	}

	if dbState.TreeHeight != nil {
		state.TreeHeight = *dbState.TreeHeight
	}

	return state
}

// StateToDBState converts a State to a DBState.
func StateToDBState(state State) types.DBState {
	snapshotRoot := state.SnapshotRoot
	root := state.Root
	chainID := state.ChainID
	nonce := state.Nonce
	originBlockNumber := state.OriginBlockNumber
	originTimestamp := state.OriginTimestamp
	proof := state.Proof
	treeHeight := state.TreeHeight

	return types.DBState{
		SnapshotRoot:      &snapshotRoot,
		Root:              &root,
		ChainID:           &chainID,
		Nonce:             &nonce,
		OriginBlockNumber: &originBlockNumber,
		OriginTimestamp:   &originTimestamp,
		Proof:             &proof,
		TreeHeight:        &treeHeight,
	}
}

// AgentsTypesStateToState converts an agentsTypes.State to a State.
func AgentsTypesStateToState(state agentsTypes.State, snapshotRoot [32]byte, proof [][]byte, treeHeight uint32) (State, error) {
	root := state.Root()

	// Convert the proof to a json
	proofJSON, err := json.Marshal(proof)
	if err != nil {
		return State{}, fmt.Errorf("failed to marshal proof: %w", err)
	}

	proofDBFormat := json.RawMessage(proofJSON)

	return State{
		SnapshotRoot:      common.BytesToHash(snapshotRoot[:]).String(),
		Root:              common.BytesToHash(root[:]).String(),
		ChainID:           state.Origin(),
		Nonce:             state.Nonce(),
		OriginBlockNumber: state.BlockNumber().Uint64(),
		OriginTimestamp:   state.Timestamp().Uint64(),
		Proof:             proofDBFormat,
		TreeHeight:        treeHeight,
	}, nil
}
