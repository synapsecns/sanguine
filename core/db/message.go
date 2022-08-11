package db

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/types"
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

// MessageDB stores messages.
type MessageDB interface {
	// RetrieveLatestCommittedMessageNonce gets the latest nonce of a committed message
	// returns ErrNoNonceForDomain if no nonce exists
	RetrieveLatestCommittedMessageNonce(ctx context.Context, domainID uint32) (nonce uint32, err error)
	// StoreMessageLatestBlockEnd stores the latest block end
	StoreMessageLatestBlockEnd(ctx context.Context, domainID uint32, blockNumber uint32) error
	// GetMessageLatestBlockEnd gets the message latest block
	// returns ErrNoStoredBlockForChain when not present
	GetMessageLatestBlockEnd(ctx context.Context, domainID uint32) (height uint32, err error)
	// StoreCommittedMessage stores a raw committed message
	StoreCommittedMessage(ctx context.Context, domainID uint32, message types.CommittedMessage) error
	// StoreSignedAttestations stores a signed attestation
	StoreSignedAttestations(ctx context.Context, attestation types.SignedAttestation) error
	// RetrieveSignedAttestationByNonce retrieves a signed attestation by nonce
	RetrieveSignedAttestationByNonce(ctx context.Context, domainID, nonce uint32) (attestation types.SignedAttestation, err error)
}

// MonitorDB stores event data for monitoring.
type MonitorDB interface {
	// StoreDispatchMessage stores a dispatch message
	StoreDispatchMessage(ctx context.Context, message types.Message) error
	// StoreAcceptedAttestation stores an accepted attestation
	StoreAcceptedAttestation(ctx context.Context, replicaDomain uint32, attestation types.Attestation) error
}

// SynapseDB combines db types.
type SynapseDB interface {
	MessageDB
	TxQueueDB
	MonitorDB
}
