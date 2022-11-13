package executor

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hashicorp/consul/sdk/freeport"
	"github.com/synapsecns/sanguine/services/scribe/api"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/base"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"github.com/synapsecns/sanguine/services/scribe/node"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Executor is the executor agent.
type Executor struct {
	// Scribe is the Scribe service.
	Scribe *node.Scribe
	// scribeConfig is the config for the Scribe.
	scribeConfig config.Config
	// clients is a mapping of chain IDs -> clients.
	clients map[uint32][]backfill.ScribeBackend
	// dbPath is the path to the database.
	dbPath string
	// dbType is the type of database.
	dbType string
	// grpcPort is the port to expose the gRPC server on.
	grpcPort uint16
	// logChans is a map from chain ID -> log channel.
	logChans map[uint32]chan *types.Log
}

// NewExecutor creates a new executor agent.
func NewExecutor(ctx context.Context, scribeConfig config.Config, clients map[uint32][]backfill.ScribeBackend, dbPath string, dbType string, grpcPort uint16) (*Executor, error) {
	eventDB, err := api.InitDB(ctx, dbType, dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not initialize db: %w", err)
	}

	executorScribe, err := node.NewScribe(eventDB, clients, scribeConfig)
	if err != nil {
		return nil, fmt.Errorf("could not create scribe: %w", err)
	}

	return &Executor{
		Scribe:       executorScribe,
		scribeConfig: scribeConfig,
		clients:      clients,
		dbPath:       dbPath,
		dbType:       dbType,
		grpcPort:     grpcPort,
		logChans:     make(map[uint32]chan *types.Log),
	}, nil
}

// Start starts the executor agent. This runs the Scribe's live node, and uses gRPC to process the logs.
func (e Executor) Start(ctx context.Context) error {
	// Start the Scribe's live node concurrently.
	go func() {
		err := e.Scribe.Start(ctx)
		if err != nil {
			logger.Warnf("could not start scribe: %s", err)
			return
		}
	}()

	ports, err := freeport.Take(2)
	if len(ports) < 2 || err != nil {
		return fmt.Errorf("could not get free port: %w", err)
	}

	// Start the GraphQL server for the Scribe, and expose the gRPC server.
	go func() {
		err := api.Start(ctx, api.Config{
			HTTPPort: uint16(ports[0]),
			Database: e.dbType,
			Path:     e.dbPath,
			GRPCPort: uint16(ports[1]),
		})
		if err != nil {
			logger.Warnf("could not start api: %s", err)
			return
		}
	}()

	// Consume the logs from gRPC.
	conn, err := grpc.Dial(fmt.Sprintf(":%d", ports[1]), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("could not dial: %w", err)
	}

	client := pbscribe.NewScribeServiceClient(conn)
	g, groupCtx := errgroup.WithContext(ctx)

	// Process each chain in a separate goroutine.
	for _, chain := range e.scribeConfig.Chains {
		g.Go(func() error {
			var responseLogs []*pbscribe.Log

			page := uint32(1)

			for {
				response, err := client.FilterLogs(groupCtx, &pbscribe.FilterLogsRequest{
					Filter: &pbscribe.LogFilter{
						ContractAddress: nil,
						ChainId:         chain.ChainID,
						BlockNumber:     nil,
						TxHash:          nil,
						TxIndex:         nil,
						BlockHash:       nil,
						Index:           nil,
						Confirmed:       nil,
					},
					Page: page,
				})
				if err != nil {
					return fmt.Errorf("could not filter logs: %w", err)
				}

				responseLogs = append(responseLogs, response.Logs...)
				page++

				// See if we need to paginate.
				if len(response.Logs) < base.PageSize {
					break
				}
			}

			// Convert the logs to the types.Log type, and put them in the channel.
			for _, log := range responseLogs {
				e.logChans[chain.ChainID] <- log.ToLog()
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not process logs: %w", err)
	}

	return nil
}
