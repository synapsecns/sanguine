package rest

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/metrics"
	baseServer "github.com/synapsecns/sanguine/core/server"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/api/config"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
)

type APIServer struct {
	cfg                 config.Config
	db                  db.ApiDB
	engine              *gin.Engine
	omnirpcClient       omniClient.RPCClient
	handler             metrics.Handler
	fastBridgeContracts map[uint32]*fastbridge.FastBridge
}

func NewAPI(
	ctx context.Context,
	cfg config.Config,
	handler metrics.Handler,
	omniRPCClient omniClient.RPCClient,
	store db.ApiDB,
) (*APIServer, error) {
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

	bridges := make(map[uint32]*fastbridge.FastBridge)
	for chainID, bridge := range cfg.Bridges {
		chainClient, err := omniRPCClient.GetChainClient(ctx, int(chainID))
		if err != nil {
			return nil, fmt.Errorf("could not create omnirpc client: %w", err)
		}
		bridges[chainID], err = fastbridge.NewFastBridge(common.HexToAddress(bridge), chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not create bridge contract: %w", err)
		}
	}

	return &APIServer{
		cfg:                 cfg,
		db:                  store,
		omnirpcClient:       omniRPCClient,
		handler:             handler,
		fastBridgeContracts: bridges,
	}, nil
}

// Run runs the rest api server.
func (r *APIServer) Run(ctx context.Context) error {

	// TODO: Use Gin Helper
	engine := gin.Default()
	h := NewHandler(r.db)

	// Apply AuthMiddleware only to the PUT route
	quotesPut := engine.Group("/quotes")
	quotesPut.Use(r.AuthMiddleware())
	quotesPut.PUT("", h.ModifyQuote)
	// engine.PUT("/quotes", h.ModifyQuote)
	// GET routes without the AuthMiddleware
	engine.GET("/quotes", h.GetQuotes)
	engine.GET("/quotes/filter", h.GetFilteredQuotes)

	r.engine = engine

	connection := baseServer.Server{}
	err := connection.ListenAndServe(ctx, fmt.Sprintf(":%s", r.cfg.Port), r.engine)
	if err != nil {
		return fmt.Errorf("could not start rest api server: %w", err)
	}

	return nil
}

type PutRequest struct {
	ID            int    `json:"id"`
	DestChainID   string `json:"dest_chain_id"`
	DestTokenAddr string `json:"dest_token_addr"`
	DestAmount    string `json:"dest_amount"`
	Price         string `json:"price"`
}

// AuthMiddleware is the Gin authentication middleware that authenticates requests using EIP191.
func (r *APIServer) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req PutRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		destChainID, err := strconv.ParseUint(req.DestChainID, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dest_chain_id"})
			c.Abort()
			return
		}

		bridge, ok := r.fastBridgeContracts[uint32(destChainID)]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "dest chain id not supported"})
			c.Abort()
			return
		}

		ops := &bind.CallOpts{Context: c}
		relayerRole := crypto.Keccak256Hash([]byte("RELAYER_ROLE"))

		// authenticate relayer signature with EIP191
		deadline := time.Now().Unix() - 1000 // TODO: Replace with some type of r.cfg.AuthExpiryDelta
		addressRecovered, err := EIP191Auth(c, deadline)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("unable to authenticate relayer: %v", err)})
			c.Abort()
			return
		}

		has, err := bridge.HasRole(ops, relayerRole, addressRecovered)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "unable to check relayer role on-chain"})
			c.Abort()
			return
		} else if !has {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "q.Relayer not an on-chain relayer"})
			c.Abort()
			return
		}

		// Log and pass to the next middleware if authentication succeeds
		fmt.Println("Authentication successful for:", addressRecovered.Hex())
		// Store the request in context after binding and validation
		c.Set("putRequest", &req)
		c.Next()
	}
}
