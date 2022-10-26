package metrics

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// nullHandler is a metrics handler that does nothing.
// it is used to allow metrics collection to be skipped.
type nullHandler struct {
}

func (n nullHandler) ConfigureHTTPClient(client *http.Client) {
	// Do nothing
}

func (n nullHandler) Gin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func (n nullHandler) Start(_ context.Context) error {
	return nil
}

// NewNullHandler creates a new null transaction handler.
func NewNullHandler() Handler {
	return &nullHandler{}
}

var _ Handler = &nullHandler{}
