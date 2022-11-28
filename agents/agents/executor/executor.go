package executor

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/prysmaticlabs/prysm/shared/trieutil"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/services/scribe/client"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
)

// Executor is the executor agent.
type Executor struct {
	// chainIDs are the chain IDs for the Executor to process.
	chainIDs []uint32
	// addresses is a map from chain ID -> address.
	addresses map[uint32]common.Address
	// scribeClient is the client to the Scribe gRPC server.
	scribeClient client.ScribeClient
	// lastLog is a map from chainID -> last log processed.
	lastLog map[uint32]logOrderInfo
	// closeConnection is a map from chain ID -> channel to close the connection.
	closeConnection map[uint32]chan bool
	// roots is a slice of merkle roots. The root at [i] is the root of nonce i.
	roots [][32]byte
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
func NewExecutor(chainIDs []uint32, addresses map[uint32]common.Address, scribeClient client.ScribeClient) (*Executor, error) {
	channels := make(map[uint32]chan *ethTypes.Log)
	closeChans := make(map[uint32]chan bool)
	originParsers := make(map[uint32]origin.Parser)
	attestationcollectorParsers := make(map[uint32]attestationcollector.Parser)

	for _, chainID := range chainIDs {
		channels[chainID] = make(chan *ethTypes.Log, 1000)
		closeChans[chainID] = make(chan bool, 1)
		originParser, err := origin.NewParser(addresses[chainID])
		if err != nil {
			return nil, fmt.Errorf("could not create origin parser: %w", err)
		}

		originParsers[chainID] = originParser
		attestationcollectorParser, err := attestationcollector.NewParser(addresses[chainID])
		if err != nil {
			return nil, fmt.Errorf("could not create attestationcollector parser: %w", err)
		}

		attestationcollectorParsers[chainID] = attestationcollectorParser

	}

	merkleTree, err := trieutil.NewTrie(treeDepth)
	if err != nil {
		return nil, fmt.Errorf("could not create merkle tree: %w", err)
	}

	return &Executor{
		chainIDs:                    chainIDs,
		addresses:                   addresses,
		scribeClient:                scribeClient,
		lastLog:                     make(map[uint32]logOrderInfo),
		closeConnection:             closeChans,
		roots:                       [][32]byte{},
		originParsers:               originParsers,
		attestationcollectorParsers: attestationcollectorParsers,
		LogChans:                    channels,
		MerkleTree:                  merkleTree,
	}, nil
}

// Start starts the executor agent. This uses gRPC to process the logs.
//
//nolint:cyclop,gocognit
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

	for _, chainID := range e.chainIDs {
		chainID := chainID

		g.Go(func() error {
			stream, err := grpcClient.StreamLogs(ctx, &pbscribe.StreamLogsRequest{
				Filter: &pbscribe.LogFilter{
					ContractAddress: &pbscribe.NullableString{Kind: &pbscribe.NullableString_Data{Data: e.addresses[chainID].Hex()}},
					ChainId:         chainID,
				},
				FromBlock: "earliest",
				ToBlock:   "latest",
			})
			if err != nil {
				return fmt.Errorf("could not stream logs: %w", err)
			}

			for {
				select {
				case <-e.closeConnection[chainID]:
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
					if !e.lastLog[chainID].verifyAfter(*log) {
						return fmt.Errorf("log is not in chronological order. last log blockNumber: %d, blockIndex: %d. this log blockNumber: %d, blockIndex: %d, txHash: %s", e.lastLog[chainID].blockNumber, e.lastLog[chainID].blockIndex, log.BlockNumber, log.Index, log.TxHash.String())
					}

					e.LogChans[chainID] <- log
					e.lastLog[chainID] = logOrderInfo{
						blockNumber: log.BlockNumber,
						blockIndex:  log.Index,
					}
				}
			}
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
}

// Listen listens to the log channel and processes the logs.
func (e Executor) Listen(ctx context.Context, chainID uint32) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case log := <-e.LogChans[chainID]:
			if log == nil {
				return fmt.Errorf("log is nil")
			}

			err := e.processLog(*log, chainID)
			if err != nil {
				return fmt.Errorf("could not process log: %w", err)
			}
		}
	}
}

// processLog processes the log and updates the merkle tree.
func (e Executor) processLog(log ethTypes.Log, chainID uint32) error {
	merkleIndex := e.MerkleTree.NumOfItems()
	leafData, err := e.logToLeaf(log, chainID)
	if err != nil {
		return fmt.Errorf("could not convert log to leaf: %w", err)
	}

	e.MerkleTree.Insert(leafData, merkleIndex)
	e.roots = append(e.roots, e.MerkleTree.Root())

	return nil
}

// getRoot returns the merkle root at the given index.
func (e Executor) getRoot(index uint64) ([32]byte, error) {
	if index >= uint64(len(e.roots)) {
		return [32]byte{}, fmt.Errorf("index out of range")
	}

	return e.roots[index], nil
}

// logToLeaf converts the log to a leaf data.
func (e Executor) logToLeaf(log ethTypes.Log, chainID uint32) ([]byte, error) {
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

		return leaf[:], nil
	} else if eventType, ok := e.attestationcollectorParsers[chainID].EventType(log); ok && eventType == 0 {
		// TODO: handle this case with attestationcollector properly.
		return nil, nil
	} else {
		return nil, fmt.Errorf("could not parse committed message")
	}

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
