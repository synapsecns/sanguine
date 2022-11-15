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
	// grpcPort is the port to expose the gRPC server on.
	grpcPort uint16
	// lastBlockProcessed is a map from chain ID -> last block processed.
	lastBlockProcessed map[uint32]uint64
	// LogChans is a map from chain ID -> log channel.
	LogChans map[uint32]chan *types.Log
}

// NewExecutor creates a new executor agent.
func NewExecutor(chainIDs []uint32, dbPath string, dbType string, grpcPort uint16) (*Executor, error) {
	channels := make(map[uint32]chan *types.Log)
	for _, chain := range chainIDs {
		channels[chain] = make(chan *types.Log, 1000)
	}

	return &Executor{
		chainIDs: chainIDs,
		dbPath:   dbPath,
		dbType:   dbType,
		grpcPort: grpcPort,
		LogChans: channels,
	}, nil
}

// Start starts the executor agent. This uses gRPC to process the logs.
func (e Executor) Start(ctx context.Context) error {
	// Start the GraphQL server for the Scribe, and expose the gRPC server.
	go func() {
		err := api.Start(ctx, api.Config{
			HTTPPort: uint16(freeport.GetPort()),
			Database: e.dbType,
			Path:     e.dbPath,
			GRPCPort: e.grpcPort,
		})
		if err != nil {
			logger.Warnf("could not start api: %s", err)

			return
		}
	}()

	// Consume the logs from gRPC.
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("localhost:%d", e.grpcPort), opts...)
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
				e.LogChans[chain] <- responseLogs[i].ToLog()
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not process logs: %w", err)
	}

	return nil
}
