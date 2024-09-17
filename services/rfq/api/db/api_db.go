// Package db provides the database interfaces and types for the RFQ API.
package db

import (
	"context"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

// Quote is the database model for a quote.
type Quote struct {
	// OriginChainID is the chain which the relayer is willing to relay from
	OriginChainID uint64 `gorm:"column:origin_chain_id;index;primaryKey"`
	// OriginTokenAddr is the token address for which the relayer willing to relay from
	OriginTokenAddr string `gorm:"column:origin_token;index;primaryKey"`
	// DestChainID is the chain which the relayer is willing to relay to
	DestChainID uint64 `gorm:"column:dest_chain_id;index;primaryKey"`
	// DestToken is the token address for which the relayer willing to relay to
	DestTokenAddr string `gorm:"column:dest_token;index;primaryKey"`
	// DestAmount is the max amount of liquidity which exists for a given destination token, provided in the destination token decimals
	DestAmount decimal.Decimal `gorm:"column:dest_amount"`
	// MaxOriginAmount is the maximum amount of origin tokens bridgeable
	MaxOriginAmount decimal.Decimal `gorm:"column:max_origin_amount"`
	// FixedFee is the fixed fee for the quote, provided in the destination token terms
	FixedFee decimal.Decimal `gorm:"column:fixed_fee"`
	// Address of the relayer providing the quote
	RelayerAddr string `gorm:"column:relayer_address;primaryKey"`
	// OriginFastBridgeAddress is the address of the fast bridge contract on the origin chain
	OriginFastBridgeAddress string `gorm:"column:origin_fast_bridge_address"`
	// DestFastBridgeAddress is the address of the fast bridge contract on the destination chain
	DestFastBridgeAddress string `gorm:"column:dest_fast_bridge_address"`
	// UpdatedAt is the time that the quote was last upserted
	UpdatedAt time.Time
}

// ActiveQuoteRequestStatus is the status of a quote request in the db.
// This is the primary mechanism for moving data through the app.
//
// TODO: consider making this an interface and exporting that.
//
// EXTREMELY IMPORTANT: DO NOT ADD NEW VALUES TO THIS ENUM UNLESS THEY ARE AT THE END.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=ActiveQuoteRequestStatus
type ActiveQuoteRequestStatus uint8

const (
	// Received means the quote request has been received by the server.
	Received ActiveQuoteRequestStatus = iota + 1
	// Pending means the quote request is pending awaiting relayer responses.
	Pending
	// Expired means the quote request has expired without any valid responses.
	Expired
	// Fulfilled means the quote request has been fulfilled.
	Fulfilled
)

// Int returns the int value of the quote request status.
func (q ActiveQuoteRequestStatus) Int() uint8 {
	return uint8(q)
}

// GormDataType implements the gorm common interface for enums.
func (q ActiveQuoteRequestStatus) GormDataType() string {
	return dbcommon.EnumDataType
}

// Scan implements the gorm common interface for enums.
func (q *ActiveQuoteRequestStatus) Scan(src any) error {
	res, err := dbcommon.EnumScan(src)
	if err != nil {
		return fmt.Errorf("could not scan %w", err)
	}
	newStatus := ActiveQuoteRequestStatus(res)
	*q = newStatus
	return nil
}

// Value implements the gorm common interface for enums.
func (q ActiveQuoteRequestStatus) Value() (driver.Value, error) {
	// nolint: wrapcheck
	return dbcommon.EnumValue(q)
}

var _ dbcommon.Enum = (*ActiveQuoteRequestStatus)(nil)

// ActiveQuoteResponseStatus is the status of a quote request in the db.
// This is the primary mechanism for moving data through the app.
//
// TODO: consider making this an interface and exporting that.
//
// EXTREMELY IMPORTANT: DO NOT ADD NEW VALUES TO THIS ENUM UNLESS THEY ARE AT THE END.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=ActiveQuoteResponseStatus
type ActiveQuoteResponseStatus uint8

const (
	// Considered means the quote request was considered by the relayer, but was not ultimately the fulfilling response.
	Considered ActiveQuoteResponseStatus = iota + 1
	// Returned means the quote request was returned by the relayer to the user.
	Returned
	// PastExpiration means the quote request was received, but past the expiration window.
	PastExpiration
	// Malformed means that the quote request was malformed.
	Malformed
	// Duplicate means that the quote request was a duplicate.
	Duplicate
)

// Int returns the int value of the quote request status.
func (q ActiveQuoteResponseStatus) Int() uint8 {
	return uint8(q)
}

// GormDataType implements the gorm common interface for enums.
func (q ActiveQuoteResponseStatus) GormDataType() string {
	return dbcommon.EnumDataType
}

// Scan implements the gorm common interface for enums.
func (q *ActiveQuoteResponseStatus) Scan(src any) error {
	res, err := dbcommon.EnumScan(src)
	if err != nil {
		return fmt.Errorf("could not scan %w", err)
	}
	newStatus := ActiveQuoteResponseStatus(res)
	*q = newStatus
	return nil
}

// Value implements the gorm common interface for enums.
func (q ActiveQuoteResponseStatus) Value() (driver.Value, error) {
	// nolint: wrapcheck
	return dbcommon.EnumValue(q)
}

var _ dbcommon.Enum = (*ActiveQuoteResponseStatus)(nil)

// ActiveQuoteRequest is the database model for an active quote request.
type ActiveQuoteRequest struct {
	RequestID         string                   `gorm:"column:request_id;primaryKey"`
	UserAddress       string                   `gorm:"column:user_address"`
	OriginChainID     uint64                   `gorm:"column:origin_chain_id"`
	OriginTokenAddr   string                   `gorm:"column:origin_token"`
	DestChainID       uint64                   `gorm:"column:dest_chain_id"`
	DestTokenAddr     string                   `gorm:"column:dest_token"`
	OriginAmount      decimal.Decimal          `gorm:"column:origin_amount"`
	ExpirationWindow  time.Duration            `gorm:"column:expiration_window"`
	CreatedAt         time.Time                `gorm:"column:created_at"`
	Status            ActiveQuoteRequestStatus `gorm:"column:status"`
	FulfilledAt       time.Time                `gorm:"column:fulfilled_at"`
	FullfilledQuoteID string                   `gorm:"column:fullfilled_quote_id"`
}

// FromUserRequest converts a model.PutUserQuoteRequest to an ActiveQuoteRequest.
func FromUserRequest(req *model.PutUserQuoteRequest, requestID string) *ActiveQuoteRequest {
	originAmount, _ := decimal.NewFromString(req.Data.OriginAmount)
	return &ActiveQuoteRequest{
		RequestID:        requestID,
		UserAddress:      req.UserAddress,
		OriginChainID:    uint64(req.Data.OriginChainID),
		OriginTokenAddr:  req.Data.OriginTokenAddr,
		DestChainID:      uint64(req.Data.DestChainID),
		DestTokenAddr:    req.Data.DestTokenAddr,
		OriginAmount:     originAmount,
		ExpirationWindow: time.Duration(req.Data.ExpirationWindow),
		CreatedAt:        time.Now(),
		Status:           Received,
	}
}

// ActiveQuoteResponse is the database model for an active quote response.
type ActiveQuoteResponse struct {
	RequestID       string                    `gorm:"column:request_id;primaryKey"`
	QuoteID         string                    `gorm:"column:quote_id"`
	OriginChainID   uint64                    `gorm:"column:origin_chain_id"`
	OriginTokenAddr string                    `gorm:"column:origin_token"`
	DestChainID     uint64                    `gorm:"column:dest_chain_id"`
	DestTokenAddr   string                    `gorm:"column:dest_token"`
	OriginAmount    decimal.Decimal           `gorm:"column:origin_amount"`
	DestAmount      decimal.Decimal           `gorm:"column:dest_amount"`
	RelayerAddr     string                    `gorm:"column:relayer_address"`
	UpdatedAt       time.Time                 `gorm:"column:updated_at"`
	Status          ActiveQuoteResponseStatus `gorm:"column:status"`
}

// FromRelayerResponse converts a model.RelayerWsQuoteResponse to an ActiveQuoteResponse.
func FromRelayerResponse(resp *model.RelayerWsQuoteResponse) *ActiveQuoteResponse {
	originAmount, _ := decimal.NewFromString(resp.Data.OriginAmount)
	destAmount, _ := decimal.NewFromString(*resp.Data.DestAmount)
	return &ActiveQuoteResponse{
		RequestID:       resp.RequestID,
		QuoteID:         resp.QuoteID,
		OriginChainID:   uint64(resp.Data.OriginChainID),
		OriginTokenAddr: resp.Data.OriginTokenAddr,
		DestChainID:     uint64(resp.Data.DestChainID),
		DestTokenAddr:   resp.Data.DestTokenAddr,
		OriginAmount:    originAmount,
		DestAmount:      destAmount,
		RelayerAddr:     *resp.Data.RelayerAddress,
		UpdatedAt:       resp.UpdatedAt,
		Status:          Considered,
	}
}

// APIDBReader is the interface for reading from the database.
type APIDBReader interface {
	// GetQuotesByDestChainAndToken gets quotes from the database by destination chain and token.
	GetQuotesByDestChainAndToken(ctx context.Context, destChainID uint64, destTokenAddr string) ([]*Quote, error)
	// GetQuotesByOriginAndDestination gets quotes from the database by origin and destination.
	GetQuotesByOriginAndDestination(ctx context.Context, originChainID uint64, originTokenAddr string, destChainID uint64, destTokenAddr string) ([]*Quote, error)
	// GetQuotesByRelayerAddress gets quotes from the database by relayer address.
	GetQuotesByRelayerAddress(ctx context.Context, relayerAddress string) ([]*Quote, error)
	// GetActiveQuoteRequests gets active quote requests from the database.
	GetActiveQuoteRequests(ctx context.Context, matchStatuses ...ActiveQuoteRequestStatus) ([]*ActiveQuoteRequest, error)
	// GetAllQuotes retrieves all quotes from the database.
	GetAllQuotes(ctx context.Context) ([]*Quote, error)
}

// APIDBWriter is the interface for writing to the database.
type APIDBWriter interface {
	// UpsertQuote upserts a quote in the database.
	UpsertQuote(ctx context.Context, quote *Quote) error
	// UpsertQuotes upserts multiple quotes in the database.
	UpsertQuotes(ctx context.Context, quotes []*Quote) error
	// InsertActiveQuoteRequest inserts an active quote request into the database.
	InsertActiveQuoteRequest(ctx context.Context, req *model.PutUserQuoteRequest, requestID string) error
	// UpdateActiveQuoteRequestStatus updates the status of an active quote request in the database.
	UpdateActiveQuoteRequestStatus(ctx context.Context, requestID string, status ActiveQuoteRequestStatus) error
	// InsertActiveQuoteResponse inserts an active quote response into the database.
	InsertActiveQuoteResponse(ctx context.Context, resp *model.RelayerWsQuoteResponse) error
	// UpdateActiveQuoteResponseStatus updates the status of an active quote response in the database.
	UpdateActiveQuoteResponseStatus(ctx context.Context, requestID string, status ActiveQuoteResponseStatus) error
}

// APIDB is the interface for the database service.
type APIDB interface {
	APIDBReader
	APIDBWriter
}
