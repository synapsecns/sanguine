package destination

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
	// ParseExecuted parses an Executed event.
	ParseExecuted(log ethTypes.Log) (origin *uint32, leaf *[32]byte, ok bool)
}

type parserImpl struct {
	// filterer is the parser filterer we use to parse events
	filterer *DestinationFilterer
}

// NewParser creates a new parser for the destination contract.
func NewParser(destinationAddress common.Address) (Parser, error) {
	parser, err := NewDestinationFilterer(destinationAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", DestinationFilterer{}, err)
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
	//destinationAttestationAccepted, err := p.filterer.ParseAttestationAccepted(log)
	//if err != nil {
	//	return nil, false
	//}
	//
	//attestation, err := types.DecodeSignedAttestation(destinationAttestationAccepted.Attestation)
	//if err != nil {
	//	return nil, false
	//}
	//return attestation.Attestation(), true
	return nil, false
}

// ParseExecuted parses an Executed event.
func (p parserImpl) ParseExecuted(log ethTypes.Log) (origin *uint32, leaf *[32]byte, ok bool) {
	destinationExecuted, err := p.filterer.ParseExecuted(log)
	if err != nil {
		return nil, nil, false
	}

	return &destinationExecuted.RemoteDomain, &destinationExecuted.MessageHash, true
}

// EventType is the type of the destination event
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// AttestationAcceptedEvent is an AttestationAccepted event.
	AttestationAcceptedEvent EventType = 0
	// ExecutedEvent is an Executed event.
	ExecutedEvent EventType = 1
)

// Int gets the int for an event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
