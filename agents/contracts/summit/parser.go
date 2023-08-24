package summit

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
)

// Parser parses events from the summit contract.
type Parser interface {
	// EventType determines the event type.
	EventType(log ethTypes.Log) (_ EventType, ok bool)
	// ParseAttestationSaved parses a AttestationSaved event.
	ParseAttestationSaved(log ethTypes.Log) (_ types.Attestation, ok bool)
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
	return EventType(len(topicMap()) + 2), false
}

// ParseAttestationSaved parses a AttesationSaved event.
func (p parserImpl) ParseAttestationSaved(log ethTypes.Log) (_ types.Attestation, ok bool) {
	summitAttestationSaved, err := p.filterer.ParseAttestationSaved(log)
	if err != nil {
		return nil, false
	}

	attestation, err := types.DecodeAttestation(summitAttestationSaved.Attestation)
	if err != nil {
		return nil, false
	}

	return attestation, true
}

// EventType is the type of the summit event
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// AttestationSavedEvent is an AttestationAccepted event.
	AttestationSavedEvent EventType = 0
)

// Int gets the int for an event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
