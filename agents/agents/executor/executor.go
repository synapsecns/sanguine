package executor

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	execTypes "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"github.com/synapsecns/sanguine/agents/config/executor"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/merkle"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	ethergoChain "github.com/synapsecns/sanguine/ethergo/chain"
	agentsConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
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
	// lightInboxParser is the light inbox parser.
	lightInboxParser lightinbox.Parser
	// inboxParser is the inbox parser.
	inboxParser inbox.Parser
	// summitParser is the summit parser.
	summitParser summit.Parser
	// logChan is the log channel.
	logChan chan *ethTypes.Log
	// merkleTree is a merkle tree for a specific origin chain.
	merkleTree *merkle.HistoricalTree
	// rpcClient is an RPC client.
	rpcClient Backend
	// boundDestination is a bound destination contract.
	boundDestination domains.DestinationContract
	// boundOrigin is a bound origin contract.
	boundOrigin domains.OriginContract
}

// Executor is the executor agent.
type Executor struct {
	// config is the executor agent config.
	config executor.Config
	// executorDB is the executor agent database.
	executorDB db.ExecutorDB
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
	// txSubmitter is the transaction submitter.
	txSubmitter submitter.TransactionSubmitter
	// retryConfig is the retry configuration for RPC calls.
	retryConfig []retry.WithBackoffConfigurator
	// NowFunc returns the current time.
	NowFunc func() time.Time
}

