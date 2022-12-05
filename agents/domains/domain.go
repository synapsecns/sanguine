package domains

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// DomainClient represents a client used for interacting
// with contracts in a particular domain. The goal of a domain is that
// the particulars of interacting with an eth vs a solana contract are abstracted
// away and can be done through a set of common interfaces.
type DomainClient interface {
	// Name gets the name of the client. This can be used for logging.
	Name() string
	// Config gets the config that was used to create the client.
	Config() config.DomainConfig
	// BlockNumber gets the latest block
	BlockNumber(ctx context.Context) (uint32, error)
	// Origin retrieves a handle for the origin contract
	Origin() OriginContract
	// AttestationCollector is the attestation collector
	AttestationCollector() AttestationCollectorContract
}

// OriginContract represents the origin contract on a particular chain.
type OriginContract interface {
	// FetchSortedMessages fetches all messages in order form lowest->highest in a given block range
	FetchSortedMessages(ctx context.Context, from uint32, to uint32) (messages []types.CommittedMessage, err error)
	// ProduceAttestation suggests an update from the origin contract
	// TODO (joe): this will be changed to "ProduceAttestations" and return an attestion per destination
	ProduceAttestation(ctx context.Context) (types.Attestation, error)
	// GetHistoricalAttestation gets the root corresponding to destination and nonce,
	// as well as the block number the message was dispatched and the current block number
	GetHistoricalAttestation(ctx context.Context, destinationID, nonce uint32) (types.Attestation, uint64, uint64, error)
}

// AttestationCollectorContract contains the interface for the attestation collector.
type AttestationCollectorContract interface {
	// AddNotary adds a notary to the attesation collector
	AddNotary(transactOpts *bind.TransactOpts, destination uint32, signer signer.Signer) error
	// SubmitAttestation submits an attestation to the attestation collector.
	SubmitAttestation(ctx context.Context, signer signer.Signer, attestation types.SignedAttestation) error
	// GetLatestNonce gets the latest nonce for the domain on the attestation collector
	GetLatestNonce(ctx context.Context, origin uint32, destination uint32, signer signer.Signer) (nonce uint32, currBlock uint64, err error)
}

// ErrNoUpdate indicates no update has been produced.
var ErrNoUpdate = errors.New("no update produced")
