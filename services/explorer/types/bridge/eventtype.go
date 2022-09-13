package bridge

// EventType is the type of the bridge event.
//
//go:generate stringer -type=EventType
type EventType uint8

const (
	// TokenDepositEvent is the token deposit event.
	TokenDepositEvent EventType = iota
	// TokenRedeemEvent is the token redeem event.
	TokenRedeemEvent
	// TokenWithdrawEvent is the token withdraw event.
	TokenWithdrawEvent
	// TokenMintEvent is the token mint event.
	TokenMintEvent
	// TokenDepositAndSwapEvent is the token deposit and swap event.
	TokenDepositAndSwapEvent
	// TokenMintAndSwapEvent is the token mint and swap event.
	TokenMintAndSwapEvent
	// TokenRedeemAndSwapEvent is the token redeem and swap event.
	TokenRedeemAndSwapEvent
	// TokenRedeemAndRemoveEvent is the token redeem and remove event.
	TokenRedeemAndRemoveEvent
	// TokenWithdrawAndRemoveEvent is the token withdraw and remove event.
	TokenWithdrawAndRemoveEvent
	// TokenRedeemV2Event is the token redeem v2 event.
	TokenRedeemV2Event
)

// AllEventTypes is a list of the event types.
func AllEventTypes() []EventType {
	return []EventType{TokenDepositEvent, TokenRedeemEvent, TokenWithdrawEvent, TokenMintEvent,
		TokenDepositAndSwapEvent, TokenMintAndSwapEvent, TokenRedeemAndSwapEvent, TokenRedeemAndRemoveEvent,
		TokenWithdrawAndRemoveEvent, TokenRedeemV2Event}
}

// Int gets the int value of the event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
