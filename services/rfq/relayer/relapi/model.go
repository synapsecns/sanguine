package relapi

// GetQuoteRequestStatusResponse contains the schema for a GET /quote response.
type GetQuoteRequestStatusResponse struct {
	Status string `json:"status"`
	TxID   string `json:"tx_id"`
	TxHash string `json:"tx_hash"`
}
