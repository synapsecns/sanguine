package relapi

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

// Handler is the REST API handler.
type Handler struct {
	db reldb.Service
}

// NewHandler creates a new REST API handler.
func NewHandler(db reldb.Service) *Handler {
	return &Handler{
		db: db, // Store the database connection in the handler
	}
}

// GetQuoteRequestStatusByTxID gets the status of a quote request, given a tx id.
func (h *Handler) GetQuoteRequestStatusByTxID(c *gin.Context) {
	txIDStr := c.Query("txID")
	if txIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must specify txID"})
		return
	}

	txIDBytes, err := hexutil.Decode(txIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid txID"})
		return
	}
	var txID [32]byte
	copy(txID[:], txIDBytes)

	quoteRequest, err := h.db.GetQuoteRequestByID(c, txID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := GetQuoteRequestStatusResponse{
		Status: quoteRequest.Status.String(),
		TxID:   hexutil.Encode(txID[:]),
		TxHash: quoteRequest.DestTxHash.String(),
	}
	c.JSON(http.StatusOK, resp)
}

// PutTxRetry retries a transaction based on tx hash.
func (h *Handler) PutTxRetry(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
