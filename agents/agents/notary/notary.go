package notary

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/agents/notary/db"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

// Notary checks the Summit for that latest states signed by guards, validates those states on origin,
// then signs and submits the snapshot to Summit.
type Notary struct {
	bondedSigner            signer.Signer
	unbondedSigner          signer.Signer
	ownerSigner             signer.Signer
	domains                 []domains.DomainClient
	summitDomain            domains.DomainClient
	destinationDomain       domains.DomainClient
	refreshInterval         time.Duration
	summitMyLatestStates    map[uint32]types.State
	summitGuardLatestStates map[uint32]types.State
	currentSnapRoot         [32]byte
	summitParser            summit.Parser
	handler                 metrics.Handler
	retryConfig             []retry.WithBackoffConfigurator
	txSubmitter             submitter.TransactionSubmitter
}

// NewNotary creates a new notary.
//
//nolint:cyclop
func NewNotary(ctx context.Context, cfg config.AgentConfig, omniRPCClient omnirpcClient.RPCClient, txDB db.NotaryDB, handler metrics.Handler) (_ Notary, err error) {
	notary := Notary{
		refreshInterval: time.Second * time.Duration(cfg.RefreshIntervalSeconds),
	}
	notary.domains = []domains.DomainClient{}

	notary.bondedSigner, err = signerConfig.SignerFromConfig(ctx, cfg.BondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error with bondedSigner, could not create notary: %w", err)
	}

	notary.unbondedSigner, err = signerConfig.SignerFromConfig(ctx, cfg.UnbondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error with unbondedSigner, could not create notary: %w", err)
	}

	// ownerSigner is optional
	notary.ownerSigner, _ = signerConfig.SignerFromConfig(ctx, cfg.OwnerSigner)

	for domainName, domain := range cfg.Domains {
		var domainClient domains.DomainClient

		chainRPCURL := omniRPCClient.GetEndpoint(int(domain.DomainID), 1)
		domainClient, err = evm.NewEVM(ctx, domainName, domain, chainRPCURL)
		if err != nil {
			return Notary{}, fmt.Errorf("failing to create evm for domain, could not create notary for: %w", err)
		}
		notary.domains = append(notary.domains, domainClient)
		if domain.DomainID == cfg.SummitDomainID {
			notary.summitDomain = domainClient
		}
		if domain.DomainID == cfg.DomainID {
			notary.destinationDomain = domainClient
		}
	}

	notary.summitMyLatestStates = make(map[uint32]types.State, len(notary.domains))
	notary.summitGuardLatestStates = make(map[uint32]types.State, len(notary.domains))

	notary.summitParser, err = summit.NewParser(common.HexToAddress(notary.summitDomain.Config().SummitAddress))
	if err != nil {
		return Notary{}, fmt.Errorf("could not create summit parser: %w", err)
	}

	notary.handler = handler
	if cfg.MaxRetrySeconds == 0 {
		cfg.MaxRetrySeconds = 30
	}

	notary.retryConfig = []retry.WithBackoffConfigurator{
		retry.WithMaxAttemptTime(time.Second * time.Duration(cfg.MaxRetrySeconds)),
	}

	notary.txSubmitter = submitter.NewTransactionSubmitter(handler, notary.unbondedSigner, omniRPCClient, txDB.SubmitterDB(), &cfg.SubmitterConfig)

	return notary, nil
}

