package types

const (
	// BaseMessageSenderOffset is the sender offset.
	BaseMessageSenderOffset = 0
	// BaseMessageRecipientOffset is the recipient offset.
	BaseMessageRecipientOffset = 32
	// BaseMessageTipsOffset is the tips offset.
	BaseMessageTipsOffset = BaseMessageRecipientOffset + 32
	// BaseMessageRequestOffset is the request offset.
	BaseMessageRequestOffset = BaseMessageTipsOffset + TipsSize
	// BaseMessageContentOffset is the content offset.
	BaseMessageContentOffset = BaseMessageRequestOffset + RequestSize
)

/// | Position   | Field     | Type    | Bytes | Description                            |
/// | ---------- | --------- | ------- | ----- | -------------------------------------- |
/// | [000..032) | sender    | bytes32 | 32    | Sender address on origin chain         |
/// | [032..064) | recipient | bytes32 | 32    | Recipient address on destination chain |
/// | [064..096) | tips      | uint256 | 32    | Encoded tips paid on origin chain      |
/// | [096..116) | request   | uint160 | 20    | Encoded request for message execution  |
/// | [104..AAA) | content   | bytes   | ??    | Content to be passed to recipient      |

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
