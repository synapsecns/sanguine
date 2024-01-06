package model

// PutQuoteRequest contains the schema for a PUT /quote request.
type PutQuoteRequest struct {
	OriginChainID           int    `json:"origin_chain_id"`
	OriginTokenAddr         string `json:"origin_token_addr"`
	DestChainID             int    `json:"dest_chain_id"`
	DestTokenAddr           string `json:"dest_token_addr"`
	DestAmount              string `json:"dest_amount"`
	MaxOriginAmount         string `json:"max_origin_amount"`
	FixedFee                string `json:"fixed_fee"`
	OriginFastBridgeAddress string `json:"origin_fast_bridge_address"`
	DestFastBridgeAddress   string `json:"dest_fast_bridge_address"`
}

// GetQuoteSpecificRequest contains the schema for a GET /quote request with specific params.
type GetQuoteSpecificRequest struct {
	OriginChainID   int    `json:"originChainId"`
	OriginTokenAddr string `json:"originTokenAddr"`
	DestChainID     int    `json:"destChainId"`
	DestTokenAddr   string `json:"destTokenAddr"`
}
