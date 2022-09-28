package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph"
	resolvers "github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/resolver"
)

const (
	// GraphqlEndpoint is the base endpoint for graphql and the endpoint for post requests to the graphql service.
	GraphqlEndpoint string = "/graphql"
	// GraphiqlEndpoint is the endpoint for the graphql user interface.
	GraphiqlEndpoint string = "/graphiql"
)

// EnableGraphql enables the scribe graphql service.
func EnableGraphql(engine *gin.Engine, consumerDB db.ConsumerDB) {
	server := handler.NewDefaultServer(
		resolvers.NewExecutableSchema(
			resolvers.Config{Resolvers: &graph.Resolver{
				DB: consumerDB,
			}},
		),
	)

	engine.GET(GraphqlEndpoint, graphqlHandler(server))
	engine.POST(GraphqlEndpoint, graphqlHandler(server))
	engine.GET(GraphiqlEndpoint, graphiqlHandler())
}
