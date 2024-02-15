package db

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/core/dbcommon"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
)

// ErrNoLatestBlockForChainID is returned when no block exists for the chain.
var ErrNoLatestBlockForChainID = errors.New("no latest block for chainId")

// Reader is the interface for reading from the database.
type Reader interface {
	PutLatestBlock(ctx context.Context, chainID, height uint64) error
}

// Writer is the interface for writing to the database.
type Writer interface {
	LatestBlockForChain(ctx context.Context, chainID uint64) (uint64, error)
}

// Service is the interface for the database service.
type Service interface {
	Reader
	Writer
	// SubmitterDB returns the submitter database service.
	SubmitterDB() submitterDB.Service
}

// SynapseRequestStatus is the status of a synapse request.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=SynapseRequestStatus
type SynapseRequestStatus uint8

const (
	// Seen is the status of a synapse request that has been seen.
	Seen SynapseRequestStatus = iota + 1
)

// Int returns the integer value of the synapse request status.
func (s SynapseRequestStatus) Int() uint8 {
	return uint8(s)
}

// GormDataType implements the gorm common interface for enums.
func (s SynapseRequestStatus) GormDataType() string {
	return dbcommon.EnumDataType
}

// Scan implements the gorm common interface for enums.
func (s *SynapseRequestStatus) Scan(src any) error {
	res, err := dbcommon.EnumScan(src)
	if err != nil {
		return fmt.Errorf("could not scan %w", err)
	}
	newStatus := SynapseRequestStatus(res)
	*s = newStatus
	return nil
}

// Value implements the gorm common interface for enums.
func (s SynapseRequestStatus) Value() (driver.Value, error) {
	// nolint: wrapcheck
	return dbcommon.EnumValue(s)
}

var _ dbcommon.Enum = (*SynapseRequestStatus)(nil)
