package types

import "math/big"

// Update represents an update emitted by a contract in a domain.
type Update interface {
	// GetMessageHash gets the hash of the message
	GetMessageHash() [32]byte
	// GetLeafIndex gets the leaf index
	GetLeafIndex() *big.Int
	// GetDestinationAndNonce gets the destination and nonce of a message
	GetDestinationAndNonce() uint64
	// GetCommittedRoot gets the committed root of the merkle tree
	GetCommittedRoot() [32]byte
	// GetMessage gets the message of in the dispatched update
	GetMessage() []byte
}
