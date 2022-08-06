package types

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// Header contains information of a message.
type Header interface {
	// Version gets the version of the header
	Version() uint16
	// OriginDomain is the origin domain of the message
	OriginDomain() uint32
	// Sender is the sender of the message
	Sender() common.Hash
	// Nonce is the nonce of the message
	Nonce() uint32
	// DestinationDomain is the destination domain of the message
	DestinationDomain() uint32
	// Recipient is the recipient of the message
	Recipient() common.Hash
	// OptimisticSeconds is the optimistic time period of the message in seconds
	OptimisticSeconds() uint32
}

type headerImpl struct {
	version           uint16
	originDomain      uint32
	sender            common.Hash
	nonce             uint32
	destinationDomain uint32
	recipient         common.Hash
	optimisticSeconds uint32
}

func (h headerImpl) Version() uint16 {
	return h.version
}

const headerVersion uint16 = 1

// NewHeader creates a new header type.
func NewHeader(originDomain uint32, sender common.Hash, nonce uint32, destinationDomain uint32, recipient common.Hash, optimisticSeconds uint32) Header {
	return &headerImpl{
		version:           headerVersion,
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
		version:           encoded.Version,
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
