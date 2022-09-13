package bridge

// EventType is the type of the bridge event.
//
//go:generate stringer -type=EventType
type EventType uint8

const (
	// DepositEvent is the token deposit event.
	DepositEvent EventType = iota
	// RedeemEvent is the token redeem event.
	RedeemEvent
	// WithdrawEvent is the token withdraw event.
	WithdrawEvent
	// MintEvent is the token mint event.
	MintEvent
	// DepositAndSwapEvent is the token deposit and swap event.
	DepositAndSwapEvent
	// MintAndSwapEvent is the token mint and swap event.
	MintAndSwapEvent
	// RedeemAndSwapEvent is the token redeem and swap event.
	RedeemAndSwapEvent
	// RedeemAndRemoveEvent is the token redeem and remove event.
	RedeemAndRemoveEvent
	// WithdrawAndRemoveEvent is the token withdraw and remove event.
	WithdrawAndRemoveEvent
	// RedeemV2Event is the token redeem v2 event.
	RedeemV2Event
)

// AllEventTypes is a list of the event types.
func AllEventTypes() []EventType {
	return []EventType{DepositEvent, RedeemEvent, WithdrawEvent, MintEvent,
		DepositAndSwapEvent, MintAndSwapEvent, RedeemAndSwapEvent, RedeemAndRemoveEvent,
		WithdrawAndRemoveEvent, RedeemV2Event}
}

// Int gets the int value of the event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
