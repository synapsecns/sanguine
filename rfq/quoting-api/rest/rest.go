package rest

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"strconv"
	"time"

	"github.com/synapsecns/sanguine/rfq/quoting-api/bindings"
	"github.com/synapsecns/sanguine/rfq/quoting-api/config"
	"github.com/synapsecns/sanguine/rfq/quoting-api/db"
	"github.com/synapsecns/sanguine/rfq/quoting-api/db/models"
	"github.com/synapsecns/sanguine/rfq/quoting-api/rest/auth"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	// EVMClient "github.com/synapsecns/sanguine/ethergo/client".

	"github.com/gin-gonic/gin"
)

type RestApiServer struct {
	cfg     *config.Config // TODO: cfg
	db      *db.Database
	engine  *gin.Engine
	bridges map[uint]*bindings.FastBridge
}

func NewRestApiServer(ctx context.Context, cfg *config.Config) (*RestApiServer, error) {
	db, err := db.NewDatabase(ctx, metrics.NewNullHandler(), true)
	if err != nil {
		return nil, fmt.Errorf("could not create db: %w", err)
	}
	engine := gin.Default()
	r := RestApiServer{cfg: cfg, db: db, engine: engine}
	return &r, nil
}

// Setup initializes rest api server routes.
func (r *RestApiServer) Setup() {
	r.engine.GET("/ping", r.Ping)
	r.engine.POST("/quote", r.CreateQuote)
	r.engine.GET("/quote", r.ReadQuotes)
	r.engine.GET("/quote/:id", r.ReadQuote)
	r.engine.PUT("/quote/:id", r.UpdateQuote)
	r.engine.DELETE("/quote/:id", r.DeleteQuote)
	r.engine.POST("/quote/:id/ping", r.PingQuote)
}

// Run runs the rest api server.
func (r *RestApiServer) Run() {
	r.engine.Run()
}

// Authenticate checks request header for EIP191 signature for a valid relayer.
func (r *RestApiServer) Authenticate(c *gin.Context, q *models.Quote) {
	// check relayer registered with contract
	bridge, ok := r.bridges[q.DestChainId]
	if !ok {
		err := fmt.Errorf("Dest chain id not supported")
		c.JSON(400, gin.H{"msg": err})
		return
	}

	// call on-chain to dest chain bridge::HasRole for relayer role
	ops := &bind.CallOpts{Context: c}
	role := crypto.Keccak256Hash([]byte("RELAYER_ROLE")) // keccak256("RELAYER_ROLE")
	relayer := common.HexToAddress(q.Relayer)

	if has, err := bridge.HasRole(ops, role, relayer); err != nil {
		err := fmt.Errorf("Unable to check relayer role on-chain")
		c.JSON(400, gin.H{"msg": err})
		return
	} else if !has {
		err := fmt.Errorf("q.Relayer not an on-chain relayer")
		c.JSON(400, gin.H{"msg": err})
		return
	}

	// authenticate relayer signature with EIP191
	deadline := time.Now().Unix() - r.cfg.AuthExpiryDelta
	auth.EIP191Auth(q.Relayer, deadline)(c)
}

// GET /ping.
func (r *RestApiServer) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"result": "pong"})
}

// POST /quote.
func (r *RestApiServer) CreateQuote(c *gin.Context) {
	var q models.Quote
	c.Bind(&q)
	r.Authenticate(c, &q)

	id, err := r.db.InsertQuote(c, &q)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"result": id})
}

// GET /quote
// query: originChainId, destChainId, originToken, destToken, originAmount, deadline (+ relayer?)
func (r *RestApiServer) ReadQuotes(c *gin.Context) {
	var req models.Request
	c.Bind(&req)
	qs, err := r.db.GetQuotes(c, &req)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"result": qs})
}

// GET /quote/{id}.
func (r *RestApiServer) ReadQuote(c *gin.Context) {
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
func (r *RestApiServer) UpdateQuote(c *gin.Context) {
	var q models.Quote
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	q.ID = uint(id) // get q.ID first from URI
	c.Bind(&q)      // binds remaining form data

	// check relayer hasn't changed
	if quote, err := r.db.GetQuote(c, q.ID); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	} else if quote.Relayer != q.Relayer {
		err := fmt.Errorf("Quote relayer not same")
		c.JSON(400, gin.H{"msg": err})
		return
	}
	r.Authenticate(c, &q)

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
func (r *RestApiServer) DeleteQuote(c *gin.Context) {
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
	r.Authenticate(c, &q)

	if err := r.db.DeleteQuote(c, q.ID); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"result": q.ID})
}

// POST /quote/{id}/ping.
func (r *RestApiServer) PingQuote(c *gin.Context) {
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
	r.Authenticate(c, &q)

	if err := r.db.UpdateQuote(c, &q); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"result": q.ID})
}
