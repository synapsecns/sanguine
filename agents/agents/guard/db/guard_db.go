package db

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	agentTypes "github.com/synapsecns/sanguine/agents/types"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
)

// GuardDBWriter is the interface for writing to the guard's database.
type GuardDBWriter interface {
	// StoreDispute stores a dispute.
	StoreDispute(
		ctx context.Context,
		disputeIndex *big.Int,
		disputeProcessedStatus agentTypes.DisputeProcessedStatus,
		guardAddress common.Address,
		notaryIndex uint32,
		notaryAddress common.Address,
	) error

	// UpdateDisputeProcessedStatus updates the DisputedProcessedStatus for a dispute.
	UpdateDisputeProcessedStatus(
		ctx context.Context,
		guardAddress *common.Address,
		notaryAddress *common.Address,
		flag agentTypes.DisputeProcessedStatus,
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
		blockNumber uint64,
	) error
}

// GuardDBReader is the interface for reading from the guard's database.
type GuardDBReader interface {
	// GetUpdateAgentStatusParameters gets eligible parameters for the updateAgentStatus() contract call.
	GetUpdateAgentStatusParameters(ctx context.Context) ([]agentTypes.AgentTree, error)
	// GetSummitBlockNumberForRoot gets the summit block number for a given agent root.
	GetSummitBlockNumberForRoot(ctx context.Context, agentRoot [32]byte) (uint64, error)
}

// GuardDB is the interface for the guard's database.
type GuardDB interface {
	GuardDBWriter
	GuardDBReader
	SubmitterDB() submitterDB.Service
}
