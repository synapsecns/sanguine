package db

import (
	"context"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
)

type ConsumerDBWriter interface {
	StoreEvent(ctx context.Context, bridgeEvent bridge.EventLog, swapEvent swap.EventLog, chainID uint32) error
}

//// ConsumerBridgeDBWriter is the interface for writing bridge events to the ConsumerDB.
//type ConsumerBridgeDBWriter interface {
//	StoreEvent(ctx context.Context, deposit bridge.EventLog, chainID uint32) error
//}
//
//// ConsumerSwapDBWriter is the interface for writing swap events to the ConsumerDB.
//type ConsumerSwapDBWriter interface {
//	StoreSwapEvent(ctx context.Context, data swap.EventLog, chainID uint32) error
//}

// ConsumerBridgeDBReader is the interface for reading bridge events from the ConsumerDB.
type ConsumerBridgeDBReader interface {
}

// ConsumerSwapDBReader is the interface for reading swap events from the ConsumerDB.
type ConsumerSwapDBReader interface {
}

// ConsumerDB is the interface for the ConsumerDB.
//
//go:generate go run github.com/vektra/mockery/v2 --name=ConsumerDB --output=mocks --case=underscore
type ConsumerDB interface {
	ConsumerDBWriter
	ConsumerBridgeDBReader
	ConsumerSwapDBReader
}