// ensureNotaryActive returns without error if the notary is active on the Summit.
// If Unknown, the notary is registered on the Summit.
// Otherwise, an error is returned.
// TODO: should do this for all agents
func (n *Notary) ensureNotaryActive(parentCtx context.Context) (err error) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "ensureNotaryActive", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(n.destinationDomain.Config().DomainID)),
		attribute.String("agent", n.bondedSigner.Address().String()),
	))
	defer metrics.EndSpanWithErr(span, err)

	var agentStatus types.AgentStatus
	contractCall := func(ctx context.Context) (err error) {
		agentStatus, err = n.summitDomain.BondingManager().GetAgentStatus(ctx, n.bondedSigner.Address())
		if err != nil {
			return fmt.Errorf("could not get agent status: %w", err)
		}
		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, n.retryConfig...)
	if err != nil {
		return fmt.Errorf("could not get agent status: %w", err)
	}
	span.AddEvent("got agent status", trace.WithAttributes(
		attribute.String("agentStatus", agentStatus.Flag().String()),
	))

	if agentStatus.Flag() == types.AgentFlagActive {
		return nil
	} else if agentStatus.Flag() == types.AgentFlagUnknown {
		return n.addAgent(ctx)
	}

	return fmt.Errorf("notary is not active on summit")
}

// addAgent calls addAgent on the BondingManager after fetching agent proof.
func (n *Notary) addAgent(parentCtx context.Context) (err error) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "addAgent")
	defer metrics.EndSpanWithErr(span, err)

	// make sure we have access to the owner
	if n.ownerSigner == nil {
		return fmt.Errorf("cannot add agent without owner signer")
	}

	// fetch the agent proof
	var proof [][32]byte
	contractCall := func(ctx context.Context) (err error) {
		proof, err = n.summitDomain.BondingManager().GetProof(ctx, n.bondedSigner.Address())
		if err != nil {
			return fmt.Errorf("could not get agent proof: %w", err)
		}
		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, n.retryConfig...)
	if err != nil {
		return err
	}
	span.AddEvent("got agent proof")

	// add the agent; we don't use submitter for now because of onlyOwner constraint
	transactor, err := n.ownerSigner.GetTransactor(ctx, big.NewInt(int64(n.summitDomain.Config().DomainID)))
	if err != nil {
		return fmt.Errorf("could not get owner transactor: %w", err)
	}
	tx, err := n.summitDomain.BondingManager().AddAgent(transactor, n.destinationDomain.Config().DomainID, n.bondedSigner.Address(), proof)
	if err != nil {
		return fmt.Errorf("could not add agent: %w", err)
	}
	span.AddEvent("submitted addAgent() tx", trace.WithAttributes(
		attribute.String("tx", tx.Hash().Hex()),
	))
	types.LogTx("NOTARY", "Called addAgent()", n.summitDomain.Config().DomainID, tx)
	return nil
}

//nolint:cyclop
func (n *Notary) loadSummitMyLatestStates(parentCtx context.Context) {
	for _, domain := range n.domains {
		ctx, span := n.handler.Tracer().Start(parentCtx, "loadSummitMyLatestStates", trace.WithAttributes(
			attribute.Int(metrics.ChainID, int(domain.Config().DomainID)),
		))

		originID := domain.Config().DomainID
		if n.destinationDomain.Config().DomainID == originID {
			continue
		}
		myLatestState, err := n.summitDomain.Summit().GetLatestAgentState(ctx, originID, n.bondedSigner)
		if err != nil {
			myLatestState = nil
			span.AddEvent("GetLatestAgentState failed", trace.WithAttributes(
				attribute.String("err", err.Error()),
			))
			continue
		}
		if myLatestState != nil && myLatestState.Nonce() > uint32(0) {
			n.summitMyLatestStates[originID] = myLatestState
			span.AddEvent("Got my summit latest state", trace.WithAttributes(
				attribute.Int("nonce", int(myLatestState.Nonce())),
				attribute.Int("originID", int(originID)),
			))
		}

		span.End()
	}
}

