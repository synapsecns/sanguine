package metrics

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// nullHandler is a metrics handler that does nothing.
// it is used to allow metrics collection to be skipped.
type nullHandler struct {
}

func (n nullHandler) StartTransaction(name string) Transaction {
	return nullTransaction{}
}

func (n nullHandler) AddGormCallbacks(db *gorm.DB) {
	// Do nothing
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

type nullTransaction struct {
}

func (n nullTransaction) NewGoroutine() Transaction {
	return n
}

func (n nullTransaction) NewGoRoutine() Transaction {
	return n
}

func (n nullTransaction) End() {
	// do nothing
}

var _ Transaction = &nullTransaction{}
