package cmd

import (
	"context"
	"os"
	"sync/atomic"
	"time"

	"github.com/synapsecns/sanguine/agents/agents/guard/db"
	"github.com/synapsecns/sanguine/agents/agents/guard/db/sql/mysql"
	"github.com/synapsecns/sanguine/agents/agents/guard/db/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/agents/guard/metadata"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/client"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm/schema"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/hedzr/log"
	"github.com/jftuga/termsize"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/agents/guard/api"
	"golang.org/x/sync/errgroup"

	// used to embed markdown.
	_ "embed"
	"fmt"

	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/core"
	"github.com/urfave/cli/v2"
)

//go:embed cmd.md
var help string

// GuardInfoCommand gets info about using the guard agent.
var GuardInfoCommand = &cli.Command{
	Name:        "guard-info",
	Description: "learn how to use guard cli",
	Action: func(c *cli.Context) error {
		fmt.Println(string(markdown.Render(help, termsize.Width(), 6)))
		return nil
	},
}

var metricsPortFlag = &cli.UintFlag{
	Name:  "metrics-port",
	Usage: "--port 5121",
	Value: 0,
}

var configFlag = &cli.StringFlag{
	Name:      "config",
	Usage:     "--config /Users/synapsecns/guard_config.yaml",
	TakesFile: true,
	Required:  true,
}

var debugFlag = &cli.BoolFlag{
	Name:  "debug",
	Usage: "--debug",
}

func createGuardParameters(ctx context.Context, c *cli.Context, metrics metrics.Handler) (guardConfig config.AgentConfig, guardDB db.GuardDB, err error) {
	guardConfig, err = config.DecodeAgentConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
	if err != nil {
		return guardConfig, nil, fmt.Errorf("failed to decode config: %w", err)
	}

	if guardConfig.DBPrefix == "" && guardConfig.DBConfig.Type == dbcommon.Mysql.String() {
		guardConfig.DBPrefix = "guard"
	}

	if guardConfig.DBConfig.Type == dbcommon.Sqlite.String() {
		guardConfig.DBPrefix = ""
	}

	guardDB, err = InitGuardDB(ctx,
		guardConfig.DBConfig.Type,
		guardConfig.DBConfig.Source,
		guardConfig.DBPrefix,
		metrics,
	)
	if err != nil {
		return guardConfig, nil, fmt.Errorf("failed to init guard db: %w", err)
	}

	return guardConfig, guardDB, nil
}

// GuardRunCommand runs the guard.
var GuardRunCommand = &cli.Command{
	Name:        "guard-run",
	Description: "runs the guard service",
	Flags:       []cli.Flag{configFlag, metricsPortFlag},
	Action: func(c *cli.Context) error {
		handler, err := metrics.NewFromEnv(c.Context, metadata.BuildInfo())
		if err != nil {
			return fmt.Errorf("failed to create metrics handler: %w", err)
		}

		guardConfig, guardDB, err := createGuardParameters(c.Context, c, handler)
		if err != nil {
			return fmt.Errorf("failed to create guard parameters: %w", err)
		}

		var baseOmniRPCClient omnirpcClient.RPCClient
		if debugFlag.IsSet() {
			baseOmniRPCClient = omnirpcClient.NewOmnirpcClient(guardConfig.BaseOmnirpcURL, handler, omnirpcClient.WithCaptureReqRes())
		} else {
			baseOmniRPCClient = omnirpcClient.NewOmnirpcClient(guardConfig.BaseOmnirpcURL, handler)
		}

		var shouldRetryAtomic atomic.Bool
		shouldRetryAtomic.Store(true)

		for shouldRetryAtomic.Load() {
			shouldRetryAtomic.Store(false)

			g, _ := errgroup.WithContext(c.Context)

			embedded := client.NewEmbeddedScribe(
				guardConfig.DBConfig.Type,
				guardConfig.DBConfig.Source,
				handler,
			)

			g.Go(func() error {
				err := embedded.Start(c.Context)
				if err != nil {
					return fmt.Errorf("failed to start embedded scribe: %w", err)
				}

				return nil
			})

			guard, err := guard.NewGuard(c.Context, guardConfig, baseOmniRPCClient, embedded.ScribeClient, guardDB, handler)
			if err != nil {
				return fmt.Errorf("failed to create guard: %w", err)
			}

			g.Go(func() error {
				err = guard.Start(c.Context)
				if err != nil {
					shouldRetryAtomic.Store(true)

					log.Errorf("Error running guard, will sleep for a minute and retry: %v", err)
					time.Sleep(60 * time.Second)
					return fmt.Errorf("failed to run guard: %w", err)
				}

				return nil
			})

			g.Go(func() error {
				err := api.Start(c.Context, uint16(c.Uint(metricsPortFlag.Name)))
				if err != nil {
					return fmt.Errorf("failed to start api: %w", err)
				}

				return nil
			})

			if err := g.Wait(); err != nil {
				return fmt.Errorf("failed to run guard: %w", err)
			}
		}

		return nil
	},
}

func init() {
	metricsPortFlag.Value = uint(freeport.GetPort())
}

// InitGuardDB initializes a database given a database type and path.
//
//nolint:cyclop
func InitGuardDB(parentCtx context.Context, database string, path string, tablePrefix string, handler metrics.Handler) (_ db.GuardDB, err error) {
	ctx, span := handler.Tracer().Start(parentCtx, "InitGuardDB", trace.WithAttributes(
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

			mysqlStore, err := mysql.NewMysqlStore(ctx, connString, handler)
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

		mysqlStore, err := mysql.NewMysqlStore(ctx, path, handler)
		if err != nil {
			return nil, fmt.Errorf("failed to create mysql store: %w", err)
		}

		return mysqlStore, nil

	default:
		return nil, fmt.Errorf("invalid database type: %s", database)
	}
}
