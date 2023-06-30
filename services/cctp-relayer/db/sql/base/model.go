package base

import (
	"github.com/synapsecns/sanguine/core/dbcommon"
)

// define common field names. See package docs  for an explanation of why we have to do this.
// note: some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	OriginTxHashFieldName = namer.GetConsistentName("OriginTxHash")
	DestTxHashFieldName = namer.GetConsistentName("DestTxHash")
	NonceFieldName = namer.GetConsistentName("DestNonce")
	OriginChainIDFieldName = namer.GetConsistentName("OriginChainID")
	DestChainIDFieldName = namer.GetConsistentName("DestChainID")
	MessageFieldName = namer.GetConsistentName("Message")
	MessageHashFieldName = namer.GetConsistentName("MessageHash")
	AttestationFieldName = namer.GetConsistentName("Attestation")
	RequestVersionFieldName = namer.GetConsistentName("RequestVersion")
	FormattedRequestFieldName = namer.GetConsistentName("FormattedRequest")
	RequestIDFieldName = namer.GetConsistentName("RequestID")
	BlockNumberFieldName = namer.GetConsistentName("BlockNumber")
	StateFieldName = namer.GetConsistentName("State")
}

var (
	// OriginTxHashFieldName gets the burn tx hash field name.
	OriginTxHashFieldName string
	// DestTxHashFieldName gets the burn tx hash field name.
	DestTxHashFieldName string
	// NonceFieldName gets the mint tx hash field name.
	NonceFieldName string
	// OriginChainIDFieldName gets the origin chain ID field name.
	OriginChainIDFieldName string
	// DestChainIDFieldName gets the destination chain ID field name.
	DestChainIDFieldName string
	// MessageFieldName gets the message field name.
	MessageFieldName string
	// MessageHashFieldName gets the message hash field name.
	MessageHashFieldName string
	// AttestationFieldName gets the signature field name.
	AttestationFieldName string
	// RequestVersionFieldName gets the request version field name.
	RequestVersionFieldName string
	// FormattedRequestFieldName gets the formatted request field name.
	FormattedRequestFieldName string
	// RequestIDFieldName gets the request id field name.
	RequestIDFieldName string
	// BlockNumberFieldName gets the block number field name.
	BlockNumberFieldName string
	// StateFieldName gets the state field name.
	StateFieldName string
)
