package executor

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/services/scribe/api"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
)

// Executor is the executor agent.
type Executor struct {
	// address is the address of the contract.
	address common.Address
	// chainID is the chain ID for the executor.
	chainID uint32
	// dbPath is the path to the database.
	dbPath string
	// dbType is the type of database.
	dbType string
	// lastLog is the last log that was processed.
	lastLog logOrderInfo
	// LogChan is a channel for logs.
	LogChan chan *types.Log
}

// logOrderInfo is a struct to keep track of the order of a log.
type logOrderInfo struct {
	blockNumber uint64
	blockIndex  uint
}

// NewExecutor creates a new executor agent.
func NewExecutor(address common.Address, chainID uint32, dbPath string, dbType string) (*Executor, error) {
	return &Executor{
		address: address,
		chainID: chainID,
		dbPath:  dbPath,
		dbType:  dbType,
		LogChan: make(chan *types.Log, 1000),
	}, nil
}

// Start starts the executor agent. This uses gRPC to process the logs.
func (e Executor) Start(ctx context.Context) error {
	// Start the GraphQL server for the Scribe, and expose the gRPC server.
	apiConfig := api.Config{
		HTTPPort: uint16(freeport.GetPort()),
		Database: e.dbType,
		Path:     e.dbPath,
		GRPCPort: uint16(freeport.GetPort()),
	}

	go func() {
		err := api.Start(ctx, apiConfig)
		if err != nil {
			logger.Warnf("could not start api: %s", err)

			return
		}
	}()

	// Consume the logs from gRPC.
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("localhost:%d", apiConfig.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("could not dial grpc: %w", err)
	}

	client := pbscribe.NewScribeServiceClient(conn)

	// Ensure that gRPC is up and running.
	healthCheck, err := client.Check(ctx, &pbscribe.HealthCheckRequest{}, grpc.WaitForReady(true))
	if err != nil {
		return fmt.Errorf("could not check: %w", err)
	}
	if healthCheck.Status != pbscribe.HealthCheckResponse_SERVING {
		return fmt.Errorf("not serving: %s", healthCheck.Status)
	}

	g, groupCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		stream, err := client.StreamLogs(groupCtx, &pbscribe.StreamLogsRequest{
			Filter: &pbscribe.LogFilter{
				ContractAddress: &pbscribe.NullableString{Kind: &pbscribe.NullableString_Data{Data: e.address.Hex()}},
				ChainId:         e.chainID,
			},
			FromBlock: "earliest",
			ToBlock:   "latest",
		})
		if err != nil {
			return fmt.Errorf("could not stream logs: %w", err)
		}

		for {
			response, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return fmt.Errorf("could not receive: %w", err)
			}

			log := response.Log.ToLog()
			if log == nil {
				return fmt.Errorf("could not convert log")
			}
			if !e.lastLog.verifyAfter(*log) {
				return fmt.Errorf("log is not in chronological order. last log blockNumber: %d, blockIndex: %d. this log blockNumber: %d, blockIndex: %d, txHash: %s", e.lastLog.blockNumber, e.lastLog.blockIndex, log.BlockNumber, log.Index, log.TxHash.String())
			}

			e.LogChan <- log
			e.lastLog = logOrderInfo{
				blockNumber: log.BlockNumber,
				blockIndex:  log.Index,
			}
		}
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not process logs: %w", err)
	}

	return nil
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
