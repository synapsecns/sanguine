package types

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// MessageBodyOffset is the message body offset.
	MessageBodyOffset = MessageHeaderSize
)

// Message is an interface that contains the message.
//
//nolint:interfacebloat
type Message interface {
	// Header gets the message header
	Header() Header
	// BaseMessage is the base message if the flag indicates the type is a base message
	BaseMessage() BaseMessage
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
	header      Header
	baseMessage BaseMessage
	body        []byte
}

const headerOffset uint16 = 0

// NewMessageFromBaseMessage creates a generic message given a base message.
func NewMessageFromBaseMessage(header Header, baseMessage BaseMessage) (Message, error) {
	if header.Flag() != MessageFlagBase {
		return nil, fmt.Errorf("header flag is not base")
	}

	baseMessageBytes, err := EncodeBaseMessage(baseMessage)
	if err != nil {
		return nil, fmt.Errorf("could not encode base message: %w", err)
	}

	return &messageImpl{
		header:      header,
		baseMessage: baseMessage,
		body:        baseMessageBytes,
	}, nil
}

// NewMessageFromManagerMessage creates a generic message given a manager message.
func NewMessageFromManagerMessage(header Header, payload []byte) (Message, error) {
	if header.Flag() != MessageFlagManager {
		return nil, fmt.Errorf("header flag is not manager")
	}

	return &messageImpl{
		header:      header,
		baseMessage: nil,
		body:        payload,
	}, nil
}

// NewMessage creates a new message from fields passed in.
// TODO: Set up different initializers. One for BaseMessage and one for ManagerMessage.
func NewMessage(header Header, baseMessage BaseMessage, body []byte) Message {
	return &messageImpl{
		header:      header,
		baseMessage: baseMessage,
		body:        body,
	}
}

func (m messageImpl) Header() Header {
	return m.header
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

func (m messageImpl) BaseMessage() BaseMessage {
	return m.baseMessage
}

func (m messageImpl) Body() []byte {
	return m.body
}

// ToLeaf converts a message to an encoded leaf.
func (m messageImpl) ToLeaf() (leaf [32]byte, err error) {
	var toHash []byte
	if m.Header().Flag() == MessageFlagBase {
		leaf, err = m.BaseMessage().Leaf()
		if err != nil {
			return common.Hash{}, fmt.Errorf("could not get leaf: %w", err)
		}

		toHash = leaf[:]
	} else {
		toHash = crypto.Keccak256(m.Body())
	}

	headerLeaf, err := m.Header().Leaf()
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not get header leaf: %w", err)
	}

	rawLeaf := crypto.Keccak256(headerLeaf[:], toHash)
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
