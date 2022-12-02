package types

import "github.com/ethereum/go-ethereum/common"

type DBMessage struct {
	// ChainID is the chain ID of the chain that the message is for.
	ChainID *uint32
	// Nonce is the nonce of the message.
	Nonce *uint32
	// Root is the root of the message.
	Root *common.Hash
	// Message is the message.
	Message *[]byte
	// Leaf is the leaf representation of the message.
	Leaf *common.Hash
	// BlockNumber is the block number of the message.
	BlockNumber *uint64
}
