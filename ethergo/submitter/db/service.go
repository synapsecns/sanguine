package db

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"golang.org/x/exp/slices"
	"math/big"
)

// Service is the interface for the tx queue database.
// note: the other files in this package (base, sqlite, mysql) provide a suggested implementation.
// you can implement these yourself. If you plan on importing them, you should wrap in your own service.
type Service interface {
	// GetNonceForChainID gets the nonce for a given chain id.
	GetNonceForChainID(ctx context.Context, fromAddress common.Address, chainID *big.Int) (nonce uint64, err error)
	// PutTX stores a tx in the database.
	PutTX(ctx context.Context, tx *types.Transaction, status Status) error
	// GetTXS gets all txs for a given address and chain id.
	GetTXS(ctx context.Context, fromAddress common.Address, chainID *big.Int, statuses ...Status) (txs []*types.Transaction, err error)
}

// Status is the status of a tx.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=Status -linecomment
type Status uint8

// Important: do not modify the order of these constants.
// if one needs to be removed, replace it with a no-op status.
const (
	// Pending is the status of a tx that has not been processed yet.
	Pending Status = iota + 1 // Pending
	// Stored is the status of a tx that has been stored.
	Stored // Stored
	// Replaced is the status of a tx that has been replaced by a new tx.
	Replaced // Replaced
	// Confirmed is the status of a tx that has been confirmed.
	Confirmed // Confirmed
)

var allStatusTypes = []Status{Pending, Replaced, Confirmed}

// check to make sure all statuses are included in all status types.
func _() {
	for i := 0; i < len(_Status_index); i++ {
		statusNum := i + 1
		status := Status(statusNum)
		if status.String() == "" {
			panic(fmt.Sprintf("invalid status: %d", status))
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

	if !slices.Contains[Status](allStatusTypes, *s) {
		return fmt.Errorf("invalid status: %d", res)
	}

	return nil
}

// Value implements the gorm Valuer interface.
func (s *Status) Value() (driver.Value, error) {
	return dbcommon.EnumValue(s)
}

var _ dbcommon.EnumInter = (*Status)(nil)

var ErrNoNonceForChain = errors.New("no nonce exists for this chain")
