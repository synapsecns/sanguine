package guard

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/agents/guard/db"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/core/metrics"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/client"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// Guard scans origins for latest state and submits snapshots to the Summit.
type Guard struct {
	bondedSigner       signer.Signer
	unbondedSigner     signer.Signer
	domains            map[uint32]domains.DomainClient
	summitDomainID     uint32
	refreshInterval    time.Duration
	summitLatestStates map[uint32]types.State
	// TODO: change to metrics type
	originLatestStates map[uint32]types.State
	handler            metrics.Handler
	grpcClient         pbscribe.ScribeServiceClient
	grpcConn           *grpc.ClientConn
	logChans           map[uint32]chan *ethTypes.Log
	inboxParser        inbox.Parser
	boundInbox         *inbox.Inbox
	boundOrigins       map[uint32]*origin.Origin
	txSubmitter        submitter.TransactionSubmitter
}

const (
	logChanSize          = 1000
	scribeConnectTimeout = 30 * time.Second
)

func makeScribeClient(parentCtx context.Context, handler metrics.Handler, url string) (*grpc.ClientConn, pbscribe.ScribeServiceClient, error) {
	ctx, cancel := context.WithTimeout(parentCtx, scribeConnectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor(otelgrpc.WithTracerProvider(handler.GetTracerProvider()))),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor(otelgrpc.WithTracerProvider(handler.GetTracerProvider()))),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("could not dial grpc: %w", err)
	}

	scribeClient := pbscribe.NewScribeServiceClient(conn)

	// Ensure that gRPC is up and running.
	healthCheck, err := scribeClient.Check(ctx, &pbscribe.HealthCheckRequest{}, grpc.WaitForReady(true))
	if err != nil {
		return nil, nil, fmt.Errorf("could not check: %w", err)
	}
	if healthCheck.Status != pbscribe.HealthCheckResponse_SERVING {
		return nil, nil, fmt.Errorf("not serving: %s", healthCheck.Status)
	}

	return conn, scribeClient, nil
}

// NewGuard creates a new guard.
//
//nolint:cyclop
func NewGuard(ctx context.Context, cfg config.AgentConfig, omniRPCClient omnirpcClient.RPCClient, scribeClient client.ScribeClient, txDB db.GuardDB, handler metrics.Handler) (_ Guard, err error) {
	guard := Guard{
		refreshInterval: time.Second * time.Duration(cfg.RefreshIntervalSeconds),
	}
	guard.domains = make(map[uint32]domains.DomainClient)

	guard.grpcConn, guard.grpcClient, err = makeScribeClient(ctx, handler, fmt.Sprintf("%s:%d", scribeClient.URL, scribeClient.Port))
	if err != nil {
		return Guard{}, fmt.Errorf("could not create scribe client: %w", err)
	}

	guard.bondedSigner, err = signerConfig.SignerFromConfig(ctx, cfg.BondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error with bondedSigner, could not create guard: %w", err)
	}

	guard.unbondedSigner, err = signerConfig.SignerFromConfig(ctx, cfg.UnbondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error with unbondedSigner, could not create guard: %w", err)
	}

	for domainName, domain := range cfg.Domains {
		var domainClient domains.DomainClient

		omnirpcClient, err := omniRPCClient.GetConfirmationsClient(ctx, int(domain.DomainID), 1)
		if err != nil {
			return Guard{}, fmt.Errorf("error with omniRPCClient, could not create guard: %w", err)
		}

		chainRPCURL := omniRPCClient.GetEndpoint(int(domain.DomainID), 1)

		domainClient, err = evm.NewEVM(ctx, domainName, domain, chainRPCURL)
		if err != nil {
			return Guard{}, fmt.Errorf("failing to create evm for domain, could not create guard for: %w", err)
		}
		guard.domains[domain.DomainID] = domainClient

		guard.logChans[domain.DomainID] = make(chan *ethTypes.Log, logChanSize)
		guard.boundOrigins[domain.DomainID], err = origin.NewOrigin(
			common.HexToAddress(domain.OriginAddress),
			omnirpcClient,
		)

		// Initializations that only need to happen on the Summit domain.
		if domain.DomainID == cfg.SummitDomainID {
			guard.summitDomainID = domain.DomainID
			// Create a new inbox parser for the summit domain.
			guard.inboxParser, err = inbox.NewParser(common.HexToAddress(domain.InboxAddress))
			if err != nil {
				return Guard{}, fmt.Errorf("could not create inbox parser: %w", err)
			}

			guard.boundInbox, err = inbox.NewInbox(
				common.HexToAddress(domain.InboxAddress),
				omnirpcClient,
			)
			if err != nil {
				return Guard{}, fmt.Errorf("could not create bonding manager: %w", err)
			}
		}
	}

	guard.summitLatestStates = make(map[uint32]types.State, len(guard.domains))
	guard.originLatestStates = make(map[uint32]types.State, len(guard.domains))

	guard.handler = handler

	guard.txSubmitter = submitter.NewTransactionSubmitter(handler, guard.unbondedSigner, omniRPCClient, txDB.SubmitterDB(), &cfg.SubmitterConfig)

	return guard, nil
}

