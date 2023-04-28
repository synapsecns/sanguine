package lightmanager

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
)

// Parser parses events from the destination contract.
type Parser interface {
	// EventType determines if an event was initiated by the bridge or the user.
	EventType(log ethTypes.Log) (_ EventType, ok bool)
	// ParseAttestationAccepted parses an AttestationAccepted event
	ParseAttestationAccepted(log ethTypes.Log) (_ types.Attestation, ok bool)
}

type parserImpl struct {
	// filterer is the parser filterer we use to parse events
	filterer *LightManagerFilterer
}

// NewParser creates a new parser for the light manager contract.
func NewParser(lightManagerAddress common.Address) (Parser, error) {
	parser, err := NewLightManagerFilterer(lightManagerAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", LightManagerFilterer{}, err)
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

// ParseAttestationAccepted parses an AttestationAccepted event.
func (p parserImpl) ParseAttestationAccepted(log ethTypes.Log) (_ types.Attestation, ok bool) {
	lightManagerAttestationAccepted, err := p.filterer.ParseAttestationAccepted(log)
	if err != nil {
		return nil, false
	}

	attestation, err := types.DecodeAttestation(lightManagerAttestationAccepted.AttPayload)
	if err != nil {
		return nil, false
	}

	return attestation, true
}

// EventType is the type of the destination event
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// AttestationAcceptedEvent is an AttestationAccepted event.
	AttestationAcceptedEvent EventType = 0
)

// Int gets the int for an event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
