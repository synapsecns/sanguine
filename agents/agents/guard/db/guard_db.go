package db

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	agentTypes "github.com/synapsecns/sanguine/agents/types"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
)

// GuardDBWriter is the interface for writing to the guard's database.
type GuardDBWriter interface {
	// StoreRelayableAgentStatus stores a relayable agent status.
	StoreRelayableAgentStatus(
		ctx context.Context,
		agentAddress common.Address,
		staleFlag agentTypes.AgentFlagType,
		updatedFlag agentTypes.AgentFlagType,
		domain uint32,
	) error

	// UpdateAgentStatusRelayedState updates the relayed state for a relayable agent status.
	UpdateAgentStatusRelayedState(
		ctx context.Context,
		agentAddress common.Address,
		state agentTypes.AgentStatusRelayedState,
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
	// GetRelayableAgentStatuses gets eligible parameters for the updateAgentStatus() contract call.
	GetRelayableAgentStatuses(ctx context.Context, chainID uint32) ([]agentTypes.AgentTree, error)
	// GetSummitBlockNumberForRoot gets the summit block number for a given agent root.
	GetSummitBlockNumberForRoot(ctx context.Context, agentRoot string) (uint64, error)
}

// GuardDB is the interface for the guard's database.
type GuardDB interface {
	GuardDBWriter
	GuardDBReader
	SubmitterDB() submitterDB.Service
}