// streamLogs uses the grpcConnection to Scribe, with a chainID and address to get all logs from that address.
func (g Guard) streamLogs(ctx context.Context, chainID uint32, address string) error {
	// TODO: Get last block number to define starting point for streamLogs.
	fromBlock := strconv.FormatUint(0, 16)

	toBlock := "latest"

	stream, err := g.grpcClient.StreamLogs(ctx, &pbscribe.StreamLogsRequest{
		Filter: &pbscribe.LogFilter{
			ContractAddress: &pbscribe.NullableString{Kind: &pbscribe.NullableString_Data{Data: address}},
			ChainId:         chainID,
		},
		FromBlock: fromBlock,
		ToBlock:   toBlock,
	})
	if err != nil {
		return fmt.Errorf("could not stream logs: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			err := stream.CloseSend()
			if err != nil {
				return fmt.Errorf("could not close stream: %w", err)
			}

			err = g.grpcConn.Close()
			if err != nil {
				return fmt.Errorf("could not close connection: %w", err)
			}

			return fmt.Errorf("context done: %w", ctx.Err())
		default:
			response, err := stream.Recv()
			if err != nil {
				return fmt.Errorf("could not receive: %w", err)
			}

			log := response.Log.ToLog()
			if log == nil {
				return fmt.Errorf("could not convert log")
			}

			g.logChans[chainID] <- log
		}
	}
}

// receiveLogs continuously receives logs from the log channel and processes them.
func (g Guard) receiveLogs(ctx context.Context, chainID uint32) error {
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context canceled: %w", ctx.Err())
		case log := <-g.logChans[chainID]:
			if log == nil {
				return fmt.Errorf("log is nil")
			}

			err := g.processSnapshot(ctx, *log)
			if err != nil {
				return fmt.Errorf("could not process log: %w", err)
			}
		}
	}
}

func (g Guard) processSnapshot(ctx context.Context, log ethTypes.Log) error {
	snapshot, agentSig, err := g.logToSnapshot(log)
	if err != nil {
		return fmt.Errorf("could not convert log to snapshot: %w", err)
	}

	if snapshot == nil {
		return nil
	}

	snapshotPayload, err := types.EncodeSnapshot(snapshot)
	if err != nil {
		return fmt.Errorf("could not encode snapshot: %w", err)
	}

	for stateIndex, state := range snapshot.States() {
		// Check the validity of each state by calling `isValidState` on each state's origin domain.
		statePayload, err := types.EncodeState(state)
		if err != nil {
			return fmt.Errorf("could not encode state: %w", err)
		}

		// TODO: Have a way to retry failed RPC calls for this check.
		isValid, err := g.domains[state.Origin()].Origin().IsValidState(
			ctx,
			statePayload,
		)
		if err != nil {
			return fmt.Errorf("could not check validity of state: %w", err)
		}

		if !isValid {
			// Call submitStateReportWithSnapshot on (each?) destination domain.
			err = g.submitStateReports(ctx, int64(stateIndex), snapshotPayload, agentSig)
		}
	}

	return nil
}

func (g Guard) submitStateReports(ctx context.Context, stateIndex int64, snapshotPaylod, snapshotSig []byte) error {
	// Call on the Summit's `BondingManager` contract.
	signature, err := g.bondedSigner.SignMessage(ctx, snapshotPaylod, false)
	if err != nil {
		return fmt.Errorf("could not sign message: %w", err)
	}

	// TODO: What is signature here?
	_, err = g.domains[g.summitDomainID].Inbox().SubmitStateReportWithSnapshot(ctx, g.bondedSigner, stateIndex, signature, snapshotPaylod, snapshotSig)
	if err != nil {
		return fmt.Errorf("could not submit state report with snapshot: %w", err)
	}

	return nil
}

// logToSnapshot converts the log to a snapshot.
func (g Guard) logToSnapshot(log ethTypes.Log) (types.Snapshot, []byte, error) {
	snapshot, domain, agentSig, ok := g.inboxParser.ParseSnapshotAccepted(log)
	if !ok {
		return nil, nil, fmt.Errorf("could not parse snapshot")
	}

	// Domain == 0 check here only qualifies Notary submitted snapshots.
	if snapshot == nil || domain == 0 || agentSig == nil {
		//nolint:nilnil
		return nil, nil, nil
	}

	return snapshot, nil, nil
}

//nolint:cyclop
func (g Guard) loadSummitLatestStates(parentCtx context.Context) {
	for _, domain := range g.domains {
		ctx, span := g.handler.Tracer().Start(parentCtx, "loadSummitLatestStates", trace.WithAttributes(
			attribute.Int("domain", int(domain.Config().DomainID)),
		))

		originID := domain.Config().DomainID
		latestState, err := g.domains[g.summitDomainID].Summit().GetLatestAgentState(ctx, originID, g.bondedSigner)
		if err != nil {
			latestState = nil
			logger.Errorf("Failed calling GetLatestAgentState for originID %d on the Summit contract: err = %v", originID, err)
			span.AddEvent("Failed calling GetLatestAgentState for originID on the Summit contract", trace.WithAttributes(
				attribute.Int("originID", int(originID)),
				attribute.String("err", err.Error()),
			))
		}
		if latestState != nil && latestState.Nonce() > uint32(0) {
			g.summitLatestStates[originID] = latestState
		}

		span.End()
	}
}

