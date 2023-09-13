package dbcommon_test

import (
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"testing"
)

// ExplicitColumnTestModel tests explicit naming
// of the column as such. E.g.: column:origin_chain_id.
type ExplicitColumnTestModel struct {
	gorm.Model
	TestField uint64 `,json:"not_test_field" gorm:"column:test_field"`
}

// ExplicitColumnTestModel tests implicit naming
// of the column as such. E.g.: origin_chain_id.
type ImplicitColumnTestModel struct {
	gorm.Model
	TestField uint64 `,json:"not_test_field" gorm:"test_field"`
}

// MultiColumnModel tests multiple tags in gorm.
type MultiColumnModel struct {
	gorm.Model
	TestField uint64 `,json:"not_test_field" gorm:"column:test_field;uniqueIndex:idx_id"`
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

// TestModel is a simple test model.
type TestModel struct {
	gorm.Model
}

type TestTabelerModel struct {
	gorm.Model
}

var tabelerName = gofakeit.Word()

func (t TestTabelerModel) TableName() string {
	return tabelerName
}

var _ schema.Tabler = &TestTabelerModel{}

type TestTablerWithNamerModel struct {
	gorm.Model
}

func (t TestTablerWithNamerModel) TableName(namer schema.Namer) string {
	return fmt.Sprintf("hi_%s", gofakeit.Word())
}

var _ schema.TablerWithNamer = &TestTablerWithNamerModel{}

// getTestModels returns a list of test models to migrate.
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func getTestModels() (allModels []interface{}) {
	allModels = append(allModels,
		&TestModel{}, &TestTabelerModel{}, &TestTablerWithNamerModel{})
	return allModels
}

type modelNameTest struct {
	namingStrategy schema.NamingStrategy
	strategyType   string
}

func (s *DbSuite) TestGetModelName() {
	namingStrategies := []modelNameTest{
		// no naming strategy
		{
			strategyType: "none",
		},
		{
			namingStrategy: schema.NamingStrategy{
				TablePrefix: "prefix_",
			},
			strategyType: "prefix",
		},
		{
			namingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			strategyType: "singular",
		},
		{
			namingStrategy: schema.NamingStrategy{
				NoLowerCase: true,
			},
			strategyType: "no_lower",
		},
		{
			namingStrategy: schema.NamingStrategy{
				NameReplacer: testReplacer{},
			},
			strategyType: "replacer",
		},
	}

	for _, namingStrategy := range namingStrategies {
		// create the db
		tmpDb := fmt.Sprintf("%s/tmp.db", filet.TmpDir(s.T(), ""))

		gdb, err := gorm.Open(sqlite.Open(tmpDb), &gorm.Config{
			Logger:         dbcommon.GetGormLogger(testDBLogger),
			NamingStrategy: namingStrategy.namingStrategy,
		})

		s.Require().NoError(err)

		err = gdb.WithContext(s.GetTestContext()).AutoMigrate(getTestModels()...)
		s.Require().NoError(err)

		for _, model := range getTestModels() {
			// TODO: example for GetModelName
			modelName, err := dbcommon.GetModelName(gdb, model)
			s.Require().NoErrorf(err, "failed to get model name for %T %s", model, namingStrategy.strategyType)

			var count uint64
			tx := gdb.Raw(fmt.Sprintf("SELECT COUNT(*) FROM %s", modelName)).Scan(&count)
			s.Require().NoErrorf(tx.Error, "failed to query model for %s", namingStrategy.strategyType)
		}
	}
}

type testReplacer struct{}

func (t testReplacer) Replace(name string) string {
	return strings.NewReplacer("model", "mod").Replace(name)
}

// NonConsistentModelName tests a model with a non-consistent name.
type NonConsistentModelName struct {
	gorm.Model
	TestField uint64 `,json:"not_test_field" gorm:"column:test_field"`
}

type OtherNonConsistentModelName struct {
	gorm.Model
	TestField uint64 `,json:"not_test_field" gorm:"column:test_field_with_different_name"`
}

func TestNonConsistentName(t *testing.T) {
	namer := dbcommon.NewNamer(getNonConsistentModels())
	Panics(t, func() {
		namer.GetConsistentName("TestField")
	})
}

func TestConsistentName(t *testing.T) {
	namer := dbcommon.NewNamer(getTestModels())
	testFieldName := namer.GetConsistentName("CreatedAt")
	Equal(t, testFieldName, "created_at")
}

// GetAllModels gets all models to migrate
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func getNonConsistentModels() (allModels []interface{}) {
	allModels = append(allModels,
		&NonConsistentModelName{}, &OtherNonConsistentModelName{},
	)
	return allModels
}

func ExampleNamer_GetConsistentName() {
	namer := dbcommon.NewNamer(getTestModels())
	fmt.Println(namer.GetConsistentName("CreatedAt"))
	// output: created_at
}
