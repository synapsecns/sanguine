package graph

import (
	serverConfig "github.com/synapsecns/sanguine/services/sinner/config/server"
	"github.com/synapsecns/sanguine/services/sinner/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the root resolver.
//
//go:generate go run github.com/synapsecns/sanguine/services/sinner/graphql/contrib/client
type Resolver struct {
	DB     db.EventDB
	Config serverConfig.Config
}
