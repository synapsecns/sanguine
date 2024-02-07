package base

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
	"gorm.io/gorm/clause"
	"math/big"
)

// StoreState stores a state.
func (s Store) StoreState(ctx context.Context, state agentsTypes.State, snapshotRoot [32]byte, proof [][]byte, stateIndex uint32, blockNumber uint64) error {
	dbState, err := AgentsTypesStateToState(state, snapshotRoot, proof, stateIndex, blockNumber)
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
func (s Store) StoreStates(ctx context.Context, states []agentsTypes.State, snapshotRoot [32]byte, proofs [][][]byte, blockNumber uint64) error {
	var dbStates []State
	for i := range states {
		state, err := AgentsTypesStateToState(states[i], snapshotRoot, proofs[i], uint32(i), blockNumber)
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
func (s Store) GetState(ctx context.Context, stateMask db.DBState) (*agentsTypes.State, error) {
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

// GetStateMetadata gets the snapshot root, proof, and tree height of a state from the database.
func (s Store) GetStateMetadata(ctx context.Context, stateMask db.DBState) (snapshotRoot *[32]byte, proof *json.RawMessage, stateIndex *uint32, err error) {
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

	snapshotRootHash := common.HexToHash(state.SnapshotRoot)
	snapshotRoot = (*[32]byte)(&snapshotRootHash)
	proof = &state.Proof
	stateIndex = &state.StateIndex

	return
}

// GetPotentialSnapshotRoots gets all snapshot roots that are greater than or equal to a specified nonce and matches
// a specified chain ID.
func (s Store) GetPotentialSnapshotRoots(ctx context.Context, chainID uint32, nonce uint32) ([]string, error) {
	var states []State

	dbTx := s.DB().WithContext(ctx).
		Model(&states).
		Where(fmt.Sprintf("%s = ? AND %s >= ?", ChainIDFieldName, NonceFieldName), chainID, nonce).
		Scan(&states)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get potential snapshot roots: %w", dbTx.Error)
	}

	var snapshotRoots []string
	for _, state := range states {
		snapshotRoots = append(snapshotRoots, state.SnapshotRoot)
	}

	return snapshotRoots, nil
}

// GetSnapshotRootsInNonceRange gets all snapshot roots for all states in a specified nonce range.
func (s Store) GetSnapshotRootsInNonceRange(ctx context.Context, chainID uint32, startNonce uint32, endNonce uint32) ([]string, error) {
	var states []State

	dbTx := s.DB().WithContext(ctx).
		Model(&states).
		Where(fmt.Sprintf("%s = ? AND %s >= ? AND %s <= ?", ChainIDFieldName, NonceFieldName, NonceFieldName), chainID, startNonce, endNonce).
		Scan(&states)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get snapshot roots in nonce range: %w", dbTx.Error)
	}

	var snapshotRoots []string
	for _, state := range states {
		snapshotRoots = append(snapshotRoots, state.SnapshotRoot)
	}

	return snapshotRoots, nil
}

// DBStateToState converts a DBState to a State.
// nolint:cyclop
func DBStateToState(dbState db.DBState) State {
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

	if dbState.StateIndex != nil {
		state.StateIndex = *dbState.StateIndex
	}

	if dbState.BlockNumber != nil {
		state.BlockNumber = *dbState.BlockNumber
	}

	if dbState.GDGasPrice != nil {
		state.GDGasPrice = *dbState.GDGasPrice
	}

	if dbState.GDDataPrice != nil {
		state.GDDataPrice = *dbState.GDDataPrice
	}

	if dbState.GDExecBuffer != nil {
		state.GDExecBuffer = *dbState.GDExecBuffer
	}

	if dbState.GDAmortAttCost != nil {
		state.GDAmortAttCost = *dbState.GDAmortAttCost
	}

	if dbState.GDEtherPrice != nil {
		state.GDEtherPrice = *dbState.GDEtherPrice
	}

	if dbState.GDMarkup != nil {
		state.GDMarkup = *dbState.GDMarkup
	}

	return state
}

// StateToDBState converts a State to a DBState.
func StateToDBState(state State) db.DBState {
	snapshotRoot := state.SnapshotRoot
	root := state.Root
	chainID := state.ChainID
	nonce := state.Nonce
	originBlockNumber := state.OriginBlockNumber
	originTimestamp := state.OriginTimestamp
	proof := state.Proof
	stateIndex := state.StateIndex
	blockNumber := state.BlockNumber
	gasPrice := state.GDGasPrice
	dataPrice := state.GDDataPrice
	execBuffer := state.GDExecBuffer
	amortAttCost := state.GDAmortAttCost
	etherPrice := state.GDEtherPrice
	markup := state.GDMarkup

	return db.DBState{
		SnapshotRoot:      &snapshotRoot,
		Root:              &root,
		ChainID:           &chainID,
		Nonce:             &nonce,
		OriginBlockNumber: &originBlockNumber,
		OriginTimestamp:   &originTimestamp,
		Proof:             &proof,
		StateIndex:        &stateIndex,
		BlockNumber:       &blockNumber,
		GDGasPrice:        &gasPrice,
		GDDataPrice:       &dataPrice,
		GDExecBuffer:      &execBuffer,
		GDAmortAttCost:    &amortAttCost,
		GDEtherPrice:      &etherPrice,
		GDMarkup:          &markup,
	}
}

// AgentsTypesStateToState converts an agentsTypes.State to a State.
func AgentsTypesStateToState(state agentsTypes.State, snapshotRoot [32]byte, proof [][]byte, stateIndex uint32, blockNumber uint64) (State, error) {
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
		StateIndex:        stateIndex,
		BlockNumber:       blockNumber,
		GDGasPrice:        state.GasData().GasPrice(),
		GDDataPrice:       state.GasData().DataPrice(),
		GDExecBuffer:      state.GasData().ExecBuffer(),
		GDAmortAttCost:    state.GasData().AmortAttCost(),
		GDEtherPrice:      state.GasData().EtherPrice(),
		GDMarkup:          state.GasData().Markup(),
	}, nil
}
