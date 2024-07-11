package l1gateway

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
	// DepositETHEvent is the event type for the DepositETH event.
	DepositETHEvent EventType = iota + 1
	// DepositERC20Event is the event type for the DepositERC20 event.
	DepositERC20Event
	// FinalizeWithdrawETHEvent is the event type for the FinalizeWithdrawETH event.
	FinalizeWithdrawETHEvent
	// FinalizeWithdrawERC20Event is the event type for the FinalizeWithdrawERC20 event.
	FinalizeWithdrawERC20Event
)

// Parser parses events from the l1gateway contracat.
type Parser interface {
	// ParseEvent parses the event from the log.
	ParseEvent(log ethTypes.Log) (_ EventType, event interface{}, ok bool)
}
type parserImpl struct {
	filterer *L1GatewayRouterFilterer
}

// NewParser creates a new parser for the l1gateway contract.
func NewParser(l1GatewayAddress common.Address) (Parser, error) {
	parser, err := NewL1GatewayRouterFilterer(l1GatewayAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", L1GatewayRouterFilterer{}, err)
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
	case DepositETHEvent:
		requested, err := p.filterer.ParseDepositETH(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, requested, true
	case DepositERC20Event:
		requested, err := p.filterer.ParseDepositERC20(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, requested, true
	case FinalizeWithdrawETHEvent:
		proven, err := p.filterer.ParseFinalizeWithdrawETH(log)
		if err != nil {
			return noOpEvent, nil, false
		}
		return eventType, proven, true
	case FinalizeWithdrawERC20Event:
		claimed, err := p.filterer.ParseFinalizeWithdrawERC20(log)
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
		DepositETHEvent:            DepositETHTopic,
		DepositERC20Event:          DepositERC20Topic,
		FinalizeWithdrawETHEvent:   FinalizeWithdrawETHTopic,
		FinalizeWithdrawERC20Event: FinalizeWithdrawERC20Topic,
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
