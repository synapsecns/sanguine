package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// BaseMessageTipsOffset is the tips offset.
	BaseMessageTipsOffset = 0
	// BaseMessageSenderOffset is the sender offset.
	BaseMessageSenderOffset = 32
	// BaseMessageRecipientOffset is the recipient offset.
	BaseMessageRecipientOffset = 64
	// BaseMessageRequestOffset is the request offset.
	BaseMessageRequestOffset = BaseMessageRecipientOffset + TipsSize
	// BaseMessageContentOffset is the content offset.
	BaseMessageContentOffset = BaseMessageRequestOffset + RequestSize
)

// BaseMessage is an interface that contains the base message.
//
//nolint:interfacebloat
type BaseMessage interface {
	// Sender address on origin chain.
	Sender() [32]byte
	// Recipient address on destination chain.
	Recipient() [32]byte
	// Tips paid on the origin chain.
	Tips() Tips
	// Request gets the message request.
	Request() Request
	// Content to be passed to recipient.
	Content() []byte

	// BodyLeaf gets the message body.
	BodyLeaf() ([]byte, error)
	// Leaf gets the message leaf.
	Leaf() ([32]byte, error)
}

// baseMessageImpl implements a base message. It is used for testutils. Real base messages are emitted by the contract.
type baseMessageImpl struct {
	sender    [32]byte
	recipient [32]byte
	tips      Tips
	request   Request
	content   []byte
}

// NewBaseMessage creates a new message from fields passed in.
func NewBaseMessage(sender, recipient [32]byte, tips Tips, request Request, content []byte) BaseMessage {
	return &baseMessageImpl{
		sender:    sender,
		recipient: recipient,
		tips:      tips,
		request:   request,
		content:   content,
	}
}

func (m baseMessageImpl) Sender() [32]byte {
	return m.sender
}

func (m baseMessageImpl) Recipient() [32]byte {
	return m.recipient
}

func (m baseMessageImpl) Tips() Tips {
	return m.tips
}

func (m baseMessageImpl) Request() Request {
	return m.request
}

func (m baseMessageImpl) Content() []byte {
	return m.content
}

func (m baseMessageImpl) BodyLeaf() ([]byte, error) {
	encodeMessage, err := EncodeBaseMessage(m)
	if err != nil {
		return nil, fmt.Errorf("failed to encode message: %w", err)
	}

	return crypto.Keccak256(encodeMessage[BaseMessageSenderOffset:]), nil
}

func (m baseMessageImpl) Leaf() ([32]byte, error) {
	bodyLeaf, err := m.BodyLeaf()
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to get body leaf: %w", err)
	}

	tipsBytes, err := EncodeTips(m.Tips())
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to encode tips: %w", err)
	}

	hashedTips := crypto.Keccak256(tipsBytes)

	return crypto.Keccak256Hash(hashedTips, bodyLeaf), nil
}
