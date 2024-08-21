package relapi

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

// GetQuoteRequestResponse is the response to a get quote request.
type GetQuoteRequestResponse struct {
	Status          string `json:"status"`
	TxID            string `json:"tx_id"`
	QuoteRequestRaw string `json:"quote_request"`
	OriginTxHash    string `json:"origin_tx_hash"`
	DestTxHash      string `json:"dest_tx_hash"`
	OriginChainID   uint32 `json:"origin_chain_id"`
	DestChainID     uint32 `json:"dest_chain_id"`
	OriginToken     string `json:"origin_token"`
	DestToken       string `json:"dest_token"`
}
