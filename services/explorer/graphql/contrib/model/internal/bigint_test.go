package internal

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/graphql-go/graphql"
)

func buildUseableSchema(inputType graphql.Input) graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"test_bigint": &graphql.Field{
					Name:        "test_bigint",
					Description: "test bigint correctly resolving stuff",
					Type:        BigInt,
					Args: graphql.FieldConfigArgument{
						"test_param": &graphql.ArgumentConfig{Type: inputType},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						arg := p.Args["test_param"]

						serialized := BigInt.Serialize(arg)

						return serialized, nil
					},
				},
			},
		}),
	})
	if err != nil {
		log.Fatalf("Got an error(s) creating the schema: %v", err)
	}

	return schema
}

func doQuery(schema graphql.Schema, query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		Context:       context.Background(),
	})

	return result
}

func unmarshalResult(result interface{}) TestResult {
	var testResult TestResult

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("Failed to marshal the result. Womp. Error: %v", err)
	}
	if err := json.Unmarshal(data, &testResult); err != nil {
		log.Fatalf("Failed to unmarshal the result. Womp. Error: %v", err)
	}

	return testResult
}

type TestResult struct {
	TestBigInt int64 `json:"test_bigint"`
}

func TestBoolTrue(t *testing.T) {
	const expected int64 = 1

	schema := buildUseableSchema(graphql.Boolean)
	query := `{ test_bigint(test_param: true) }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestBoolFalse(t *testing.T) {
	const expected int64 = 0

	schema := buildUseableSchema(graphql.Boolean)
	query := `{ test_bigint(test_param: false) }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestPositiveInt(t *testing.T) {
	const expected int64 = 42

	schema := buildUseableSchema(graphql.Int)
	query := `{ test_bigint(test_param: 42) }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestNegativeInt(t *testing.T) {
	const expected int64 = -42

	schema := buildUseableSchema(graphql.Int)
	query := `{ test_bigint(test_param: -42) }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestBigIntPositive(t *testing.T) {
	const expected int64 = 42000000000000

	schema := buildUseableSchema(BigInt)
	query := `{ test_bigint(test_param: 42000000000000) }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestBigIntNegative(t *testing.T) {
	const expected int64 = -42000000000000

	schema := buildUseableSchema(BigInt)
	query := `{ test_bigint(test_param: -42000000000000) }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestFloatPositiveRoundDown(t *testing.T) {
	const expected int64 = 42

	schema := buildUseableSchema(graphql.Float)
	query := `{ test_bigint(test_param: 42.3) }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestFloatPositiveRoundUp(t *testing.T) {
	const expected int64 = 43

	schema := buildUseableSchema(graphql.Float)
	query := `{ test_bigint(test_param: 42.7) }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestFloatNegativeRoundDown(t *testing.T) {
	const expected int64 = -42

	schema := buildUseableSchema(graphql.Float)
	query := `{ test_bigint(test_param: -42.3) }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestFloatNegativeRoundUp(t *testing.T) {
	const expected int64 = -43

	schema := buildUseableSchema(graphql.Float)
	query := `{ test_bigint(test_param: -42.7) }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestStringIntPositive(t *testing.T) {
	const expected int64 = 42

	schema := buildUseableSchema(graphql.String)
	query := `{ test_bigint(test_param: "42") }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestStringIntNegative(t *testing.T) {
	const expected int64 = -42

	schema := buildUseableSchema(graphql.String)
	query := `{ test_bigint(test_param: "-42") }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestStringFloatPositive(t *testing.T) {
	const expected int64 = 42

	schema := buildUseableSchema(graphql.String)
	query := `{ test_bigint(test_param: "42.3") }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestStringFloatNegative(t *testing.T) {
	const expected int64 = -42

	schema := buildUseableSchema(graphql.String)
	query := `{ test_bigint(test_param: "-42.3") }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestStringFloatNegativeRoundDown(t *testing.T) {
	const expected int64 = -42

	schema := buildUseableSchema(graphql.String)
	query := `{ test_bigint(test_param: "-42.3") }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestStringFloatNegativeRoundUp(t *testing.T) {
	const expected int64 = -43

	schema := buildUseableSchema(graphql.String)
	query := `{ test_bigint(test_param: "-42.7") }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestStringFloatPositiveRoundDown(t *testing.T) {
	const expected int64 = 42

	schema := buildUseableSchema(graphql.String)
	query := `{ test_bigint(test_param: "42.3") }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}

func TestStringFloatPositiveRoundUp(t *testing.T) {
	const expected int64 = 43

	schema := buildUseableSchema(graphql.String)
	query := `{ test_bigint(test_param: "42.7") }`
	result := doQuery(schema, query)

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}

	testResult := unmarshalResult(result.Data)

	if testResult.TestBigInt != expected {
		t.Logf("Failed result: %d", testResult.TestBigInt)
		t.Fail()
	}
}
