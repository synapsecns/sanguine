package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Header contains information of a message.
type Header interface {
	// OriginDomain is the origin domain of the message
	OriginDomain() uint32
	// Nonce is the nonce of the message
	Nonce() uint32
	// DestinationDomain is the destination domain of the message
	DestinationDomain() uint32
	// OptimisticSeconds is the optimistic time period of the message in seconds
	OptimisticSeconds() uint32
}

type headerImpl struct {
	originDomain      uint32
	nonce             uint32
	destinationDomain uint32
	optimisticSeconds uint32
}

// NewHeader creates a new header type.
func NewHeader(originDomain uint32, nonce uint32, destinationDomain uint32, optimisticSeconds uint32) Header {
	return &headerImpl{
		originDomain:      originDomain,
		nonce:             nonce,
		destinationDomain: destinationDomain,
		optimisticSeconds: optimisticSeconds,
	}
}

// DecodeHeader decodes a header from a byte slice.
func DecodeHeader(header []byte) (Header, error) {
	reader := bytes.NewReader(header)

	var encoded headerEncoder

	err := binary.Read(reader, binary.BigEndian, &encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode header: %w", err)
	}

	decoded := headerImpl{
		originDomain:      encoded.OriginDomain,
		nonce:             encoded.Nonce,
		destinationDomain: encoded.DestinationDomain,
		optimisticSeconds: encoded.OptimisticSeconds,
	}

	return decoded, nil
}

func (h headerImpl) OriginDomain() uint32 {
	return h.originDomain
}

func (h headerImpl) Nonce() uint32 {
	return h.nonce
}

func (h headerImpl) DestinationDomain() uint32 {
	return h.destinationDomain
}

func (h headerImpl) OptimisticSeconds() uint32 {
	return h.optimisticSeconds
}
