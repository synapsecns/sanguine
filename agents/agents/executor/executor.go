package executor

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/prysmaticlabs/prysm/shared/trieutil"
	"github.com/synapsecns/sanguine/agents/agents/executor/config"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	execTypes "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/services/scribe/client"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"math/big"
	"strconv"
)

// ChainExecutor is a struct that contains the necessary information for each chain level executor.
type ChainExecutor struct {
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
	// merkleTrees is a map from destination chain ID -> merkle tree.
	merkleTrees map[uint32]*trieutil.SparseMerkleTrie
	// client is an RPC client.
	client Backend
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
	// chainExecutors is a map from chain ID -> chain executor.
	chainExecutors map[uint32]*ChainExecutor
}

// logOrderInfo is a struct to keep track of the order of a log.
type logOrderInfo struct {
	blockNumber uint64
	blockIndex  uint
}

const treeDepth uint64 = 32

const logChanSize = 1000

// NewExecutor creates a new executor agent.
func NewExecutor(ctx context.Context, config config.Config, executorDB db.ExecutorDB, scribeClient client.ScribeClient, clients map[uint32]Backend) (*Executor, error) {
	chainExecutors := make(map[uint32]*ChainExecutor)
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", scribeClient.URL, scribeClient.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	for _, chain := range config.Chains {
		originParser, err := origin.NewParser(common.HexToAddress(chain.OriginAddress))
		if err != nil {
			return nil, fmt.Errorf("could not create origin parser: %w", err)
		}

		destinationParser, err := destination.NewParser(common.HexToAddress(chain.DestinationAddress))
		if err != nil {
			return nil, fmt.Errorf("could not create destination parser: %w", err)
		}

		chainExecutors[chain.ChainID] = &ChainExecutor{
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
			merkleTrees:       make(map[uint32]*trieutil.SparseMerkleTrie),
			client:            clients[chain.ChainID],
		}

		for _, destination := range config.Chains {
			if destination.ChainID == chain.ChainID {
				continue
			}

			tree, err := trieutil.NewTrie(treeDepth)
			if err != nil {
				return nil, fmt.Errorf("could not create merkle tree: %w", err)
			}

			chainExecutors[chain.ChainID].merkleTrees[destination.ChainID] = tree
		}
	}

	return &Executor{
		config:         config,
		executorDB:     executorDB,
		scribeClient:   scribeClient,
		grpcConn:       conn,
		grpcClient:     grpcClient,
		chainExecutors: chainExecutors,
	}, nil
}

// Start starts the executor agent. This uses gRPC to process the logs.
func (e Executor) Start(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	for _, chain := range e.config.Chains {
		chain := chain

		g.Go(func() error {
			return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chain, originContract)
		})

		g.Go(func() error {
			return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chain, destinationContract)
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error when streaming logs: %w", err)
	}

	return nil
}

// Stop stops the executor agent.
func (e Executor) Stop(chainID uint32) {
	e.chainExecutors[chainID].closeConnection <- true
	e.chainExecutors[chainID].stopListenChan <- true
}

