package types

// AgentStatusRelayedState represents the state of a RelayableAgentStatus model.
type AgentStatusRelayedState uint8

const (
	// Queued is when an agent status has been updated on Summit, but has not been relayed to the remote chain.
	Queued AgentStatusRelayedState = iota
	// Relayed is when the agent status has been relayed to the remote chain.
	Relayed
)
