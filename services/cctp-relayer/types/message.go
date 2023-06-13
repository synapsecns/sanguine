package types

// Message is the information about a message parsed by the CCTPRelayer.
type Message struct {
	// Hash of USDC burn transaction
	OriginTxHash string `gorm:"column:send_tx_hash"`
	// Hash of USDC mint transaction
	DestTxHash string `gorm:"column:recv_tx_hash"`
	// Chain ID of the origin chain
	OriginChainID uint32 `gorm:"column:origin_chain_id"`
	// Chain ID of the destination chain
	DestChainID uint32 `gorm:"column:dest_chain_id"`
	// Raw bytes of message produced by Circle's MessageTransmitter
	Message []byte `gorm:"column:message"`
	// Keccak256 hash of message bytes
	MessageHash string `gorm:"column:message_hash"`
	// Attestation produced by Circle's API: https://developers.circle.com/stablecoin/reference/getattestation
	Signature []byte `gorm:"column:signature"`
	// Version of the request
	RequestVersion uint32 `gorm:"column:request_version"`
	// Formatted request produced by SynapseCCTP
	FormattedRequest []byte `gorm:"column:formatted_request"`
	// BlockNumber is the block number.
	BlockNumber uint64 `gorm:"column:block_number"`
	// State is the state of the message.
	State MessageState `gorm:"column:state"`
}
