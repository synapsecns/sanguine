package base

import "github.com/synapsecns/sanguine/core/dbcommon"

func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	chainIDFieldName = namer.GetConsistentName("ChainID")
	blockNumberFieldName = namer.GetConsistentName("BlockNumber")
	transactionIDFieldName = namer.GetConsistentName("TransactionID")
	statusFieldName = namer.GetConsistentName("Status")
}

var (
	// chainIDFieldName gets the chain id field name.
	chainIDFieldName string
	// blockNumberFieldName is the name of the block number field.
	blockNumberFieldName string
	// transactionIDFieldName is the name of the transaction field.
	transactionIDFieldName string
	// statusIDFieldName is the name of the status field.
	statusFieldName string
)
