// Package db provides the database interface for the screener-api.
package db

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"gorm.io/gorm/schema"
)

// BlacklistedAddress is a blacklisted address.
type BlacklistedAddress struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	Id      string `gorm:"column:id;primary_key"`
	TypeReq string `gorm:"column:typereq"`
	Data    string `gorm:"column:data"`
	Address string `gorm:"column:address"`
	Network string `gorm:"column:network"`
	Tag     string `gorm:"column:tag"`
	Remark  string `gorm:"column:remark"`
}

func (b BlacklistedAddress) GormDataType() string {
	return "json"
}

func (b BlacklistedAddress) Value() (driver.Value, error) {
	if b == (BlacklistedAddress{}) {
		return nil, nil
	}
	ba, err := json.Marshal(b)
	return string(ba), err
}

func (b *BlacklistedAddress) Scan(val interface{}) error {
	if val == nil {
		*b = BlacklistedAddress{}
		return nil
	}
	var ba []byte
	switch v := val.(type) {
	case []byte:
		ba = v
	case string:
		ba = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}
	rd := bytes.NewReader(ba)
	decoder := json.NewDecoder(rd)
	decoder.UseNumber()
	err := decoder.Decode(b)
	//nolint: wrapcheck
	return err
}

// AddressIndicators is the address indicators for a given address.
type AddressIndicators struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	// Address is the address to be screened
	Address string `gorm:"column:address;primary_key"`
	// RiskIndicators is the list of categories for the address
	Indicators addressRiskIndicators `gorm:"column:indicators"`
}

// addressRiskIndicator is a risk indicator for an address
// it wraps the trmlabs.AddressRiskIndicator struct.
type addressRiskIndicators []trmlabs.AddressRiskIndicator

// GormDataType returns the data type for the column.
func (a addressRiskIndicators) GormDataType() string {
	return "json"
}

// Value return json value, implement driver.Valuer interface.
func (a addressRiskIndicators) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	ba, err := json.Marshal(a)
	return string(ba), err
}

// Scan scan value into Jsonb, implements sql.Scanner interface.
func (a *addressRiskIndicators) Scan(val interface{}) error {
	if val == nil {
		*a = make(addressRiskIndicators, 0)
		return nil
	}
	var ba []byte
	switch v := val.(type) {
	case []byte:
		ba = v
	case string:
		ba = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}
	t := addressRiskIndicators{}
	rd := bytes.NewReader(ba)
	decoder := json.NewDecoder(rd)
	decoder.UseNumber()
	err := decoder.Decode(&t)
	*a = t
	//nolint: wrapcheck
	return err
}

// ToTRMLabs converts the address risk indicators to trmlabs.AddressRiskIndicator.
func (a addressRiskIndicators) ToTRMLabs() (res []trmlabs.AddressRiskIndicator) {
	if a == nil {
		return nil
	}

	res = make([]trmlabs.AddressRiskIndicator, len(a))
	copy(res, a)

	return res
}

// MakeRecord creates a new address indicators record.
func MakeRecord(address string, records []trmlabs.AddressRiskIndicator) *AddressIndicators {
	indicators := make(addressRiskIndicators, len(records))
	copy(indicators, records)

	return &AddressIndicators{
		Address:    strings.ToLower(address),
		Indicators: indicators,
	}
}

var _ schema.GormDataTypeInterface = addressRiskIndicators{}
var _ driver.Value = addressRiskIndicators{}
var _ sql.Scanner = &addressRiskIndicators{}

var _ schema.GormDataTypeInterface = BlacklistedAddress{}
var _ driver.Value = BlacklistedAddress{}
var _ sql.Scanner = &BlacklistedAddress{}
