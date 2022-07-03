package db

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/types"
	"math/big"
)

// TxQueueDB contains an interface for storing transactions currently being processed.
type TxQueueDB interface {
	// StoreRawTx stores a raw transaction
	StoreRawTx(ctx context.Context, tx *ethTypes.Transaction, chainID *big.Int, from common.Address) error
	// StoreProcessedTx stores a tx that has already been processed
	StoreProcessedTx(ctx context.Context, tx *ethTypes.Transaction) error
	// GetNonceForChainID gets the latest nonce for a sender
	GetNonceForChainID(ctx context.Context, fromAddress common.Address, chainID *big.Int) (nonce uint64, err error)
}

// MessageDB contains the synapse db.
type MessageDB interface {
	// StoreCommittedMessage stores a committed message.
	StoreCommittedMessage(committedMessage types.CommittedMessage) error
	// StoreLatestMessage stores a raw committed message building off the leaf index
	StoreLatestMessage(committedMessage types.CommittedMessage) error
	// MessageByNonce retreives a raw committed message by its leaf hash.
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

	// RetrieveLatestRoot retreives latest root
	RetrieveLatestRoot() (common.Hash, error)
	// StoreLatestRoot stores the latest root
	StoreLatestRoot(latestRoot common.Hash) error

	// RetrieveProducedUpdate retrieves a produced update for a root
	RetrieveProducedUpdate(root common.Hash) (types.SignedUpdate, error)
	// StoreProducedUpdate stores a produced update
	StoreProducedUpdate(previousRoot common.Hash, update types.SignedUpdate) error
}
