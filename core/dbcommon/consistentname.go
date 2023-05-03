package dbcommon

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"

	"github.com/fatih/structtag"
	"gorm.io/gorm/schema"
)

// getGormFieldName gets a gorm field name by name
// it panics if the field cannot be found or if the tag is not available
//
// note: this should only be called on init. See package docs for an explanation
// of this pattern. Ideally - all struct names map 1:1 with a field name consistently.
// please use caution before using this directly anywhere
//
// TODO: it'd be ideal if more of this could be done by the gorm schema parser.
func getGormFieldName(model interface{}, fieldName string) string {
	field, ok := reflect.TypeOf(model).Elem().FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("could not get field %s on struct %s. Does this field exist?", fieldName, reflect.TypeOf(model)))
	}

	tags, err := structtag.Parse(string(field.Tag))
	if err != nil {
		panic(fmt.Errorf("could not parse struct tag: %s, got error: %w", string(field.Tag), err))
	}

	gormTag, err := tags.Get("gorm")
	// no tag at all, return the field name
	if err != nil {
		return schema.NamingStrategy{}.ColumnName("", fieldName)
	}

	// create a empty schema, this is only used for setting errors in the parse field function
	// these errors are not exported, but currently will not happen here if migrations are successful
	// (a pre-requisite to running the node). See: https://git.io/JW6by for details
	emptySchema := schema.Schema{}
	parsedField := emptySchema.ParseField(field)

	// if the field (column like type) can be parsed, return it otherwise
	// the user is using gorm:COLUMN_NAME pattern and we can return the tag name.
	if parsedField.DBName == "" {
		return gormTag.Name
	}

	return parsedField.DBName
}

// Namer is used to pull consistent names for fields from models. This prevents inconsistency by panicing on breaking changes.
// It also allows us to avoid using xModelFieldName, yModelFieldName by taking advantage of introspection.
type Namer struct {
	models []interface{}
}

// NewNamer creates a new namer.
func NewNamer(models []interface{}) Namer {
	return Namer{models: models}
}

// GetConsistentName makes sure a `ColumnName` has a common column tag against all models
// we use this so we don't have to define bridgeRedeemModelFieldName, originChainFieldName, etc.
// panic's on an inconsistency. Note: this should only be called from init.
//
// after this is called, you can safely call getGormFieldName(AnyModelWithField, fieldName)
// if all fields cannot make this guarantee, they should be separated out to avoid developer confusion.
//
// this returns the result of getGormFieldName for a valid model. It also panics if no model is valid
// todo consider generating getters for all field names.
func (n Namer) GetConsistentName(fieldName string) string {
	var (
		// lastFoundModel is the interface where the model was last found
		// this is used for logging differences between two models
		lastFoundModel interface{}
		// lastTagName in the loop
		lastTagName string
		// whether or not a tag name has been found
		foundTagName bool
	)

	for _, model := range n.models {
		// check if the model has the field
		_, ok := reflect.TypeOf(model).Elem().FieldByName(fieldName)
		if !ok {
			continue
		}
		newTagName := getGormFieldName(model, fieldName)
		if !foundTagName {
			foundTagName = true
			lastTagName = newTagName
			lastFoundModel = model
			continue
		}

		if newTagName != lastTagName {
			panic(fmt.Sprintf("error asserting name consistency on field name %s. Model %s has tag name set "+
				"to %s, but model %s has tag name set to %s. See getConsistentName/package docs for details",
				fieldName, reflect.TypeOf(model).Name(), newTagName, reflect.TypeOf(lastFoundModel), lastTagName))
		}
		lastFoundModel = model
	}

	if !foundTagName {
		panic(fmt.Errorf("did not find any models with field name %s", fieldName))
	}

	return getGormFieldName(lastFoundModel, fieldName)
}

// GetModelName returns the name of the model.
func GetModelName(db *gorm.DB, model interface{}) (string, error) {
	if reflect.ValueOf(model).Kind() != reflect.Ptr {
		return "", fmt.Errorf("model must be a pointer, type is %T", model)
	}

	tableNameStmt := db.Model(model).Statement
	err := tableNameStmt.Parse(model)
	if err != nil {
		return "", fmt.Errorf("failed to parse model: %w", err)
	}

	return tableNameStmt.Schema.Table, nil
}
