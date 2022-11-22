package executor

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	// LogChans is a mapping from chain ID -> log channel.
	LogChans map[uint32]chan *types.Log
	// closeConnection is a channel to close the connection.
	closeConnection chan bool
}

// logOrderInfo is a struct to keep track of the order of a log.
type logOrderInfo struct {
	blockNumber uint64
	blockIndex  uint
}

// NewExecutor creates a new executor agent.
func NewExecutor(chainIDs []uint32, addresses map[uint32]common.Address, scribeClient client.ScribeClient) (*Executor, error) {
	channels := make(map[uint32]chan *types.Log)
	for _, chainID := range chainIDs {
		channels[chainID] = make(chan *types.Log, 1000)
	}

	return &Executor{
		chainIDs:        chainIDs,
		addresses:       addresses,
		scribeClient:    scribeClient,
		lastLog:         make(map[uint32]logOrderInfo),
		LogChans:        channels,
		closeConnection: make(chan bool, 1),
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

	g, gCtx := errgroup.WithContext(ctx)

	for _, chainID := range e.chainIDs {
		chainID := chainID

		g.Go(func() error {
			stream, err := grpcClient.StreamLogs(gCtx, &pbscribe.StreamLogsRequest{
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
				case <-e.closeConnection:
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
func (e Executor) Stop() {
	e.closeConnection <- true
}

func (l logOrderInfo) verifyAfter(log types.Log) bool {
	if log.BlockNumber < l.blockNumber {
		return false
	}

	if log.BlockNumber == l.blockNumber {
		return log.Index > l.blockIndex
	}

	return true
}
