package model

// GetQuoteResponse contains the schema for a GET /quote response.
type GetQuoteResponse struct {
	// OriginChainID is the chain which the relayer is willing to relay from
	OriginChainID int `json:"origin_chain_id"`
	// OriginTokenAddr is the token address for which the relayer willing to relay from
	OriginTokenAddr string `json:"origin_token_addr"`
	// DestChainID is the chain which the relayer is willing to relay to
	DestChainID int `json:"dest_chain_id"`
	// DestToken is the token address for which the relayer willing to relay to
	DestTokenAddr string `json:"dest_token_addr"`
	// DestAmount is the max amount of liquidity which exists for a given destination token, provided in the destination token decimals
	DestAmount string `json:"dest_amount"`
	// MaxOriginAmount is the maximum amount of origin tokens bridgeable
	MaxOriginAmount string `json:"max_origin_amount"`
	// FixedFee is the fixed fee for the quote, provided in the destination token terms
	FixedFee string `json:"fixed_fee"`
	// Address of the relayer providing the quote
	RelayerAddr string `json:"relayer_addr"`
	// OriginFastBridgeAddress is the address of the fast bridge contract on the origin chain
	OriginFastBridgeAddress string `json:"origin_fast_bridge_address"`
	// DestFastBridgeAddress is the address of the fast bridge contract on the destination chain
	DestFastBridgeAddress string `json:"dest_fast_bridge_address"`
	// UpdatedAt is the time that the quote was last upserted
	UpdatedAt string `json:"updated_at"`
}
