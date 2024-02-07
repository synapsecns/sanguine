package cmd

import (
	"context"
	"github.com/synapsecns/sanguine/agents/agents/notary/db"
	"github.com/synapsecns/sanguine/agents/agents/notary/db/sql/mysql"
	"github.com/synapsecns/sanguine/agents/agents/notary/db/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/agents/notary/metadata"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm/schema"
	"os"
	"sync/atomic"
	"time"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/hedzr/log"
	"github.com/jftuga/termsize"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/agents/notary/api"
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

// NotaryInfoCommand gets info about using the notary agent.
var NotaryInfoCommand = &cli.Command{
	Name:        "notary-info",
	Description: "learn how to use notary cli",
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
	Usage:     "--config /Users/synapsecns/notary_config.yaml",
	TakesFile: true,
	Required:  true,
}

var debugFlag = &cli.BoolFlag{
	Name:  "debug",
	Usage: "--debug",
}

func createNotaryParameters(ctx context.Context, c *cli.Context, metrics metrics.Handler) (notaryConfig config.AgentConfig, notaryDB db.NotaryDB, err error) {
	notaryConfig, err = config.DecodeAgentConfig(core.ExpandOrReturnPath(c.String(configFlag.Name)))
	if err != nil {
		return notaryConfig, nil, fmt.Errorf("failed to decode config: %w", err)
	}

	if notaryConfig.DBPrefix == "" && notaryConfig.DBConfig.Type == dbcommon.Mysql.String() {
		notaryConfig.DBPrefix = "notary"
	}

	if notaryConfig.DBConfig.Type == dbcommon.Sqlite.String() {
		notaryConfig.DBPrefix = ""
	}

	notaryDB, err = InitNotaryDB(
		ctx,
		notaryConfig.DBConfig.Type,
		notaryConfig.DBConfig.Source,
		notaryConfig.DBPrefix,
		metrics,
	)
	if err != nil {
		return notaryConfig, nil, fmt.Errorf("failed to init db: %w", err)
	}

	return notaryConfig, notaryDB, nil
}

// NotaryRunCommand runs the notary.
var NotaryRunCommand = &cli.Command{
	Name:        "notary-run",
	Description: "runs the notary service",
	Flags:       []cli.Flag{configFlag, metricsPortFlag, debugFlag},
	Action: func(c *cli.Context) error {
		metricsProvider, err := metrics.NewFromEnv(c.Context, metadata.BuildInfo())
		if err != nil {
			return fmt.Errorf("failed to create metrics handler: %w", err)
		}

		notaryConfig, notaryDB, err := createNotaryParameters(c.Context, c, metricsProvider)
		if err != nil {
			return fmt.Errorf("failed to create notary parameters: %w", err)
		}

		var baseOmniRPCClient omnirpcClient.RPCClient
		if debugFlag.IsSet() {
			baseOmniRPCClient = omnirpcClient.NewOmnirpcClient(notaryConfig.BaseOmnirpcURL, metricsProvider, omnirpcClient.WithCaptureReqRes())
		} else {
			baseOmniRPCClient = omnirpcClient.NewOmnirpcClient(notaryConfig.BaseOmnirpcURL, metricsProvider)
		}

		var shouldRetryAtomic atomic.Bool
		shouldRetryAtomic.Store(true)

		for shouldRetryAtomic.Load() {
			shouldRetryAtomic.Store(false)

			g, _ := errgroup.WithContext(c.Context)

			notary, err := notary.NewNotary(c.Context, notaryConfig, baseOmniRPCClient, notaryDB, metricsProvider)
			if err != nil {
				return fmt.Errorf("failed to create notary: %w", err)
			}

			g.Go(func() error {
				err = notary.Start(c.Context)
				if err != nil {
					shouldRetryAtomic.Store(true)

					log.Errorf("Error running notary, will sleep for a minute and retry: %v", err)
					time.Sleep(60 * time.Second)
					return fmt.Errorf("failed to create notary: %w", err)
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
				return fmt.Errorf("failed to run notary: %w", err)
			}
		}

		return nil
	},
}

func init() {
	metricsPortFlag.Value = uint(freeport.GetPort())
}

// InitNotaryDB initializes a database given a database type and path.
//
//nolint:cyclop
func InitNotaryDB(parentCtx context.Context, database string, path string, tablePrefix string, handler metrics.Handler) (_ db.NotaryDB, err error) {
	ctx, span := handler.Tracer().Start(parentCtx, "InitNotaryDB", trace.WithAttributes(
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
