package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/api/config"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

// Handler is the REST API handler.
type Handler struct {
	db  db.APIDB
	cfg config.Config
}

// NewHandler creates a new REST API handler.
func NewHandler(db db.APIDB, cfg config.Config) *Handler {
	return &Handler{
		db:  db, // Store the database connection in the handler
		cfg: cfg,
	}
}

// APIVersionMiddleware adds the X-API-Version header to the response with the current version # from versions.json file.
func APIVersionMiddleware(serverVersion string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Api-Version", serverVersion)
		c.Next()
	}
}

// ModifyQuote upserts a quote
//
// PUT /quotes
// @dev Protected Method: Authentication is handled through middleware in server.go.
// nolint: cyclop
// @Summary Upsert quote
// @Schemes
// @Description upsert a quote from relayer.
// @Param request body model.PutRelayerQuoteRequest true "query params"
// @Tags quotes
// @Accept json
// @Produce json
// @Success 200
// @Header 200 {string} X-Api-Version "API Version Number - See docs for more info"
// @Router /quotes [put].
func (h *Handler) ModifyQuote(c *gin.Context) {
	// Retrieve the request from context
	req, exists := c.Get("putRequest")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not found"})
		return
	}
	relayerAddr, exists := c.Get("relayerAddr")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No relayer address recovered from signature"})
		return
	}
	putRequest, ok := req.(*model.PutRelayerQuoteRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request type"})
		return
	}

	dbQuote, err := parseDBQuote(h.cfg, *putRequest, relayerAddr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.db.UpsertQuote(c, dbQuote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// ModifyBulkQuotes upserts multiple quotes
//
// PUT /bulk_quotes
// @dev Protected Method: Authentication is handled through middleware in server.go.
// nolint: cyclop
// @Summary Upsert quotes
// @Schemes
// @Description upsert bulk quotes from relayer.
// @Param request body model.PutBulkQuotesRequest true "query params"
// @Tags quotes
// @Accept json
// @Produce json
// @Success 200
// @Header 200 {string} X-Api-Version "API Version Number - See docs for more info"
// @Router /bulk_quotes [put].
func (h *Handler) ModifyBulkQuotes(c *gin.Context) {
	// Retrieve the request from context
	req, exists := c.Get("putRequest")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not found"})
		return
	}
	relayerAddr, exists := c.Get("relayerAddr")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No relayer address recovered from signature"})
		return
	}
	putRequest, ok := req.(*model.PutBulkQuotesRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request type"})
		return
	}

	dbQuotes := []*db.Quote{}
	for _, quoteReq := range putRequest.Quotes {
		dbQuote, err := parseDBQuote(h.cfg, quoteReq, relayerAddr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quote request"})
			return
		}
		dbQuotes = append(dbQuotes, dbQuote)
	}

	err := h.db.UpsertQuotes(c, dbQuotes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

//nolint:gosec
func parseDBQuote(cfg config.Config, putRequest model.PutRelayerQuoteRequest, relayerAddr interface{}) (*db.Quote, error) {
	destAmount, err := decimal.NewFromString(putRequest.DestAmount)
	if err != nil {
		return nil, fmt.Errorf("invalid DestAmount")
	}
	maxOriginAmount, err := decimal.NewFromString(putRequest.MaxOriginAmount)
	if err != nil {
		return nil, fmt.Errorf("invalid MaxOriginAmount")
	}
	fixedFee, err := decimal.NewFromString(putRequest.FixedFee)
	if err != nil {
		return nil, fmt.Errorf("invalid FixedFee")
	}

	err = validateFastBridgeAddresses(cfg, putRequest)
	if err != nil {
		return nil, fmt.Errorf("invalid fast bridge addresses: %w", err)
	}

	// nolint: forcetypeassert
	return &db.Quote{
		OriginChainID:   uint64(putRequest.OriginChainID),
		OriginTokenAddr: putRequest.OriginTokenAddr,
		DestChainID:     uint64(putRequest.DestChainID),
		DestTokenAddr:   putRequest.DestTokenAddr,
		DestAmount:      destAmount,
		MaxOriginAmount: maxOriginAmount,
		FixedFee:        fixedFee,
		//nolint: forcetypeassert
		RelayerAddr:             relayerAddr.(string),
		OriginFastBridgeAddress: putRequest.OriginFastBridgeAddress,
		DestFastBridgeAddress:   putRequest.DestFastBridgeAddress,
	}, nil
}

func validateFastBridgeAddresses(cfg config.Config, putRequest model.PutRelayerQuoteRequest) error {
	// Check V1 contracts
	isV1Origin := common.HexToAddress(cfg.FastBridgeContractsV1[uint32(putRequest.OriginChainID)]) == common.HexToAddress(putRequest.OriginFastBridgeAddress)
	isV1Dest := common.HexToAddress(cfg.FastBridgeContractsV1[uint32(putRequest.DestChainID)]) == common.HexToAddress(putRequest.DestFastBridgeAddress)

	// Check V2 contracts
	isV2Origin := common.HexToAddress(cfg.FastBridgeContractsV2[uint32(putRequest.OriginChainID)]) == common.HexToAddress(putRequest.OriginFastBridgeAddress)
	isV2Dest := common.HexToAddress(cfg.FastBridgeContractsV2[uint32(putRequest.DestChainID)]) == common.HexToAddress(putRequest.DestFastBridgeAddress)

	// Valid if both addresses match either V1 or V2
	if (isV1Origin && isV1Dest) || (isV2Origin && isV2Dest) {
		return nil
	}

	return fmt.Errorf("origin and destination fast bridge addresses must match either V1 or V2")
}

//nolint:gosec
func quoteResponseFromDBQuote(dbQuote *db.Quote) *model.GetQuoteResponse {
	return &model.GetQuoteResponse{
		OriginChainID:           int(dbQuote.OriginChainID),
		OriginTokenAddr:         dbQuote.OriginTokenAddr,
		DestChainID:             int(dbQuote.DestChainID),
		DestTokenAddr:           dbQuote.DestTokenAddr,
		DestAmount:              dbQuote.DestAmount.String(),
		MaxOriginAmount:         dbQuote.MaxOriginAmount.String(),
		FixedFee:                dbQuote.FixedFee.String(),
		RelayerAddr:             dbQuote.RelayerAddr,
		OriginFastBridgeAddress: dbQuote.OriginFastBridgeAddress,
		DestFastBridgeAddress:   dbQuote.DestFastBridgeAddress,
		UpdatedAt:               dbQuote.UpdatedAt.Format(time.RFC3339),
	}
}

// GetQuotes retrieves all quotes from the database.
// GET /quotes.
// nolint: cyclop
// PingExample godoc
// @Summary Get quotes
// @Schemes
// @Param   originChainID     path    int     false        "origin chain id to filter quotes by"
// @Param   originTokenAddr   path    string     false        "origin chain id to filter quotes by"
// @Param   destChainID     path    int     false        "destination chain id to filter quotes by"
// @Param   destTokenAddr   path    string     false        "destination token address to filter quotes by"
// @Param   relayerAddr   path    string     false        "relayer address to filter quotes by"
// @Description get quotes from all relayers.
// @Tags quotes
// @Accept json
// @Produce json
// @Success 200 {array} model.GetQuoteResponse
// @Header 200 {string} X-Api-Version "API Version Number - See docs for more info"
// @Router /quotes [get].
func (h *Handler) GetQuotes(c *gin.Context) {
	originChainIDStr := c.Query("originChainID")
	originTokenAddr := c.Query("originTokenAddr")
	destChainIDStr := c.Query("destChainId")
	destTokenAddr := c.Query("destTokenAddr")
	relayerAddr := c.Query("relayerAddr")

	// TODO (aureliusbtc): rewrite this if
	//nolint: gocritic, nestif
	var dbQuotes []*db.Quote
	var err error
	if originChainIDStr != "" && originTokenAddr != "" && destChainIDStr != "" && destTokenAddr != "" {
		destChainID, err := strconv.ParseUint(destChainIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid destChainId"})
			return
		}

		originChainID, err := strconv.ParseUint(originChainIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid originChainID"})
			return
		}

		dbQuotes, err = h.db.GetQuotesByOriginAndDestination(c, originChainID, originTokenAddr, destChainID, destTokenAddr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else if relayerAddr != "" {
		dbQuotes, err = h.db.GetQuotesByRelayerAddress(c, relayerAddr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		dbQuotes, err = h.db.GetAllQuotes(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Filter quotes
	dbQuotes = filterQuoteAge(h.cfg, dbQuotes)

	// Convert quotes from db model to api model
	quotes := make([]*model.GetQuoteResponse, len(dbQuotes))
	for i, dbQuote := range dbQuotes {
		quotes[i] = quoteResponseFromDBQuote(dbQuote)
	}
	c.JSON(http.StatusOK, quotes)
}

// GetOpenQuoteRequests retrieves all open quote requests.
// GET /open_quote_requests
// @Summary Get open quote requests
// @Description Get all open quote requests that are currently in Received or Pending status.
// @Tags quotes
// @Accept json
// @Produce json
// @Success 200 {array} model.GetOpenQuoteRequestsResponse
// @Header 200 {string} X-Api-Version "API Version Number - See docs for more info"
// @Router /open_quote_requests [get].
func (h *Handler) GetOpenQuoteRequests(c *gin.Context) {
	dbQuotes, err := h.db.GetActiveQuoteRequests(c, db.Received, db.Pending)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	quotes := make([]*model.GetOpenQuoteRequestsResponse, len(dbQuotes))
	for i, dbQuote := range dbQuotes {
		quotes[i] = dbActiveQuoteRequestToModel(dbQuote)
	}
	c.JSON(http.StatusOK, quotes)
}

func dbActiveQuoteRequestToModel(dbQuote *db.ActiveQuoteRequest) *model.GetOpenQuoteRequestsResponse {
	return &model.GetOpenQuoteRequestsResponse{
		UserAddress:       dbQuote.UserAddress,
		OriginChainID:     dbQuote.OriginChainID,
		OriginTokenAddr:   dbQuote.OriginTokenAddr,
		DestChainID:       dbQuote.DestChainID,
		DestTokenAddr:     dbQuote.DestTokenAddr,
		OriginAmountExact: dbQuote.OriginAmountExact.String(),
		ExpirationWindow:  int(dbQuote.ExpirationWindow.Milliseconds()),
		CreatedAt:         dbQuote.CreatedAt,
	}
}

// GetContracts retrieves all contracts api is currently enabled on.
// GET /contracts.
// PingExample godoc
// @Summary Get contract addresses
// @Description get quotes from all relayers.
// @Tags quotes
// @Accept json
// @Produce json
// @Success 200 {array} model.GetContractsResponse
// @Header 200 {string} X-Api-Version "API Version Number - See docs for more info"
// @Router /contracts [get].
func (h *Handler) GetContracts(c *gin.Context) {
	c.JSON(http.StatusOK, model.GetContractsResponse{
		ContractsV1: h.cfg.FastBridgeContractsV1,
		ContractsV2: h.cfg.FastBridgeContractsV2,
	})
}

func filterQuoteAge(cfg config.Config, dbQuotes []*db.Quote) []*db.Quote {
	maxAge := cfg.GetMaxQuoteAge()

	thresh := time.Now().Add(-maxAge)
	var filteredQuotes []*db.Quote
	for _, quote := range dbQuotes {
		if quote.UpdatedAt.After(thresh) {
			filteredQuotes = append(filteredQuotes, quote)
		}
	}

	return filteredQuotes
}
