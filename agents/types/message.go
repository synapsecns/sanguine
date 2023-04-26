package types

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Message is an interface that contains metadata.
//
//nolint:interfacebloat
type Message interface {
	// Header gets the message header
	Header() Header
	// Tips gets the tips
	Tips() Tips
	// Body gets the message body
	Body() []byte

	// OriginDomain returns the Slip-44 ID
	OriginDomain() uint32
	// Nonce is the count of all previous messages to the destination
	Nonce() uint32
	// DestinationDomain is the slip-44 id of the destination
	DestinationDomain() uint32
	// ToLeaf converts a leaf to a keccac256
	ToLeaf() (leaf [32]byte, err error)
	// OptimisticSeconds gets the optimistic seconds count
	OptimisticSeconds() uint32
}

// messageImpl implements a message. It is used for testutils. Real messages are emitted by the contract.
type messageImpl struct {
	version uint16
	header  Header
	tips    Tips
	body    []byte
}

const headerOffset uint16 = 0

// NewMessage creates a new message from fields passed in.
func NewMessage(header Header, tips Tips, body []byte) Message {
	return &messageImpl{
		header: header,
		tips:   tips,
		body:   body,
	}
}

func (m messageImpl) Header() Header {
	return m.header
}

func (m messageImpl) Tips() Tips {
	return m.tips
}

// DecodeMessage decodes a message from a byte slice.
func DecodeMessage(message []byte) (Message, error) {
	reader := bytes.NewReader(message)

	var encoded messageEncoder

	err := binary.Read(reader, binary.BigEndian, &encoded)
	if err != nil {
		return nil, fmt.Errorf("could not parse encoded: %w", err)
	}

	rawHeader := message[headerOffset : encoded.HeaderLength+headerOffset]

	header, err := DecodeHeader(rawHeader)
	if err != nil {
		return nil, fmt.Errorf("could not decode header: %w", err)
	}

	rawTips := message[headerOffset+encoded.HeaderLength : headerOffset+encoded.HeaderLength+encoded.TipsLength]
	unmarshalledTips, err := DecodeTips(rawTips)
	if err != nil {
		return nil, fmt.Errorf("could not decode unmarshalledTips: %w", err)
	}

	dataSize := binary.Size(encoded)

	// make sure we can get the body of the message
	if dataSize > len(message) {
		return nil, fmt.Errorf("message too small, expected at least %d, got %d", dataSize, len(message))
	}

	rawBody := message[headerOffset+encoded.HeaderLength+encoded.TipsLength:]

	decoded := messageImpl{
		version: encoded.Version,
		body:    rawBody,
		header:  header,
		tips:    unmarshalledTips,
	}

	return decoded, nil
}

func (m messageImpl) Version() uint16 {
	return m.version
}

func (m messageImpl) OriginDomain() uint32 {
	return m.Header().OriginDomain()
}

func (m messageImpl) Nonce() uint32 {
	return m.Header().Nonce()
}

func (m messageImpl) DestinationDomain() uint32 {
	return m.Header().DestinationDomain()
}

func (m messageImpl) OptimisticSeconds() uint32 {
	return m.Header().OptimisticSeconds()
}

func (m messageImpl) Body() []byte {
	return m.body
}

// ToLeaf converts a message to an encoded leaf.
func (m messageImpl) ToLeaf() (leaf [32]byte, err error) {
	encodedMessage, err := EncodeMessage(m)
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not encode message: %w", err)
	}

	rawLeaf := crypto.Keccak256(encodedMessage)
	copy(leaf[:], rawLeaf)
	return leaf, nil
}

// CommittedMessage is the message that got committed.
type CommittedMessage interface {
	// Message is the fully detailed message that was committed
	Message() []byte
	// Leaf gets a leaf
	Leaf() [32]byte
	// Encode encodes a message
	Encode() ([]byte, error)
}

// NewCommittedMessage creates a new committed message.
func NewCommittedMessage(message []byte) CommittedMessage {
	return committedMessageImpl{
		message: message,
	}
}

// commitedMessageImpl contains the implementation of a committed message.
type committedMessageImpl struct {
	committedRoot common.Hash
	message       []byte
}

// CommittedMessageEncoder is used to export fields for struct encoding.
type CommittedMessageEncoder struct {
	CommittedRoot common.Hash
	Message       []byte
}

// DecodeCommittedMessage decodes a committed message into a struct.
func DecodeCommittedMessage(rawMessage []byte) (CommittedMessage, error) {
	dec := gob.NewDecoder(bytes.NewReader(rawMessage))

	var msg CommittedMessageEncoder
	err := dec.Decode(&msg)
	if err != nil {
		return nil, fmt.Errorf("could not decode message: %w", err)
	}

	decoded := committedMessageImpl{
		committedRoot: msg.CommittedRoot,
		message:       msg.Message,
	}

	return decoded, nil
}

// Encode encodes a committed message.
// Deprecated: will be removed.
func (c committedMessageImpl) Encode() ([]byte, error) {
	newCommittedMessage := CommittedMessageEncoder{
		CommittedRoot: c.committedRoot,
		Message:       c.message,
	}

	var res bytes.Buffer
	enc := gob.NewEncoder(&res)

	err := enc.Encode(newCommittedMessage)
	if err != nil {
		return nil, fmt.Errorf("could not encode %T: %w", newCommittedMessage, err)
	}
	return res.Bytes(), nil
}

// Leaf gets the leaf.
func (c committedMessageImpl) Leaf() (leaf [32]byte) {
	rawLeaf := crypto.Keccak256(c.message)
	copy(leaf[:], rawLeaf)

	return leaf
}

// Message gets the message.
func (c committedMessageImpl) Message() []byte {
	return c.message
}
