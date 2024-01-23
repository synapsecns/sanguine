package db

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"golang.org/x/exp/slices"
	"math/big"
)

// Service is the interface for the tx queue database.
// note: the other files in this package (base, sqlite, mysql) provide a suggested implementation.
// you can implement these yourself. If you plan on importing them, you should wrap in your own service.
//
//go:generate go run github.com/vektra/mockery/v2 --name Service --output ./mocks --case=underscore
type Service interface {
	// GetNonceForChainID gets the nonce for a given chain id.
	GetNonceForChainID(ctx context.Context, fromAddress common.Address, chainID *big.Int) (nonce uint64, err error)
	// PutTXS stores a tx in the database.
	PutTXS(ctx context.Context, txs ...TX) error
	// GetTXS gets all txs for a given address and chain id. If chain id is nil, it will get all txs for the address.
	GetTXS(ctx context.Context, fromAddress common.Address, chainID *big.Int, statuses ...Status) (txs []TX, err error)
	// MarkAllBeforeOrAtNonceReplacedOrConfirmed marks all txs for a given chain id and address before a given nonce as replaced or confirmed.
	// TODO: cleaner function name
	MarkAllBeforeOrAtNonceReplacedOrConfirmed(ctx context.Context, signer common.Address, chainID *big.Int, nonce uint64) error
	// DBTransaction executes a transaction on the database.
	// the function passed in will be passed a new service that is scoped to the transaction.
	DBTransaction(ctx context.Context, f TransactionFunc) error
	// GetAllTXAttemptByStatus gets all txs for a given address and chain id with a given status.
	GetAllTXAttemptByStatus(ctx context.Context, fromAddress common.Address, chainID *big.Int, matchStatuses ...Status) (txs []TX, err error)
	// GetNonceStatus returns the nonce status for a given nonce by aggregating all attempts and finding the highest status.
	GetNonceStatus(ctx context.Context, fromAddress common.Address, chainID *big.Int, nonce uint64) (status Status, err error)
	// GetNonceAttemptsByStatus gets all txs for a given address and chain id with a given status and nonce.
	GetNonceAttemptsByStatus(ctx context.Context, fromAddress common.Address, chainID *big.Int, nonce uint64, matchStatuses ...Status) (txs []TX, err error)
}

// TransactionFunc is a function that can be passed to DBTransaction.
type TransactionFunc func(ctx context.Context, svc Service) error

// Status is the status of a tx.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=Status -linecomment
type Status uint8

// Important: do not modify the order of these constants.
// if one needs to be removed, replace it with a no-op status.
// additionally, due to the GetMaxNoncestatus function, statuses are currently assumed to be in order.
// if you need to modify this functionality, please update that function. to reflect that the highest status
// isno longer the expected end status.
const (
	// Pending is the status of a tx that has not been processed yet.
	Pending Status = iota + 1 // Pending
	// Stored is the status of a tx that has been stored.
	Stored // Stored
	// Submitted is the status of a tx that has been submitted.
	Submitted // Submitted
	// FailedSubmit is the status of a tx that has failed to submit.
	FailedSubmit // Failed
	// ReplacedOrConfirmed is the status of a tx that has been replaced by a new tx or confirmed. The actual status will be set later.
	ReplacedOrConfirmed // ReplacedOrConfirmed
	// Replaced is the status of a tx that has been replaced by a new tx.
	Replaced // Replaced
	// Confirmed is the status of a tx that has been confirmed.
	Confirmed // Confirmed
)

var allStatusTypes = []Status{Pending, Stored, Submitted, FailedSubmit, ReplacedOrConfirmed, Replaced, Confirmed}

// AllStatusTypes returns all status types.
// it is exported for testing purposes
//
// These are guaranteed to be in order.
func AllStatusTypes() []Status {
	return allStatusTypes
}

// check to make sure all statuses are included in all status types.
func _() {
	for i := 0; i < len(_Status_index); i++ {
		statusNum := i + 1
		status := Status(statusNum)
		if status.String() == "" {
			panic(fmt.Sprintf("invalid status: %d", status))
		}
		if status.String() != AllStatusTypes()[i].String() {
			panic(fmt.Sprintf("status string and all status types do not match: %s, %s", status.String(), AllStatusTypes()[i]))
		}
	}
}

// Int returns the uint8 representation of the status.
func (s Status) Int() uint8 {
	return uint8(s)
}

// GormDataType returns the gorm data type for the status.
func (s Status) GormDataType() string {
	return dbcommon.EnumDataType
}

// Scan implements the gorm Scanner interface.
func (s *Status) Scan(src interface{}) error {
	res, err := dbcommon.EnumScan(src)
	if err != nil {
		return fmt.Errorf("could not scan status: %w", err)
	}
	newStatus := Status(res)
	*s = newStatus

	if !slices.Contains[[]Status](allStatusTypes, *s) {
		return fmt.Errorf("invalid status: %d", res)
	}

	return nil
}

// Value implements the gorm Valuer interface.
func (s *Status) Value() (driver.Value, error) {
	//nolint: wrapcheck
	return dbcommon.EnumValue(s)
}

var _ dbcommon.EnumInter = (*Status)(nil)

var (
	// ErrNoNonceForChain is the error returned when there is no nonce for a given chain id.
	ErrNoNonceForChain = errors.New("no nonce exists for this chain")
	// ErrNonceNotExist is the error returned when a nonce does not exist.
	ErrNonceNotExist = errors.New("nonce does not exist")
)
