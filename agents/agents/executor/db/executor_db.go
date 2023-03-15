package db

import (
	"context"
	"encoding/json"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
)

// ExecutorDBWriter is the interface for writing to the executor database.
type ExecutorDBWriter interface {
	// StoreMessage stores a message in the database.
	StoreMessage(ctx context.Context, message agentsTypes.Message, blockNumber uint64, attestationNonce uint32, minimumTime uint64) error
	// ExecuteMessage marks a message as executed in the database.
	ExecuteMessage(ctx context.Context, messageMask types.DBMessage) error
	// SetMinimumTime sets the minimum time of a message.
	SetMinimumTime(ctx context.Context, messageMask types.DBMessage, attestationNonce uint32, minimumTime uint64) error

	// StoreAttestation stores an attestation.
	StoreAttestation(ctx context.Context, attestation agentsTypes.Attestation, destination uint32, destinationBlockNumber, destinationTimestamp uint64) error

	// StoreState stores a state.
	StoreState(ctx context.Context, state agentsTypes.State, snapshotRoot [32]byte, proof [][]byte, treeHeight, stateIndex uint32) error
	// StoreStates stores multiple states with the same snapshot root.
	StoreStates(ctx context.Context, states []agentsTypes.State, snapshotRoot [32]byte, proofs [][][]byte, treeHeight uint32) error
}

// ExecutorDBReader is the interface for reading from the executor database.
//
//nolint:interfacebloat
type ExecutorDBReader interface {
	// GetMessage gets a message from the database.
	GetMessage(ctx context.Context, messageMask types.DBMessage) (*agentsTypes.Message, error)
	// GetMessages gets messages from the database, paginated and ordered in ascending order by nonce.
	GetMessages(ctx context.Context, messageMask types.DBMessage, page int) ([]agentsTypes.Message, error)
	// GetBlockNumber gets the block number of a message from the database.
	GetBlockNumber(ctx context.Context, messageMask types.DBMessage) (uint64, error)
	// GetLastBlockNumber gets the last block number that had a message in the database.
	GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
	// GetExecutableMessages gets executable messages from the database.
	GetExecutableMessages(ctx context.Context, messageMask types.DBMessage, currentTime uint64, page int) ([]agentsTypes.Message, error)
	// GetUnsetMinimumTimeMessages gets messages from the database that have not had their minimum time set.
	GetUnsetMinimumTimeMessages(ctx context.Context, messageMask types.DBMessage, page int) ([]agentsTypes.Message, error)
	// GetMessageMinimumTime gets the minimum time for a message to be executed.
	GetMessageMinimumTime(ctx context.Context, messageMask types.DBMessage) (*uint64, error)

	// GetAttestation gets an attestation that has fields matching the attestation mask.
	GetAttestation(ctx context.Context, attestationMask types.DBAttestation) (*agentsTypes.Attestation, error)
	// GetAttestationBlockNumber gets the block number of an attestation.
	GetAttestationBlockNumber(ctx context.Context, attestationMask types.DBAttestation) (*uint64, error)
	// GetAttestationTimestamp gets the timestamp of an attestation.
	GetAttestationTimestamp(ctx context.Context, attestationMask types.DBAttestation) (*uint64, error)
	// GetAttestationMinimumTimestampAndNonce takes a list of snapshot roots and returns the timestamp and nonce of the attestation with the lowest block number.
	GetAttestationMinimumTimestampAndNonce(ctx context.Context, attestationMask types.DBAttestation, snapshotRoots []string) (*uint64, *uint32, error)
	// GetEarliestSnapshotFromAttestation takes a list of snapshot roots, checks which one has the lowest block number, and returns that snapshot root back.
	GetEarliestSnapshotFromAttestation(ctx context.Context, attestationMask types.DBAttestation, snapshotRoots []string) (*[32]byte, error)

	// GetState gets a state from the database.
	GetState(ctx context.Context, stateMask types.DBState) (*agentsTypes.State, error)
	// GetStateMetadata gets the snapshot root, proof, and tree height of a state from the database.
	GetStateMetadata(ctx context.Context, stateMask types.DBState) (snapshotRoot *[32]byte, proof *json.RawMessage, treeHeight *uint32, stateIndex *uint32, err error)
	// GetPotentialSnapshotRoots gets all snapshot roots that are greater than or equal to a specified nonce and matches
	// a specified chain ID.
	GetPotentialSnapshotRoots(ctx context.Context, chainID uint32, nonce uint32) ([]string, error)
	// GetSnapshotRootsInNonceRange gets all snapshot roots for all states in a specified nonce range.
	GetSnapshotRootsInNonceRange(ctx context.Context, chainID uint32, startNonce uint32, endNonce uint32) ([]string, error)
}

// ExecutorDB is the interface for the executor database.
type ExecutorDB interface {
	ExecutorDBWriter
	ExecutorDBReader
}
