// Package main generates graphql object code from contracts that are used to send messages by users
// contracts must be added here in order to eb generated
package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
	"github.com/synapsecns/sanguine/services/explorer/contracts/user"
	"github.com/synapsecns/sanguine/services/explorer/graphql/contrib/model/internal"
)

func main() {
	messageFormats := user.GetMessageFormats()
	schemaConfig := graphql.SchemaConfig{}
	for name, model := range messageFormats {
		gqlType, err := internal.GetGraphQLObject(model.DataType)
		if err != nil {
			panic(fmt.Errorf("could not generate graphql object for %s (type %T): %v", name, model, err))
		}

		schemaConfig.Types = append(schemaConfig.Types, gqlType)
		// required for intorpsection
		schemaConfig.Query = graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				name: &graphql.Field{
					Type: gqlType,
				},
			},
		})
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(fmt.Errorf("could not generate graphql schema: %v", err))
	}
	res := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: testutil.IntrospectionQuery,
	})

	if res.HasErrors() {
		panic(res.Errors)
	}

	fmt.Println(res.Data)
}
