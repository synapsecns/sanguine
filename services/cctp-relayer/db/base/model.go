package base

import (
	"github.com/synapsecns/sanguine/core/dbcommon"
)

// define common field names. See package docs  for an explanation of why we have to do this.
// note: some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	BurnTxHashFieldName = namer.GetConsistentName("BurnTxHash")
	OriginChainIDFieldName = namer.GetConsistentName("OriginChainID")
	DestChainIDFieldName = namer.GetConsistentName("DestChainID")
	MessageFieldName = namer.GetConsistentName("Message")
	MessageHashFieldName = namer.GetConsistentName("MessageHash")
	SignatureFieldName = namer.GetConsistentName("Signature")
	RequestVersionFieldName = namer.GetConsistentName("RequestVersion")
	FormattedRequestFieldName = namer.GetConsistentName("FormattedRequest")
	BlockNumberFieldName = namer.GetConsistentName("BlockNumber")
}

var (
	// BurnTxHashFieldName gets the burn tx hash field name.
	BurnTxHashFieldName string
	// OriginChainIDFieldName gets the origin chain ID field name.
	OriginChainIDFieldName string
	// DestChainIDFieldName gets the destination chain ID field name.
	DestChainIDFieldName string
	// MessageFieldName gets the message field name.
	MessageFieldName string
	// MessageHashFieldName gets the message hash field name.
	MessageHashFieldName string
	// SignatureFieldName gets the signature field name.
	SignatureFieldName string
	// RequestVersionFieldName gets the request version field name.
	RequestVersionFieldName string
	// FormattedRequestFieldName gets the formatted request field name.
	FormattedRequestFieldName string
	// BlockNumberFieldName gets the block number field name.
	BlockNumberFieldName string
)
