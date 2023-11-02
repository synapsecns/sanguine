package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	serverConfig "github.com/synapsecns/sanguine/services/sinner/config/server"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/db/sql/mysql"
	"github.com/synapsecns/sanguine/services/sinner/db/sql/sqlite"
	gqlServer "github.com/synapsecns/sanguine/services/sinner/graphql/server"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"os"
)

var logger = log.Logger("sinner-api")
var errNoPort = errors.New("port not specified, must be between 1 and 65535")

// Start starts the api server for sinner.
func Start(ctx context.Context, cfg serverConfig.Config, handler metrics.Handler) error {
	if cfg.HTTPPort == 0 {
		return errNoPort
	}
	logger.Warnf("starting api server")
	router := ginhelper.New(logger)
	// wrap gin with metrics
	router.GET(metrics.MetricsPath, gin.WrapH(handler.Handler()))

	eventDB, err := InitDB(ctx, cfg.DBType, cfg.DBPath, handler, cfg.SkipMigrations)
	if err != nil {
		return fmt.Errorf("could not initialize database: %w", err)
	}

	router.Use(handler.Gin())
	gqlServer.EnableGraphql(router, eventDB, cfg, handler)
	fmt.Printf("started graphiql gqlServer on port: http://localhost:%d/graphiql\n", cfg.HTTPPort)
	g, ctx := errgroup.WithContext(ctx)

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", fmt.Sprintf(":%d", cfg.HTTPPort))
	if err != nil {
		return fmt.Errorf("could not listen on port %d", cfg.HTTPPort)
	}

	g.Go(func() error {
		//nolint: gosec
		err := http.Serve(listener, router)
		if err != nil {
			return fmt.Errorf("could not serve http: %w", err)
		}

		return nil
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}

// InitDB initializes a database given a database type and path.
// TODO: use enum for database type.
func InitDB(ctx context.Context, dbTypeStr string, path string, metrics metrics.Handler, skipMigrations bool) (db.EventDB, error) {
	logger.Warnf("Starting database connection from api")

	dbType, err := dbcommon.DBTypeFromString(dbTypeStr)
	if err != nil {
		return nil, fmt.Errorf("invalid databaseType type: %s", dbTypeStr)
	}

	// nolint:exhaustive
	switch dbType {
	case dbcommon.Sqlite:
		sqliteStore, err := sqlite.NewSqliteStore(ctx, path, metrics, skipMigrations)
		if err != nil {
			return nil, fmt.Errorf("failed to create sqlite store: %w", err)
		}

		metrics.AddGormCallbacks(sqliteStore.DB())

		return sqliteStore, nil
	case dbcommon.Mysql:
		if os.Getenv("OVERRIDE_MYSQL") != "" {
			dbname := os.Getenv("MYSQL_DATABASE")
			connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", core.GetEnv("MYSQL_USER", "root"), os.Getenv("MYSQL_PASSWORD"), core.GetEnv("MYSQL_HOST", "127.0.0.1"), core.GetEnvInt("MYSQL_PORT", 3306), dbname)
			mysqlStore, err := mysql.NewMysqlStore(ctx, connString, metrics, skipMigrations)
			if err != nil {
				return nil, fmt.Errorf("failed to create mysql store: %w", err)
			}

			metrics.AddGormCallbacks(mysqlStore.DB())

			return mysqlStore, nil
		}

		mysqlStore, err := mysql.NewMysqlStore(ctx, path, metrics, skipMigrations)
		if err != nil {
			return nil, fmt.Errorf("failed to create mysql store: %w", err)
		}

		return mysqlStore, nil
	default:
		return nil, fmt.Errorf("invalid databaseType type: %s", dbTypeStr)
	}
}