//nolint:cyclop
func (n *Notary) loadSummitGuardLatestStates(parentCtx context.Context) {
	for _, d := range n.domains {
		domain := d
		ctx, span := n.handler.Tracer().Start(parentCtx, "loadSummitGuardLatestStates", trace.WithAttributes(
			attribute.Int(metrics.ChainID, int(domain.Config().DomainID)),
		))

		originID := domain.Config().DomainID

		// TODO: Wrap this with a retry loop if we deviate from the current `Start` behavior.
		guardLatestState, err := n.summitDomain.Summit().GetLatestState(ctx, originID)
		if err != nil {
			guardLatestState = nil
			span.AddEvent("GetLatestState failed", trace.WithAttributes(
				attribute.String("err", err.Error()),
			))
		}
		if guardLatestState != nil && guardLatestState.Nonce() > uint32(0) {
			n.summitGuardLatestStates[originID] = guardLatestState
			span.AddEvent("Got guard latest state", trace.WithAttributes(
				attribute.Int("nonce", int(guardLatestState.Nonce())),
				attribute.Int("originID", int(originID)),
			))
		}

		span.End()
	}
}

//nolint:cyclop
func (n *Notary) loadLatestSummitAttestation(parentCtx context.Context) (types.NotaryAttestation, error) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "loadNotaryLatestAttestation", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(n.destinationDomain.Config().DomainID)),
		attribute.String("snapRoot", common.BytesToHash(n.currentSnapRoot[:]).String()),
	))
	defer span.End()

	// Fetch the attestation nonce corresponding to the current snapRoot.
	var attNonce uint32
	contractCall := func(ctx context.Context) (err error) {
		attNonce, err = n.summitDomain.Destination().GetAttestationNonce(ctx, n.currentSnapRoot)
		if err != nil {
			return fmt.Errorf("could not get attestation nonce: %w", err)
		}
		return nil
	}
	err := retry.WithBackoff(ctx, contractCall, n.retryConfig...)
	if err != nil {
		span.AddEvent("GetLatestNotaryAttestation failed", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
		return nil, err
	}
	span.AddEvent("got attestation nonce", trace.WithAttributes(
		attribute.Int("nonce", int(attNonce)),
	))

	// Fetch the attestation and corresponding metadata for the attestation nonce.
	var attestation types.NotaryAttestation
	contractCall = func(ctx context.Context) (err error) {
		attestation, err = n.summitDomain.Summit().GetAttestation(ctx, uint32(attNonce))
		if err != nil {
			return fmt.Errorf("could not get attestation: %w", err)
		}
		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, n.retryConfig...)
	if err != nil {
		span.AddEvent("GetLatestNotaryAttestation failed", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
		return nil, err
	}

	if attNonce == 0 {
		return nil, nil
	}

	types.LogTx("NOTARY", fmt.Sprintf("Loaded attestation with nonce %d, snapshotRoot %s", attNonce, common.BytesToHash(n.currentSnapRoot[:]).String()), n.destinationDomain.Config().DomainID, nil)
	return attestation, nil
}

func (n *Notary) isAlreadySubmitted(parentCtx context.Context, attestation types.NotaryAttestation) (bool, error) {
	snapRoot := attestation.Attestation().SnapshotRoot()
	ctx, span := n.handler.Tracer().Start(parentCtx, "isAlreadySubmitted", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(n.destinationDomain.Config().DomainID)),
		attribute.String("snapRoot", common.BytesToHash(snapRoot[:]).String()),
	))
	defer span.End()

	// Fetch the attestation nonce corresponding to the given snapRoot.
	var attNonce uint32
	contractCall := func(ctx context.Context) (err error) {
		attNonce, err = n.destinationDomain.Destination().GetAttestationNonce(ctx, snapRoot)
		if err != nil {
			return fmt.Errorf("could not get attestation nonce: %w", err)
		}
		return nil
	}
	err := retry.WithBackoff(ctx, contractCall, n.retryConfig...)
	if err != nil {
		span.AddEvent("isAlreadySubmitted failed", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
		return false, err
	}
	span.AddEvent("got attestation nonce", trace.WithAttributes(
		attribute.Int("attNonce", int(attNonce)),
	))

	return attNonce > 0, nil
}

