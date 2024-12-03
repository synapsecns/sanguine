package chain

import (
	"encoding/binary"
	"math/big"

	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridgev2"
)

const (
	// Field sizes in bytes.
	sizeVersion = 2
	sizeChainID = 4
	sizeAddress = 20
	sizeUint256 = 32

	// Field offsets in bytes.
	offsetVersion            = 0
	offsetOriginChainID      = offsetVersion + sizeVersion
	offsetDestChainID        = offsetOriginChainID + sizeChainID
	offsetOriginSender       = offsetDestChainID + sizeChainID
	offsetDestRecipient      = offsetOriginSender + sizeAddress
	offsetOriginToken        = offsetDestRecipient + sizeAddress
	offsetDestToken          = offsetOriginToken + sizeAddress
	offsetOriginAmount       = offsetDestToken + sizeAddress
	offsetDestAmount         = offsetOriginAmount + sizeUint256
	offsetOriginFeeAmount    = offsetDestAmount + sizeUint256
	offsetDeadline           = offsetOriginFeeAmount + sizeUint256
	offsetNonce              = offsetDeadline + sizeUint256
	offsetExclusivityRelayer = offsetNonce + sizeUint256
	offsetExclusivityEndTime = offsetExclusivityRelayer + sizeAddress
	offsetZapNative          = offsetExclusivityEndTime + sizeUint256
	offsetZapData            = offsetZapNative + sizeUint256
)

// Helper function to properly encode uint256.
func padUint256(b *big.Int) []byte {
	// Convert big.Int to bytes
	bytes := b.Bytes()
	// Create 32-byte array (initialized to zeros)
	result := make([]byte, 32)
	// Copy bytes to right side of array (left-pad with zeros)
	copy(result[32-len(bytes):], bytes)
	return result
}

// EncodeBridgeTx encodes a bridge transaction into a byte array.
func EncodeBridgeTx(tx fastbridgev2.IFastBridgeV2BridgeTransactionV2) ([]byte, error) {
	// Initialize with total size including ZapData
	result := make([]byte, offsetZapData+len(tx.ZapData))

	// Version
	result[offsetVersion] = 0
	result[offsetVersion+1] = 2

	// Chain IDs
	binary.BigEndian.PutUint32(result[offsetOriginChainID:offsetOriginChainID+sizeChainID], tx.OriginChainId)
	binary.BigEndian.PutUint32(result[offsetDestChainID:offsetDestChainID+sizeChainID], tx.DestChainId)

	// Addresses
	copy(result[offsetOriginSender:offsetOriginSender+sizeAddress], tx.OriginSender.Bytes())
	copy(result[offsetDestRecipient:offsetDestRecipient+sizeAddress], tx.DestRecipient.Bytes())
	copy(result[offsetOriginToken:offsetOriginToken+sizeAddress], tx.OriginToken.Bytes())
	copy(result[offsetDestToken:offsetDestToken+sizeAddress], tx.DestToken.Bytes())

	// uint256 values
	copy(result[offsetOriginAmount:offsetOriginAmount+sizeUint256], padUint256(tx.OriginAmount))
	copy(result[offsetDestAmount:offsetDestAmount+sizeUint256], padUint256(tx.DestAmount))
	copy(result[offsetOriginFeeAmount:offsetOriginFeeAmount+sizeUint256], padUint256(tx.OriginFeeAmount))
	copy(result[offsetDeadline:offsetDeadline+sizeUint256], padUint256(tx.Deadline))
	copy(result[offsetNonce:offsetNonce+sizeUint256], padUint256(tx.Nonce))

	// Exclusivity address
	copy(result[offsetExclusivityRelayer:offsetExclusivityRelayer+sizeAddress], tx.ExclusivityRelayer.Bytes())

	// More uint256 values
	copy(result[offsetExclusivityEndTime:offsetExclusivityEndTime+sizeUint256], padUint256(tx.ExclusivityEndTime))
	copy(result[offsetZapNative:offsetZapNative+sizeUint256], padUint256(tx.ZapNative))

	// Replace append with copy for ZapData
	copy(result[offsetZapData:], tx.ZapData)

	return result, nil
}
