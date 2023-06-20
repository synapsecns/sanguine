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
	// AgentFlagFraudulent means agent has been accused of fraud and is awaiting fraud resolution.
	AgentFlagFraudulent
	// AgentFlagSlashed means that the agent has been slashed and will not be a part of agent set in the next epoch.
	AgentFlagSlashed
)
