package cctp

// EventType is the type of the cctp event.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint8

const (
	// CircleRequestSentEvent is emitted when a Circle token is sent with an attached action request.
	CircleRequestSentEvent EventType = iota
	// CircleRequestFulfilledEvent is emitted when a Circle token is received with an attached action request.
	CircleRequestFulfilledEvent
)

// AllEventTypes is a list of the event types.
func AllEventTypes() []EventType {
	return []EventType{CircleRequestSentEvent, CircleRequestFulfilledEvent}
}

// Int gets the int value of the event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
