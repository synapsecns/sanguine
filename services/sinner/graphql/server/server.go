package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/friendsofgo/graphiql"
	"github.com/gin-gonic/gin"
)

func graphqlHandler(server *handler.Server) gin.HandlerFunc {
	return gin.WrapH(server)
}

func graphiqlHandler() gin.HandlerFunc {
	h, err := graphiql.NewGraphiqlHandler(GraphqlEndpoint)
	if err != nil {
		panic(err)
	}

	return gin.WrapH(h)
}
