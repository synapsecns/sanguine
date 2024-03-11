package types

// Message is the information about a message parsed by the CCTPRelayer.
type Message struct {
	// Hash of USDC burn transaction
	OriginTxHash string `gorm:"column:origin_tx_hash"`
	// Hash of USDC mint transaction
	DestTxHash string `gorm:"column:dest_tx_hash"`
	// Nonce of USDC mint transaction
	DestNonce int `gorm:"column:dest_nonce"`
	// Chain ID of the origin chain
	OriginChainID uint32 `gorm:"column:origin_chain_id"`
	// Chain ID of the destination chain
	DestChainID uint32 `gorm:"column:dest_chain_id"`
	// Raw bytes of message produced by Circle's MessageTransmitter
	Message []byte `gorm:"column:message"`
	// Keccak256 hash of message bytes
	MessageHash string `gorm:"column:message_hash;primaryKey"`
	// RequestID of the message
	RequestID string `gorm:"column:request_id"`
	// Attestation produced by Circle's API: https://developers.circle.com/stablecoin/reference/getattestation
	Attestation []byte `gorm:"column:attestation"`
	// Version of the request
	RequestVersion uint32 `gorm:"column:request_version"`
	// Formatted request produced by SynapseCCTP
	FormattedRequest []byte `gorm:"column:formatted_request"`
	// BlockNumber is the block number.
	BlockNumber uint64 `gorm:"column:block_number"`
	// State is the state of the message.
	State MessageState `gorm:"column:state"`
}
