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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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

const meterName = "github.com/synapsecns/sanguine/services/rfq/api/rest"

func getCurrentVersion() (string, error) {
	if len(APIversions.Versions) == 0 {
		return "", fmt.Errorf("no versions found")
	}

	return APIversions.Versions[0].Version, nil
}

// QuoterAPIServer is a struct that holds the configuration, database connection, gin engine, RPC client, metrics handler, and fast bridge contracts.
// It is used to initialize and run the API server.
type QuoterAPIServer struct {
	cfg                 config.Config
	db                  db.APIDB
	engine              *gin.Engine
	upgrader            websocket.Upgrader
	omnirpcClient       omniClient.RPCClient
	handler             metrics.Handler
	meter               metric.Meter
	fastBridgeContracts map[uint32]*fastbridge.FastBridge
	roleCache           map[uint32]*ttlcache.Cache[string, bool]
	// relayAckCache contains a set of transactionID values that reflect
	// transactions that have been acked for relay
	relayAckCache *ttlcache.Cache[string, string]
	// ackMux is a mutex used to ensure that only one transaction id can be acked at a time.
	ackMux sync.Mutex
	// latestQuoteAgeGauge is a gauge that records the age of the latest quote
	latestQuoteAgeGauge metric.Float64ObservableGauge
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

	q := &QuoterAPIServer{
		cfg:                 cfg,
		db:                  store,
		omnirpcClient:       omniRPCClient,
		handler:             handler,
		meter:               handler.Meter(meterName),
		fastBridgeContracts: bridges,
		roleCache:           roles,
		relayAckCache:       relayAckCache,
		ackMux:              sync.Mutex{},
	}

	// Prometheus metrics setup
	var err error
	q.latestQuoteAgeGauge, err = q.meter.Float64ObservableGauge("latest_quote_age")
	if err != nil {
		return nil, fmt.Errorf("could not create latest quote age gauge: %w", err)
	}

	_, err = q.meter.RegisterCallback(q.recordLatestQuoteAge, q.latestQuoteAgeGauge)
	if err != nil {
		return nil, fmt.Errorf("could not register callback: %w", err)
	}

	return q, nil
}

const (
	// QuoteRoute is the API endpoint for handling quote related requests.
	QuoteRoute = "/quotes"
	// BulkQuotesRoute is the API endpoint for handling bulk quote related requests.
	BulkQuotesRoute = "/bulk_quotes"
	// AckRoute is the API endpoint for handling relay ack related requests.
	AckRoute = "/ack"
	// ContractsRoute is the API endpoint for returning a list fo contracts.
	ContractsRoute = "/contracts"
	// QuoteRequestsRoute is the API endpoint for handling active quote requests via websocket.
	QuoteRequestsRoute = "/quote_requests"
	cacheInterval      = time.Minute
)

var logger = log.Logger("rfq-api")

