// Package internal is a modified version of https://github.com/ThundR67/straf
// build to support the generation of graphql objects from structs
// the original package is modified to support types used by etheruem
package internal

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"reflect"
	"time"

	"github.com/graphql-go/graphql"
)

// GetGraphQLObject Converts struct into graphql object
func GetGraphQLObject(object interface{}) (*graphql.Object, error) {
	objectType := reflect.TypeOf(object)
	fields, err := convertStruct(objectType)

	output := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   objectType.Name(),
			Fields: fields,
		},
	)

	if err != nil {
		err = fmt.Errorf("Error While Converting Struct To GraphQL Object: %v", err)
		return &graphql.Object{}, err
	}

	return output, nil
}

// convertStructToObject converts simple struct to graphql object
func convertStructToObject(
	objectType reflect.Type) (*graphql.Object, error) {

	fields, err := convertStruct(objectType)
	if err != nil {
		err = fmt.Errorf(
			"Error while converting type %v to graphql fields: %v",
			objectType,
			err,
		)
		return &graphql.Object{}, err
	}

	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:   objectType.Name(),
			Fields: fields,
		},
	), nil
}

// convertStruct converts struct to graphql fields
func convertStruct(objectType reflect.Type) (graphql.Fields, error) {
	fields := graphql.Fields{}

	for i := 0; i < objectType.NumField(); i++ {
		currentField := objectType.Field(i)
		fieldType, err := getFieldType(currentField)
		if err != nil {
			err = fmt.Errorf(
				"Error while converting type %v to graphQL Type: %v",
				currentField.Type,
				err,
			)
			return graphql.Fields{}, err
		}

		fields[currentField.Name] = &graphql.Field{
			Name:              currentField.Name,
			Type:              fieldType,
			DeprecationReason: getTagValue(currentField, "deprecationReason"),
			Description:       getTagValue(currentField, "description"),
		}
	}

	return fields, nil
}

// getFieldType Converts object to a graphQL field type
func getFieldType(object reflect.StructField) (graphql.Output, error) {

	isID, ok := object.Tag.Lookup("unique")
	if isID == "true" && ok {
		return graphql.ID, nil
	}

	objectType := object.Type
	if objectType.Kind() == reflect.Struct {
		return convertStructToObject(objectType)

	} else if objectType.Kind() == reflect.Slice &&
		objectType.Elem().Kind() == reflect.Struct {

		elemType, err := convertStructToObject(objectType.Elem())
		return graphql.NewList(elemType), err

	} else if objectType.Kind() == reflect.Slice {
		elemType, err := convertSimpleType(objectType.Elem())
		return graphql.NewList(elemType), err
	}

	return convertSimpleType(objectType)
}

// convertSimpleType converts simple type  to graphql field
func convertSimpleType(objectType reflect.Type) (*graphql.Scalar, error) {

	typeMap := map[reflect.Kind]*graphql.Scalar{
		reflect.String:                          graphql.String,
		reflect.Bool:                            graphql.Boolean,
		reflect.Int:                             graphql.Int,
		reflect.Uint:                            graphql.Int,
		reflect.Int8:                            graphql.Int,
		reflect.Uint8:                           graphql.Int,
		reflect.Int16:                           graphql.Int,
		reflect.Uint16:                          graphql.Int,
		reflect.Int32:                           graphql.Int,
		reflect.Uint32:                          graphql.Int,
		reflect.Int64:                           graphql.Int,
		reflect.Uint64:                          graphql.Int,
		reflect.Float32:                         graphql.Float,
		reflect.Float64:                         graphql.Float,
		reflect.TypeOf(time.Time{}).Kind():      graphql.DateTime,
		reflect.TypeOf(&big.Int{}).Kind():       BigInt,
		reflect.TypeOf(common.Address{}).Kind(): graphql.String,
	}

	graphqlType, ok := typeMap[objectType.Kind()]

	if !ok {
		return &graphql.Scalar{}, fmt.Errorf("Invalid Type %s", objectType.Kind().String())
	}

	return graphqlType, nil
}

// getTagValue returns tag value of a struct
func getTagValue(objectType reflect.StructField, tagName string) string {
	tag := objectType.Tag
	value, ok := tag.Lookup(tagName)
	if !ok {
		return ""
	}
	return value
}
