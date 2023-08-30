package base

import (
	"encoding/json"

	agentTypes "github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

// define common field names. See package docs  for an explanation of why we have to do this.
// note: some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	AgentRootFieldName = namer.GetConsistentName("AgentRoot")
	AgentAddressFieldName = namer.GetConsistentName("AgentAddress")
	BlockNumberFieldName = namer.GetConsistentName("BlockNumber")
	AgentStatusRelayedStateFieldName = namer.GetConsistentName("AgentStatusRelayedState")
	DomainFieldName = namer.GetConsistentName("Domain")
	UpdatedFlagFieldName = namer.GetConsistentName("UpdatedFlag")
}

var (
	// AgentRootFieldName is the field name of the agent root.
	AgentRootFieldName string
	// AgentAddressFieldName gets the agent address field name.
	AgentAddressFieldName string
	// BlockNumberFieldName gets the agent block number field name.
	BlockNumberFieldName string
	// AgentStatusRelayedStateFieldName gets the relayable agent status field name.
	AgentStatusRelayedStateFieldName string
	// DomainFieldName gets the agent domain field name.
	DomainFieldName string
	// UpdatedFlagFieldName gets the updated flag field name.
	UpdatedFlagFieldName string
)

// RelayableAgentStatus is used for tracking agent statuses that are out of
// sync and need to be relayed to a remote chain.
type RelayableAgentStatus struct {
	AgentAddress string `gorm:"column:agent_address"`
	// StaleFlag is the old flag that needs to be updated.
	StaleFlag agentTypes.AgentFlagType `gorm:"column:stale_flag"`
	// UpdatedFlag is the new flag value that should be relayed.
	UpdatedFlag agentTypes.AgentFlagType `gorm:"column:updated_flag"`
	// Domain is the domain of the agent status.
	Domain uint32 `gorm:"column:domain"`
	// AgentStatusRelayedState is the state of the relayable agent status.
	AgentStatusRelayedState agentTypes.AgentStatusRelayedState `gorm:"column:agent_status_relayed_state"`
}

// AgentTree is the state of an agent tree on Summit.
type AgentTree struct {
	// AgentRoot is the root of the agent tree.
	AgentRoot string `gorm:"column:agent_root;primaryKey"`
	// AgentAddress is the address of the agent for the Merkle proof.
	AgentAddress string `gorm:"column:agent_address;primaryKey"`
	// BlockNumber is the block number that the agent tree was updated (on summit).
	BlockNumber uint64 `gorm:"column:block_number"`
	// Proof is the agent tree proof.
	Proof json.RawMessage `gorm:"column:proof"`
}

// AgentRoot is the state of the agent roots on summit.
type AgentRoot struct {
	// AgentRoot is the root of the agent tree.
	AgentRoot string `gorm:"column:agent_root;primaryKey"`
	// BlockNumber is the block number that the agent tree was updated.
	BlockNumber uint64 `gorm:"column:block_number"`
}
