package graph

import (
	"github.com/synapsecns/sanguine/services/explorer/consumer"
	"github.com/synapsecns/sanguine/services/explorer/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the root resolver.
//
//go:generate go run github.com/synapsecns/sanguine/services/explorer/graphql/contrib
type Resolver struct {
	DB      db.ConsumerDB
	Fetcher consumer.Fetcher
}
