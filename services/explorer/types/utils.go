// Package types hold supplementary types for the explorer service.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher/tokenprice"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokendata"
	bridgeContract "github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	cctpContract "github.com/synapsecns/sanguine/services/explorer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/explorer/db"
)

// ServerParsers is a custom type for holding parsers for the server.
type ServerParsers struct {
	BridgeParsers map[uint32]*parser.BridgeParser
	CCTParsers    map[uint32]*parser.CCTPParser
}

// ServerRefs is a custom type for holding refs for the server.
type ServerRefs struct {
	BridgeRefs map[uint32]*bridgeContract.BridgeRef
	CCTPRefs   map[uint32]*cctpContract.CCTPRef
}

// ParserConfig is a custom type for initializing parser.
type ParserConfig struct {
	// ContractAddress is the address of the contract.
	ContractAddress common.Address
	// ConsumerDB is the database to store parsed data in.
	ConsumerDB db.ConsumerDB
	// CoinGeckoIDs is the mapping of token id to coin gecko ID
	CoinGeckoIDs map[string]string
	// ConsumerFetcher is the ScribeFetcher for sender and timestamp.
	ConsumerFetcher fetcher.ScribeFetcher
	// TokenDataService contains the token data service/cache
	TokenDataService tokendata.Service
	// TokenPriceService contains the token price service/cache
	TokenPriceService tokenprice.Service
	// FromAPI is true if the parser is being called from the API.
	FromAPI bool
}