func (n *Notary) shouldNotaryRegisteredOnDestination(parentCtx context.Context) (bool, bool) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "shouldNotaryRegisteredOnDestination", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(n.destinationDomain.Config().DomainID)),
	))
	defer span.End()
	var bondingManagerAgentRoot [32]byte
	contractCall := func(ctx context.Context) (err error) {
		bondingManagerAgentRoot, err = n.summitDomain.BondingManager().GetAgentRoot(ctx)
		if err != nil {
			return fmt.Errorf("could not get agent root: %w", err)
		}

		return nil
	}
	err := retry.WithBackoff(ctx, contractCall, n.retryConfig...)
	if err != nil {
		span.AddEvent("GetAgentRoot failed on bonding manager", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))

		return false, false
	}

	var destinationLightManagerAgentRoot [32]byte
	contractCall = func(ctx context.Context) (err error) {
		destinationLightManagerAgentRoot, err = n.destinationDomain.LightManager().GetAgentRoot(ctx)
		if err != nil {
			fmt.Printf("could not get agent root: %f\n", err)
			return fmt.Errorf("could not get agent root: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, n.retryConfig...)
	if err != nil {
		span.AddEvent("GetAgentRoot failed on destination light manager", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))

		return false, false
	}

	if bondingManagerAgentRoot != destinationLightManagerAgentRoot {
		// We need to wait until destination has same agent root as the synapse chain.
		return false, false
	}

	var agentStatus types.AgentStatus
	contractCall = func(ctx context.Context) (err error) {
		agentStatus, err = n.destinationDomain.LightManager().GetAgentStatus(ctx, n.bondedSigner.Address())
		if err != nil {
			return fmt.Errorf("could not get agent status: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, n.retryConfig...)
	if err != nil {
		span.AddEvent("GetAgentStatus failed", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))

		return false, false
	}

	if agentStatus.Flag() == types.AgentFlagUnknown {
		// Here we want to add the Notary and proceed with sending to destination
		return true, true
	} else if agentStatus.Flag() == types.AgentFlagActive {
		// Here we already added the Notary and can proceed with sending to destination
		return false, true
	}
	return false, false
}

//nolint:cyclop
func (n *Notary) isValidOnOrigin(parentCtx context.Context, state types.State, domain domains.DomainClient) bool {
	if state == nil {
		return false
	}

	stateRoot := state.Root()
	ctx, span := n.handler.Tracer().Start(parentCtx, "isValidOnOrigin", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(domain.Config().DomainID)),
		attribute.Int(metrics.Nonce, int(state.Nonce())),
		attribute.String("stateRoot", common.Bytes2Hex(stateRoot[:])),
	))

	defer span.End()

	var stateOnOrigin types.State
	contractCall := func(ctx context.Context) (err error) {
		stateOnOrigin, err = domain.Origin().SuggestState(ctx, state.Nonce())
		if err != nil {
			return fmt.Errorf("could not suggest state: %w", err)
		}

		return nil
	}
	err := retry.WithBackoff(ctx, contractCall, n.retryConfig...)
	if err != nil {
		span.AddEvent("SuggestState failed", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
		// return false since we weren't able to validate the state on the origin
		return false
	}

	if stateOnOrigin.Root() != state.Root() {
		span.AddEvent("State roots do not equal")
		return false
	}

	if stateOnOrigin.Origin() != state.Origin() {
		span.AddEvent("State origins do not equal")
		return false
	}

	if stateOnOrigin.Nonce() != state.Nonce() {
		span.AddEvent("State nonces do not equal")
		return false
	}

	if stateOnOrigin.BlockNumber() == nil {
		span.AddEvent("State on origin had nil block number")
		return false
	}

	if state.BlockNumber() == nil {
		span.AddEvent("State to validate had nil block number")
		return false
	}

	if stateOnOrigin.BlockNumber().Uint64() != state.BlockNumber().Uint64() {
		span.AddEvent("State block numbers do not equal")
		return false
	}

	if stateOnOrigin.Timestamp() == nil {
		span.AddEvent("State on origin had nil time stamp")
		return false
	}

	if state.Timestamp() == nil {
		span.AddEvent("State to validate had nil time stamp")
		return false
	}

	if stateOnOrigin.Timestamp().Uint64() != state.Timestamp().Uint64() {
		span.AddEvent("State timestamps do not equal")
		return false
	}

	stateOnOriginHash, err := stateOnOrigin.Hash()
	if err != nil {
		span.AddEvent("Error computing state on origin hash")
		return false
	}

	stateHash, err := state.Hash()
	if err != nil {
		span.AddEvent("Error computing state on summit hash")
		return false
	}

	if stateOnOriginHash != stateHash {
		span.AddEvent("State hashes do not equal")
		return false
	}

	return true
}

//nolint:cyclop
func (n *Notary) getLatestSnapshot(parentCtx context.Context) (types.Snapshot, map[uint32]types.State) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "getLatestSnapshot")
	defer span.End()

	statesToSubmit := make(map[uint32]types.State, len(n.domains))
	for _, domain := range n.domains {
		_, stateSpan := n.handler.Tracer().Start(ctx, "loading state", trace.WithAttributes(
			attribute.Int(metrics.ChainID, int(domain.Config().DomainID)),
		))

		originID := domain.Config().DomainID
		summitMyLatest, ok := n.summitMyLatestStates[originID]
		if !ok || summitMyLatest == nil || summitMyLatest.Nonce() == 0 {
			summitMyLatest = nil
		}
		summitGuardLatest, ok := n.summitGuardLatestStates[originID]
		if !ok || summitGuardLatest == nil || summitGuardLatest.Nonce() == 0 {
			continue
		}

		if summitMyLatest != nil && summitMyLatest.Nonce() >= summitGuardLatest.Nonce() {
			// Here this notary already submitted this state
			continue
		}
		if !n.isValidOnOrigin(ctx, summitGuardLatest, domain) {
			stateSpan.AddEvent("state not valid on origin", trace.WithAttributes(
				attribute.Int("nonce", int(summitGuardLatest.Nonce())),
			))
			continue
		}
		statesToSubmit[originID] = summitGuardLatest
		stateSpan.AddEvent("registering guard's latest summit state", trace.WithAttributes(
			attribute.Int("originID", int(originID)),
			attribute.Int("nonce", int(summitGuardLatest.Nonce())),
		))

		stateSpan.End()
	}
	snapshotStates := make([]types.State, 0, len(statesToSubmit))
	for _, state := range statesToSubmit {
		if state.Nonce() == 0 {
			continue
		}
		snapshotStates = append(snapshotStates, state)
	}
	if len(snapshotStates) > 0 {
		snapshot := types.NewSnapshot(snapshotStates)
		snapRoot, _, _ := snapshot.SnapshotRootAndProofs()
		span.AddEvent("got snapshot", trace.WithAttributes(
			attribute.String("snapRoot", common.BytesToHash(snapRoot[:]).String()),
		))
		return snapshot, statesToSubmit
	}
	//nolint:nilnil
	return nil, nil
}

