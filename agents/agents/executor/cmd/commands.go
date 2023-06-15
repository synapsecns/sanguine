package cmd

import (
	"context"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/jftuga/termsize"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	"github.com/synapsecns/sanguine/agents/agents/executor/api"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/mysql"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/agents/executor/metadata"
	"github.com/synapsecns/sanguine/core/metrics"
	scribeAPI "github.com/synapsecns/sanguine/services/scribe/api"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/client"
	scribeCmd "github.com/synapsecns/sanguine/services/scribe/cmd"
	"github.com/synapsecns/sanguine/services/scribe/node"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm/schema"

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

// ExecutorInfoCommand gets info about using the executor agent.
var ExecutorInfoCommand = &cli.Command{
	Name:        "executor-info",
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

var metricsPortFlag = &cli.UintFlag{
	Name:  "metrics-port",
	Usage: "--port 5121",
	Value: 0,
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

var scribeURL = &cli.StringFlag{
	Name:  "scribe-url",
	Usage: "--scribe-url <url>",
}

// nolint:dupword
func createExecutorParameters(c *cli.Context, metrics metrics.Handler) (executorConfig config.Config, executorDB db.ExecutorDB, clients map[uint32]executor.Backend, err error) {
	executorConfig, err = config.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
	if err != nil {
		return executorConfig, nil, nil, fmt.Errorf("failed to decode config: %w", err)
	}

	if executorConfig.DBPrefix == "" && c.String(dbFlag.Name) == "mysql" {
		executorConfig.DBPrefix = "executor"
	}

	if executorConfig.DBPrefix != "" && c.String(dbFlag.Name) == "sqlite" {
		executorConfig.DBPrefix = ""
	}

	executorDB, err = InitExecutorDB(c.Context, c.String(dbFlag.Name), c.String(pathFlag.Name), executorConfig.DBPrefix, metrics)
	if err != nil {
		return executorConfig, nil, nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	clients = make(map[uint32]executor.Backend)
	/* for _, execClient := range executorConfig.Chains {
		rpcDial, err := rpc.DialContext(c.Context, fmt.Sprintf("%s/%d/rpc/%d", executorConfig.BaseOmnirpcURL, 1, execClient.ChainID))
		if err != nil {
			return executorConfig, nil, nil, fmt.Errorf("failed to dial rpc: %w", err)
		}

		ethClient := ethclient.NewClient(rpcDial)
		clients[execClient.ChainID] = ethClient
	} */
	for _, execClient := range executorConfig.Chains {
		rpcDial, err := rpc.DialContext(c.Context, execClient.TempRPC)
		if err != nil {
			return executorConfig, nil, nil, fmt.Errorf("failed to dial rpc: %w", err)
		}

		ethClient := ethclient.NewClient(rpcDial)
		clients[execClient.ChainID] = ethClient
	}

	return executorConfig, executorDB, clients, nil
}

// ExecutorRunCommand runs the executor.
var ExecutorRunCommand = &cli.Command{
	Name:        "executor-run",
	Description: "runs the executor service",
	Flags: []cli.Flag{configFlag, dbFlag, pathFlag, scribeTypeFlag, metricsPortFlag,
		// The flags below are used when `scribeTypeFlag` is set to "embedded".
		scribeDBFlag, scribePathFlag,
		// The flags below are used when `scribeTypeFlag` is set to "remote".
		scribePortFlag, scribeURL},
	Action: func(c *cli.Context) error {
		metricsProvider := metrics.Get()

		executorConfig, executorDB, clients, err := createExecutorParameters(c, metricsProvider)
		if err != nil {
			return err
		}

		var scribeClient client.ScribeClient

		g, _ := errgroup.WithContext(c.Context)

		switch c.String(scribeTypeFlag.Name) {
		case "embedded":
			eventDB, err := scribeAPI.InitDB(c.Context, c.String(scribeDBFlag.Name), c.String(scribePathFlag.Name), metricsProvider, false)
			if err != nil {
				return fmt.Errorf("failed to initialize database: %w", err)
			}

			scribeClients := make(map[uint32][]backfill.ScribeBackend)

			for _, client := range executorConfig.EmbeddedScribeConfig.Chains {
				for confNum := 1; confNum <= scribeCmd.MaxConfirmations; confNum++ {
					backendClient, err := backfill.DialBackend(c.Context, fmt.Sprintf("%s/%d/rpc/%d", executorConfig.BaseOmnirpcURL, confNum, client.ChainID), metricsProvider)
					if err != nil {
						return fmt.Errorf("could not start client for %s", fmt.Sprintf("%s/1/rpc/%d", executorConfig.BaseOmnirpcURL, client.ChainID))
					}

					scribeClients[client.ChainID] = append(scribeClients[client.ChainID], backendClient)
				}
			}

			scribe, err := node.NewScribe(eventDB, scribeClients, executorConfig.EmbeddedScribeConfig, metricsProvider)
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

			embedded := client.NewEmbeddedScribe(c.String(scribeDBFlag.Name), c.String(scribePathFlag.Name), metricsProvider)

			g.Go(func() error {
				err := embedded.Start(c.Context)
				if err != nil {
					return fmt.Errorf("failed to start embedded scribe: %w", err)
				}

				return nil
			})

			scribeClient = embedded.ScribeClient
		case "remote":
			scribeClient = client.NewRemoteScribe(uint16(c.Uint(scribePortFlag.Name)), c.String(scribeURL.Name), metricsProvider).ScribeClient
		default:
			return fmt.Errorf("invalid scribe type: %s", c.String(scribeTypeFlag.Name))
		}

		handler, err := metrics.NewFromEnv(c.Context, metadata.BuildInfo())
		if err != nil {
			return fmt.Errorf("failed to create metrics handler: %w", err)
		}

		executor, err := executor.NewExecutor(c.Context, executorConfig, executorDB, scribeClient, clients, handler)
		if err != nil {
			return fmt.Errorf("failed to create executor: %w", err)
		}

		g.Go(func() error {
			err := api.Start(c.Context, uint16(c.Uint(metricsPortFlag.Name)))
			if err != nil {
				return fmt.Errorf("failed to start api: %w", err)
			}

			return nil
		})

		g.Go(func() error {
			err := executor.Run(c.Context)
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

func init() {
	metricsPortFlag.Value = uint(freeport.GetPort())
}

// InitExecutorDB initializes a database given a database type and path.
//
//nolint:cyclop
func InitExecutorDB(parentCtx context.Context, database string, path string, tablePrefix string, handler metrics.Handler) (_ db.ExecutorDB, err error) {
	ctx, span := handler.Tracer().Start(parentCtx, "InitExecutorDB", trace.WithAttributes(
		attribute.String("database", database),
		attribute.String("path", path),
		attribute.String("tablePrefix", tablePrefix),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	switch {
	case database == "sqlite":
		sqliteStore, err := sqlite.NewSqliteStore(ctx, path, handler)
		if err != nil {
			return nil, fmt.Errorf("failed to create sqlite store: %w", err)
		}

		return sqliteStore, nil

	case database == "mysql":
		if os.Getenv("OVERRIDE_MYSQL") != "" {
			dbname := os.Getenv("MYSQL_DATABASE")
			connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", core.GetEnv("MYSQL_USER", "root"), os.Getenv("MYSQL_PASSWORD"), core.GetEnv("MYSQL_HOST", "127.0.0.1"), core.GetEnvInt("MYSQL_PORT", 3306), dbname)

			mysqlStore, err := mysql.NewMysqlStore(ctx, connString, handler)
			if err != nil {
				return nil, fmt.Errorf("failed to create mysql store: %w", err)
			}

			return mysqlStore, nil
		}

		namingStrategy := schema.NamingStrategy{
			TablePrefix: fmt.Sprintf("%s_", tablePrefix),
		}

		mysql.NamingStrategy = namingStrategy

		mysqlStore, err := mysql.NewMysqlStore(ctx, path, handler)
		if err != nil {
			return nil, fmt.Errorf("failed to create mysql store: %w", err)
		}

		return mysqlStore, nil

	default:
		return nil, fmt.Errorf("invalid database type: %s", database)
	}
}
