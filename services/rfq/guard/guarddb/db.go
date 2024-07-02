package guarddb

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/synapsecns/sanguine/ethergo/listener/db"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/dbcommon"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
)

var (
	// ErrNoProvenForID means the proven was not found.
	ErrNoProvenForID = errors.New("no proven found for tx id")
)

// Writer is the interface for writing to the database.
type Writer interface {
	// StorePendingProven stores a pending proven.
	StorePendingProven(ctx context.Context, proven PendingProven) error
	// UpdatePendingProvenStatus updates the status of a pending proven.
	UpdatePendingProvenStatus(ctx context.Context, id [32]byte, status PendingProvenStatus) error
}

// Reader is the interface for reading from the database.
type Reader interface {
	// GetPendingProvenByID gets a quote request by id. Should return ErrNoProvenForID if not found
	GetPendingProvenByID(ctx context.Context, id [32]byte) (*PendingProven, error)
}

// Service is the interface for the database service.
type Service interface {
	Reader
	// SubmitterDB returns the submitter database service.
	SubmitterDB() submitterDB.Service
	Writer
	db.ChainListenerDB
}

// PendingProven is the pending proven object.
type PendingProven struct {
	TransactionID [32]byte
	TxHash        common.Hash
	Status        PendingProvenStatus
}

// PendingProvenStatus is the status of a quote request in the db.
// This is the primary mechanism for moving data through the app.
//
// TODO: consider making this an interface and exporting that.
//
// EXTREMELY IMPORTANT: DO NOT ADD NEW VALUES TO THIS ENUM UNLESS THEY ARE AT THE END.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=PendingProvenStatus
type PendingProvenStatus uint8

const (
	// ProveCalled means the prove() function has been called.
	ProveCalled PendingProvenStatus = iota + 1
	// Validated means the prove() call has been properly validated on the dest chain.
	Validated
	// DisputePending means dispute() has been called in the event of an invalid prove().
	DisputePending
	// Disputed means the dispute() call has been confirmed.
	Disputed
)

// Int returns the int value of the quote request status.
func (q PendingProvenStatus) Int() uint8 {
	return uint8(q)
}

// GormDataType implements the gorm common interface for enums.
func (q PendingProvenStatus) GormDataType() string {
	return dbcommon.EnumDataType
}

// Scan implements the gorm common interface for enums.
func (q *PendingProvenStatus) Scan(src any) error {
	res, err := dbcommon.EnumScan(src)
	if err != nil {
		return fmt.Errorf("could not scan %w", err)
	}
	newStatus := PendingProvenStatus(res)
	*q = newStatus
	return nil
}

// Value implements the gorm common interface for enums.
func (q PendingProvenStatus) Value() (driver.Value, error) {
	// nolint: wrapcheck
	return dbcommon.EnumValue(q)
}

var _ dbcommon.Enum = (*PendingProvenStatus)(nil)
