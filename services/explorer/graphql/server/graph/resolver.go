package graph

import (
	etherClient "github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/services/explorer/api/cache"
	serverConfig "github.com/synapsecns/sanguine/services/explorer/config/server"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/scribe"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/types"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the root resolver.
//
//go:generate go run github.com/synapsecns/sanguine/services/explorer/graphql/contrib/client
type Resolver struct {
	DB          db.ConsumerDB
	Fetcher     scribe.IScribeFetcher
	Cache       cache.Service
	Clients     map[uint32]etherClient.EVM
	Parsers     *types.ServerParsers
	Refs        *types.ServerRefs
	SwapFilters map[string]*swap.SwapFlashLoanFilterer
	Config      serverConfig.Config
}
