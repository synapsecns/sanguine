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
	DisputeIndexFieldName = namer.GetConsistentName("DisputeIndex")
	AgentAddressFieldName = namer.GetConsistentName("AgentAddress")
	BlockNumberFieldName = namer.GetConsistentName("BlockNumber")
	DisputeProcessedStatusFieldName = namer.GetConsistentName("DisputeProcessedStatus")
	NotaryAddressFieldName = namer.GetConsistentName("NotaryAddress")
}

var (
	// AgentRootFieldName is the field name of the agent root.
	AgentRootFieldName string
	// DisputeIndexFieldName is the field name of the agent root.
	DisputeIndexFieldName string
	// AgentAddressFieldName gets the agent address field name.
	AgentAddressFieldName string
	// BlockNumberFieldName gets the agent block number field name.
	BlockNumberFieldName string
	// DisputeProcessedStatusFieldName gets the dispute processed status field name.
	DisputeProcessedStatusFieldName string
	// NotaryAddressFieldName gets the notary address field name.
	NotaryAddressFieldName string
)

// Dispute is a dispute between two agents.
type Dispute struct {
	// DisputeIndex is the index of the dispute on the BondingManager.
	DisputeIndex uint64 `gorm:"column:dispute_index;primaryKey"`
	// DisputeProcessedStatus indicates the status of the dispute.
	DisputeProcessedStatus agentTypes.DisputeProcessedStatus `gorm:"column:dispute_processed_status"`
	// GuardAddress is the address of the guard.
	GuardAddress string `gorm:"column:guard_address"`
	// NotaryIndex is the index of the notary on the BondingManager.
	NotaryIndex uint64 `gorm:"column:notary_index"`
	// NotaryAddress is the address of the notary.
	NotaryAddress string `gorm:"column:notary_address"`
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
