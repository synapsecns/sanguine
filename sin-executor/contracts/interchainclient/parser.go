package interchainclient

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

// EventType is the type of the module watcher
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// InterchainTransactionSentEvent is an EventType that represents an interchain transaction sent event.
	InterchainTransactionSentEvent EventType = iota + 1
	// InterchainOptionsV1Event is an EventType that represents an interchain options v1 event.
	InterchainOptionsV1Event
)

// Parser parses events from the module contract.
type Parser interface {
	// ParseEvent parses the event from the log.
	ParseEvent(log ethTypes.Log) (_ EventType, event interface{}, ok bool)
}

type parserImpl struct {
	filterer *InterchainClientV1Filterer
}

// NewParser creates a new parser for the fastbridge contract.
func NewParser(synapseModuleAddress common.Address) (Parser, error) {
	parser, err := NewInterchainClientV1Filterer(synapseModuleAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", InterchainClientV1Filterer{}, err)
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
	case InterchainTransactionSentEvent:
		event, err := p.filterer.ParseInterchainTransactionSent(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, event, true
	case InterchainOptionsV1Event:
		event, err := p.filterer.ParseInterchainOptionsV1(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, event, true
	}
	return eventType, nil, false
}