// logOrderInfo is a struct to keep track of the order of a log.
type logOrderInfo struct {
	blockNumber uint64
	blockIndex  uint
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

// NewExecutor creates a new executor agent.
//
//nolint:cyclop
func NewExecutor(ctx context.Context, config executor.Config, executorDB db.ExecutorDB, scribeClient client.ScribeClient, omniRPCClient omnirpcClient.RPCClient, handler metrics.Handler) (*Executor, error) {
	chainExecutors := make(map[uint32]*chainExecutor)

	conn, grpcClient, err := makeScribeClient(ctx, handler, fmt.Sprintf("%s:%d", scribeClient.URL, scribeClient.Port))
	if err != nil {
		return nil, fmt.Errorf("could not create scribe client: %w", err)
	}

	executorSigner, err := agentsConfig.SignerFromConfig(ctx, config.UnbondedSigner)
	if err != nil {
		return nil, fmt.Errorf("could not create signer: %w", err)
	}

	txSubmitter := submitter.NewTransactionSubmitter(handler, executorSigner, omniRPCClient, executorDB.SubmitterDB(), &config.SubmitterConfig)

	if config.MaxRetrySeconds == 0 {
		config.MaxRetrySeconds = 30
	}

	retryConfig := []retry.WithBackoffConfigurator{
		retry.WithMaxAttemptTime(time.Second * time.Duration(config.MaxRetrySeconds)),
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

		lightInboxParser, err := lightinbox.NewParser(common.HexToAddress(chain.LightInboxAddress))
		if err != nil {
			return nil, fmt.Errorf("could not create destination parser: %w", err)
		}

		var inboxParser inbox.Parser
		var summitParser summit.Parser

		if config.SummitChainID == chain.ChainID {
			inboxParser, err = inbox.NewParser(common.HexToAddress(config.InboxAddress))
			if err != nil {
				return nil, fmt.Errorf("could not create inbox parser: %w", err)
			}

			summitParser, err = summit.NewParser(common.HexToAddress(config.SummitAddress))
			if err != nil {
				return nil, fmt.Errorf("could not create summit parser: %w", err)
			}
		} else {
			inboxParser = nil
			summitParser = nil
		}

		evmClient, err := omniRPCClient.GetConfirmationsClient(ctx, int(chain.ChainID), 1)
		if err != nil {
			return nil, fmt.Errorf("could not get evm client: %w", err)
		}

		chainClient, err := ethergoChain.NewFromURL(ctx, omniRPCClient.GetEndpoint(int(chain.ChainID), 1))
		if err != nil {
			return nil, fmt.Errorf("could not create chain client: %w", err)
		}

		boundDestination, err := evm.NewDestinationContract(ctx, chainClient, common.HexToAddress(chain.DestinationAddress))
		if err != nil {
			return nil, fmt.Errorf("could not bind destination contract: %w", err)
		}

		boundOrigin, err := evm.NewOriginContract(ctx, chainClient, common.HexToAddress(chain.OriginAddress))
		if err != nil {
			return nil, fmt.Errorf("could not bind origin contract: %w", err)
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
			closeConnection:  make(chan bool, 1),
			stopListenChan:   make(chan bool, 1),
			originParser:     originParser,
			lightInboxParser: lightInboxParser,
			inboxParser:      inboxParser,
			summitParser:     summitParser,
			logChan:          make(chan *ethTypes.Log, logChanSize),
			merkleTree:       tree,
			rpcClient:        evmClient,
			boundDestination: boundDestination,
			boundOrigin:      boundOrigin,
		}
	}

	return &Executor{
		config:         config,
		executorDB:     executorDB,
		grpcConn:       conn,
		grpcClient:     grpcClient,
		signer:         executorSigner,
		chainExecutors: chainExecutors,
		handler:        handler,
		txSubmitter:    txSubmitter,
		retryConfig:    retryConfig,
		NowFunc:        time.Now,
	}, nil
}

// Run starts the executor agent. It calls `Start` and `Listen`.
func (e Executor) Run(parentCtx context.Context) error {
	g, ctx := errgroup.WithContext(parentCtx)

	g.Go(func() error {
		err := e.txSubmitter.Start(ctx)
		if err != nil {
			err = fmt.Errorf("could not start tx submitter: %w", err)
		}
		return err
	})

	// Listen for snapshotAccepted events on the inbox.
	g.Go(func() error {
		return e.streamLogs(ctx, e.grpcClient, e.grpcConn, e.config.SummitChainID, e.config.InboxAddress, execTypes.InboxContract)
	})

	// Listen for attestationSaved events on the summit.
	g.Go(func() error {
		return e.streamLogs(ctx, e.grpcClient, e.grpcConn, e.config.SummitChainID, e.config.SummitAddress, execTypes.SummitContract)
	})

	for _, chain := range e.config.Chains {
		chain := chain

		// Listen for sent events on origins.
		g.Go(func() error {
			return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chain.ChainID, chain.OriginAddress, execTypes.OriginContract)
		})

		// Listen for attestationAccepted events on destination.
		g.Go(func() error {
			return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chain.ChainID, chain.LightInboxAddress, execTypes.LightInboxContract)
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

	ctx, span := e.handler.Tracer().Start(parentCtx, "Execute", trace.WithAttributes(
		attribute.Int(metrics.Origin, int(originDomain)),
		attribute.Int(metrics.Destination, int(destinationDomain)),
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
	state, err := e.executorDB.GetEarliestStateInRange(ctx, originDomain, destinationDomain, *nonce, maximumNonce)
	if err != nil {
		return false, fmt.Errorf("could not get earliest attestation nonce: %w", err)
	}

	if state == nil {
		return false, nil
	}

	var proof [][]byte
	contractCall := func(ctx context.Context) error {
		proof, err = e.chainExecutors[message.OriginDomain()].merkleTree.MerkleProof(*nonce-1, (*state).Nonce())
		if err != nil {
			return fmt.Errorf("could not get merkle proof: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, e.retryConfig...)
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
	stateMask := db.DBState{
		Root:    &stateRootString,
		ChainID: &origin,
		Nonce:   &stateNonce,
	}

	_, snapshotProof, stateIndex, err := e.executorDB.GetStateMetadata(ctx, stateMask)
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

	_, err = e.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(destinationDomain)), func(transactor *bind.TransactOpts) (tx *ethTypes.Transaction, err error) {
		tx, err = e.chainExecutors[message.DestinationDomain()].boundDestination.Execute(
			transactor,
			message,
			originProof,
			snapshotProofB32,
			big.NewInt(int64(*stateIndex)),
			uint64(1000000),
		)
		if err != nil {
			return nil, fmt.Errorf("could not execute message: %w", err)
		}

		return
	})
	if err != nil {
		return false, fmt.Errorf("could not submit transaction: %w", err)
	}

	return true, nil
}

// verifyMessageMerkleProof verifies a message against the merkle tree at the state of the given nonce.
func (e Executor) verifyMessageMerkleProof(message types.Message) (bool, error) {
	var root []byte
	contractCall := func(ctx context.Context) error {
		var err error
		root, err = e.chainExecutors[message.OriginDomain()].merkleTree.Root(message.Nonce())
		if err != nil {
			return fmt.Errorf("could not get root: %w", err)
		}

		return nil
	}
	err := retry.WithBackoff(context.Background(), contractCall, e.retryConfig...)
	if err != nil {
		return false, fmt.Errorf("could not get root: %w", err)
	}

	var proof [][]byte
	contractCall = func(ctx context.Context) error {
		proof, err = e.chainExecutors[message.OriginDomain()].merkleTree.MerkleProof(message.Nonce()-1, message.Nonce())
		if err != nil {
			return fmt.Errorf("could not get merkle proof: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(context.Background(), contractCall, e.retryConfig...)
	if err != nil {
		return false, fmt.Errorf("could not get merkle proof: %w", err)
	}

	leaf, err := message.ToLeaf()
	if err != nil {
		return false, fmt.Errorf("could not convert message to leaf: %w", err)
	}

	inTree := merkle.VerifyMerkleProof(root, leaf[:], message.Nonce()-1, proof, merkle.MessageTreeHeight)

	return inTree, nil
}

// verifyStateMerkleProof verifies that a state is in the snapshot merkle tree.
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

	stateMask := db.DBState{
		Root:    &root,
		ChainID: &chainID,
	}

	snapshotRoot, proof, stateIndex, err := e.executorDB.GetStateMetadata(ctx, stateMask)
	if err != nil {
		return false, fmt.Errorf("could not get snapshot root: %w", err)
	}

	if snapshotRoot == nil || proof == nil || stateIndex == nil {
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

	inTree := merkle.VerifyMerkleProof((*snapshotRoot)[:], leaf[:], (*stateIndex)*2, proofBytes, merkle.SnapshotTreeHeight)

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
		span.AddEvent("Determine execution status", trace.WithAttributes(attribute.Bool("should_execute", msgNonce != nil)))
		metrics.EndSpanWithErr(span, err)
	}()

	messageMask := db.DBMessage{
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

	var currentTime uint64
	chainCall := func(ctx context.Context) error {
		var err error
		latestHeader, err := e.chainExecutors[chainID].rpcClient.HeaderByNumber(ctx, nil)
		if err != nil {
			return fmt.Errorf("could not get latest header: %w", err)
		}

		if latestHeader == nil {
			return fmt.Errorf("latest header is nil")
		}

		currentTime = latestHeader.Time

		return nil
	}
	err = retry.WithBackoff(ctx, chainCall, e.retryConfig...)
	if err != nil {
		return nil, fmt.Errorf("could not get latest header: %w", err)
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

	messageMask := db.DBMessage{
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

	merkleTree := merkle.NewTreeFromItems(rawMessages, merkle.MessageTreeHeight)

	return merkleTree, nil
}

// checkIfExecuted checks if a message has been executed.
// TODO: Needs to be tested.
func (e Executor) checkIfExecuted(parentCtx context.Context, message types.Message) (_ bool, err error) {
	ctx, span := e.handler.Tracer().Start(parentCtx, "checkIfExecuted", trace.WithAttributes(
		attribute.Int(metrics.Origin, int(message.OriginDomain())),
		attribute.Int(metrics.Destination, int(message.DestinationDomain())),
		attribute.Int(metrics.Nonce, int(message.Nonce())),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	var executed uint8
	contractCall := func(ctx context.Context) error {
		var err error
		executed, err = e.chainExecutors[message.DestinationDomain()].boundDestination.MessageStatus(ctx, message)
		if err != nil {
			return fmt.Errorf("could not get executed status: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, e.retryConfig...)
	if err != nil {
		return false, fmt.Errorf("could not get executed status: %w", err)
	}

	if execTypes.MessageStatusType(executed) == execTypes.Success {
		span.AddEvent("message executed")
		return true, nil
	}

	span.AddEvent("message not executed")
	return false, nil
}

// streamLogs uses gRPC to stream logs into a channel.
//
//nolint:cyclop
func (e Executor) streamLogs(ctx context.Context, grpcClient pbscribe.ScribeServiceClient, conn *grpc.ClientConn, chainID uint32, address string, contractType execTypes.ContractType) error {
	lastStoredBlock, err := e.executorDB.GetLastBlockNumber(ctx, chainID, contractType)
	if err != nil {
		return fmt.Errorf("could not get last stored block: %w", err)
	}

	fromBlock := strconv.FormatUint(lastStoredBlock, 16)

	toBlock := "latest"

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
			if err != nil {
				return fmt.Errorf("could not receive: %w", err)
			}

			log := response.Log.ToLog()
			if log == nil {
				return fmt.Errorf("could not convert log")
			}

			// We do not use a span context here because this is just meant to track transactions coming in.
			_, span := e.handler.Tracer().Start(ctx, "executor.streamLog", trace.WithAttributes(
				attribute.Int(metrics.ChainID, int(chainID)),
				attribute.String(metrics.Contract, contractType.String()),
				attribute.String(metrics.TxHash, log.TxHash.String()),
			))

			select {
			case <-ctx.Done():
				return fmt.Errorf("context canceled: %w", ctx.Err())
			case e.chainExecutors[chainID].logChan <- log:
			}
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
	datatypeInterface, err := e.logToInterface(log, chainID)
	if err != nil {
		return fmt.Errorf("could not convert log to interface: %w", err)
	}
	if datatypeInterface == nil {
		return nil
	}

	ctx, span := e.handler.Tracer().Start(parentCtx, "processLog", trace.WithAttributes(
		attribute.Int(metrics.ChainID, int(chainID)),
		attribute.String(metrics.TxHash, log.TxHash.String()),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	switch datatype := datatypeInterface.(type) {
	case types.Message:
		return e.processMessage(ctx, datatype, log.BlockNumber)
	case types.Snapshot:
		return e.processSnapshot(ctx, datatype, log.BlockNumber)
	case types.Attestation:
		return e.processAttestation(ctx, datatype, chainID, log.BlockNumber)
	default:
		return fmt.Errorf("type not supported")
	}
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
			currentTime := uint64(e.NowFunc().Unix())

			messageMask := db.DBMessage{
				ChainID: &chainID,
			}

			for {
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
						attribute.Bool(metrics.MessageExecuted, messageExecuted),
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
					executedMessageMask := db.DBMessage{
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
			messageMask := db.DBMessage{
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

				minimumTimestamp, err := e.executorDB.GetTimestampForMessage(ctx, chainID, destinationDomain, nonce)
				if err != nil {
					return fmt.Errorf("could not get timestamp for message: %w", err)
				}

				if minimumTimestamp == nil {
					continue
				}

				setMessageMask := db.DBMessage{
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
