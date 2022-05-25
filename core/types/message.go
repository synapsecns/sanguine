package types

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Message is an interface that contains metadata.
type Message interface {
	// Origin returns the Slip-44 ID
	Origin() uint32
	// Sender is the address of the sender
	Sender() common.Hash
	// Nonce is the count of all previous messages to the destination
	Nonce() uint32
	// Destination is the slip-44 id of the destination
	Destination() uint32
	// Recipient is the address of the recipient
	Recipient() common.Hash
	// Body is the message contents
	Body() []byte
	// Encode encodes a message
	Encode() ([]byte, error)
	// ToLeaf converts a leaf to a keccac256
	ToLeaf() (leaf [32]byte, err error)
	// DestinationAndNonce gets the destination and nonce encoded into a single field
	DestinationAndNonce() uint64
}

// messageImpl implements a message. It is used for testing. Real messages are emitted by the contract.
type messageImpl struct {
	origin      uint32
	sender      common.Hash
	nonce       uint32
	destination uint32
	recipient   common.Hash
	body        []byte
}

// NewMessage creates a new message from fields passed in.
func NewMessage(origin uint32, sender common.Hash, nonce uint32, destination uint32, body []byte, recipient common.Hash) Message {
	return &messageImpl{
		origin:      origin,
		sender:      sender,
		nonce:       nonce,
		body:        body,
		destination: destination,
		recipient:   recipient,
	}
}

// DecodeMessage decodes a message from a byte slice.
func DecodeMessage(message []byte) (Message, error) {
	dec := gob.NewDecoder(bytes.NewReader(message))

	var msg MessageEncoder
	err := dec.Decode(&msg)
	if err != nil {
		return nil, fmt.Errorf("could not decode message: %w", err)
	}

	decoded := messageImpl{
		origin:      msg.Origin,
		sender:      msg.Sender,
		nonce:       msg.Nonce,
		destination: msg.Destination,
		body:        msg.Body,
	}

	return decoded, nil
}

func (m messageImpl) Origin() uint32 {
	return m.origin
}

func (m messageImpl) Sender() common.Hash {
	return m.sender
}

func (m messageImpl) Nonce() uint32 {
	return m.nonce
}

func (m messageImpl) Destination() uint32 {
	return m.destination
}

func (m messageImpl) Recipient() common.Hash {
	return m.recipient
}

// DestinationAndNonce gets the destination and nonce encoded in a single field
// TODO: statically assert 32 bit fields here.
func (m messageImpl) DestinationAndNonce() uint64 {
	return CombineDestinationAndNonce(m.destination, m.nonce)
}

// CombineDestinationAndNonce combines a destination and nonce.
func CombineDestinationAndNonce(destination, nonce uint32) uint64 {
	dest := uint64(destination)
	return ((dest) << 32) | uint64(nonce)
}

func (m messageImpl) Body() []byte {
	return m.body
}

// MessageEncoder is used for message marshaling/unmarshalling. It is exported
// for encoding only.
type MessageEncoder struct {
	Origin      uint32
	Sender      common.Hash
	Nonce       uint32
	Destination uint32
	Body        []byte
	Recipient   common.Hash
}

// Encode encodes the message to a bytes
// TODO: this should use a helper message once contract abis are ready.
func (m messageImpl) Encode() ([]byte, error) {
	newMessage := MessageEncoder{
		Origin:      m.origin,
		Sender:      m.sender,
		Nonce:       m.nonce,
		Destination: m.destination,
		Body:        m.body,
	}

	var res bytes.Buffer
	enc := gob.NewEncoder(&res)

	err := enc.Encode(newMessage)
	if err != nil {
		return nil, fmt.Errorf("could not encode %T: %w", m, err)
	}
	return res.Bytes(), nil
}

// ToLeaf converts a message to an encoded leaf.
func (m messageImpl) ToLeaf() (leaf [32]byte, err error) {
	encodedMessage, err := m.Encode()
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not encode message: %w", err)
	}

	rawLeaf := crypto.Keccak256(encodedMessage)
	copy(leaf[:], rawLeaf)
	return leaf, nil
}

// CommittedMessage is the message that got committed.
type CommittedMessage interface {
	// LeafIndex is the index at which the message is committed
	LeafIndex() uint32
	// CommitedRoot is the current root when the message was committed.
	CommitedRoot() common.Hash
	// Message is the fully detailed message that was committed
	Message() []byte
	// Leaf gets a leaf
	Leaf() [32]byte
	// Encode encodes a message
	Encode() ([]byte, error)
}

// NewCommittedMessage creates a new committed message.
func NewCommittedMessage(leafIndex uint32, committedRoot common.Hash, message []byte) CommittedMessage {
	return committedMessageImpl{
		leafIndex:     leafIndex,
		committedRoot: committedRoot,
		message:       message,
	}
}

// commitedMessageImpl contains the implementation of a committed message.
type committedMessageImpl struct {
	leafIndex     uint32
	committedRoot common.Hash
	message       []byte
}

// CommittedMessageEncoder is used to export fields for struct encoding.
type CommittedMessageEncoder struct {
	LeafIndex     uint32
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
		leafIndex:     msg.LeafIndex,
		committedRoot: msg.CommittedRoot,
		message:       msg.Message,
	}

	return decoded, nil
}

// Encode encodes a committed message.
func (c committedMessageImpl) Encode() ([]byte, error) {
	newCommittedMessage := CommittedMessageEncoder{
		LeafIndex:     c.leafIndex,
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

// LeafIndex gets the index of the leaf.
func (c committedMessageImpl) LeafIndex() uint32 {
	return c.leafIndex
}

// CommitedRoot gets the root.
func (c committedMessageImpl) CommitedRoot() common.Hash {
	return c.committedRoot
}

// Message gets the message.
func (c committedMessageImpl) Message() []byte {
	return c.message
}
