package base

import "github.com/synapsecns/sanguine/core/dbcommon"

func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	transactionIDFieldName = namer.GetConsistentName("TransactionID")
	statusFieldName = namer.GetConsistentName("Status")
}

var (
	// transactionIDFieldName is the name of the transaction field.
	transactionIDFieldName string
	// statusIDFieldName is the name of the status field.
	statusFieldName string
)
