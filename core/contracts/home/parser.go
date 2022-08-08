package home

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/types"
	legacyTypes "github.com/synapsecns/synapse-node/pkg/types"
)

// Parser parses events from the home contract.
type Parser interface {
	// EventType determines if an event was initiated by the bridge or the user.
	EventType(log ethTypes.Log) (_ EventType, ok bool)
	// ParseDispatch parses a dispatch event
	ParseDispatch(log ethTypes.Log) (_ types.CommittedMessage, ok bool)
}

type parserImpl struct {
	// filterer is the parser filterer we use to parse events
	filterer *HomeFilterer
}

// NewParser creates a new parser for the home contract.
func NewParser(homeAddress common.Address) (Parser, error) {
	parser, err := NewHomeFilterer(homeAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", HomeFilterer{}, err)
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
	return EventType(len(legacyTypes.AllEventTypes()) + 2), false
}

// ParseDispatch parses an update event.
func (p parserImpl) ParseDispatch(log ethTypes.Log) (_ types.CommittedMessage, ok bool) {
	dispatch, err := p.filterer.ParseDispatch(log)
	if err != nil {
		return nil, false
	}

	leafIndex := uint32(dispatch.LeafIndex.Int64())

	commitedMessage := types.NewCommittedMessage(leafIndex, dispatch.Message)

	return commitedMessage, true
}

// EventType is the type of the home event
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// DispatchEvent is a dispatch event.
	DispatchEvent EventType = 0
)

// Int gets the int for an event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
