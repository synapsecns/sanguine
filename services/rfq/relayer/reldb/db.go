package reldb

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"

	"github.com/synapsecns/sanguine/ethergo/listener/db"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/dbcommon"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
)

// Writer is the interface for writing to the database.
type Writer interface {
	// StoreQuoteRequest stores a quote request. If one already exists, only  the status will be updated
	// TODO: find a better way to describe this in the name
	StoreQuoteRequest(ctx context.Context, request QuoteRequest) error
	// StoreRebalance stores a rebalance.
	StoreRebalance(ctx context.Context, rebalance Rebalance) error
	// UpdateQuoteRequestStatus updates the status of a quote request
	UpdateQuoteRequestStatus(ctx context.Context, id [32]byte, status QuoteRequestStatus, prevStatus *QuoteRequestStatus) error
	// UpdateRebalance updates the status of a rebalance action.
	// If the origin is supplied, it will be used to update the ID for the corresponding rebalance model.
	UpdateRebalance(ctx context.Context, rebalance Rebalance, updateID bool) error
	// UpdateLatestRebalance updates a rebalance action.
	// This is meant to be used in cases where a consistent rebalance ID cannot be determined across chains.
	UpdateLatestRebalance(ctx context.Context, rebalance Rebalance) error
	// UpdateDestTxHash updates the dest tx hash of a quote request
	UpdateDestTxHash(ctx context.Context, id [32]byte, destTxHash common.Hash) error
	// UpdateRelayNonce updates the relay nonce of a quote request
	UpdateRelayNonce(ctx context.Context, id [32]byte, nonce uint64) error
}

// Reader is the interface for reading from the database.
type Reader interface {
	// GetQuoteRequestByID gets a quote request by id. Should return ErrNoQuoteForID if not found
	GetQuoteRequestByID(ctx context.Context, id [32]byte) (*QuoteRequest, error)
	// GetQuoteRequestByOriginTxHash gets a quote request by origin tx hash. Should return ErrNoQuoteForTxHash if not found
	GetQuoteRequestByOriginTxHash(ctx context.Context, txHash common.Hash) (*QuoteRequest, error)
	// GetQuoteResultsByStatus gets quote results by status
	GetQuoteResultsByStatus(ctx context.Context, matchStatuses ...QuoteRequestStatus) (res []QuoteRequest, _ error)
	// GetPendingRebalances checks fetches all pending rebalances that involve the given chainIDs.
	GetPendingRebalances(ctx context.Context, chainIDs ...uint64) ([]*Rebalance, error)
	// GetRebalance gets a rebalance by ID. Should return ErrNoRebalanceForID if not found.
	GetRebalanceByID(ctx context.Context, rebalanceID string) (*Rebalance, error)
	// GetDBStats gets the database stats.
	GetDBStats(ctx context.Context) (*sql.DBStats, error)
	// GetStatusCounts gets the counts of quote requests by status.
	GetStatusCounts(ctx context.Context, matchStatuses ...QuoteRequestStatus) (map[QuoteRequestStatus]int, error)
}

// Service is the interface for the database service.
type Service interface {
	Reader
	// SubmitterDB returns the submitter database service.
	SubmitterDB() submitterDB.Service
	Writer
	db.ChainListenerDB
}

var (
	// ErrNoQuoteForID means the quote was not found.
	ErrNoQuoteForID = errors.New("no quote found for tx id")
	// ErrNoQuoteForTxHash means the quote was not found.
	ErrNoQuoteForTxHash = errors.New("no quote found for tx hash")
	// ErrNoRebalanceForID means the rebalance was not found.
	ErrNoRebalanceForID = errors.New("no rebalance found for id")
)

// QuoteRequest is the quote request object.
type QuoteRequest struct {
	BlockNumber         uint64
	OriginTokenDecimals uint8
	RawRequest          []byte
	DestTokenDecimals   uint8
	TransactionID       [32]byte
	Sender              common.Address
	Transaction         fastbridge.IFastBridgeBridgeTransaction
	// Status is the quote request status
	Status       QuoteRequestStatus
	OriginTxHash common.Hash
	DestTxHash   common.Hash
	// RelayNonce is the nonce for the relay transaction.
	RelayNonce uint64
	Reason     string
}

// GetOriginIDPair gets the origin chain id and token address pair.
// for some reason, this is specified as [chainid]-[tokenaddr] in the config.
// this represents the origin pair.
func (q QuoteRequest) GetOriginIDPair() string {
	return fmt.Sprintf("%d-%s", q.Transaction.OriginChainId, q.Transaction.OriginToken.Hex())
}

