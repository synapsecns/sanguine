package executor

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/services/scribe/api"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/base"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Executor is the executor agent.
type Executor struct {
	// chainIDs is the chain IDs that the executor is processing.
	chainIDs []uint32
	// dbPath is the path to the database.
	dbPath string
	// dbType is the type of database.
	dbType string
	// lastLog is a map from chain ID -> logOrder to make sure logs are chronological.
	lastLog map[uint32]logOrder
	// LogChans is a map from chain ID -> log channel.
	LogChans map[uint32]chan *types.Log
}

// logOrder is a struct to keep track of the order of a log.
type logOrder struct {
	blockNumber uint64
	blockIndex  uint
}

// NewExecutor creates a new executor agent.
func NewExecutor(chainIDs []uint32, dbPath string, dbType string) (*Executor, error) {
	lastLog := make(map[uint32]logOrder)
	channels := make(map[uint32]chan *types.Log)

	for _, chain := range chainIDs {
		channels[chain] = make(chan *types.Log, 1000)
	}

	return &Executor{
		chainIDs: chainIDs,
		dbPath:   dbPath,
		dbType:   dbType,
		lastLog:  lastLog,
		LogChans: channels,
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
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("localhost:%d", apiConfig.GRPCPort), opts...)
	if err != nil {
		return fmt.Errorf("could not dial grpc: %w", err)
	}

	client := pbscribe.NewScribeServiceClient(conn)

	healthCheck, err := client.Check(ctx, &pbscribe.HealthCheckRequest{}, grpc.WaitForReady(true))
	if err != nil {
		return fmt.Errorf("could not check: %w", err)
	}
	if healthCheck.Status != pbscribe.HealthCheckResponse_SERVING {
		return fmt.Errorf("not serving: %s", healthCheck.Status)
	}

	g, groupCtx := errgroup.WithContext(ctx)

	// Process each chain in a separate goroutine.
	for _, chain := range e.chainIDs {
		chain := chain

		g.Go(func() error {
			var responseLogs []*pbscribe.Log

			page := uint32(1)

			for {
				response, err := client.FilterLogs(groupCtx, &pbscribe.FilterLogsRequest{
					Filter: &pbscribe.LogFilter{
						ChainId: chain,
					},
					Page: page,
				})
				if err != nil {
					return fmt.Errorf("could not filter logs: %w", err)
				}

				responseLogs = append(responseLogs, response.Logs...)

				// See if we do not need to get the next page.
				if len(response.Logs) < base.PageSize {
					break
				}

				page++
			}

			// Convert the logs to the types.Log type, and put them in the channel.
			for i := len(responseLogs) - 1; i >= 0; i-- {
				log := responseLogs[i].ToLog()
				if log == nil {
					return fmt.Errorf("could not convert log")
				}
				if !e.lastLog[chain].verifyAfter(*log) {
					return fmt.Errorf("log is not in chronological order. last log blockNumber: %d, blockIndex: %d. this log blockNumber: %d, blockIndex: %d, txHash: %s", e.lastLog[chain].blockNumber, e.lastLog[chain].blockIndex, log.BlockNumber, log.Index, log.TxHash.String())
				}

				e.LogChans[chain] <- log
				e.lastLog[chain] = logOrder{
					blockNumber: log.BlockNumber,
					blockIndex:  log.Index,
				}
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not process logs: %w", err)
	}

	return nil
}

func (l logOrder) verifyAfter(log types.Log) bool {
	if log.BlockNumber < l.blockNumber {
		return false
	}

	if log.BlockNumber == l.blockNumber {
		return log.Index > l.blockIndex
	}

	return true
}
