package db

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-datastore"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/core/dbcommon"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
)

// ErrNoLatestBlockForChainID is returned when no block exists for the chain.
var ErrNoLatestBlockForChainID = errors.New("no latest block for chainId")

// Reader is the interface for reading from the database.
type Reader interface {
	PutLatestBlock(ctx context.Context, chainID, height uint64) error
	GetQuoteResultsByStatus(ctx context.Context, matchStatuses ...SynapseRequestStatus) (res []SignRequest, _ error)
}

// Writer is the interface for writing to the database.
type Writer interface {
	LatestBlockForChain(ctx context.Context, chainID uint64) (uint64, error)
}

// Datstores contain the datastores for the p2p comms.
type Datstores interface {
	DatastoreForSigner(address common.Address) (datastore.Batching, error)
	GlobalDatastore() (datastore.Batching, error)
}

// Service is the interface for the database service.
type Service interface {
	Reader
	Writer
	Datstores
	// SubmitterDB returns the submitter database service.
	SubmitterDB() submitterDB.Service
	StoreInterchainTransactionReceived(ctx context.Context, sr synapsemodule.SynapseModuleModuleMessageSent) error
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

type SignRequest struct {
	// TXHash is the hash of the transaction
	TXHash common.Hash
	// TransactionID is the ID of the transaction
	TransactionID common.Hash
	// Transaction is the transaction
	Transaction []byte
	// Nonce is the nonce of the raw evm tx
	Nonce uint64
	// Status is the status of the transaction
	Status SynapseRequestStatus
	// Sender is the sender of the transaction
	Sender common.Address
	// Receiver is the receiver of the transaction
	Receiver common.Address
	// OriginChainID is the chain id the transaction hash was sent on
	OriginChainID int
	// DestinationChainID is the chain id the transaction hash will be sent on
	DestinationChainID int
	// ModuleRequiredResponses is the number of responses required for the module
	ModuleRequiredResponses int
	// Signature is the signature of the transaction
	Signature map[common.Address][]byte
}
