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
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/services/scribe/client"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"strconv"
	"sync"
)

// Executor is the executor agent.
type Executor struct {
	// config is the executor agent config.
	config config.Config
	// executorDB is the executor agent database.
	executorDB db.ExecutorDB
	// scribeClient is the client to the Scribe gRPC server.
	scribeClient client.ScribeClient
	// lastLog is a map from chainID -> last log processed.
	lastLog map[uint32]logOrderInfo
	// lastLogMutex is a mutex for the lastLog map.
	lastLogMutex *sync.Mutex
	// closeConnection is a map from chain ID -> channel to close the connection.
	closeConnection map[uint32]chan bool
	// stopListenChan is a map from chain ID -> channel to stop listening for logs.
	stopListenChan map[uint32]chan bool
	// originParsers is a map from chain ID -> origin parser.
	originParsers map[uint32]origin.Parser
	// attestationcollectorParsers is a map from chain ID -> attestationcollector parser.
	attestationcollectorParsers map[uint32]attestationcollector.Parser
	// LogChans is a mapping from chain ID -> log channel.
	LogChans map[uint32]chan *ethTypes.Log
	// MerkleTree is the merkle tree.
	MerkleTree *trieutil.SparseMerkleTrie
}

// logOrderInfo is a struct to keep track of the order of a log.
type logOrderInfo struct {
	blockNumber uint64
	blockIndex  uint
}

const treeDepth uint64 = 32

// NewExecutor creates a new executor agent.
func NewExecutor(config config.Config, executorDB db.ExecutorDB, scribeClient client.ScribeClient) (*Executor, error) {
	channels := make(map[uint32]chan *ethTypes.Log)
	closeChans := make(map[uint32]chan bool)
	stopListenChans := make(map[uint32]chan bool)
	originParsers := make(map[uint32]origin.Parser)
	attestationcollectorParsers := make(map[uint32]attestationcollector.Parser)

	for _, chain := range config.Chains {
		channels[chain.ChainID] = make(chan *ethTypes.Log, 1000)
		closeChans[chain.ChainID] = make(chan bool, 1)
		stopListenChans[chain.ChainID] = make(chan bool, 1)
		originParser, err := origin.NewParser(common.HexToAddress(chain.OriginAddress))
		if err != nil {
			return nil, fmt.Errorf("could not create origin parser: %w", err)
		}

		originParsers[chain.ChainID] = originParser
		attestationcollectorParser, err := attestationcollector.NewParser(common.HexToAddress(chain.AttestationCollectorAddress))
		if err != nil {
			return nil, fmt.Errorf("could not create attestationcollector parser: %w", err)
		}

		attestationcollectorParsers[chain.ChainID] = attestationcollectorParser
	}

	merkleTree, err := trieutil.NewTrie(treeDepth)
	if err != nil {
		return nil, fmt.Errorf("could not create merkle tree: %w", err)
	}

	return &Executor{
		config:                      config,
		executorDB:                  executorDB,
		scribeClient:                scribeClient,
		lastLog:                     make(map[uint32]logOrderInfo),
		lastLogMutex:                &sync.Mutex{},
		closeConnection:             closeChans,
		stopListenChan:              stopListenChans,
		originParsers:               originParsers,
		attestationcollectorParsers: attestationcollectorParsers,
		LogChans:                    channels,
		MerkleTree:                  merkleTree,
	}, nil
}

