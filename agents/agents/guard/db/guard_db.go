package db

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
)

// GuardDBWriter is the interface for writing to the guard's database.
type GuardDBWriter interface {
	// StoreDispute stores a dispute.
	StoreDispute(
		ctx context.Context,
		disputeIndex *big.Int,
		disputeProcessedStatus guard.DisputeProcessedStatus,
		guardAddress common.Address,
		notaryIndex uint32,
		notaryAddress common.Address,
	) error

	// StoreAgentTree stores an agent tree.
	StoreAgentTree(
		ctx context.Context,
		agentRoot [32]byte,
		agentAddress common.Address,
		blockNumber uint64,
		proof [][32]byte,
	) error

	// StoreAgentRoot stores an agent root.
	StoreAgentRoot(
		ctx context.Context,
		agentRoot [32]byte,
		chainID uint32,
		blockNumber uint64,
	) error
}

// GuardDBReader is the interface for reading from the guard's database.
type GuardDBReader interface {
	// GetUpdateAgentStatusParameters gets eligible parameters for the updateAgentStatus() contract call.
	GetUpdateAgentStatusParameters(ctx context.Context) ([]AgentTree, error)
	// GetLatestConfirmedSummitBlockNumber gets the latest confirmed summit block number.
	GetLatestConfirmedSummitBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
}

// GuardDB is the interface for the guard's database.
type GuardDB interface {
	GuardDBWriter
	GuardDBReader
	SubmitterDB() submitterDB.Service
}
