package types

// AgentFlagType is the type for the Agent Status Flag.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=AgentFlagType -linecomment
type AgentFlagType uint8

const (
	AgentFlagUnknown AgentFlagType = iota
	AgentFlagActive
	AgentFlagUnstaking
	AgentFlagResting
	AgentFlagFraudulent
	AgentFlagSlashed
)
