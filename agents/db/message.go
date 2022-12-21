package db

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
)

// TxQueueDB contains an interface for storing transactions currently being processed.
//
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
// nolint
// TODO (joe): This needs to be refactored after we do the GlobalRegistry stuff
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
	StoreAcceptedAttestation(ctx context.Context, attestation types.Attestation) error
	// GetDelinquentMessage gets messages that were sent, but never received
	GetDelinquentMessages(ctx context.Context, destinationDomain uint32) ([]types.Message, error)
}

// InProgressAttestationDB stores in-progress attesations.
// nolint
type InProgressAttestationDB interface {
	// RetrieveLatestCachedNonce gets the latest nonce cached for a particular origin-destination pair
	// returns ErrNoNonceForDomain if no nonce exists.
	RetrieveLatestCachedNonce(ctx context.Context, originID, destinationID uint32) (nonce uint32, err error)
	// StoreNewInProgressAttestation stores a in-progress attestation only if it hasn't already been stored
	StoreNewInProgressAttestation(ctx context.Context, attestation types.Attestation, originDispathBlockNumber uint64) error
	// UpdateSignature sets the signature of the in-progress Attestation.
	UpdateSignature(ctx context.Context, inProgressAttestation types.InProgressAttestation) error
	// UpdateSubmittedToAttestationCollectorTime sets the time attestation was sent to Attestation Collector.
	UpdateSubmittedToAttestationCollectorTime(ctx context.Context, inProgressAttestation types.InProgressAttestation) error
	// MarkConfirmedOnAttestationCollector confirms that we posted the signed attestation on the Attestation Collector.
	MarkConfirmedOnAttestationCollector(ctx context.Context, inProgressAttestation types.InProgressAttestation) error
	// RetrieveInProgressAttestation retrieves an in-progress attestation by <origin, destination, nonce>.
	RetrieveInProgressAttestation(ctx context.Context, originID, destinationID, nonce uint32) (inProgressAttestation types.InProgressAttestation, err error)
	// RetrieveOldestUnsignedInProgressAttestation retrieves the oldest in-progress attestation that has not yet been signed.
	RetrieveOldestUnsignedInProgressAttestation(ctx context.Context, originID, destinationID uint32) (inProgressAttestation types.InProgressAttestation, err error)
	// RetrieveOldestUnsubmittedSignedInProgressAttestation retrieves the oldest in-progress attestation that has been signed but not yet submitted.
	RetrieveOldestUnsubmittedSignedInProgressAttestation(ctx context.Context, originID, destinationID uint32) (inProgressAttestation types.InProgressAttestation, err error)
	// RetrieveOldestUnconfirmedSubmittedInProgressAttestation retrieves the oldest in-progress attestation that has been signed and submitted but not yet confirmed on the AttestationCollector.
	RetrieveOldestUnconfirmedSubmittedInProgressAttestation(ctx context.Context, originID, destinationID uint32) (inProgressAttestation types.InProgressAttestation, err error)
	// RetrieveNewestConfirmedInProgressAttestation retrieves the newest in-progress attestation that has been confirmed on the AttestationCollector.
	RetrieveNewestConfirmedInProgressAttestation(ctx context.Context, originID, destinationID uint32) (inProgressAttestation types.InProgressAttestation, err error)
}

// SynapseDB combines db types.
type SynapseDB interface {
	MessageDB
	TxQueueDB
	MonitorDB
	InProgressAttestationDB
}
