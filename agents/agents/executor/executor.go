package executor

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/consul/sdk/freeport"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/scribe/api"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/node"
)

// Executor is the executor agent.
type Executor struct {
	// Scribe is the scribe service.
	Scribe *node.Scribe
	// dbPath is the path to the database.
	dbPath string
	// dbType is the type of database.
	dbType string
	// grpcPort is the port to expose the gRPC server on.
	grpcPort uint16
}

// NewExecutor creates a new executor agent.
func NewExecutor(ctx context.Context, configPath string, dbPath string, dbType string, grpcPort uint16) (*Executor, error) {
	var clients map[uint32][]backfill.ScribeBackend

	scribeConfig, err := config.DecodeConfig(core.ExpandOrReturnPath(configPath))
	if err != nil {
		return nil, fmt.Errorf("could not decode config: %w", err)
	}

	eventDB, err := api.InitDB(ctx, dbType, dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not initialize db: %w", err)
	}

	clients = make(map[uint32][]backfill.ScribeBackend)

	for _, client := range scribeConfig.Chains {
		backendClient, err := ethclient.DialContext(ctx, fmt.Sprintf("%s/1/rpc/%d", scribeConfig.RPCURL, client.ChainID))
		if err != nil {
			return nil, fmt.Errorf("could not start client for %s", fmt.Sprintf("%s/1/rpc/%d", scribeConfig.RPCURL, client.ChainID))
		}

		backendClientWConfirmations, err := ethclient.DialContext(ctx, fmt.Sprintf("%s/2/rpc/%d", scribeConfig.RPCURL, client.ChainID))
		if err != nil {
			return nil, fmt.Errorf("could not start client for %s", fmt.Sprintf("%s/4/rpc/%d", scribeConfig.RPCURL, client.ChainID))
		}

		clients[client.ChainID] = append(clients[client.ChainID], backendClient)
		clients[client.ChainID] = append(clients[client.ChainID], backendClientWConfirmations)
	}

	executorScribe, err := node.NewScribe(eventDB, clients, scribeConfig)
	if err != nil {
		return nil, fmt.Errorf("could not create scribe: %w", err)
	}

	return &Executor{
		Scribe:   executorScribe,
		dbPath:   dbPath,
		dbType:   dbType,
		grpcPort: grpcPort,
	}, nil
}

// Start starts the executor agent. This runs the Scribe's live node, and uses gRPC to process the logs.
func (e Executor) Start(ctx context.Context) error {
	// Start the Scribe's live node.
	err := e.Scribe.Start(ctx)
	if err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}

	// Start the GraphQL server for the Scribe, and expose the gRPC server.
	ports, err := freeport.Take(1)
	if len(ports) < 1 || err != nil {
		return fmt.Errorf("could not get free port: %w", err)
	}

	err = api.Start(ctx, api.Config{
		HTTPPort: uint16(ports[0]),
		Database: e.dbType,
		Path:     e.dbPath,
		GRPCPort: e.grpcPort,
	})
	if err != nil {
		return fmt.Errorf("could not start api: %w", err)
	}

	// Consume the logs from gRPC.

	return nil
}
