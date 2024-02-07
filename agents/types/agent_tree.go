package types

import "github.com/ethereum/go-ethereum/common"

// AgentTree is a version of the database AgentTree model with solidity-compatible types.
type AgentTree struct {
	// AgentRoot is the root of the agent tree.
	AgentRoot string
	// AgentAddress is the address of the agent for the Merkle proof.
	AgentAddress common.Address
	// AgentDomain is the domain of the agent.
	AgentDomain uint32
	// UpdatedAgentFlag is the updated agent flag corresponding to the agent tree.
	UpdatedAgentFlag AgentFlagType
	// BlockNumber is the block number that the agent tree was updated (on summit).
	BlockNumber uint64
	// Proof is the agent tree proof.
	Proof [][32]byte
}