//nolint:cyclop
func (g Guard) loadOriginLatestStates(parentCtx context.Context) {
	for _, domain := range g.domains {
		ctx, span := g.handler.Tracer().Start(parentCtx, "loadOriginLatestStates", trace.WithAttributes(
			attribute.Int("domain", int(domain.Config().DomainID)),
		))

		originID := domain.Config().DomainID
		latestState, err := domain.Origin().SuggestLatestState(ctx)
		if err != nil {
			latestState = nil
			logger.Errorf("Failed calling SuggestLatestState for originID %d on the Origin contract: %v", originID, err)
			span.AddEvent("Failed calling SuggestLatestState for originID on the Origin contract", trace.WithAttributes(
				attribute.Int("originID", int(originID)),
				attribute.String("err", err.Error()),
			))
		} else if latestState == nil || latestState.Nonce() == uint32(0) {
			logger.Errorf("No latest state found for origin id %d", originID)
			span.AddEvent("No latest state found for origin id", trace.WithAttributes(
				attribute.Int("originID", int(originID)),
			))
		}
		if latestState != nil {
			// TODO: if overwriting, end span and start a new one
			g.originLatestStates[originID] = latestState
		}

		span.End()
	}
}

//nolint:cyclop
func (g Guard) getLatestSnapshot() (types.Snapshot, map[uint32]types.State) {
	statesToSubmit := make(map[uint32]types.State, len(g.domains))
	for _, domain := range g.domains {
		originID := domain.Config().DomainID
		summitLatest, ok := g.summitLatestStates[originID]
		if !ok || summitLatest == nil || summitLatest.Nonce() == 0 {
			summitLatest = nil
		}
		originLatest, ok := g.originLatestStates[originID]
		if !ok || originLatest == nil || originLatest.Nonce() == 0 {
			continue
		}
		if summitLatest != nil && summitLatest.Nonce() >= originLatest.Nonce() {
			// Here this guard already submitted this state
			continue
		}
		// TODO: add event for submitting that state
		statesToSubmit[originID] = originLatest
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
func (g Guard) submitLatestSnapshot(parentCtx context.Context) {
	summitDomain := g.domains[g.summitDomainID]

	ctx, span := g.handler.Tracer().Start(parentCtx, "submitLatestSnapshot", trace.WithAttributes(
		attribute.Int("domain", int(g.summitDomainID)),
	))

	defer func() {
		span.End()
	}()

	snapshot, statesToSubmit := g.getLatestSnapshot()
	if snapshot == nil {
		return
	}

	snapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(ctx, g.bondedSigner)

	//nolint:nestif
	if err != nil {
		logger.Errorf("Error signing snapshot: %v", err)
		span.AddEvent("Error signing snapshot", trace.WithAttributes(
			attribute.String("err", err.Error()),
		))
	} else {
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(g.summitDomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			rawSig, err := types.EncodeSignature(snapshotSignature)
			if err != nil {
				return nil, fmt.Errorf("failed to encode signature: %w", err)
			}

			tx, err = summitDomain.Inbox().GetContractRef().SubmitSnapshot(transactor, encodedSnapshot, rawSig)
			if err != nil {
				if strings.Contains(err.Error(), "nonce too low") {
					summitDomain.Inbox().GetNonceManager().ClearNonce(g.unbondedSigner.Address())
				}
				return nil, fmt.Errorf("failed to submit snapshot: %w", err)
			}

			return
		})
		if err != nil {
			logger.Errorf("Failed to submit snapshot to inbox: %v", err)
			span.AddEvent("Failed to submit snapshot to inbox", trace.WithAttributes(
				attribute.String("err", err.Error()),
			))
		} else {
			for originID, state := range statesToSubmit {
				g.summitLatestStates[originID] = state
			}
		}
	}
}

// Start starts the guard.
//
//nolint:cyclop
func (g Guard) Start(parentCtx context.Context) error {
	// First initialize a map to track what was the last state signed by this guard
	g.loadSummitLatestStates(parentCtx)

	group, ctx := errgroup.WithContext(parentCtx)

	group.Go(func() error {
		err := g.txSubmitter.Start(ctx)
		if err != nil {
			err = fmt.Errorf("could not start tx submitter: %w", err)
		}
		return err
	})

	group.Go(func() error {
		for {
			select {
			// parent loop terminated
			case <-ctx.Done():
				logger.Info("Guard exiting without error")
				return nil
			case <-time.After(g.refreshInterval):
				g.loadOriginLatestStates(ctx)
				g.submitLatestSnapshot(ctx)
			}
		}
	})

	if err := group.Wait(); err != nil {
		return fmt.Errorf("guard error: %w", err)
	}

	return nil
}
