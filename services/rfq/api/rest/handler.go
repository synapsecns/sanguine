package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

// Handler is the REST API handler.
type Handler struct {
	db db.APIDB
}

// NewHandler creates a new REST API handler.
func NewHandler(db db.APIDB) *Handler {
	return &Handler{
		db: db, // Store the database connection in the handler
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

	destAmount, err := decimal.NewFromString(putRequest.DestAmount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DestAmount"})
		return
	}
	maxOriginAmount, err := decimal.NewFromString(putRequest.MaxOriginAmount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DestAmount"})
		return
	}
	fixedFee, err := decimal.NewFromString(putRequest.FixedFee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid FixedFee"})
		return
	}
	// nolint: forcetypeassert
	quote := &db.Quote{
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
	}
	err = h.db.UpsertQuote(c, quote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
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

	// Convert quotes from db model to api model
	quotes := make([]*model.GetQuoteResponse, len(dbQuotes))
	for i, dbQuote := range dbQuotes {
		quotes[i] = model.QuoteResponseFromDbQuote(dbQuote)
	}
	c.JSON(http.StatusOK, quotes)
}
