package domains

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
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
	// Summit is the summit
	Summit() SummitContract
	// Destination retrieves a handle for the destination contract
	Destination() DestinationContract
}

// OriginContract represents the origin contract on a particular chain.
type OriginContract interface {
	// FetchSortedMessages fetches all messages in order form lowest->highest in a given block range
	FetchSortedMessages(ctx context.Context, from uint32, to uint32) (messages []types.CommittedMessage, err error)
	// ProduceAttestation suggests an update from the origin contract
	// TODO (joe): this will be changed to "ProduceAttestations" and return an attestation per destination
	ProduceAttestation(ctx context.Context) (types.Attestation, error)
	// GetHistoricalAttestation gets the root corresponding to destination and nonce,
	// as well as the block number the message was dispatched and the current block number
	GetHistoricalAttestation(ctx context.Context, destinationID, nonce uint32) (types.Attestation, uint64, error)
	// SuggestLatestState gets the latest state on the origin
	SuggestLatestState(ctx context.Context) (types.State, error)
	// SuggestState gets the state on the origin with the given nonce if it exists
	SuggestState(ctx context.Context, nonce uint32) (types.State, error)
}

// SummitContract contains the interface for the summit.
type SummitContract interface {
	// AddAgent adds an agent (guard or notary) to the summit
	AddAgent(transactOpts *bind.TransactOpts, domain uint32, signer signer.Signer) error
	// SubmitSnapshot submits a snapshot to the summit.
	SubmitSnapshot(ctx context.Context, signer signer.Signer, encodedSnapshot []byte, signature signer.Signature) error
	// GetLatestState gets the latest state signed by any guard for the given origin
	GetLatestState(ctx context.Context, origin uint32) (types.State, error)
	// GetLatestAgentState gets the latest state signed by the bonded signer for the given origin
	GetLatestAgentState(ctx context.Context, origin uint32, bondedAgentSigner signer.Signer) (types.State, error)
}

// DestinationContract contains the interface for the destination.
type DestinationContract interface {
	// SubmitAttestation submits an attestation to the destination.
	SubmitAttestation(ctx context.Context, signer signer.Signer, attestation types.SignedAttestation) error
	// Execute executes a message on the destination.
	Execute(ctx context.Context, signer signer.Signer, message types.Message, proof [32][32]byte, index *big.Int) error
	// SubmittedAt retrieves the time a given Merkle root from the given origin was submitted on the destination.
	SubmittedAt(ctx context.Context, origin uint32, root [32]byte) (*time.Time, error)
}

// TestClientContract contains the interface for the test client.
type TestClientContract interface {
	// SendMessage sends a message through the TestClient.
	SendMessage(ctx context.Context, signer signer.Signer, destination uint32, recipient common.Address, optimisticSeconds uint32, message []byte) error
}

// PingPongClientContract contains the interface for the ping pong test client.
type PingPongClientContract interface {
	// DoPing sends a ping message through the PingPongClient.
	DoPing(ctx context.Context, signer signer.Signer, destination uint32, recipient common.Address, pings uint16) error
	WatchPingSent(ctx context.Context, sink chan<- *pingpongclient.PingPongClientPingSent) (event.Subscription, error)
	WatchPongReceived(ctx context.Context, sink chan<- *pingpongclient.PingPongClientPongReceived) (event.Subscription, error)
}

// ErrNoUpdate indicates no update has been produced.
var ErrNoUpdate = errors.New("no update produced")
