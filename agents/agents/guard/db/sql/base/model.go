package base

import (
	"encoding/json"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

// define common field names. See package docs  for an explanation of why we have to do this.
// note: some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	AgentRootFieldName = namer.GetConsistentName("AgentRoot")
	DisputeIndexFieldName = namer.GetConsistentName("DisputeIndex")
	ChainIDFieldName = namer.GetConsistentName("ChainID")
}

var (
	// AgentRootFieldName is the field name of the agent root.
	AgentRootFieldName string
	// DisputeIndexFieldName is the field name of the agent root.
	DisputeIndexFieldName string
	// ChainIDFieldName gets the chain id field name.
	ChainIDFieldName string
)

// Dispute is a dispute between two agents.
// TODO: Change guard index and notary index to addresses? Requires an additional call.
type Dispute struct {
	// AgentRoot is the root of the agent tree.
	AgentRoot string `gorm:"column:agent_root;primaryKey"`
	// DisputeIndex is the index of the dispute on the BondingManager.
	DisputeIndex uint64 `gorm:"column:dispute_index;primaryKey"`
	// Resolved is if the dispute has been resolved.
	Resolved bool `gorm:"column:resolved"`
	// GuardIndex is the index of the guard on the BondingManager.
	GuardIndex uint64 `gorm:"column:guard_index"`
	// GuardAddress is the address of the guard.
	GuardAddress string `gorm:"column:guard_address"`
	//NotaryIndex is the index of the notary on the BondingManager.
	NotaryIndex uint64 `gorm:"column:notary_index"`
	// NotaryAddress is the address of the notary.
	NotaryAddress string `gorm:"column:notary_address"`
}

// AgentTree is the state of an agent tree on Summit.
type AgentTree struct {
	// AgentRoot is the root of the agent tree.
	AgentRoot string `gorm:"column:agent_root;primaryKey"`
	// BlockNumber is the block number that the agent tree was updated on Summit.
	BlockNumber uint64 `gorm:"column:block_number"`
	// Proof is the agent tree proof.
	Proof json.RawMessage `gorm:"column:proof"`
}

// RemoteAgentRoot is the state of an agent tree on a remote chain.
type RemoteAgentRoot struct {
	// AgentRoot is the root of the agent tree accepted on the remote chain.
	AgentRoot string `gorm:"column:agent_root;primaryKey"`
	// ChainID is the chain id of the remote chain that has seen the agent root.
	ChainID uint32 `gorm:"column:chain_id;primaryKey"`
}
