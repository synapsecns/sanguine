package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/bindings"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/config"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/db"
)

// NewTestRestAPIServer creates a new test rest api server exposing the constructor.
func NewTestRestAPIServer(cfg *config.Config, db *db.Database, engine *gin.Engine, bridges map[uint]*bindings.FastBridge) *APIServer {
	return &APIServer{
		cfg:     cfg,
		db:      db,
		engine:  engine,
		bridges: bridges,
	}
}

// Engine exports the gin engine for testing.
func (r *APIServer) Engine() *gin.Engine {
	return r.engine
}

// DB exports the database for testing.
func (r *APIServer) DB() *db.Database {
	return r.db
}
