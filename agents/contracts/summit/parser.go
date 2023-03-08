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

// TODO (joeallen): parse events from summit
// Parser parses events from the summit contract.
type Parser interface {
	// EventType determines if an event was initiated by the bridge or the user.
	EventType(log ethTypes.Log) (_ EventType, ok bool)
	// ParseStateAccepted parses an StateAccepted event
	ParseSnapshotAccepted(log ethTypes.Log) (_ types.Snapshot, ok bool)
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

// ParseSnapshotAccepted parses an SnapshotAccepted event.
func (p parserImpl) ParseSnapshotAccepted(log ethTypes.Log) (_ types.Snapshot, ok bool) {
	summitSnapshot, err := p.filterer.ParseSnapshotAccepted(log)
	if err != nil {
		return nil, false
	}

	snapshot, err := types.DecodeSnapshot(summitSnapshot.Snapshot)
	if err != nil {
		return nil, false
	}

	return snapshot, true
}

// EventType is the type of the summit events
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// SnapshotAcceptedEvent is a SnapshotAccepted event.
	SnapshotAcceptedEvent EventType = 0
)

// Int gets the int for an event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}

// AllEventTypes contains all event types.
var AllEventTypes = []EventType{SnapshotAcceptedEvent}
