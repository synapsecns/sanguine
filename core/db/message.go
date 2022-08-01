package db

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/types"
	"math/big"
)

// TxQueueDB contains an interface for storing transactions currently being processed.
//go:generate go run github.com/vektra/mockery/v2 --name TxQueueDB --output ./mocks --case=underscore
type TxQueueDB interface {
	// StoreRawTx stores a raw transaction
	StoreRawTx(ctx context.Context, tx *ethTypes.Transaction, chainID *big.Int, from common.Address) error
	// StoreProcessedTx stores a tx that has already been processed
	StoreProcessedTx(ctx context.Context, tx *ethTypes.Transaction) error
	// GetNonceForChainID gets the latest nonce for a sender
	GetNonceForChainID(ctx context.Context, fromAddress common.Address, chainID *big.Int) (nonce uint64, err error)
}

type NewMessageDB interface {
	// RetrieveLatestNonce gets the latest nonce of a committed message
	// returns ErrNoNonceForDomain if no nonce exists
	RetrieveLatestNonce(ctx context.Context, domainID uint32) (nonce uint32, err error)
	// StoreMessageLatestBlockEnd stores the latest block end
	StoreMessageLatestBlockEnd(ctx context.Context, domainID uint32, blockNumber uint32) error
	// GetMessageLatestBlockEnd gets the message latest block
	// returns ErrNoStoredBlockForChain when not precent
	GetMessageLatestBlockEnd(ctx context.Context, domainID uint32) (height uint32, err error)
	// StoreCommittedMessage stores a raw committed message building off the leaf index
	StoreCommittedMessage(ctx context.Context, domainID uint32, message types.CommittedMessage) error
}

type SynapseDB interface {
	NewMessageDB
	TxQueueDB
}

// MessageDB contains the synapse db.
type MessageDB interface {
	// StoreCommittedMessage stores a committed message.
	StoreCommittedMessage(committedMessage types.CommittedMessage) error
	// StoreLatestMessage stores a raw committed message building off the leaf index
	StoreLatestMessage(committedMessage types.CommittedMessage) error
	// MessageByNonce retrieves a raw committed message by its leaf hash.
	MessageByNonce(destination, nonce uint32) (types.CommittedMessage, error)
	// MessageByLeaf fetches a message by leaf
	MessageByLeaf(leaf common.Hash) (types.CommittedMessage, error)
	// MessageByLeafIndex fetches a message by leaf the index of it's leaf
	MessageByLeafIndex(leafIndex uint32) (types.CommittedMessage, error)
	// StoreProof stores a proof of the lead index
	StoreProof(leafIndex uint32, proof types.Proof) error
	// ProofByLeafIndex gets a proof by it's leaf index
	ProofByLeafIndex(leafIndex uint32) (types.Proof, error)

	// StoreIndexedHeight stores the indexed height
	StoreIndexedHeight(domain string, height uint32) error
	// GetIndexedHeight gets the indexed height for a domain
	GetIndexedHeight(domain string) (uint32, error)
	// UpdateLatestLeafIndex sets the latest leaf
	UpdateLatestLeafIndex(leafIndex uint32) error
	// StoreMessageLatestBlockEnd stores the latest block end
	StoreMessageLatestBlockEnd(blockNumber uint32) error
	// GetMessageLatestBlockEnd gets the message latest block
	GetMessageLatestBlockEnd() (height uint32, err error)

	// RetrieveLatestRoot retrieves latest root
	RetrieveLatestRoot() (common.Hash, error)
	// StoreLatestRoot stores the latest root
	StoreLatestRoot(latestRoot common.Hash) error
}
