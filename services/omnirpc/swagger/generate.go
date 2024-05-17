// Package swagger registers the swagger generated spec.
package swagger

import _ "embed" // embed is required for go:embed

//go:generate go run github.com/synapsecns/sanguine/services/omnirpc/swagger/generator
//go:generate ./gen.sh

// OpenAPI is the openapi specification for the omnirpc service.
//
//go:embed openapi.yaml
var OpenAPI []byte
