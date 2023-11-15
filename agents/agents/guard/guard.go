package guard

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/synapsecns/sanguine/agents/contracts/bondingmanager"
	"github.com/synapsecns/sanguine/agents/contracts/lightmanager"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/agents/guard/db"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
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
	originLatestStates   map[uint32]types.State
	handler              metrics.Handler
	grpcClient           pbscribe.ScribeServiceClient
	grpcConn             *grpc.ClientConn
	logChans             map[uint32]chan *ethTypes.Log
	inboxParser          inbox.Parser
	lightInboxParser     lightinbox.Parser
	bondingManagerParser bondingmanager.Parser
	lightManagerParser   lightmanager.Parser
	boundOrigins         map[uint32]*origin.Origin
	txSubmitter          submitter.TransactionSubmitter
	retryConfig          []retry.WithBackoffConfigurator
	guardDB              db.GuardDB
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
func NewGuard(ctx context.Context, cfg config.AgentConfig, omniRPCClient omnirpcClient.RPCClient, scribeClient client.ScribeClient, guardDB db.GuardDB, handler metrics.Handler) (guard *Guard, err error) {
	guard = &Guard{
		refreshInterval: time.Second * time.Duration(cfg.RefreshIntervalSeconds),
		domains:         make(map[uint32]domains.DomainClient),
		logChans:        make(map[uint32]chan *ethTypes.Log),
		boundOrigins:    make(map[uint32]*origin.Origin),
	}

	guard.grpcConn, guard.grpcClient, err = makeScribeClient(ctx, handler, fmt.Sprintf("%s:%d", scribeClient.URL, scribeClient.Port))
	if err != nil {
		return nil, fmt.Errorf("could not create scribe client: %w", err)
	}

	guard.bondedSigner, err = signerConfig.SignerFromConfig(ctx, cfg.BondedSigner)
	if err != nil {
		return nil, fmt.Errorf("error with bondedSigner, could not create guard: %w", err)
	}

	guard.unbondedSigner, err = signerConfig.SignerFromConfig(ctx, cfg.UnbondedSigner)
	if err != nil {
		return nil, fmt.Errorf("error with unbondedSigner, could not create guard: %w", err)
	}

	// Set up evm utilities for each domain
	for domainName, domain := range cfg.Domains {
		omnirpcClient, err := omniRPCClient.GetConfirmationsClient(ctx, int(domain.DomainID), 1)
		if err != nil {
			return nil, fmt.Errorf("error with omniRPCClient, could not create guard: %w", err)
		}

		chainRPCURL := omniRPCClient.GetDefaultEndpoint(int(domain.DomainID))
		domainClient, err := evm.NewEVM(ctx, domainName, domain, chainRPCURL)
		if err != nil {
			return nil, fmt.Errorf("failing to create evm for domain, could not create guard for: %w", err)
		}
		guard.domains[domain.DomainID] = domainClient

		guard.logChans[domain.DomainID] = make(chan *ethTypes.Log, logChanSize)
		guard.boundOrigins[domain.DomainID], err = origin.NewOrigin(
			common.HexToAddress(domain.OriginAddress),
			omnirpcClient,
		)
		if err != nil {
			return nil, fmt.Errorf("could not create origin: %w", err)
		}

		// Initialize contract parsers for the summit domain.
		if domain.DomainID == cfg.SummitDomainID {
			guard.summitDomainID = domain.DomainID

			guard.inboxParser, err = inbox.NewParser(common.HexToAddress(domain.InboxAddress))
			if err != nil {
				return nil, fmt.Errorf("could not create inbox parser: %w", err)
			}

			guard.lightInboxParser, err = lightinbox.NewParser(common.HexToAddress(domain.LightInboxAddress))
			if err != nil {
				return nil, fmt.Errorf("could not create inbox parser: %w", err)
			}

			guard.bondingManagerParser, err = bondingmanager.NewParser(common.HexToAddress(domain.BondingManagerAddress))
			if err != nil {
				return nil, fmt.Errorf("could not create bonding manager parser: %w", err)
			}

			guard.lightManagerParser, err = lightmanager.NewParser(common.HexToAddress(domain.LightManagerAddress))
			if err != nil {
				return nil, fmt.Errorf("could not create light manager parser: %w", err)
			}
		}
	}

	_, ok := guard.domains[guard.summitDomainID]
	if !ok {
		return nil, fmt.Errorf("summit domain not set: %d", guard.summitDomainID)
	}

	guard.summitLatestStates = make(map[uint32]types.State, len(guard.domains))
	guard.originLatestStates = make(map[uint32]types.State, len(guard.domains))
	guard.handler = handler
	guard.txSubmitter = submitter.NewTransactionSubmitter(handler, guard.unbondedSigner, omniRPCClient, guardDB.SubmitterDB(), &cfg.SubmitterConfig)

	if cfg.MaxRetrySeconds == 0 {
		cfg.MaxRetrySeconds = 60
	}

	guard.retryConfig = []retry.WithBackoffConfigurator{
		retry.WithMaxAttemptTime(time.Second * time.Duration(cfg.MaxRetrySeconds)),
	}
	guard.guardDB = guardDB

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
		response, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("could not receive: %w", err)
		}

		log := response.Log.ToLog()
		if log == nil {
			return fmt.Errorf("could not convert log")
		}

		select {
		case <-ctx.Done():
			fmt.Printf("context done on chain %d addr %s\n", chainID, address)
			err := stream.CloseSend()
			if err != nil {
				return fmt.Errorf("could not close stream: %w", err)
			}

			err = g.grpcConn.Close()
			if err != nil {
				return fmt.Errorf("could not close connection: %w", err)
			}

			fmt.Printf("exiting on chain %d addr %s\n", chainID, address)
			return fmt.Errorf("context done: %w", ctx.Err())
		case g.logChans[chainID] <- log:
			logger.Info("Received log with topic: %s", log.Topics[0].String())
		}
	}
}

