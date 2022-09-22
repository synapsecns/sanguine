package db

import (
	"context"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
)

// ConsumerDBWriter is the interface for writing to the ConsumerDB.
type ConsumerDBWriter interface {
	// StoreEvent stores an event.
	StoreEvent(ctx context.Context, bridgeEvent bridge.EventLog, swapEvent swap.EventLog, chainID uint32, tokenID *string) error
}

// ConsumerDBReader is the interface for reading events from the ConsumerDB.
type ConsumerDBReader interface {
	// ReadBlockNumberByChainID reads an event from the database by chainID.
	ReadBlockNumberByChainID(ctx context.Context, eventType int8, chainID uint32) (*uint64, error)
}

// ConsumerDB is the interface for the ConsumerDB.
//
//go:generate go run github.com/vektra/mockery/v2 --name=ConsumerDB --output=mocks --case=underscore
type ConsumerDB interface {
	ConsumerDBWriter
	ConsumerDBReader
}
