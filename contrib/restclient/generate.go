package restclient

import (
	// this is used to generate the client code.
	_ "github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate types,client -o client.gen.go -package restclient ../../packages/rest-api/swagger.json
