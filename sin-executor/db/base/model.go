package base

import "github.com/synapsecns/sanguine/core/dbcommon"

func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	chainIDFieldName = namer.GetConsistentName("ChainID")
	blockNumberFieldName = namer.GetConsistentName("BlockNumber")
	statusFieldName = namer.GetConsistentName("Status")
	transactionIDFieldName = namer.GetConsistentName("TransactionID")
}

var (
	// chainIDFieldName gets the chain id field name.
	chainIDFieldName string
	// blockNumberFieldName is the name of the block number field.
	blockNumberFieldName string
	// statusFieldName is the name of the status field.
	statusFieldName string
	// transactionIDFieldName is the name of the transaction id field.
	transactionIDFieldName string
)
