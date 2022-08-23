package types

import "github.com/ethereum/go-ethereum/common"

// Log is the contract log event.
type Log interface {
	// Address gets the address that generated the event
	Address() common.Address
	// ChainID gets the chain id of the contract that generated the event
	ChainID() uint32
	// PrimaryTopic gets the primary topic of the event, Topics[0]
	PrimaryTopic() common.Hash
	// TopicA gets the first topic of the event, Topics[1]
	TopicA() common.Hash
	// TopicB gets the second topic of the event, Topics[2]
	TopicB() common.Hash
	// TopicC gets the third topic of the event, Topics[3]
	TopicC() common.Hash
	// Data gets the data provided by the contract
	Data() []byte
	// BlockNumber gets the block in which the transaction was included
	BlockNumber() uint64
	// TxHash gets the hash of the transaction
	TxHash() common.Hash
	// TxIndex gets the index of the transaction in the block
	TxIndex() uint64
	// BlockHash gets the hash of the block in which the transaction was included
	BlockHash() common.Hash
	// Index gets the index of the log in the block
	Index() uint64
	// Removed gets true if this log was reverted due to a chain re-organization
	Removed() bool
}

type log struct {
	// address of the contract
	address common.Address
	// chain id
	chainID uint32
	// primary topic of the log, topics[0]
	primaryTopic common.Hash
	// first topic of the log, topics[1]
	topicA common.Hash
	// second topic of the log, topics[2]
	topicB common.Hash
	// third topic of the log, topics[3]
	topicC common.Hash
	// data provided by the contract
	data []byte
	// block in which the transaction was included
	blockNumber uint64
	// hash of the transaction
	txHash common.Hash
	// index of the transaction in the block
	txIndex uint64
	// hash of the block in which the transaction was included
	blockHash common.Hash
	// index of the log in the block
	index uint64
	// removed is true if this log was reverted due to a chain re-organization
	removed bool
}

// NewLog creates a new log.
func NewLog(
	address common.Address,
	chainID uint32,
	primaryTopic,
	topicA,
	topicB,
	topicC common.Hash,
	data []byte,
	blockNumber uint64,
	txHash common.Hash,
	txIndex uint64,
	blockHash common.Hash,
	index uint64,
	removed bool,
) Log {
	return log{
		address:      address,
		chainID:      chainID,
		primaryTopic: primaryTopic,
		topicA:       topicA,
		topicB:       topicB,
		topicC:       topicC,
		data:         data,
		blockNumber:  blockNumber,
		txHash:       txHash,
		txIndex:      txIndex,
		blockHash:    blockHash,
		index:        index,
		removed:      removed,
	}
}

func (l log) Address() common.Address {
	return l.address
}

func (l log) ChainID() uint32 {
	return l.chainID
}

func (l log) PrimaryTopic() common.Hash {
	return l.primaryTopic
}

func (l log) TopicA() common.Hash {
	return l.topicA
}

func (l log) TopicB() common.Hash {
	return l.topicB
}

func (l log) TopicC() common.Hash {
	return l.topicC
}

func (l log) Data() []byte {
	return l.data
}

func (l log) BlockNumber() uint64 {
	return l.blockNumber
}

func (l log) TxHash() common.Hash {
	return l.txHash
}

func (l log) TxIndex() uint64 {
	return l.txIndex
}

func (l log) BlockHash() common.Hash {
	return l.blockHash
}

func (l log) Index() uint64 {
	return l.index
}

func (l log) Removed() bool {
	return l.removed
}
