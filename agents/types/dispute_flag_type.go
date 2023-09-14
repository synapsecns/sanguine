package types

// DisputeFlagType is the type for the Dispute Status Flag.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=DisputeFlagType -linecomment
type DisputeFlagType uint8

const (
	// DisputeFlagNone means agent is not in dispute.
	DisputeFlagNone DisputeFlagType = iota
	// DisputeFlagPending means agent is in unresolved dispute.
	DisputeFlagPending
	// DisputeFlagSlashed means agent was in dispute that lead to agent being slashed.
	DisputeFlagSlashed
)
