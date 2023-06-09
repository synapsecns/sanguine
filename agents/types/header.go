package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

// MessageFlag indicates if the message is normal "Base" message or "Manager" message.
type MessageFlag uint8

const (
	// MessageFlagBase is the normal message that has tips.
	MessageFlagBase MessageFlag = iota
	// MessageFlagManager is manager message and will not have tips.
	MessageFlagManager
)

const (
	// MessageFlagSize if the size in bytes of a message flag.
	MessageFlagSize = 1
	// MessageHeaderSize is the size in bytes of a message header.
	MessageHeaderSize = 4*4 + MessageFlagSize
)

// Header contains information of a message.
type Header interface {
	// Flag is the flag of the message
	Flag() MessageFlag
	// OriginDomain is the origin domain of the message
	OriginDomain() uint32
	// Nonce is the nonce of the message
	Nonce() uint32
	// DestinationDomain is the destination domain of the message
	DestinationDomain() uint32
	// OptimisticSeconds is the optimistic time period of the message in seconds
	OptimisticSeconds() uint32

	// Leaf is the leaf of the header.
	Leaf() ([32]byte, error)
}

type headerImpl struct {
	flag              MessageFlag
	originDomain      uint32
	nonce             uint32
	destinationDomain uint32
	optimisticSeconds uint32
}

// NewHeader creates a new header type.
func NewHeader(flag MessageFlag, originDomain uint32, nonce uint32, destinationDomain uint32, optimisticSeconds uint32) Header {
	return &headerImpl{
		flag:              flag,
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
		flag:              encoded.Flag,
		originDomain:      encoded.OriginDomain,
		nonce:             encoded.Nonce,
		destinationDomain: encoded.DestinationDomain,
		optimisticSeconds: encoded.OptimisticSeconds,
	}

	return decoded, nil
}

func (h headerImpl) Flag() MessageFlag {
	return h.flag
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

func (h headerImpl) Leaf() ([32]byte, error) {
	var paddedHeader [32]byte
	bytesHeader, err := EncodeHeader(h)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to encode header: %w", err)
	}

	copy(paddedHeader[:], bytesHeader)
	headerHash := crypto.Keccak256(paddedHeader[:])

	fmt.Println("bytesHeader: ", paddedHeader)
	fmt.Println("len paddedfHeader: ", len(paddedHeader))

	var headerHash32Byte [32]byte
	copy(headerHash32Byte[:], headerHash)

	return headerHash32Byte, nil
}
