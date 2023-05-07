package dbcommon_test

import (
	"context"
	"database/sql/driver"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/ipfs/go-log"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"testing"
)

var testDBLogger = log.Logger("dbcommon")

// TestEnum tests the default providers for an enum type.
func (s DbSuite) TestEnum() {
	res, err := RunEnumExample(s.GetTestContext(), fmt.Sprintf("%s/sql.db", filet.TmpDir(s.T(), "")))
	Nil(s.T(), err)

	for i, fruit := range AllFruits {
		Equal(s.T(), fruit.Int(), res[i].Fruit.Int())
	}

	Equal(s.T(), len(res), len(AllFruits))
}

// ExampleEnum demonstrates example use of the enum interface.
// this implementation can be confusing, so there's an example below.
func ExampleEnum() {
	res, err := RunEnumExample(context.Background(), fmt.Sprintf("%s/sql.db", os.TempDir()))
	if err != nil {
		panic(err)
	}

	for _, res := range res {
		fmt.Printf("got result %s \n", res.Fruit.String())
	}
}

// RunEnumExample is used to separate out tests from the example.
func RunEnumExample(ctx context.Context, dbDir string) (res []InventoryModel, err error) {
	gdb, err := gorm.Open(sqlite.Open(dbDir), &gorm.Config{
		Logger: dbcommon.GetGormLogger(testDBLogger),
	})
	if err != nil {
		return res, fmt.Errorf("could not open db: %w", err)
	}

	// migrate the inventory model
	err = gdb.WithContext(ctx).AutoMigrate(&InventoryModel{})
	if err != nil {
		return res, fmt.Errorf("could not migrate db: %w", err)
	}

	for _, fruit := range AllFruits {
		tx := gdb.WithContext(ctx).Create(&InventoryModel{
			Fruit: fruit,
		})

		if tx.Error != nil {
			return res, fmt.Errorf("could not insert fruit: %w", err)
		}
	}

	tx := gdb.WithContext(ctx).Find(&res)
	if tx.Error != nil {
		return res, fmt.Errorf("could not query db: %w", err)
	}

	return res, nil
}

// InventoryModel is an example model for of an inventory table for fruit.
type InventoryModel struct {
	gorm.Model
	// fruit is the fruit we're storing
	Fruit Fruit
}

// you should use ints rather than iota's when interacting with the database.
const (
	// Apple is an example implementing enum.
	Apple Fruit = 0
	// Pear is a n example implementing enum.
	Pear Fruit = 1
)

var AllFruits = []Fruit{Apple, Pear}

type Fruit uint8

// String gets a string of the enum
// in a production setting, generater should be used.
// see: https://pkg.go.dev/golang.org/x/tools/cmd/stringer for details
func (f Fruit) String() string {
	switch f {
	case Apple:
		return "Apple"
	case Pear:
		return "Pear"
	}
	return ""
}

// Int get the integer value of the fruit.
func (f Fruit) Int() uint8 {
	return uint8(f)
}

// GormDataType is the gorm data type.
func (f Fruit) GormDataType() string {
	return dbcommon.EnumDataType
}

// Scan will scan the fruit into the db.
func (f *Fruit) Scan(src interface{}) error {
	res, err := dbcommon.EnumScan(src)
	if err != nil {
		return fmt.Errorf("could not scan: %w", err)
	}
	newFruit := Fruit(res)
	*f = newFruit
	return nil
}

// nolint: wrapcheck
func (f *Fruit) Value() (driver.Value, error) {
	return dbcommon.EnumValue(f)
}

var _ dbcommon.EnumInter = (*Fruit)(nil)

type testEnum uint8

func (t testEnum) Int() uint8 {
	return uint8(t)
}

const (
	testEnumValue1 testEnum = 1
	testEnumValue2 testEnum = 2
)

func TestEnumValue(t *testing.T) {
	tests := []struct {
		name    string
		enum    dbcommon.EnumInter
		want    int64
		wantErr error
	}{
		{
			name: "Valid enum value",
			enum: testEnumValue1,
			want: 1,
		},
		{
			name: "Valid enum value",
			enum: testEnumValue2,
			want: 2,
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbcommon.EnumValue(tt.enum)
			if tt.wantErr != nil {
				ErrorIs(t, err, tt.wantErr)
			} else {
				Nil(t, err)
				Equal(t, tt.want, got)
			}
		})
	}
}

func TestEnumScan(t *testing.T) {
	tests := []struct {
		name    string
		src     interface{}
		want    uint8
		wantErr string
	}{
		{
			name: "Valid int64 value",
			src:  int64(1),
			want: 1,
		},
		{
			name: "Valid int32 value",
			src:  int32(2),
			want: 2,
		},
		{
			name:    "Invalid type",
			src:     "invalid",
			want:    0,
			wantErr: "could not scan enum: converting driver.Value type string (\"invalid\") to a int32: invalid syntax",
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbcommon.EnumScan(tt.src)
			if tt.wantErr != "" {
				Error(t, err)
				EqualError(t, err, tt.wantErr)
			} else {
				NoError(t, err)
				Equal(t, tt.want, got)
			}
		})
	}
}
