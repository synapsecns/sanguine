// Package rest provides RESTful API services for RFQ
package rest

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ipfs/go-log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/synapsecns/sanguine/core/ginhelper"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/core/metrics"
	baseServer "github.com/synapsecns/sanguine/core/server"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/api/config"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/docs"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
)

// QuoterAPIServer is a struct that holds the configuration, database connection, gin engine, RPC client, metrics handler, and fast bridge contracts.
// It is used to initialize and run the API server.
type QuoterAPIServer struct {
	cfg                 config.Config
	db                  db.APIDB
	engine              *gin.Engine
	omnirpcClient       omniClient.RPCClient
	handler             metrics.Handler
	fastBridgeContracts map[uint32]*fastbridge.FastBridge
	roleCache           map[uint32]*ttlcache.Cache[string, bool]
	// relayAckCache contains a set of transactionID values that reflect
	// transactions that have been acked for relay
	relayAckCache *ttlcache.Cache[string, string]
	// ackMux is a mutex used to ensure that only one transaction id can be acked at a time.
	ackMux sync.Mutex
}

// NewAPI holds the configuration, database connection, gin engine, RPC client, metrics handler, and fast bridge contracts.
// It is used to initialize and run the API server.
func NewAPI(
	ctx context.Context,
	cfg config.Config,
	handler metrics.Handler,
	omniRPCClient omniClient.RPCClient,
	store db.APIDB,
) (*QuoterAPIServer, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}
	if handler == nil {
		return nil, fmt.Errorf("handler is nil")
	}
	if omniRPCClient == nil {
		return nil, fmt.Errorf("omniRPCClient is nil")
	}
	if store == nil {
		return nil, fmt.Errorf("store is nil")
	}

	docs.SwaggerInfo.Title = "RFQ Quoter API"

	bridges := make(map[uint32]*fastbridge.FastBridge)
	roles := make(map[uint32]*ttlcache.Cache[string, bool])
	for chainID, bridge := range cfg.Bridges {
		chainClient, err := omniRPCClient.GetChainClient(ctx, int(chainID))
		if err != nil {
			return nil, fmt.Errorf("could not create omnirpc client: %w", err)
		}
		bridges[chainID], err = fastbridge.NewFastBridge(common.HexToAddress(bridge), chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not create bridge contract: %w", err)
		}

		// create the roles cache
		roles[chainID] = ttlcache.New[string, bool](
			ttlcache.WithTTL[string, bool](cacheInterval),
		)
		roleCache := roles[chainID]
		go roleCache.Start()
		go func() {
			<-ctx.Done()
			roleCache.Stop()
		}()
	}

	// create the relay ack cache
	relayAckCache := ttlcache.New[string, string](
		ttlcache.WithTTL[string, string](cfg.GetRelayAckTimeout()),
		ttlcache.WithDisableTouchOnHit[string, string](),
	)
	go relayAckCache.Start()
	go func() {
		<-ctx.Done()
		relayAckCache.Stop()
	}()

	return &QuoterAPIServer{
		cfg:                 cfg,
		db:                  store,
		omnirpcClient:       omniRPCClient,
		handler:             handler,
		fastBridgeContracts: bridges,
		roleCache:           roles,
		relayAckCache:       relayAckCache,
		ackMux:              sync.Mutex{},
	}, nil
}

const (
	// QuoteRoute is the API endpoint for handling quote related requests.
	QuoteRoute = "/quotes"
	// AckRoute is the API endpoint for handling relay ack related requests.
	AckRoute      = "/ack"
	cacheInterval = time.Minute
)

var logger = log.Logger("rfq-api")

// Run runs the quoter api server.
func (r *QuoterAPIServer) Run(ctx context.Context) error {
	// TODO: Use Gin Helper
	engine := ginhelper.New(logger)
	h := NewHandler(r.db)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Apply AuthMiddleware only to the PUT routes
	quotesPut := engine.Group(QuoteRoute)
	quotesPut.Use(r.AuthMiddleware())
	quotesPut.PUT("", h.ModifyQuote)
	ackPut := engine.Group(AckRoute)
	ackPut.Use(r.AuthMiddleware())
	ackPut.PUT("", r.PutRelayAck)

	// GET routes without the AuthMiddleware
	// engine.PUT("/quotes", h.ModifyQuote)
	engine.GET(QuoteRoute, h.GetQuotes)

	r.engine = engine

	connection := baseServer.Server{}
	fmt.Printf("starting api at http://localhost:%s\n", r.cfg.Port)
	err := connection.ListenAndServe(ctx, fmt.Sprintf(":%s", r.cfg.Port), r.engine)
	if err != nil {
		return fmt.Errorf("could not start rest api server: %w", err)
	}

	return nil
}

