package api

import (
	"context"
	"embed"
	"fmt"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/mysql"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	gqlServer "github.com/synapsecns/sanguine/services/scribe/graphql/server"
	"github.com/synapsecns/sanguine/services/scribe/grpc/server"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"os"
	"time"
)

// HealthCheck is the health check endpoint.
const HealthCheck string = "/health-check"

//go:embed static
var static embed.FS

// Config contains the config for the api.
type Config struct {
	// HTTPPort is the http port for the api
	HTTPPort uint16
	// Database is the database type
	// TODO: should be enum
	Database string
	// Path is the path to the database or db connection
	// TODO: should be renamed
	Path string
	// GRPCPort is the path to the grpc service
	GRPCPort uint16
}

// Start starts the api server.
func Start(ctx context.Context, cfg Config) error {
	router := gin.New()

	router.Use(helmet.Default())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS"},
		MaxAge:          12 * time.Hour,
	}))

	// initialize the database
	eventDB, err := InitDB(ctx, cfg.Database, cfg.Path)
	if err != nil {
		return fmt.Errorf("could not initialize database: %w", err)
	}

	gqlServer.EnableGraphql(router, eventDB)
	grpcServer, err := server.SetupGRPCServer(ctx, router, eventDB)
	if err != nil {
		return fmt.Errorf("could not create grpc server: %w", err)
	}

	router.GET(HealthCheck, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	router.GET("static", gin.WrapH(http.FileServer(http.FS(static))))

	fmt.Printf("started graphiql gqlServer on port: http://localhost:%d/graphiql\n", cfg.HTTPPort)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		connection := baseServer.Server{}
		err = connection.ListenAndServe(ctx, fmt.Sprintf(":%d", cfg.HTTPPort), router)
		if err != nil {
			return fmt.Errorf("could not start gqlServer: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		var lc net.ListenConfig
		listener, err := lc.Listen(ctx, "tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
		if err != nil {
			return fmt.Errorf("could not start listener: %w", err)
		}

		err = grpcServer.Serve(listener)
		if err != nil {
			return fmt.Errorf("could not start grpc server: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		grpcServer.Stop()
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
