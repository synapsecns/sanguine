package relapi

// GetQuoteRequestStatusResponse contains the schema for a GET /quote response.
type GetQuoteRequestStatusResponse struct {
	Status        string `json:"status"`
	TxID          string `json:"tx_id"`
	OriginTxHash  string `json:"origin_tx_hash"`
	OriginChainID uint32 `json:"origin_chain_id"`
	DestTxHash    string `json:"dest_tx_hash"`
	DestChainID   uint32 `json:"dest_chain_id"`
}

// GetTxRetryResponse contains the schema for a PUT /tx/retry response.
type GetTxRetryResponse struct {
	TxID      string `json:"tx_id"`
	ChainID   uint32 `json:"chain_id"`
	Nonce     uint64 `json:"nonce"`
	GasAmount string `json:"gas_amount"`
}

// PutRelayAckResponse contains the schema for a POST /relay/ack response.
type PutRelayAckResponse struct {
	TxID           string `json:"tx_id"`
	ShouldRelay    bool   `json:"should_relay"`
	RelayerAddress string `json:"relayer_address"`
}