//nolint:cyclop
func (n *Notary) submitLatestSnapshot(parentCtx context.Context) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "submitLatestSnapshot", trace.WithAttributes(
		attribute.String("snapRoot", common.BytesToHash(n.currentSnapRoot[:]).String()),
		attribute.Int("domain", int(n.destinationDomain.Config().DomainID)),
	))
	defer span.End()

	snapshot, statesToSubmit := n.getLatestSnapshot(ctx)
	if snapshot == nil {
		span.AddEvent("no snapshot to submit")
		return
	}

	n.currentSnapRoot, _, _ = snapshot.SnapshotRootAndProofs()
	attNonce, err := n.summitDomain.Destination().GetAttestationNonce(ctx, n.currentSnapRoot)
	if err != nil {
		span.AddEvent(fmt.Sprintf("Could not get attestation nonce for snapRoot: %s", common.BytesToHash(n.currentSnapRoot[:])), trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
		return
	}
	span.AddEvent("got attestation nonce", trace.WithAttributes(
		attribute.Int("attNonce", int(attNonce)),
	))

	// if the snapshot root has a corresponding attestation, no need to submit it
	if attNonce > 0 {
		return
	}

	snapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(ctx, n.bondedSigner)
	//nolint:nestif
	if err != nil {
		span.AddEvent("Error signing snapshot", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
	} else {
		snapshotRoot, _, _ := snapshot.SnapshotRootAndProofs()
		span.AddEvent("submitting snapshot", trace.WithAttributes(
			attribute.Int("numStates", len(statesToSubmit)),
			attribute.String("snapRoot", common.BytesToHash(snapshotRoot[:]).String()),
		))
		_, err := n.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(n.summitDomain.Config().DomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = n.summitDomain.Inbox().SubmitSnapshot(transactor, encodedSnapshot, snapshotSignature)
			if err != nil {
				return nil, fmt.Errorf("could not submit snapshot: %w", err)
			}
			span.AddEvent("Submitted snapshot tx", trace.WithAttributes(
				attribute.String("tx", tx.Hash().Hex()),
			))
			types.LogTx("NOTARY", fmt.Sprintf("Submitted snapshot with snapRoot: %v", common.BytesToHash(n.currentSnapRoot[:]).String()), n.summitDomain.Config().DomainID, tx)
			return
		})
		if err != nil {
			span.AddEvent("Error submitting snapshot", trace.WithAttributes(
				attribute.String("err", err.Error()),
			))
			return
		}
		// update our view of summit states
		for originID, state := range statesToSubmit {
			n.summitMyLatestStates[originID] = state
		}
	}
}

