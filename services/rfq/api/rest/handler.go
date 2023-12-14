package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
)

type Handler struct {
	db db.ApiDB
}

func NewHandler(db db.ApiDB) *Handler {
	return &Handler{
		db: db, // Store the database connection in the handler
	}
}

// PUT /quotes
// @dev Protected Method: Authentication is handled through middleware in server.go.
func (h *Handler) ModifyQuote(c *gin.Context) {
	// Retrieve the request from context
	req, exists := c.Get("putRequest")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not found"})
		return
	}
	putRequest, ok := req.(*PutRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request type"})
		return
	}

	destChainID, err := strconv.ParseUint(putRequest.DestChainID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DestChainID"})
		return
	}
	destAmount, err := decimal.NewFromString(putRequest.DestAmount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DestAmount"})
		return
	}
	price, err := decimal.NewFromString(putRequest.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Price"})
		return
	}
	quote := &db.Quote{
		ID:            uint64(putRequest.ID),
		DestChainID:   destChainID,
		DestTokenAddr: putRequest.DestTokenAddr,
		DestAmount:    destAmount,
		Price:         price,
	}
	err = h.db.UpsertQuote(quote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// GET /quotes.
func (h *Handler) GetQuotes(c *gin.Context) {
	destChainIdStr := c.Query("destChainId")
	destTokenAddr := c.Query("destTokenAddr")

	if destChainIdStr != "" && destTokenAddr != "" {
		destChainId, err := strconv.ParseUint(destChainIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid destChainId"})
			return
		}

		quotes, err := h.db.GetQuotesByDestChainAndToken(destChainId, destTokenAddr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, quotes)
	} else {
		// Pseudocode for retrieving all quotes from the database
		quotes, err := h.db.GetAllQuotes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, quotes)
	}
	c.Status(http.StatusOK)
	// Implement logic to fetch and return quotes
}

// GET /quotes?destChainId=&destTokenAddr=&destAmount=.
func (h *Handler) GetFilteredQuotes(c *gin.Context) {
	// Implement logic to fetch and return filtered quotes
}
