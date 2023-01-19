package cmd

import (
	"context"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/jftuga/termsize"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/mysql"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/sqlite"
	scribeAPI "github.com/synapsecns/sanguine/services/scribe/api"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/client"
	scribeCmd "github.com/synapsecns/sanguine/services/scribe/cmd"
	"github.com/synapsecns/sanguine/services/scribe/node"
	"golang.org/x/sync/errgroup"

	// used to embed markdown.
	_ "embed"
	"fmt"
	"github.com/synapsecns/sanguine/agents/agents/executor/config"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/core"
	"github.com/urfave/cli/v2"
	"os"
)

//go:embed cmd.md
var help string

// infoCommand gets info about using the executor agent.
var infoCommand = &cli.Command{
	Name:        "info",
	Description: "learn how to use executor cli",
	Action: func(c *cli.Context) error {
		fmt.Println(string(markdown.Render(help, termsize.Width(), 6)))
		return nil
	},
}

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Users/synapsecns/config.yaml",
	TakesFile: true,
	Required:  true,
}

var dbFlag = &cli.StringFlag{
	Name:     "db",
	Usage:    "--db <sqlite> or <mysql>",
	Value:    "sqlite",
	Required: true,
}

var pathFlag = &cli.StringFlag{
	Name:     "path",
	Usage:    "--path <path/to/database> or <database url>",
	Value:    "",
	Required: true,
}

var scribeTypeFlag = &cli.StringFlag{
	Name:     "scribe-type",
	Usage:    "--scribe-type <embedded> or <remote>",
	Required: true,
}

var scribeDBFlag = &cli.StringFlag{
	Name:  "scribe-db",
	Usage: "--scribe-db <sqlite> or <mysql>",
}

var scribePathFlag = &cli.StringFlag{
	Name:  "scribe-path",
	Usage: "--scribe-path <path/to/database> or <database url>",
}

var scribePortFlag = &cli.UintFlag{
	Name:  "scribe-port",
	Usage: "--scribe-port <port>",
	Value: 0,
}

var scribeGrpcPortFlag = &cli.UintFlag{
	Name:  "scribe-grpc-port",
	Usage: "--scribe-grpc-port <port>",
}

var scribeURL = &cli.StringFlag{
	Name:  "scribe-url",
	Usage: "--scribe-url <url>",
}

