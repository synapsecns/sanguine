package model

// GetQuoteResponse contains the schema for a GET /quote response.
type GetQuoteResponse struct {
	// OriginChainID is the chain which the relayer is willing to relay from
	OriginChainID int `gorm:"column:origin_chain_id;index;primaryKey"`
	// OriginTokenAddr is the token address for which the relayer willing to relay from
	OriginTokenAddr string `gorm:"column:origin_token;index;primaryKey"`
	// DestChainID is the chain which the relayer is willing to relay to
	DestChainID int `gorm:"column:dest_chain_id;index;primaryKey"`
	// DestToken is the token address for which the relayer willing to relay to
	DestTokenAddr string `gorm:"column:dest_token;index;primaryKey"`
	// DestAmount is the max amount of liquidity which exists for a given destination token, provided in the destination token decimals
	DestAmount float64 `gorm:"column:dest_amount"`
	// MaxOriginAmount is the maximum amount of origin tokens bridgeable
	MaxOriginAmount float64 `gorm:"column:max_origin_amount"`
	// FixedFee is the fixed fee for the quote, provided in the destination token terms
	FixedFee float64 `gorm:"column:fixed_fee"`
	// Address of the relayer providing the quote
	RelayerAddr string `gorm:"column:relayer_address;primaryKey"`
	// OriginFastBridgeAddress is the address of the fast bridge contract on the origin chain
	OriginFastBridgeAddress string `gorm:"column:origin_fast_bridge_address"`
	// DestFastBridgeAddress is the address of the fast bridge contract on the destination chain
	DestFastBridgeAddress string `gorm:"column:dest_fast_bridge_address"`
	// UpdatedAt is the time that the quote was last upserted
	UpdatedAt string
}
