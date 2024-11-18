// Package rest provides RESTful API services for RFQ
package rest

import (
	"context"
	"encoding/json"
	"math/big"

	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ipfs/go-log"
	"github.com/puzpuzpuz/xsync"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
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
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridgev2"
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
	cfg                   config.Config
	db                    db.APIDB
	engine                *gin.Engine
	upgrader              websocket.Upgrader
	omnirpcClient         omniClient.RPCClient
	handler               metrics.Handler
	meter                 metric.Meter
	fastBridgeContractsV1 map[uint32]*fastbridge.FastBridge
	fastBridgeContractsV2 map[uint32]*fastbridgev2.FastBridgeV2
	roleCacheV1           map[uint32]*ttlcache.Cache[string, bool]
	roleCacheV2           map[uint32]*ttlcache.Cache[string, bool]
	// relayAckCache contains a set of transactionID values that reflect
	// transactions that have been acked for relay
	relayAckCache *ttlcache.Cache[string, string]
	// ackMux is a mutex used to ensure that only one transaction id can be acked at a time.
	ackMux sync.Mutex
	// latestQuoteAgeGauge is a gauge that records the age of the latest quote.
	latestQuoteAgeGauge metric.Float64ObservableGauge
	// wsClients maintains a mapping of connection ID to a channel for sending quote requests.
	wsClients     *xsync.MapOf[string, WsClient]
	pubSubManager PubSubManager
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

	fastBridgeContractsV1 := make(map[uint32]*fastbridge.FastBridge)
	rolesV1 := make(map[uint32]*ttlcache.Cache[string, bool])
	for chainID, contract := range cfg.FastBridgeContractsV1 {
		chainClient, err := omniRPCClient.GetChainClient(ctx, int(chainID))
		if err != nil {
			return nil, fmt.Errorf("could not create omnirpc client: %w", err)
		}
		fastBridgeContractsV1[chainID], err = fastbridge.NewFastBridge(common.HexToAddress(contract), chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not create bridge contract: %w", err)
		}

		// create the roles cache
		rolesV1[chainID] = ttlcache.New[string, bool](
			ttlcache.WithTTL[string, bool](cacheInterval),
		)
		roleCache := rolesV1[chainID]
		go roleCache.Start()
		go func() {
			<-ctx.Done()
			roleCache.Stop()
		}()
	}

	fastBridgeContractsV2 := make(map[uint32]*fastbridgev2.FastBridgeV2)
	rolesV2 := make(map[uint32]*ttlcache.Cache[string, bool])
	for chainID, contract := range cfg.FastBridgeContractsV2 {
		chainClient, err := omniRPCClient.GetChainClient(ctx, int(chainID))
		if err != nil {
			return nil, fmt.Errorf("could not create omnirpc client: %w", err)
		}
		fastBridgeContractsV2[chainID], err = fastbridgev2.NewFastBridgeV2(common.HexToAddress(contract), chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not create bridge contract: %w", err)
		}

		// create the roles cache
		rolesV2[chainID] = ttlcache.New[string, bool](
			ttlcache.WithTTL[string, bool](cacheInterval),
		)
		roleCache := rolesV2[chainID]
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
		cfg:                   cfg,
		db:                    store,
		omnirpcClient:         omniRPCClient,
		handler:               handler,
		meter:                 handler.Meter(meterName),
		fastBridgeContractsV1: fastBridgeContractsV1,
		fastBridgeContractsV2: fastBridgeContractsV2,
		roleCacheV1:           rolesV1,
		roleCacheV2:           rolesV2,
		relayAckCache:         relayAckCache,
		ackMux:                sync.Mutex{},
		wsClients:             xsync.NewMapOf[WsClient](),
		pubSubManager:         NewPubSubManager(),
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
	// ContractsRoute is the API endpoint for returning a list of contracts.
	ContractsRoute = "/contracts"
	// RFQStreamRoute is the API endpoint for handling active quote requests via websocket.
	RFQStreamRoute = "/rfq_stream"
	// RFQRoute is the API endpoint for handling RFQ requests.
	RFQRoute = "/rfq"
	// ChainsHeader is the header for specifying chains during a websocket handshake.
	ChainsHeader = "Chains"
	// AuthorizationHeader is the header for specifying the authorization.
	AuthorizationHeader = "Authorization"
	cacheInterval       = time.Minute
)

var logger = log.Logger("rfq-api")

// Run runs the quoter api server.
func (r *QuoterAPIServer) Run(ctx context.Context) error {
	engine := ginhelper.New(logger)
	h := NewHandler(r.db, r.cfg)

	versionNumber, versionNumErr := getCurrentVersion()
	if versionNumErr != nil {
		return fmt.Errorf("could not get current API version: %w", versionNumErr)
	}
	engine.Use(APIVersionMiddleware(versionNumber))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Authenticated routes
	quotesPut := engine.Group(QuoteRoute)
	quotesPut.Use(r.AuthMiddleware())
	quotesPut.PUT("", h.ModifyQuote)
	bulkQuotesPut := engine.Group(BulkQuotesRoute)
	bulkQuotesPut.Use(r.AuthMiddleware())
	bulkQuotesPut.PUT("", h.ModifyBulkQuotes)
	ackPut := engine.Group(AckRoute)
	ackPut.Use(r.AuthMiddleware())
	ackPut.PUT("", r.PutRelayAck)
	openQuoteRequestsGet := engine.Group(RFQRoute)
	openQuoteRequestsGet.Use(r.AuthMiddleware())
	openQuoteRequestsGet.GET("", h.GetOpenQuoteRequests)

	// WebSocket route
	wsRoute := engine.Group(RFQStreamRoute)
	wsRoute.Use(r.AuthMiddleware())
	wsRoute.GET("", func(c *gin.Context) {
		r.GetActiveRFQWebsocket(ctx, c)
	})

	// Unauthenticated routes
	engine.GET(QuoteRoute, h.GetQuotes)
	engine.GET(ContractsRoute, h.GetContracts)
	engine.PUT(RFQRoute, r.PutRFQRequest)

	// WebSocket upgrader
	r.upgrader = websocket.Upgrader{
		CheckOrigin: func(_ *http.Request) bool {
			return true // TODO: Implement a more secure check
		},
	}

	r.engine = engine

	// Start the main HTTP server
	connection := baseServer.Server{}
	fmt.Printf("starting api at http://localhost:%s\n", r.cfg.Port)

	err := connection.ListenAndServe(ctx, fmt.Sprintf(":%s", r.cfg.Port), r.engine)
	if err != nil {
		return fmt.Errorf("could not start rest api server: %w", err)
	}

	return nil
}

// AuthMiddleware is the Gin authentication middleware that authenticates requests using EIP191.
//
//nolint:gosec
func (r *QuoterAPIServer) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loggedRequest interface{}
		var err error
		destChainIDs := []uint32{}

		// Parse the dest chain id from the request
		switch c.Request.URL.Path {
		case QuoteRoute:
			var req model.PutRelayerQuoteRequest
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
		case RFQRoute, RFQStreamRoute:
			chainsHeader := c.GetHeader(ChainsHeader)
			if chainsHeader != "" {
				var chainIDs []int
				err = json.Unmarshal([]byte(chainsHeader), &chainIDs)
				if err == nil {
					for _, chainID := range chainIDs {
						destChainIDs = append(destChainIDs, uint32(chainID))
					}
				}
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
			addr, err := r.checkRoleParallel(c, destChainID)
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

type roleContract interface {
	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)
}

func (r *QuoterAPIServer) checkRoleParallel(c *gin.Context, destChainID uint32) (addressRecovered common.Address, err error) {
	g := new(errgroup.Group)
	var v1Addr, v2Addr common.Address
	var v1Err, v2Err error

	g.Go(func() error {
		v1Addr, v1Err = r.checkRole(c, destChainID, true)
		return v1Err
	})

	g.Go(func() error {
		v2Addr, v2Err = r.checkRole(c, destChainID, false)
		return v2Err
	})

	err = g.Wait()
	if v1Addr != (common.Address{}) {
		return v1Addr, nil
	}
	if v2Addr != (common.Address{}) {
		return v2Addr, nil
	}
	if err != nil {
		return common.Address{}, fmt.Errorf("role check failed: %w", err)
	}

	return common.Address{}, fmt.Errorf("role check failed for both v1 and v2")
}

func (r *QuoterAPIServer) checkRole(c *gin.Context, destChainID uint32, useV1 bool) (addressRecovered common.Address, err error) {
	var bridge roleContract
	var roleCache *ttlcache.Cache[string, bool]
	var ok bool
	if useV1 {
		bridge, ok = r.fastBridgeContractsV1[destChainID]
		if !ok {
			err = fmt.Errorf("dest chain id not supported: %d", destChainID)
			return addressRecovered, err
		}
		roleCache = r.roleCacheV1[destChainID]
	} else {
		bridge, ok = r.fastBridgeContractsV2[destChainID]
		if !ok {
			err = fmt.Errorf("dest chain id not supported: %d", destChainID)
			return addressRecovered, err
		}
		roleCache = r.roleCacheV2[destChainID]
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
	cachedRoleItem := roleCache.Get(addressRecovered.Hex())
	var hasRole bool

	if cachedRoleItem == nil || cachedRoleItem.IsExpired() {
		// Cache miss or expired, check on-chain
		hasRole, err = bridge.HasRole(ops, relayerRole, addressRecovered)
		if err != nil {
			return addressRecovered, fmt.Errorf("unable to check relayer role on-chain: %w", err)
		}
		// Update cache
		roleCache.Set(addressRecovered.Hex(), hasRole, cacheInterval)
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
// @Param request body model.PutRelayerQuoteRequest true "query params"
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

// GetActiveRFQWebsocket handles the WebSocket connection for active quote requests.
// GET /rfq_stream.
// @Summary Listen for Active RFQs
// @Schemes
// @Description Establish a WebSocket connection to listen for streaming active quote requests.
// @Tags quotes
// @Produce json
// @Success 101 {string} string "Switching Protocols"
// @Header 101 {string} X-Api-Version "API Version Number - See docs for more info"
// @Router /rfq_stream [get].
func (r *QuoterAPIServer) GetActiveRFQWebsocket(ctx context.Context, c *gin.Context) {
	ctx, span := r.handler.Tracer().Start(ctx, "GetActiveRFQWebsocket")
	defer func() {
		metrics.EndSpan(span)
	}()

	ws, err := r.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("Failed to set websocket upgrade", "error", err)
		return
	}

	// use the relayer address as the ID for the connection
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

	span.SetAttributes(
		attribute.String("relayer_address", relayerAddr),
	)

	// only one connection per relayer allowed
	_, ok = r.wsClients.Load(relayerAddr)
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "relayer already connected"})
		return
	}

	defer func() {
		// cleanup ws registry
		r.wsClients.Delete(relayerAddr)
	}()

	client := newWsClient(relayerAddr, ws, r.pubSubManager, r.handler)
	r.wsClients.Store(relayerAddr, client)
	span.AddEvent("registered ws client")
	err = client.Run(ctx)
	if err != nil {
		logger.Error("Error running websocket client", "error", err)
	}
}

const (
	quoteTypeActive  = "active"
	quoteTypePassive = "passive"
)

// PutRFQRequest handles a user request for a quote.
// PUT /rfq.
// @Summary Initiate an Active RFQ
// @Schemes
// @Description Initiate an Active Request-For-Quote and return the best quote available.
// @Param request body model.PutRFQRequest true "Initiate an Active Request-For-Quote"
// @Tags quotes
// @Accept json
// @Produce json
// @Success 200 {object} model.PutRFQResponse
// @Header 200 {string} X-Api-Version "API Version Number - See docs for more info"
// @Router /rfq [put].
//
//nolint:cyclop
func (r *QuoterAPIServer) PutRFQRequest(c *gin.Context) {
	var req model.PutRFQRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestID := uuid.New().String()
	ctx, span := r.handler.Tracer().Start(c.Request.Context(), "PutRFQRequest", trace.WithAttributes(
		attribute.String("request_id", requestID),
	))
	defer func() {
		metrics.EndSpan(span)
	}()

	err = r.db.InsertActiveQuoteRequest(ctx, &req, requestID)
	if err != nil {
		logger.Warnf("Error inserting active quote request: %w", err)
	}

	var isActiveRFQ bool
	for _, quoteType := range req.QuoteTypes {
		if quoteType == quoteTypeActive {
			isActiveRFQ = true
			break
		}
	}
	span.SetAttributes(attribute.Bool("is_active_rfq", isActiveRFQ))

	// if specified, fetch the active quote. always consider passive quotes
	var activeQuote *model.QuoteData
	if isActiveRFQ {
		activeQuote = r.handleActiveRFQ(ctx, &req, requestID)
		if activeQuote != nil && activeQuote.DestAmount != nil {
			span.SetAttributes(attribute.String("active_quote_dest_amount", *activeQuote.DestAmount))
		}
	}
	passiveQuote, err := r.handlePassiveRFQ(ctx, &req)
	if err != nil {
		logger.Error("Error handling passive RFQ", "error", err)
	}
	if passiveQuote != nil && passiveQuote.DestAmount != nil {
		span.SetAttributes(attribute.String("passive_quote_dest_amount", *passiveQuote.DestAmount))
	}
	quote := getBestQuote(activeQuote, passiveQuote)
	var quoteType string
	if quote == activeQuote {
		quoteType = quoteTypeActive
	} else if quote == passiveQuote {
		quoteType = quoteTypePassive
	}

	// build and return the response
	resp := getQuoteResponse(ctx, quote, quoteType)
	c.JSON(http.StatusOK, resp)
}

func getQuoteResponse(ctx context.Context, quote *model.QuoteData, quoteType string) (resp model.PutRFQResponse) {
	span := trace.SpanFromContext(ctx)

	destAmount := big.NewInt(0)
	if quote != nil && quote.DestAmount != nil {
		amt, ok := destAmount.SetString(*quote.DestAmount, 10)
		if ok {
			destAmount = amt
		}
	}
	if destAmount.Sign() <= 0 {
		span.AddEvent("no quotes found")
		resp = model.PutRFQResponse{
			Success: false,
			Reason:  "no quotes found",
		}
	} else {
		span.SetAttributes(
			attribute.String("quote_type", quoteType),
			attribute.String("quote_dest_amount", *quote.DestAmount),
		)
		resp = model.PutRFQResponse{
			Success:        true,
			QuoteType:      quoteType,
			QuoteID:        quote.QuoteID,
			DestAmount:     *quote.DestAmount,
			RelayerAddress: *quote.RelayerAddress,
		}
	}

	return resp
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
