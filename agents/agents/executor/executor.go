package executor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	evmClient "github.com/synapsecns/sanguine/ethergo/client"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"math/big"
	"strconv"
	"time"

	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/agents/contracts/summit"

	agentsConfig "github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/agents/executor/config"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	execTypes "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/merkle"
	"github.com/synapsecns/sanguine/services/scribe/client"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// chainExecutor is a struct that contains the necessary information for each chain level executor.
type chainExecutor struct {
	// chainID is the chain ID of the chain that this executor is responsible for.
	chainID uint32
	// lastLog is the last log that was processed.
	lastLog *logOrderInfo
	// closeConnection is a channel that is used to close the connection.
	closeConnection chan bool
	// stopListenChan is a channel that is used to stop listening to the log channel.
	stopListenChan chan bool
	// originParser is the origin parser.
	originParser origin.Parser
	// destinationParser is the destination parser.
	destinationParser destination.Parser
	// summitParser is the summit parser.
	summitParser *summit.Parser
	// logChan is the log channel.
	logChan chan *ethTypes.Log
	// merkleTree is a merkle tree for a specific origin chain.
	merkleTree *merkle.HistoricalTree
	// rpcClient is an RPC client.
	rpcClient Backend
	// boundDestination is a bound destination contract.
	boundDestination domains.DestinationContract
}

// Executor is the executor agent.
type Executor struct {
	// config is the executor agent config.
	config config.Config
	// executorDB is the executor agent database.
	executorDB db.ExecutorDB
	// scribeClient is the client to the Scribe gRPC server.
	scribeClient client.ScribeClient
	// grpcClient is the gRPC client.
	grpcClient pbscribe.ScribeServiceClient
	// grpcConn is the gRPC connection.
	grpcConn *grpc.ClientConn
	// signer is the signer.
	signer signer.Signer
	// chainExecutors is a map from chain ID -> chain executor.
	chainExecutors map[uint32]*chainExecutor
	// handler is the metrics handler.
	handler metrics.Handler
}

// logOrderInfo is a struct to keep track of the order of a log.
type logOrderInfo struct {
	blockNumber uint64
	blockIndex  uint
}

