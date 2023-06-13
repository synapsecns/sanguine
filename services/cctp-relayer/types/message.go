package types

import "github.com/ethereum/go-ethereum/common"

// Message is the information about a message parsed by the CCTPRelayer.
type Message struct {
	// Hash of USDC burn transaction
	BurnTxHash common.Hash
	// Chain ID of the origin chain
	OriginChainID uint32
	// Chain ID of the destination chain
	DestChainID uint32
	// Raw bytes of message produced by Circle's MessageTransmitter
	Message []byte
	// Keccak256 hash of message bytes
	MessageHash common.Hash
	// Attestation produced by Circle's API: https://developers.circle.com/stablecoin/reference/getattestation
	Signature []byte
	// Version of the request
	RequestVersion uint32
	// Formatted request produced by SynapseCCTP
	FormattedRequest []byte
	// BlockNumber is the block number.
	BlockNumber uint64 `gorm:"column:block_number"`
}
