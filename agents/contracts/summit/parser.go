package summit

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
)

// init serves as a static assertion that AllEventTypes are accounted for.
func init() {
	if len(_EventType_index)-1 != len(AllEventTypes) {
		panic("add new events to all event types")
	}
}

// Parser parses events from the summit contract.
type Parser interface {
	// EventType is the event type.
	EventType(log ethTypes.Log) (_ EventType, ok bool)
	// ParseSnapshotAccepted parses a SnapshotAccepted event.
	ParseSnapshotAccepted(log ethTypes.Log) (_ types.Snapshot, domain uint32, ok bool)
}

type parserImpl struct {
	// filterer is the parser filterer we use to parse events
	filterer *SummitFilterer
}

// NewParser creates a new parser for the summit contract.
func NewParser(summitAddress common.Address) (Parser, error) {
	parser, err := NewSummitFilterer(summitAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", SummitFilterer{}, err)
	}

	return &parserImpl{filterer: parser}, nil
}

func (p parserImpl) EventType(log ethTypes.Log) (_ EventType, ok bool) {
	for _, logTopic := range log.Topics {
		eventType := eventTypeFromTopic(logTopic)
		if eventType == nil {
			continue
		}

		return *eventType, true
	}
	// return an unknown event to avoid cases where user failed to check the event type
	return EventType(len(AllEventTypes) + 2), false
}

// ParseSnapshotAccepted parses a SnapshotAccepted event.
func (p parserImpl) ParseSnapshotAccepted(log ethTypes.Log) (_ types.Snapshot, domain uint32, ok bool) {
	snapshotAccepted, err := p.filterer.ParseSnapshotAccepted(log)
	if err != nil {
		return nil, 0, false
	}

	snapshot, err := types.DecodeSnapshot(snapshotAccepted.Snapshot)
	if err != nil {
		return nil, 0, false
	}

	return snapshot, snapshotAccepted.Domain, true
}

// EventType is the type of the attestation collector events
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// SnapshotAcceptedEvent is a SnapshotAccepted event.
	SnapshotAcceptedEvent EventType = iota
)

// Int gets the int for an event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}

// AllEventTypes contains all event types.
var AllEventTypes = []EventType{SnapshotAcceptedEvent}
