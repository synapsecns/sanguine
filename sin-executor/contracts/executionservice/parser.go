package executionservice

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
	// ExecutionRequestedEvent is an EventType that represents an execution requested event.
	ExecutionRequestedEvent EventType = iota + 1
)

// Parser parses events from the module contract.
type Parser interface {
	// ParseEvent parses the event from the log.
	ParseEvent(log ethTypes.Log) (_ EventType, event interface{}, ok bool)
}

type parserImpl struct {
	filterer *ExecutionServiceFilterer
}

// NewParser creates a new parser for the fastbridge contract.
func NewParser(synapseModuleAddress common.Address) (Parser, error) {
	parser, err := NewExecutionServiceFilterer(synapseModuleAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", ExecutionServiceFilterer{}, err)
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
	case ExecutionRequestedEvent:
		event, err := p.filterer.ParseExecutionRequested(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, event, true
	}
	return eventType, nil, false
}