// Run runs the quoter api server.
func (r *QuoterAPIServer) Run(ctx context.Context) error {
	// TODO: Use Gin Helper
	engine := ginhelper.New(logger)
	h := NewHandler(r.db, r.cfg)

	versionNumber, versionNumErr := getCurrentVersion()
	if versionNumErr != nil {
		return fmt.Errorf("could not get current API version: %w", versionNumErr)
	}
	engine.Use(APIVersionMiddleware(versionNumber))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Apply AuthMiddleware only to the PUT routes
	quotesPut := engine.Group(QuoteRoute)
	quotesPut.Use(r.AuthMiddleware())
	quotesPut.PUT("", h.ModifyQuote)
	bulkQuotesPut := engine.Group(BulkQuotesRoute)
	bulkQuotesPut.Use(r.AuthMiddleware())
	bulkQuotesPut.PUT("", h.ModifyBulkQuotes)
	ackPut := engine.Group(AckRoute)
	ackPut.Use(r.AuthMiddleware())
	ackPut.PUT("", r.PutRelayAck)
	activeRFQGet := engine.Group(QuoteRequestsRoute)
	activeRFQGet.Use(r.AuthMiddleware())
	activeRFQGet.GET("", func(c *gin.Context) {
		r.GetActiveRFQWebsocket(ctx, c)
	})

	// GET routes without the AuthMiddleware
	// engine.PUT("/quotes", h.ModifyQuote)
	engine.GET(QuoteRoute, h.GetQuotes)

	engine.GET(ContractsRoute, h.GetContracts)

	// WebSocket upgrader
	r.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // TODO: Implement a more secure check
		},
	}

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
		var err error
		destChainIDs := []uint32{}

		// Parse the dest chain id from the request
		switch c.Request.URL.Path {
		case QuoteRoute:
			var req model.PutQuoteRequest
			err = c.BindJSON(&req)
			if err == nil {
				destChainIDs = append(destChainIDs, uint32(req.DestChainID))
				loggedRequest = &req
			}
		case BulkQuotesRoute:
			var req model.PutBulkQuotesRequest
			err = c.BindJSON(&req)
			if err == nil {
				for _, quote := range req.Quotes {
					destChainIDs = append(destChainIDs, uint32(quote.DestChainID))
				}
				loggedRequest = &req
			}
		case AckRoute:
			var req model.PutAckRequest
			err = c.BindJSON(&req)
			if err == nil {
				destChainIDs = append(destChainIDs, uint32(req.DestChainID))
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
		var addressRecovered *common.Address
		for _, destChainID := range destChainIDs {
			addr, err := r.checkRole(c, destChainID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
				c.Abort()
				return
			}
			if addressRecovered == nil {
				addressRecovered = &addr
			} else if *addressRecovered != addr {
				c.JSON(http.StatusBadRequest, gin.H{"msg": "relayer address mismatch"})
				c.Abort()
				return
			}
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

	// Check and update cache
	cachedRoleItem := r.roleCache[destChainID].Get(addressRecovered.Hex())
	var hasRole bool

	if cachedRoleItem == nil || cachedRoleItem.IsExpired() {
		// Cache miss or expired, check on-chain
		hasRole, err = bridge.HasRole(ops, relayerRole, addressRecovered)
		if err != nil {
			return addressRecovered, fmt.Errorf("unable to check relayer role on-chain: %w", err)
		}
		// Update cache
		r.roleCache[destChainID].Set(addressRecovered.Hex(), hasRole, cacheInterval)
	} else {
		// Use cached value
		hasRole = cachedRoleItem.Value()
	}

	// Verify role
	if !hasRole {
		return addressRecovered, fmt.Errorf("relayer not an on-chain relayer")
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
// @Header 200 {string} X-Api-Version "API Version Number - See docs for more info"
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
	shouldRelay := ack == nil || common.HexToAddress(relayerAddr).Hex() == common.HexToAddress(ack.Value()).Hex()
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

func (r *QuoterAPIServer) recordLatestQuoteAge(ctx context.Context, observer metric.Observer) (err error) {
	if r.handler == nil || r.latestQuoteAgeGauge == nil {
		return nil
	}

	quotes, err := r.db.GetAllQuotes(ctx)
	if err != nil {
		return fmt.Errorf("could not get latest quote age: %w", err)
	}

	ageByRelayer := make(map[string]float64)
	for _, quote := range quotes {
		age := time.Since(quote.UpdatedAt).Seconds()
		prevAge, ok := ageByRelayer[quote.RelayerAddr]
		if !ok || age < prevAge {
			ageByRelayer[quote.RelayerAddr] = age
		}
	}

	for relayer, age := range ageByRelayer {
		opts := metric.WithAttributes(
			attribute.String("relayer", relayer),
		)
		observer.ObserveFloat64(r.latestQuoteAgeGauge, age, opts)
	}

	return nil
}

// GetActiveRFQWebsocket handles the WebSocket connection for active quote requests.
func (r *QuoterAPIServer) GetActiveRFQWebsocket(ctx context.Context, c *gin.Context) {
	ws, err := r.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("Failed to set websocket upgrade", "error", err)
		return
	}
	defer ws.Close()

	// pass the run context here in case the server is shutdown
	r.handleWebSocket(ctx, ws)
}

const (
	pingOp         = "ping"
	pongOp         = "pong"
	requestQuoteOp = "request_quote"
	sendQuoteOp    = "send_quote"
)

// Update handleWebSocket to accept the context
func (r *QuoterAPIServer) handleWebSocket(ctx context.Context, conn *websocket.Conn) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Read message from WebSocket
			_, message, err := conn.ReadMessage()
			if err != nil {
				logger.Error("Error reading WebSocket message", "error", err)
				return
			}

			// Process the message
			response, err := r.processQuoteRequest(message)
			if err != nil {
				logger.Error("Error processing quote request", "error", err)
				continue
			}

			// Send response back through WebSocket
			if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
				logger.Error("Error writing WebSocket message", "error", err)
				return
			}
		}
	}
}

func (r *QuoterAPIServer) processQuoteRequest(message []byte) ([]byte, error) {
	var wsMessage model.ActiveRFQMessage
	err := json.Unmarshal(message, &wsMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal WebSocket message: %w", err)
	}

	switch wsMessage.Op {
	case pingOp:
		return json.Marshal(model.ActiveRFQMessage{
			Op:      pongOp,
			Success: true,
		})

	case requestQuoteOp:
		var quoteRequest model.QuoteRequest
		err := json.Unmarshal(message, &quoteRequest)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal quote request: %w", err)
		}

		// Process the quote request and generate a response
		quoteResponse, err := r.generateQuoteResponse(quoteRequest)
		if err != nil {
			return json.Marshal(model.ActiveRFQMessage{
				Op:      sendQuoteOp,
				Content: err.Error(),
				Success: false,
			})
		}

		return json.Marshal(model.ActiveRFQMessage{
			Op:      sendQuoteOp,
			Content: quoteResponse,
			Success: true,
		})

	default:
		return json.Marshal(model.ActiveRFQMessage{
			Content: "Unknown operation",
			Success: false,
		})
	}
}

func (r *QuoterAPIServer) generateQuoteResponse(request model.QuoteRequest) (model.QuoteResponse, error) {
	// TODO: Implement actual quote generation logic
	// This is a placeholder implementation
	quoteResponse := model.QuoteResponse{
		RequestID: request.RequestID,
		QuoteID:   uuid.New().String(),
		Data: model.QuoteResponseData{
			OriginChainID:           request.Data.OriginChainID,
			DestChainID:             request.Data.DestChainID,
			OriginTokenAddr:         request.Data.OriginTokenAddr,
			DestTokenAddr:           request.Data.DestTokenAddr,
			MaxOriginAmount:         request.Data.MaxOriginAmount,
			DestAmount:              "0",                                          // TODO: Calculate actual destination amount
			FixedFee:                "0",                                          // TODO: Calculate actual fee
			RelayerAddress:          "0x1234567890123456789012345678901234567890", // TODO: Use actual relayer address
			OriginFastBridgeAddress: "0x0987654321098765432109876543210987654321", // TODO: Use actual origin fast bridge address
			DestFastBridgeAddress:   "0x5432109876543210987654321098765432109876", // TODO: Use actual destination fast bridge address
		},
		UpdatedAt: time.Now(),
	}

	return quoteResponse, nil
}
