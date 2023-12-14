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
	fmt.Println("Context:", ctx)
	fmt.Println("Config:", cfg)
	fmt.Println("Handler:", handler)
	fmt.Println("OmniRPCClient:", omniRPCClient)
	fmt.Println("Store:", store)

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
	h := NewHandler()

	engine.Use(r.AuthMiddleware()).PUT("/quotes", h.ModifyQuote)
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
	ID            string `json:"id"`
	DestChainID   string `json:"dest_chain_id"`
	DestTokenAddr string `json:"dest_token_addr"`
	DestAmount    string `json:"dest_amount"`
	Price         string `json:"price"`
}

func (r *APIServer) Authenticate(c *gin.Context) (err error) {

	var req PutRequest
	fmt.Println("Request:", req)
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	destChainID, err := strconv.ParseUint(req.DestChainID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dest_chain_id"})
		return err
	}
	// Now you can access DestChainID as req.DestChainID
	fmt.Println("DestChainID:", req.DestChainID)
	// check relayer registered with contract
	bridge, ok := r.fastBridgeContracts[uint32(destChainID)]
	if !ok {
		err = fmt.Errorf("dest chain id not supported")
		c.JSON(400, gin.H{"msg": err})
		return err
	}

	// // call on-chain to dest chain bridge::HasRole for relayer role
	ops := &bind.CallOpts{Context: c}
	relayer_role := crypto.Keccak256Hash([]byte("RELAYER_ROLE"))

	// authenticate relayer signature with EIP191
	deadline := time.Now().Unix() - 1000 // TODO: Replace with some type of r.cfg.AuthExpiryDelta
	addressRecovered, err := EIP191Auth(c, deadline)
	if err != nil {
		return fmt.Errorf("unable to authenticate relayer: %w", err)
	}

	var has bool

	if has, err = bridge.HasRole(ops, relayer_role, addressRecovered); err != nil {
		err = fmt.Errorf("unable to check relayer role on-chain")
		c.JSON(400, gin.H{"msg": err})
		return err
	}

	if !has {
		err = fmt.Errorf("q.Relayer not an on-chain relayer")
		c.JSON(400, gin.H{"msg": err})
		return err
	}

	fmt.Printf("HAS:" + fmt.Sprintf("%t", has))
	return nil
}

// AuthMiddleware is a placeholder for Gin authentication middleware.
func (r *APIServer) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement the actual authentication logic here.

		// For now, just log a message or pass through.
		fmt.Println("AuthMiddleware: Logic to be implemented")
		err := r.Authenticate(c)
		if err != nil {
			fmt.Println("Auth Error:", err)
		}
		// Pass on to the next-in-chain
		c.Next()
	}
}
