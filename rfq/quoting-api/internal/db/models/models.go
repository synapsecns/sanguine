// Package models contains teh database models
package models

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// GetAllModels gets all models to migrate.
func GetAllModels() (allModels []interface{}) {
	return []interface{}{&Request{}, &Quote{}}
}

// Request is the request to get a quote.
type Request struct {
	OriginChainID uint            `form:"origin_chain_id"`
	DestChainID   uint            `form:"dest_chain_id"`
	OriginToken   string          `form:"origin_token"`
	DestToken     string          `form:"dest_token"`
	OriginAmount  decimal.Decimal `form:"origin_amount" gorm:"type:numeric"`
	UpdatedAtLast time.Time       `form:"updated_at_last"` // for liveness checks on q.UpdatedAt
}

// Quote represents bindings/FastBridge.go::IFastBridgeBridgeTransaction.
type Quote struct {
	ID            uint            `gorm:"primaryKey" json:"id" uri:"id"`
	Relayer       string          `json:"relayer" binding:"required"`
	OriginChainID uint            `json:"origin_chain_id" binding:"required"`
	DestChainID   uint            `json:"dest_chain_id" binding:"required"`
	OriginToken   string          `json:"origin_token" binding:"required"`
	DestToken     string          `json:"dest_token" binding:"required"`
	OriginAmount  decimal.Decimal `json:"origin_amount" binding:"required" gorm:"type:numeric"`
	DestAmount    decimal.Decimal `json:"dest_amount" binding:"required" gorm:"type:numeric"`
	Price         decimal.Decimal `json:"price" gorm:"type:numeric"` // price = destAmount <quote> / originAmount <base>
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeletedAt     gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
}

// Validate validates the quote attributes.
func (q *Quote) Validate() (err error) {
	// zero address and zero int string
	zeroAddress := (common.Address{}).Hex()
	zeroDecimal := decimal.Zero

	//nolint: gocritic
	// TODO: make this a switch for clarity + readability
	if q.OriginChainID == q.DestChainID {
		err = fmt.Errorf("invalid quote: q.OriginChainID == q.DestChainID")
		return
	} else if q.OriginToken == zeroAddress || q.DestToken == zeroAddress {
		err = fmt.Errorf("invalid quote: q.Tokens == address(0)")
		return
	} else if q.OriginAmount == zeroDecimal || q.DestAmount == zeroDecimal {
		err = fmt.Errorf("invalid quote: q.Amounts == 0")
		return
	} else if q.Relayer == zeroAddress {
		err = fmt.Errorf("invalid quote: q.Relayer == address(0)") // TODO: test
		return
	}
	return
}

// BeforeCreate hook to set price on insert.
func (q *Quote) BeforeCreate(tx *gorm.DB) (err error) {
	// cannot create quote with existing ID
	if q.ID != 0 {
		err = fmt.Errorf("invalid quote: created q.ID != 0")
		return
	}
	if err = q.Validate(); err != nil {
		return
	}

	q.Price = q.GetPrice()
	return
}

// BeforeSave hook to set price on update.
func (q *Quote) BeforeSave(tx *gorm.DB) (err error) {
	if err = q.Validate(); err != nil {
		return
	}

	q.Price = q.GetPrice()
	return
}

// GetPrice is a util to get the quote price inferred from dest, origin amounts.
func (q *Quote) GetPrice() (price decimal.Decimal) {
	// calculate price using decimals
	price = q.DestAmount.Div(q.OriginAmount)
	return
}