// AuthMiddleware is the Gin authentication middleware that authenticates requests using EIP191.
func (r *QuoterAPIServer) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loggedRequest interface{}
		var destChainID uint32
		var err error

		// Parse the dest chain id from the request
		switch c.Request.URL.Path {
		case QuoteRoute:
			var req model.PutQuoteRequest
			err = c.BindJSON(&req)
			if err == nil {
				destChainID = uint32(req.DestChainID)
				loggedRequest = &req
			}
		case AckRoute:
			var req model.PutAckRequest
			err = c.BindJSON(&req)
			if err == nil {
				destChainID = uint32(req.DestChainID)
				loggedRequest = &req
			}
		default:
			err = fmt.Errorf("unexpected request path: %s", c.Request.URL.Path)
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Authenticate and fetch the address from the request
		addressRecovered, err := r.checkRole(c, destChainID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			c.Abort()
			return
		}

		// Log and pass to the next middleware if authentication succeeds
		// Store the request in context after binding and validation
		c.Set("putRequest", loggedRequest)
		c.Set("relayerAddr", addressRecovered.Hex())
		c.Next()
	}
}

func (r *QuoterAPIServer) checkRole(c *gin.Context, destChainID uint32) (addressRecovered common.Address, err error) {
	bridge, ok := r.fastBridgeContracts[destChainID]
	if !ok {
		err = fmt.Errorf("dest chain id not supported: %d", destChainID)
		return addressRecovered, err
	}

	ops := &bind.CallOpts{Context: c}
	relayerRole := crypto.Keccak256Hash([]byte("RELAYER_ROLE"))

	// authenticate relayer signature with EIP191
	deadline := time.Now().Unix() - 1000 // TODO: Replace with some type of r.cfg.AuthExpiryDelta
	addressRecovered, err = EIP191Auth(c, deadline)
	if err != nil {
		err = fmt.Errorf("unable to authenticate relayer: %w", err)
		return addressRecovered, err
	}

	hasRole := r.roleCache[destChainID].Get(addressRecovered.Hex())

	if hasRole == nil || hasRole.IsExpired() {
		has, roleErr := bridge.HasRole(ops, relayerRole, addressRecovered)
		if roleErr == nil {
			r.roleCache[destChainID].Set(addressRecovered.Hex(), has, cacheInterval)
		}

		if roleErr != nil {
			err = fmt.Errorf("unable to check relayer role on-chain")
			return addressRecovered, err
		} else if !has {
			err = fmt.Errorf("q.Relayer not an on-chain relayer")
			return addressRecovered, err
		}
	}
	return addressRecovered, nil
}

// PutRelayAck checks if a relay is pending or not.
// Note that the ack is not binding; that is, any relayer can still relay the transaction
// on chain if they ignore the response to this call.
// Also, this will not work if the API is run on multiple servers, since there is no inter-server
// communication to maintain the cache.
//
// PUT /ack.
// @dev Protected Method: Authentication is handled through middleware in server.go.
// @Summary Relay ack
// @Schemes
// @Description cache an ack request to synchronize relayer actions.
// @Param request body model.PutQuoteRequest true "query params"
// @Tags ack
// @Accept json
// @Produce json
// @Success 200
// @Router /ack [put].
func (r *QuoterAPIServer) PutRelayAck(c *gin.Context) {
	req, exists := c.Get("putRequest")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not found"})
		return
	}
	rawRelayerAddr, exists := c.Get("relayerAddr")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No relayer address recovered from signature"})
		return
	}
	relayerAddr, ok := rawRelayerAddr.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid relayer address type"})
		return
	}
	ackReq, ok := req.(*model.PutAckRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request type"})
		return
	}

	// If the tx id is already in the cache, it should not be relayed.
	// Otherwise, insert the current relayer's address into the cache.
	r.ackMux.Lock()
	ack := r.relayAckCache.Get(ackReq.TxID)
	shouldRelay := ack == nil
	if shouldRelay {
		r.relayAckCache.Set(ackReq.TxID, relayerAddr, ttlcache.DefaultTTL)
	} else {
		relayerAddr = ack.Value()
	}
	r.ackMux.Unlock()

	resp := relapi.PutRelayAckResponse{
		TxID:           ackReq.TxID,
		ShouldRelay:    shouldRelay,
		RelayerAddress: relayerAddr,
	}
	c.JSON(http.StatusOK, resp)
}
