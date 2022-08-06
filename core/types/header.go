package types

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// Header contains information of a message.
type Header interface {
	// The origin domain of the message
	OriginDomain() uint32
	// The sender of the message
	Sender() common.Hash
	// The nonce of the message
	Nonce() uint32
	// The destination domain of the message
	DestinationDomain() uint32
	// The recipient of the message
	Recipient() common.Hash
	// The optimistic time period of the message in seconds
	OptimisticSeconds() uint32
}

type headerImpl struct {
	originDomain      uint32
	sender            common.Hash
	nonce             uint32
	destinationDomain uint32
	recipient         common.Hash
	optimisticSeconds uint32
}

// NewHeader creates a new header type.
func NewHeader(originDomain uint32, sender common.Hash, nonce uint32, destinationDomain uint32, recipient common.Hash, optimisticSeconds uint32) Header {
	return &headerImpl{
		originDomain:      originDomain,
		sender:            sender,
		nonce:             nonce,
		destinationDomain: destinationDomain,
		recipient:         recipient,
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
		sender:            encoded.Sender,
		nonce:             encoded.Nonce,
		destinationDomain: encoded.DestinationDomain,
		recipient:         encoded.Recipient,
		optimisticSeconds: encoded.OptimisticSeconds,
	}

	return decoded, nil
}

func (h headerImpl) OriginDomain() uint32 {
	return h.originDomain
}

func (h headerImpl) Sender() common.Hash {
	return h.sender
}

func (h headerImpl) Nonce() uint32 {
	return h.nonce
}

func (h headerImpl) DestinationDomain() uint32 {
	return h.destinationDomain
}

func (h headerImpl) Recipient() common.Hash {
	return h.recipient
}

func (h headerImpl) OptimisticSeconds() uint32 {
	return h.optimisticSeconds
}