//nolint:cyclop
func (n *Notary) registerNotaryOnDestination(parentCtx context.Context) bool {
	ctx, span := n.handler.Tracer().Start(parentCtx, "registerNotaryOnDestination", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(n.destinationDomain.Config().DomainID)),
	))
	defer span.End()

	var agentProof [][32]byte
	contractCall := func(ctx context.Context) (err error) {
		agentProof, err = n.summitDomain.BondingManager().GetProof(ctx, n.bondedSigner.Address())
		if err != nil {
			return fmt.Errorf("could not get agent proof: %w", err)
		}

		return nil
	}
	err := retry.WithBackoff(ctx, contractCall, n.retryConfig...)
	if err != nil {
		span.AddEvent("GetProof on bonding manager failed", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))

		return false
	}

	var agentStatus types.AgentStatus
	contractCall = func(ctx context.Context) (err error) {
		agentStatus, err = n.summitDomain.BondingManager().GetAgentStatus(ctx, n.bondedSigner.Address())
		if err != nil {
			return fmt.Errorf("could not get agent status: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, n.retryConfig...)
	if err != nil {
		span.AddEvent("GetAgentStatus on bonding manager failed", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))

		return false
	}
	span.AddEvent("Dispatching notary registration to submitter", trace.WithAttributes(
		attribute.String("agentStatus", agentStatus.Flag().String()),
		attribute.String("agent", n.bondedSigner.Address().String()),
	))
	_, err = n.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(n.destinationDomain.Config().DomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
		tx, err = n.destinationDomain.LightManager().UpdateAgentStatus(
			transactor,
			n.bondedSigner.Address(),
			agentStatus,
			agentProof,
		)
		if err != nil {
			return nil, fmt.Errorf("could not update agent status: %w", err)
		}
		if tx != nil {
			span.AddEvent("Submitted notary registration tx", trace.WithAttributes(
				attribute.String("tx", tx.Hash().Hex()),
			))
		}

		return
	})
	if err != nil {
		span.AddEvent("Error updating agent status", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
		return false
	}
	return true
}

//nolint:cyclop,unused
func (n *Notary) submitAttestation(parentCtx context.Context) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "submitAttestation", trace.WithAttributes(
		attribute.String("snapRoot", common.BytesToHash(n.currentSnapRoot[:]).String()),
		attribute.Int("domain", int(n.destinationDomain.Config().DomainID)),
	))
	defer span.End()

	attestation, err := n.loadLatestSummitAttestation(ctx)
	if err != nil {
		logger.Warnf("Error loading latest summit attestation: %v\n", err)
		return
	}
	if attestation == nil {
		span.AddEvent("no attestation to submit")
		return
	}

	alreadySubmitted, err := n.isAlreadySubmitted(ctx, attestation)
	if err != nil {
		logger.Warnf("Error checking if attestation already submitted: %v\n", err)
		return
	}
	if alreadySubmitted {
		snapRoot := attestation.Attestation().SnapshotRoot()
		span.AddEvent("attestation already submitted on destination", trace.WithAttributes(
			attribute.String("snapRoot", common.BytesToHash(snapRoot[:]).String()),
		))
		return
	}

	attestationSignature, _, _, err := attestation.Attestation().SignAttestation(ctx, n.bondedSigner, true)
	if err != nil {
		logger.Errorf("Error signing attestation: %v", err)
		span.AddEvent("Error signing attestation", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
	} else {
		snapRoot := attestation.Attestation().SnapshotRoot()
		span.AddEvent("dispatching attestation to submitter", trace.WithAttributes(
			attribute.String("snapRoot", common.BytesToHash(snapRoot[:]).String()),
			attribute.Int("attNonce", int(attestation.Attestation().Nonce())),
		))
		_, err = n.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(n.destinationDomain.Config().DomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = n.destinationDomain.LightInbox().SubmitAttestation(
				transactor,
				attestation.AttPayload(),
				attestationSignature,
				attestation.AgentRoot(),
				attestation.SnapGas(),
			)
			if err != nil {
				return nil, fmt.Errorf("could not submit attestation: %w", err)
			}
			snapRootStr := common.BytesToHash(snapRoot[:]).String()
			types.LogTx("NOTARY", fmt.Sprintf("Submitted attestation with snapRoot: %s", snapRootStr), n.destinationDomain.Config().DomainID, tx)
			span.AddEvent("Submitted transaction", trace.WithAttributes(
				attribute.String("tx", tx.Hash().Hex()),
			))

			return
		})
		if err != nil {
			span.AddEvent("Error submitting attestation", trace.WithAttributes(
				attribute.String("err", err.Error()),
			))
			return
		}
	}
}