// Listen listens to the log channel and processes the logs. Requires Start to be called first.
func (e Executor) Listen(ctx context.Context, chainID uint32) error {
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

// GetRoot returns the merkle root at the given nonce.
func (e Executor) GetRoot(ctx context.Context, nonce uint32, chainID uint32, destination uint32) ([32]byte, error) {
	if nonce == 0 || nonce > uint32(e.chainExecutors[chainID].merkleTrees[destination].NumOfItems()) {
		return [32]byte{}, fmt.Errorf("nonce is out of range")
	}

	messageMask := execTypes.DBMessage{
		ChainID:     &chainID,
		Destination: &destination,
		Nonce:       &nonce,
	}
	root, err := e.executorDB.GetRoot(ctx, messageMask)
	if err != nil {
		return [32]byte{}, fmt.Errorf("could not get message: %w", err)
	}

	return root, nil
}

// BuildTreeFromDB builds the merkle tree from the database's messages. This function will
// reset the current merkle tree and replace it with the one built from the database.
// This function should also not be called while Start or Listen are running.
func (e Executor) BuildTreeFromDB(ctx context.Context, chainID uint32, destination uint32) error {
	var allMessages []types.Message

	messageMask := execTypes.DBMessage{
		ChainID:     &chainID,
		Destination: &destination,
	}
	page := 1

	for {
		messages, err := e.executorDB.GetMessages(ctx, messageMask, page)
		if err != nil {
			return fmt.Errorf("could not get messages: %w", err)
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
			return fmt.Errorf("could not convert message to leaf: %w", err)
		}

		rawMessages[i] = rawMessage[:]
	}

	merkleTree, err := trieutil.GenerateTrieFromItems(rawMessages, treeDepth)
	if err != nil {
		return fmt.Errorf("could not generate trie from items: %w", err)
	}

	e.chainExecutors[chainID].merkleTrees[destination] = merkleTree

	return nil
}

type contractType int

const (
	originContract contractType = iota
	destinationContract
	other
)

// VerifyMessageNonce verifies a message against the merkle tree at the state of the given nonce.
func (e Executor) VerifyMessageNonce(ctx context.Context, nonce uint32, message types.Message, chainID uint32, destination uint32) (bool, error) {
	root, err := e.GetRoot(ctx, nonce, chainID, destination)
	if err != nil {
		return false, fmt.Errorf("could not get root: %w", err)
	}

	proof, err := e.GetLatestNonceProof(nonce, chainID, destination)
	if err != nil {
		return false, fmt.Errorf("could not get latest nonce proof: %w", err)
	}

	leaf, err := message.ToLeaf()
	if err != nil {
		return false, fmt.Errorf("could not convert message to leaf: %w", err)
	}

	inTree := trieutil.VerifyMerkleBranch(root[:], leaf[:], int(nonce-1), proof, treeDepth)

	return inTree, nil
}

// VerifyOptimisticPeriod verifies that the optimistic period is valid.
func (e Executor) VerifyOptimisticPeriod(ctx context.Context, message types.Message) (bool, error) {
	chainID := message.OriginDomain()
	destination := message.DestinationDomain()
	nonce := message.Nonce()
	attestationMask := execTypes.DBAttestation{
		ChainID:     &chainID,
		Destination: &destination,
		Nonce:       &nonce,
	}
	attestation, err := e.executorDB.GetAttestation(ctx, attestationMask)
	if err != nil {
		return false, fmt.Errorf("could not get attestation: %w", err)
	}

	if attestation == nil {
		return false, nil
	}

	root := (*attestation).Root()
	rootToHash := common.BytesToHash(root[:])
	attestationMask.Root = &rootToHash
<<<<<<< HEAD
	attestationBlockNumber, err := e.executorDB.GetAttestationBlockNumber(ctx, attestationMask)
	if err != nil {
		return false, fmt.Errorf("could not get attestation block number: %w", err)
	}

	if attestationBlockNumber == nil {
		return false, nil
	}

	header, err := e.chainExecutors[destination].client.HeaderByNumber(ctx, big.NewInt(int64(*attestationBlockNumber)))
	if err != nil {
		return false, fmt.Errorf("could not get header by number: %w", err)
	}

	attestationTimestamp := header.Time
=======
	attestationTime, err := e.executorDB.GetAttestationBlockTime(ctx, attestationMask)
	if err != nil {
		return false, fmt.Errorf("could not get attestation block time: %w", err)
	}

	if attestationTime == nil {
		return false, nil
	}

>>>>>>> master
	latestHeader, err := e.chainExecutors[destination].client.HeaderByNumber(ctx, nil)
	if err != nil {
		return false, fmt.Errorf("could not get latest header: %w", err)
	}

	currentTime := latestHeader.Time
<<<<<<< HEAD
	if attestationTimestamp+uint64(message.OptimisticSeconds()) > currentTime {
=======
	if *attestationTime+uint64(message.OptimisticSeconds()) > currentTime {
>>>>>>> master
		return false, nil
	}

	return true, nil
}

// GetLatestNonceProof returns the merkle proof for a nonce, with a tree where that nonce is the last item added.
// This is done by copying the current merkle tree's items and generating a new tree with the items from the range
// [0, nonce).
func (e Executor) GetLatestNonceProof(nonce, chainID, destination uint32) ([][]byte, error) {
	if nonce == 0 || nonce > uint32(e.chainExecutors[chainID].merkleTrees[destination].NumOfItems()) {
		return nil, fmt.Errorf("nonce is out of range")
	}

	items := e.chainExecutors[chainID].merkleTrees[destination].Items()
	tree, err := trieutil.GenerateTrieFromItems(items[:nonce], treeDepth)
	if err != nil {
		return nil, fmt.Errorf("could not generate trie: %w", err)
	}

	proof, err := tree.MerkleProof(int(nonce - 1))
	if err != nil {
		return nil, fmt.Errorf("could not get merkle proof: %w", err)
	}

	return proof, nil
}

// streamLogs uses gRPC to stream logs into a channel.
//
//nolint:cyclop
func (e Executor) streamLogs(ctx context.Context, grpcClient pbscribe.ScribeServiceClient, conn *grpc.ClientConn, chain config.ChainConfig, contract contractType) error {
	var address string

	//nolint:exhaustive
	switch contract {
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
	stream, err := grpcClient.StreamLogs(ctx, &pbscribe.StreamLogsRequest{
		Filter: &pbscribe.LogFilter{
			ContractAddress: &pbscribe.NullableString{Kind: &pbscribe.NullableString_Data{Data: address}},
			ChainId:         chain.ChainID,
		},
		FromBlock: fromBlock,
		ToBlock:   "latest",
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
//nolint:cyclop
func (e Executor) processLog(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	logType := e.logType(log, chainID)

	switch logType {
	case originContract:
		message, err := e.logToMessage(log, chainID)
		if err != nil {
			return fmt.Errorf("could not convert log to leaf: %w", err)
		}

		if message == nil {
			return nil
		}

		destination := (*message).DestinationDomain()
		merkleIndex := e.chainExecutors[chainID].merkleTrees[destination].NumOfItems()
		leaf, err := (*message).ToLeaf()
		if err != nil {
			return fmt.Errorf("could not convert message to leaf: %w", err)
		}

		// Make sure the nonce of the message is being inserted at the right index.
		if uint32(merkleIndex)+1 != (*message).Nonce() {
			return fmt.Errorf("nonce of message is not equal to the merkle index: %d != %d", (*message).Nonce(), merkleIndex+1)
		}

		e.chainExecutors[chainID].merkleTrees[destination].Insert(leaf[:], merkleIndex)
		root := e.chainExecutors[chainID].merkleTrees[destination].Root()
		err = e.executorDB.StoreMessage(ctx, *message, root, log.BlockNumber)
		if err != nil {
			return fmt.Errorf("could not store message: %w", err)
		}
	case destinationContract:
		attestation, err := e.logToAttestation(log, chainID)
		if err != nil {
			return fmt.Errorf("could not convert log to attestation: %w", err)
		}

		if attestation == nil {
			return nil
		}

<<<<<<< HEAD
		err = e.executorDB.StoreAttestation(ctx, *attestation, log.BlockNumber)
=======
		logHeader, err := e.chainExecutors[(*attestation).Destination()].client.HeaderByNumber(ctx, big.NewInt(int64(log.BlockNumber)))
		if err != nil {
			return fmt.Errorf("could not get log header: %w", err)
		}

		err = e.executorDB.StoreAttestation(ctx, *attestation, log.BlockNumber, logHeader.Time)
>>>>>>> master
		if err != nil {
			return fmt.Errorf("could not store attestation: %w", err)
		}
	case other:
		logger.Warnf("the log's event type is not supported")
	default:
		return fmt.Errorf("log type not supported")
	}

	return nil
}

// logToMessage converts the log to a leaf data.
func (e Executor) logToMessage(log ethTypes.Log, chainID uint32) (*types.Message, error) {
	committedMessage, ok := e.chainExecutors[chainID].originParser.ParseDispatch(log)
	if !ok {
		return nil, fmt.Errorf("could not parse committed message")
	}

	message, err := types.DecodeMessage(committedMessage.Message())
	if err != nil {
		return nil, fmt.Errorf("could not decode message: %w", err)
	}

	return &message, nil
}

// logToAttestation converts the log to an attestation.
func (e Executor) logToAttestation(log ethTypes.Log, chainID uint32) (*types.Attestation, error) {
	attestation, ok := e.chainExecutors[chainID].destinationParser.ParseAttestationAccepted(log)
	if !ok {
		return nil, fmt.Errorf("could not parse attestation")
	}

	return &attestation, nil
}

// logType determines whether a log is a `Dispatch` from Origin.sol or `AttestationAccepted` from Destination.sol.
func (e Executor) logType(log ethTypes.Log, chainID uint32) contractType {
	contract := other

	if eventType, ok := e.chainExecutors[chainID].originParser.EventType(log); ok && eventType == origin.DispatchEvent {
		contract = originContract
	}

	if eventType, ok := e.chainExecutors[chainID].destinationParser.EventType(log); ok && eventType == destination.AttestationAcceptedEvent {
		contract = destinationContract
	}

	return contract
}

func (l logOrderInfo) verifyAfter(log ethTypes.Log) bool {
	if log.BlockNumber < l.blockNumber {
		return false
	}

	if log.BlockNumber == l.blockNumber {
		return log.Index > l.blockIndex
	}

	return true
}