// receiveLogs continuously receives logs from the log channel and processes them.
func (g Guard) receiveLogs(ctx context.Context, chainID uint32) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("exiting receiveLogs on chainID %d\n", chainID)
			return fmt.Errorf("context canceled: %w", ctx.Err())
		case log := <-g.logChans[chainID]:
			if log == nil {
				return fmt.Errorf("log is nil")
			}

			err := g.handleLog(ctx, *log, chainID)
			if err != nil {
				//TODO: how to handle error here?
				logger.Errorf("could not process log: %v", err)
			}
		}
	}
}

func (g Guard) handleLog(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	switch {
	case isSnapshotAcceptedEvent(g.inboxParser, log):
		return g.handleSnapshotAccepted(ctx, log)
	case isAttestationAcceptedEvent(g.lightInboxParser, log):
		return g.handleAttestationAccepted(ctx, log)
	case isReceiptAcceptedEvent(g.inboxParser, log):
		return g.handleReceiptAccepted(ctx, log)
	case isStatusUpdatedEvent(g.bondingManagerParser, log):
		return g.handleStatusUpdated(ctx, log, chainID)
	case isRootUpdatedEvent(g.bondingManagerParser, log):
		return g.handleRootUpdated(ctx, log, chainID)
	}
	return nil
}

//nolint:cyclop
func (g Guard) loadSummitLatestStates(parentCtx context.Context) {
	for _, domain := range g.domains {
		ctx, span := g.handler.Tracer().Start(parentCtx, "loadSummitLatestStates", trace.WithAttributes(
			attribute.Int(metrics.ChainID, int(domain.Config().DomainID)),
		))

		originID := domain.Config().DomainID

		var latestState types.State
		var err error
		contractCall := func(ctx context.Context) error {
			latestState, err = g.domains[g.summitDomainID].Summit().GetLatestAgentState(ctx, originID, g.bondedSigner)
			if err != nil {
				return fmt.Errorf("failed calling GetLatestAgentState for originID %d on the Summit contract: err = %w", originID, err)
			}

			return nil
		}
		err = retry.WithBackoff(ctx, contractCall, g.retryConfig...)
		if err == nil && latestState.Nonce() > uint32(0) {
			g.summitLatestStates[originID] = latestState
		}

		span.End()
	}
}

//nolint:cyclop
func (g Guard) loadOriginLatestStates(parentCtx context.Context) {
	fmt.Println("loadOriginLatestStates")
	for _, d := range g.domains {
		domain := d
		ctx, span := g.handler.Tracer().Start(parentCtx, "loadOriginLatestStates", trace.WithAttributes(
			attribute.Int(metrics.ChainID, int(domain.Config().DomainID)),
		))

		originID := domain.Config().DomainID

		var latestState types.State
		contractCall := func(ctx context.Context) (err error) {
			latestState, err = domain.Origin().SuggestLatestState(ctx)
			if err != nil {
				return fmt.Errorf("failed calling GetLatestAgentState for originID %d on the Summit contract: err = %w", originID, err)
			}

			return nil
		}
		err := retry.WithBackoff(ctx, contractCall, g.retryConfig...)
		if err != nil {
			latestState = nil
			logger.Errorf("Failed calling SuggestLatestState for originID %d on the Origin contract: %v", originID, err)
			span.AddEvent("failed calling SuggestLatestState for originID on the Origin contract", trace.WithAttributes(
				attribute.Int("originID", int(originID)),
				attribute.String(metrics.Error, err.Error()),
			))
		} else if latestState == nil || latestState.Nonce() == uint32(0) {
			logger.Errorf("No latest state found for origin id %d", originID)
			span.AddEvent("no latest state found for origin id", trace.WithAttributes(
				attribute.Int("originID", int(originID)),
			))
		}
		if latestState != nil {
			// TODO: if overwriting, end span and start a new one
			g.originLatestStates[originID] = latestState
			span.AddEvent("set latest state", trace.WithAttributes(
				attribute.Int(metrics.Origin, int(originID)),
				attribute.Int("nonce", int(latestState.Nonce())),
			))
		}

		span.End()
	}
}

