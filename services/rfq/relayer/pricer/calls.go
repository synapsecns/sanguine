package pricer

import (
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/util"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
)

type callType int

const (
	// claimCallType is the call type for claim.
	claimCallType callType = iota + 1
	// proveCallType is the call type for prove.
	proveCallType
	// relayCallType is the call type for relay.
	relayCallType
)

func (c callType) String() string {
	switch c {
	case claimCallType:
		return "claim"
	case proveCallType:
		return "prove"
	case relayCallType:
		return "relay"
	}
	return ""
}

func getCall(transactor *bind.TransactOpts, bridge *fastbridge.FastBridgeRef, cType callType) (call *ethereum.CallMsg, err error) {
	var tx *types.Transaction
	switch cType {
	case claimCallType:
		tx, err = bridge.Claim(transactor, []byte{}, common.HexToAddress(""))
	case proveCallType:
		tx, err = bridge.Prove(transactor, []byte{}, [32]byte{})
	case relayCallType:
		tx, err = bridge.Relay(transactor, []byte{})
	default:
		return nil, fmt.Errorf("unknown call type: %d", cType)
	}
	if err != nil {
		return nil, fmt.Errorf("could not get tx with type %s: %w", cType.String(), err)
	}
	call, err = util.TxToCall(tx)
	if err != nil {
		return nil, fmt.Errorf("could not get call: %w", err)
	}
	if call == nil {
		return nil, fmt.Errorf("call is nil")
	}
	return call, nil
}
