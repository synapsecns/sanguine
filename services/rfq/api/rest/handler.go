package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// PUT /quotes
func (h *Handler) ModifyQuote(c *gin.Context) {
	c.Status(http.StatusOK)
	// var quote db.Quote
	// if err := c.BindJSON(&quote); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// err := db.ApiDB.UpsertQuote(&quote)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.Status(http.StatusOK)
}

// GET /quotes
func (h *Handler) GetQuotes(c *gin.Context) {
	c.Status(http.StatusOK)
	// Implement logic to fetch and return quotes
}

// GET /quotes?destChainId=&destTokenAddr=&destAmount=
func (h *Handler) GetFilteredQuotes(c *gin.Context) {
	// Implement logic to fetch and return filtered quotes
}