//nolint:cyclop
func (g Guard) getLatestSnapshot(parentCtx context.Context) (types.Snapshot, map[uint32]types.State) {
	fmt.Println("getLatestSnapshot")
	_, span := g.handler.Tracer().Start(parentCtx, "getLatestSnapshot", trace.WithAttributes(
		stateMapToAttribute("summitLatestStates", g.summitLatestStates),
		stateMapToAttribute("originLatestStates", g.originLatestStates),
	))
	defer func() {
		span.End()
	}()

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
		span.AddEvent("got origin state to submit", trace.WithAttributes(
			attribute.Int(metrics.Origin, int(originID)),
			attribute.Int("nonce", int(originLatest.Nonce())),
		))
	}
	snapshotStates := make([]types.State, 0, len(statesToSubmit))
	for _, state := range statesToSubmit {
		if state.Nonce() == 0 {
			continue
		}
		snapshotStates = append(snapshotStates, state)
	}
	span.AddEvent("got latest states for snapshot", trace.WithAttributes(stateSliceToAttribute("snapshotStates", snapshotStates)))
	if len(snapshotStates) > 0 {
		snapshot := types.NewSnapshot(snapshotStates)
		snapRoot, _, _ := snapshot.SnapshotRootAndProofs()
		span.SetAttributes(attribute.String(metrics.SnapRoot, common.BytesToHash(snapRoot[:]).String()))
		return snapshot, statesToSubmit
	}
	//nolint:nilnil
	return nil, nil
}

//nolint:cyclop
func (g Guard) submitLatestSnapshot(parentCtx context.Context) {
	fmt.Println("submitLatestSnapshot")
	summitDomain := g.domains[g.summitDomainID]

	ctx, span := g.handler.Tracer().Start(parentCtx, "submitLatestSnapshot", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(g.summitDomainID)),
	))

	defer func() {
		span.End()
	}()

	snapshot, statesToSubmit := g.getLatestSnapshot(ctx)
	if snapshot == nil {
		return
	}

	snapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(ctx, g.bondedSigner)

	//nolint:nestif
	if err != nil {
		logger.Errorf("Error signing snapshot: %v", err)
		span.AddEvent("Error signing snapshot", trace.WithAttributes(
			attribute.String(metrics.Error, err.Error()),
		))
	} else {
		snapshotRoot, _, _ := snapshot.SnapshotRootAndProofs()
		fmt.Printf("submitting guard snapshot with root: %v\n", common.BytesToHash(snapshotRoot[:]).String())
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(g.summitDomainID)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
			tx, err = summitDomain.Inbox().SubmitSnapshot(transactor, encodedSnapshot, snapshotSignature)
			if err != nil {
				return nil, fmt.Errorf("failed to submit snapshot: %w", err)
			}
			types.LogTx("GUARD", "Submitted snapshot", g.summitDomainID, tx)

			return
		})
		if err != nil {
			logger.Errorf("Failed to submit snapshot to inbox: %v", err)
			span.AddEvent("Failed to submit snapshot to inbox", trace.WithAttributes(
				attribute.String(metrics.Error, err.Error()),
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
		return g.streamLogs(ctx, g.summitDomainID, g.domains[g.summitDomainID].Config().InboxAddress)
	})

	group.Go(func() error {
		return g.streamLogs(ctx, g.summitDomainID, g.domains[g.summitDomainID].Config().BondingManagerAddress)
	})

	group.Go(func() error {
		return g.receiveLogs(ctx, g.summitDomainID)
	})

	for _, domain := range g.domains {
		domainCfg := domain.Config()
		if domainCfg.DomainID == g.summitDomainID {
			continue
		}

		group.Go(func() error {
			return g.streamLogs(ctx, domainCfg.DomainID, domainCfg.LightInboxAddress)
		})

		group.Go(func() error {
			return g.streamLogs(ctx, domainCfg.DomainID, domainCfg.LightManagerAddress)
		})

		group.Go(func() error {
			return g.receiveLogs(ctx, domainCfg.DomainID)
		})
	}

	group.Go(func() error {
		fmt.Println("starting guard parent loop")
		for {
			select {
			// parent loop terminated
			case <-ctx.Done():
				logger.Info("Guard exiting without error")
				return nil
			case <-time.After(g.refreshInterval):
				g.loadOriginLatestStates(ctx)
				g.submitLatestSnapshot(ctx)
				err := g.updateAgentStatuses(ctx)
				if err != nil {
					fmt.Println("exiting refresh loop")
					return err
				}
			}
		}
	})

	err := group.Wait()
	if err != nil {
		fmt.Printf("guard exiting with error: %v\n", err)
		return fmt.Errorf("guard error: %w", err)
	}

	fmt.Println("guard exiting")
	return nil
}
