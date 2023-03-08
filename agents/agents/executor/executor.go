package executor

import (
	"context"
	"errors"
	"fmt"
	"github.com/jpillora/backoff"
	"io"
	"math/big"
	"strconv"
	"time"

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
	ethergoChain "github.com/synapsecns/sanguine/ethergo/chain"
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
	// logChan is the log channel.
	logChan chan *ethTypes.Log
	// merkleTree is a merkle tree for a specific origin chain.
	merkleTree *merkle.HistoricalTree
	// rpcClient is an RPC client.
	rpcClient Backend
	// boundDestination is a bound destination contract.
	boundDestination domains.DestinationContract
	// executed is a map from hash(origin chain ID, destination chain ID, nonce) -> bool.
	executed map[[32]byte]bool
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
}

// logOrderInfo is a struct to keep track of the order of a log.
type logOrderInfo struct {
	blockNumber uint64
	blockIndex  uint
}

const logChanSize = 1000

// NewExecutor creates a new executor agent.
//
//nolint:cyclop
func NewExecutor(ctx context.Context, config config.Config, executorDB db.ExecutorDB, scribeClient client.ScribeClient, clients map[uint32]Backend) (*Executor, error) {
	chainExecutors := make(map[uint32]*chainExecutor)
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", scribeClient.URL, scribeClient.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("could not dial grpc: %w", err)
	}

	grpcClient := pbscribe.NewScribeServiceClient(conn)

	// Ensure that gRPC is up and running.
	healthCheck, err := grpcClient.Check(ctx, &pbscribe.HealthCheckRequest{}, grpc.WaitForReady(true))
	if err != nil {
		return nil, fmt.Errorf("could not check: %w", err)
	}
	if healthCheck.Status != pbscribe.HealthCheckResponse_SERVING {
		return nil, fmt.Errorf("not serving: %s", healthCheck.Status)
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

		chainRPCURL := fmt.Sprintf("%s/1/rpc/%d", config.BaseOmnirpcURL, chain.ChainID)

		underlyingClient, err := ethergoChain.NewFromURL(ctx, chainRPCURL)
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
			logChan:           make(chan *ethTypes.Log, logChanSize),
			merkleTree:        tree,
			rpcClient:         clients[chain.ChainID],
			boundDestination:  boundDestination,
			executed:          make(map[[32]byte]bool),
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
	}, nil
}

