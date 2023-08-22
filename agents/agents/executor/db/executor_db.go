package db

import (
	"context"
	"encoding/json"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
)

// ExecutorDBWriter is the interface for writing to the executor database.
type ExecutorDBWriter interface {
	// StoreMessage stores a message in the database.
	StoreMessage(ctx context.Context, message agentsTypes.Message, blockNumber uint64, minimumTimeSet bool, minimumTime uint64) error
	// ExecuteMessage marks a message as executed in the database.
	ExecuteMessage(ctx context.Context, messageMask DBMessage) error
	// SetMinimumTime sets the minimum time of a message.
	SetMinimumTime(ctx context.Context, messageMask DBMessage, minimumTime uint64) error

	// StoreAttestation stores an attestation.
	StoreAttestation(ctx context.Context, attestation agentsTypes.Attestation, destination uint32, destinationBlockNumber, destinationTimestamp uint64) error

	// StoreState stores a state.
	StoreState(ctx context.Context, state agentsTypes.State, snapshotRoot [32]byte, proof [][]byte, stateIndex uint32, blockNumber uint64) error
	// StoreStates stores multiple states with the same snapshot root.
	StoreStates(ctx context.Context, states []agentsTypes.State, snapshotRoot [32]byte, proofs [][][]byte, blockNumber uint64) error
}

// ExecutorDBReader is the interface for reading from the executor database.
//
//nolint:interfacebloat
type ExecutorDBReader interface {
	// GetMessage gets a message from the database.
	GetMessage(ctx context.Context, messageMask DBMessage) (*agentsTypes.Message, error)
	// GetMessages gets messages from the database, paginated and ordered in ascending order by nonce.
	GetMessages(ctx context.Context, messageMask DBMessage, page int) ([]agentsTypes.Message, error)
	// GetBlockNumber gets the block number of a message from the database.
	GetBlockNumber(ctx context.Context, messageMask DBMessage) (uint64, error)
	// GetLastBlockNumber gets the last block number that had a message in the database.
	GetLastBlockNumber(ctx context.Context, chainID uint32, contractType types.ContractType) (uint64, error)
	// GetExecutableMessages gets executable messages from the database.
	GetExecutableMessages(ctx context.Context, messageMask DBMessage, currentTime uint64, page int) ([]agentsTypes.Message, error)
	// GetUnsetMinimumTimeMessages gets messages from the database that have not had their minimum time set.
	GetUnsetMinimumTimeMessages(ctx context.Context, messageMask DBMessage, page int) ([]agentsTypes.Message, error)
	// GetMessageMinimumTime gets the minimum time for a message to be executed.
	GetMessageMinimumTime(ctx context.Context, messageMask DBMessage) (*uint64, error)

	// GetAttestation gets an attestation that has fields matching the attestation mask.
	GetAttestation(ctx context.Context, attestationMask DBAttestation) (*agentsTypes.Attestation, error)
	// GetAttestationBlockNumber gets the block number of an attestation.
	GetAttestationBlockNumber(ctx context.Context, attestationMask DBAttestation) (*uint64, error)
	// GetAttestationTimestamp gets the timestamp of an attestation.
	GetAttestationTimestamp(ctx context.Context, attestationMask DBAttestation) (*uint64, error)
	// GetEarliestSnapshotFromAttestation takes a list of snapshot roots, checks which one has the lowest block number, and returns that snapshot root back.
	GetEarliestSnapshotFromAttestation(ctx context.Context, attestationMask DBAttestation, snapshotRoots []string) (*[32]byte, error)

	// GetState gets a state from the database.
	GetState(ctx context.Context, stateMask DBState) (*agentsTypes.State, error)
	// GetStateMetadata gets the snapshot root, proof, and tree height of a state from the database.
	GetStateMetadata(ctx context.Context, stateMask DBState) (snapshotRoot *[32]byte, proof *json.RawMessage, stateIndex *uint32, err error)
	// GetPotentialSnapshotRoots gets all snapshot roots that are greater than or equal to a specified nonce and matches
	// a specified chain ID.
	GetPotentialSnapshotRoots(ctx context.Context, chainID uint32, nonce uint32) ([]string, error)
	// GetSnapshotRootsInNonceRange gets all snapshot roots for all states in a specified nonce range.
	GetSnapshotRootsInNonceRange(ctx context.Context, chainID uint32, startNonce uint32, endNonce uint32) ([]string, error)

	// GetTimestampForMessage gets the timestamp for a message. This is done in multiple logical steps:
	// 1. Get all potential snapshot roots for the message (all snapshot roots that are associated to states with
	// the same chain ID and a nonce greater than or equal to the message nonce).
	// 2. Get the minimum destination block number for all attestations that are associated to the potential snapshot roots.
	// 3. Return the timestamp of the attestation with the minimum destination block number.
	GetTimestampForMessage(ctx context.Context, chainID, destination, nonce uint32) (*uint64, error)
	// GetEarliestStateInRange gets the earliest state with the same snapshot root as an attestation within a nonce range.
	// 1. Get all states that are within a nonce range.
	// 2. Get the state with the earliest attestation associated to it.
	GetEarliestStateInRange(ctx context.Context, chainID, destination, startNonce, endNonce uint32) (*agentsTypes.State, error)
}

// ExecutorDB is the interface for the executor database.
type ExecutorDB interface {
	ExecutorDBWriter
	ExecutorDBReader
	SubmitterDB() submitterDB.Service
}
