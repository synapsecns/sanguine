package graph

import "github.com/synapsecns/sanguine/services/scribe/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the root resolver.
//
//go:generate go run github.com/synapsecns/sanguine/services/scribe/graphql/server/contrib
type Resolver struct {
	DB db.EventDB
}