const (
	logChanSize          = 1000
	rpcRetry             = 7
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

// NewExecutor creates a new executor agent.
//
//nolint:cyclop
func NewExecutor(ctx context.Context, config config.Config, executorDB db.ExecutorDB, scribeClient client.ScribeClient, clients map[uint32]Backend, handler metrics.Handler) (*Executor, error) {
	chainExecutors := make(map[uint32]*chainExecutor)

	conn, grpcClient, err := makeScribeClient(ctx, handler, fmt.Sprintf("%s:%d", scribeClient.URL, scribeClient.Port))
	if err != nil {
		return nil, fmt.Errorf("could not create scribe client: %w", err)
	}

	executorSigner, err := agentsConfig.SignerFromConfig(ctx, config.UnbondedSigner)
	if err != nil {
		return nil, fmt.Errorf("could not create signer: %w", err)
	}

	if config.ExecuteInterval == 0 {
		config.ExecuteInterval = 2
	}

	if config.SetMinimumTimeInterval == 0 {
		config.SetMinimumTimeInterval = 2
	}

	for _, chain := range config.Chains {
		originParser, err := origin.NewParser(common.HexToAddress(chain.OriginAddress))
		if err != nil {
			return nil, fmt.Errorf("could not create origin parser: %w", err)
		}

		destinationParser, err := destination.NewParser(common.HexToAddress(chain.DestinationAddress))
		if err != nil {
			return nil, fmt.Errorf("could not create destination parser: %w", err)
		}

		var summitParserRef *summit.Parser

		if config.SummitChainID == chain.ChainID {
			summitParser, err := summit.NewParser(common.HexToAddress(config.SummitAddress))
			if err != nil {
				return nil, fmt.Errorf("could not create summit parser: %w", err)
			}

			summitParserRef = &summitParser
		}

		// chainRPCURL := fmt.Sprintf("%s/1/rpc/%d", config.BaseOmnirpcURL, chain.ChainID)

		underlyingClient, err := evmClient.DialBackendChainID(ctx, big.NewInt(int64(chain.ChainID)), chain.TempRPC, handler, evmClient.Capture(true))
		if err != nil {
			return nil, fmt.Errorf("could not get evm: %w", err)
		}

		boundDestination, err := evm.NewDestinationContract(ctx, underlyingClient, common.HexToAddress(chain.DestinationAddress))
		if err != nil {
			return nil, fmt.Errorf("could not bind destination contract: %w", err)
		}

		tree, err := newTreeFromDB(ctx, chain.ChainID, executorDB)
		if err != nil {
			return nil, fmt.Errorf("could not get tree from db: %w", err)
		}

		chainExecutors[chain.ChainID] = &chainExecutor{
			chainID: chain.ChainID,
			lastLog: &logOrderInfo{
				blockNumber: 0,
				blockIndex:  0,
			},
			closeConnection:   make(chan bool, 1),
			stopListenChan:    make(chan bool, 1),
			originParser:      originParser,
			destinationParser: destinationParser,
			summitParser:      summitParserRef,
			logChan:           make(chan *ethTypes.Log, logChanSize),
			merkleTree:        tree,
			rpcClient:         clients[chain.ChainID],
			boundDestination:  boundDestination,
		}
	}

	return &Executor{
		config:         config,
		executorDB:     executorDB,
		scribeClient:   scribeClient,
		grpcConn:       conn,
		grpcClient:     grpcClient,
		signer:         executorSigner,
		chainExecutors: chainExecutors,
		handler:        handler,
	}, nil
}

// Run starts the executor agent. It calls `Start` and `Listen`.
func (e Executor) Run(parentCtx context.Context) error {
	g, ctx := errgroup.WithContext(parentCtx)

	// Listen for snapshotAcceptedEvents on summit.
	g.Go(func() error {
		return e.streamLogs(ctx, e.grpcClient, e.grpcConn, e.config.SummitChainID, e.config.SummitAddress, nil, contractEventType{
			contractType: execTypes.SummitContract,
			eventType:    snapshotAcceptedEvent,
		})
	})

	for _, chain := range e.config.Chains {
		chain := chain

		// Listen for dispatchEvents on origin.
		g.Go(func() error {
			return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chain.ChainID, chain.OriginAddress, nil, contractEventType{
				contractType: execTypes.OriginContract,
				eventType:    dispatchedEvent,
			})
		})

		// Listen for attestationAcceptedEvents on destination.
		g.Go(func() error {
			return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chain.ChainID, chain.DestinationAddress, nil, contractEventType{
				contractType: execTypes.DestinationContract,
				eventType:    attestationAcceptedEvent,
			})
		})

		g.Go(func() error {
			return e.receiveLogs(ctx, chain.ChainID)
		})

		g.Go(func() error {
			return e.setMinimumTime(ctx, chain.ChainID)
		})

		g.Go(func() error {
			return e.executeExecutable(ctx, chain.ChainID)
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error in executor agent: %w", err)
	}

	return nil
}

// Stop stops the executor agent.
func (e Executor) Stop(chainID uint32) {
	e.chainExecutors[chainID].closeConnection <- true
	e.chainExecutors[chainID].stopListenChan <- true
}

// Execute calls execute on `Destination.sol` on the destination chain, after verifying the message.
// TODO: Use multi-call to batch execute.
//
//nolint:cyclop
func (e Executor) Execute(parentCtx context.Context, message types.Message) (_ bool, err error) {
	originDomain := message.OriginDomain()
	destinationDomain := message.DestinationDomain()

	testLeaf, err := message.ToLeaf()
	if err != nil {
		return false, fmt.Errorf("could not get test leaf: %w", err)
	}

	ctx, span := e.handler.Tracer().Start(parentCtx, "Execute", trace.WithAttributes(
		attribute.Int(metrics.Origin, int(originDomain)),
		attribute.String("leaf", common.Bytes2Hex(testLeaf[:])),
		attribute.Int(metrics.Nonce, int(message.Nonce())),
		attribute.Int(metrics.Destination, int(message.DestinationDomain())),
	))

	span.AddEvent("sanity check", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(1)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	nonce, err := e.verifyMessageOptimisticPeriod(ctx, message)
	if err != nil {
		return false, fmt.Errorf("could not verify optimistic period: %w", err)
	}

	if nonce == nil {
		return false, nil
	}

	maximumNonce := e.chainExecutors[message.OriginDomain()].merkleTree.NumOfItems()
	state, err := e.executorDB.GetEarliestStateInRange(ctx, originDomain, destinationDomain, *nonce, maximumNonce, e.config.DBPrefix)
	if err != nil {
		return false, fmt.Errorf("could not get earliest attestation nonce: %w", err)
	}

	if state == nil {
		return false, nil
	}

	proof, err := e.chainExecutors[message.OriginDomain()].merkleTree.MerkleProof(*nonce-1, (*state).Nonce())

	if err != nil {
		return false, fmt.Errorf("could not get merkle proof: %w", err)
	}

	verifiedMessageProof, err := e.verifyMessageMerkleProof(message)
	if err != nil {
		return false, fmt.Errorf("could not verify merkle proof: %w", err)
	}

	if !verifiedMessageProof {
		return false, nil
	}

	verifiedStateProof, err := e.verifyStateMerkleProof(ctx, *state)
	if err != nil {
		return false, fmt.Errorf("could not verify state merkle proof: %w", err)
	}

	if !verifiedStateProof {
		return false, nil
	}

	root := (*state).Root()
	stateRootString := common.BytesToHash(root[:]).String()
	origin := (*state).Origin()
	stateNonce := (*state).Nonce()
	stateMask := execTypes.DBState{
		Root:    &stateRootString,
		ChainID: &origin,
		Nonce:   &stateNonce,
	}

	_, snapshotProof, _, stateIndex, err := e.executorDB.GetStateMetadata(ctx, stateMask)
	if err != nil {
		return false, fmt.Errorf("could not get state index: %w", err)
	}

	if snapshotProof == nil || stateIndex == nil {
		return false, nil
	}

	var originProof [32][32]byte
	for i, p := range proof {
		copy(originProof[i][:], p)
	}

	var snapshotProofBytes [][]byte
	err = json.Unmarshal(*snapshotProof, &snapshotProofBytes)
	if err != nil {
		return false, fmt.Errorf("could not unmarshal proof: %w", err)
	}

	var snapshotProofB32 [][32]byte
	for _, p := range snapshotProofBytes {
		var p32 [32]byte
		copy(p32[:], p)
		snapshotProofB32 = append(snapshotProofB32, p32)
	}

	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    30 * time.Millisecond,
		Max:    3 * time.Second,
	}

	timeout := time.Duration(0)

	for {
		select {
		case <-ctx.Done():
			return false, fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(timeout):
			if b.Attempt() >= rpcRetry {
				return false, fmt.Errorf("could not execute message after %f attempts", b.Attempt())
			}

			err = e.chainExecutors[message.DestinationDomain()].boundDestination.Execute(ctx, e.signer, message, originProof, snapshotProofB32, big.NewInt(int64(*stateIndex)))
			if err != nil {
				timeout = b.Duration()
				span.AddEvent("error when executing", trace.WithAttributes(
					attribute.Int(metrics.ChainID, int(message.DestinationDomain())),
					attribute.String("error", err.Error()),
					attribute.Float64("timeout", timeout.Seconds()),
				))
				continue
			}

			return true, nil
		}
	}
}

type eventType int

const (
	// Origin's Dispatched event.
	dispatchedEvent eventType = iota
	// Destination's AttestationAccepted event.
	attestationAcceptedEvent
	// Summit's SnapshotAccepted event.
	snapshotAcceptedEvent
	otherEvent
)

type contractEventType struct {
	contractType execTypes.ContractType
	eventType    eventType
}

// verifyMessageMerkleProof verifies a message against the merkle tree at the state of the given nonce.
func (e Executor) verifyMessageMerkleProof(message types.Message) (bool, error) {
	root, err := e.chainExecutors[message.OriginDomain()].merkleTree.Root(message.Nonce())
	if err != nil {
		return false, fmt.Errorf("could not get root: %w", err)
	}

	proof, err := e.chainExecutors[message.OriginDomain()].merkleTree.MerkleProof(message.Nonce()-1, message.Nonce())
	if err != nil {
		return false, fmt.Errorf("could not get merkle proof: %w", err)
	}

	leaf, err := message.ToLeaf()
	if err != nil {
		return false, fmt.Errorf("could not convert message to leaf: %w", err)
	}

	inTree := merkle.VerifyMerkleProof(root, leaf[:], message.Nonce()-1, proof, merkle.MessageTreeDepth)

	return inTree, nil
}

// verifyStateMerkleProof verifies that a state is in the snapshot merkle tree.
//
//nolint:unused
func (e Executor) verifyStateMerkleProof(parentCtx context.Context, state types.State) (_ bool, err error) {
	stateRoot := state.Root()
	root := common.BytesToHash(stateRoot[:]).String()
	chainID := state.Origin()

	ctx, span := e.handler.Tracer().Start(parentCtx, "verifyStateMerkleProof", trace.WithAttributes(
		attribute.String("root", root),
		attribute.Int(metrics.ChainID, int(chainID)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	stateMask := execTypes.DBState{
		Root:    &root,
		ChainID: &chainID,
	}

	snapshotRoot, proof, treeHeight, stateIndex, err := e.executorDB.GetStateMetadata(ctx, stateMask)
	if err != nil {
		return false, fmt.Errorf("could not get snapshot root: %w", err)
	}

	if snapshotRoot == nil || proof == nil || treeHeight == nil || stateIndex == nil {
		return false, nil
	}

	leaf, _, err := state.SubLeaves()
	if err != nil {
		return false, fmt.Errorf("could not hash state: %w", err)
	}

	var proofBytes [][]byte
	err = json.Unmarshal(*proof, &proofBytes)
	if err != nil {
		return false, fmt.Errorf("could not unmarshal proof: %w", err)
	}

	inTree := merkle.VerifyMerkleProof((*snapshotRoot)[:], leaf[:], (*stateIndex)*2, proofBytes, *treeHeight)

	return inTree, nil
}

// verifyMessageOptimisticPeriod verifies that the optimistic period is valid.
func (e Executor) verifyMessageOptimisticPeriod(parentCtx context.Context, message types.Message) (msgNonce *uint32, err error) {
	chainID := message.OriginDomain()
	destinationDomain := message.DestinationDomain()
	nonce := message.Nonce()

	ctx, span := e.handler.Tracer().Start(parentCtx, "verifyMessageOptimisticPeriod", trace.WithAttributes(
		attribute.Int(metrics.Origin, int(chainID)),
		attribute.Int(metrics.Destination, int(destinationDomain)),
		attribute.Int(metrics.Nonce, int(nonce)),
	))

	defer func() {
		span.AddEvent("determine execution status", trace.WithAttributes(attribute.Bool("shouldExecute", msgNonce != nil)))
		metrics.EndSpanWithErr(span, err)
	}()

	messageMask := execTypes.DBMessage{
		ChainID:     &chainID,
		Destination: &destinationDomain,
		Nonce:       &nonce,
	}

	messageMinimumTime, err := e.executorDB.GetMessageMinimumTime(ctx, messageMask)
	if err != nil {
		return nil, fmt.Errorf("could not get attestation block time: %w", err)
	}

	if messageMinimumTime == nil {
		//nolint:nilnil
		return nil, nil
	}

	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    30 * time.Millisecond,
		Max:    3 * time.Second,
	}

	timeout := time.Duration(0)

	var currentTime uint64

retryLoop:
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(timeout):
			if b.Attempt() >= rpcRetry {
				return nil, fmt.Errorf("could not get latest header: %w", err)
			}

			latestHeader, err := e.chainExecutors[destinationDomain].rpcClient.HeaderByNumber(ctx, nil)
			if err != nil {
				timeout = b.Duration()

				continue
			}

			currentTime = latestHeader.Time

			break retryLoop
		}
	}

	if *messageMinimumTime > currentTime {
		//nolint:nilnil
		return nil, nil
	}

	return &nonce, nil
}

// newTreeFromDB creates a new merkle tree from the database's messages.
func newTreeFromDB(ctx context.Context, chainID uint32, executorDB db.ExecutorDB) (*merkle.HistoricalTree, error) {
	var allMessages []types.Message

	messageMask := execTypes.DBMessage{
		ChainID: &chainID,
	}
	page := 1

	for {
		messages, err := executorDB.GetMessages(ctx, messageMask, page)
		if err != nil {
			return nil, fmt.Errorf("could not get messages: %w", err)
		}
		if len(messages) == 0 {
			break
		}

		allMessages = append(allMessages, messages...)
		page++
	}

	rawMessages := make([][]byte, len(allMessages))

	for i, message := range allMessages {
		rawMessage, err := message.ToLeaf()
		if err != nil {
			return nil, fmt.Errorf("could not convert message to leaf: %w", err)
		}

		rawMessages[i] = rawMessage[:]
	}

	merkleTree := merkle.NewTreeFromItems(rawMessages, merkle.MessageTreeDepth)

	return merkleTree, nil
}

// checkIfExecuted checks if a message has been executed.
func (e Executor) checkIfExecuted(parentCtx context.Context, message types.Message) (_ bool, err error) {
	ctx, span := e.handler.Tracer().Start(parentCtx, "checkIfExecuted", trace.WithAttributes(
		attribute.Int(metrics.Origin, int(message.OriginDomain())),
		attribute.Int(metrics.Destination, int(message.DestinationDomain())),
		attribute.Int(metrics.Nonce, int(message.Nonce())),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    30 * time.Millisecond,
		Max:    3 * time.Second,
	}

	timeout := time.Duration(0)

	for {
		select {
		case <-ctx.Done():
			return false, fmt.Errorf("context cancelled: %w", ctx.Err())
		case <-time.After(timeout):
			if b.Attempt() >= 5 {
				return false, fmt.Errorf("could not get executed status: %w", ctx.Err())
			}

			executed, err := e.chainExecutors[message.DestinationDomain()].boundDestination.MessageStatus(ctx, message)
			if err != nil {
				timeout = b.Duration()
				span.AddEvent("could not get executed status",
					trace.WithAttributes(attribute.String("error", err.Error())),
					trace.WithAttributes(attribute.String("timeout", timeout.String())),
				)
				continue
			}

			if executed {
				span.AddEvent("message executed")
				return true, nil
			}

			span.AddEvent("message not executed")
			return false, nil
		}
	}
}

// streamLogs uses gRPC to stream logs into a channel.
//
//nolint:cyclop
func (e Executor) streamLogs(ctx context.Context, grpcClient pbscribe.ScribeServiceClient, conn *grpc.ClientConn, chainID uint32, address string, toBlockNumber *uint64, contractEvent contractEventType) error {
	lastStoredBlock, err := e.executorDB.GetLastBlockNumber(ctx, chainID, contractEvent.contractType)
	if err != nil {
		return fmt.Errorf("could not get last stored block: %w", err)
	}

	fromBlock := strconv.FormatUint(lastStoredBlock, 16)

	toBlock := "latest"
	if toBlockNumber != nil {
		toBlock = strconv.FormatUint(*toBlockNumber, 16)
	}

	stream, err := grpcClient.StreamLogs(ctx, &pbscribe.StreamLogsRequest{
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
		case <-e.chainExecutors[chainID].closeConnection:
			err := stream.CloseSend()
			if err != nil {
				return fmt.Errorf("could not close stream: %w", err)
			}

			err = conn.Close()
			if err != nil {
				return fmt.Errorf("could not close connection: %w", err)
			}

			return nil
		default:
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				logger.Errorf("stream closed: %s", err)
				return nil
			}
			if err != nil {
				return fmt.Errorf("could not receive: %w", err)
			}

			log := response.Log.ToLog()
			if log == nil {
				logger.Errorf("could not convert log: %s", response.Log.TxHash.String())
				return fmt.Errorf("could not convert log")
			}

			_, span := e.handler.Tracer().Start(ctx, "executor.streamLog", trace.WithAttributes(
				attribute.Int(metrics.ChainID, int(chainID)),
				attribute.Int("contract", int(contractEvent.contractType)),
				attribute.Int("event", int(contractEvent.eventType)),
				attribute.String(metrics.TxHash, log.TxHash.String()),
			))

			//if contractEvent.eventType == dispatchedEvent && !e.chainExecutors[chainID].lastLog.verifyAfter(*log) {
			//	logger.Errorf("log is not in chronological order. last log blockNumber: %d, blockIndex: %d. this log blockNumber: %d, blockIndex: %d, txHash: %s", e.chainExecutors[chainID].lastLog.blockNumber, e.chainExecutors[chainID].lastLog.blockIndex, log.BlockNumber, log.Index, log.TxHash.String())
			//
			//	continue
			//}

			select {
			case <-ctx.Done():
				return fmt.Errorf("context cancelled: %w", ctx.Err())
			case e.chainExecutors[chainID].logChan <- log:
			}
			span.AddEvent("log sent to channel")
			e.chainExecutors[chainID].lastLog.blockNumber = log.BlockNumber
			e.chainExecutors[chainID].lastLog.blockIndex = log.Index

			span.End()
		}
	}
}

// processLog processes the log and updates the merkle tree.
//
//nolint:cyclop,gocognit
func (e Executor) processLog(parentCtx context.Context, log ethTypes.Log, chainID uint32) (err error) {
	contractEvent := e.logType(log, chainID)

	ctx, span := e.handler.Tracer().Start(parentCtx, "processLog", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(chainID)),
		attribute.Int("contract", int(contractEvent.contractType)),
		attribute.Int("event", int(contractEvent.eventType)),
		attribute.String(metrics.TxHash, log.TxHash.String()),
	))

	span.AddEvent("am i getting here", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(chainID)),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	switch contractEvent.contractType {
	case execTypes.OriginContract:
		span.AddEvent("origin", trace.WithAttributes(attribute.String("step", "a")))
		message, err := e.logToMessage(log, chainID)
		if err != nil {
			return fmt.Errorf("could not convert log to leaf: %w", err)
		}
		span.AddEvent("origin", trace.WithAttributes(attribute.String("step", "b")))

		if message == nil {
			return nil
		}

		span.AddEvent("origin", trace.WithAttributes(attribute.String("step", "c")))

		merkleIndex := e.chainExecutors[chainID].merkleTree.NumOfItems()
		leaf, err := (*message).ToLeaf()
		if err != nil {
			return fmt.Errorf("could not convert message to leaf: %w", err)
		}
		span.AddEvent("origin", trace.WithAttributes(attribute.String("step", "d")))
		logger.Errorf("message nonce is %d chain id is %d", (*message).Nonce(), chainID)

		// Make sure the nonce of the message is being inserted at the right index.
		switch {
		case merkleIndex+1 > (*message).Nonce():
			return nil
		case merkleIndex+1 < (*message).Nonce():
			return fmt.Errorf("nonce is not correct. expected: %d, got: %d", merkleIndex+1, (*message).Nonce())
		default:
		}

		span.AddEvent("origin", trace.WithAttributes(attribute.String("step", "e")))

		e.chainExecutors[chainID].merkleTree.Insert(leaf[:])

		err = e.executorDB.StoreMessage(ctx, *message, log.BlockNumber, false, 0)
		if err != nil {
			return fmt.Errorf("could not store message: %w", err)
		}
		span.AddEvent("origin", trace.WithAttributes(attribute.String("step", "f")))
	case execTypes.DestinationContract:
		span.AddEvent("destination", trace.WithAttributes(attribute.String("step", "a")))
		attestation, err := e.logToAttestation(log, chainID)
		if err != nil {
			return fmt.Errorf("could not convert log to attestation: %w", err)
		}
		span.AddEvent("destination", trace.WithAttributes(attribute.String("step", "b")))
		if attestation == nil {
			return nil
		}

		b := &backoff.Backoff{
			Factor: 2,
			Jitter: true,
			Min:    30 * time.Millisecond,
			Max:    3 * time.Second,
		}

		timeout := time.Duration(0)

		var logHeader *ethTypes.Header

	retryLoop:
		for {
			select {
			case <-ctx.Done():
				return fmt.Errorf("context canceled: %w", ctx.Err())
			case <-time.After(timeout):
				span.AddEvent("destination", trace.WithAttributes(attribute.String("step", "c")))
				if b.Attempt() >= rpcRetry {
					return fmt.Errorf("could not get log header: %w", err)
				}
				logHeader, err = e.chainExecutors[chainID].rpcClient.HeaderByNumber(ctx, big.NewInt(int64(log.BlockNumber)))
				if err != nil {
					timeout = b.Duration()

					continue
				}
				span.AddEvent("destination", trace.WithAttributes(attribute.String("step", "d")))

				break retryLoop
			}
		}

		err = e.executorDB.StoreAttestation(ctx, *attestation, chainID, log.BlockNumber, logHeader.Time)
		if err != nil {
			return fmt.Errorf("could not store attestation: %w", err)
		}
		span.AddEvent("destination", trace.WithAttributes(attribute.String("step", "e")))
	case execTypes.SummitContract:
		//nolint:gocritic,exhaustive
		span.AddEvent("summit", trace.WithAttributes(attribute.String("step", "a")))
		span.AddEvent("summit", trace.WithAttributes(attribute.String("step", "b")))
		snapshot, err := e.logToSnapshot(log, chainID)
		if err != nil {
			return fmt.Errorf("could not convert log to snapshot: %w", err)
		}
		span.AddEvent("summit", trace.WithAttributes(attribute.String("step", "c")))
		if snapshot == nil {
			return nil
		}
		span.AddEvent("summit", trace.WithAttributes(attribute.String("step", "e")))

		snapshotRoot, proofs, err := (*snapshot).SnapshotRootAndProofs()
		if err != nil {
			return fmt.Errorf("could not get snapshot root and proofs: %w", err)
		}
		span.AddEvent("summit", trace.WithAttributes(attribute.String("step", "f")))

		treeHeight := (*snapshot).TreeHeight()

		err = e.executorDB.StoreStates(ctx, (*snapshot).States(), snapshotRoot, proofs, treeHeight, log.BlockNumber)
		if err != nil {
			return fmt.Errorf("could not store states: %w", err)
		}
		span.AddEvent("summit", trace.WithAttributes(attribute.String("step", "g")))
	case execTypes.Other:
		span.AddEvent("Other contract event")
	default:
		return fmt.Errorf("log type not supported")
	}

	return nil
}

