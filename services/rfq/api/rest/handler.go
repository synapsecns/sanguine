package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
// @Param request body model.PutQuoteRequest true "query params"
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
	putRequest, ok := req.(*model.PutQuoteRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request type"})
		return
	}

	dbQuote, err := parseDBQuote(*putRequest, relayerAddr)
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
		dbQuote, err := parseDBQuote(quoteReq, relayerAddr)
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

func parseDBQuote(putRequest model.PutQuoteRequest, relayerAddr interface{}) (*db.Quote, error) {
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
		quotes[i] = model.QuoteResponseFromDbQuote(dbQuote)
	}
	c.JSON(http.StatusOK, quotes)
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
	// Convert quotes from db model to api model
	contracts := make(map[uint32]string)
	for chainID, address := range h.cfg.Bridges {
		contracts[chainID] = address
	}
	c.JSON(http.StatusOK, model.GetContractsResponse{Contracts: contracts})
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
