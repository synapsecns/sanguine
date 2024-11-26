package fastbridge

// EventType is the type of the rfq event.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint8

const (
	// BridgeRequestedEvent is emitted when a RFQ request is broadcasted.
	BridgeRequestedEvent EventType = iota
	// BridgeRelayedEvent is emitted when a RFQ request is relayed to the destination chain.
	BridgeRelayedEvent
	// BridgeRefundedEvent is emitted when a RFQ request is refunded.
	BridgeRefundedEvent
	// BridgeProvenEvent is emitted when a RFQ request is proven.
	BridgeProvenEvent
	// BridgeClaimedEvent is emitted when a RFQ request is claimed.
	BridgeClaimedEvent
)

// AllEventTypes is a list of the event types.
func AllEventTypes() []EventType {
	return []EventType{BridgeRequestedEvent, BridgeRelayedEvent, BridgeRefundedEvent, BridgeProvenEvent, BridgeClaimedEvent}
}

// Int gets the int value of the event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
