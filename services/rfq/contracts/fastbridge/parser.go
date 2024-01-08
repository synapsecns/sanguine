package fastbridge

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

// EventType is the type of the bridge watcher
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// BridgeRequestedEvent is the event type for the BridgeRequested event.
	BridgeRequestedEvent EventType = iota + 1
	// BridgeRelayedEvent is the event type for the BridgeRelayed event.
	BridgeRelayedEvent
	// BridgeProofProvidedEvent is the event type for the BridgeProofProvided event.
	BridgeProofProvidedEvent
	// BridgeDepositClaimedEvent is the event type for the BridgeDepositClaimed event.
	BridgeDepositClaimedEvent
)

// Parser parses events from the fastbridge contracat.
type Parser interface {
	// ParseEvent parses the event from the log.
	ParseEvent(log ethTypes.Log) (_ EventType, event interface{}, ok bool)
}
type parserImpl struct {
	filterer *FastBridgeFilterer
}

// NewParser creates a new parser for the fastbridge contract.
func NewParser(fastBridgeAddress common.Address) (Parser, error) {
	parser, err := NewFastBridgeFilterer(fastBridgeAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", FastBridgeFilterer{}, err)
	}

	return &parserImpl{filterer: parser}, nil
}

// nolint: cyclop
func (p parserImpl) ParseEvent(log ethTypes.Log) (_ EventType, event interface{}, ok bool) {
	// return an unknown event to avoid cases where user failed to check the event type
	// make it high enough to make it obvious (we start iotas at +1, see uber style guide for details)
	noOpEvent := EventType(len(topicMap()) + 2)

	if len(log.Topics) == 0 {
		return noOpEvent, nil, false
	}
	nillableEventType := eventTypeFromTopic(log.Topics[0])
	if nillableEventType == nil {
		return noOpEvent, nil, false
	}

	eventType := *nillableEventType

	switch eventType {
	case BridgeRequestedEvent:
		requested, err := p.filterer.ParseBridgeRequested(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, requested, true
	case BridgeRelayedEvent:
		requested, err := p.filterer.ParseBridgeRelayed(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, requested, true
	case BridgeProofProvidedEvent:
		proven, err := p.filterer.ParseBridgeProofProvided(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, proven, true
	case BridgeDepositClaimedEvent:
		claimed, err := p.filterer.ParseBridgeDepositClaimed(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, claimed, true
	}

	return eventType, nil, true
}