// Run starts the executor agent. It calls `Start` and `Listen`.
func (e Executor) Run(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	// Backfill executed messages.
	for _, chain := range e.config.Chains {
		chain := chain

		g.Go(func() error {
			return e.markAsExecuted(ctx, chain)
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not backfill executed messages: %w", err)
	}

	for _, chain := range e.config.Chains {
		chain := chain

		g.Go(func() error {
			return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chain, nil, contractEventType{
				contractType: originContract,
				eventType:    dispatchEvent,
			})
		})

		g.Go(func() error {
			return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chain, nil, contractEventType{
				contractType: destinationContract,
				eventType:    attestationAcceptedEvent,
			})
		})

		g.Go(func() error {
			return e.receiveLogs(ctx, chain.ChainID)
		})

		for _, destinationChain := range e.config.Chains {
			destinationChain := destinationChain
			if destinationChain.ChainID == chain.ChainID {
				continue
			}
			g.Go(func() error {
				return e.setMinimumTime(ctx, chain.ChainID, destinationChain.ChainID)
			})
		}

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
func (e Executor) Execute(ctx context.Context, message types.Message) (bool, error) {
	nonce, err := e.verifyMessageOptimisticPeriod(ctx, message)
	if err != nil {
		return false, fmt.Errorf("could not verify optimistic period: %w", err)
	}

	if nonce == nil {
		return false, nil
	}

	originDomain := message.OriginDomain()
	destinationDomain := message.DestinationDomain()
	attestationMask := execTypes.DBAttestation{
		ChainID:     &originDomain,
		Destination: &destinationDomain,
	}
	maximumNonce := e.chainExecutors[message.OriginDomain()].merkleTree.NumOfItems()
	itemCountNonce, err := e.executorDB.GetEarliestAttestationsNonceInNonceRange(ctx, attestationMask, *nonce, maximumNonce)
	if err != nil {
		return false, fmt.Errorf("could not get earliest attestation nonce: %w", err)
	}

	if itemCountNonce == nil {
		return false, nil
	}

	proof, err := e.chainExecutors[message.OriginDomain()].merkleTree.MerkleProof(*nonce-1, *itemCountNonce)

	if err != nil {
		return false, fmt.Errorf("could not get merkle proof: %w", err)
	}

	verified, err := e.verifyMessageMerkleProof(message)
	if err != nil {
		return false, fmt.Errorf("could not verify merkle proof: %w", err)
	}

	if !verified {
		return false, nil
	}

	index := big.NewInt(int64(*nonce - 1))

	var proofB32 [32][32]byte
	for i, p := range proof {
		copy(proofB32[i][:], p)
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
			if b.Attempt() >= 5 {
				return false, fmt.Errorf("could not execute message after %f attempts", b.Attempt())
			}

			err = e.chainExecutors[message.DestinationDomain()].boundDestination.Execute(ctx, e.signer, message, proofB32, index)
			if err != nil {
				timeout = b.Duration()
				logger.Errorf("got error %v when trying to execute the message on chain %d. trying again in %f seconds", err, message.DestinationDomain(), timeout.Seconds())
				continue
			}

			return true, nil
		}
	}
}

type contractType int

type eventType int

const (
	originContract contractType = iota
	destinationContract
	summitContract
	other
)

const (
	// Origin's Dispatch event.
	dispatchEvent eventType = iota
	// Destination's AttestationAccepted event.
	attestationAcceptedEvent
	// Destination's AttestationExecuted event.
	executedEvent
	// Summit's SnapshotAccepted event.
	//nolint:deadcode,varcheck
	snapshotAcceptedEvent
	otherEvent
)

type contractEventType struct {
	contractType contractType
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

// verifyMessageOptimisticPeriod verifies that the optimistic period is valid.
func (e Executor) verifyMessageOptimisticPeriod(ctx context.Context, message types.Message) (*uint32, error) {
	chainID := message.OriginDomain()
	destinationDomain := message.DestinationDomain()
	nonce := message.Nonce()
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

	latestHeader, err := e.chainExecutors[destinationDomain].rpcClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("could not get latest header: %w", err)
	}

	currentTime := latestHeader.Time
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

// markAsExecuted marks a message as executed via the `executed` mapping.
func (e Executor) markAsExecuted(ctx context.Context, chain config.ChainConfig) error {
	latestHeader, err := e.chainExecutors[chain.ChainID].rpcClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not get latest header: %w", err)
	}

	blockNumber := latestHeader.Number.Uint64()

	return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chain, &blockNumber, contractEventType{
		contractType: destinationContract,
		eventType:    executedEvent,
	})
}

