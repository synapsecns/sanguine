package origin

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
)

// Parser parses events from the origin contract.
type Parser interface {
	// EventType determines if an event was initiated by the bridge or the user.
	EventType(log ethTypes.Log) (_ EventType, ok bool)
	// ParseDispatched parses a dispatched event
	ParseDispatched(log ethTypes.Log) (_ types.CommittedMessage, ok bool)
}

type parserImpl struct {
	// filterer is the parser filterer we use to parse events
	filterer *OriginFilterer
}

// NewParser creates a new parser for the origin contract.
func NewParser(originAddress common.Address) (Parser, error) {
	parser, err := NewOriginFilterer(originAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", OriginFilterer{}, err)
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
	return EventType(len(topicMap()) + 2), false
}

// ParseDispatched parses an update event.
func (p parserImpl) ParseDispatched(log ethTypes.Log) (_ types.CommittedMessage, ok bool) {
	dispatched, err := p.filterer.ParseDispatched(log)
	if err != nil {
		return nil, false
	}

	commitedMessage := types.NewCommittedMessage(dispatched.Message)

	return commitedMessage, true
}

// EventType is the type of the origin event
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// DispatchedEvent is a dispatched event.
	DispatchedEvent EventType = 0
)

// Int gets the int for an event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
