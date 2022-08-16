package attestationcollector

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/types"
	legacyTypes "github.com/synapsecns/synapse-node/pkg/types"
)

// Parser parses events from the attestation collector contract.
type Parser interface {
	// EventType determines if an event was initiated by the bridge or the user.
	EventType(log ethTypes.Log) (_ EventType, ok bool)
	// ParseAttestationSubmitted parses an AttestationSubmitted event
	ParseAttestationSubmitted(log ethTypes.Log) (_ types.AttestationSubmitted, ok bool)
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
	return EventType(len(legacyTypes.AllEventTypes()) + 2), false
}

// ParseAttestationSubmitted parses an AttestationSubmitted event.
func (p parserImpl) ParseAttestationSubmitted(log ethTypes.Log) (_ types.AttestationSubmitted, ok bool) {
	attestation, err := p.filterer.ParseAttestationSubmitted(log)
	if err != nil {
		return nil, false
	}

	attestationSubmitted := types.NewAttestationSubmitted(
		attestation.Notary.Hash(),
		attestation.Attestation,
	)

	return attestationSubmitted, true
}

// EventType is the type of the attestation collector events
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType uint

const (
	// AttestationSubmittedEvent is a AttestationSubmitted event.
	AttestationSubmittedEvent EventType = 0
)

// Int gets the int for an event type.
func (i EventType) Int() uint8 {
	return uint8(i)
}