// receiveLogs receives logs from the log channel and processes them.
func (e Executor) receiveLogs(ctx context.Context, chainID uint32) error {
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context canceled: %w", ctx.Err())
		case <-e.chainExecutors[chainID].stopListenChan:
			return nil
		case log := <-e.chainExecutors[chainID].logChan:
			if log == nil {
				return fmt.Errorf("log is nil")
			}

			err := e.processLog(ctx, *log, chainID)
			if err != nil {
				return fmt.Errorf("could not process log: %w", err)
			}
		}
	}
}

// executeExecutable executes executable messages in the database.
//
//nolint:gocognit,cyclop
func (e Executor) executeExecutable(parentCtx context.Context, chainID uint32) (err error) {
	backoffInterval := time.Duration(0)
	for {
		select {
		case <-parentCtx.Done():
			return fmt.Errorf("context canceled: %w", parentCtx.Err())
		case <-time.After(backoffInterval):
			backoffInterval = time.Duration(e.config.ExecuteInterval) * time.Second

			page := 1
			currentTime := uint64(time.Now().Unix())

			messageMask := execTypes.DBMessage{
				ChainID: &chainID,
			}

			for {
				if chainID == 137 {
					fmt.Printf("")
				}
				messages, err := e.executorDB.GetExecutableMessages(parentCtx, messageMask, currentTime, page)
				if err != nil {
					return fmt.Errorf("could not get executable messages: %w", err)
				}

				if len(messages) == 0 {
					break
				}

				ctx, span := e.handler.Tracer().Start(parentCtx, "executeExecutable", trace.WithAttributes(
					attribute.Int(metrics.ChainID, int(chainID)),
					attribute.Int("num_messages", len(messages)),
					attribute.Int(metrics.Page, page),
				))

				for _, message := range messages {
					messageExecuted, err := e.checkIfExecuted(ctx, message)
					if err != nil {
						return fmt.Errorf("could not check if message was executed: %w", err)
					}

					span.AddEvent("checked if message was executed", trace.WithAttributes(
						attribute.Int(metrics.ChainID, int(chainID)),
						attribute.Int(metrics.Nonce, int(message.Nonce())),
						attribute.Bool("message_executed", messageExecuted),
					))

					if !messageExecuted {
						executed, err := e.Execute(ctx, message)
						if err != nil {
							logger.Errorf("could not execute message, retrying: %s", err)
							continue
						}

						if !executed {
							continue
						}
					}

					destinationDomain := message.DestinationDomain()
					nonce := message.Nonce()
					executedMessageMask := execTypes.DBMessage{
						ChainID:     &chainID,
						Destination: &destinationDomain,
						Nonce:       &nonce,
					}
					err = e.executorDB.ExecuteMessage(ctx, executedMessageMask)
					if err != nil {
						return fmt.Errorf("could not execute message: %w", err)
					}
				}

				metrics.EndSpanWithErr(span, err)

				page++
			}
		}
	}
}

