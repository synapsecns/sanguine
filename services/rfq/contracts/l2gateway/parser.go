package l2gateway

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

// EventType is the type of the bridge watcher
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// WithdrawETHEvent is the event type for the WithdrawETH event.
	WithdrawETHEvent EventType = iota + 2
	// WithdrawERC20Event is the event type for the WithdrawERC20 event.
	WithdrawERC20Event
	// FinalizeDepositETHEvent is the event type for the FinalizeDepositETH event.
	FinalizeDepositETHEvent
	// FinalizeDepositERC20Event is the event type for the FinalizeDepositERC20 event.
	FinalizeDepositERC20Event
)

// Parser parses events from the l2gateway contracat.
type Parser interface {
	// ParseEvent parses the event from the log.
	ParseEvent(log ethTypes.Log) (_ EventType, event interface{}, ok bool)
}
type parserImpl struct {
	filterer *L2GatewayRouterFilterer
}

// NewParser creates a new parser for the l2gateway contract.
func NewParser(l2GatewayAddress common.Address) (Parser, error) {
	parser, err := NewL2GatewayRouterFilterer(l2GatewayAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", L2GatewayRouterFilterer{}, err)
	}

	return &parserImpl{filterer: parser}, nil
}

// nolint: cyclop
func (p parserImpl) ParseEvent(log ethTypes.Log) (_ EventType, event interface{}, ok bool) {
	// return an unknown event to avoid cases where user failed to check the event type
	// make it high enough to make it obvious (we start iotas at +1, see uber style guide for details)
	noOpEvent := EventType(len(topicMap()) + 1)

	if len(log.Topics) == 0 {
		return noOpEvent, nil, false
	}
	nillableEventType := eventTypeFromTopic(log.Topics[0])
	if nillableEventType == nil {
		return noOpEvent, nil, false
	}

	eventType := *nillableEventType

	switch eventType {
	case WithdrawETHEvent:
		requested, err := p.filterer.ParseWithdrawETH(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, requested, true
	case WithdrawERC20Event:
		requested, err := p.filterer.ParseWithdrawERC20(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, requested, true
	case FinalizeDepositETHEvent:
		proven, err := p.filterer.ParseFinalizeDepositETH(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, proven, true
	case FinalizeDepositERC20Event:
		claimed, err := p.filterer.ParseFinalizeDepositERC20(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, claimed, true
	}

	return eventType, nil, true
}

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		WithdrawETHEvent:          WithdrawETHTopic,
		WithdrawERC20Event:        WithdrawERC20Topic,
		FinalizeDepositETHEvent:   FinalizeDepositETHTopic,
		FinalizeDepositERC20Event: FinalizeDepositERC20Topic,
	}
}

// eventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func eventTypeFromTopic(ogTopic common.Hash) *EventType {
	for eventType, topic := range topicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}
