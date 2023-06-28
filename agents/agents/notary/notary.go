package notary

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"

	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"golang.org/x/sync/errgroup"
)

// Notary checks the Summit for that latest states signed by guards, validates those states on origin,
// then signs and submits the snapshot to Summit.
type Notary struct {
	bondedSigner                       signer.Signer
	unbondedSigner                     signer.Signer
	domains                            []domains.DomainClient
	summitDomain                       domains.DomainClient
	destinationDomain                  domains.DomainClient
	refreshInterval                    time.Duration
	summitMyLatestStates               map[uint32]types.State
	summitGuardLatestStates            map[uint32]types.State
	myLatestNotaryAttestation          types.NotaryAttestation
	didSubmitMyLatestNotaryAttestation bool
	summitParser                       summit.Parser
	lastSummitBlock                    uint64
	handler                            metrics.Handler
}

// NewNotary creates a new notary.
//
//nolint:cyclop
func NewNotary(ctx context.Context, cfg config.AgentConfig, handler metrics.Handler) (_ Notary, err error) {
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

	for domainName, domain := range cfg.Domains {
		var domainClient domains.DomainClient

		chainRPCURL := fmt.Sprintf("%s/confirmations/1/rpc/%d", cfg.BaseOmnirpcURL, domain.DomainID)
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

	return notary, nil
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
		}
		if myLatestState != nil && myLatestState.Nonce() > uint32(0) {
			n.summitMyLatestStates[originID] = myLatestState
		}

		span.End()
	}
}

//nolint:cyclop
func (n *Notary) loadSummitGuardLatestStates(parentCtx context.Context) {
	for _, domain := range n.domains {
		ctx, span := n.handler.Tracer().Start(parentCtx, "loadSummitGuardLatestStates", trace.WithAttributes(
			attribute.Int(metrics.ChainID, int(domain.Config().DomainID)),
		))

		originID := domain.Config().DomainID

		guardLatestState, err := n.summitDomain.Summit().GetLatestState(ctx, originID)
		if err != nil {
			guardLatestState = nil
			span.AddEvent("GetLatestState failed", trace.WithAttributes(
				attribute.String("err", err.Error()),
			))
		}
		if guardLatestState != nil && guardLatestState.Nonce() > uint32(0) {
			n.summitGuardLatestStates[originID] = guardLatestState
		}

		span.End()
	}
}

//nolint:cyclop
func (n *Notary) loadNotaryLatestAttestation(parentCtx context.Context) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "loadNotaryLatestAttestation", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(n.destinationDomain.Config().DomainID)),
	))
	defer span.End()

	latestNotaryAttestation, err := n.summitDomain.Summit().GetLatestNotaryAttestation(ctx, n.bondedSigner)
	if err != nil {
		span.AddEvent("GetLatestNotaryAttestation failed", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
	}
	if latestNotaryAttestation != nil {
		if n.myLatestNotaryAttestation == nil ||
			latestNotaryAttestation.Attestation().SnapshotRoot() != n.myLatestNotaryAttestation.Attestation().SnapshotRoot() {
			n.myLatestNotaryAttestation = latestNotaryAttestation
			n.didSubmitMyLatestNotaryAttestation = false
		}
	}
}

func (n *Notary) shouldNotaryRegisteredOnDestination(parentCtx context.Context) (bool, bool) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "shouldNotaryRegisteredOnDestination", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(n.destinationDomain.Config().DomainID)),
	))
	defer span.End()

	bondingManagerAgentRoot, err := n.summitDomain.BondingManager().GetAgentRoot(ctx)
	if err != nil {
		span.AddEvent("GetAgentRoot failed on bonding manager", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
		return false, false
	}

	destinationLightManagerAgentRoot, err := n.destinationDomain.LightManager().GetAgentRoot(ctx)
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

	agentStatus, err := n.destinationDomain.LightManager().GetAgentStatus(ctx, n.bondedSigner)
	if err != nil {
		span.AddEvent("GetAgentStatus failed", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
		return false, false
	}
	if types.AgentFlagType(agentStatus.Flag()) == types.AgentFlagUnknown {
		// Here we want to add the Notary and proceed with sending to destination
		return true, true
	} else if types.AgentFlagType(agentStatus.Flag()) == types.AgentFlagActive {
		// Here we already added the Notary and can proceed with sending to destination
		return false, true
	}
	return false, false
}

func (n *Notary) checkDidSubmitNotaryLatestAttestation(parentCtx context.Context) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "checkDidSubmitNotaryLatestAttestation", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(n.destinationDomain.Config().DomainID)),
	))
	defer span.End()

	if n.myLatestNotaryAttestation == nil {
		n.didSubmitMyLatestNotaryAttestation = false
		return
	}

	if n.didSubmitMyLatestNotaryAttestation {
		return
	}

	attNonce, err := n.destinationDomain.Destination().GetAttestationNonce(ctx, n.myLatestNotaryAttestation.Attestation().SnapshotRoot())
	if err != nil {
		span.AddEvent("checkDidSubmitNotaryLatestAttestation failed", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
	}
	if attNonce > 0 {
		n.didSubmitMyLatestNotaryAttestation = true
	}
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

	stateOnOrigin, err := domain.Origin().SuggestState(ctx, state.Nonce())
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
	statesToSubmit := make(map[uint32]types.State, len(n.domains))
	for _, domain := range n.domains {
		ctx, span := n.handler.Tracer().Start(parentCtx, "getLatestSnapshot", trace.WithAttributes(
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
			span.AddEvent("State not valid on origin", trace.WithAttributes(
				attribute.Int("nonce", int(summitGuardLatest.Nonce())),
			))
			continue
		}
		statesToSubmit[originID] = summitGuardLatest

		span.End()
	}
	snapshotStates := make([]types.State, 0, len(statesToSubmit))
	for _, state := range statesToSubmit {
		if state.Nonce() == 0 {
			continue
		}
		if state.Origin() == n.destinationDomain.Config().DomainID {
			continue
		}
		snapshotStates = append(snapshotStates, state)
	}
	if len(snapshotStates) > 0 {
		return types.NewSnapshot(snapshotStates), statesToSubmit
	}
	//nolint:nilnil
	return nil, nil
}

//nolint:cyclop
func (n *Notary) submitLatestSnapshot(parentCtx context.Context) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "submitLatestSnapshot")
	defer span.End()

	snapshot, statesToSubmit := n.getLatestSnapshot(ctx)
	if snapshot == nil {
		return
	}

	snapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(ctx, n.bondedSigner)
	if err != nil {
		span.AddEvent("Error signing snapshot", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
	} else {
		logger.Infof("Notary submitting snapshot to summit")
		err := n.summitDomain.Inbox().SubmitSnapshot(ctx, n.unbondedSigner, encodedSnapshot, snapshotSignature)
		if err != nil {
			span.AddEvent("Error submitting snapshot", trace.WithAttributes(
				attribute.String("err", err.Error()),
			))
		} else {
			for originID, state := range statesToSubmit {
				n.summitMyLatestStates[originID] = state
			}
		}
	}
}