// streamLogs uses gRPC to stream logs into a channel.
//
//nolint:cyclop
func (e Executor) streamLogs(ctx context.Context, grpcClient pbscribe.ScribeServiceClient, conn *grpc.ClientConn, chain config.ChainConfig, toBlockNumber *uint64, contractEvent contractEventType) error {
	var address string

	//nolint:exhaustive
	switch contractEvent.contractType {
	case originContract:
		address = chain.OriginAddress
	case destinationContract:
		address = chain.DestinationAddress
	default:
		return fmt.Errorf("contract type not supported")
	}

	lastStoredBlock, err := e.executorDB.GetLastBlockNumber(ctx, chain.ChainID)
	if err != nil {
		return fmt.Errorf("could not get last stored block: %w", err)
	}

	fromBlock := strconv.FormatUint(lastStoredBlock, 10)

	toBlock := "latest"
	if toBlockNumber != nil {
		toBlock = strconv.FormatUint(*toBlockNumber, 10)
	}

	stream, err := grpcClient.StreamLogs(ctx, &pbscribe.StreamLogsRequest{
		Filter: &pbscribe.LogFilter{
			ContractAddress: &pbscribe.NullableString{Kind: &pbscribe.NullableString_Data{Data: address}},
			ChainId:         chain.ChainID,
		},
		FromBlock: fromBlock,
		ToBlock:   toBlock,
	})
	if err != nil {
		return fmt.Errorf("could not stream logs: %w", err)
	}

	for {
		select {
		case <-e.chainExecutors[chain.ChainID].closeConnection:
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
				return nil
			}
			if err != nil {
				return fmt.Errorf("could not receive: %w", err)
			}

			log := response.Log.ToLog()
			if log == nil {
				return fmt.Errorf("could not convert log")
			}

			// If we are filtering for `executed` events, we do not need to `verifyAfter`
			// since we are backfilling.
			if contractEvent.eventType == executedEvent {
				e.chainExecutors[chain.ChainID].logChan <- log

				continue
			}
			if !e.chainExecutors[chain.ChainID].lastLog.verifyAfter(*log) {
				logger.Warnf("log is not in chronological order. last log blockNumber: %d, blockIndex: %d. this log blockNumber: %d, blockIndex: %d, txHash: %s", e.chainExecutors[chain.ChainID].lastLog.blockNumber, e.chainExecutors[chain.ChainID].lastLog.blockIndex, log.BlockNumber, log.Index, log.TxHash.String())

				continue
			}

			e.chainExecutors[chain.ChainID].logChan <- log
			e.chainExecutors[chain.ChainID].lastLog.blockNumber = log.BlockNumber
			e.chainExecutors[chain.ChainID].lastLog.blockIndex = log.Index
		}
	}
}

