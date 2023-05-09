package summit

// EventType is the type of the summit event.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint8

const (
	// ReceiptAccepted when a new message is sent.
	ReceiptAccepted EventType = iota
	// SnapshotAccepted when a snapshot is accepted.
	SnapshotAccepted
	// ReceiptConfirmed when a message is confirmed.
	ReceiptConfirmed
	// TipAwarded when a tip is awarded.
	TipAwarded
)

// AllEventTypes is a list of the event types.
func AllEventTypes() []EventType {
	return []EventType{ReceiptAccepted, SnapshotAccepted, ReceiptConfirmed, TipAwarded}
}

// Int gets the int value of the event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