// Start starts the notary.
//
//nolint:cyclop
func (n *Notary) Start(parentCtx context.Context) error {
	g, ctx := errgroup.WithContext(parentCtx)

	logger.Info("Starting the notary")

	// Ensure that this notary is active on the Summit
	err := n.ensureNotaryActive(ctx)
	if err != nil {
		return err
	}

	// Load latest states from Summit
	n.loadSummitMyLatestStates(ctx)

	g.Go(func() error {
		err := n.txSubmitter.Start(ctx)
		if err != nil {
			err = fmt.Errorf("could not start tx submitter: %w", err)
		}
		return err
	})

	g.Go(func() error {
		for {
			select {
			// parent loop terminated
			case <-ctx.Done():
				logger.Info("Notary exiting without error")
				return nil
			case <-time.After(n.refreshInterval):
				n.loadSummitGuardLatestStates(ctx)
				n.submitLatestSnapshot(ctx)
				shouldRegisterNotary, shouldSendToDestination := n.shouldNotaryRegisteredOnDestination(ctx)
				didRegisterAgent := true
				if shouldRegisterNotary {
					didRegisterAgent = n.registerNotaryOnDestination(ctx)
				}
				if shouldSendToDestination && didRegisterAgent {
					n.submitAttestation(ctx)
				}
			}
		}
	})

	err = g.Wait()
	if err != nil {
		logger.Errorf("Notary exiting with error: %v", err)
		return fmt.Errorf("could not start the notary: %w", err)
	}

	logger.Info("Notary exiting without error")
	return nil
}
