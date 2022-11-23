package messagebus

// EventType is the type of the swap event.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint8

const (
	// ExecutedEvent is the message executed (sent) event.
	ExecutedEvent EventType = iota
	// MessageSentEvent is the message sent event.
	MessageSentEvent
	// CallRevertedEvent is when a call is reverted.
	CallRevertedEvent
)

// AllEventTypes is a list of the event types.
func AllEventTypes() []EventType {
	return []EventType{ExecutedEvent, MessageSentEvent, CallRevertedEvent}
}

// Int gets the int value of the event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
