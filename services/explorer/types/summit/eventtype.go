package summit

// EventType is the type of the summit event.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint8

const (
	// ReceiptAcceptedEvent when a new message is sent.
	ReceiptAcceptedEvent EventType = iota
	// SnapshotAcceptedEvent when a snapshot is accepted.
	SnapshotAcceptedEvent
	// ReceiptConfirmedEvent when a message is confirmed.
	ReceiptConfirmedEvent
	// TipAwardedEvent when a tip is awarded.
	TipAwardedEvent
)

// AllEventTypes is a list of the event types.
func AllEventTypes() []EventType {
	return []EventType{ReceiptAcceptedEvent, SnapshotAcceptedEvent, ReceiptConfirmedEvent, TipAwardedEvent}
}

// Int gets the int value of the event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
