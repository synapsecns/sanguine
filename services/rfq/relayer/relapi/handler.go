package relapi

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

// Handler is the REST API handler.
type Handler struct {
	db     reldb.Service
	chains map[uint32]*chain.Chain
}

// NewHandler creates a new REST API handler.
func NewHandler(db reldb.Service, chains map[uint32]*chain.Chain) *Handler {
	return &Handler{
		db:     db, // Store the database connection in the handler
		chains: chains,
	}
}

// GetHealth returns a successful response to signify the API is up and running.
func (h *Handler) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

const unspecifiedTxHash = "Must specify 'hash' (corresponding to origin tx)"

// GetQuoteRequestStatusByTxHash gets the status of a quote request, given an origin tx hash.
func (h *Handler) GetQuoteRequestStatusByTxHash(c *gin.Context) {
	txHashStr := c.Query("hash")
	if txHashStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": unspecifiedTxHash})
		return
	}

	txHash := common.HexToHash(txHashStr)
	quoteRequest, err := h.db.GetQuoteRequestByOriginTxHash(c, txHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := GetQuoteRequestStatusResponse{
		Status:       quoteRequest.Status.String(),
		TxID:         hexutil.Encode(quoteRequest.TransactionID[:]),
		OriginTxHash: quoteRequest.OriginTxHash.String(),
		DestTxHash:   quoteRequest.DestTxHash.String(),
	}
	c.JSON(http.StatusOK, resp)
}

// GetQuoteRequestStatusByTxID gets the status of a quote request, given a tx id.
func (h *Handler) GetQuoteRequestStatusByTxID(c *gin.Context) {
	txIDStr := c.Query("id")
	if txIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must specify 'id'"})
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
		Status:       quoteRequest.Status.String(),
		TxID:         hexutil.Encode(quoteRequest.TransactionID[:]),
		OriginTxHash: quoteRequest.OriginTxHash.String(),
		DestTxHash:   quoteRequest.DestTxHash.String(),
	}
	c.JSON(http.StatusOK, resp)
}

// GetTxRetry retries a transaction based on tx hash.
func (h *Handler) GetTxRetry(c *gin.Context) {
	txHashStr := c.Query("hash")
	if txHashStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": unspecifiedTxHash})
		return
	}

	txHash := common.HexToHash(txHashStr)
	quoteRequest, err := h.db.GetQuoteRequestByOriginTxHash(c, txHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	chainID := quoteRequest.Transaction.DestChainId
	chain, ok := h.chains[chainID]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No contract found for chain: %d", chainID)})
		return
	}

	// `quoteRequest == nil` case should be handled by the db query above
	nonce, gasAmount, err := chain.SubmitRelay(c, *quoteRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("could not submit relay: %s", err.Error())})
		return
	}

	resp := GetTxRetryResponse{
		TxID:      hexutil.Encode(quoteRequest.TransactionID[:]),
		ChainID:   chainID,
		Nonce:     nonce,
		GasAmount: gasAmount.String(),
	}
	c.JSON(http.StatusOK, resp)
}
