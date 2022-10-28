package api

import (
	"context"
	"fmt"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/metrics"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/services/explorer/consumer"
	"github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	gqlServer "github.com/synapsecns/sanguine/services/explorer/graphql/server"
	"github.com/synapsecns/sanguine/services/explorer/testutil/clickhouse"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

// HealthCheck is the health check endpoint.
const HealthCheck string = "/health-check"

// Config contains the config for the api.
type Config struct {
	// HTTPPort is the http port for the api
	HTTPPort uint16
	// Address is the address of the database
	Address string
	// ScribeURL is the url of the scribe service
	ScribeURL string
}

// Start starts the api server.
func Start(ctx context.Context, cfg Config) error {
	router := gin.New()
	router.Use(helmet.Default())
	router.Use(gin.Recovery())
	router.Use(metrics.Get().Gin())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS"},
		MaxAge:          12 * time.Hour,
	}))

	// initialize the database
	consumerDB, err := InitDB(ctx, cfg.Address)
	if err != nil {
		return fmt.Errorf("could not initialize database: %w", err)
	}

	// get the fetcher
	fetcher := consumer.NewFetcher(client.NewClient(http.DefaultClient, cfg.ScribeURL))

	gqlServer.EnableGraphql(router, consumerDB, *fetcher)

	router.GET(HealthCheck, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

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

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}

// InitDB initializes a database given a database type and path.
func InitDB(ctx context.Context, address string) (db.ConsumerDB, error) {
	// TODO add connection to Google Cloud hosted clickhouse
	if address == "default" {
		cleanup, port, err := clickhouse.NewClickhouseStore("explorer")
		if cleanup == nil {
			return nil, fmt.Errorf("clickhouse spin up failure, no open port found: %w", err)
		}
		if port == nil || err != nil {
			cleanup()
			return nil, fmt.Errorf("clickhouse spin up failure, no open port found: %w", err)
		}
		address = "clickhouse://clickhouse_test:clickhouse_test@localhost:" + fmt.Sprintf("%d", *port) + "/clickhouse_test"
	}
	clickhouseDB, err := sql.OpenGormClickhouse(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	return clickhouseDB, nil
}
