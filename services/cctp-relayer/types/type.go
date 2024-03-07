package types

// MessageType represents the type of a CCTP transfer.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=MessageType
type MessageType uint

const (
	// SynapseMessageType indicates the request was generated on a SynapseCCTP contract.
	SynapseMessageType MessageType = iota + 1
	// CircleMessageType indicates that the request was generated on a native Circle contract.
	CircleMessageType
)
