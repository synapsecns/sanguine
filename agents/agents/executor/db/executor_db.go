package db

import (
	"context"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
)

// ExecutorDBWriter is the interface for writing to the executor database.
type ExecutorDBWriter interface {
	// StoreMessage stores a message in the database.
	StoreMessage(ctx context.Context, message agentsTypes.Message, blockNumber uint64, minimumTimeSet bool, minimumTime uint64) error
	// ExecuteMessage marks a message as executed in the database.
	ExecuteMessage(ctx context.Context, messageMask types.DBMessage) error
	// SetMinimumTime sets the minimum time of a message.
	SetMinimumTime(ctx context.Context, messageMask types.DBMessage, minimumTime uint64) error

	// StoreAttestation stores an attestation.
	StoreAttestation(ctx context.Context, attestation agentsTypes.Attestation, blockNumber uint64, blockTime uint64) error
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

	// GetAttestation gets an attestation from the database.
	GetAttestation(ctx context.Context, attestationMask types.DBAttestation) (*agentsTypes.Attestation, error)
	// GetAttestationBlockNumber gets the block number of an attestation.
	GetAttestationBlockNumber(ctx context.Context, attestationMask types.DBAttestation) (*uint64, error)
	// GetAttestationBlockTime gets the block time of an attestation.
	GetAttestationBlockTime(ctx context.Context, attestationMask types.DBAttestation) (*uint64, error)
	// GetAttestationForNonceOrGreater gets the lowest nonce attestation that is greater than or equal to the given nonce.
	GetAttestationForNonceOrGreater(ctx context.Context, attestationMask types.DBAttestation) (nonce *uint32, blockTime *uint64, err error)
	// GetAttestationsAboveOrEqualNonce gets attestations in a nonce range.
	GetAttestationsAboveOrEqualNonce(ctx context.Context, attestationMask types.DBAttestation, minNonce uint32, page int) ([]types.DBAttestation, error)
	// GetEarliestAttestationsNonceInNonceRange gets the earliest attestation (by block number) in a nonce range.
	GetEarliestAttestationsNonceInNonceRange(ctx context.Context, attestationMask types.DBAttestation, minNonce uint32, maxNonce uint32) (*uint32, error)
}

// ExecutorDB is the interface for the executor database.
type ExecutorDB interface {
	ExecutorDBWriter
	ExecutorDBReader
}
