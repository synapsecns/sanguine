package relapi

// GetQuoteRequestStatusResponse contains the schema for a GET /quote response.
type GetQuoteRequestStatusResponse struct {
	Status string `json:"status"`
	TxID   string `json:"tx_id"`
	TxHash string `json:"tx_hash"`
}

// PutTxRetryResponse contains the schema for a PUT /tx/retry response.
type PutTxRetryResponse struct {
	TxID      string `json:"tx_id"`
	ChainID   uint32 `json:"chain_id"`
	Nonce     uint64 `json:"nonce"`
	GasAmount string `json:"gas_amount"`
}
