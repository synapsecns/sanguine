package dbcommon_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"gorm.io/gorm"
	"testing"
)

// ExplicitColumnTestModel tests explicit naming
// of the column as such. E.g.: column:origin_chain_id.
type ExplicitColumnTestModel struct {
	gorm.Model
	TestField uint64 `gorm:"column:test_field" ,json:"not_test_field"`
}

// ExplicitColumnTestModel tests implicit naming
// of the column as such. E.g.: origin_chain_id.
type ImplicitColumnTestModel struct {
	gorm.Model
	TestField uint64 `gorm:"test_field" ,json:"not_test_field"`
}

// MultiColumnModel tests multiple tags in gorm.
type MultiColumnModel struct {
	gorm.Model
	TestField uint64 `gorm:"column:test_field;uniqueIndex:idx_id" ,json:"not_test_field"`
}

// TestGetGormFieldName tests getting the gorm field by name.
func TestGetGormFieldName(t *testing.T) {
	fieldName := dbcommon.GetGormFieldName(&ExplicitColumnTestModel{}, "TestField")
	Equal(t, fieldName, "test_field")

	// test implicit naming
	fieldName = dbcommon.GetGormFieldName(&ImplicitColumnTestModel{}, "TestField")
	Equal(t, fieldName, "test_field")

	fieldName = dbcommon.GetGormFieldName(&MultiColumnModel{}, "TestField")
	Equal(t, fieldName, "test_field")
}
