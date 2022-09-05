package graph

import "github.com/synapsecns/sanguine/services/scribe/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// // go:generate go run github.com/synapsecns/sanguine/services/scribe/server/contrib
//
//go:generate go run github.com/99designs/gqlgen generate
type Resolver struct {
	DB db.EventDB
}
