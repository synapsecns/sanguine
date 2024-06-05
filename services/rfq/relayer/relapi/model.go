package relapi

// GetQuoteRequestStatusResponse contains the schema for a GET /quote response.
type GetQuoteRequestStatusResponse struct {
	Status       string `json:"status"`
	TxID         string `json:"tx_id"`
	OriginTxHash string `json:"origin_tx_hash"`
	DestTxHash   string `json:"dest_tx_hash"`
}

// GetTxRetryResponse contains the schema for a PUT /tx/retry response.
type GetTxRetryResponse struct {
	TxID      string `json:"tx_id"`
	ChainID   uint32 `json:"chain_id"`
	Nonce     uint64 `json:"nonce"`
	GasAmount string `json:"gas_amount"`
}
