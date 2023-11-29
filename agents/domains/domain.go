package domains

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethTypes "github.com/ethereum/go-ethereum/core/types"

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
	// IsValidState checks if the given state is valid on its origin.
	IsValidState(ctx context.Context, statePayload []byte) (isValid bool, err error)
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
	// IsValidAttestation checks if the given attestation is valid on the summit
	IsValidAttestation(ctx context.Context, attestation []byte) (bool, error)
	// GetNotarySnapshot gets the snapshot payload corresponding to a given attestation
	GetNotarySnapshot(ctx context.Context, attPayload []byte) (types.Snapshot, error)
	// GetAttestation gets an attestation for a given attestationNonce.
	GetAttestation(ctx context.Context, attNonce uint32) (types.NotaryAttestation, error)
}

// InboxContract contains the interface for the inbox.
type InboxContract interface {
	// SubmitStateReportWithSnapshot reports to the inbox that a state within a snapshot is invalid.
	SubmitStateReportWithSnapshot(transactor *bind.TransactOpts, stateIndex uint8, signature signer.Signature, snapPayload []byte, snapSignature []byte) (tx *ethTypes.Transaction, err error)
	// SubmitSnapshot submits a snapshot to the inbox (via the Inbox).
	SubmitSnapshot(transactor *bind.TransactOpts, encodedSnapshot []byte, signature signer.Signature) (tx *ethTypes.Transaction, err error)
	// VerifyAttestation verifies a snapshot on the inbox.
	VerifyAttestation(transactor *bind.TransactOpts, attestation []byte, attSignature []byte) (tx *ethTypes.Transaction, err error)
	// SubmitStateReportWithAttestation submits a state report corresponding to an attesation for an invalid state.
	SubmitStateReportWithAttestation(transactor *bind.TransactOpts, stateIndex uint8, signature signer.Signature, snapPayload, attPayload, attSignature []byte) (tx *ethTypes.Transaction, err error)
	// SubmitReceipt submits a receipt to the inbox.
	SubmitReceipt(transactor *bind.TransactOpts, rcptPayload []byte, rcptSignature signer.Signature, paddedTips *big.Int, headerHash [32]byte, bodyHash [32]byte) (tx *ethTypes.Transaction, err error)
	// VerifyReceipt verifies a receipt on the inbox.
	VerifyReceipt(transactor *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (tx *ethTypes.Transaction, err error)
	// SubmitReceiptReport submits a receipt report to the inbox.
	SubmitReceiptReport(transactor *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte, rrSignature []byte) (tx *ethTypes.Transaction, err error)
}

// BondingManagerContract contains the interface for the bonding manager.
type BondingManagerContract interface {
	// AddAgent adds an agent to the bonding manager.
	AddAgent(transactor *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (tx *ethTypes.Transaction, err error)
	// GetAgentStatus returns the current agent status for the given agent.
	GetAgentStatus(ctx context.Context, address common.Address) (types.AgentStatus, error)
	// GetAgentRoot gets the current agent root
	GetAgentRoot(ctx context.Context) ([32]byte, error)
	// GetProof gets the proof that the agent is in the Agent Merkle Tree
	GetProof(ctx context.Context, address common.Address) ([][32]byte, error)
	// DisputeStatus gets the dispute status for the given agent.
	DisputeStatus(ctx context.Context, address common.Address) (disputeStatus DisputeStatus, err error)
	// GetDispute gets the dispute for a given dispute index.
	// TODO: Add more returned values here as needed.
	GetDispute(ctx context.Context, index *big.Int) (err error)
	// GetDisputeStatus gets the dispute status for the given agent.
	GetDisputeStatus(ctx context.Context, agent common.Address) (disputeStatus types.DisputeStatus, err error)
	// CompleteSlashing completes the slashing of an agent.
	CompleteSlashing(transactor *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (tx *ethTypes.Transaction, err error)
	// GetAgent gets an agent status and address for a given agent index.
	GetAgent(ctx context.Context, index *big.Int) (types.AgentStatus, common.Address, error)
}

// DestinationContract contains the interface for the destination.
type DestinationContract interface {
	// Execute executes a message on the destination.
	Execute(transactor *bind.TransactOpts, message types.Message, originProof [32][32]byte, snapshotProof [][32]byte, index uint8, gasLimit uint64) (tx *ethTypes.Transaction, err error) // AttestationsAmount retrieves the number of attestations submitted to the destination.
	// AttestationsAmount retrieves the number of attestations submitted to the destination.
	AttestationsAmount(ctx context.Context) (uint64, error)
	// GetAttestationNonce gets the nonce of the attestation by snap root
	GetAttestationNonce(ctx context.Context, snapRoot [32]byte) (uint32, error)
	// MessageStatus takes a message and returns whether it has been executed or not. 0: None, 1: Failed, 2: Success.
	MessageStatus(ctx context.Context, message types.Message) (uint8, error)
	// IsValidReceipt checks if the given receipt is valid on the destination
	IsValidReceipt(ctx context.Context, rcptPayload []byte) (bool, error)
	// PassAgentRoot passes the agent root to the destination.
	PassAgentRoot(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error)
}

// LightInboxContract contains the interface for the light inbox.
type LightInboxContract interface {
	// SubmitStateReportWithSnapshot reports to the inbox that a state within a snapshot is invalid.
	SubmitStateReportWithSnapshot(transactor *bind.TransactOpts, stateIndex uint8, signature signer.Signature, snapPayload []byte, snapSignature []byte) (tx *ethTypes.Transaction, err error)
	// SubmitAttestation submits an attestation to the destination chain (via the light inbox contract)
	SubmitAttestation(
		transactor *bind.TransactOpts,
		attPayload []byte,
		signature signer.Signature,
		agentRoot [32]byte,
		snapGas []*big.Int,
	) (tx *ethTypes.Transaction, err error)
	// SubmitStateReportWithAttestation submits a state report corresponding to an attesation for an invalid state.
	SubmitStateReportWithAttestation(transactor *bind.TransactOpts, stateIndex uint8, signature signer.Signature, snapPayload, attPayload, attSignature []byte) (tx *ethTypes.Transaction, err error)
	// VerifyStateWithSnapshot verifies a state within a snapshot.
	VerifyStateWithSnapshot(transactor *bind.TransactOpts, stateIndex uint8, snapPayload []byte, snapSignature []byte) (tx *ethTypes.Transaction, err error)
	// SubmitAttestationReport submits an attestation report to the inbox (via the light inbox contract)
	SubmitAttestationReport(transactor *bind.TransactOpts, attestation, arSignature, attSignature []byte) (tx *ethTypes.Transaction, err error)
	// VerifyStateWithAttestation verifies a state with attestation.
	VerifyStateWithAttestation(transactor *bind.TransactOpts, stateIndex uint8, snapPayload []byte, attPayload []byte, attSignature []byte) (tx *ethTypes.Transaction, err error)
	// VerifyReceipt verifies a receipt on the inbox.
	VerifyReceipt(transactor *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (tx *ethTypes.Transaction, err error)
}

// LightManagerContract contains the interface for the light manager.
type LightManagerContract interface {
	// GetAgentStatus returns the current agent status for the given agent.
	GetAgentStatus(ctx context.Context, address common.Address) (types.AgentStatus, error)
	// GetAgentRoot gets the current agent root
	GetAgentRoot(ctx context.Context) ([32]byte, error)
	// UpdateAgentStatus updates the agent status on the remote chain.
	UpdateAgentStatus(
		transactor *bind.TransactOpts,
		agentAddress common.Address,
		agentStatus types.AgentStatus,
		agentProof [][32]byte) (*ethTypes.Transaction, error)
	// GetDispute gets the dispute for a given dispute index.
	// TODO: Add more returned values here as needed.
	GetDispute(ctx context.Context, index *big.Int) (err error)
}

// TestClientContract contains the interface for the test client.
type TestClientContract interface {
	// SendMessage sends a message through the TestClient.
	SendMessage(ctx context.Context, signer signer.Signer, destination uint32, recipient common.Address, optimisticSeconds uint32, gasLimit uint64, version uint32, message []byte) error
}

// PingPongClientContract contains the interface for the ping pong test client.
type PingPongClientContract interface {
	// DoPing sends a ping message through the PingPongClient.
	DoPing(ctx context.Context, signer signer.Signer, destination uint32, recipient common.Address, pings uint16) (*ethTypes.Transaction, error)
	WatchPingSent(ctx context.Context, sink chan<- *pingpongclient.PingPongClientPingSent) (event.Subscription, error)
	WatchPongReceived(ctx context.Context, sink chan<- *pingpongclient.PingPongClientPongReceived) (event.Subscription, error)
}

// ErrNoUpdate indicates no update has been produced.
var ErrNoUpdate = errors.New("no update produced")
