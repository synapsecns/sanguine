package cmd

import (
	"context"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/jftuga/termsize"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	"github.com/synapsecns/sanguine/agents/agents/executor/api"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/sql/mysql"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/agents/executor/metadata"
	execConfig "github.com/synapsecns/sanguine/agents/config/executor"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	scribeAPI "github.com/synapsecns/sanguine/services/scribe/api"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/client"
	scribeCmd "github.com/synapsecns/sanguine/services/scribe/cmd"
	"github.com/synapsecns/sanguine/services/scribe/service"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm/schema"

	// used to embed markdown.
	_ "embed"
	"fmt"
	"os"

	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/core"
	"github.com/urfave/cli/v2"
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

var metricsPortFlag = &cli.UintFlag{
	Name:  "metrics-port",
	Usage: "--port 5121",
	Value: 0,
}

var debugFlag = &cli.BoolFlag{
	Name:  "debug",
	Usage: "--debug",
}

func createExecutorParameters(ctx context.Context, c *cli.Context, metrics metrics.Handler) (executorConfig execConfig.Config, executorDB db.ExecutorDB, err error) {
	executorConfig, err = execConfig.DecodeConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
	if err != nil {
		return executorConfig, nil, fmt.Errorf("failed to decode config: %w", err)
	}

	if executorConfig.DBPrefix == "" && executorConfig.DBConfig.Type == dbcommon.Mysql.String() {
		executorConfig.DBPrefix = "executor"
	}

	if executorConfig.DBConfig.Type == dbcommon.Sqlite.String() {
		executorConfig.DBPrefix = ""
	}

	executorDB, err = InitExecutorDB(
		ctx,
		executorConfig.DBConfig.Type,
		executorConfig.DBConfig.Source,
		executorConfig.DBPrefix,
		metrics,
	)
	if err != nil {
		return executorConfig, nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return executorConfig, executorDB, nil
}

// ExecutorRunCommand runs the executor.
var ExecutorRunCommand = &cli.Command{
	Name:        "executor-run",
	Description: "runs the executor service",
	Flags:       []cli.Flag{configFlag, metricsPortFlag, debugFlag},
	Action: func(c *cli.Context) error {

		var scribeClient client.ScribeClient

		g, ctx := errgroup.WithContext(c.Context)

		handler, err := metrics.NewFromEnv(ctx, metadata.BuildInfo())
		if err != nil {
			return fmt.Errorf("failed to create metrics handler: %w", err)
		}

		executorConfig, executorDB, err := createExecutorParameters(ctx, c, handler)
		if err != nil {
			return err
		}

		switch executorConfig.ScribeConfig.Type {
		case "embedded":
			eventDB, err := scribeAPI.InitDB(
				ctx,
				executorConfig.DBConfig.Type,
				executorConfig.DBConfig.Source,
				handler,
				false,
			)
			if err != nil {
				return fmt.Errorf("failed to initialize database: %w", err)
			}

			scribeClients := make(map[uint32][]backend.ScribeBackend)

			for _, client := range executorConfig.ScribeConfig.EmbeddedScribeConfig.Chains {
				for confNum := 1; confNum <= scribeCmd.MaxConfirmations; confNum++ {
					backendClient, err := backend.DialBackend(ctx, fmt.Sprintf("%s/%d/rpc/%d", executorConfig.ScribeConfig.EmbeddedScribeConfig.RPCURL, confNum, client.ChainID), handler)
					if err != nil {
						return fmt.Errorf("could not start client for %s", fmt.Sprintf("%s/1/rpc/%d", executorConfig.ScribeConfig.EmbeddedScribeConfig.RPCURL, client.ChainID))
					}

					scribeClients[client.ChainID] = append(scribeClients[client.ChainID], backendClient)
				}
			}

			scribe, err := service.NewScribe(eventDB, scribeClients, executorConfig.ScribeConfig.EmbeddedScribeConfig, handler)
			if err != nil {
				return fmt.Errorf("failed to initialize scribe: %w", err)
			}

			g.Go(func() error {
				err := scribe.Start(ctx)
				if err != nil {
					return fmt.Errorf("failed to start scribe: %w", err)
				}

				return nil
			})

			embedded := client.NewEmbeddedScribe(
				executorConfig.ScribeConfig.EmbeddedDBConfig.Type,
				executorConfig.DBConfig.Source,
				handler,
			)

			g.Go(func() error {
				err := embedded.Start(ctx)
				if err != nil {
					return fmt.Errorf("failed to start embedded scribe: %w", err)
				}

				return nil
			})

			scribeClient = embedded.ScribeClient
		case "remote":
			scribeClient = client.NewRemoteScribe(
				uint16(executorConfig.ScribeConfig.Port),
				executorConfig.ScribeConfig.URL,
				handler,
			).ScribeClient
		default:
			return fmt.Errorf("invalid scribe type: %s", executorConfig.ScribeConfig.Type)
		}

		var baseOmniRPCClient omnirpcClient.RPCClient
		if debugFlag.IsSet() {
			baseOmniRPCClient = omnirpcClient.NewOmnirpcClient(executorConfig.BaseOmnirpcURL, handler, omnirpcClient.WithCaptureReqRes())
		} else {
			baseOmniRPCClient = omnirpcClient.NewOmnirpcClient(executorConfig.BaseOmnirpcURL, handler)
		}

		executor, err := executor.NewExecutor(ctx, executorConfig, executorDB, scribeClient, baseOmniRPCClient, handler)
		if err != nil {
			return fmt.Errorf("failed to create executor: %w", err)
		}

		g.Go(func() error {
			err := api.Start(ctx, uint16(c.Uint(metricsPortFlag.Name)))
			if err != nil {
				return fmt.Errorf("failed to start api: %w", err)
			}

			return nil
		})

		g.Go(func() error {
			err := executor.Run(ctx)
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
	case database == dbcommon.Sqlite.String():
		sqliteStore, err := sqlite.NewSqliteStore(ctx, path, handler, false)
		if err != nil {
			return nil, fmt.Errorf("failed to create sqlite store: %w", err)
		}

		return sqliteStore, nil

	case database == dbcommon.Mysql.String():
		if os.Getenv("OVERRIDE_MYSQL") != "" {
			dbname := os.Getenv("MYSQL_DATABASE")
			connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", core.GetEnv("MYSQL_USER", "root"), os.Getenv("MYSQL_PASSWORD"), core.GetEnv("MYSQL_HOST", "127.0.0.1"), core.GetEnvInt("MYSQL_PORT", 3306), dbname)

			mysqlStore, err := mysql.NewMysqlStore(ctx, connString, handler, false)
			if err != nil {
				return nil, fmt.Errorf("failed to create mysql store: %w", err)
			}

			return mysqlStore, nil
		}

		if tablePrefix != "" {
			mysql.NamingStrategy = schema.NamingStrategy{
				TablePrefix: fmt.Sprintf("%s_", tablePrefix),
			}
		}

		mysqlStore, err := mysql.NewMysqlStore(ctx, path, handler, false)
		if err != nil {
			return nil, fmt.Errorf("failed to create mysql store: %w", err)
		}

		return mysqlStore, nil

	default:
		return nil, fmt.Errorf("invalid database type: %s", database)
	}
}
