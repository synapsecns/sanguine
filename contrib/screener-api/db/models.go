package db

import (
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// AddressIndicators is the address indicators for a given address
type AddressIndicators struct {
	gorm.Model
	// Address is the address to be screened
	Address string `gorm:"column:address;primary_key"`
	// RiskIndicators is the list of categories for the address
	addressRiskIndicator `gorm:"column:categories"`
}

// addressRiskIndicator is a risk indicator for an address
// it wraps the trmlabs.AddressRiskIndicator struct
type addressRiskIndicator struct {
	trmlabs.AddressRiskIndicator
}

// GormDataType returns the data type for the column
func (a addressRiskIndicator) GormDataType() string {
	return "json"
}

var _ schema.GormDataTypeInterface = addressRiskIndicator{}
