package types

// AgentFlagType is the type for the Agent Status Flag.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=AgentFlagType -linecomment
type AgentFlagType uint8

const (
	// AgentFlagUnknown means agent is not part of agent set.
	AgentFlagUnknown AgentFlagType = iota
	// AgentFlagActive means agent is active.
	AgentFlagActive
	// AgentFlagUnstaking means agent is unstaking.
	AgentFlagUnstaking
	// AgentFlagResting means agent has staked but not currently doing any work.
	AgentFlagResting
	// AgentFlagFraudulent means agent has been found to be fraudulent but has not yet been slashed.
	AgentFlagFraudulent
	// AgentFlagSlashed means that the agent was found fraudulent and has been slashed.
	AgentFlagSlashed
)
