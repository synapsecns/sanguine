package bridge

import (
	"database/sql/driver"
	"fmt"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

// EventType is the type of the bridge event.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
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

// BridgeInitiated determines whether or not the event type is initiated by the bridge
// (as opposed to the user).
func (i EventType) BridgeInitiated() bool {
	switch i {
	case DepositEvent, RedeemEvent, RedeemAndRemoveEvent, DepositAndSwapEvent, RedeemAndSwapEvent, RedeemV2Event:
		return false
	case WithdrawEvent, MintEvent, MintAndSwapEvent, WithdrawAndRemoveEvent:
		return true
	}
	panic("unknown event")
}

// GormDataType gets the data type to use for gorm.
func (i EventType) GormDataType() string {
	return dbcommon.EnumDataType
}

// Scan gets the type to insert into the db.
func (i *EventType) Scan(src interface{}) error {
	res, err := dbcommon.EnumScan(src)
	if err != nil {
		return fmt.Errorf("could not scan value: %w", err)
	}
	newEventType := EventType(res)
	*i = newEventType
	return nil
}

// Value gets the value to use for the db.
func (i EventType) Value() (driver.Value, error) {
	drvr, err := dbcommon.EnumValue(i)
	if err != nil {
		return nil, fmt.Errorf("could not get value: %w", err)
	}
	return drvr, nil
}
