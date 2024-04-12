package db

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/core/dbcommon"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"math/big"
)

// ErrNoLatestBlockForChainID is returned when no block exists for the chain.
var ErrNoLatestBlockForChainID = errors.New("no latest block for chainId")

// Reader is the interface for reading from the database.
type Reader interface {
	LatestBlockForChain(ctx context.Context, chainID uint64) (uint64, error)
	GetInterchainTXsByStatus(ctx context.Context, statuses ...ExecutableStatus) ([]TransactionSent, error)
}

// Writer is the interface for writing to the database.
type Writer interface {
	PutLatestBlock(ctx context.Context, chainID, height uint64) error
	StoreInterchainTransaction(ctx context.Context, originChainID *big.Int, interchainTx *interchainclient.InterchainClientV1InterchainTransactionSent, options *interchainclient.OptionsV1, encodedTX []byte) error
	UpdateInterchainTransactionStatus(ctx context.Context, transactionid [32]byte, statuses ExecutableStatus) error
}

// Service is the interface for the database service.
type Service interface {
	Reader
	Writer
	// SubmitterDB returns the submitter database service.
	SubmitterDB() submitterDB.Service
}

// ExecutableStatus is the status of a synapse request.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=ExecutableStatus
type ExecutableStatus uint8

const (
	// Seen is the status of a synapse request that has been seen.
	Seen ExecutableStatus = iota + 1
	// Ready is the status of a synapse request that is ready to be executed.
	Ready
	// Executed is the status of a synapse request that has been executed.
	Executed
)

// Int returns the integer value of the synapse request status.
func (s ExecutableStatus) Int() uint8 {
	return uint8(s)
}

// GormDataType implements the gorm common interface for enums.
func (s ExecutableStatus) GormDataType() string {
	return dbcommon.EnumDataType
}

// Scan implements the gorm common interface for enums.
func (s *ExecutableStatus) Scan(src any) error {
	res, err := dbcommon.EnumScan(src)
	if err != nil {
		return fmt.Errorf("could not scan %w", err)
	}
	newStatus := ExecutableStatus(res)
	*s = newStatus
	return nil
}

// Value implements the gorm common interface for enums.
func (s ExecutableStatus) Value() (driver.Value, error) {
	// nolint: wrapcheck
	return dbcommon.EnumValue(s)
}

var _ dbcommon.Enum = (*ExecutableStatus)(nil)
