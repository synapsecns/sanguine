package lightinbox

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
)

// Parser parses events from the light inbox contract.
type Parser interface {
	// EventType determines if an event was initiated by the bridge or the user.
	EventType(log ethTypes.Log) (_ EventType, ok bool)
	// ParseAttestationAccepted parses an AttestationAccepted event
	ParseAttestationAccepted(log ethTypes.Log) (_ *types.AttestationWithMetadata, err error)
}

type parserImpl struct {
	// filterer is the parser filterer we use to parse events
	filterer *LightInboxFilterer
}

// NewParser creates a new parser for the light inbox contract.
func NewParser(lightInboxAddress common.Address) (Parser, error) {
	parser, err := NewLightInboxFilterer(lightInboxAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", LightInboxFilterer{}, err)
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
func (p parserImpl) ParseAttestationAccepted(log ethTypes.Log) (_ *types.AttestationWithMetadata, err error) {
	lightInboxAttestationAccepted, err := p.filterer.ParseAttestationAccepted(log)
	if err != nil {
		return nil, fmt.Errorf("could not parse attestation accepted: %w", err)
	}

	attestationData, err := types.NewAttestationWithMetadata(
		lightInboxAttestationAccepted.AttPayload,
		lightInboxAttestationAccepted.Domain,
		lightInboxAttestationAccepted.Notary,
		lightInboxAttestationAccepted.AttSignature,
	)
	if err != nil {
		return nil, fmt.Errorf("could not create attestation with metadata from payload: %w", err)
	}

	return attestationData, nil
}

// EventType is the type of the light inbox event
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
