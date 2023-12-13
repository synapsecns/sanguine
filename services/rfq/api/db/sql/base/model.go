package base

import (
	"time"

	"github.com/shopspring/decimal"
)

// define common field names. See package docs  for an explanation of why we have to do this.
// note: some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {

}

var ()

type Quote struct {
	// ID is the unique identifier saved of each quote provided
	ID uint64 `gorm:"column:id;primaryKey;"`
	// DestChainID is the chain which the relayer is willing to provide liquidity for
	DestChainID uint64 `gorm:"column:dest_chain_id;index"`
	// DestToken is the token address for which the relayer is providing liquidity
	DestTokenAddr string `gorm:"column:token;index"`
	// DestAmount is the max amount of liquidity which exists for a given destination token, provided in the destination token decimals
	DestAmount decimal.Decimal `gorm:"column:dest_amount"`
	// Price is the price per origin token provided for which a relayer is indicating willingness to relay
	Price decimal.Decimal `gorm:"column:price"`
	// UpdatedAt is the time that the quote was last upserted
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
