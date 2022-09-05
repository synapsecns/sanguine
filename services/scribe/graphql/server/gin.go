package server

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	syn_server "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/mysql"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server/graph"
	resolvers "github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/resolver"
)

const (
	// GraphqlEndpoint is the base endpoint for graphql and the endpoint for post requests to the graphql service.
	GraphqlEndpoint string = "/graphql"
	// GraphiqlEndpoint is the endpoint for the graphql user interface.
	GraphiqlEndpoint string = "/graphiql"
)

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

// Start starts the server and initializes the database.
func Start(ctx context.Context, port uint16, database string, path string) error {
	// initialize the database
	db, err := initDB(ctx, database, path)
	if err != nil {
		return fmt.Errorf("could not initialize database: %w", err)
	}
	server := handler.NewDefaultServer(
		resolvers.NewExecutableSchema(
			resolvers.Config{Resolvers: &graph.Resolver{
				DB: db,
			}},
		),
	)

	router := gin.New()

	router.Use(helmet.Default())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS"},
		MaxAge:          12 * time.Hour,
	}))

	router.GET(GraphqlEndpoint, graphqlHandler(server))
	router.POST(GraphqlEndpoint, graphqlHandler(server))
	router.GET(GraphiqlEndpoint, graphiqlHandler())

	fmt.Printf("started graphiql server on port: http://localhost:%d/graphiql\n", port)

	connection := syn_server.Server{}
	err = connection.ListenAndServe(ctx, fmt.Sprintf(":%d", port), router)
	if err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}

	return nil
}