// Start starts the executor agent. This uses gRPC to process the logs.
func (e Executor) Start(ctx context.Context) error {
	// Consume the logs from gRPC.
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", e.scribeClient.URL, e.scribeClient.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("could not dial grpc: %w", err)
	}

	grpcClient := pbscribe.NewScribeServiceClient(conn)

	// Ensure that gRPC is up and running.
	healthCheck, err := grpcClient.Check(ctx, &pbscribe.HealthCheckRequest{}, grpc.WaitForReady(true))
	if err != nil {
		return fmt.Errorf("could not check: %w", err)
	}
	if healthCheck.Status != pbscribe.HealthCheckResponse_SERVING {
		return fmt.Errorf("not serving: %s", healthCheck.Status)
	}

	g, _ := errgroup.WithContext(ctx)

	for _, chain := range e.config.Chains {
		chain := chain

		g.Go(func() error {
			return e.streamLogs(ctx, grpcClient, conn, chain, originContract)
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error when streaming logs: %w", err)
	}

	return nil
}

// Stop stops the executor agent.
func (e Executor) Stop(chainID uint32) {
	e.closeConnection[chainID] <- true
	e.stopListenChan[chainID] <- true
}

// Listen listens to the log channel and processes the logs. Requires Start to be called first.
func (e Executor) Listen(ctx context.Context, chainID uint32) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-e.stopListenChan[chainID]:
			return nil
		case log := <-e.LogChans[chainID]:
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
func (e Executor) GetRoot(ctx context.Context, nonce uint32, chainID uint32) (*[32]byte, error) {
	if nonce == 0 || nonce > uint32(e.MerkleTree.NumOfItems()) {
		return nil, fmt.Errorf("nonce is out of range")
	}

	messageMask := execTypes.DBMessage{
		ChainID: &chainID,
		Nonce:   &nonce,
	}
	message, err := e.executorDB.GetMessage(ctx, messageMask)
	if err != nil {
		return nil, fmt.Errorf("could not get message: %w", err)
	}
	if message == nil {
		return nil, fmt.Errorf("no message found for chainID %d and nonce %d", chainID, nonce)
	}

	return (*[32]byte)(message.Root.Bytes()), nil
}

// BuildTreeFromDB builds the merkle tree from the database's messages. This function will
// reset the current merkle tree and replace it with the one built from the database.
// This function should also not be called while Start or Listen are running.
func (e *Executor) BuildTreeFromDB(ctx context.Context, chainID uint32) error {
	merkleTree, err := trieutil.NewTrie(treeDepth)
	if err != nil {
		return fmt.Errorf("could not create merkle tree: %w", err)
	}

	nonce := uint32(1)

	for {
		messageMask := execTypes.DBMessage{
			ChainID: &chainID,
			Nonce:   &nonce,
		}
		message, err := e.executorDB.GetMessage(ctx, messageMask)
		if err != nil {
			return fmt.Errorf("could not get message: %w", err)
		}

		if message == nil {
			break
		}

		merkleTree.Insert(message.Root.Bytes(), int(nonce-1))

		nonce++
	}

	e.MerkleTree = merkleTree

	return nil
}

type contractType int

const (
	originContract contractType = iota
	attestationcollectorContract
)

// streamLogs uses gRPC to stream logs into a channel.
//
//nolint:cyclop
func (e Executor) streamLogs(ctx context.Context, grpcClient pbscribe.ScribeServiceClient, conn *grpc.ClientConn, chain config.ChainConfig, contract contractType) error {
	var address string
	switch contract {
	case originContract:
		address = chain.OriginAddress
	case attestationcollectorContract:
		address = chain.AttestationCollectorAddress
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
		case <-e.closeConnection[chain.ChainID]:
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
			if !e.lastLog[chain.ChainID].verifyAfter(*log) {
				return fmt.Errorf("log is not in chronological order. last log blockNumber: %d, blockIndex: %d. this log blockNumber: %d, blockIndex: %d, txHash: %s", e.lastLog[chain.ChainID].blockNumber, e.lastLog[chain.ChainID].blockIndex, log.BlockNumber, log.Index, log.TxHash.String())
			}

			e.LogChans[chain.ChainID] <- log
			e.lastLogMutex.Lock()
			e.lastLog[chain.ChainID] = logOrderInfo{
				blockNumber: log.BlockNumber,
				blockIndex:  log.Index,
			}
			e.lastLogMutex.Unlock()
		}
	}
}

// processLog processes the log and updates the merkle tree.
func (e Executor) processLog(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	merkleIndex := e.MerkleTree.NumOfItems()
	message, err := e.logToMessage(log, chainID)
	if err != nil {
		return fmt.Errorf("could not convert log to leaf: %w", err)
	}
	if message == nil {
		return nil
	}

	e.MerkleTree.Insert(message.Leaf.Bytes(), merkleIndex)
	root := e.MerkleTree.Root()
	rootHash := common.BytesToHash(root[:])
	message.Root = &rootHash
	err = e.executorDB.StoreMessage(ctx, *message)
	if err != nil {
		return fmt.Errorf("could not store message: %w", err)
	}

	return nil
}

// logToMessage converts the log to a leaf data.
func (e Executor) logToMessage(log ethTypes.Log, chainID uint32) (*execTypes.DBMessage, error) {
	var dbMessage *execTypes.DBMessage
	if eventType, ok := e.originParsers[chainID].EventType(log); ok && eventType == origin.DispatchEvent {
		committedMessage, ok := e.originParsers[chainID].ParseDispatch(log)
		if !ok {
			return nil, fmt.Errorf("could not parse committed message")
		}

		message, err := types.DecodeMessage(committedMessage.Message())
		if err != nil {
			return nil, fmt.Errorf("could not decode message: %w", err)
		}

		leaf, err := message.ToLeaf()
		if err != nil {
			return nil, fmt.Errorf("could not convert message to leaf: %w", err)
		}

		messageChainID := chainID
		messageNonce := message.Nonce()
		messageMessage := message.Body()
		messageLeaf := common.BytesToHash(leaf[:])
		messageBlockNumber := log.BlockNumber

		dbMessage = &execTypes.DBMessage{
			ChainID:     &messageChainID,
			Nonce:       &messageNonce,
			Root:        nil,
			Message:     &messageMessage,
			Leaf:        &messageLeaf,
			BlockNumber: &messageBlockNumber,
		}

		return dbMessage, nil
	} else if eventType, ok := e.attestationcollectorParsers[chainID].EventType(log); ok && eventType == 0 {
		// TODO: handle this case with attestationcollector properly.
		//nolint:nilnil
		return nil, nil
	}

	logger.Warnf("could not match the log's event type")

	//nolint:nilnil
	return nil, nil
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
