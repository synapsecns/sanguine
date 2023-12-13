// Package rest implements the rest driver
package rest

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/bindings"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/config"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/db"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/db/models"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/rest/auth"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/gin-gonic/gin"
)

// APIServer is the server object.
// TODO: this should be an interface.
type APIServer struct {
	cfg     *config.Config // TODO: cfg
	db      *db.Database   // TODO: carefully consider why this is a pointer
	engine  *gin.Engine
	bridges map[uint]*bindings.FastBridge
}

// NewRestAPIServer creates a new instance of the rest api server.
func NewRestAPIServer(ctx context.Context, cfg *config.Config) (*APIServer, error) {
	apiDB, err := db.NewDatabase(ctx, metrics.NewNullHandler(), false, cfg.DBType, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("could not create db: %w", err)
	}
	engine := gin.Default()
	r := APIServer{cfg: cfg, db: apiDB, engine: engine}
	return &r, nil
}

// Setup initializes rest api server routes.
func (r *APIServer) Setup() {
	r.engine.GET("/ping", r.ping)
	r.engine.POST("/quote", r.createQuote)
	r.engine.GET("/quote", r.readQuotes)
	r.engine.GET("/quote/:id", r.readQuote)
	r.engine.PUT("/quote/:id", r.updateQuote)
	r.engine.DELETE("/quote/:id", r.deleteQuote)
	r.engine.POST("/quote/:id/ping", r.pingQuote)
}

// Run runs the rest api server.
func (r *APIServer) Run() error {
	err := r.engine.Run()
	if err != nil {
		return fmt.Errorf("could not run rest api server: %w", err)
	}
	return nil
}

// Authenticate checks request header for EIP191 signature for a valid relayer.
// TODO: this should be moved to a middleware package.
func (r *APIServer) Authenticate(c *gin.Context, q *models.Quote) (err error) {
	// check relayer registered with contract
	bridge, ok := r.bridges[q.DestChainID]
	if !ok {
		err = fmt.Errorf("dest chain id not supported")
		c.JSON(400, gin.H{"msg": err})
		return err
	}

	// call on-chain to dest chain bridge::HasRole for relayer role
	ops := &bind.CallOpts{Context: c}
	// TODO: change me to RELAYER_ROLE
	role := crypto.Keccak256Hash([]byte("FILLER_ROLE")) // keccak256("RELAYER_ROLE")
	relayer := common.HexToAddress(q.Relayer)

	var has bool
	if has, err = bridge.HasRole(ops, role, relayer); err != nil {
		err = fmt.Errorf("unable to check relayer role on-chain")
		c.JSON(400, gin.H{"msg": err})
		return err
	} else if !has {
		err = fmt.Errorf("q.Relayer not an on-chain relayer")
		c.JSON(400, gin.H{"msg": err})
		return err
	}

	// authenticate relayer signature with EIP191
	deadline := time.Now().Unix() - r.cfg.AuthExpiryDelta
	err = auth.EIP191Auth(c, q.Relayer, deadline)
	if err != nil {
		return fmt.Errorf("unable to authenticate relayer: %w", err)
	}

	return nil
}

// GET /ping.
func (r *APIServer) ping(c *gin.Context) {
	c.JSON(200, gin.H{"result": "pong"})
}

// POST /quote.
func (r *APIServer) createQuote(c *gin.Context) {
	var q models.Quote
	err := c.Bind(&q)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	err = r.Authenticate(c, &q)
	if err != nil {
		return
	}

	id, err := r.db.InsertQuote(c, &q)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"result": id})
}

// GET /quote
// query: originChainId, destChainId, originToken, destToken, originAmount, deadline (+ relayer?)
func (r *APIServer) readQuotes(c *gin.Context) {
	var req models.Request
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	qs, err := r.db.GetQuotes(c, &req)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"result": qs})
}

// GET /quote/{id}.
func (r *APIServer) readQuote(c *gin.Context) {
	var q models.Quote
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	q.ID = uint(id) // get q.ID first from URI

	q, err = r.db.GetQuote(c, q.ID)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"result": q})
}

// PUT /quote/{id}.
func (r *APIServer) updateQuote(c *gin.Context) {
	var q models.Quote
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	q.ID = uint(id) // get q.ID first from URI

	err = c.Bind(&q) // binds remaining form data
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	// check relayer hasn't changed
	if quote, err := r.db.GetQuote(c, q.ID); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	} else if quote.Relayer != q.Relayer {
		err := fmt.Errorf("quote relayer not same")
		c.JSON(400, gin.H{"msg": err})
		return
	}
	err = r.Authenticate(c, &q)
	if err != nil {
		return
	}

	if uint(id) != q.ID {
		err := fmt.Errorf(":id != quote.ID")
		c.JSON(400, gin.H{"msg": err})
		return
	} else if err := r.db.UpdateQuote(c, &q); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"result": q.ID})
}

// DELETE /quote/{id}.
func (r *APIServer) deleteQuote(c *gin.Context) {
	var q models.Quote
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	q.ID = uint(id) // get q.ID first from URI

	// get quote to authenticate relayer
	q, err = r.db.GetQuote(c, q.ID)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	err = r.Authenticate(c, &q)
	if err != nil {
		return
	}

	if err := r.db.DeleteQuote(c, q.ID); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"result": q.ID})
}

// POST /quote/{id}/ping.
func (r *APIServer) pingQuote(c *gin.Context) {
	var q models.Quote
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	q.ID = uint(id) // get q.ID first from URI

	// first get quote then resave with same public info
	// to modify *only* updatedAt
	q, err = r.db.GetQuote(c, q.ID)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	err = r.Authenticate(c, &q)
	if err != nil {
		return
	}

	if err := r.db.UpdateQuote(c, &q); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"result": q.ID})
}
