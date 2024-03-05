package types

// MessageType represents the type of a CCTP transfer.
type MessageType int

const (
	// SynapseMessageType indicates the request was generated on a SynapseCCTP contract.
	SynapseMessageType MessageType = iota + 1
	// CircleMessageType indicates that the request was generated on a native Circle contract.
	CircleMessageType
)

func (m MessageType) String() string {
	switch m {
	case SynapseMessageType:
		return "Synapse"
	case CircleMessageType:
		return "Circle"
	}
	return ""
}