// codebeat:disable[CYCLO,DEPTH,LOC]
//
//nolint:cyclop
func (n *Notary) registerNotaryOnDestination(parentCtx context.Context) bool {
	ctx, span := n.handler.Tracer().Start(parentCtx, "registerNotaryOnDestination")
	defer span.End()

	agentProof, err := n.summitDomain.BondingManager().GetProof(ctx, n.bondedSigner)
	if err != nil {
		logger.Errorf("Error getting agent proof: %v", err)
		span.AddEvent("Error getting agent proof", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
		return false
	}
	agentStatus, err := n.summitDomain.BondingManager().GetAgentStatus(ctx, n.bondedSigner)
	if err != nil {
		span.AddEvent("GetAgentStatus on bonding manager failed", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
		return false
	}
	err = n.destinationDomain.LightManager().UpdateAgentStatus(
		ctx,
		n.unbondedSigner,
		n.bondedSigner,
		agentStatus,
		agentProof)
	if err != nil {
		span.AddEvent("Error updating agent status", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
		return false
	}
	return true
}

//nolint:cyclop,unused
func (n *Notary) submitMyLatestAttestation(parentCtx context.Context) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "submitMyLatestAttestation")
	defer span.End()

	if n.myLatestNotaryAttestation == nil {
		return
	}

	if n.didSubmitMyLatestNotaryAttestation {
		return
	}

	attestationSignature, _, _, err := n.myLatestNotaryAttestation.Attestation().SignAttestation(ctx, n.bondedSigner)
	if err != nil {
		logger.Errorf("Error signing attestation: %v", err)
		span.AddEvent("Error signing attestation", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
	} else {
		err = n.destinationDomain.LightInbox().SubmitAttestation(
			ctx,
			n.unbondedSigner,
			n.myLatestNotaryAttestation.AttPayload(),
			attestationSignature,
			n.myLatestNotaryAttestation.AgentRoot(),
			n.myLatestNotaryAttestation.SnapGas())
		if err != nil {
			span.AddEvent("Error submitting attestation", trace.WithAttributes(
				attribute.String("err", err.Error()),
			))
		}
	}
}

// Start starts the notary.
//
// codebeat:disable[CYCLO,DEPTH]
//
//nolint:cyclop
func (n *Notary) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	logger.Info("Starting the notary")

	// Setting latestBlock on summit chain
	latestBlockNUmber, err := n.summitDomain.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("could not get latest block number from Summit: %w", err)
	}
	// Try starting from previous day
	n.lastSummitBlock = uint64(latestBlockNUmber)
	if n.lastSummitBlock > 3000 {
		n.lastSummitBlock = uint64(latestBlockNUmber) - uint64(3000)
	} else {
		n.lastSummitBlock = uint64(0)
	}

	logger.Infof("Notary loadSummitMyLatestStates")
	n.loadSummitMyLatestStates(ctx)

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
				n.loadNotaryLatestAttestation(ctx)
				n.checkDidSubmitNotaryLatestAttestation(ctx)
				shouldRegisterNotary, shouldSendToDestination := n.shouldNotaryRegisteredOnDestination(ctx)
				didRegisterAgent := true
				if shouldRegisterNotary {
					didRegisterAgent = n.registerNotaryOnDestination(ctx)
				}
				if shouldSendToDestination && didRegisterAgent {
					n.submitMyLatestAttestation(ctx)
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
