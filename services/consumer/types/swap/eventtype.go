package swap

// EventType is the type of the swap event.
//
//go:generate stringer -type=EventType
type EventType uint8

const (
	// TokenSwapEvent is the token swap event.
	TokenSwapEvent EventType = iota
	// AddLiquidityEvent is the add liquidity event.
	AddLiquidityEvent
	// RemoveLiquidityEvent is the remove liquidity event.
	RemoveLiquidityEvent
	// RemoveLiquidityOneEvent is the remove liquidity one event.
	RemoveLiquidityOneEvent
	// RemoveLiquidityImbalanceEvent is the remove liquidity imbalance event.
	RemoveLiquidityImbalanceEvent
	// NewAdminFeeEvent is the new admin fee event.
	NewAdminFeeEvent
	// NewSwapFeeEvent is the new swap fee event.
	NewSwapFeeEvent
	// RampAEvent is the ramp A event.
	RampAEvent
	// StopRampAEvent is the stop ramp A event.
	StopRampAEvent
)

// AllEventTypes is a list of the event types.
func AllEventTypes() []EventType {
	return []EventType{TokenSwapEvent, AddLiquidityEvent, RemoveLiquidityEvent, RemoveLiquidityOneEvent,
		RemoveLiquidityImbalanceEvent, NewAdminFeeEvent, NewSwapFeeEvent, RampAEvent, StopRampAEvent}
}

// Int gets the int value of the event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
