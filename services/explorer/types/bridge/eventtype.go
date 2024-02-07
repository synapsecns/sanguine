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
	DepositEvent EventType = iota // Origin
	// RedeemEvent is the token redeem event.
	RedeemEvent // Origin
	// WithdrawEvent is the token withdraw event.
	WithdrawEvent // Destination
	// MintEvent is the token mint event.
	MintEvent // Destination
	// DepositAndSwapEvent is the token deposit and swap event.
	DepositAndSwapEvent // Origin
	// MintAndSwapEvent is the token mint and swap event.
	MintAndSwapEvent // Destination
	// RedeemAndSwapEvent is the token redeem and swap event.
	RedeemAndSwapEvent // Origin
	// RedeemAndRemoveEvent is the token redeem and remove event.
	RedeemAndRemoveEvent // Origin
	// WithdrawAndRemoveEvent is the token withdraw and remove event.
	WithdrawAndRemoveEvent // Destination
	// RedeemV2Event is the token redeem v2 event.
	RedeemV2Event // Origin
	// CircleRequestSentEvent is emitted when the origin bridge event is executed using the cctp contract.
	CircleRequestSentEvent // Origin
	// CircleRequestFulfilledEvent is emitted when the destination bridge event is executed using the cctp contract.
	CircleRequestFulfilledEvent // Destination
)

// AllEventTypes is a list of the event types.
func AllEventTypes() []EventType {
	return []EventType{DepositEvent, RedeemEvent, WithdrawEvent, MintEvent,
		DepositAndSwapEvent, MintAndSwapEvent, RedeemAndSwapEvent, RedeemAndRemoveEvent,
		WithdrawAndRemoveEvent, RedeemV2Event, CircleRequestSentEvent, CircleRequestFulfilledEvent}
}

// GetEventType gets the str/clear text event type from EventType.
//
// nolint:cyclop
func GetEventType(eventType uint8) string {
	switch eventType {
	case 0:
		return "DepositEvent"
	case 1:
		return "RedeemEvent"
	case 2:
		return "WithdrawEvent"
	case 3:
		return "MintEvent"
	case 4:
		return "DepositAndSwapEvent"
	case 5:
		return "MintAndSwapEvent"
	case 6:
		return "RedeemAndSwapEvent"
	case 7:
		return "RedeemAndRemoveEvent"
	case 8:
		return "WithdrawAndRemoveEvent"
	case 9:
		return "RedeemV2Event"
	case 10:
		return "CircleRequestSentEvent"
	case 11:
		return "CircleRequestFulfilledEvent"
	default:
		return "Unknown"
	}
}

// Int gets the int value of the event type.
func (i EventType) Int() uint8 {
	return uint8(i)
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
