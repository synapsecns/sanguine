package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	bridgeContract "github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	cctpContract "github.com/synapsecns/sanguine/services/explorer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"math/big"
)

type ServerParsers struct {
	BridgeParsers map[uint32]*parser.BridgeParser
	CCTParsers    map[uint32]*parser.CCTPParser
}

type ServerRefs struct {
	BridgeRefs map[uint32]*bridgeContract.BridgeRef
	CCTPRefs   map[uint32]*cctpContract.CCTPRef
}

type IFaceBridgeEvent struct {
	IFace       bridge.EventLog
	BridgeEvent *sql.BridgeEvent
}

type SwapReplacementData struct {
	Address common.Address
	Amount  *big.Int
}
