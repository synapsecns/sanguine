package destination

// EventType is the type of the summit event.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint8

const (
	// AttestationAccepted when a snapshot is accepted by the Destination contract.
	AttestationAccepted EventType = iota
	// AgentRootAccepted agent root accepted.
	AgentRootAccepted
)

// AllEventTypes is a list of the event types.
func AllEventTypes() []EventType {
	return []EventType{AttestationAccepted, AgentRootAccepted}
}

// Int gets the int value of the event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
