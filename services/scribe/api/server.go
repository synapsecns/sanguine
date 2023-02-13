package api

import (
	"context"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-log"
	"github.com/soheilhy/cmux"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/mysql"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	gqlServer "github.com/synapsecns/sanguine/services/scribe/graphql/server"
	"github.com/synapsecns/sanguine/services/scribe/grpc/server"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"os"
)

//go:embed static
var static embed.FS

// Config contains the config for the api.
type Config struct {
	// Port is the http port for the api.
	Port uint16
	// Database is the database type.
	// TODO: should be enum
	Database string
	// Path is the path to the database or db connection.
	// TODO: should be renamed
	Path string
	// OmniRPCURL is the url of the omnirpc service.
	OmniRPCURL string
}

var logger = log.Logger("scribe-api")

// Start starts the api server.
func Start(ctx context.Context, cfg Config) error {
	router := ginhelper.New(logger)

	eventDB, err := InitDB(ctx, cfg.Database, cfg.Path)
	if err != nil {
		return fmt.Errorf("could not initialize database: %w", err)
	}

	gqlServer.EnableGraphql(router, eventDB, cfg.OmniRPCURL)
	grpcServer, err := server.SetupGRPCServer(ctx, router, eventDB)
	if err != nil {
		return fmt.Errorf("could not create grpc server: %w", err)
	}

	router.GET("static", gin.WrapH(http.FileServer(http.FS(static))))
	fmt.Printf("started graphiql gqlServer on port: http://localhost:%d/graphiql\n", cfg.Port)
	g, ctx := errgroup.WithContext(ctx)

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return fmt.Errorf("could not listen on port %d", cfg.Port)
	}

	m := cmux.New(listener)
	httpListener := m.Match(cmux.HTTP1Fast())
	// fallback to grpc
	grpcListener := m.Match(cmux.Any())

	g.Go(func() error {
		//nolint: gosec
		// TODO: consider setting timeouts here:  https://ieftimov.com/posts/make-resilient-golang-net-http-servers-using-timeouts-deadlines-context-cancellation/
		err := http.Serve(httpListener, router)
		if err != nil {
			return fmt.Errorf("could not serve http: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		err = grpcServer.Serve(grpcListener)
		if err != nil {
			return fmt.Errorf("could not start grpc server: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		err := m.Serve()
		if err != nil {
			return fmt.Errorf("could not start server: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		grpcServer.Stop()
		m.Close()

		return nil
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}

// InitDB initializes a database given a database type and path.
func InitDB(ctx context.Context, database string, path string) (db.EventDB, error) {
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
