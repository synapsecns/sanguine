package model

import "time"

// PutRelayerQuoteRequest contains the schema for a PUT /quote request.
type PutRelayerQuoteRequest struct {
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

// PutBulkQuotesRequest contains the schema for a PUT /quote request.
type PutBulkQuotesRequest struct {
	Quotes []PutRelayerQuoteRequest `json:"quotes"`
}

// PutAckRequest contains the schema for a PUT /ack request.
type PutAckRequest struct {
	TxID        string `json:"tx_id"`
	DestChainID int    `json:"dest_chain_id"`
}

// GetQuoteSpecificRequest contains the schema for a GET /quote request with specific params.
type GetQuoteSpecificRequest struct {
	OriginChainID   int    `json:"originChainId"`
	OriginTokenAddr string `json:"originTokenAddr"`
	DestChainID     int    `json:"destChainId"`
	DestTokenAddr   string `json:"destTokenAddr"`
}

// PutRFQRequest represents a user request for quote.
type PutRFQRequest struct {
	UserAddress  string    `json:"user_address"`
	IntegratorID string    `json:"integrator_id"`
	QuoteTypes   []string  `json:"quote_types"`
	Data         QuoteData `json:"data"`
}

// QuoteRequest represents a request for a quote.
type QuoteRequest struct {
	RequestID string    `json:"request_id"`
	Data      QuoteData `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

// QuoteData represents the data within a quote request.
type QuoteData struct {
	OriginChainID    int     `json:"origin_chain_id"`
	DestChainID      int     `json:"dest_chain_id"`
	OriginTokenAddr  string  `json:"origin_token_addr"`
	DestTokenAddr    string  `json:"dest_token_addr"`
	OriginAmount     string  `json:"origin_amount"`
	ExpirationWindow int64   `json:"expiration_window"`
	ZapData          string  `json:"zap_data"`
	ZapNative        string  `json:"zap_native"`
	DestAmount       *string `json:"dest_amount"`
	RelayerAddress   *string `json:"relayer_address"`
	QuoteID          *string `json:"quote_id"`
}

// WsRFQRequest represents a request for a quote to a relayer.
type WsRFQRequest struct {
	RequestID string    `json:"request_id"`
	Data      QuoteData `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

// SubscribeActiveRFQRequest represents a request to subscribe to active quotes.
// Note that this request is not actually bound to the request body, but rather the chain IDs
// are encoded under the ChainsHeader.
type SubscribeActiveRFQRequest struct {
	ChainIDs []int `json:"chain_ids"`
}

// NewWsRFQRequest creates a new WsRFQRequest.
func NewWsRFQRequest(data QuoteData, requestID string) *WsRFQRequest {
	return &WsRFQRequest{
		RequestID: requestID,
		Data:      data,
		CreatedAt: time.Now(),
	}
}
