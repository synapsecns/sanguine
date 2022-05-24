package types

import (
	"bytes"
	"encoding/gob"
	"github.com/ethereum/go-ethereum/common"
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
	Destination() common.Hash
	// Body is the message contents
	Body() []byte
}

// messageImpl implements a message. It is used for testing. Real messages are emitted by teh contract
type messageImpl struct {
	origin      uint32
	sender      common.Hash
	nonce       uint32
	destination common.Hash
	body        []byte
}

// NewMessage creates a new message from fields passed in
func NewMessage(origin uint32, sender common.Hash, nonce uint32, destination common.Hash, body []byte) Message {
	return &messageImpl{
		origin:      origin,
		sender:      sender,
		nonce:       nonce,
		body:        body,
		destination: destination,
	}
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

func (m messageImpl) Destination() common.Hash {
	return m.destination
}

func (m messageImpl) Body() []byte {
	return m.body
}

// Encode encodes the message to a bytes
// TODO: this should use a helper message once contract abis are ready
func (m messageImpl) Encode() ([]byte, error) {
	var res bytes.Buffer
	enc := gob.NewEncoder(&res)

	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	return res.Bytes(), nil
}

// CommittedMessage is the message that got committed.
type CommittedMessage interface {
	// LeafIndex is the index at which the message is committed
	LeafIndex() uint32
	// CommitedRoot is the current root when the message was commited.
	CommitedRoot() common.Hash
	// Message is the fully detailed message that was commited
	Message() []byte
}
