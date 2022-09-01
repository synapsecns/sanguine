package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/friendsofgo/graphiql"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/services/scribe/server/graph"
	resolvers "github.com/synapsecns/sanguine/services/scribe/server/graph/resolver"
)

var server = handler.NewDefaultServer(
	resolvers.NewExecutableSchema(
		resolvers.Config{Resolvers: &graph.Resolver{}},
	),
)

func graphqlHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	}
}

func graphiqlHandler() gin.HandlerFunc {
	h, _ := graphiql.NewGraphiqlHandler(GraphqlEndpoint)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
