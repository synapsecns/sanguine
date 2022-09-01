package server

import (
	"fmt"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	// GraphqlEndpoint is the base endpoint for graphql and the endpoint for post requests to the graphql service
	GraphqlEndpoint string = "/graphql"
	// GraphiqlEndpoint is the endpoint for the graphql user interface
	GraphiqlEndpoint string = "/graphiql"
)

// Start starts the server
// other scribe stuff in here
func Start(port uint16) error {
	router := gin.New()

	router.Use(helmet.Default())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS"},
		MaxAge:          12 * time.Hour,
	}))

	router.GET(GraphqlEndpoint, graphqlHandler())
	router.POST(GraphqlEndpoint, graphqlHandler())
	router.GET(GraphiqlEndpoint, graphiqlHandler())

	fmt.Printf("started graphiql server on port: http://localhost:%d/graphiql\n", port)

	// TODO: respect context cancellations
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
