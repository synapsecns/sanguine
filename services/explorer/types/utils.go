// Package types hold supplementary types for the explorer service.
package types

import (
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	bridgeContract "github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	cctpContract "github.com/synapsecns/sanguine/services/explorer/contracts/cctp"
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
