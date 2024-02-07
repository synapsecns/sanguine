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
	retryConfig                        []retry.WithBackoffConfigurator
	txSubmitter                        submitter.TransactionSubmitter
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

	var latestNotaryAttestation types.NotaryAttestation
	contractCall := func(ctx context.Context) (err error) {
		latestNotaryAttestation, err = n.summitDomain.Summit().GetLatestNotaryAttestation(ctx, n.bondedSigner)
		if err != nil {
			return fmt.Errorf("could not get latest notary attestation: %w", err)
		}

		return nil
	}
	err := retry.WithBackoff(ctx, contractCall, n.retryConfig...)
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

	var attNonce uint32
	contractCall := func(ctx context.Context) (err error) {
		attNonce, err = n.destinationDomain.Destination().GetAttestationNonce(ctx, n.myLatestNotaryAttestation.Attestation().SnapshotRoot())
		if err != nil {
			return fmt.Errorf("could not get attestation nonce: %w", err)
		}

		return nil
	}
	err := retry.WithBackoff(ctx, contractCall, n.retryConfig...)
	if err != nil {
		span.AddEvent("GetAttestationNonce failed", trace.WithAttributes(
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

	//nolint:nestif
	if err != nil {
		span.AddEvent("Error signing snapshot", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
	} else {
		logger.Infof("Notary submitting snapshot to summit")
		_, err := n.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(n.summitDomain.Config().DomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = n.summitDomain.Inbox().SubmitSnapshot(transactor, encodedSnapshot, snapshotSignature)
			if err != nil {
				return nil, fmt.Errorf("could not submit snapshot: %w", err)
			}

			return
		})
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

//nolint:cyclop
func (n *Notary) registerNotaryOnDestination(parentCtx context.Context) bool {
	ctx, span := n.handler.Tracer().Start(parentCtx, "registerNotaryOnDestination")
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
func (n *Notary) submitMyLatestAttestation(parentCtx context.Context) {
	ctx, span := n.handler.Tracer().Start(parentCtx, "submitMyLatestAttestation")
	defer span.End()

	if n.myLatestNotaryAttestation == nil {
		return
	}

	if n.didSubmitMyLatestNotaryAttestation {
		return
	}

	attestationSignature, _, _, err := n.myLatestNotaryAttestation.Attestation().SignAttestation(ctx, n.bondedSigner, true)
	if err != nil {
		logger.Errorf("Error signing attestation: %v", err)
		span.AddEvent("Error signing attestation", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
	} else {
		_, err = n.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(n.destinationDomain.Config().DomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = n.destinationDomain.LightInbox().SubmitAttestation(
				transactor,
				n.myLatestNotaryAttestation.AttPayload(),
				attestationSignature,
				n.myLatestNotaryAttestation.AgentRoot(),
				n.myLatestNotaryAttestation.SnapGas(),
			)
			if err != nil {
				return nil, fmt.Errorf("could not submit attestation: %w", err)
			}

			return
		})
		if err != nil {
			span.AddEvent("Error submitting attestation", trace.WithAttributes(
				attribute.String("err", err.Error()),
			))
		}
	}
}

// Start starts the notary.
//
//nolint:cyclop
func (n *Notary) Start(parentCtx context.Context) error {
	g, ctx := errgroup.WithContext(parentCtx)

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