// setMinimumTime sets the minimum time for the message to be executed by checking for associated attestations.
//
//nolint:gocognit,cyclop
func (e Executor) setMinimumTime(parentCtx context.Context, chainID uint32) (err error) {
	backoffInterval := time.Duration(0)

	for {
		select {
		case <-parentCtx.Done():
			return fmt.Errorf("context canceled: %w", parentCtx.Err())
		case <-time.After(backoffInterval):
			backoffInterval = time.Duration(e.config.SetMinimumTimeInterval) * time.Second

			page := 1
			messageMask := execTypes.DBMessage{
				ChainID: &chainID,
			}

			var unsetMessages []types.Message

			// Get all unset messages.
			for {
				messages, err := e.executorDB.GetUnsetMinimumTimeMessages(parentCtx, messageMask, page)
				if err != nil {
					return fmt.Errorf("could not get messages without minimum time: %w", err)
				}

				if len(messages) == 0 {
					break
				}

				unsetMessages = append(unsetMessages, messages...)

				page++
			}

			if len(unsetMessages) == 0 {
				continue
			}

			ctx, span := e.handler.Tracer().Start(parentCtx, "setMinimumTime", trace.WithAttributes(
				attribute.Int(metrics.ChainID, int(chainID)),
				attribute.Int("num_unset_messages", len(unsetMessages)),
				attribute.Int(metrics.Page, page),
			))

			for _, message := range unsetMessages {
				nonce := message.Nonce()
				destinationDomain := message.DestinationDomain()

				minimumTimestamp, err := e.executorDB.GetTimestampForMessage(ctx, chainID, destinationDomain, nonce, e.config.DBPrefix)
				if err != nil {
					return fmt.Errorf("could not get timestamp for message: %w", err)
				}

				if minimumTimestamp == nil {
					continue
				}

				setMessageMask := execTypes.DBMessage{
					ChainID:     &chainID,
					Destination: &destinationDomain,
					Nonce:       &nonce,
				}

				err = e.executorDB.SetMinimumTime(ctx, setMessageMask, *minimumTimestamp+uint64(message.OptimisticSeconds()))
				if err != nil {
					return fmt.Errorf("could not set minimum time: %w", err)
				}
			}

			metrics.EndSpanWithErr(span, err)
		}
	}
}
