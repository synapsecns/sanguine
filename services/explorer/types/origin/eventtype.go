package origin

// EventType is the type of the swap event.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint8

const (
	// SentEvent when a new message is sent.
	SentEvent EventType = iota
)

// AllEventTypes is a list of the event types.
func AllEventTypes() []EventType {
	return []EventType{SentEvent}
}

// Int gets the int value of the event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