// processLog processes the log and updates the merkle tree.
//
//nolint:cyclop,gocognit
func (e Executor) processLog(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	contractEvent := e.logType(log, chainID)

	switch contractEvent.contractType {
	case originContract:
		message, err := e.logToMessage(log, chainID)
		if err != nil {
			return fmt.Errorf("could not convert log to leaf: %w", err)
		}

		if message == nil {
			return nil
		}

		destination := (*message).DestinationDomain()
		merkleIndex := e.chainExecutors[chainID].merkleTree.NumOfItems()
		leaf, err := (*message).ToLeaf()
		if err != nil {
			return fmt.Errorf("could not convert message to leaf: %w", err)
		}

		// Make sure the nonce of the message is being inserted at the right index.
		if merkleIndex+1 != (*message).Nonce() {
			return fmt.Errorf("nonce of message is not equal to the merkle index: %d != %d", (*message).Nonce(), merkleIndex+1)
		}

		e.chainExecutors[chainID].merkleTree.Insert(leaf[:])

		messageNonce := (*message).Nonce()
		attestationMask := execTypes.DBAttestation{
			ChainID:     &chainID,
			Destination: &destination,
			Nonce:       &messageNonce,
		}
		nonce, blockTime, err := e.executorDB.GetAttestationForNonceOrGreater(ctx, attestationMask)
		if err != nil {
			return fmt.Errorf("could not get attestation for nonce or greater: %w", err)
		}
		if nonce != nil && blockTime != nil {
			err = e.executorDB.StoreMessage(ctx, *message, log.BlockNumber, true, *blockTime+uint64((*message).OptimisticSeconds()))
			if err != nil {
				return fmt.Errorf("could not store message: %w", err)
			}
		} else {
			err = e.executorDB.StoreMessage(ctx, *message, log.BlockNumber, false, 0)
			if err != nil {
				return fmt.Errorf("could not store message: %w", err)
			}
		}
	case destinationContract:
		//nolint:exhaustive
		switch contractEvent.eventType {
		case attestationAcceptedEvent:
			// TODO: Change to populating the attestation table with the new attestation format.
			attestation, err := e.logToAttestation(log, chainID)
			if err != nil {
				return fmt.Errorf("could not convert log to attestation: %w", err)
			}

			if attestation == nil {
				return nil
			}

			logHeader, err := e.chainExecutors[(*attestation).Destination()].rpcClient.HeaderByNumber(ctx, big.NewInt(int64(log.BlockNumber)))
			if err != nil {
				return fmt.Errorf("could not get log header: %w", err)
			}

			err = e.executorDB.StoreAttestation(ctx, *attestation, log.BlockNumber, logHeader.Time)
			if err != nil {
				return fmt.Errorf("could not store attestation: %w", err)
			}
		case executedEvent:
			originDomain, messageLeaf, ok := e.chainExecutors[chainID].destinationParser.ParseExecuted(log)
			if !ok || originDomain == nil || messageLeaf == nil {
				return fmt.Errorf("could not parse executed event")
			}

			e.chainExecutors[chainID].executed[*messageLeaf] = true
		case otherEvent:
			logger.Warnf("the log's event type is not supported")
		default:
			return fmt.Errorf("log type not supported")
		}
	case summitContract:
		// TODO: Parse the snapshot and put everything into the state table.
	case other:
		logger.Warnf("the log's event type is not supported")
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
func (e Executor) executeExecutable(ctx context.Context, chainID uint32) error {
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(time.Duration(e.config.ExecuteInterval) * time.Second):
			page := 1
			currentTime := uint64(time.Now().Unix())

			messageMask := execTypes.DBMessage{
				ChainID: &chainID,
			}

			for {
				messages, err := e.executorDB.GetExecutableMessages(ctx, messageMask, currentTime, page)
				if err != nil {
					return fmt.Errorf("could not get executable messages: %w", err)
				}

				if len(messages) == 0 {
					break
				}

				for _, message := range messages {
					leaf, err := message.ToLeaf()
					if err != nil {
						return fmt.Errorf("could not convert message to leaf: %w", err)
					}

					destinationDomain := message.DestinationDomain()

					if !e.chainExecutors[destinationDomain].executed[leaf] {
						executed, err := e.Execute(ctx, message)
						if err != nil {
							logger.Errorf("could not execute message, retrying: %s", err)
							continue
						}

						if !executed {
							continue
						}
					}

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

				page++
			}
		}
	}
}

// setMinimumTime sets the minimum time for the message to be executed by checking for associated attestations.
//
//nolint:gocognit,cyclop
func (e Executor) setMinimumTime(ctx context.Context, chainID uint32, destinationDomain uint32) error {
	// TODO: Make for origin-dest, not just origin
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(time.Duration(e.config.SetMinimumTimeInterval) * time.Second):
			page := 1
			messageMask := execTypes.DBMessage{
				ChainID:     &chainID,
				Destination: &destinationDomain,
			}

			var unsetMessages []types.Message

			// Get all unset messages.
			for {
				messages, err := e.executorDB.GetUnsetMinimumTimeMessages(ctx, messageMask, page)
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

			page = 1
			minNonce := unsetMessages[0].Nonce()
			attestationMask := execTypes.DBAttestation{
				ChainID:     &chainID,
				Destination: &destinationDomain,
			}

			var attestations []execTypes.DBAttestation

			// Get all attestations for the unset messages.
			for {
				atts, err := e.executorDB.GetAttestationsAboveOrEqualNonce(ctx, attestationMask, minNonce, page)
				if err != nil {
					return fmt.Errorf("could not get attestations: %w", err)
				}

				if len(atts) == 0 {
					break
				}

				attestations = append(attestations, atts...)

				page++
			}

			err := e.setMinimumTimes(ctx, unsetMessages, attestations)
			if err != nil {
				return fmt.Errorf("could not set minimum times: %w", err)
			}
		}
	}
}