// GetDestIDPair gets the destination chain id and token address pair.
// for some reason, this is specified as [chainid]-[tokenaddr] in the config.
// this represents the destination pair.
func (q QuoteRequest) GetDestIDPair() string {
	return fmt.Sprintf("%d-%s", q.Transaction.DestChainId, q.Transaction.DestToken.Hex())
}

// QuoteRequestStatus is the status of a quote request in the db.
// This is the primary mechanism for moving data through the app.
//
// TODO: consider making this an interface and exporting that.
//
// EXTREMELY IMPORTANT: DO NOT ADD NEW VALUES TO THIS ENUM UNLESS THEY ARE AT THE END.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=QuoteRequestStatus
type QuoteRequestStatus uint8

const (
	// Seen means the quote request has been seen by the relayer, but not processed or committed to.
	Seen QuoteRequestStatus = iota + 1
	// NotEnoughInventory means the relayer does not have enough inventory to process the request.
	// This can be retried at a later time.
	NotEnoughInventory
	// DeadlineExceeded means the quote request has exceeded the deadline.
	// This is a terminal state.
	DeadlineExceeded
	// WillNotProcess means the relayer will not process the request for some reason.
	// This is a terminal state.
	WillNotProcess
	// CommittedPending means the relayer has committed liquidity to the request to the chain, but it is not yet confirmed on chain.
	CommittedPending
	// CommittedConfirmed means the relayer has committed liquidity to the request to the chain, and original bridge tx has been confirmed on chain.
	CommittedConfirmed
	// RelayStarted means the relayer has called Relay() on the destination chain.
	RelayStarted
	// RelayCompleted means the relayer has called Relay() on the destination chain, and the tx has been confirmed on chain.
	RelayCompleted
	// ProvePosting means the relayer has called Prove() on the origin chain.
	ProvePosting
	// ProvePosted means the relayer has called Prove() on the origin chain, and the tx has been confirmed on chain.
	ProvePosted
	// ClaimPending means the relayer has called Claim() on the origin chain.
	ClaimPending
	// ClaimCompleted means the relayer has called Claim() on the origin chain, and the tx has been confirmed on chain.
	ClaimCompleted
	// RelayRaceLost means another relayer has relayed the tx.
	RelayRaceLost
)

// Int returns the int value of the quote request status.
func (q QuoteRequestStatus) Int() uint8 {
	return uint8(q)
}

// GormDataType implements the gorm common interface for enums.
func (q QuoteRequestStatus) GormDataType() string {
	return dbcommon.EnumDataType
}

// Scan implements the gorm common interface for enums.
func (q *QuoteRequestStatus) Scan(src any) error {
	res, err := dbcommon.EnumScan(src)
	if err != nil {
		return fmt.Errorf("could not scan %w", err)
	}
	newStatus := QuoteRequestStatus(res)
	*q = newStatus
	return nil
}

// Value implements the gorm common interface for enums.
func (q QuoteRequestStatus) Value() (driver.Value, error) {
	// nolint: wrapcheck
	return dbcommon.EnumValue(q)
}

var _ dbcommon.Enum = (*QuoteRequestStatus)(nil)

// Rebalance represents a rebalance action.
type Rebalance struct {
	RebalanceID     *string
	Origin          uint64
	Destination     uint64
	OriginAmount    *big.Int
	OriginTokenAddr common.Address
	Status          RebalanceStatus
	OriginTxHash    common.Hash
	DestTxHash      common.Hash
	TokenName       string
}

// RebalanceStatus is the status of a rebalance action in the db.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=RebalanceStatus
type RebalanceStatus uint8

const (
	// RebalanceInitiated means the rebalance transaction has been initiated.
	RebalanceInitiated RebalanceStatus = iota + 1
	// RebalancePending means the rebalance transaction has been confirmed on the origin.
	RebalancePending
	// RebalanceCompleted means the rebalance transaction has been confirmed on the destination.
	RebalanceCompleted
)

// Int returns the int value of the quote request status.
func (r RebalanceStatus) Int() uint8 {
	return uint8(r)
}

// GormDataType implements the gorm common interface for enums.
func (r RebalanceStatus) GormDataType() string {
	return dbcommon.EnumDataType
}

// Scan implements the gorm common interface for enums.
func (r *RebalanceStatus) Scan(src any) error {
	res, err := dbcommon.EnumScan(src)
	if err != nil {
		return fmt.Errorf("could not scan %w", err)
	}
	newStatus := RebalanceStatus(res)
	*r = newStatus
	return nil
}

// Value implements the gorm common interface for enums.
func (r RebalanceStatus) Value() (driver.Value, error) {
	// nolint: wrapcheck
	return dbcommon.EnumValue(r)
}

var _ dbcommon.Enum = (*RebalanceStatus)(nil)
