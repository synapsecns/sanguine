package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/ravilushqa/otelgqlgen"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server/graph"
	resolvers "github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/resolver"
)

const (
	// GraphqlEndpoint is the base endpoint for graphql and the endpoint for post requests to the graphql service.
	GraphqlEndpoint string = "/graphql"
	// GraphiqlEndpoint is the endpoint for the graphql user interface.
	GraphiqlEndpoint string = "/graphiql"
)

// EnableGraphql enables the scribe graphql service.
func EnableGraphql(engine *gin.Engine, eventDB db.EventDB, omniRPCURL string, metrics metrics.Handler) {
	server := handler.NewDefaultServer(
		resolvers.NewExecutableSchema(
			resolvers.Config{Resolvers: &graph.Resolver{
				DB:         eventDB,
				OmniRPCURL: omniRPCURL,
				Metrics:    metrics,
			}},
		),
	)
	// TODO; investigate WithCreateSpanFromFields(predicate)
	server.Use(otelgqlgen.Middleware(otelgqlgen.WithTracerProvider(metrics.GetTracerProvider())))

	engine.GET(GraphqlEndpoint, graphqlHandler(server))
	engine.POST(GraphqlEndpoint, graphqlHandler(server))
	engine.GET(GraphiqlEndpoint, graphiqlHandler())
}