func createExecutorParameters(c *cli.Context) (executorConfig config.Config, executorDB db.ExecutorDB, clients map[uint32]executor.Backend, err error) {
	executorConfig, err = config.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
	if err != nil {
		return executorConfig, nil, nil, fmt.Errorf("failed to decode config: %w", err)
	}

	executorDB, err = InitExecutorDB(c.Context, c.String(dbFlag.Name), c.String(pathFlag.Name))
	if err != nil {
		return executorConfig, nil, nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	clients = make(map[uint32]executor.Backend)
	for _, execClient := range executorConfig.Chains {
		rpcDial, err := rpc.DialContext(c.Context, fmt.Sprintf("%s/confirmations/%d/rpc/%d", executorConfig.BaseOmnirpcURL, 1, execClient.ChainID))
		if err != nil {
			return executorConfig, nil, nil, fmt.Errorf("failed to dial rpc: %w", err)
		}

		ethClient := ethclient.NewClient(rpcDial)
		clients[execClient.ChainID] = ethClient
	}

	return executorConfig, executorDB, clients, nil
}

var runCommand = &cli.Command{
	Name:        "run",
	Description: "runs the executor service",
	Flags: []cli.Flag{configFlag, dbFlag, pathFlag, scribeTypeFlag,
		// The flags below are used when `scribeTypeFlag` is set to "embedded".
		scribeDBFlag, scribePathFlag,
		// The flags below are used when `scribeTypeFlag` is set to "remote".
		scribePortFlag, scribeGrpcPortFlag, scribeURL},
	Action: func(c *cli.Context) error {
		executorConfig, executorDB, clients, err := createExecutorParameters(c)
		if err != nil {
			return err
		}

		var scribeClient client.ScribeClient

		g, _ := errgroup.WithContext(c.Context)

		switch c.String(scribeTypeFlag.Name) {
		case "embedded":
			eventDB, err := scribeAPI.InitDB(c.Context, c.String(scribeDBFlag.Name), c.String(scribePathFlag.Name))
			if err != nil {
				return fmt.Errorf("failed to initialize database: %w", err)
			}

			scribeClients := make(map[uint32][]backfill.ScribeBackend)

			for _, client := range executorConfig.EmbeddedScribeConfig.Chains {
				for confNum := 1; confNum <= scribeCmd.MaxConfirmations; confNum++ {
					backendClient, err := backfill.DialBackend(c.Context, fmt.Sprintf("%s/confirmations/%d/rpc/%d", executorConfig.EmbeddedScribeConfig.RPCURL, confNum, client.ChainID))
					if err != nil {
						return fmt.Errorf("could not start client for %s", fmt.Sprintf("%s/%s/%d/rpc/%d", executorConfig.EmbeddedScribeConfig.RPCURL, "confirmations", confNum, client.ChainID))
					}

					scribeClients[client.ChainID] = append(scribeClients[client.ChainID], backendClient)
				}
			}

			scribe, err := node.NewScribe(eventDB, scribeClients, executorConfig.EmbeddedScribeConfig)
			if err != nil {
				return fmt.Errorf("failed to initialize scribe: %w", err)
			}

			g.Go(func() error {
				err := scribe.Start(c.Context)
				if err != nil {
					return fmt.Errorf("failed to start scribe: %w", err)
				}

				return nil
			})

			embedded := client.NewEmbeddedScribe(c.String(scribeDBFlag.Name), c.String(scribePathFlag.Name))

			g.Go(func() error {
				err := embedded.Start(c.Context)
				if err != nil {
					return fmt.Errorf("failed to start embedded scribe: %w", err)
				}

				return nil
			})

			scribeClient = embedded.ScribeClient
		case "remote":
			scribeClient = client.NewRemoteScribe(uint16(c.Uint(scribePortFlag.Name)), uint16(c.Uint(scribeGrpcPortFlag.Name)), c.String(scribeURL.Name)).ScribeClient
		default:
			return fmt.Errorf("invalid scribe type")
		}

		executor, err := executor.NewExecutor(c.Context, executorConfig, executorDB, scribeClient, clients)
		if err != nil {
			return fmt.Errorf("failed to create executor: %w", err)
		}

		g.Go(func() error {
			err = executor.Run(c.Context)
			if err != nil {
				return fmt.Errorf("failed to run executor: %w", err)
			}

			return nil
		})

		if err := g.Wait(); err != nil {
			return fmt.Errorf("failed to run executor: %w", err)
		}

		return nil
	},
}

// InitExecutorDB initializes a database given a database type and path.
func InitExecutorDB(ctx context.Context, database string, path string) (db.ExecutorDB, error) {
	switch {
	case database == "sqlite":
		sqliteStore, err := sqlite.NewSqliteStore(ctx, path)
		if err != nil {
			return nil, fmt.Errorf("failed to create sqlite store: %w", err)
		}

		return sqliteStore, nil

	case database == "mysql":
		if os.Getenv("OVERRIDE_MYSQL") != "" {
			dbname := os.Getenv("MYSQL_DATABASE")
			connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", core.GetEnv("MYSQL_USER", "root"), os.Getenv("MYSQL_PASSWORD"), core.GetEnv("MYSQL_HOST", "127.0.0.1"), core.GetEnvInt("MYSQL_PORT", 3306), dbname)
			mysqlStore, err := mysql.NewMysqlStore(ctx, connString)
			if err != nil {
				return nil, fmt.Errorf("failed to create mysql store: %w", err)
			}

			return mysqlStore, nil
		}

		mysqlStore, err := mysql.NewMysqlStore(ctx, path)
		if err != nil {
			return nil, fmt.Errorf("failed to create mysql store: %w", err)
		}

		return mysqlStore, nil

	default:
		return nil, fmt.Errorf("invalid database type: %s", database)
	}
}
