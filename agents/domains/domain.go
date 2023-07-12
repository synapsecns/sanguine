package domains

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
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
	// LightManager retrieves a handle for the light manager contract
	LightManager() LightManagerContract
	// BondingManager retrieves a handle for the bonding manager contract
	BondingManager() BondingManagerContract
	// LightInbox retrieves a handle for the light inbox contract
	LightInbox() LightInboxContract
	// Inbox retrieves a handle for the inbox contract
	Inbox() InboxContract
}

// OriginContract represents the origin contract on a particular chain.
type OriginContract interface {
	// FetchSortedMessages fetches all messages in order form lowest->highest in a given block range
	FetchSortedMessages(ctx context.Context, from uint32, to uint32) (messages []types.Message, err error)
	// SuggestLatestState gets the latest state on the origin
	SuggestLatestState(ctx context.Context) (types.State, error)
	// SuggestState gets the state on the origin with the given nonce if it exists
	SuggestState(ctx context.Context, nonce uint32) (types.State, error)
}

// SummitContract contains the interface for the summit.
type SummitContract interface {
	// GetLatestState gets the latest state signed by any guard for the given origin
	GetLatestState(ctx context.Context, origin uint32) (types.State, error)
	// GetLatestAgentState gets the latest state signed by the bonded signer for the given origin
	GetLatestAgentState(ctx context.Context, origin uint32, bondedAgentSigner signer.Signer) (types.State, error)
	// GetLatestNotaryAttestation gets the latest notary attestation signed by the notary and posted on Summit.
	GetLatestNotaryAttestation(ctx context.Context, notarySigner signer.Signer) (types.NotaryAttestation, error)
	// WatchAttestationSaved looks for attesation saved events
	WatchAttestationSaved(ctx context.Context, sink chan<- *summit.SummitAttestationSaved) (event.Subscription, error)
}

// InboxContract contains the interface for the inbox.
type InboxContract interface {
	// SubmitSnapshot submits a snapshot to the inbox (via the Inbox).
	SubmitSnapshot(ctx context.Context, signer signer.Signer, encodedSnapshot []byte, signature signer.Signature) error
}

// BondingManagerContract contains the interface for the bonding manager.
type BondingManagerContract interface {
	// GetAgentStatus returns the current agent status for the given agent.
	GetAgentStatus(ctx context.Context, signer signer.Signer) (types.AgentStatus, error)
	// GetAgentRoot gets the current agent root
	GetAgentRoot(ctx context.Context) ([32]byte, error)
	// GetProof gets the proof that the agent is in the Agent Merkle Tree
	GetProof(ctx context.Context, bondedAgentSigner signer.Signer) ([][32]byte, error)
}

// DestinationContract contains the interface for the destination.
type DestinationContract interface {
	// Execute executes a message on the destination.
	Execute(ctx context.Context, signer signer.Signer, message types.Message, originProof [32][32]byte, snapshotProof [][32]byte, index *big.Int, gasLimit uint64) error
	// AttestationsAmount retrieves the number of attestations submitted to the destination.
	AttestationsAmount(ctx context.Context) (uint64, error)
	// GetAttestationNonce gets the nonce of the attestation by snap root
	GetAttestationNonce(ctx context.Context, snapRoot [32]byte) (uint32, error)
	// MessageStatus takes a message and returns whether it has been executed or not. 0: None, 1: Failed, 2: Success.
	MessageStatus(ctx context.Context, message types.Message) (uint8, error)
}

// LightInboxContract contains the interface for the light inbox.
type LightInboxContract interface {
	// SubmitAttestation submits an attestation to the destination chain (via the light inbox contract)
	SubmitAttestation(
		ctx context.Context,
		signer signer.Signer,
		attPayload []byte,
		signature signer.Signature,
		agentRoot [32]byte,
		snapGas []*big.Int,
	) error
}

// LightManagerContract contains the interface for the light manager.
type LightManagerContract interface {
	// GetAgentStatus returns the current agent status for the given agent.
	GetAgentStatus(ctx context.Context, signer signer.Signer) (types.AgentStatus, error)
	// GetAgentRoot gets the current agent root
	GetAgentRoot(ctx context.Context) ([32]byte, error)
	// UpdateAgentStatus updates the agent status on the remote chain.
	UpdateAgentStatus(
		ctx context.Context,
		unbondedSigner signer.Signer,
		bondedSigner signer.Signer,
		agentStatus types.AgentStatus,
		agentProof [][32]byte) error
}

// TestClientContract contains the interface for the test client.
type TestClientContract interface {
	// SendMessage sends a message through the TestClient.
	SendMessage(ctx context.Context, signer signer.Signer, destination uint32, recipient common.Address, optimisticSeconds uint32, gasLimit uint64, version uint32, message []byte) error
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
