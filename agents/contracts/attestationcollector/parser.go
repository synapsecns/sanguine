package attestationcollector

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
)

// init serves as a static assertion that AllEventTypes are accounted for.
func init() {
	if len(_EventType_index)-1 != len(AllEventTypes) {
		panic("add new events to alle vent types")
	}
}

// Parser parses events from the attestation collector contract.
type Parser interface {
	// EventType determines if an event was initiated by the bridge or the user.
	EventType(log ethTypes.Log) (_ EventType, ok bool)
	// ParseAttestationAccepted parses an AttestationAccepted event
	ParseAttestationAccepted(log ethTypes.Log) (_ types.Attestation, ok bool)
}

type parserImpl struct {
	// filterer is the parser filterer we use to parse events
	filterer *AttestationCollectorFilterer
}

// NewParser creates a new parser for the attestation collector contract.
func NewParser(attestationCollectorAddress common.Address) (Parser, error) {
	parser, err := NewAttestationCollectorFilterer(attestationCollectorAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", AttestationCollectorFilterer{}, err)
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

// ParseAttestationAccepted parses an AttestationAccepted event.
func (p parserImpl) ParseAttestationAccepted(log ethTypes.Log) (_ types.Attestation, ok bool) {
	// TODO (joeallen): FIX ME
	// attestationCollectorAttestation, err := p.filterer.ParseAttestationAccepted(log)
	// if err != nil {
	//	return nil, false
	//}

	// attestation, err := types.DecodeAttestation(attestationCollectorAttestation.Attestation)
	// if err != nil {
	//	return nil, false
	//}

	// return attestation, true
	return nil, false
}

// EventType is the type of the attestation collector events
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// AttestationAcceptedEvent is a AttestationAccepted event.
	AttestationAcceptedEvent EventType = 0
)

// Int gets the int for an event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}

// AllEventTypes contains all event types.
var AllEventTypes = []EventType{AttestationAcceptedEvent}
