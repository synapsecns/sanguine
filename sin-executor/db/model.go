package db

import "github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"

type TransactionSent struct {
	// EncodedTX is the encoded transaction.
	EncodedTX []byte
	// Transaction is the transaction that was sent.
	interchainclient.InterchainClientV1InterchainTransactionSent
	// Status is the status of the transaction.
	Status ExecutableStatus
}
