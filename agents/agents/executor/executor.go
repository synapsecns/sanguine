package executor

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/prysmaticlabs/prysm/shared/trieutil"
	"github.com/synapsecns/sanguine/agents/agents/executor/config"
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
	// config is the executor agent config.
	config config.Config
	// scribeClient is the client to the Scribe gRPC server.
	scribeClient client.ScribeClient
	// lastLogs is a map from chain ID -> last log processed.
	lastLogs map[uint32]*logOrderInfo
	// closeConnection is a map from chain ID -> channel to close the connection.
	closeConnection map[uint32]chan bool
	// roots is a map from chain ID -> slice of merkle roots. The root at [i] is the root of nonce i.
	roots map[uint32]map[uint32][][32]byte
	// originParsers is a map from chain ID -> origin parser.
	originParsers map[uint32]origin.Parser
	// attestationCollectorParser is an attestationCollector parser.
	attestationCollectorParser attestationcollector.Parser
	// attestationCollectorChainID is the chain ID of the Synapse chain.
	attestationCollectorChainID uint32
	// attestationCollectorAddress is the address of the attestation collector contract.
	attestationCollectorAddress common.Address
	// logChans is a mapping from chain ID -> log channel.
	logChans map[uint32]chan *ethTypes.Log
	// merkleTrees is a map from chain ID -> destination domain -> merkle tree.
	merkleTrees map[uint32]map[uint32]*trieutil.SparseMerkleTrie
}

// logOrderInfo is a struct to keep track of the order of a log.
type logOrderInfo struct {
	blockNumber uint64
	blockIndex  uint
}

const treeDepth uint64 = 32

const logChanSize = 1000

// NewExecutor creates a new executor agent.
func NewExecutor(config config.Config, scribeClient client.ScribeClient) (*Executor, error) {
	lastLogs := make(map[uint32]*logOrderInfo)
	channels := make(map[uint32]chan *ethTypes.Log)
	closeConnectionChans := make(map[uint32]chan bool)
	roots := make(map[uint32]map[uint32][][32]byte)
	originParsers := make(map[uint32]origin.Parser)
	attestationCollectorParser, err := attestationcollector.NewParser(common.HexToAddress(config.AttestationCollectorAddress))
	if err != nil {
		return nil, fmt.Errorf("could not create attestationcollector parser: %w", err)
	}

	merkleTrees := make(map[uint32]map[uint32]*trieutil.SparseMerkleTrie)

	for _, chain := range config.Chains {
		lastLogs[chain.ChainID] = &logOrderInfo{
			blockNumber: 0,
			blockIndex:  0,
		}
		channels[chain.ChainID] = make(chan *ethTypes.Log, logChanSize)
		closeConnectionChans[chain.ChainID] = make(chan bool, 1)
		originParser, err := origin.NewParser(common.HexToAddress(chain.OriginAddress))
		if err != nil {
			return nil, fmt.Errorf("could not create origin parser: %w", err)
		}

		originParsers[chain.ChainID] = originParser

		roots[chain.ChainID] = make(map[uint32][][32]byte)
		merkleTrees[chain.ChainID] = make(map[uint32]*trieutil.SparseMerkleTrie)

		for _, destination := range config.Chains {
			if destination.ChainID == chain.ChainID {
				continue
			}
			roots[chain.ChainID][destination.ChainID] = [][32]byte{}

			merkleTree, err := trieutil.NewTrie(treeDepth)
			if err != nil {
				return nil, fmt.Errorf("could not create merkle tree: %w", err)
			}

			merkleTrees[chain.ChainID][destination.ChainID] = merkleTree
		}
	}

	return &Executor{
		config:                      config,
		scribeClient:                scribeClient,
		lastLogs:                    lastLogs,
		closeConnection:             closeConnectionChans,
		roots:                       roots,
		originParsers:               originParsers,
		attestationCollectorParser:  attestationCollectorParser,
		attestationCollectorChainID: config.AttestationCollectorChainID,
		attestationCollectorAddress: common.HexToAddress(config.AttestationCollectorAddress),
		logChans:                    channels,
		merkleTrees:                 merkleTrees,
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
}

// Listen listens to the log channel and processes the logs. Requires Start to be called first.
func (e Executor) Listen(ctx context.Context, chainID uint32) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case log := <-e.logChans[chainID]:
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
		address = e.attestationCollectorAddress.String()
	default:
		return fmt.Errorf("contract type not supported")
	}

	stream, err := grpcClient.StreamLogs(ctx, &pbscribe.StreamLogsRequest{
		Filter: &pbscribe.LogFilter{
			ContractAddress: &pbscribe.NullableString{Kind: &pbscribe.NullableString_Data{Data: address}},
			ChainId:         chain.ChainID,
		},
		FromBlock: "earliest",
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
			if !e.lastLogs[chain.ChainID].verifyAfter(*log) {
				return fmt.Errorf("log is not in chronological order. last log blockNumber: %d, blockIndex: %d. this log blockNumber: %d, blockIndex: %d, txHash: %s", e.lastLogs[chain.ChainID].blockNumber, e.lastLogs[chain.ChainID].blockIndex, log.BlockNumber, log.Index, log.TxHash.String())
			}

			e.logChans[chain.ChainID] <- log
			e.lastLogs[chain.ChainID].blockNumber = log.BlockNumber
			e.lastLogs[chain.ChainID].blockIndex = log.Index
		}
	}
}

// processLog processes the log and updates the merkle tree.
func (e Executor) processLog(log ethTypes.Log, chainID uint32) error {
	leafData, destination, err := e.logToLeaf(log, chainID)
	if err != nil {
		return fmt.Errorf("could not convert log to leaf: %w", err)
	}
	if leafData == nil {
		return nil
	}

	merkleIndex := e.merkleTrees[chainID][destination].NumOfItems()

	e.merkleTrees[chainID][destination].Insert(leafData, merkleIndex)
	e.roots[chainID][destination] = append(e.roots[chainID][destination], e.merkleTrees[chainID][destination].Root())

	return nil
}

// GetRoot returns the merkle root at the given index.
func (e Executor) GetRoot(index uint64, chainID uint32, destination uint32) ([32]byte, error) {
	if index >= uint64(len(e.roots[chainID][destination])) {
		return [32]byte{}, fmt.Errorf("index out of range")
	}

	return e.roots[chainID][destination][index], nil
}

// logToLeaf converts the log to a leaf data.
func (e Executor) logToLeaf(log ethTypes.Log, chainID uint32) ([]byte, uint32, error) {
	if eventType, ok := e.originParsers[chainID].EventType(log); ok && eventType == origin.DispatchEvent {
		committedMessage, ok := e.originParsers[chainID].ParseDispatch(log)
		if !ok {
			return nil, 0, fmt.Errorf("could not parse committed message")
		}

		message, err := types.DecodeMessage(committedMessage.Message())
		if err != nil {
			return nil, 0, fmt.Errorf("could not decode message: %w", err)
		}

		leaf, err := message.ToLeaf()
		if err != nil {
			return nil, 0, fmt.Errorf("could not convert message to leaf: %w", err)
		}

		return leaf[:], message.DestinationDomain(), nil
	}

	logger.Warnf("could not match the log's event type")

	return nil, 0, nil
}

func (l logOrderInfo) verifyAfter(log ethTypes.Log) bool {
	if log.BlockNumber < l.blockNumber {
		return false
	}

	if log.BlockNumber == l.blockNumber {
		// TODO: duplicates
		return log.Index > l.blockIndex
	}

	return true
}
