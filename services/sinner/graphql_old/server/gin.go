package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-gonic/gin"
	"github.com/ravilushqa/otelgqlgen"
	"github.com/synapsecns/sanguine/core/metrics"
	serverConfig "github.com/synapsecns/sanguine/services/sinner/config/server"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/fetcher"
	"github.com/synapsecns/sanguine/services/sinner/graphql_old/server/graph"
	resolvers "github.com/synapsecns/sanguine/services/sinner/graphql_old/server/graph/resolver"
	"time"
)

const (
	// GraphqlEndpoint is the base endpoint for graphql_old and the endpoint for post requests to the graphql_old service.
	GraphqlEndpoint string = "/graphql_old"
	// GraphiqlEndpoint is the endpoint for the graphql_old user interface.
	GraphiqlEndpoint string = "/graphiql"
)

// EnableGraphql enables the sinner graphql_old service.
func EnableGraphql(engine *gin.Engine, consumerDB db.EventDB, fetcher fetcher.ScribeFetcher, config serverConfig.Config, handler metrics.Handler) {
	server := createServer(
		resolvers.NewExecutableSchema(
			resolvers.Config{Resolvers: &graph.Resolver{
				DB:      consumerDB,
				Fetcher: fetcher,
				Config:  config,
			}},
		),
	)
	// TODO; investigate WithCreateSpanFromFields(predicate)
	server.Use(otelgqlgen.Middleware(otelgqlgen.WithTracerProvider(handler.GetTracerProvider())))

	engine.GET(GraphqlEndpoint, graphqlHandler(server))
	engine.POST(GraphqlEndpoint, graphqlHandler(server))
	engine.GET(GraphiqlEndpoint, graphiqlHandler())
}

// Create a server without introspection.
func createServer(es graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(es)
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.FixedComplexityLimit(300)) // Prevent ddos
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	return srv
}
