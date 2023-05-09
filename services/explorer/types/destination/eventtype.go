package destination

// EventType is the type of the summit event.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint8

const (
	// AttestationAcceptedEvent when a snapshot is accepted by the Destination contract.
	AttestationAcceptedEvent EventType = iota
	// AgentRootAcceptedEvent agent root accepted.
	AgentRootAcceptedEvent
)

// AllEventTypes is a list of the event types.
func AllEventTypes() []EventType {
	return []EventType{AttestationAcceptedEvent, AgentRootAcceptedEvent}
}

// Int gets the int value of the event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
