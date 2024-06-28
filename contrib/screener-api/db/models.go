// Package db provides the database interface for the screener-api.
package db

import (
	"time"
)

// BlacklistedAddress is a blacklisted address.
type BlacklistedAddress struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	ID      string `gorm:"column:id;primaryKey" json:"id"`
	Type    string `gorm:"column:type"          json:"type"`
	Address string `gorm:"column:address"       json:"address"`
	Network string `gorm:"column:network"       json:"network"`
	Tag     string `gorm:"column:tag"           json:"tag"`
	Remark  string `gorm:"column:remark"        json:"remark"`
}
