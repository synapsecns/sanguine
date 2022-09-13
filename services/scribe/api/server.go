package api

import (
	"context"
	"fmt"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	syn_server "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/mysql"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server"
	"time"
)

// HealthCheck is the health check endpoint.
const HealthCheck string = "/health-check"

func initDB(ctx context.Context, database string, path string) (db.EventDB, error) {
	switch {
	case database == "sqlite":
		sqliteStore, err := sqlite.NewSqliteStore(ctx, path)
		if err != nil {
			return nil, fmt.Errorf("failed to create sqlite store: %w", err)
		}
		return sqliteStore, nil
	case database == "mysql":
		mysqlStore, err := mysql.NewMysqlStore(ctx, path)
		if err != nil {
			return nil, fmt.Errorf("failed to create mysql store: %w", err)
		}
		return mysqlStore, nil
	default:
		return nil, fmt.Errorf("invalid database type: %s", database)
	}
}

func Start(ctx context.Context, port uint16, database string, path string) error {
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
	eventDB, err := initDB(ctx, database, path)
	if err != nil {
		return fmt.Errorf("could not initialize database: %w", err)
	}

	server.EnableGraphql(router, eventDB)
	router.GET(HealthCheck, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	fmt.Printf("started graphiql server on port: http://localhost:%d/graphiql\n", port)

	connection := syn_server.Server{}
	err = connection.ListenAndServe(ctx, fmt.Sprintf(":%d", port), router)
	if err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}
	return nil
}
