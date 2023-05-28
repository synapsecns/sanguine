package summit

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

// Parser parses events from the summit contract.
type Parser interface {
	// ParseAttestationSaved parses a AttestationSaved event.
	ParseAttestationSaved(log ethTypes.Log) (_ []byte, ok bool)
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

// ParseAttestationSaved parses a AttesationSaved event.
func (p parserImpl) ParseAttestationSaved(log ethTypes.Log) (_ []byte, ok bool) {
	summitAttestationSaved, err := p.filterer.ParseAttestationSaved(log)
	if err != nil {
		return nil, false
	}

	if err != nil {
		return nil, false
	}

	return summitAttestationSaved.Attestation, true
}
