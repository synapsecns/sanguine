package db

import (
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

// TX is a superset of transaction that includes the gas price.
type TX struct {
	// inherited from types.Transaction
	*types.Transaction
	// creationTime is the time the transaction was last updated
	// this field is unexported to prevent it from being set outside of the package
	// instead we create a unsafe setter function that is only used by implementing dbs package
	creationTime time.Time
	// Status is the status of the transaction
	Status Status
}

// NewTX creates a new TX for use in the db package.
func NewTX(tx *types.Transaction, status Status) TX {
	return TX{
		Transaction: tx,
		Status:      status,
	}
}

// CreationTime is the time the transaction was last updated.
func (t *TX) CreationTime() time.Time {
	return t.creationTime
}

// UnsafeSetCreationTime is an unsafe setter for the creation time
// it is called unsafe to force you to read this comment telling you
// this should only be called if you are creating a db implementation of this package.
// this should not be called in the submmiter itself or any other package.
func (t *TX) UnsafeSetCreationTime(creationTime time.Time) {
	t.creationTime = creationTime
}
