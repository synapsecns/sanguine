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
	relayerAddr, exists := c.Get("relayerAddr")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No relayer address recovered from signature"})
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
	originChainID, err := strconv.ParseUint(putRequest.OriginChainID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DestChainID"})
		return
	}
	maxOriginAmount, err := decimal.NewFromString(putRequest.MaxOriginAmount)
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
		OriginChainID:   originChainID,
		OriginTokenAddr: putRequest.OriginTokenAddr,
		DestChainID:     destChainID,
		DestTokenAddr:   putRequest.DestTokenAddr,
		DestAmount:      destAmount,
		Price:           price,
		MaxOriginAmount: maxOriginAmount,
		RelayerAddr:     relayerAddr.(string),
	}
	err = h.db.UpsertQuote(c, quote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// GET /quotes.
func (h *Handler) GetQuotes(c *gin.Context) {
	originChainIDStr := c.Query("originChainID")
	originTokenAddr := c.Query("originTokenAddr")
	destChainIDStr := c.Query("destChainId")
	destTokenAddr := c.Query("destTokenAddr")
	relayerAddr := c.Query("relayerAddr")

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

		quotes, err := h.db.GetQuotesByOriginAndDestination(c, originChainID, originTokenAddr, destChainID, destTokenAddr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, quotes)
	} else if relayerAddr != "" {
		quotes, err := h.db.GetQuotesByRelayerAddress(c, relayerAddr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, quotes)
	} else {
		// Pseudocode for retrieving all quotes from the database
		quotes, err := h.db.GetAllQuotes(c)
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
